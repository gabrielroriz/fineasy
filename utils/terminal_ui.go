package utils

import (
	"fmt"
	"strings"
)

func TerminalUIPrintBold(str string) {
	fmt.Printf("\033[1m%s\033[0m", str)
}

func MakeBold(str string, a ...any) string {
	return fmt.Sprintf("\033[1m%s\033[0m", fmt.Sprintf(str, a...))
}

func TerminalUIPrintGreen(str string) {
	fmt.Printf("\033[32;1m%s\033[0m", str)
}

func TerminalPrintOnSameLine(str string, a ...any) {
	fmt.Printf("\033[1A\033[K%s\033[0m", fmt.Sprintf(str, a...))
}

func TerminalUIPrintTable(keys []string, values [][]string) {

	// Size of each column
	sizes := make([]int, len(keys))

	// Define sizes and give one space before each value
	for line := 0; line < len(values); line++ {
		row := values[line]

		// First reference of column size is key
		for column := 0; column < len(row); column++ {
			sizes[column] = len(keys[column])
		}

		for column := 0; column < len(row); column++ {
			// Add a space before each value
			values[line][column] = fmt.Sprintf(" %s", values[line][column])

			// Find largest value
			if sizes[column] < len(row[column]) {
				sizes[column] = len(row[column]) + 2
			}

		}
	}

	var totalSize int

	for i := 0; i < len(sizes); i++ {

		// To centeralize column title
		diff := sizes[i] - len(keys[i])
		if diff%2 != 0 {
			sizes[i]++
		}

		totalSize += sizes[i]

	}

	totalSize += len(keys) + 1

	//draw first strong line
	fmt.Println(drawLine("╒", "═", "╕", '╤', totalSize, sizes))

	for i := 0; i < len(keys); i++ {

		if i == 0 {
			fmt.Print("│")
		}

		// fmt.Printf("\nSizes[i]: %d", sizes[i])
		// fmt.Printf("\nLen(Keys[i]): %d", len(keys[i]))
		// time.Sleep(time.Second * 3)

		emptySpace := strings.Repeat(" ", (sizes[i]-len(keys[i]))/2)

		fmt.Print(fmt.Sprintf("\033[1m%s%s%s\033[0m│", emptySpace, keys[i], emptySpace))

		if i == len(keys)-1 {
			fmt.Print("\n")
		}
	}

	// Draw second strong line
	fmt.Println(drawLine("╞", "═", "╡", '╪', totalSize, sizes))

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			if j == 0 {
				fmt.Print("│")
			}

			fmt.Print(fmt.Sprintf("%-*s│", sizes[j], values[i][j]))

			if j == len(values[i])-1 {
				fmt.Print("\n")
			}
		}

	}

	// Draw last line
	fmt.Println(drawLine("└", "─", "┘", '┴', totalSize, sizes))
}

func drawLine(leftEdge string, common string, rightEdge string, division rune, totalSize int, sizes []int) string {

	line := fmt.Sprintf("%s%s%s", leftEdge, fmt.Sprintf("%s", strings.Repeat(common, totalSize-2)), rightEdge)

	startingAt := 1

	for i := 0; i < len(sizes); i++ {
		if i != len(sizes)-1 {
			line = replaceAtIndex(line, division, startingAt+sizes[i])
			startingAt += sizes[i] + 1
		}
	}

	return line
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
