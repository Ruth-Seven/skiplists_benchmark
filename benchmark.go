package main

import (
	"flag"
	"log"

	"github.com/Ruth-Seven/skiplists_benchmark/skiplist"
	"github.com/Ruth-Seven/skiplists_benchmark/tools"
)

var (
	// The range of values for n passed to the individual benchmarks
	start  = flag.Int("start", 100, "the lowest times of test operator")
	end    = flag.Int("end", 1000, "the larget times of test operator")
	factor = flag.Float64("factor", 1.5, "the  mutilple factor of start")
	output = flag.String("folder", "result", "the lowest times of test operator")
)

const (
	csvfile = "result.csv"
)

// runIterations executes the tests in a loop with the given parameters
func runIterations(name string, f func(int)) {
	tools.Tee(name + ",")
	for i := *start; i <= *end; i = int(float64(i) * (*factor)) {
		f(i)
		tools.Tee(",")
	}
	tools.Tee("\n")
}

func main() {
	flag.Parse()
	// first, print the CSV header with iteration counts
	runIterations("iterations", tools.Iterations)
	allFunctions := append(skiplist.MapFunctions, skiplist.SeanFunctions...)
	allFunctions = append(allFunctions, skiplist.ZhenjlFunctions...)
	allFunctions = append(allFunctions, skiplist.ColFunctions...)
	allFunctions = append(allFunctions, skiplist.MtchavezFunctions...)
	allFunctions = append(allFunctions, skiplist.HuanduFunctions...)
	// allFunctions = append(allFunctions, skiplist.RyszardFunctions...)
	allFunctions = append(allFunctions, skiplist.MtFunctions...)
	allFunctions = append(allFunctions, skiplist.RuthFunctions...)

	for _, f := range allFunctions {
		runIterations(tools.FuncName(f), f)
	}

	var folder = "./" + *output
	if err := tools.CloseTee(folder, csvfile); err != nil {
		log.Fatal(err)
	}
}
