package hangman

import(
	"testing"
)

func TestGame(t *testing.T) {
	game := New()
	totalTries := game.GetTotalTries()
	if totalTries != 0 {
		t.Error("Total tries must be equal 0")
	}
}

func TestTry(t *testing.T) {
	game := New()
	game.SetWord("Hello")
	game.Try("h")
	game.Try("i")
	game.Try("u")
	totalTries := game.GetTotalTries()

	if totalTries != 2 {
		t.Error("Only an error produced")
	}
}

func TestGetWord(t *testing.T) {
	game := New()
	game.SetWord(" HeLlO\n")
	if game.GetWord() != "hello" {
		t.Error("Word is different")
	}
}

func TestGetUncompleteWord(t *testing.T) {
	game := New()
	game.SetWord("test")
	game.Try("s")

	if game.GetUncompleteWord() != "__s_" {
		t.Error("Try haven't been submited")
	}
}

func TestHasFinished(t *testing.T) {
	game := New()
	game.SetWord("Hello")
	game.Try("h")
	game.Try("i")
	game.Try("u")
	if game.HasFinished() {
		t.Error("Wrong. Word \"Hello\" must be completed")
	}

	game.Try("e")
	game.Try("l")
	game.Try("o")
	if !game.HasFinished() {
		t.Error("Game must be finished")
	}
}