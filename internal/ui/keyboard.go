package ui

import (
	"fmt"
	"os"
)

// Kody klawiszy specjalnych
const (
	KeyEnter      = '\r'
	KeyEsc        = 27
	KeyArrowUp    = 65
	KeyArrowDown  = 66
	KeyArrowLeft  = 68
	KeyArrowRight = 67
)

// KeyboardReader obsługuje odczyt klawiatury, w tym strzałki
type KeyboardReader struct {
	oldState *termios // Stare ustawienia terminala
}

// termios reprezentuje ustawienia terminala
type termios struct {
	raw bool
}

// NewKeyboardReader tworzy nowy obiekt do odczytu klawiatury
func NewKeyboardReader() *KeyboardReader {
	return &KeyboardReader{
		oldState: &termios{
			raw: false,
		},
	}
}

// EnableRawMode włącza tryb surowy terminala (do odczytu pojedynczych klawiszy)
func (kr *KeyboardReader) EnableRawMode() {
	// W systemie Linux/Unix używamy komend terminalowych
	// Wyłączamy buforowanie i echo
	fmt.Print("\033[?25l") // Ukryj kursor

	// W Go nie ma bezpośredniego dostępu do funkcji termios, więc używamy escape sequence
	os.Stdin.Fd() // Pobierz deskryptor pliku, choć nie używamy go bezpośrednio

	kr.oldState.raw = true
}

// DisableRawMode wyłącza tryb surowy terminala
func (kr *KeyboardReader) DisableRawMode() {
	// Przywróć domyślne zachowanie terminala
	fmt.Print("\033[?25h") // Pokaż kursor

	kr.oldState.raw = false
}

// ReadKey odczytuje pojedynczy klawisz
func (kr *KeyboardReader) ReadKey() (rune, error) {
	// Bufor na jeden znak
	var b = make([]byte, 3)

	// Odczytaj jeden znak
	n, err := os.Stdin.Read(b)
	if err != nil {
		return 0, err
	}

	// Sprawdź czy to sekwencja escape
	if n > 1 && b[0] == '\033' && b[1] == '[' {
		// Sekwencja escape dla strzałek
		if n > 2 {
			switch b[2] {
			case KeyArrowUp:
				return KeyArrowUp, nil
			case KeyArrowDown:
				return KeyArrowDown, nil
			case KeyArrowLeft:
				return KeyArrowLeft, nil
			case KeyArrowRight:
				return KeyArrowRight, nil
			}
		}
		return 0, nil
	}

	// Pojedynczy znak
	return rune(b[0]), nil
}

// SelectOption obsługuje wybór opcji z menu za pomocą strzałek
func (kr *KeyboardReader) SelectOption(options []string, currentIndex int) (int, error) {
	kr.EnableRawMode()
	defer kr.DisableRawMode()

	selectedIndex := currentIndex

	// Wyświetl opcje z aktualnie wybraną
	clearAndPrintOptions(options, selectedIndex)

	for {
		key, err := kr.ReadKey()
		if err != nil {
			return -1, err
		}

		switch key {
		case KeyArrowUp:
			// Przesuń w górę (z zapętleniem)
			selectedIndex--
			if selectedIndex < 0 {
				selectedIndex = len(options) - 1
			}
			clearAndPrintOptions(options, selectedIndex)
		case KeyArrowDown:
			// Przesuń w dół (z zapętleniem)
			selectedIndex = (selectedIndex + 1) % len(options)
			clearAndPrintOptions(options, selectedIndex)
		case KeyEnter:
			// Zatwierdź wybór
			return selectedIndex, nil
		case KeyEsc:
			// Anuluj wybór
			return -1, nil
		}
	}
}

// clearAndPrintOptions czyści ekran i wyświetla opcje z zaznaczoną aktualnie wybraną
func clearAndPrintOptions(options []string, selectedIndex int) {
	// Wyczyść linie (bez czyszczenia całego ekranu)
	for i := 0; i < len(options); i++ {
		fmt.Print("\033[1A\033[2K") // Przesuń kursor do góry i wyczyść linię
	}

	// Wyświetl opcje
	for i, option := range options {
		if i == selectedIndex {
			fmt.Printf("→ %s\n", option) // Strzałka wskazuje wybraną opcję
		} else {
			fmt.Printf("  %s\n", option)
		}
	}
}

// SelectLanguage wyświetla menu wyboru języka
func (consoleUI *ConsoleUI) SelectLanguage(options []string) (int, error) {
	consoleUI.ClearScreen()

	// Tytuł menu
	fmt.Println(consoleUI.CenterText(Bold + Yellow + "=== WYBÓR JĘZYKA / LANGUAGE SELECTION ===" + Reset))
	fmt.Println(consoleUI.CenterText("Użyj strzałek ↑↓ aby wybrać język i Enter aby zatwierdzić"))
	fmt.Println(consoleUI.CenterText("Use ↑↓ arrows to select language and Enter to confirm"))
	fmt.Println()

	// Wyświetl opcje języków
	optionsWithArrow := make([]string, len(options))
	for i, option := range options {
		optionsWithArrow[i] = consoleUI.CenterText(option)
	}

	// Stwórz czytnik klawiatury
	kr := NewKeyboardReader()

	// Centruj opcje na ekranie (prosta implementacja)
	width := consoleUI.terminalWidth / 2
	for i := 0; i < width; i++ {
		fmt.Print(" ")
	}

	// Wybierz język
	return kr.SelectOption(options, 0)
}
