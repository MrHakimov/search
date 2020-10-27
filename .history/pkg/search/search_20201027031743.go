package search

import (
	"context"
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

// FindMatchesInFile finds all phrase occurrences
func FindMatchesInFile(phrase, file string, findingAll bool) []Result {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	var result []Result = nil
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

			result := FindMatchesInFile(phrase, filename, true)

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

// Any is the main function for finding one of the occurrences of phrase in given list of files
func Any(ctx context.Context, phrase string, files []string) <-chan Result {
	ch := make(chan Result)
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	var result Result
	for i, file := range files {
		current := FindMatchesInFile(phrase, file, false)

		if len(current) > 0 {
			result = current[0]
			cancel()
		}
	}

	wg.Add(1)

	go func(ctx context.Context, filename string, i int, ch chan<- Result) {
		defer wg.Done()

		result := FindMatchesInFile(phrase, filename, false)

		if len(result) > 0 {
			var forCh Result = result[0]
			ch <- forCh
			cancel()
		}
	}(ctx, file, i, ch)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	cancel()

	return ch
}
