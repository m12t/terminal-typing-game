# Terminal Typing Game
This is a humble typing game I made to help me get better at touch typing. But since it's public, here's a fittingly small description for this small program:

## Overview
- A "target" will pop up on the terminal, you have to type it out as quickly and accurately as possible. While you're not timed by the game, your accuracy is tracked and you'll get a message `MISS!` in red letters whenever you miss a character in the target.
- Missing doesn't advance the cursor, which means you'll stay on a letter until you get it right. There's also no going backwards to redo a line. The cursor only increases, and the only way to move the cursor is by typing the correct character.
- After you successfully type out a target, a new one will immediately replace it for the cycle to repeat over and over until you exit the game by pressing the `esc` character. Upon exiting, you'll see your accuracy percentage, as well as the raw number of hits and total characters attempted over the course of the session.


## Basic Usage
1. clone the repo
2. cd into it
3. run `go run main.go` and optionally invoke your [configuration](#configuration) of choice.
    - For example, a "hard mode" might look like:
        - `go run main.go -mode=fullASCII -case=mixed -length=15 -variable=true`
    - If you only want to practice your numbers, below is a nice zen-like choice:
        - `go run main.go -mode=num -length=1`

## Configuration
1. There are several command-line arguments you can give the file upon execution that will dictate its behavior:
    - `mode` one of `{'alpha', 'num', 'alphanum', 'fullASCII'}`, default `alphanum`
        - `mode` determines the character bank that your targets will come from:
            - `alpha`: letters `[a-z]` (the letter case is modified by another flag, `case`)
            - `alphanum`: letters and numbers `[a-z0-9]`
            - `num`: numbers only `[0-9]`
            - `fullASCII`: letters, numbers, and symbols. For an exact list, it's [ASCII characters 33-126](https://www.cs.cmu.edu/~pattis/15-1XX/common/handouts/ascii.html)
    - `case` one of `{'lower', 'upper', 'mixed'}`, default `lower`
        - `case` modifies the letters in `mode`, if alpha characters were chosen
            - `lower`: lowercase letters only
            - `upper`: uppercase letters only
            - `mixed`: both lowercase and uppercase letters
    - `length` (integer), default `7`
        - `length` determines the max length of the target. It's modified by `variable`
    - `variable` (boolean), default `true`
        - `variable` determines whether the target length can vary randomly between runs. If `true`, length can vary between `[1, length]` else, the target's length will be constantly `length` characters for every run
