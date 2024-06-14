package tree

import (
	"strings"

	"github.com/sirupsen/logrus"
)

/*
 attention:
			!!!!!!!!!!!!!!!!!!!!!
			this file is diff: (head vs base) diff ( (head vs org) && (base vs org) )
			!!!!!!!!!!!!!!!!!!!!!
*/

// StringDataFlowItem string items
type StringDataFlowItem struct {
	ListOrg, ListBase, ListHead          []*Diff3EmitItem
	HeadDataMap, BaseDataMap, OrgDataMap map[string]*Diff3EmitItem
	Org2HeadIndexMap, Org2BaseIndexMap   map[int]int
	Head2OrgMap, Base2OrgMap             map[string]*Diff3EmitItem // to mapping base/head vs org diff
	BaseExistMap, HeadExistMap           map[string]*Diff3EmitItem
}

// Diff3EmitItem output item
type Diff3EmitItem struct {
	// Required:true
	ID string `json:"id"`
	// Required:true
	Content string `json:"content"`
	// Required:true
	DiffType uint8 `json:"diff_type"`
	// Required:true
	HeadInfo string `json:"head_info"`
}

// DiffFileType represents the type of a DiffFile.
type DiffFileType uint8

// DiffFileType possible values.
const (
	DiffFileAdd DiffFileType = iota + 1
	DiffFileChange
	DiffFileDel
	DiffFileRename
	DiffFileCopy
	DiffFileMove
	DiffFileMoveIn
	DiffFileMoveOut
	DiffFileMoveInAndChange
	DiffFileMoveOutAndChange
	DiffFileMoveInAndRename
	DiffFileMoveOutAndRename
)

func (d *StringDataFlowItem) Diff3Core() []*Diff3EmitItem {
	return d.generateTrunc(d.Org2BaseIndexMap, d.Org2HeadIndexMap)
}

// GenerateTrunc generate trunc
func (d *StringDataFlowItem) generateTrunc(mapBase, mapHead map[int]int) []*Diff3EmitItem {
	var (
		ans                          = make([]*Diff3EmitItem, 0)
		matchO, matchBase, matchHead = -1, -1, -1 // to record previous matched index
		travers                      func(start int)
		emitTrunc                    func(startA, startHead, endBase, endHead int)
	)

	emitTrunc = func(startBase, startHead, endBase, endHead int) {
		if startBase < endBase || startHead < endHead {
			var (
				baseList = make([]*Diff3EmitItem, 0)
				headList = make([]*Diff3EmitItem, 0)
			)
			for i := startBase; i < endBase; i++ {
				baseList = append(baseList, d.ListBase[i])
			}
			for i := startHead; i < endHead; i++ {
				headList = append(headList, d.ListHead[i])
			}
			ans = append(ans, d.getTruncDiff(headList, baseList)...)
		}
	}

	travers = func(start int) {
		var o, headIdx, baseIdx = findNextMatch(mapBase, mapHead, start, len(d.ListOrg))
		if o == -1 { // not matched, emit truncs
			emitTrunc(matchBase+1, matchHead+1, len(d.ListBase), len(d.ListHead))
			return
		}
		// matched
		if baseIdx-1 > matchBase || headIdx-1 > matchHead {
			emitTrunc(matchBase+1, matchHead+1, baseIdx, headIdx)
		}

		ans = append(ans, d.dealChangeEvent(d.ListOrg[o].ID)...)

		// update previous matched index
		matchO, matchBase, matchHead = o, baseIdx, headIdx
		travers(matchO + 1)
	}

	// begin from idx=0
	travers(0)
	return ans
}

func (d *StringDataFlowItem) cmpHeadAndBaseObj(id string) []*Diff3EmitItem {
	return d.cmpTwoObj(d.HeadDataMap[id], d.BaseDataMap[id])
}

func (d *StringDataFlowItem) cmpHeadAndOrgObj(id string) []*Diff3EmitItem {
	return d.cmpTwoObj(d.OrgDataMap[id], d.HeadDataMap[id])
}

func (d *StringDataFlowItem) cmpBaseAndOrgObj(id string) []*Diff3EmitItem {
	return d.cmpTwoObj(d.OrgDataMap[id], d.BaseDataMap[id])
}

