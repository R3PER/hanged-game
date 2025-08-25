package game

import (
	"strings"
)

// Poziomy trudności
const (
	EasyLevel   = 8 // 8 prób
	MediumLevel = 6 // 6 prób
	HardLevel   = 4 // 4 próby
)

// GameState reprezentuje stan gry
type GameState int

// Możliwe stany gry
const (
	Playing GameState = iota
	Won
	Lost
)

// Game reprezentuje pojedynczą rozgrywkę
type Game struct {
	Word           string    // Słowo do odgadnięcia
	GuessedLetters []rune    // Odgadnięte litery
	WrongGuesses   []rune    // Błędne próby
	MaxAttempts    int       // Maksymalna liczba prób
	Points         int       // Punkty zdobyte w grze
	State          GameState // Aktualny stan gry
}

// NewGame tworzy nową grę
func NewGame(word string, difficultyLevel int) *Game {
	maxAttempts := MediumLevel // Domyślnie średni poziom trudności

	// Ustawienie poziomu trudności
	switch difficultyLevel {
	case 1:
		maxAttempts = EasyLevel
	case 2:
		maxAttempts = MediumLevel
	case 3:
		maxAttempts = HardLevel
	}

	return &Game{
		Word:           strings.ToLower(word),
		GuessedLetters: []rune{},
		WrongGuesses:   []rune{},
		MaxAttempts:    maxAttempts,
		Points:         0,
		State:          Playing,
	}
}

// GetWordWithGuesses zwraca słowo z widocznymi odgadniętymi literami
func (g *Game) GetWordWithGuesses() string {
	result := ""
	for _, char := range g.Word {
		if g.isGuessed(char) {
			result += string(char)
		} else {
			result += "_"
		}
		result += " "
	}
	return strings.TrimSpace(result)
}

// isGuessed sprawdza czy litera została już odgadnięta
func (g *Game) isGuessed(letter rune) bool {
	normalizedLetter := NormalizeGuess(letter)
	for _, guessed := range g.GuessedLetters {
		if NormalizeGuess(guessed) == normalizedLetter {
			return true
		}
	}
	return false
}

// isWrongGuess sprawdza czy litera znajduje się w liście błędnych prób
func (g *Game) isWrongGuess(letter rune) bool {
	normalizedLetter := NormalizeGuess(letter)
	for _, wrong := range g.WrongGuesses {
		if NormalizeGuess(wrong) == normalizedLetter {
			return true
		}
	}
	return false
}

// Guess dokonuje próby odgadnięcia litery
func (g *Game) Guess(letter rune) bool {
	// Jeśli gra się skończyła, zwróć false
	if g.State != Playing {
		return false
	}

	// Jeśli litera została już odgadnięta lub jest błędną próbą, zwróć false
	normalizedLetter := NormalizeGuess(letter)
	if g.isGuessed(normalizedLetter) || g.isWrongGuess(normalizedLetter) {
		return false
	}

	// Sprawdź czy litera znajduje się w słowie
	letterInWord := false
	for _, char := range g.Word {
		if NormalizeGuess(char) == normalizedLetter {
			letterInWord = true
			break
		}
	}

	if letterInWord {
		g.GuessedLetters = append(g.GuessedLetters, letter)

		// Dodaj punkty za odgadniętą literę
		g.Points += 10

		// Sprawdź czy wszystkie litery zostały odgadnięte
		allGuessed := true
		for _, char := range g.Word {
			if !g.isGuessed(char) {
				allGuessed = false
				break
			}
		}

		if allGuessed {
			g.State = Won
			// Bonus za wygraną
			g.Points += 50

			// Bonus za pozostałe próby
			remainingAttempts := g.MaxAttempts - len(g.WrongGuesses)
			g.Points += remainingAttempts * 5
		}
	} else {
		g.WrongGuesses = append(g.WrongGuesses, letter)

		// Odejmij punkty za błędną próbę
		g.Points -= 5

		// Sprawdź czy przekroczono maksymalną liczbę prób
		if len(g.WrongGuesses) >= g.MaxAttempts {
			g.State = Lost
		}
	}

	return true
}

// GetRemainingAttempts zwraca liczbę pozostałych prób
func (g *Game) GetRemainingAttempts() int {
	return g.MaxAttempts - len(g.WrongGuesses)
}

// GetWrongGuesses zwraca listę błędnych prób
func (g *Game) GetWrongGuesses() string {
	var result strings.Builder
	for _, letter := range g.WrongGuesses {
		result.WriteRune(letter)
		result.WriteRune(' ')
	}
	return strings.TrimSpace(result.String())
}

// GetProgress zwraca procentowy postęp odgadnięcia słowa
func (g *Game) GetProgress() float64 {
	totalLetters := 0
	uniqueLetters := make(map[rune]bool)

	for _, char := range g.Word {
		if IsPolishLetter(char) {
			totalLetters++
			uniqueLetters[NormalizeGuess(char)] = true
		}
	}

	if totalLetters == 0 {
		return 0
	}

	guessedLetters := 0
	for _, char := range g.Word {
		if g.isGuessed(char) {
			guessedLetters++
		}
	}

	return float64(guessedLetters) / float64(totalLetters) * 100
}
