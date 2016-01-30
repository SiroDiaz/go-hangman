Hangman
=======

Go-hangman is a simple package containing the game module that gets words from an external service API.

### Game

to init the game you must create a main package with the following
logic(You can modify messages):

~~~ go
package main

import "github.com/SiroDiaz/go-hangman"
import "fmt"
import "strings"
import "strconv"

func main() {
	defer func () {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	game := hangman.New()
	var letter string
	
	for !game.HasFinished() {
		fmt.Print("Insert a character: ")
		fmt.Scanf("%s", &letter)
		letter = strings.TrimSpace(letter)
		if len(letter) > 1 {
			panic("Letter must be one character, no more")
		}

		game.Try(letter)
		game.PrintUncompleteWord()
	}

	if game.GetTotalTries() < hangman.GetMaxTries() {
		fmt.Println("You won! you got: "+
			strconv.Itoa(game.GetTotalTries()) +
			" fails")
	} else {
		fmt.Println("You lose :(. The word is: "+ game.GetWord())
	}
}
~~~


