package search

import (
	"context"
	"fmt"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {
	fmt.Println("Started testing...")

	ch := All(context.Background(), "ipsum", []string{"../../data/data.txt"})

	s, err := <-ch

	if !err {
		t.Error(err)
	}

	log.Println("=======>>>>>", s)

}
