package search

import (
	"context"
	"log"
	"testing"
)

func TestAll_user(t *testing.T) {
	log.Println("Started testing...")

	ch := All(context.Background(), "ipsum", []string{"../../data.txt"})

	s, ok := <-ch

	if !ok {
		t.Errorf(" function All error => %v", ok)
	}

	log.Println("=======>>>>>", s)

}
