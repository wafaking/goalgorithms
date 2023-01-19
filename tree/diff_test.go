package tree

import "testing"

type fileItem struct {
	fileA, fileO, fileB []string
	expect              []string
}

var fileList = []fileItem{
	{ //1. A is the same to B
		/*
				Alice               Original            Bob
			  1. celery           1. celery           1. celery         A
			  2. salmon           3. salmon           3. salmon
			  3. tomatoes         5. tomatoes         5. tomatoes       C
			  6. wine             6. wine             6. wine           E
		*/
		fileA: []string{
			"celery",
			"salmon",
			"tomatoes",
			"wine",
		},
		fileO: []string{
			"celery",
			"salmon",
			"tomatoes",
			"wine",
		},
		fileB: []string{
			"celery",
			"salmon",
			"tomatoes",
			"wine",
		},
		expect: []string{
			"celery",
			"salmon",
			"tomatoes",
			"wine",
		},
	}, // 1  A is the same to B

	{ //2. A is absolute different from B
		/*
				Alice               Original            Bob
			  1. AAAAA1  		  2. OOOOO1           2. BBBBB1         B
			  2. AAAAA2           3. OOOOO2           3. BBBBB2
			  3. AAAAA3			  4. OOOOO3           4. BBBBB3
		*/
		fileA: []string{
			"AAAAA1",
			"AAAAA2",
			"AAAAA3",
		},
		fileO: []string{
			"OOOOO1",
			"OOOOO2",
			"OOOOO3",
		},
		fileB: []string{
			"BBBBB1",
			"BBBBB2",
			"BBBBB3",
		},
		expect: []string{
			"<<<<<<< fileA",
			"AAAAA1",
			"AAAAA2",
			"AAAAA3",
			"=======",
			"BBBBB1",
			"BBBBB2",
			"BBBBB3",
			">>>>>>> fileB",
		},
	}, // 2  A is absolute different from B

	{ // 3：开头匹配，中间有匹配，结尾不匹配
		/*
				Alice               Original            Bob

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
		fileA: []string{
			"celery",
			"salmon",
			"tomatoes",
			"garlic",
			"onions",
			"wine",
		},
		fileO: []string{
			"celery",
			"garlic",
			"onions",
			"salmon",
			"tomatoes",
			"wine",
		},
		fileB: []string{
			"celery",
			"salmon",
			"garlic",
			"onions",
			"tomatoes",
			"wine",
		},
		expect: []string{
			"celery",
			"<<<<<<< fileA",
			"salmon",
			"=======",
			"salmon",
			"garlic",
			"onions",
			">>>>>>> fileB",
			"tomatoes",
			"<<<<<<< fileA",
			"garlic",
			"onions",
			"=======",
			">>>>>>> fileB",
			"wine",
		},
	}, // 3

	{
		/*
				Alice			Original		 Bob
			  1. celery			1. celery		1. celery
			  2. bab			2.bob			2. bob            A
			  -----------------------------------------------------------
								  2. garlic           2. salmon		B
			  2. salmon           3. onions           3. garlic
								  4. salmon           4. onions
			  -----------------------------------------------------------
			  3. tomatoes         5. tomatoes         5. tomatoes       C
			  -----------------------------------------------------------
			  4. garlic                                                 D
			  5. onions
			  -----------------------------------------------------------
			  6. wine             6. wine             6. wine           E
			  6. wine1            6. wine1           6. wine1           E
		*/
		fileA: []string{
			"celery",
			"bob",
			"salmon",
			"tomatoes",
			"garlic",
			"onions",
			"wine",
			"wine1",
		},
		fileO: []string{
			"celery",
			"bob",
			"garlic",
			"onions",
			"salmon",
			"tomatoes",
			"wine",
			"wine1",
		},
		fileB: []string{
			"celery",
			"bob",
			"salmon",
			"garlic",
			"onions",
			"tomatoes",
			"wine",
			"wine1",
		},
		expect: []string{
			"celery",
			"bob",
			"<<<<<<< fileA",
			"salmon",
			"=======",
			"salmon",
			"garlic",
			"onions",
			">>>>>>> fileB",
			"tomatoes",
			"<<<<<<< fileA",
			"garlic",
			"onions",
			"=======",
			">>>>>>> fileB",
			"wine",
			"wine1",
		},
	}, // 4

	{
		/*
				Alice			Original		 Bob

			  					  2. garlic1           2. salmon         B
			  2. salmon           3. onions1           3. garlic
								  4. salmon1           4. onions
			  -----------------------------------------------------------
			  3. tomatoes         5. tomatoes         5. tomatoes       C
			  -----------------------------------------------------------
			  4. garlic                                                 D
			  5. onions
			  -----------------------------------------------------------
			  6. wine             6. wine             6. wine           E
		*/
		fileA: []string{
			"salmon",
			"tomatoes",
			"garlic",
			"onions",
			"wine",
		},
		fileO: []string{
			"garlic1",
			"onions1",
			"salmon1",
			"tomatoes",
			"wine",
		},
		fileB: []string{
			"salmon",
			"garlic",
			"onions",
			"tomatoes",
			"wine",
		},
		expect: []string{
			"<<<<<<< fileA",
			"salmon",
			"=======",
			"salmon",
			"garlic",
			"onions",
			">>>>>>> fileB",
			"tomatoes",
			"<<<<<<< fileA",
			"garlic",
			"onions",
			"=======",
			">>>>>>> fileB",
			"wine",
		},
	}, // 5

	{
		/*
				Alice				Original		 	Bob
			  					  2. garlic1           2. salmon         B
			  2. salmon           3. onions1           3. garlic
								  4. salmon1           4. onions
			  -----------------------------------------------------------
			  3. tomatoes         5. tomatoes         5. tomatoes       C
			  -----------------------------------------------------------
			  4. garlic                                                 D
			  5. onions
		*/
		fileA: []string{
			"salmon",
			"tomatoes",
			"garlic",
			"onions",
		},
		fileO: []string{
			"garlic1",
			"onions1",
			"salmon1",
			"tomatoes",
		},
		fileB: []string{
			"salmon",
			"garlic",
			"onions",
			"tomatoes",
		},
		expect: []string{
			"<<<<<<< fileA",
			"salmon",
			"=======",
			"salmon",
			"garlic",
			"onions",
			">>>>>>> fileB",
			"tomatoes",
			"<<<<<<< fileA",
			"garlic",
			"onions",
			"=======",
			">>>>>>> fileB",
		},
	}, // 6

	{
		/* 7. o's lines number is more than A and B
			Alice			Original		 Bob

		  					  2. garlic           2. salmon         B
		  2. salmon           3. onions           3. garlic
							  4. salmon           4. onions
		  -----------------------------------------------------------
		  3. tomatoes         5. tomatoes         5. tomatoes       C
		  -----------------------------------------------------------
		  					  6. wine
		*/
		fileA: []string{
			"salmon",
			"tomatoes",
		},
		fileO: []string{
			"garlic",
			"onions",
			"salmon",
			"tomatoes",
			"wine",
		},
		fileB: []string{
			"salmon",
			"garlic",
			"onions",
			"tomatoes",
		},
		expect: []string{
			"<<<<<<< fileA",
			"salmon",
			"=======",
			"salmon",
			"garlic",
			"onions",
			">>>>>>> fileB",
			"tomatoes",
		},
	}, // 7. o's lines number is more than A and B

	{
		/*
				Alice			Original		 Bob

			  					  2. garlic1          2. garlic         B
			  2. salmon           3. onions           3. salmon
								  4. salmon1          4. onions
			  -----------------------------------------------------------
			  3. tomatoes         5. tomatoes         5. tomatoes       C
			  -----------------------------------------------------------
			  4. garlic                               6. apple         D
			  5. onions								  7. wine
		*/
		fileA: []string{
			"salmon",
			"tomatoes",
			"garlic",
			"onions",
		},
		fileO: []string{
			"garlic1",
			"onions",
			"salmon1",
			"tomatoes",
		},
		fileB: []string{
			"garlic",
			"salmon",
			"onions",
			"tomatoes",
			"apple",
			"wine",
		},
		expect: []string{
			"<<<<<<< fileA",
			"salmon",
			"=======",
			"garlic",
			"salmon",
			"onions",
			">>>>>>> fileB",
			"tomatoes",
			"<<<<<<< fileA",
			"garlic",
			"onions",
			"=======",
			"apple",
			"wine",
			">>>>>>> fileB",
		},
	}, // 8. o's lines number is less than A and B

}

func TestDiff3(t *testing.T) {
	for idx, item := range fileList {
		//if idx < 2 {
		//	continue
		//}
		res := Diff3(item.fileO, item.fileA, item.fileB)
		k, ok := stringSliceEqual(res, item.expect)
		if !ok {
			t.Logf("falied,idx:%d:%d,\n res:%#v,\n expect:%#v", idx, k, res, item.expect)
		}
	}
}

func stringSliceEqual(sli1, sli2 []string) (int, bool) {
	if len(sli1) != len(sli2) {
		return -1, false
	}

	if sli1 == nil || sli2 == nil {
		return -1, false
	}
	for k, v := range sli2 {
		if v != sli1[k] {
			return k, false
		}
	}
	return -1, true
}
