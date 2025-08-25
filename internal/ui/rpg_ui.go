package ui

import (
	"fmt"
	"strings"

	"github.com/r3per/hanged-game/internal/game"
)

// RPGCharacterUI wyświetla informacje o postaci RPG
type RPGCharacterUI struct {
	rpgLevel *game.RPGLevel
	quests   []game.RPGQuest
}

// NewRPGCharacterUI tworzy nowy interfejs postaci RPG
func NewRPGCharacterUI(rpgLevel *game.RPGLevel, quests []game.RPGQuest) *RPGCharacterUI {
	return &RPGCharacterUI{
		rpgLevel: rpgLevel,
		quests:   quests,
	}
}

// SetQuests ustawia zadania
func (rui *RPGCharacterUI) SetQuests(quests []game.RPGQuest) {
	rui.quests = quests
}

// DrawRPGBox rysuje ramkę w stylu RPG
func (ui *ConsoleUI) DrawRPGBox(title string, content []string, width int) string {
	if width < 20 {
		width = 20
	}

	// Utwórz górną ramkę
	topBorder := "╔" + strings.Repeat("═", width-2) + "╗"
	titleLine := "║ " + title + strings.Repeat(" ", width-4-len(stripANSI(title))) + " ║"
	separator := "╠" + strings.Repeat("═", width-2) + "╣"

	// Utwórz zawartość
	contentLines := make([]string, 0)
	for _, line := range content {
		strippedLine := stripANSI(line)
		if len(strippedLine) > width-4 {
			// Podziel długie linie
			for i := 0; i < len(strippedLine); i += width - 4 {
				end := i + width - 4
				if end > len(strippedLine) {
					end = len(strippedLine)
				}
				contentLines = append(contentLines, "║ "+line[i:end]+strings.Repeat(" ", width-4-(end-i))+" ║")
			}
		} else {
			contentLines = append(contentLines, "║ "+line+strings.Repeat(" ", width-4-len(strippedLine))+" ║")
		}
	}

	// Utwórz dolną ramkę
	bottomBorder := "╚" + strings.Repeat("═", width-2) + "╝"

	// Połącz wszystkie elementy
	result := topBorder + "\n" + titleLine + "\n" + separator + "\n"
	result += strings.Join(contentLines, "\n") + "\n"
	result += bottomBorder

	return result
}

// PrintCharacterInfo wyświetla informacje o postaci
func (rui *RPGCharacterUI) PrintCharacterInfo(consoleUI *ConsoleUI) {
	// Przygotuj informacje o postaci
	charInfoContent := []string{
		Bold + "Poziom: " + Reset + fmt.Sprintf("%d", rui.rpgLevel.Level),
		Bold + "Doświadczenie: " + Reset + fmt.Sprintf("%d/%d", rui.rpgLevel.Experience, rui.rpgLevel.NextLevelXP),
		Bold + "Postęp do następnego poziomu: " + Reset + fmt.Sprintf("%.1f%%", rui.rpgLevel.GetXPProgress()),
		"",
		Bold + Yellow + "Atrybuty:" + Reset,
		Bold + "Inteligencja: " + Reset + fmt.Sprintf("%d (+%.1f%% szansy na podpowiedź)",
			rui.rpgLevel.Attributes.Intelligence,
			rui.rpgLevel.Attributes.GetIntelligenceBonus()*100),
		Bold + "Szczęście: " + Reset + fmt.Sprintf("%d (+%.1f%% szansy na uniknięcie błędu)",
			rui.rpgLevel.Attributes.Luck,
			rui.rpgLevel.Attributes.GetLuckBonus()*100),
		Bold + "Percepcja: " + Reset + fmt.Sprintf("%d (+%d pkt za trafienie)",
			rui.rpgLevel.Attributes.Perception,
			rui.rpgLevel.Attributes.GetPerceptionBonus()),
		Bold + "Odporność: " + Reset + fmt.Sprintf("%d (+%d dodatkowych prób)",
			rui.rpgLevel.Attributes.Resilience,
			rui.rpgLevel.Attributes.GetResilienceBonus()),
	}

	// Wyświetl ramkę z informacjami o postaci
	charBox := consoleUI.DrawRPGBox("Informacje o Postaci", charInfoContent, 60)
	fmt.Println(consoleUI.CenterText(charBox))
	fmt.Println()
}

