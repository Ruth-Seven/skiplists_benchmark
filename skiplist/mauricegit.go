package skiplist

import (
	"fmt"
	"math/rand"
	"time"

	mtSkiplist "github.com/MauriceGit/skiplist"
	"github.com/Ruth-Seven/skiplists_benchmark/tools"
)

type element uint64

func (e element) ExtractKey() float64 {
	return float64(e)
}
func (e element) String() string {
	return fmt.Sprintf("%03d", e)
}

func mtInserts(n int) {
	list := mtSkiplist.New()
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}
}

func mtWorstInserts(n int) {
	list := mtSkiplist.New()
	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Insert(element(i))
	}
}

func mtRandomInserts(n int) {
	list := mtSkiplist.New()

	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)
	for _, e := range rList {
		list.Insert(element(e))
	}
}

func mtAvgSearch(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(element(i))
	}
}

func mtSearchEnd(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		_, _ = list.Find(element(n - 1))
	}
}

func mtDelete(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Delete(element(i))
	}
}

func mtWorstDelete(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	defer tools.TimeTrack(time.Now(), n)

	for i := 0; i < n; i++ {
		list.Delete(element(n - i))
	}
}

func mtRandomDelete(n int) {
	list := mtSkiplist.New()

	for i := 0; i < n; i++ {
		list.Insert(element(n - i))
	}

	rList := rand.Perm(n)

	defer tools.TimeTrack(time.Now(), n)

	for _, e := range rList {
		list.Delete(element(e))
	}
}

var MtFunctions = []func(int){mtInserts, mtWorstInserts, mtRandomInserts,
	mtAvgSearch, mtSearchEnd, mtDelete, mtWorstDelete, mtRandomDelete}