// cmpTwoObj to diff two object
func (d *StringDataFlowItem) cmpTwoObj(head, base *Diff3EmitItem) []*Diff3EmitItem {
	if head == nil || base == nil {
		return nil
	}

	var item = *base
	if head.Content != base.Content {
		item.DiffType = uint8(DiffFileChange)
		item.HeadInfo = head.Content
	}
	return []*Diff3EmitItem{&item}
}

// findNextMatch if matches, return o, A, B
func findNextMatch(mapBase, mapHead map[int]int, start, total int) (int, int, int) {
	for o := start; o < total; o++ {
		idxHead, okHead := mapHead[o]
		idxBase, okBase := mapBase[o]
		if okBase && okHead {
			return o, idxHead, idxBase
		}
	}
	return -1, -1, -1
}

func GenMatchMap(base, head []string) map[int]int {
	var m = make(map[int]int, 0)
	outPut := GenerateDiff(base, head)
	var headIdx, baseIdx = 0, 0
	for _, item := range outPut {
		item = strings.TrimSpace(item)
		if strings.HasPrefix(item, "+") {
			headIdx++
		} else if strings.HasPrefix(item, "-") {
			baseIdx++
		} else {
			m[baseIdx] = headIdx
			baseIdx++
			headIdx++
		}
	}
	return m
}

// getTruncDiff maybe there are some same items in trunc
func (d *StringDataFlowItem) getTruncDiff(headList, baseList []*Diff3EmitItem) []*Diff3EmitItem {
	var (
		ans        = make([]*Diff3EmitItem, 0)
		headIDs, _ = prepareList(headList)
		baseIDs, _ = prepareList(baseList)
		output     = GenerateDiff(baseIDs, headIDs)
	)

	for _, v := range output {
		if strings.HasPrefix(v, "+") {
			v = strings.TrimSpace(strings.TrimPrefix(v, "+"))
			ans = append(ans, d.dealHeadAddEvent(v)...)

		} else if strings.HasPrefix(v, "-") {
			v = strings.TrimSpace(strings.TrimPrefix(v, "-"))
			ans = append(ans, d.dealBaseDelEvent(v)...)

		} else {
			v = strings.TrimSpace(v)
			ans = append(ans, d.dealChangeEvent(v)...)
		}
	}
	return ans
}

