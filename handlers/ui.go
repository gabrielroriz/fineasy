package handlers

import (
	"fmt"
	"strings"
)

func PrintBold(str string) {
	fmt.Printf("\033[1m%s\033[0m", str)
}

func PrintSuccess(str string) {
	fmt.Printf("\033[32;1m%s\033[0m", str)
}

func PrintTable(keys []string, values [][]string) {

	sizes := make([]int, len(keys))

	//define sizes
	for line := 0; line < len(values); line++ {

		row := values[line]
		for column := 0; column < len(row); column++ {
			if sizes[column] < len(row[column]) {
				sizes[column] = len(row[column]) + 3
			}

		}
	}

	var totalSize int
	for i := 0; i < len(sizes); i++ {
		totalSize += sizes[i]
	}

	totalSize += len(keys) + 1

	normalLine := fmt.Sprintf("%s", strings.Repeat("-", totalSize))
	strongLine := fmt.Sprintf("%s", strings.Repeat("=", totalSize))

	//header table
	fmt.Println(strongLine)

	for i := 0; i < len(keys); i++ {

		if i == 0 {
			fmt.Print("|")
		}

		fmt.Print(fmt.Sprintf("\033[1m%-*s\033[0m|", sizes[i], keys[i]))

		if i == len(keys)-1 {
			fmt.Print("\n")
		}
	}

	fmt.Println(strongLine)

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			if j == 0 {
				fmt.Print("|")
			}

			fmt.Print(fmt.Sprintf("%-*s|", sizes[j], values[i][j]))

			if j == len(values[i])-1 {
				fmt.Print("\n")
			}
		}

	}

	fmt.Println(normalLine)

	// fmt.Println(sizes)
}
