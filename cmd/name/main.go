package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/paulbellamy/name"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(name.Generate())
}
