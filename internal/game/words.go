package game

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

// WordsManager zarządza słowami do gry
type WordsManager struct {
	words []string
}

// NewWordsManager tworzy nowy manager słów
func NewWordsManager(filePath string) (*WordsManager, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()
		if word != "" {
			words = append(words, strings.ToLower(strings.TrimSpace(word)))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &WordsManager{words: words}, nil
}

// GetRandomWord zwraca losowe słowo z listy
func (wm *WordsManager) GetRandomWord() string {
	rand.Seed(time.Now().UnixNano())
	return wm.words[rand.Intn(len(wm.words))]
}

// ContainsPolishChars sprawdza czy słowo zawiera polskie znaki
func ContainsPolishChars(word string) bool {
	polishChars := []rune{'ą', 'ć', 'ę', 'ł', 'ń', 'ó', 'ś', 'ź', 'ż'}
	for _, char := range word {
		for _, polish := range polishChars {
			if char == polish {
				return true
			}
		}
	}
	return false
}

// NormalizeGuess normalizuje literę wprowadzoną przez użytkownika
func NormalizeGuess(guess rune) rune {
	normalized := unicode.ToLower(guess)

	// Mapowanie polskich znaków na podstawowe
	switch normalized {
	case 'ą':
		return 'a'
	case 'ć':
		return 'c'
	case 'ę':
		return 'e'
	case 'ł':
		return 'l'
	case 'ń':
		return 'n'
	case 'ó':
		return 'o'
	case 'ś':
		return 's'
	case 'ź', 'ż':
		return 'z'
	default:
		return normalized
	}
}

// IsPolishLetter sprawdza czy znak jest polską literą
func IsPolishLetter(r rune) bool {
	return unicode.IsLetter(r) || r == 'ą' || r == 'ć' || r == 'ę' || r == 'ł' ||
		r == 'ń' || r == 'ó' || r == 'ś' || r == 'ź' || r == 'ż'
}
