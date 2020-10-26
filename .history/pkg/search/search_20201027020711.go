package search

import (
	"context"
	"search/pkg/types"
	"sync"
)

func All(ctx context.Context, phrase string, files []string) <-chan []types.Result {
	ch := make(chan []types.Result)
	defer close(ch)

	wg := sync.WaitGroup{}

	ctx, cancel := context.WithCancel(ctx)

	for i, file := range files {
		wg.Add(1)

		go func(ctx context.Context, file string, i int, ch chan<- []types.Result) {
			defer wg.Done()

			allMatches := FindAllMatchTextInFile(phrase, file)

			if len(allMatches) > 0 {
				ch <- allMatches
			}
		}(ctx, file, i, ch)
	}

	wg.Wait()
	cancel()
	return ch
}
