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
	writer := ansi.NewWriter(os.Stdout)
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

func updateWindowSize(writer *ansi.Writer) error {
	writer.ClearScreen()
	writer.SetCursorPos(0, 0)
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return err
	}

	cols := int(ws.Col)
	rows := int(ws.Row)
	fmt.Printf("Rows: %d, Cols: %d", rows, cols)

	fmt.Print(ansi.SetCursorPos(rows, 0)) // move the cursor to the bottom of the screen
	return nil
}
