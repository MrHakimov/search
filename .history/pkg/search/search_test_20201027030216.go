package search

import (
	"context"
	"testing"
)

func TestAll_success(t *testing.T) {
	ch := All(context.Background(), "ipsum", []string{"../../data/data.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}
}

func TestAny_success(t *testing.T) {
	ch := Any(context.Background(), "ipsum", []string{"../../data/data.txt"})

	_, err := <-ch

	if !err {
		t.Error(err)
	}

	if err {
		t.Error(err)
	}
}
