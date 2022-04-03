package main

import (
	"fmt"
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	writer := ansi.NewConsole(os.Stdout)
	_ = updateWindowSize(writer)
	//
	sigwinch := make(chan os.Signal, 1) // Set signal handler
	defer close(sigwinch)
	signal.Notify(sigwinch, syscall.SIGWINCH)
	go func() {
		for {
			select {
			case <-sigwinch:
				updateWindowSize(writer)
				//default:
			}
		}
	}()

	for {
		time.Sleep(time.Hour)
	}
}

func getConsoleDimensions() (rows int, cols int, err error) {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	cols = int(ws.Col)
	rows = int(ws.Row)
	return rows, cols, nil
}

func updateWindowSize(writer *ansi.Console) error {
	writer.ClearScreen()
	writer.SetCursorPos(0, 0)
	rows, cols, err := getConsoleDimensions()
	if err != nil {
		return err
	}

	fmt.Printf("Rows: %d, Cols: %d", rows, cols)

	fmt.Print(ansi.SetCursorPos(rows, 0)) // move the cursor to the bottom of the screen
	return nil
}
