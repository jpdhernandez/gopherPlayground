package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jpdhernandez/gopherPlayground/stringutils"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filePath := os.Args[1]

	absPath, _ := filepath.Abs(filePath)
	dat, err := ioutil.ReadFile(absPath)
	check(err)

	splitGoProverbs := strings.Split(string(dat), "\n")

	for _, proverb := range splitGoProverbs {
		fmt.Println(proverb)

		for key, val := range stringutils.GetCharCountMap(proverb) {
			s := fmt.Sprintf("'%s'=%d ", key, val)

			fmt.Print(s)
		}

		fmt.Printf("\n\n")
	}
}