// PrintQuestLog wyświetla dziennik zadań
func (rui *RPGCharacterUI) PrintQuestLog(consoleUI *ConsoleUI) {
	// Przygotuj informacje o zadaniach
	questContent := []string{
		Bold + Yellow + "Aktywne Zadania:" + Reset,
	}

	activeQuestsCount := 0
	for _, quest := range rui.quests {
		if !quest.Completed {
			progressPercent := float64(quest.Progress) / float64(quest.Target) * 100
			questLine := Bold + quest.Name + ": " + Reset + quest.Description
			progressLine := fmt.Sprintf("Postęp: %d/%d (%.1f%%) - Nagroda: %d XP",
				quest.Progress, quest.Target, progressPercent, quest.Reward)

			questContent = append(questContent, questLine)
			questContent = append(questContent, progressLine)
			questContent = append(questContent, "")
			activeQuestsCount++
		}
	}

	if activeQuestsCount == 0 {
		questContent = append(questContent, "Brak aktywnych zadań")
	}

	// Dodaj ukończone zadania
	questContent = append(questContent, "")
	questContent = append(questContent, Bold+Green+"Ukończone Zadania:"+Reset)

	completedQuestsCount := 0
	for _, quest := range rui.quests {
		if quest.Completed {
			questContent = append(questContent, Bold+Green+"✓ "+quest.Name+Reset+": "+quest.Description)
			completedQuestsCount++
		}
	}

	if completedQuestsCount == 0 {
		questContent = append(questContent, "Brak ukończonych zadań")
	}

	// Wyświetl ramkę z informacjami o zadaniach
	questBox := consoleUI.DrawRPGBox("Dziennik Zadań", questContent, 70)
	fmt.Println(consoleUI.CenterText(questBox))
	fmt.Println()
}

// PrintInventory wyświetla ekwipunek
func (rui *RPGCharacterUI) PrintInventory(consoleUI *ConsoleUI) {
	// Przygotuj informacje o ekwipunku
	inventoryContent := []string{
		Bold + Yellow + fmt.Sprintf("Przedmioty (%d/%d):",
			len(rui.rpgLevel.Inventory.Items),
			rui.rpgLevel.Inventory.MaxCapacity) + Reset,
		"",
	}

	if len(rui.rpgLevel.Inventory.Items) == 0 {
		inventoryContent = append(inventoryContent, "Ekwipunek jest pusty")
	} else {
		for i, item := range rui.rpgLevel.Inventory.Items {
			// Koloruj przedmioty w zależności od rzadkości
			rarityColor := White
			switch item.Rarity {
			case "common":
				rarityColor = White
			case "uncommon":
				rarityColor = Green
			case "rare":
				rarityColor = Blue
			case "epic":
				rarityColor = Purple
			case "legendary":
				rarityColor = Yellow
			}

			usedStatus := ""
			if item.Used {
				usedStatus = Red + " [UŻYTY]" + Reset
			}

			itemLine := fmt.Sprintf("%d. %s%s%s%s - %s",
				i+1,
				rarityColor,
				Bold,
				item.Name,
				Reset,
				item.Description)

			inventoryContent = append(inventoryContent, itemLine+usedStatus)

			// Dodaj efekty przedmiotu
			effectsLine := "   Efekty: "
			for j, effect := range item.Effects {
				if j > 0 {
					effectsLine += ", "
				}

				switch effect.Type {
				case "reveal_letter":
					effectsLine += fmt.Sprintf("Odkryj %d literę", effect.Value)
				case "extra_life":
					effectsLine += fmt.Sprintf("+%d życie", effect.Value)
				case "intelligence_boost":
					effectsLine += fmt.Sprintf("+%d inteligencji", effect.Value)
				case "luck_boost":
					effectsLine += fmt.Sprintf("+%d szczęścia", effect.Value)
				default:
					effectsLine += fmt.Sprintf("%s: %d", effect.Type, effect.Value)
				}
			}

			inventoryContent = append(inventoryContent, effectsLine)
			inventoryContent = append(inventoryContent, "")
		}
	}

	// Wyświetl ramkę z informacjami o ekwipunku
	inventoryBox := consoleUI.DrawRPGBox("Ekwipunek", inventoryContent, 65)
	fmt.Println(consoleUI.CenterText(inventoryBox))
	fmt.Println()
}

