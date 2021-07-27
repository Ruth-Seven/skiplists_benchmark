package tools

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"reflect"
	"runtime"
	"time"
)

// tools.TimeTrack will print out the number of nanoseconds since the start time divided by n
// Useful for printing out how long each iteration took in a benchmark
func TimeTrack(start time.Time, n int) {
	loopNS := time.Since(start).Nanoseconds() / int64(n)
	Tee(fmt.Sprint(loopNS))
}

// iterations is used to print out the CSV header with iteration counts
func Iterations(n int) {
	Tee(fmt.Sprint(n))
}

// funcName returns just the function name of a string, given any function at all
func FuncName(f func(int)) string {
	longFunc := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()[5:]
	_, file := path.Split(longFunc)
	return file
}

func FolderName(a, b, c int) string {
	return fmt.Sprintf("Benchs%3es%3ee%3e", float32(a), float32(b), float32(c))
}

var buf *bytes.Buffer

func Tee(s string) {
	fmt.Print(s)
	if buf == nil {
		buf = bytes.NewBuffer(nil)
	}
	buf.Write([]byte(s))
}

func CloseTee(folder, file string) error {
	if buf == nil {
		return fmt.Errorf("had not called Tee before")
	}
	p := "./" + folder + "/" + file
	if err := os.MkdirAll(path.Dir(p), 0755); err != nil {
		return err
	}
	f, err := os.Create(path.Clean(p))
	if err != nil {
		if err != os.ErrExist {
			fmt.Print("ee")
			return err
		}
	}
	_, err = buf.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}
