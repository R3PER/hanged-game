package ui

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/r3per/hanged-game/internal/game"
)

// Kody kolorów ANSI
const (
	Reset    = "\033[0m"
	Bold     = "\033[1m"
	Red      = "\033[31m"
	Green    = "\033[32m"
	Yellow   = "\033[33m"
	Blue     = "\033[34m"
	Purple   = "\033[35m"
	Cyan     = "\033[36m"
	White    = "\033[37m"
	BgRed    = "\033[41m"
	BgGreen  = "\033[42m"
	BgYellow = "\033[43m"
	BgBlue   = "\033[44m"
	BgPurple = "\033[45m"
	BgCyan   = "\033[46m"
	BgWhite  = "\033[47m"
)

// Domyślna szerokość terminala
const (
	DefaultTerminalWidth = 80
)

// ConsoleUI reprezentuje interfejs użytkownika konsoli
type ConsoleUI struct {
	reader        *bufio.Reader
	hangman       *game.HangmanDrawing
	showProgress  bool
	terminalWidth int
}

// NewConsoleUI tworzy nowy interfejs użytkownika konsoli
func NewConsoleUI() *ConsoleUI {
	return &ConsoleUI{
		reader:        bufio.NewReader(os.Stdin),
		hangman:       game.NewHangmanDrawing(),
		showProgress:  true,
		terminalWidth: DefaultTerminalWidth,
	}
}

// CenterText centruje tekst w konsoli
func (ui *ConsoleUI) CenterText(text string) string {
	lines := strings.Split(text, "\n")
	centeredLines := make([]string, len(lines))

	for i, line := range lines {
		// Usuń kody ANSI podczas obliczania długości
		visibleLen := utf8.RuneCountInString(stripANSI(line))

		if visibleLen < ui.terminalWidth {
			padding := (ui.terminalWidth - visibleLen) / 2
			centeredLines[i] = strings.Repeat(" ", padding) + line
		} else {
			centeredLines[i] = line
		}
	}

	return strings.Join(centeredLines, "\n")
}

// stripANSI usuwa kody ANSI z tekstu
func stripANSI(text string) string {
	// Proste rozwiązanie - usuwamy wszystkie sekwencje zaczynające się od \033[ i kończące na literze
	result := ""
	inAnsi := false

	for _, r := range text {
		if inAnsi {
			if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
				inAnsi = false
			}
		} else {
			if r == '\033' {
				inAnsi = true
			} else {
				result += string(r)
			}
		}
	}

	return result
}

// PrintTitle wyświetla tytuł gry
func (ui *ConsoleUI) PrintTitle() {
	logo := Bold + Yellow + `
    ▄█    █▄       ▄████████ ███▄▄▄▄      ▄██████▄     ▄████████ ████████▄  
   ███    ███     ███    ███ ███▀▀▀██▄   ███    ███   ███    ███ ███   ▀███ 
   ███    ███     ███    ███ ███   ███   ███    █▀    ███    █▀  ███    ███ 
  ▄███▄▄▄▄███▄▄   ███    ███ ███   ███  ▄███         ▄███▄▄▄     ███    ███ 
 ▀▀███▀▀▀▀███▀  ▀███████████ ███   ███ ▀▀███ ████▄  ▀▀███▀▀▀     ███    ███ 
   ███    ███     ███    ███ ███   ███   ███    ███   ███    █▄  ███    ███ 
   ███    ███     ███    ███ ███   ███   ███    ███   ███    ███ ███   ▄███ 
   ███    █▀      ███    █▀   ▀█   █▀    ████████▀    ██████████ ████████▀  
` + Reset
	fmt.Println(logo)
}

// ClearScreen czyści ekran konsoli
func (ui *ConsoleUI) ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

// PrintGameState wyświetla aktualny stan gry
func (ui *ConsoleUI) PrintGameState(g *game.Game) {
	// Wyświetl rysunek wisielca
	hangmanDrawing := White + ui.hangman.GetDrawing(len(g.WrongGuesses)) + Reset
	fmt.Println(ui.CenterText(hangmanDrawing))

	// Wyświetl słowo z odgadniętymi literami
	wordWithGuesses := g.GetWordWithGuesses()
	fmt.Println(ui.CenterText(Bold + Blue + "\nSłowo: " + White + wordWithGuesses + Reset))

	// Wyświetl błędne próby
	wrongGuesses := g.GetWrongGuesses()
	if wrongGuesses != "" {
		fmt.Println(ui.CenterText(Bold + Red + "Błędne próby: " + White + wrongGuesses + Reset))
	}

	// Wyświetl pozostałe próby
	remainingAttempts := g.GetRemainingAttempts()
	fmt.Println(ui.CenterText(Bold + Yellow + "Pozostałe próby: " + White + fmt.Sprintf("%d", remainingAttempts) + Reset))

	// Wyświetl punkty
	fmt.Println(ui.CenterText(Bold + Green + "Punkty: " + White + fmt.Sprintf("%d", g.Points) + Reset))

	// Wyświetl postęp (opcjonalnie)
	if ui.showProgress {
		progress := g.GetProgress()
		progressText := fmt.Sprintf(Bold+Cyan+"Postęp: "+White+"%.1f%%"+Reset, progress)
		fmt.Println(ui.CenterText(progressText))
	}

	fmt.Println()
}

