package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func rgbDecToHex(rgb string) string {
	var r, g, b int
	fmt.Sscanf(rgb, "Color=%d,%d,%d", &r, &g, &b)
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: ", os.Args[0], "file.colorscheme")
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	colors := make(map[string]string)
	key := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			if line[0] == '[' {
				key = line
			} else if line[0] == 'C' {
				colors[key] = rgbDecToHex(line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	colorName := []string{"black", "red", "green", "yellow", "blue", "magenta", "cyan", "white"}
	fmt.Println("\t/* 8 normal colors */")
	for i := 0; i < 8; i++ {
		fmt.Printf("\t[%d] = \"%s\", /* %-7s */\n", i, colors["[Color"+strconv.Itoa(i)+"]"], colorName[i])
	}

	fmt.Println("\n\t/* 8 bright colors */")
	for i := 0; i < 8; i++ {
		fmt.Printf("\t[%2d] = \"%s\", /* %-7s */\n", i+8, colors["[Color"+strconv.Itoa(i)+"Intense]"], colorName[i])
	}

	fmt.Println("\n\t/* special colors */")
	fmt.Printf("\t[%d] = \"%s\", /* %s */\n", 256, colors["[Background]"], "background")
	fmt.Printf("\t[%d] = \"%s\", /* %s */\n", 257, colors["[Foreground]"], "foreground")
}
