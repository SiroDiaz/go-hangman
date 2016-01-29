package hangman

import(
	"net/http"
	"io/ioutil"
	"strings"
	"fmt"
)

const max_tries = 6
const url = "http://randomword.setgetgo.com/get.php"
const separator = '_'

type game struct {
	tries int
	completeWord string
	userWord []byte
}

// initialize a new game and return a pointer of the structure.
func New() *game {
	game := new(game)
	game.tries = 0
	game.completeWord = getRandomWord()
	game.userWord = make([]byte, len(game.completeWord))
	for i := 0; i < len(game.completeWord); i++ {
		game.userWord[i] = separator
	}

	return game
}

// Obtain a random word from a service.
func getRandomWord() string {
	res, err := http.Get(url)
	if err != nil {
		panic("Error requesting the page: "+ url)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic("Error reading the body response")
	}

	return strings.ToLower(string(body))
}


// Returns maximum tries that user can do.
func GetMaxTries() int {
	return max_tries
}

// Returns the word to be guessed
func (g game) GetWord() string {
	return g.completeWord
}

// Returns the total of tries that user has done.
func (g game) GetTotalTries() int {
	return g.tries
}

// Returns the current status word with underscores
// in characters that haven't been found.
func (g *game) GetUncompleteWord() string {
	return string(g.userWord)
}

// Sets the word to be guessed.
func (g *game) SetWord(word string) {
	word = strings.TrimSpace(word)
	if len(word) == 0 {
		panic("Empty word used")
	}

	g.completeWord = strings.ToLower(word)
	g.userWord = make([]byte, len(word))
	for i := 0; i < len(word); i++ {
		g.userWord[i] = separator
	}
}

// Returns true if the number of tries has been overcome
// or if the word has been completed.
func (g *game) HasFinished() bool {
	return g.tries >= max_tries || !strings.Contains(string(g.userWord), string(separator))
}

// Makes a new try. If the letter isn't contained then
// increments the number of tries.
func (g *game) Try(letter string) {
	if strings.Contains(g.completeWord, letter) {
		letterChar := []byte(letter)[0]
		for i := 0; i < len(g.userWord); i++ {
			if g.completeWord[i] == letterChar {
				g.userWord[i] = letterChar
			}
		}
	} else {
		g.tries++
	}
}

// Util method that prints with format the uncompleted word.
func (g game) PrintUncompleteWord() {
	for _, v := range string(g.userWord) {
		fmt.Print(string(v) +" ")
	}
	fmt.Println()
}
