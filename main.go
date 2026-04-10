package main

import (
	"fmt"
	"os"
	"strings"
)

// func to handle error
func usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING]")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <substring to be colored> something")
}

func main() {

	// check if arg is less than two or greater than 5 return error
	if len(os.Args) < 2 || len(os.Args) > 5 {
		usage()
		return
	}

	banneFile := "standard.txt"

	var substring string
	input := ""
	color := ""
	reset := "\033[0m"

	// check if args == two thne input is index one
	if len(os.Args) == 2 {
		input = os.Args[1]
	}

	//   check if len of arg is 3 and start with flag color then color it  esle input index one and banner index 2
	if len(os.Args) == 3 {
		if strings.HasPrefix(os.Args[1], "--color=") {
			// color whole string
			colorName := strings.TrimPrefix(os.Args[1], "--color=")
			color = getColorCode(colorName)
			if color == "" {
				fmt.Println("Color not found on the list")
				return
			}
			input = os.Args[2]
			substring = ""
		} else {
			// return input(string) and banner index
			input = os.Args[1]
			banneFile = os.Args[2]
		}
	}

	// check if arg is equal 4 then coloer is index one substring index 2 string index 3
	if len(os.Args) == 4 {
		if !strings.HasPrefix(os.Args[1], "--color=") {
			usage()
			return
		}

		colorName := strings.TrimPrefix(os.Args[1], "--color=")
		color = getColorCode(colorName)
		if color == "" {
			fmt.Println("Color not found on the list")
			return
		}

		substring = os.Args[2]
		input = os.Args[3]
	}

	// adding fs to code if arg is equal 5 then color is index 1 substring index 2 string index 3 and fs index 4
	if len(os.Args) == 5 {
		if !strings.HasPrefix(os.Args[1], "--color=") {
			usage()
			return
		}

		colorName := strings.TrimPrefix(os.Args[1], "--color=")
		color = getColorCode(colorName)
		if color == "" {
			fmt.Println("Color not found on the list")
			return
		}

		substring = os.Args[2]
		input = os.Args[3]
		banneFile = os.Args[4]
	}

	// add ".txt" to the bannerfile if missing
	if !strings.HasSuffix(banneFile, ".txt") {
		banneFile += ".txt"
	}

	// Read banner
	data, err := os.ReadFile(banneFile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	banner := string(data)
	banner = strings.ReplaceAll(banner, "\r", "")
	bannerTwo := strings.Split(banner, "\n")

	// Split input
	line := strings.Split(input, "\\n")

	for i, word := range line {
		if word == "" && i == 0 {
			continue
		}
		if word == "" {
			fmt.Println()
			continue
		}

		for row := 1; row < 9; row++ {
			for index, ascii := range word {
				if ascii < 32 || ascii > 126 {
					fmt.Println("user input out of range")
					return
				}

				start := (int(ascii) - 32) * 9

				if color != "" && substring != "" && shouldColor(word, substring, index) {
					fmt.Print(color + bannerTwo[start+row] + reset)
				} else if color != "" && substring == "" {
					fmt.Print(color + bannerTwo[start+row] + reset)
				} else {
					fmt.Print(bannerTwo[start+row])
				}
			}
			fmt.Println()
		}
	}
}

// func use for color
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
