package skiplist

import (
	"math/rand"
	"time"

	"github.com/Ruth-Seven/skiplists_benchmark/tools"
	huaSkiplist "github.com/huandu/skiplist"
)

func huanduInserts(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}
}

func huanduWorstInserts(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}
}

func huanduRandomInserts(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)
	rList := rand.Perm(n)
	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		list.Set(e, testByteString)
	}
}

func huanduAvgSearch(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(i)
	}
}

func huanduSearchEnd(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Get(n - 1)
	}
}

func huanduDelete(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(i)
	}
}

func huanduWorstDelete(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_ = list.Remove(n - i)
	}
}

func huanduRandomDelete(n int) {
	list := huaSkiplist.New(huaSkiplist.Int)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	rList := rand.Perm(n)
	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		_ = list.Remove(e)
	}
}

var HuanduFunctions = []func(int){huanduInserts, huanduWorstInserts, huanduRandomInserts,
	huanduAvgSearch, huanduSearchEnd, huanduDelete, huanduWorstDelete, huanduRandomDelete}
