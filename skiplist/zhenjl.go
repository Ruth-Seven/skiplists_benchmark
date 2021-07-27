package skiplist

import (
	"math/rand"
	"time"

	"github.com/skiplist-survey/tools"
	zheSkiplist "github.com/zhenjl/skiplist"
)

func zhenjlInserts(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)
	defer tools.TimeTrack(time.Now(), n)
	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}
}

func zhenjlWorstInserts(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(i, testByteString)
	}
}

func zhenjlRandomInserts(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)
	for _, e := range rList {
		list.Insert(e, testByteString)
	}
}

func zhenjlAvgSearch(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Select(i)
	}
}

func zhenjlSearchEnd(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Select(n - 1)
	}
}

func zhenjlDelete(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(i)
	}
}

func zhenjlWorstDelete(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(n - i)
	}
}

func zhenjlRandomDelete(n int) {
	list := zheSkiplist.New(zheSkiplist.BuiltinLessThan)

	for i := 0; i < n; i++ {
		list.Insert(n-i, testByteString)
	}

	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		_, _ = list.Delete(e)
	}
}

var ZhenjlFunctions = []func(int){zhenjlInserts, zhenjlWorstInserts, zhenjlRandomInserts,
	zhenjlAvgSearch, zhenjlSearchEnd, zhenjlDelete, zhenjlWorstDelete, zhenjlRandomDelete}
