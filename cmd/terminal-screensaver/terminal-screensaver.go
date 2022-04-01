package main

import (
	"fmt"
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"golang.org/x/sys/unix"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	writer := ansi.NewWriter(os.Stdout)

	//defer writer.ClearScreen()

	rows, cols, _ := getConsoleDimensions()
	_ = updateWindowSize(writer, rows, cols)
	//
	signalInterrupt := make(chan os.Signal, 1)
	defer close(signalInterrupt)
	signal.Notify(signalInterrupt, syscall.SIGINT)

	signalWinSizeChange := make(chan os.Signal, 1)
	defer close(signalWinSizeChange)
	signal.Notify(signalWinSizeChange, syscall.SIGWINCH)

	go func() {
		for {
			select {
			case <-signalWinSizeChange:
				rows, cols, _ := getConsoleDimensions()
				_ = updateWindowSize(writer, rows, cols)
			case <-signalInterrupt:
				resetConsole(writer)
				os.Exit(0)
				//default:
			}
		}
	}()

	go func() {
		for {
			// create a random row/col position
			randRow := rand.Intn(rows)
			randCol := rand.Intn(cols)
			writer.SetCursorPos(randRow, randCol)

			// create a random rgb color
			r := rand.Intn(255)
			g := rand.Intn(255)
			b := rand.Intn(255)

			// set the cursor position and background color
			writer.SetBackground(r, g, b)
			writer.SetForeground(255-r, 255-g, 255-b)
			writer.Print(RandomRune())

			// sleep a bit
			time.Sleep(25 * time.Millisecond)
		}

	}()

	for {
		time.Sleep(time.Hour)
	}
}

func resetConsole(writer *ansi.Writer) {
	writer.ResetStyle()
	writer.ClearScreen()
	writer.SetCursorPos(0, 0)
	writer.SetCursorVisible(true)

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

func updateWindowSize(writer *ansi.Writer, rows int, cols int) error {
	writer.ResetStyle()
	writer.ClearScreen()
	//writer.SetCursorPos(0, 0)
	writer.SetCursorVisible(false)
	//fmt.Printf("Rows: %d, Cols: %d", rows, cols)

	fmt.Print(ansi.SetCursorPos(rows, 0)) // move the cursor to the bottom of the screen
	return nil
}

//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//var letterRunes = []rune("01")
var letterRunes = []rune("0123456789!@#$%^&*()?><~ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func RandomRune() string {

	return string(letterRunes[rand.Intn(len(letterRunes))])
}
