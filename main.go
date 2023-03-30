package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	tb "github.com/nsf/termbox-go"
)

func main() {
	err := tb.Init()
	if err != nil {
		panic(err)
	}
	defer tb.Close()

	fmt.Println("Welcome!\n>>> The goal of the game is to type the number you see")
	fmt.Println("    appear on screen as fast as you can while still")
	fmt.Println("    as accurate as possible!")
	fmt.Println("\n>>> Your overall accuracy is show when you quit.")
	fmt.Println("\n>>> To end the game, simply press `esc` at any time.")
	fmt.Println("\n\t\tPress any key to begin!")
	tb.PollEvent()

	rand.Seed(time.Now().UnixNano())
	var hits, total int
	target := strconv.Itoa(rand.Intn(10))
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Print("\033[H\033[2J") // clear the terminal
		fmt.Println(target)
		event := tb.PollEvent() // blocking
		if event.Key == tb.KeyEsc {
			fmt.Print("\033[H\033[2J")
			fmt.Printf("%.2f accuracy %d/%d\n", float64(hits)/float64(total)*100.0, hits, total)
			time.Sleep(5 * time.Second) // give the user time to view the results before exiting.
			return
		}
		total++
		if string(event.Ch) == target {
			hits++
			target = strconv.Itoa(rand.Intn(10))
		} else {
			fmt.Println("MISS!")
		}
	}
}