// dealHeadAddEvent to handle the add head element of diff output
func (d *StringDataFlowItem) dealHeadAddEvent(id string) (ans []*Diff3EmitItem) {

	item, ok := d.HeadExistMap[id]
	if ok { // has been handled
		if item.DiffType > 0 { // need add
			ans = append(ans, item)
		}
		delete(d.BaseExistMap, id)
		return
	}

	// not exist in baseExistMap: has not handled
	var obj = &Diff3EmitItem{ID: id, Content: d.HeadDataMap[id].Content}

	if _, ok := d.OrgDataMap[id]; ok { // exist in org

		if _, ok := d.BaseDataMap[id]; ok { // exist in head

			// base is changed or not diffed with org
			// this block is similar with del base, so use Head2OrgMap
			diffHeadItem := d.Head2OrgMap[id]

			switch DiffFileType(diffHeadItem.DiffType) {
			case 0: // head not changed: base move or move&&modify, so add base is ok
				d.BaseExistMap[id] = &Diff3EmitItem{
					ID:       id,
					Content:  d.BaseDataMap[id].Content,
					DiffType: 100,
					HeadInfo: "",
				}

			case DiffFileChange: // head modify: base moved

				list := d.cmpHeadAndBaseObj(id)
				if list[0].DiffType == 0 { // base moved but has the same with head: add base is ok

					d.BaseExistMap[id] = &Diff3EmitItem{
						ID: id, Content: d.BaseDataMap[id].Content,
						DiffType: 100,
					}

				} else { // to modify base: so the result is only modify
					d.BaseExistMap[id] = &Diff3EmitItem{
						ID: id, Content: d.BaseDataMap[id].Content,
						DiffType: uint8(DiffFileChange),
						HeadInfo: d.HeadDataMap[id].Content,
					}
				}

			case DiffFileAdd, DiffFileDel: // head moved

				diffBaseItem := d.Base2OrgMap[id]
				switch DiffFileType(diffBaseItem.DiffType) {
				case 0: // base no change
					list := d.cmpHeadAndOrgObj(id) // or: head vs base
					if list[0].DiffType == 0 {     // only move
						obj.DiffType = uint8(DiffFileMoveIn)
						d.BaseExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.BaseDataMap[id].Content,
							DiffType: uint8(DiffFileMoveOut),
						}
					} else { // modify
						obj.DiffType = uint8(DiffFileMoveInAndChange)
						d.BaseExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.BaseDataMap[id].Content,
							DiffType: uint8(DiffFileMoveOutAndChange),
						}
					}
					ans = append(ans, obj)
					return

				case DiffFileChange: // base modify

					list := d.cmpHeadAndBaseObj(id)
					if list[0].DiffType == 0 {
						obj.DiffType = uint8(DiffFileMoveIn)
						d.BaseExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.BaseDataMap[id].Content,
							DiffType: uint8(DiffFileMoveOut),
						}
					} else {
						obj.DiffType = uint8(DiffFileMoveInAndChange)
						d.BaseExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.BaseDataMap[id].Content,
							DiffType: uint8(DiffFileMoveOutAndChange),
						}
					}
					ans = append(ans, obj)

				case DiffFileAdd, DiffFileDel: // base also moved
					// if head modify

					list := d.cmpHeadAndBaseObj(id)
					if list[0].DiffType == 0 {
						obj.DiffType = uint8(DiffFileMoveIn)
						d.BaseExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.BaseDataMap[id].Content,
							DiffType: uint8(DiffFileMoveOut),
						}

					} else {

						list2 := d.cmpHeadAndOrgObj(id)
						if list2[0].DiffType == 0 { // base modify
							obj.DiffType = uint8(DiffFileMoveIn)
							d.BaseExistMap[id] = &Diff3EmitItem{
								ID:       id,
								Content:  d.BaseDataMap[id].Content,
								DiffType: uint8(DiffFileMoveOut),
							}
						} else { // head modify
							obj.DiffType = uint8(DiffFileMoveInAndChange)
							d.BaseExistMap[id] = &Diff3EmitItem{
								ID:       id,
								Content:  d.BaseDataMap[id].Content,
								DiffType: uint8(DiffFileMoveOutAndChange),
							}
						}

					}

					ans = append(ans, obj)
					return
				}

			case DiffFileRename:
				// TODO: for file, it's a special case to compare...

			}

		}

	} else { // not exist in org: head added

		// if exist in base
		if _, ok := d.BaseDataMap[id]; ok { // exist in base: has updated?

			list := d.cmpHeadAndBaseObj(id)
			if list[0].DiffType == 0 { // no change: only move
				obj.DiffType = uint8(DiffFileMoveIn)
				d.BaseExistMap[id] = &Diff3EmitItem{
					ID:       id,
					Content:  d.BaseDataMap[id].Content,
					DiffType: uint8(DiffFileMoveOut),
					HeadInfo: "",
				}
			} else { // has changed: modify && moveIn
				obj.DiffType = uint8(DiffFileMoveInAndChange)
				d.BaseExistMap[id] = &Diff3EmitItem{
					ID:       id,
					Content:  d.BaseDataMap[id].Content,
					DiffType: uint8(DiffFileMoveOutAndChange),
					HeadInfo: "",
				}
			}
			ans = append(ans, obj)
			return

		}
		// else { // not exist in base: add directly
		obj.DiffType = uint8(DiffFileAdd)
		ans = append(ans, obj)
		return
		// }

	}
	return
}

