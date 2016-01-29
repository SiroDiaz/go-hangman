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

func GetMaxTries() int {
	return max_tries
}

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

// 
func (g game) GetWord() string {
	return g.completeWord
}

func (g game) GetTotalTries() int {
	return g.tries
}

// 
func (g *game) GetUncompleteWord() string {
	return string(g.userWord)
}

// 
func (g *game) HasFinished() bool {
	return g.tries >= max_tries || !strings.Contains(string(g.userWord), string(separator))
}

// 
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

// 
func (g *game) PrintUncompleteWord() {
	for _, v := range string(g.userWord) {
		fmt.Print(string(v) +" ")
	}
	fmt.Println()
}
