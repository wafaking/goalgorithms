package tree

import (
	"encoding/json"
	"strings"
	"testing"
)

type fileItem struct {
	ListBase, ListOrg, ListHead []*Diff3EmitItem
	Expect                      []*Diff3EmitItem
	// diff                        []string
}

var fileList = []fileItem{
	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "c"},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 7, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b"},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 2, HeadInfo: "b"},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
			{ID: "B", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 7, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "c"},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 7, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "A", Content: ""},
			{ID: "C", Content: "c"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 7, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "c"},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 7, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b"},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 2, HeadInfo: "b"},
			{ID: "C", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: "d1"},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: "d2"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "d1", DiffType: 2, HeadInfo: "d2"},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "M", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "M", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "c", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: "b"},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "B", Content: "b", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "M", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "M", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 3, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: "c"},
			{ID: "D", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "c", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
			{ID: "D", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "C", Content: "c", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b"},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 2, HeadInfo: "b"},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b"},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b1"},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b2"},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "b1", DiffType: 2, HeadInfo: "b2"},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
			{ID: "B", Content: "b1"},
			{ID: "D", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b2"},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "b1", DiffType: 2, HeadInfo: "b2"},
			{ID: "D", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 7, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
			{ID: "B", Content: ""},
			{ID: "D", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 7, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 8, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "B", Content: ""},
			{ID: "E", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "E", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: "b"},
			{ID: "A", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "B", Content: "b", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b"},
			{ID: "C", Content: ""},
			{ID: "E", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
			{ID: "D", Content: "d"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "b", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "c"},
			{ID: "E", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "d", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: "a"},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: "a"},
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "c"},
			{ID: "A", Content: "a", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: "a"},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: "c"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "a", DiffType: 3, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "A", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 3, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "A", Content: "a"},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "B", Content: "b"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "B", Content: "", DiffType: 2, HeadInfo: "b"},
			{ID: "A", Content: "a", DiffType: 3, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
			{ID: "E", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "D", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "E", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: "a"},
			{ID: "E", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: "b"},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "a", DiffType: 0, HeadInfo: ""},
			{ID: "E", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 3, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "B", Content: "b"},
			{ID: "E", Content: ""},
			{ID: "A", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: "a"},
			{ID: "D", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
			{ID: "B", Content: "b", DiffType: 0, HeadInfo: ""},
			{ID: "E", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 2, HeadInfo: "a"},
		},
	},

	{
		ListBase: []*Diff3EmitItem{},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "B", Content: ""},
			{ID: "A", Content: ""},
			{ID: "C", Content: ""},
			{ID: "D", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "D", Content: "", DiffType: 1, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: "c"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "c", DiffType: 9, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 10, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "1"},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: "2"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "2", DiffType: 9, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 10, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "2"},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: "2"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "2", DiffType: 7, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 8, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: "2"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "2", DiffType: 9, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 10, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: "1"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "1", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: "1"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "1"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "1", DiffType: 0, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: "2"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "1"},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "2", DiffType: 2, HeadInfo: "1"},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: "2"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 3, HeadInfo: "2"},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: "1"},
		},
		Expect: []*Diff3EmitItem{

			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "1"},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: "1"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "1", DiffType: 3, HeadInfo: ""},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},

	{
		ListBase: []*Diff3EmitItem{
			{ID: "C", Content: ""},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		ListOrg: []*Diff3EmitItem{
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
			{ID: "C", Content: ""},
		},
		ListHead: []*Diff3EmitItem{
			{ID: "C", Content: "1"},
			{ID: "A", Content: ""},
			{ID: "B", Content: ""},
		},
		Expect: []*Diff3EmitItem{
			{ID: "C", Content: "", DiffType: 2, HeadInfo: "1"},
			{ID: "A", Content: "", DiffType: 0, HeadInfo: ""},
			{ID: "B", Content: "", DiffType: 0, HeadInfo: ""},
		},
	},
}

