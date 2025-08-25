package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/r3per/hanged-game/internal/game"
	"github.com/r3per/hanged-game/internal/localization"
	"github.com/r3per/hanged-game/internal/storage"
	"github.com/r3per/hanged-game/internal/ui"
)

const (
	WordsFilePath      = "data/words.txt"
	StatsFilePath      = "data/stats.json"
	LanguageConfigPath = "data/language.txt"
)

var (
	difficultyLevel = 2 // Domyślnie średni poziom trudności
)

func main() {
	// Upewnij się, że katalog data istnieje
	dataDir := filepath.Dir(StatsFilePath)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		os.MkdirAll(dataDir, 0755)
	}

	// Inicjalizacja menedżera języka
	langManager := localization.NewLanguageManager()

	// Wczytaj preferencje językowe, jeśli istnieją
	if language, err := loadLanguagePreference(LanguageConfigPath); err == nil && language != "" {
		switch language {
		case "pl":
			langManager.SetLanguage(localization.Polish)
		case "en":
			langManager.SetLanguage(localization.English)
		}
	}

	// Pobierz teksty w aktualnym języku
	txt := langManager.GetText()

	// Inicjalizacja menedżera słów
	wordsManager, err := game.NewWordsManager(WordsFilePath)
	if err != nil {
		fmt.Printf("Błąd podczas ładowania słów: %v\n", err)
		os.Exit(1)
	}

	// Inicjalizacja menedżera statystyk
	statsManager, err := storage.NewStatsManager(StatsFilePath)
	if err != nil {
		fmt.Printf("Błąd podczas ładowania statystyk: %v\n", err)
		os.Exit(1)
	}

	// Inicjalizacja interfejsu użytkownika
	consoleUI := ui.NewConsoleUI()

	// Inicjalizacja interfejsu RPG
	rpgLevel := game.NewRPGLevel()
	quests := game.GenerateBasicQuests()
	rpgUI := ui.NewRPGCharacterUI(rpgLevel, quests)

	// Główna pętla programu
	for {
		// Pobierz teksty w aktualnym języku (na wypadek zmiany języka)
		txt = langManager.GetText()

		consoleUI.ClearScreen()
		consoleUI.PrintTitle()

		// Wyświetl menu główne z RPG UI
		rpgUI.PrintRPGMainMenu(consoleUI)

		option := consoleUI.GetMenuOption()

		switch option {
		case 1: // Nowa gra
			playGame(consoleUI, wordsManager, statsManager, langManager, rpgLevel)
		case 2: // Wybierz poziom trudności
			selectDifficulty(consoleUI, txt)
		case 3: // Pokaż statystyki
			showStats(consoleUI, statsManager, txt)
		case 4: // Pokaż ekwipunek
			rpgUI.PrintInventory(consoleUI)
			consoleUI.WaitForEnter()
		case 5: // Pokaż dziennik zadań
			rpgUI.PrintQuestLog(consoleUI)
			consoleUI.WaitForEnter()
		case 6: // Sklep z przedmiotami
			showItemShop(consoleUI, rpgLevel, rpgUI)
		case 7: // Wybierz język
			selectLanguage(consoleUI, langManager)
			// Zapisz preferencje językowe
			saveLanguagePreference(LanguageConfigPath, string(langManager.CurrentLanguage))
		case 8: // Wyjście
			fmt.Println(consoleUI.CenterText(txt.Messages.PressEnterToContinue))
			return
		default:
			fmt.Println(consoleUI.CenterText(txt.Messages.InvalidOption))
			consoleUI.WaitForEnter()
		}
	}
}

