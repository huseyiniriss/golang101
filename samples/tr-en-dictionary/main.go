package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Word struct {
	English string
	Turkish string
}

func main() {
	data, err := readData()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Welcome to the dictionary")
	for {
		fmt.Println("Enter q to quit")
		fmt.Print("Enter a word: ")
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
		findMean(data, input)
	}
}

func findMean(data []Word, input string) {
	for _, word := range data {
		if strings.ToLower(word.English) == strings.ToLower(input) || strings.ToLower(word.Turkish) == strings.ToLower(input) {
			fmt.Println("Tr meaning:", word.Turkish)
			fmt.Println("En meaning:", word.English)
			return
		}
	}
	fmt.Print("Meaning not found\n\n")
}

func readData() ([]Word, error) {
	data := make([]Word, 0)
	path := "samples/tr-en-dictionary/data.txt"
	file, err := os.Open(path)

	if err != nil {
		return data, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " - ")
		word := Word{English: line[0], Turkish: line[1]}
		data = append(data, word)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
		return data, err
	}
	return data, nil
}