func TestDiff3Core(t *testing.T) {
	for idx, item := range fileList {
		// if idx != 28 {
		//	continue
		// }
		var (
			headIDs, headMap              = prepareList(item.ListHead)
			orgIDs, orgMap                = prepareList(item.ListOrg)
			baseIDs, baseMap              = prepareList(item.ListBase)
			org2HeadIndexMap, head2OrgMap = getMatchedIndexMap(headIDs, orgIDs, headMap, orgMap)
			org2BaseIndexMap, base2OrgMap = getMatchedIndexMap(baseIDs, orgIDs, baseMap, orgMap)
			// head2OrgMap                   = getMatchedMap(org2HeadIndexMap, item.ListHead, item.ListOrg)
			// base2OrgMap                   = getMatchedMap(org2BaseIndexMap, item.ListBase, item.ListOrg)
		)

		var obj = &StringDataFlowItem{
			ListOrg:          item.ListOrg,
			ListBase:         item.ListBase,
			ListHead:         item.ListHead,
			HeadDataMap:      headMap,
			BaseDataMap:      baseMap,
			OrgDataMap:       orgMap,
			Org2HeadIndexMap: org2HeadIndexMap,
			Org2BaseIndexMap: org2BaseIndexMap,
			Head2OrgMap:      head2OrgMap,
			Base2OrgMap:      base2OrgMap,
			BaseExistMap:     make(map[string]*Diff3EmitItem, 0),
			HeadExistMap:     make(map[string]*Diff3EmitItem, 0),
		}
		res := obj.Diff3Core()
		k, ok := Diff3EmitItemSliceAssert(res, item.Expect)
		if !ok {
			resStr, _ := json.Marshal(res)
			expectStr, err := json.Marshal(item)
			if err != nil {
				t.Fatal(err)
			}
			t.Fatalf("falied,idx:%d:%d,\n res:%s,\n Expect:%s", idx, k, string(resStr), string(expectStr))
		}

	}
}

func Diff3EmitItemSliceAssert(res, expect []*Diff3EmitItem) (int, bool) {
	if len(res) != len(expect) {
		return -1, false
	}

	if res == nil || expect == nil {
		return -1, false
	}
	for k, v := range expect {
		if v.ID != res[k].ID ||
			v.Content != res[k].Content ||
			v.DiffType != res[k].DiffType ||
			v.HeadInfo != res[k].HeadInfo {
			return k, false
		}
	}
	return -1, true
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

func getMatchedIndexMap(head, base []string, headMap, baseMap map[string]*Diff3EmitItem) (map[int]int, map[string]*Diff3EmitItem) {
	var (
		m                = make(map[int]int, 0)
		obj2OrgMap       = make(map[string]*Diff3EmitItem, 0)
		outPut           = GenerateDiff(base, head)
		headIdx, baseIdx = 0, 0
	)

	for _, v := range outPut {
		v = strings.TrimSpace(v)
		if strings.HasPrefix(v, "+") {
			headIdx++
			v = strings.TrimSpace(strings.TrimPrefix(v, "+"))
			temp := *headMap[v]
			temp.DiffType = uint8(DiffFileAdd)
			obj2OrgMap[v] = &temp

		} else if strings.HasPrefix(v, "-") {
			baseIdx++
			v = strings.TrimSpace(strings.TrimPrefix(v, "-"))
			temp := *baseMap[v]
			temp.DiffType = uint8(DiffFileDel)
			obj2OrgMap[v] = &temp

		} else {
			m[baseIdx] = headIdx
			baseIdx++
			headIdx++

			v = strings.TrimSpace(v)
			item := &Diff3EmitItem{
				ID:       v,
				DiffType: 0,
			}
			if headMap[v].Content != baseMap[v].Content {
				item.DiffType = uint8(DiffFileChange)
			}
			obj2OrgMap[v] = item

		}
	}
	return m, obj2OrgMap
}