// playGame prowadzi rozgrywkę
func playGame(consoleUI *ui.ConsoleUI, wordsManager *game.WordsManager, statsManager *storage.StatsManager, langManager *localization.LanguageManager, rpgLevel *game.RPGLevel) {
	txt := langManager.GetText()

	// Utwórz nową grę
	g := consoleUI.SetupGame(wordsManager)

	// Główna pętla gry
	for g.State == game.Playing {
		consoleUI.ClearScreen()
		consoleUI.PrintGameState(g)

		// Pobierz literę od użytkownika
		letter := consoleUI.GetLetterInput()

		// Dokonaj próby odgadnięcia
		g.Guess(letter)
	}

	// Wyświetl wynik gry
	consoleUI.ClearScreen()
	consoleUI.PrintGameState(g)

	if g.State == game.Won {
		// Użyj przetłumaczonych tekstów do komunikatu o wygranej
		message := ui.BgGreen + ui.Bold + txt.Messages.Congratulations + " " + txt.Messages.YouWon + " " + g.Word + ui.Reset
		fmt.Println(consoleUI.CenterText(message))

		pointsMsg := ui.Bold + ui.Green + txt.Messages.YouEarned + " " + fmt.Sprintf("%d", g.Points) + " " + txt.Messages.Points + "!" + ui.Reset
		fmt.Println(consoleUI.CenterText(pointsMsg))
		// Zapisz wynik jako wygraną
		statsManager.AddGameResult(g.Word, "win", g.Points, g.MaxAttempts)

		// Dodaj doświadczenie
		leveledUp, levelsGained := rpgLevel.AddExperience(g.Points)

		// Aktualizuj zadania (questy)
		updatedQuests, questXP := game.UpdateQuests(game.GenerateBasicQuests(), "win_games", 1)
		if questXP > 0 {
			rpgLevel.AddExperience(questXP)
		}

		// Jeśli gracz awansował na wyższy poziom, wyświetl informację
		if leveledUp {
			// Utwórz nowy interfejs RPG z zaktualizowanymi danymi
			rpgUI := ui.NewRPGCharacterUI(rpgLevel, updatedQuests)

			// Wyświetl informację o awansie na wyższy poziom
			consoleUI.ClearScreen()
			rpgUI.PrintLevelUpNotification(consoleUI, rpgLevel.Level)

			// Jeśli awansował o więcej niż jeden poziom, wyświetl informację
			if levelsGained > 1 {
				fmt.Println(consoleUI.CenterText(ui.Bold + ui.Green +
					fmt.Sprintf("Awansowałeś o %d poziomów!", levelsGained) + ui.Reset))
			}

			consoleUI.WaitForEnter()
		}

		// Sprawdź, czy ukończono jakieś zadania i wyświetl informację
		for _, quest := range updatedQuests {
			if quest.Completed && quest.Progress == quest.Target {
				rpgUI := ui.NewRPGCharacterUI(rpgLevel, updatedQuests)
				consoleUI.ClearScreen()
				rpgUI.PrintQuestCompleteNotification(consoleUI, quest)
				consoleUI.WaitForEnter()
				break
			}
		}
	} else {
		// Użyj przetłumaczonych tekstów do komunikatu o przegranej
		message := ui.BgRed + ui.Bold + txt.Messages.YouLost + " " + g.Word + ui.Reset
		fmt.Println(consoleUI.CenterText(message))

		pointsMsg := ui.Bold + ui.Red + txt.Messages.YouEarned + " " + fmt.Sprintf("%d", g.Points) + " " + txt.Messages.Points + "." + ui.Reset
		fmt.Println(consoleUI.CenterText(pointsMsg))
		// Zapisz wynik jako przegraną
		statsManager.AddGameResult(g.Word, "lose", g.Points, g.MaxAttempts)
	}

	consoleUI.WaitForEnter()
}

// selectDifficulty pozwala wybrać poziom trudności
func selectDifficulty(consoleUI *ui.ConsoleUI, txt localization.Translations) {
	consoleUI.ClearScreen()
	consoleUI.PrintDifficultyMenu()

	option := consoleUI.GetMenuOption()
	if option >= 1 && option <= 3 {
		difficultyLevel = option
		fmt.Printf(consoleUI.CenterText(txt.Messages.DifficultySet+" %d\n"), difficultyLevel)
	} else {
		fmt.Println(consoleUI.CenterText(txt.Messages.DefaultDifficulty))
	}

	consoleUI.WaitForEnter()
}

