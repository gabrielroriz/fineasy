package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func main() {

	// var db *database.DBConfig
	// var err error

	// for db == nil {
	// 	db, err = database.InitDB()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		database.SetDBConfig(handlers.InsertDBConfig())
	// 	} else {
	// 		defer db.DB.Close()
	// 	}
	// }

	// command := ""

	// // Stops waiting when user inserts "\q"
	// for command != "\\q" {

	// 	fmt.Print("\nfineasy> ")
	// 	fmt.Scanf("%s", &command)

	// 	switch command {

	// 	// Lists
	// 	case "lf":
	// 		handlers.ListFlows()

	// 	case "lc":
	// 		handlers.ListCategories()

	// 	case "lw":
	// 		handlers.ListWallets()

	// 	case "ls":
	// 		handlers.ListSources()

	// 	// Insertions
	// 	case "if":
	// 		handlers.InsertFlow()

	// 	case "ic":
	// 		handlers.InsertCategory()

	// 	case "iw":
	// 		handlers.InsertWallet()

	// 	case "is":
	// 		handlers.InsertSource()

	// 	default:

	// 	}

	// }

	options := []string{"Listar Flows", "Listar Categorias", "Listar Wallets", "Sair"}
	selected := 0

	// Entrar no modo de entrada de terminal raw
	oldState, err := makeRaw()
	if err != nil {
		fmt.Println("Erro ao configurar o terminal:", err)
		return
	}
	defer restoreTerminal(oldState)

	for {
		// Limpa a tela
		clearScreen()

		// Exibe o menu
		fmt.Println("Use as setas para navegar e Enter para selecionar:")
		for i, option := range options {
			if i == selected {
				fmt.Printf("\033[32m> %s\033[0m\n", option) // Destaca a opção selecionada
			} else {
				fmt.Printf("  %s\n", option)
			}
		}

		// Captura a entrada do usuário
		input := make([]byte, 3)
		_, err := os.Stdin.Read(input)
		if err != nil {
			fmt.Println("Erro ao ler entrada:", err)
			return
		}

		// fmt.Printf("%d", input[0])
		// time.Sleep(3 * time.Second)

		// Processa as teclas
		if bytes.Equal(input, []byte{27, 91, 65}) { // Seta para cima
			if selected > 0 {
				selected--
			}
		} else if bytes.Equal(input, []byte{27, 91, 66}) { // Seta para baixo
			if selected < len(options)-1 {
				selected++
			}
		} else if bytes.Equal(input, []byte{10, 0, 0}) { // Enter
			fmt.Printf("Você selecionou: %s\n", options[selected])
			return
		}
	}
}

// Limpa a tela do terminal
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// Set Raw Mode on terminal
// "O modo raw (raw mode) no terminal é um estado em que a entrada do teclado é transmitida diretamente para o programa sem qualquer processamento intermediário"
func makeRaw() (*unix.Termios, error) {
	fd := int(os.Stdin.Fd())
	oldState, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		return nil, err
	}
	newState := *oldState
	newState.Lflag &^= unix.ICANON | unix.ECHO
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, &newState); err != nil {
		return nil, err
	}
	return oldState, nil
}

// Restaura o estado anterior do terminal
func restoreTerminal(oldState *unix.Termios) {
	fd := int(os.Stdin.Fd())
	_ = unix.IoctlSetTermios(fd, unix.TCSETS, oldState)
}
