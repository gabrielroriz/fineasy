package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

// Text Colors
const (
	TerminalTextColorRed     = "\033[31m"
	TerminalTextColorGreen   = "\033[32m"
	TerminalTextColorBlue    = "\033[34m"
	TerminalTextColorCyan    = "\033[36m"
	TerminalTextColorMagenta = "\033[35m"
	TerminalTextColorYellow  = "\033[33m"
	TerminalTextColorBlack   = "\033[30m"
	TerminalTextColorWhite   = "\033[37m"
)

// Text Styles
const (
	TerminalTextColorReset = "\033[0;0m" // Restore original color
	TerminalTextBold       = "\033[1m"
	TerminalTextReversed   = "\033[2m"
)

// Background Colors
const (
	TerminalBackgroundBlack   = "\033[40m"
	TerminalBackgroundRed     = "\033[41m"
	TerminalBackgroundGreen   = "\033[42m"
	TerminalBackgroundYellow  = "\033[43m"
	TerminalBackgroundBlue    = "\033[44m"
	TerminalBackgroundMagenta = "\033[45m"
	TerminalBackgroundCyan    = "\033[46m"
	TerminalBackgroundWhite   = "\033[47m"
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

/**
** Terminal Table UI
 */

func TerminalUIPrintTableModeOnlyView(keys []string, values [][]string) {
	printTable(keys, values, false, 0)
}

func printTable(keys []string, values [][]string, selectionMode bool, selected int) {

	tab := selectionMode

	// Size of each column
	sizes := make([]int, len(keys))

	// First reference of column size is key
	for column := 0; column < len(keys); column++ {
		sizes[column] = len(keys[column])
	}

	// Define sizes and give one space before each value
	for line := 0; line < len(values); line++ {
		row := values[line]

		for column := 0; column < len(row); column++ {

			// If there isn't a space before each value, add
			if rune(values[line][column][0]) != ' ' {
				// TODO: Do not modify 'values'
				values[line][column] = fmt.Sprintf(" %s", values[line][column])
			}

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

	// Draw first strong line
	fmt.Println(drawLine(treatLeftMostString("╒", tab, false), "═", "╕", totalSize))

	// Print Keys
	for i := 0; i < len(keys); i++ {

		if i == 0 {
			fmt.Print(treatLeftMostString("│", tab, false))
		}

		emptySpace := strings.Repeat(" ", (sizes[i]-len(keys[i]))/2)

		fmt.Print(fmt.Sprintf("%s%s%s%s%s", TerminalTextBold, emptySpace, keys[i], emptySpace, TerminalTextColorReset))

		if i == len(keys)-1 {
			fmt.Print("\n")
		}
	}

	// Draw second strong line
	fmt.Println(drawLine(treatLeftMostString("╞", tab, false), "═", "╡", totalSize))

	// Print Values
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[i]); j++ {
			if j == 0 {
				fmt.Print(treatLeftMostString("│", tab, selected == i))
			}

			if i == selected && selectionMode {
				fmt.Print(fmt.Sprintf("%s%s%-*s%s│", TerminalTextBold, TerminalBackgroundGreen, sizes[j] /* Dinamical size */, values[i][j], TerminalTextColorReset))
			} else {
				fmt.Print(fmt.Sprintf("%-*s│", sizes[j] /* Dinamical size */, values[i][j]))
			}

			if j == len(values[i])-1 {
				fmt.Print("\n")
			}
		}

	}

	// Draw last line
	fmt.Println(drawLine(treatLeftMostString("└", tab, false), "─", "┘", totalSize))
}

func TerminalUIPrintTableSelectionMode(keys []string, values [][]string) *string {
	keyArrowDown := []byte{27, 91, 66}
	keyArrowUp := []byte{27, 91, 65}
	keyEnter := []byte{10, 0, 0}

	selected := 0
	options := values[0]

	// Terminal Raw Mode activate
	oldState, err := TerminalModoRaw()
	if err != nil {
		fmt.Println("Erro ao configurar o terminal:", err)
		return nil
	}

	defer TerminalRestore(oldState)

	for {
		// Limpa a tela
		TerminalClearScreen()

		// Exibe o menu
		fmt.Println("Use as setas para navegar e Enter para selecionar:")

		printTable(keys, values, true, selected)

		// Captura a entrada do usuário
		input := make([]byte, 3)
		_, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			return nil
		}

		// Processa as teclas
		if bytes.Equal(input, keyArrowUp) { // Seta para cima
			if selected > 0 {
				selected--
			} else if selected == 0 {
				selected = len(options) - 1
			}

		} else if bytes.Equal(input, keyArrowDown) { // Seta para baixo
			if selected < len(options)-1 {
				selected++
			} else if selected == len(options)-1 {
				selected = 0
			}

		} else if bytes.Equal(input, keyEnter) { // Enter
			return &options[selected]
		}
	}
}

func drawLine(leftEdge string, common string, rightEdge string, totalSize int) string {
	return fmt.Sprintf("%s%s%s", leftEdge, fmt.Sprintf("%s", strings.Repeat(common, totalSize-2)), rightEdge)
}

func spaces(count int) string {
	return fmt.Sprintf("%*s", count, "")
}

func treatLeftMostString(leftmostString string, tab bool, selected bool) string {
	if tab {
		tabWidth := 8 // Default tab width

		char := " "
		if selected {
			char = fmt.Sprintf("%s%s>%s", TerminalTextBold, TerminalTextColorGreen, TerminalTextColorReset) // Character to place in the middle
		}
		padding := (tabWidth / 2)
		return fmt.Sprintf("%s%s%s%s", spaces(padding-1), char, spaces(tabWidth-padding), leftmostString)
	} else {
		return leftmostString
	}
}
