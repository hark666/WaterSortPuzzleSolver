package main

import (
	"fmt"
	"strconv"
)

func main() {
	InitialiseLevelBuffers(100)
	shortest := true
	levelNumber := 117
	fmt.Println("Solving Level "+strconv.Itoa(levelNumber)+", find shortest:", shortest)
	fmt.Println(len(levels))
	if (levelNumber - 1) > len(levels) {
		fmt.Println("We cant resolve this level because we dont have it")
	} else {
		var level = BuildLevel(levels[levelNumber-1])
		sol := level.Solve(shortest, true)
		for i, s := range sol {
			fmt.Printf("Move %2d: %2d -> %2d\n", i+1, s[0]+1, s[1]+1)
		}
	}
}
