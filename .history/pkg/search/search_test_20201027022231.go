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

	s, ok := <-ch

	if !ok {
		t.Errorf(" function All error => %v", ok)
	}

	log.Println("=======>>>>>", s)

}
