package main

import (
	"fmt"
	"strings"

	"github.com/jpdhernandez/gopherPlayground/constants"
	"github.com/jpdhernandez/gopherPlayground/stringutils"
)

func main() {
	splitGoProverbs := strings.Split(constants.GoProverbs, "\n")

	for i, proverb := range splitGoProverbs {
		fmt.Printf("%d. ", i+1)
		fmt.Print(proverb)
		fmt.Printf(" (WC: %d)\n", stringutils.WordCountByString(proverb))
	}
}
