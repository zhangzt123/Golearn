package kfk

import (
	"sync"
	"testing"
)

var sw sync.WaitGroup

func TestKfk(t *testing.T) {
	sw.Add(1)
	go provider(sw)
	//go consumer(sw)
	sw.Wait()
}
