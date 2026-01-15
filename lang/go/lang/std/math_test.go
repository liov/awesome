package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"testing"
)

func TestUint64(t *testing.T) {
	log.Println(uint64(3) * uint64(65) / uint64(5))
	a4 := uint64(3) * uint64(32) / uint64(5)
	a5 := uint64(3) * uint64(33) / uint64(5)

	log.Println("1", a4)
	log.Println("1", a5)
	log.Println("1", a4+a5)

}
func TestInt64N(t *testing.T) {
	fmt.Println(rand.Int64N(1000))
	fmt.Println(rand.Int64N(1000))
	fmt.Println(rand.Int64N(1000))
	fmt.Println(rand.Int64N(1000))
	fmt.Println(rand.Int64N(1000))
	fmt.Println(rand.Int64N(1000))
}
