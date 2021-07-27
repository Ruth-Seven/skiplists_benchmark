package skiplist

import (
	"math/rand"
	"time"

	"github.com/skiplist-survey/tools"
)

type MockList map[int64][]byte

func getMockList() MockList {
	return make(map[int64][]byte)
}

func (l *MockList) Insert(key int64, val []byte) {
	(*l)[key] = val
}

func (l *MockList) Find(key int64) ([]byte, bool) {
	if res, ok := (*l)[key]; !ok {
		return nil, false
	} else {
		return res, true
	}
}

func (l *MockList) Delete(key int64) {
	delete((*l), key)
}
func mapInserts(n int) {
	list := getMockList()
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(int64(n-i), testByteString)
	}
}

func mapWorstInserts(n int) {
	list := getMockList()
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(int64(i), testByteString)
	}
}

func mapRandomInserts(n int) {
	list := getMockList()

	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)
	for _, e := range rList {
		list.Insert(int64(e), testByteString)
	}
}

func mapAvgSearch(n int) {
	list := getMockList()

	for i := 0; i < n; i++ {
		list.Insert(int64(n-i), testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(int64(i))
	}
}

func mapSearchEnd(n int) {
	list := getMockList()

	for i := 0; i < n; i++ {
		list.Insert(int64(n-i), testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(int64(i))
	}
}

func mapDelete(n int) {
	list := getMockList()

	for i := 0; i < n; i++ {
		list.Insert(int64(n-i), testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Delete(int64(i))
	}
}

func mapWorstDelete(n int) {
	list := getMockList()

	for i := 0; i < n; i++ {
		list.Insert(int64(n-i), testByteString)
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Delete(int64(n - i))
	}
}

func mapRandomDelete(n int) {
	list := getMockList()

	for i := 0; i < n; i++ {
		list.Insert(int64(n-i), testByteString)
	}

	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		list.Delete(int64(e))
	}
}

var MapFunctions = []func(int){mapInserts, mapWorstInserts, mapRandomInserts,
	mapAvgSearch, mapSearchEnd, mapDelete, mapWorstDelete, mapRandomDelete}
