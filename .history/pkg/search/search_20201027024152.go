package search

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"sync"
)

// Result is type which can be used to store resulting data
type Result struct {
	Phrase  string
	Line    string
	LineNum int64
	ColNum  int64
}

// FindAllMatchTextInFile finds all phrase occurrences
func FindAllMatchTextInFile(phrase, file string, findingAll bool) (result []Result) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	for i, line := range strings.Split(string(data), "\n") {
		if strings.Contains(line, phrase) {
			found := Result{
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

// All is the main function for finding occurrences of phrase in given list of files
func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i, file := range files {
		wg.Add(1)

		go func(ctx context.Context, filename string, i int, ch chan<- []Result) {
			defer wg.Done()

			result := FindAllMatchTextInFile(phrase, filename, true)

			if len(result) > 0 {
				ch <- result
			}
		}(ctx, file, i, ch)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cancel()
	return ch
}
