package main

import (
	"flag"
	"fmt"
	"github.com/ptgoetz/go-ansi/pkg/ansi"
	"golang.org/x/sys/unix"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type config struct {
	// flags

	// params
	delay         time.Duration
	remainingArgs []string
	runes         string
}

var args config

func parseArgs() config {
	args = config{}
	flag.DurationVar(&args.delay, "delay", 25*time.Millisecond, "The delay between console updates.")
	flag.StringVar(&args.runes, "charset", "123ABC", "Character to use for display.")

	log.Printf("Parsing args : %v", os.Args)
	flag.Parse()
	args.remainingArgs = flag.Args()

	return args
}

func main() {
	config := parseArgs()
	rand.Seed(time.Now().UnixNano())

	console := ansi.NewConsole(os.Stdout)

	rows, cols, _ := getConsoleDimensions()
	_ = initWindow(console, rows, cols)

	signalInterrupt := make(chan os.Signal, 1)
	defer close(signalInterrupt)
	signal.Notify(signalInterrupt, syscall.SIGINT)

	signalWinSizeChange := make(chan os.Signal, 1)
	defer close(signalWinSizeChange)
	signal.Notify(signalWinSizeChange, syscall.SIGWINCH)

	resizeChan := make(chan [2]int, 1)

	// SIGWINCH, SIGINT goroutine
	go func() {
		for {
			select {
			case <-signalWinSizeChange:
				rows, cols, _ := getConsoleDimensions()
				resizeChan <- [2]int{rows, cols}
				//_ = initWindow(console, rows, cols)
			case <-signalInterrupt:
				console.Reset()
				os.Exit(0)
				//default:
			}
		}
	}()

	go func() {
		for {
			select {
			case dims := <-resizeChan:
				rows = dims[0]
				cols = dims[1]
				console.Reset()
				initWindow(console, rows, cols)
			default:
			}
			// create a random row/col position
			randRow := rand.Intn(rows)
			randCol := rand.Intn(cols)
			console.SetCursorPos(randRow, randCol)

			// create a random rgb color
			r := rand.Intn(255)
			g := rand.Intn(255)
			b := rand.Intn(255)

			// set the cursor position and background color
			console.SetBackground(r, g, b)
			console.SetForeground(255-r, 255-g, 255-b)
			console.Print(RandomRune(config.runes))

			// sleep a bit
			time.Sleep(config.delay)
		}

	}()

	for {
		//time.Sleep(time.Hour)
	}
}

func getConsoleDimensions() (rows int, cols int, err error) {
	ws, err := unix.IoctlGetWinsize(syscall.Stdout, unix.TIOCGWINSZ)
	if err != nil {
		return 0, 0, err
	}
	cols = int(ws.Col) + 1
	rows = int(ws.Row) + 1
	return rows, cols, nil
}

func initWindow(console *ansi.Console, rows int, cols int) error {
	console.ResetStyle()
	console.ClearScreen()
	//console.SetCursorPos(0, 0)
	console.SetCursorVisible(false)
	//fmt.Printf("Rows: %d, Cols: %d", rows, cols)

	fmt.Print(ansi.SetCursorPos(rows, 0)) // move the cursor to the bottom of the screen
	return nil
}

//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
//var letterRunes = []rune("01")
//var letterRunes = []rune("0123456789!@#$%^&*()?><~ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomRune(charset string) string {
	return string(charset[rand.Intn(len(charset))])
}
