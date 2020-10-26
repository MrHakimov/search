package search

import (
	"context"
	"fmt"
	"testing"
)

func TestAll_success(t *testing.T) {
	fmt.Println("Started testing...")

	ch := All(context.Background(), "ipsum", []string{"../../data/data.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}
}

func TestAll_success(t *testing.T) {
	fmt.Println("Started testing...")

	ch := All(context.Background(), "ipsum", []string{"../../data/data.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}
}
