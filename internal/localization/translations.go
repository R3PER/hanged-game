package localization

// Language reprezentuje kod języka
type Language string

const (
	Polish  Language = "pl"
	English Language = "en"
)

// Translations zawiera wszystkie tłumaczenia dla gry
type Translations struct {
	Language           Language
	MainMenu           MainMenuTranslations
	DifficultyMenu     DifficultyMenuTranslations
	GamePlay           GamePlayTranslations
	Messages           MessagesTranslations
	RPG                RPGTranslations
	LanguageSelection  LanguageSelectionTranslations
	LanguageSelfName   string // Nazwa języka w tym języku (np. "Polski", "English")
	LanguageNativeName string // Nazwa języka po angielsku (np. "Polish", "English")
}

// MainMenuTranslations zawiera tłumaczenia dla menu głównego
type MainMenuTranslations struct {
	Title          string
	NewGame        string
	Difficulty     string
	Statistics     string
	Inventory      string
	QuestLog       string
	Shop           string
	Language       string
	Exit           string
	SelectOption   string
	PressEnter     string
	CharacterLevel string
}

// DifficultyMenuTranslations zawiera tłumaczenia dla menu wyboru trudności
type DifficultyMenuTranslations struct {
	Title            string
	Easy             string
	Medium           string
	Hard             string
	SelectDifficulty string
}

// GamePlayTranslations zawiera tłumaczenia dla rozgrywki
type GamePlayTranslations struct {
	Word              string
	WrongGuesses      string
	RemainingAttempts string
	Points            string
	Progress          string
	EnterLetter       string
	InvalidCharacter  string
}

// MessagesTranslations zawiera tłumaczenia dla komunikatów
type MessagesTranslations struct {
	Congratulations      string
	YouWon               string
	YouLost              string
	WordWas              string
	YouEarned            string
	Points               string
	PressEnterToContinue string
	InvalidOption        string
	DifficultySet        string
	DefaultDifficulty    string
}

// RPGTranslations zawiera tłumaczenia dla elementów RPG
type RPGTranslations struct {
	CharacterInfo       string
	Level               string
	Experience          string
	ProgressToNextLevel string
	Attributes          string
	Intelligence        string
	Luck                string
	Perception          string
	Resilience          string
	IntelligenceBonus   string
	LuckBonus           string
	PerceptionBonus     string
	ResilienceBonus     string
	Inventory           string
	Items               string
	Effects             string
	QuestLog            string
	ActiveQuests        string
	CompletedQuests     string
	NoActiveQuests      string
	NoCompletedQuests   string
	Progress            string
	Reward              string
	ItemShop            string
	AvailableItems      string
	Price               string
	YourXP              string
	ReturnToMainMenu    string
	SelectItemToBuy     string
	LevelUp             string
	Congratulations     string
	NewAttributes       string
	QuestCompleted      string
}

// LanguageSelectionTranslations zawiera tłumaczenia dla wyboru języka
type LanguageSelectionTranslations struct {
	Title          string
	SelectLanguage string
}

// LanguageManager zarządza tłumaczeniami
type LanguageManager struct {
	CurrentLanguage Language
	Translations    map[Language]Translations
}

// GetText zwraca tekst w aktualnym języku
func (lm *LanguageManager) GetText() Translations {
	return lm.Translations[lm.CurrentLanguage]
}

// SetLanguage ustawia aktualny język
func (lm *LanguageManager) SetLanguage(lang Language) {
	lm.CurrentLanguage = lang
}

