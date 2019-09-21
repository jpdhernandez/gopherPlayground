package main

import (
	"fmt"
	"strings"

	"github.com/jpdhernandez/gopherPlayground/constants"
	"github.com/jpdhernandez/gopherPlayground/stringutils"
)

func main() {
	splitGoProverbs := strings.Split(constants.GoProverbs, "\n")

	for _, proverb := range splitGoProverbs {
		fmt.Println(proverb)

		for key, val := range stringutils.GetCharCountMap(proverb) {
			s := fmt.Sprintf("'%s'=%d ", key, val)

			fmt.Print(s)
		}

		fmt.Printf("\n\n")
	}
}