// PrintRPGGameStats wyświetla statystyki gry w stylu RPG
func (rui *RPGCharacterUI) PrintRPGGameStats(consoleUI *ConsoleUI, g *game.Game) {
	// Dodatkowe atrybuty z bonusami RPG
	extraLives := rui.rpgLevel.Attributes.GetResilienceBonus()
	pointsBonus := rui.rpgLevel.Attributes.GetPerceptionBonus()

	// Przygotuj informacje o grze
	gameStatsContent := []string{
		Bold + Blue + "Słowo: " + White + g.GetWordWithGuesses() + Reset,
		"",
		Bold + Yellow + "Pozostałe próby: " + White + fmt.Sprintf("%d (+%d)", g.GetRemainingAttempts(), extraLives) + Reset,
		Bold + Green + "Punkty: " + White + fmt.Sprintf("%d (+%d za trafienie)", g.Points, pointsBonus) + Reset,
	}

	// Dodaj informacje o błędnych próbach
	wrongGuesses := g.GetWrongGuesses()
	if wrongGuesses != "" {
		gameStatsContent = append(gameStatsContent, "")
		gameStatsContent = append(gameStatsContent, Bold+Red+"Błędne próby: "+White+wrongGuesses+Reset)
	}

	// Wyświetl ramkę z informacjami o grze
	gameStatsBox := consoleUI.DrawRPGBox("Status Gry", gameStatsContent, 50)
	fmt.Println(consoleUI.CenterText(gameStatsBox))
}

// PrintRPGMainMenu wyświetla menu główne w stylu RPG
func (rui *RPGCharacterUI) PrintRPGMainMenu(consoleUI *ConsoleUI) {
	// Przygotuj opcje menu
	menuContent := []string{
		Bold + "1. " + Reset + "Nowa gra",
		Bold + "2. " + Reset + "Wybierz poziom trudności",
		Bold + "3. " + Reset + "Pokaż statystyki",
		Bold + "4. " + Reset + "Pokaż ekwipunek",
		Bold + "5. " + Reset + "Pokaż dziennik zadań",
		Bold + "6. " + Reset + "Sklep z przedmiotami",
		Bold + "7. " + Reset + "Wyjście",
		"",
		Bold + Yellow + "Poziom postaci: " + Reset + fmt.Sprintf("%d | XP: %d/%d",
			rui.rpgLevel.Level, rui.rpgLevel.Experience, rui.rpgLevel.NextLevelXP),
	}

	// Wyświetl ramkę z menu
	menuBox := consoleUI.DrawRPGBox("MENU GŁÓWNE", menuContent, 40)
	fmt.Println(consoleUI.CenterText(menuBox))
	fmt.Print(consoleUI.CenterText(Bold + "\nWybierz opcję: " + Reset))
}

// PrintLevelUpNotification wyświetla informację o awansie na wyższy poziom
func (rui *RPGCharacterUI) PrintLevelUpNotification(consoleUI *ConsoleUI, newLevel int) {
	notificationContent := []string{
		Bold + Yellow + "Gratulacje!" + Reset,
		"",
		fmt.Sprintf("Osiągnąłeś %d poziom!", newLevel),
		"",
		Bold + "Nowe atrybuty:" + Reset,
		fmt.Sprintf("Inteligencja: %d", rui.rpgLevel.Attributes.Intelligence),
		fmt.Sprintf("Szczęście: %d", rui.rpgLevel.Attributes.Luck),
		fmt.Sprintf("Percepcja: %d", rui.rpgLevel.Attributes.Perception),
		fmt.Sprintf("Odporność: %d", rui.rpgLevel.Attributes.Resilience),
	}

	// Wyświetl ramkę z powiadomieniem
	levelUpBox := consoleUI.DrawRPGBox("AWANS POZIOMU", notificationContent, 40)
	fmt.Println(consoleUI.CenterText(levelUpBox))
}

