package utils

import (
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func TerminalClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

// Set Raw Mode on terminal
// "O modo raw (raw mode) no terminal é um estado em que a entrada do teclado é transmitida diretamente para o programa sem qualquer processamento intermediário"
func TerminalModoRaw() (*unix.Termios, error) {
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
func TerminalRestore(oldState *unix.Termios) {
	fd := int(os.Stdin.Fd())
	_ = unix.IoctlSetTermios(fd, unix.TCSETS, oldState)
}
