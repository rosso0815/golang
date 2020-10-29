package play

import (
	"log"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestRandom(t *testing.T) {

	log.Println("@@@ TestRandom")

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	log.Println("My favorite number is", r1.Intn(100))
	log.Println("Pi=", math.Pi)

}