// NewLanguageManager tworzy nowy menedżer języków
func NewLanguageManager() *LanguageManager {
	translations := make(map[Language]Translations)

	// Polski
	translations[Polish] = Translations{
		Language:           Polish,
		LanguageSelfName:   "Polski",
		LanguageNativeName: "Polish",
		MainMenu: MainMenuTranslations{
			Title:          "MENU GŁÓWNE",
			NewGame:        "Nowa gra",
			Difficulty:     "Wybierz poziom trudności",
			Statistics:     "Pokaż statystyki",
			Inventory:      "Pokaż ekwipunek",
			QuestLog:       "Pokaż dziennik zadań",
			Shop:           "Sklep z przedmiotami",
			Language:       "Wybierz język",
			Exit:           "Wyjście",
			SelectOption:   "Wybierz opcję:",
			PressEnter:     "Naciśnij Enter, aby kontynuować...",
			CharacterLevel: "Poziom postaci:",
		},
		DifficultyMenu: DifficultyMenuTranslations{
			Title:            "POZIOM TRUDNOŚCI",
			Easy:             "Łatwy (8 prób)",
			Medium:           "Średni (6 prób)",
			Hard:             "Trudny (4 próby)",
			SelectDifficulty: "Wybierz poziom trudności:",
		},
		GamePlay: GamePlayTranslations{
			Word:              "Słowo:",
			WrongGuesses:      "Błędne próby:",
			RemainingAttempts: "Pozostałe próby:",
			Points:            "Punkty:",
			Progress:          "Postęp:",
			EnterLetter:       "Podaj literę:",
			InvalidCharacter:  "Nieprawidłowy znak. Wprowadź literę alfabetu.",
		},
		Messages: MessagesTranslations{
			Congratulations:      "GRATULACJE!",
			YouWon:               "Odgadłeś słowo:",
			YouLost:              "PRZEGRAŁEŚ! Słowo to:",
			WordWas:              "Słowo to:",
			YouEarned:            "Zdobyłeś",
			Points:               "punktów",
			PressEnterToContinue: "Naciśnij Enter, aby kontynuować...",
			InvalidOption:        "Nieprawidłowa opcja. Spróbuj ponownie.",
			DifficultySet:        "Ustawiono poziom trudności:",
			DefaultDifficulty:    "Nieprawidłowa opcja. Pozostawiono domyślny poziom trudności.",
		},
		RPG: RPGTranslations{
			CharacterInfo:       "Informacje o Postaci",
			Level:               "Poziom:",
			Experience:          "Doświadczenie:",
			ProgressToNextLevel: "Postęp do następnego poziomu:",
			Attributes:          "Atrybuty:",
			Intelligence:        "Inteligencja:",
			Luck:                "Szczęście:",
			Perception:          "Percepcja:",
			Resilience:          "Odporność:",
			IntelligenceBonus:   "szansy na podpowiedź",
			LuckBonus:           "szansy na uniknięcie błędu",
			PerceptionBonus:     "pkt za trafienie",
			ResilienceBonus:     "dodatkowych prób",
			Inventory:           "Ekwipunek",
			Items:               "Przedmioty",
			Effects:             "Efekty:",
			QuestLog:            "Dziennik Zadań",
			ActiveQuests:        "Aktywne Zadania:",
			CompletedQuests:     "Ukończone Zadania:",
			NoActiveQuests:      "Brak aktywnych zadań",
			NoCompletedQuests:   "Brak ukończonych zadań",
			Progress:            "Postęp:",
			Reward:              "Nagroda:",
			ItemShop:            "Sklep z Przedmiotami",
			AvailableItems:      "Dostępne przedmioty:",
			Price:               "Cena:",
			YourXP:              "Twoje XP:",
			ReturnToMainMenu:    "Powrót do menu głównego",
			SelectItemToBuy:     "Wybierz przedmiot do kupienia:",
			LevelUp:             "AWANS POZIOMU",
			Congratulations:     "Gratulacje!",
			NewAttributes:       "Nowe atrybuty:",
			QuestCompleted:      "ZADANIE UKOŃCZONE",
		},
		LanguageSelection: LanguageSelectionTranslations{
			Title:          "WYBÓR JĘZYKA",
			SelectLanguage: "Wybierz język:",
		},
	}

	// English
	translations[English] = Translations{
		Language:           English,
		LanguageSelfName:   "English",
		LanguageNativeName: "English",
		MainMenu: MainMenuTranslations{
			Title:          "MAIN MENU",
			NewGame:        "New game",
			Difficulty:     "Select difficulty level",
			Statistics:     "Show statistics",
			Inventory:      "Show inventory",
			QuestLog:       "Show quest log",
			Shop:           "Item shop",
			Language:       "Select language",
			Exit:           "Exit",
			SelectOption:   "Select option:",
			PressEnter:     "Press Enter to continue...",
			CharacterLevel: "Character level:",
		},
		DifficultyMenu: DifficultyMenuTranslations{
			Title:            "DIFFICULTY LEVEL",
			Easy:             "Easy (8 attempts)",
			Medium:           "Medium (6 attempts)",
			Hard:             "Hard (4 attempts)",
			SelectDifficulty: "Select difficulty level:",
		},
		GamePlay: GamePlayTranslations{
			Word:              "Word:",
			WrongGuesses:      "Wrong guesses:",
			RemainingAttempts: "Remaining attempts:",
			Points:            "Points:",
			Progress:          "Progress:",
			EnterLetter:       "Enter a letter:",
			InvalidCharacter:  "Invalid character. Enter a letter of the alphabet.",
		},
		Messages: MessagesTranslations{
			Congratulations:      "CONGRATULATIONS!",
			YouWon:               "You guessed the word:",
			YouLost:              "YOU LOST! The word was:",
			WordWas:              "The word was:",
			YouEarned:            "You earned",
			Points:               "points",
			PressEnterToContinue: "Press Enter to continue...",
			InvalidOption:        "Invalid option. Try again.",
			DifficultySet:        "Difficulty level set to:",
			DefaultDifficulty:    "Invalid option. Default difficulty level kept.",
		},
		RPG: RPGTranslations{
			CharacterInfo:       "Character Information",
			Level:               "Level:",
			Experience:          "Experience:",
			ProgressToNextLevel: "Progress to next level:",
			Attributes:          "Attributes:",
			Intelligence:        "Intelligence:",
			Luck:                "Luck:",
			Perception:          "Perception:",
			Resilience:          "Resilience:",
			IntelligenceBonus:   "chance for a hint",
			LuckBonus:           "chance to avoid a mistake",
			PerceptionBonus:     "pts for a hit",
			ResilienceBonus:     "additional attempts",
			Inventory:           "Inventory",
			Items:               "Items",
			Effects:             "Effects:",
			QuestLog:            "Quest Log",
			ActiveQuests:        "Active Quests:",
			CompletedQuests:     "Completed Quests:",
			NoActiveQuests:      "No active quests",
			NoCompletedQuests:   "No completed quests",
			Progress:            "Progress:",
			Reward:              "Reward:",
			ItemShop:            "Item Shop",
			AvailableItems:      "Available items:",
			Price:               "Price:",
			YourXP:              "Your XP:",
			ReturnToMainMenu:    "Return to main menu",
			SelectItemToBuy:     "Select item to buy:",
			LevelUp:             "LEVEL UP",
			Congratulations:     "Congratulations!",
			NewAttributes:       "New attributes:",
			QuestCompleted:      "QUEST COMPLETED",
		},
		LanguageSelection: LanguageSelectionTranslations{
			Title:          "LANGUAGE SELECTION",
			SelectLanguage: "Select language:",
		},
	}

	return &LanguageManager{
		CurrentLanguage: Polish, // Domyślny język
		Translations:    translations,
	}
}
