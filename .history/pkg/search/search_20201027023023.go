// package search

// import (
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"search/pkg/types"
// 	"strings"
// )
package search

import (
	"context"
	"fmt"
	"io/ioutil"
	"search/pkg/types"
	"strings"
	"sync"
)

func FindAllMatchTextInFile(phrase, file string, findingAll bool) (result []types.Result) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	for i, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, phrase) {
			found := types.Result{
				Phrase:  phrase,
				Line:    line,
				LineNum: int64(i + 1),
				ColNum:  int64(strings.Index(line, phrase) + 1),
			}

			result = append(result, found)

			if !findingAll {
				return result
			}
		}
	}

	fmt.Println(result)
	return result
}

// func All(ctx context.Context, phrase string, files []string) <-chan []types.Result {
// 	ch := make(chan []types.Result)
// 	wg := sync.WaitGroup{}

// 	ctx, cancel := context.WithCancel(ctx)

// 	for i, file := range files {
// 		wg.Add(1)
// 		fmt.Print("FILE: ")
// 		fmt.Println(file)

// 		go func(ctx context.Context, file string, i int, ch chan<- []types.Result) {
// 			defer wg.Done()

// 			allMatches := FindAllMatchTextInFile(phrase, file, true)

// 			if len(allMatches) > 0 {
// 				ch <- allMatches
// 			}
// 		}(ctx, file, i, ch)
// 	}

// 	go func() {
// 		defer close(ch)
// 		wg.Wait()
// 	}()

// 	cancel()
// 	return nil
// }

// //Result ..
// type Result struct {
// 	Phrase  string
// 	Line    string
// 	LineNum int64
// 	ColNum  int64
// }

// //All ...
func All(ctx context.Context, phrase string, files []string) <-chan []types.Result {
	ch := make(chan []types.Result)
	wg := sync.WaitGroup{}

	//var results []Result

	ctx, cancel := context.WithCancel(ctx)

	for i := 0; i < len(files); i++ {
		wg.Add(1)

		go func(ctx context.Context, filename string, i int, ch chan<- []types.Result) {
			defer wg.Done()

			res := FindAllMatchTextInFile(phrase, filename, true)

			if len(res) > 0 {
				ch <- res
			}

		}(ctx, files[i], i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()

	}()

	cancel()
	return ch
}

// //FindAllMatchTextInFile ...
// func FindAllMatchTextInFile(phrase, fileName string) (res []Result) {

// 	data, err := ioutil.ReadFile(fileName)
// 	if err != nil {
// 		log.Println("error not opened file err => ", err)
// 		return res
// 	}

// 	file := string(data)

// 	temp := strings.Split(file, "\n")

// 	for i, line := range temp {
// 		//fmt.Println("[", i+1, "]\t", line)
// 		if strings.Contains(line, phrase) {

// 			r := Result{
// 				Phrase:  phrase,
// 				Line:    line,
// 				LineNum: int64(i + 1),
// 				ColNum:  int64(strings.Index(line, phrase)) + 1,
// 			}

// 			res = append(res, r)
// 		}
// 	}

// 	return res
// }