// PrintWinMessage wyświetla wiadomość o wygranej
func (ui *ConsoleUI) PrintWinMessage(g *game.Game) {
	message := BgGreen + Bold + "GRATULACJE! Odgadłeś słowo: " + g.Word + Reset
	fmt.Println(ui.CenterText(message))

	pointsMsg := Bold + Green + "Zdobyłeś " + fmt.Sprintf("%d", g.Points) + " punktów!" + Reset
	fmt.Println(ui.CenterText(pointsMsg))
}

// PrintLoseMessage wyświetla wiadomość o przegranej
func (ui *ConsoleUI) PrintLoseMessage(g *game.Game) {
	message := BgRed + Bold + "PRZEGRAŁEŚ! Słowo to: " + g.Word + Reset
	fmt.Println(ui.CenterText(message))

	pointsMsg := Bold + Red + "Zdobyłeś " + fmt.Sprintf("%d", g.Points) + " punktów." + Reset
	fmt.Println(ui.CenterText(pointsMsg))
}

// PrintMenu wyświetla menu główne
func (ui *ConsoleUI) PrintMenu() {
	fmt.Println(ui.CenterText(Bold + Yellow + "=== MENU GŁÓWNE ===" + Reset))
	fmt.Println(ui.CenterText(Bold + "1. " + Reset + "Nowa gra"))
	fmt.Println(ui.CenterText(Bold + "2. " + Reset + "Wybierz poziom trudności"))
	fmt.Println(ui.CenterText(Bold + "3. " + Reset + "Pokaż statystyki"))
	fmt.Println(ui.CenterText(Bold + "4. " + Reset + "Wyjście"))
	fmt.Print(ui.CenterText(Bold + "\nWybierz opcję: " + Reset))
}

// PrintDifficultyMenu wyświetla menu wyboru poziomu trudności
func (ui *ConsoleUI) PrintDifficultyMenu() {
	fmt.Println(ui.CenterText(Bold + Yellow + "=== POZIOM TRUDNOŚCI ===" + Reset))
	fmt.Println(ui.CenterText(Bold + "1. " + Reset + "Łatwy (8 prób)"))
	fmt.Println(ui.CenterText(Bold + "2. " + Reset + "Średni (6 prób)"))
	fmt.Println(ui.CenterText(Bold + "3. " + Reset + "Trudny (4 próby)"))
	fmt.Print(ui.CenterText(Bold + "\nWybierz poziom trudności: " + Reset))
}

// GetInput pobiera wejście od użytkownika
func (ui *ConsoleUI) GetInput() string {
	input, _ := ui.reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

// GetMenuOption pobiera opcję menu od użytkownika
func (ui *ConsoleUI) GetMenuOption() int {
	input := ui.GetInput()
	option := 0
	fmt.Sscanf(input, "%d", &option)
	return option
}

// GetLetterInput pobiera literę od użytkownika
func (ui *ConsoleUI) GetLetterInput() rune {
	for {
		fmt.Print(ui.CenterText(Bold + "Podaj literę: " + Reset))
		input := ui.GetInput()

		if input == "" {
			continue
		}

		// Pobierz pierwszą literę z wejścia
		r, _ := utf8.DecodeRuneInString(input)
		if game.IsPolishLetter(r) {
			return r
		}

		fmt.Println(ui.CenterText(Red + "Nieprawidłowy znak. Wprowadź literę alfabetu." + Reset))
	}
}

// ToggleProgressDisplay przełącza wyświetlanie postępu
func (ui *ConsoleUI) ToggleProgressDisplay() {
	ui.showProgress = !ui.showProgress
}

// WaitForEnter czeka na naciśnięcie klawisza Enter
func (ui *ConsoleUI) WaitForEnter() {
	fmt.Print(ui.CenterText(Bold + "\nNaciśnij Enter, aby kontynuować..." + Reset))
	ui.GetInput()
}

// SetupGame konfiguruje nową grę
func (ui *ConsoleUI) SetupGame(wordsManager *game.WordsManager) *game.Game {
	// Wybierz poziom trudności
	ui.ClearScreen()
	ui.PrintDifficultyMenu()

	difficultyLevel := ui.GetMenuOption()
	if difficultyLevel < 1 || difficultyLevel > 3 {
		difficultyLevel = 2 // Domyślnie średni poziom
	}

	// Wybierz losowe słowo
	word := wordsManager.GetRandomWord()

	// Utwórz nową grę
	return game.NewGame(word, difficultyLevel)
}