// dealBaseDelEvent to handle the del base element of diff output
func (d *StringDataFlowItem) dealBaseDelEvent(id string) (ans []*Diff3EmitItem) {

	item, ok := d.BaseExistMap[id]
	if ok { // has been handled
		if item.DiffType == 100 { // only base moved down
			item.DiffType = 0
			ans = append(ans, item)
		} else if item.DiffType > 0 { // need add
			ans = append(ans, item)
		}
		// else { // ignore
		//	// do nothing
		// }
		delete(d.BaseExistMap, id)
		return
	}

	// not exist in baseExistMap: has not handled

	var obj = &Diff3EmitItem{ID: id, Content: d.BaseDataMap[id].Content}
	if _, ok := d.OrgDataMap[id]; ok { // exist in org

		if _, ok := d.HeadDataMap[id]; ok { // exist in head

			// if head vs org has changed?
			diffHeadItem := d.Head2OrgMap[id]
			// head is changed or not diffed with org

			switch DiffFileType(diffHeadItem.DiffType) {
			case 0: // no change: like the base added
				ans = append(ans, obj)
				d.HeadExistMap[id] = &Diff3EmitItem{ID: id, DiffType: 0}
				return

			case DiffFileChange:
				// TODO: for file it's may be seems complicated,eg:A rename, B update，A VS B is to +/-
				list := d.cmpHeadAndBaseObj(id)
				if list[0].DiffType > 0 { // changed: modify
					obj.DiffType = uint8(DiffFileChange)
					obj.HeadInfo = d.HeadDataMap[id].Content
				}
				// else { // no change
				//	// do nothing
				// }

				d.HeadExistMap[id] = &Diff3EmitItem{ID: id, DiffType: 0}
				ans = append(ans, obj)
				return

			case DiffFileAdd, DiffFileDel:
				if d.OrgDataMap[id].Content == d.HeadDataMap[id].Content { // head only move
					// whether base update or not, the diff result is move
					// TODO: notice to frontend to select the base content if agree with head's move option.
					obj.DiffType = uint8(DiffFileMoveOut)
					d.HeadExistMap[id] = &Diff3EmitItem{
						ID:       id,
						Content:  d.HeadDataMap[id].Content,
						DiffType: uint8(DiffFileMoveIn),
					}

				} else { // head move && update content
					list := d.cmpHeadAndBaseObj(id)
					if list[0].DiffType > 0 { // changed: modify
						obj.DiffType = uint8(DiffFileMoveOutAndChange)
						d.HeadExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.HeadDataMap[id].Content,
							DiffType: uint8(DiffFileMoveInAndChange),
						}

					} else { // no change
						obj.DiffType = uint8(DiffFileMoveOut)
						d.HeadExistMap[id] = &Diff3EmitItem{
							ID:       id,
							Content:  d.HeadDataMap[id].Content,
							DiffType: uint8(DiffFileMoveIn),
						}
					}
				}

				ans = append(ans, obj)
				return

			case DiffFileRename:
				// TODO: for file it's may be seems complicated,eg:A rename, B update，A VS B is to +/-
				logrus.Warnf("dealBaseDel has complex rename diffType")
			}

		} else { // not exist in head:del in head

			obj.DiffType = uint8(DiffFileDel)
			ans = append(ans, obj)
		}

	} else { // not exist in org: base added

		// todo : is show the element be moved?
		// if exist in head
		if headObj, ok := d.HeadDataMap[id]; ok { // exist in head: has updated?
			list := d.cmpHeadAndBaseObj(id)
			if list[0].DiffType > 0 { // has changed: modify
				obj.DiffType = list[0].DiffType
			}
			d.HeadExistMap[id] = &Diff3EmitItem{
				ID:       id,
				Content:  headObj.Content,
				DiffType: 0, // ignore if head obj has difference again
			}

		}
		ans = append(ans, obj)
		return

	}
	return
}

// dealChangeEvent to handle the possible case betweent base and head
func (d *StringDataFlowItem) dealChangeEvent(id string) (ans []*Diff3EmitItem) {
	var obj = *(d.BaseDataMap[id])
	if _, ok := d.OrgDataMap[id]; ok { // exist in org

		list := d.cmpBaseAndOrgObj(id)
		if list[0].DiffType == 0 {
			return d.cmpHeadAndBaseObj(id)
		}

		list2 := d.cmpHeadAndOrgObj(id)
		if list2[0].DiffType == 0 { // head not modify
			return []*Diff3EmitItem{&obj}
		}
		return d.cmpHeadAndBaseObj(id)

	}
	return d.cmpHeadAndBaseObj(id)
}

func prepareList(list []*Diff3EmitItem) ([]string, map[string]*Diff3EmitItem) {
	var (
		l = make([]string, 0, len(list))
		m = make(map[string]*Diff3EmitItem)
	)

	for _, v := range list {
		l = append(l, v.ID)
		m[v.ID] = v
	}
	return l, m
}

/*
		Alice               Original            Bb
	  1. celery           1. celery           1. celery         A
	  -----------------------------------------------------------
						  					2. garlic           2. salmon         B
	  2. salmon           3. onions           3. garlic
						  					4. salmon           4. onions
	  -----------------------------------------------------------
	  3. tomatoes         5. tomatoes         5. tomatoes       C
	  -----------------------------------------------------------
	  4. garlic                                                 D
	  5. onions
	  -----------------------------------------------------------
	  6. wine             6. wine             6. wine           E
*/
