package search

import (
	"context"
	"fmt"
	"io/ioutil"
	"search/pkg/types"
	"strings"
	"sync"
)

func All(ctx context.Context, phrase string, files []string) <-chan []types.Result {
	ch := make(chan []types.Result)
	defer close(ch)

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i, file := range files {
		wg.Add(1)
		fmt.Print("FILE: ")
		fmt.Println(file)

		go func(ctx context.Context, file string, i int, ch chan<- []types.Result) {
			defer wg.Done()

			allMatches := FindAllMatchTextInFile(phrase, file, true)

			if len(allMatches) > 0 {
				ch <- allMatches
			}
		}(ctx, file, i, ch)
	}

	wg.Wait()
	cancel()
	return ch
}

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
				ColNum:  int64(strings.Index(line, phrase)) + 1,
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
