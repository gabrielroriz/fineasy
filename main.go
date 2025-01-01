package main

import (
	"bytes"
	"fineasy/database"
	"fineasy/handlers"
	"fineasy/utils"
	"fmt"
	"os"
)

// Definindo um tipo constante
type Option string

// Valores possíveis para o tipo Status
const OptionListFlow = "List Flows"
const OptionListCategories = "List Categorias"
const OptionListWallets = "List Wallets"
const OptionListSources = "List Sources"
const OptionInsertFlow = "Insert Flow"
const OptionInsertCategory = "Insert Categoria"
const OptionInsertWallet = "Insert Wallet"
const OptionInsertSource = "Insert Source"
const OptionGoBack = "Go Back"

func main() {

	var db *database.DBConfig
	var err error

	for db == nil {
		db, err = database.InitDB()
		if err != nil {
			fmt.Println(err)
			database.SetDBConfig(handlers.InsertDBConfig())
		} else {
			defer db.DB.Close()
		}
	}

	// Debug Mode:
	// handlers.InsertFlow()

	command := ""
	for command != "\\q" {
		fmt.Print("\nfineasy> ")
		fmt.Scanf("%s", &command)

		switch command {
		case "opt":
			commandOptions()
		}
	}
}

func commandOptions() {
	allOptions := []string{
		OptionListFlow,
		OptionListCategories,
		OptionListWallets,
		OptionListSources,
		OptionInsertFlow,
		OptionInsertCategory,
		OptionInsertWallet,
		OptionInsertSource,
		OptionGoBack,
	}

	selected := *showMenu(allOptions)

	switch selected {

	// Lists
	case OptionListFlow:
		handlers.ListFlows()

	case OptionListCategories:
		handlers.ListCategories()

	case OptionListWallets:
		handlers.ListWallets()

	case OptionListSources:
		handlers.ListSources()

	// Insertions
	case OptionInsertFlow:
		handlers.InsertFlow()

	case OptionInsertCategory:
		handlers.InsertCategory()

	case OptionInsertWallet:
		handlers.InsertWallet()

	case OptionInsertSource:
		handlers.InsertSource()

	default:

	}
}

func showMenu(options []string) *string {
	KeyArrowDown := []byte{27, 91, 66}
	KeyArrowUp := []byte{27, 91, 65}
	KeyEnter := []byte{10, 0, 0}

	selected := 0

	// Entrar no modo de entrada de terminal raw
	oldState, err := utils.TerminalModoRaw()
	if err != nil {
		fmt.Println("Erro ao configurar o terminal:", err)
		return nil
	}

	defer utils.TerminalRestore(oldState)

	for {
		// Limpa a tela
		utils.TerminalClearScreen()

		// Exibe o menu
		fmt.Println("Use as setas para navegar e Enter para selecionar:")
		for i, option := range options {
			makeBold := "\033[1m"
			makeGreen := "\033[32m>"
			if i == selected {
				fmt.Printf("%s%s %s\033[0m\n", makeBold, makeGreen, option) // Destaca a opção selecionada
			} else {
				fmt.Printf("  %s\n", option)
			}
		}

		// Captura a entrada do usuário
		input := make([]byte, 3)
		_, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			return nil
		}

		// Processa as teclas
		if bytes.Equal(input, KeyArrowUp) { // Seta para cima
			if selected > 0 {
				selected--
			} else if selected == 0 {
				selected = len(options) - 1
			}

		} else if bytes.Equal(input, KeyArrowDown) { // Seta para baixo
			if selected < len(options)-1 {
				selected++
			} else if selected == len(options)-1 {
				selected = 0
			}

		} else if bytes.Equal(input, KeyEnter) { // Enter
			return &options[selected]
		}
	}
}