// PrintQuestCompleteNotification wyświetla informację o ukończeniu zadania
func (rui *RPGCharacterUI) PrintQuestCompleteNotification(consoleUI *ConsoleUI, quest game.RPGQuest) {
	notificationContent := []string{
		Bold + Green + "Zadanie ukończone!" + Reset,
		"",
		Bold + quest.Name + Reset,
		quest.Description,
		"",
		fmt.Sprintf("Nagroda: %d XP", quest.Reward),
	}

	// Wyświetl ramkę z powiadomieniem
	questBox := consoleUI.DrawRPGBox("ZADANIE UKOŃCZONE", notificationContent, 50)
	fmt.Println(consoleUI.CenterText(questBox))
}

// PrintItemShop wyświetla sklep z przedmiotami
func (rui *RPGCharacterUI) PrintItemShop(consoleUI *ConsoleUI, shopItems []game.RPGItem, playerXP int) {
	// Przygotuj informacje o sklepie
	shopContent := []string{
		Bold + Yellow + "Dostępne przedmioty:" + Reset,
		"",
	}

	// Ceny przedmiotów (zależne od rzadkości)
	prices := map[string]int{
		"common":    50,
		"uncommon":  100,
		"rare":      250,
		"epic":      500,
		"legendary": 1000,
	}

	for i, item := range shopItems {
		// Koloruj przedmioty w zależności od rzadkości
		rarityColor := White
		price := 0

		switch item.Rarity {
		case "common":
			rarityColor = White
			price = prices["common"]
		case "uncommon":
			rarityColor = Green
			price = prices["uncommon"]
		case "rare":
			rarityColor = Blue
			price = prices["rare"]
		case "epic":
			rarityColor = Purple
			price = prices["epic"]
		case "legendary":
			rarityColor = Yellow
			price = prices["legendary"]
		}

		// Sprawdź, czy gracz może kupić przedmiot
		canBuy := playerXP >= price
		priceText := fmt.Sprintf("%d XP", price)
		if !canBuy {
			priceText = Red + priceText + Reset
		}

		itemLine := fmt.Sprintf("%d. %s%s%s%s - %s",
			i+1,
			rarityColor,
			Bold,
			item.Name,
			Reset,
			item.Description)

		shopContent = append(shopContent, itemLine)
		shopContent = append(shopContent, fmt.Sprintf("   Cena: %s", priceText))

		// Dodaj efekty przedmiotu
		effectsLine := "   Efekty: "
		for j, effect := range item.Effects {
			if j > 0 {
				effectsLine += ", "
			}

			switch effect.Type {
			case "reveal_letter":
				effectsLine += fmt.Sprintf("Odkryj %d literę", effect.Value)
			case "extra_life":
				effectsLine += fmt.Sprintf("+%d życie", effect.Value)
			case "intelligence_boost":
				effectsLine += fmt.Sprintf("+%d inteligencji", effect.Value)
			case "luck_boost":
				effectsLine += fmt.Sprintf("+%d szczęścia", effect.Value)
			default:
				effectsLine += fmt.Sprintf("%s: %d", effect.Type, effect.Value)
			}
		}

		shopContent = append(shopContent, effectsLine)
		shopContent = append(shopContent, "")
	}

	shopContent = append(shopContent, Bold+Yellow+"Twoje XP: "+Reset+fmt.Sprintf("%d", playerXP))
	shopContent = append(shopContent, Bold+"0. "+Reset+"Powrót do menu głównego")

	// Wyświetl ramkę z informacjami o sklepie
	shopBox := consoleUI.DrawRPGBox("Sklep z Przedmiotami", shopContent, 70)
	fmt.Println(consoleUI.CenterText(shopBox))
	fmt.Print(consoleUI.CenterText(Bold + "\nWybierz przedmiot do kupienia: " + Reset))
}
