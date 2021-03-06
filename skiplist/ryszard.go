package skiplist

import (
	"math/rand"
	"time"

	"github.com/Ruth-Seven/skiplists_benchmark/tools"
	rysSkiplist "github.com/ryszard/goskiplist/skiplist"
)

func ryszardInserts(n int) {
	list := rysSkiplist.NewIntMap()
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}
}

func ryszardWorstInserts(n int) {
	list := rysSkiplist.NewIntMap()
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Set(i, testByteString)
	}
}

func ryszardRandomInserts(n int) {
	list := rysSkiplist.NewIntMap()
	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		list.Set(e, testByteString)
	}
}

func ryszardAvgSearch(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Get(i)
	}
}

func ryszardSearchEnd(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Get(n)
	}
}

func ryszardDelete(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(i)
	}
}

func ryszardWorstDelete(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Delete(n - i)
	}
}

func ryszardRandomDelete(n int) {
	list := rysSkiplist.NewIntMap()

	for i := 0; i < n; i++ {
		list.Set(n-i, testByteString)
	}

	rList := rand.Perm(n)
	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		_, _ = list.Delete(e)
	}
}

var RyszardFunctions = []func(int){ryszardInserts, ryszardWorstInserts, ryszardRandomInserts,
	ryszardAvgSearch, ryszardSearchEnd, ryszardDelete, ryszardWorstDelete, ryszardRandomDelete}
