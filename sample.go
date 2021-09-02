package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		variableToTrace := rand.Int()
		fmt.Println(variableToTrace)
		time.Sleep(time.Second)
	}
}
