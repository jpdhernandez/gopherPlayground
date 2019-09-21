package main

import (
	"fmt"
	"strings"

	"github.com/jpdhernandez/gopherPlayground/constants"
)

func main() {
	splitGoProverbs := strings.Split(constants.GoProverbs, "\n")

	for i, proverb := range splitGoProverbs {
		fmt.Printf("%d. ", i+1)
		fmt.Println(proverb)
	}
}
