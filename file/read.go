package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("cryto.csv")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		item := strings.Split(line, ",")

		// fmt.PrintIn("Line: %s\n", line)
		fmt.Printf(item[1])

		count++
	}
}