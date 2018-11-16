package handlers

import (
	"fmt"
)

func PrintBold(str string) {
	fmt.Printf("\033[1m%s\033[0m", str)
}

func PrintSuccess(str string) {
	fmt.Printf("\033[32;1m%s\033[0m", str)
}
