package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	tb "github.com/nsf/termbox-go"
)

// TODO: incorporate some sort of timing

//   - we actually need to use unicode since termbox returns unicode.
//     luckily, the first 127 characters of unicode == first 127 of ASCII
const (
	asciiLow, asciiHigh           = 33, 127
	numLow, numHigh               = 48, 57
	upperAlphaLow, upperAlphaHigh = 65, 90
	lowerAlphaLow, lowerAlphaHigh = 97, 122
)

func generateTarget(bank *[]rune, targetLen int, isVariableLen bool) ([]rune, int) {
	if isVariableLen {
		targetLen = rand.Intn(targetLen) + 1
	}
	target := make([]rune, 0, targetLen)

	for i := 0; i < targetLen; i++ {
		target = append(target, (*bank)[rand.Intn(len(*bank))])
	}
	return target, len(target)
}

func buildCharBank(mode, letterCase *string) *[]rune {
	var bank []rune
	if *mode == "fullASCII" {
		for i := asciiLow; i <= asciiHigh; i++ {
			bank = append(bank, rune(i))
		}
		return &bank
	}
	if *mode == "alpha" || *mode == "alphanum" {
		if *letterCase == "lower" || *letterCase == "mixed" {
			for i := lowerAlphaLow; i <= lowerAlphaHigh; i++ {
				bank = append(bank, rune(i))
			}
		}
		if *letterCase == "upper" || *letterCase == "mixed" {
			for i := upperAlphaLow; i <= lowerAlphaHigh; i++ {
				bank = append(bank, rune(i))
			}
		}
	}
	if *mode == "num" || *mode == "alphanum" {
		for i := numLow; i <= numHigh; i++ {
			bank = append(bank, rune(i))
		}
	}
	return &bank
}

func main() {

	modePtr := flag.String("mode", "num", "mode: {'alpha', 'num', 'alphanum', 'fullASCII'}")
	casePtr := flag.String("case", "mixed", "case: {'lower', 'upper', 'mixed'}")
	maxLenPtr := flag.Int("length", 1, "max target string length (or target string length if `variable` flag is False)")
	variablePtr := flag.Bool("variable", true, "generate variable length target string for each run over range [1, length]")
	flag.Parse()

	charBank := buildCharBank(modePtr, casePtr)

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
	var idx, hits, total int
	target, size := generateTarget(charBank, *maxLenPtr, *variablePtr)
	for {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\033[H\033[2J") // clear the terminal
		fmt.Println(string(target))
		event := tb.PollEvent() // blocking
		if event.Key == tb.KeyEsc {
			fmt.Print("\033[H\033[2J")
			fmt.Printf("%.2f accuracy %d/%d characters\n", float64(hits)/float64(total)*100.0, hits, total)
			time.Sleep(5 * time.Second) // give the user time to view the results before exiting.
			return
		}
		total++
		if event.Ch == target[idx] {
			hits++
			idx++
			if idx == size {
				// the full target has been matched. Reset the idx to 0 and generate a new target
				target, size = generateTarget(charBank, *maxLenPtr, *variablePtr)
				idx = 0
			}
		} else {
			fmt.Println("MISS!")
		}
	}
}
