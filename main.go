package main

import (
	ascii "ascii/functions"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . something standard")
		return
	}
	Banner := "files/standard.txt"
	if len(args) == 3 {
		if strings.HasSuffix(args[2], ".txt") {
			Banner = "files/" + args[2]
		} else {
			Banner = "files/" + args[2] + ".txt"
		}
	}
	file, err := os.Open(Banner)
	if err != nil {
		log.Fatal("Error opening the file")
	}
	defer file.Close()
	SliceRunes := []string{}
	MapAscii := map[rune][]string{}
	count := 0
	espace := ' '
	myscanner := bufio.NewScanner(file)
	for myscanner.Scan() {
		text := myscanner.Text()
		if text != "" {
			SliceRunes = append(SliceRunes, text)
			count++
		}
		if count == 8 {
			MapAscii[espace] = SliceRunes
			espace++
			SliceRunes = []string{}
			count = 0
		}
	}
	argument := args[1]
	if len(argument) == 0 {
		return
	}
	if ascii.Isprintable(argument) {
		log.Fatal("Isprintable charachtere not allowed")
	}
	Splitslice := strings.Split(argument, "\\n")
	if strings.ReplaceAll(argument, "\\n", "") == "" {
		for i := 0; i < strings.Count(argument, "\\n"); i++ {
			fmt.Println()
		}
		return
	}
	result := ascii.PrintAscii(Splitslice, MapAscii)
	fmt.Println(result)
}
