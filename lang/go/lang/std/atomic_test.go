package main

import (
	"log"
	"sync/atomic"
	"testing"
)

func TestAddUint64(t *testing.T) {
	var a uint64 = 10
	atomic.AddUint64(&a, ^uint64(0))
	log.Println(a)
}
