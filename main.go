package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 && len(os.Args) != 3 && len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
		return
	}

	data, err := os.ReadFile("standard.txt")

	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	var substring string
	input := ""
	color := ""
	reset := "\033[0m"

	if len(os.Args) == 2 {
		input = os.Args[1]
	}

	if len(os.Args) == 3 {

		if !strings.HasPrefix(os.Args[1], "--color=") {
			fmt.Println("usage: go run . option, colorName")
			return
		}
		input = os.Args[2]
	}

	if len(os.Args) == 4 {
		substring = os.Args[2]
		input = os.Args[3]
	}

	if len(os.Args) == 3 || len(os.Args) == 4 {
		colorName := strings.TrimPrefix(os.Args[1], "--color=")
		color = getColorCode(colorName)

		if color == "" {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <substring to be colored> something")
			return
		}
	}

	banner := string(data)

	banner = strings.ReplaceAll(banner, "\r", "")

	bannerTwo := strings.Split(banner, "\n")

	// input := os.Args[1]

	line := strings.Split(input, "\\n")

	for i, word := range line {
		if word == "" && i == 0 {
			continue
		}
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 1; i < 9; i++ {
			for index, ascii := range word {
				if ascii < 32 || ascii > 126 {
					fmt.Println("user input out of range")
					return
				}

				start := (int(ascii) - 32) * 9
				if color != "" && substring != "" && shouldColor(word, substring, index) {
					fmt.Print(color + bannerTwo[start+i] + reset)
				} else if color != "" && substring == "" {
					fmt.Print(color + bannerTwo[start+i] + reset)
				} else {
					fmt.Print(bannerTwo[start+i])
				}
			}
			fmt.Println()
		}
	}
}
func getColorCode(colorName string) string {
	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"yellow": "\033[33m",
		"blue":   "\033[34m",
		"purple": "\033[35m",
		"cyan":   "\033[36m",
		"white":  "\033[37m",
		"orange": "\033[38;2;255;165;0m",
	}

	return colors[strings.ToLower(colorName)]
}

func shouldColor(line string, substring string, index int) bool {
	if substring == "" {
		return true
	}

	for start := 0; start <= len(line)-len(substring); start++ {
		if line[start:start+len(substring)] == substring {
			end := start + len(substring)
			if index >= start && index < end {
				return true
			}
		}
	}

	return false
}