// showStats wyświetla statystyki gracza
func showStats(consoleUI *ui.ConsoleUI, statsManager *storage.StatsManager, txt localization.Translations) {
	consoleUI.ClearScreen()
	fmt.Println(consoleUI.CenterText(ui.Bold + ui.Yellow + "=== " + txt.MainMenu.Statistics + " ===" + ui.Reset))

	stats := statsManager.GetStats()
	winRate := statsManager.GetWinRate()
	avgScore := statsManager.GetAverageScore()

	// Przetłumaczone teksty dla statystyk
	statTexts := []string{
		ui.Bold + "Rozegrane gry: " + ui.Reset + fmt.Sprintf("%d", stats.GamesPlayed),
		ui.Bold + "Wygrane gry: " + ui.Reset + fmt.Sprintf("%d", stats.GamesWon),
		ui.Bold + "Współczynnik wygranych: " + ui.Reset + fmt.Sprintf("%.1f%%", winRate),
		ui.Bold + "Łączna liczba punktów: " + ui.Reset + fmt.Sprintf("%d", stats.TotalPoints),
		ui.Bold + "Średni wynik: " + ui.Reset + fmt.Sprintf("%.1f", avgScore),
		ui.Bold + "Najwyższy wynik: " + ui.Reset + fmt.Sprintf("%d", stats.HighestScore),
	}

	// Użyj tłumaczeń dla angielskiej wersji
	if txt.Language == localization.English {
		statTexts = []string{
			ui.Bold + "Games played: " + ui.Reset + fmt.Sprintf("%d", stats.GamesPlayed),
			ui.Bold + "Games won: " + ui.Reset + fmt.Sprintf("%d", stats.GamesWon),
			ui.Bold + "Win rate: " + ui.Reset + fmt.Sprintf("%.1f%%", winRate),
			ui.Bold + "Total points: " + ui.Reset + fmt.Sprintf("%d", stats.TotalPoints),
			ui.Bold + "Average score: " + ui.Reset + fmt.Sprintf("%.1f", avgScore),
			ui.Bold + "Highest score: " + ui.Reset + fmt.Sprintf("%d", stats.HighestScore),
		}
	}

	for _, text := range statTexts {
		fmt.Println(consoleUI.CenterText(text))
	}

	// Wyświetl ostatnie gry
	lastGames := statsManager.GetLastGames(5)
	if len(lastGames) > 0 {
		// Przetłumaczony nagłówek dla ostatnich gier
		recentGamesHeader := "\n=== OSTATNIE GRY ==="
		if txt.Language == localization.English {
			recentGamesHeader = "\n=== RECENT GAMES ==="
		}

		fmt.Println(consoleUI.CenterText(ui.Bold + ui.Yellow + recentGamesHeader + ui.Reset))
		for i, game := range lastGames {
			resultColor := ui.Red
			resultText := "PRZEGRANA"
			if game.Result == "win" {
				resultColor = ui.Green
				resultText = "WYGRANA"
			}

			// Tłumaczenia dla angielskiej wersji
			if txt.Language == localization.English {
				if game.Result == "win" {
					resultText = "WON"
				} else {
					resultText = "LOST"
				}
			}

			difficultyText := "Średni"
			switch game.Difficulty {
			case 8: // game.EasyLevel
				difficultyText = txt.DifficultyMenu.Easy
			case 6: // game.MediumLevel
				difficultyText = txt.DifficultyMenu.Medium
			case 4: // game.HardLevel
				difficultyText = txt.DifficultyMenu.Hard
			}

			gameText := fmt.Sprintf("%d. %s: %s [%s] - %s%s%s (%d pkt)",
				i+1,
				game.Date.Format("02.01.2006 15:04"),
				game.Word,
				difficultyText,
				resultColor,
				resultText,
				ui.Reset,
				game.Points)

			fmt.Println(consoleUI.CenterText(gameText))
		}
	}

	consoleUI.WaitForEnter()
}

// selectLanguage pozwala wybrać język
func selectLanguage(consoleUI *ui.ConsoleUI, langManager *localization.LanguageManager) {
	// Utwórz listę dostępnych języków
	languages := []string{
		langManager.Translations[localization.Polish].LanguageSelfName,
		langManager.Translations[localization.English].LanguageSelfName,
	}

	// Wybierz język
	selectedIndex, err := consoleUI.SelectLanguage(languages)
	if err != nil || selectedIndex < 0 {
		return
	}

	// Ustaw wybrany język
	switch selectedIndex {
	case 0:
		langManager.SetLanguage(localization.Polish)
	case 1:
		langManager.SetLanguage(localization.English)
	}
}

// showItemShop wyświetla sklep z przedmiotami
func showItemShop(consoleUI *ui.ConsoleUI, rpgLevel *game.RPGLevel, rpgUI *ui.RPGCharacterUI) {
	// Generuj przedmioty sklepowe
	shopItems := game.GenerateBasicItems()

	// Wyświetl sklep
	rpgUI.PrintItemShop(consoleUI, shopItems, rpgLevel.Experience)

	// Pobierz wybór gracza
	option := consoleUI.GetMenuOption()

	// Jeśli wybrano powrót do menu, zakończ funkcję
	if option <= 0 || option > len(shopItems) {
		return
	}

	// Logika kupowania przedmiotów mogłaby być tutaj zaimplementowana
	// Ale na razie ją pomijamy
}

// loadLanguagePreference wczytuje preferencje językowe
func loadLanguagePreference(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// saveLanguagePreference zapisuje preferencje językowe
func saveLanguagePreference(filePath string, language string) error {
	return os.WriteFile(filePath, []byte(language), 0644)
}
