package game

import (
	"math"
	"time"
)

// RPGLevel reprezentuje poziom gracza w systemie RPG
type RPGLevel struct {
	Level       int            // Aktualny poziom gracza
	Experience  int            // Aktualne doświadczenie
	NextLevelXP int            // Wymagane doświadczenie do następnego poziomu
	Attributes  *RPGAttributes // Atrybuty gracza
	Inventory   *RPGInventory  // Ekwipunek gracza
}

// RPGAttributes reprezentuje atrybuty gracza
type RPGAttributes struct {
	Intelligence int // Inteligencja - zwiększa szansę na podpowiedź
	Luck         int // Szczęście - zmniejsza szansę na utratę życia przy błędzie
	Perception   int // Percepcja - zwiększa liczbę punktów za odgadnięcie litery
	Resilience   int // Odporność - dodaje dodatkowe próby
}

// RPGItem reprezentuje przedmiot w grze
type RPGItem struct {
	ID          string          // Unikalny identyfikator przedmiotu
	Name        string          // Nazwa przedmiotu
	Description string          // Opis przedmiotu
	Type        string          // Typ przedmiotu (consumable, equipment, etc.)
	Rarity      string          // Rzadkość przedmiotu (common, rare, epic, legendary)
	Effects     []RPGItemEffect // Efekty przedmiotu
	Used        bool            // Czy przedmiot został już użyty
}

// RPGItemEffect reprezentuje efekt przedmiotu
type RPGItemEffect struct {
	Type      string    // Typ efektu (reveal_letter, extra_life, etc.)
	Value     int       // Wartość efektu
	Duration  int       // Czas trwania efektu (0 = jednorazowy, >0 = liczba tur)
	ExpiresAt time.Time // Czas wygaśnięcia efektu
}

// RPGInventory reprezentuje ekwipunek gracza
type RPGInventory struct {
	Items       []RPGItem // Lista przedmiotów w ekwipunku
	MaxCapacity int       // Maksymalna liczba przedmiotów w ekwipunku
}

// RPGQuest reprezentuje zadanie w grze
type RPGQuest struct {
	ID          string // Unikalny identyfikator zadania
	Name        string // Nazwa zadania
	Description string // Opis zadania
	Objective   string // Cel zadania (np. "Odgadnij 5 słów")
	Progress    int    // Postęp w wykonaniu zadania
	Target      int    // Docelowa wartość do osiągnięcia
	Completed   bool   // Czy zadanie zostało ukończone
	Reward      int    // Nagroda XP za ukończenie zadania
}

// NewRPGLevel tworzy nowy obiekt poziomu RPG
func NewRPGLevel() *RPGLevel {
	return &RPGLevel{
		Level:       1,
		Experience:  0,
		NextLevelXP: 100, // Początkowy próg XP
		Attributes: &RPGAttributes{
			Intelligence: 1,
			Luck:         1,
			Perception:   1,
			Resilience:   1,
		},
		Inventory: &RPGInventory{
			Items:       []RPGItem{},
			MaxCapacity: 10,
		},
	}
}

// AddExperience dodaje doświadczenie i awansuje poziom jeśli to konieczne
func (rl *RPGLevel) AddExperience(xp int) (bool, int) {
	rl.Experience += xp
	leveledUp := false
	levelsGained := 0

	// Sprawdź, czy zdobyto wystarczająco dużo XP do awansu
	for rl.Experience >= rl.NextLevelXP {
		rl.Level++
		levelsGained++
		leveledUp = true

		// Odejmij XP potrzebne do poprzedniego poziomu
		rl.Experience -= rl.NextLevelXP

		// Oblicz nowy próg XP (rosnący wykładniczo)
		rl.NextLevelXP = int(float64(rl.NextLevelXP) * 1.5)

		// Dodaj punkty atrybutów przy awansie
		rl.AddAttributePoints(2)
	}

	return leveledUp, levelsGained
}

// AddAttributePoints dodaje punkty atrybutów
func (rl *RPGLevel) AddAttributePoints(points int) {
	// W prostej implementacji rozdzielamy punkty równomiernie
	// W bardziej zaawansowanej gracz mógłby samodzielnie przydzielać punkty
	rl.Attributes.Intelligence += points / 4
	rl.Attributes.Luck += points / 4
	rl.Attributes.Perception += points / 4
	rl.Attributes.Resilience += points / 4

	// Jeśli liczba punktów nie jest podzielna przez 4, dodaj resztę do inteligencji
	remainder := points % 4
	rl.Attributes.Intelligence += remainder
}

// AddItem dodaje przedmiot do ekwipunku
func (inv *RPGInventory) AddItem(item RPGItem) bool {
	if len(inv.Items) >= inv.MaxCapacity {
		return false // Ekwipunek pełny
	}

	inv.Items = append(inv.Items, item)
	return true
}

// UseItem używa przedmiotu i zwraca jego efekty
func (inv *RPGInventory) UseItem(itemID string) ([]RPGItemEffect, bool) {
	for i, item := range inv.Items {
		if item.ID == itemID && !item.Used {
			// Oznacz przedmiot jako użyty (dla przedmiotów jednorazowych)
			if item.Type == "consumable" {
				inv.Items[i].Used = true

				// Opcjonalnie usuń zużyty przedmiot
				// inv.Items = append(inv.Items[:i], inv.Items[i+1:]...)
			}

			return item.Effects, true
		}
	}

	return nil, false
}

// GetIntelligenceBonus zwraca bonus za inteligencję (szansa na podpowiedź)
func (attr *RPGAttributes) GetIntelligenceBonus() float64 {
	// Każdy punkt inteligencji daje 2% szansy na podpowiedź
	return float64(attr.Intelligence) * 0.02
}

// GetLuckBonus zwraca bonus za szczęście (szansa na uniknięcie utraty życia)
func (attr *RPGAttributes) GetLuckBonus() float64 {
	// Każdy punkt szczęścia daje 1.5% szansy na uniknięcie utraty życia
	return float64(attr.Luck) * 0.015
}

// GetPerceptionBonus zwraca bonus za percepcję (dodatkowe punkty)
func (attr *RPGAttributes) GetPerceptionBonus() int {
	// Każdy punkt percepcji zwiększa zdobyte punkty o 1
	return attr.Perception
}

// GetResilienceBonus zwraca bonus za odporność (dodatkowe próby)
func (attr *RPGAttributes) GetResilienceBonus() int {
	// Co 3 punkty odporności dają dodatkową próbę
	return attr.Resilience / 3
}

// GetXPProgress zwraca procentowy postęp do następnego poziomu
func (rl *RPGLevel) GetXPProgress() float64 {
	return math.Min(float64(rl.Experience)/float64(rl.NextLevelXP)*100, 100)
}

// GenerateBasicItems tworzy podstawowe przedmioty do gry
func GenerateBasicItems() []RPGItem {
	return []RPGItem{
		{
			ID:          "potion_hint",
			Name:        "Mikstura Podpowiedzi",
			Description: "Odkrywa losową literę w aktualnym słowie",
			Type:        "consumable",
			Rarity:      "common",
			Effects: []RPGItemEffect{
				{
					Type:  "reveal_letter",
					Value: 1,
				},
			},
		},
		{
			ID:          "scroll_extra_life",
			Name:        "Zwój Dodatkowego Życia",
			Description: "Dodaje jedną dodatkową próbę",
			Type:        "consumable",
			Rarity:      "uncommon",
			Effects: []RPGItemEffect{
				{
					Type:  "extra_life",
					Value: 1,
				},
			},
		},
		{
			ID:          "amulet_wisdom",
			Name:        "Amulet Mądrości",
			Description: "Zwiększa inteligencję o 2 podczas noszenia",
			Type:        "equipment",
			Rarity:      "rare",
			Effects: []RPGItemEffect{
				{
					Type:  "intelligence_boost",
					Value: 2,
				},
			},
		},
		{
			ID:          "ring_fortune",
			Name:        "Pierścień Fortuny",
			Description: "Zwiększa szczęście o 3 podczas noszenia",
			Type:        "equipment",
			Rarity:      "rare",
			Effects: []RPGItemEffect{
				{
					Type:  "luck_boost",
					Value: 3,
				},
			},
		},
	}
}

// GenerateBasicQuests tworzy podstawowe zadania do gry
func GenerateBasicQuests() []RPGQuest {
	return []RPGQuest{
		{
			ID:          "quest_novice",
			Name:        "Początkujący Odgadywacz",
			Description: "Odgadnij poprawnie 3 słowa",
			Objective:   "win_games",
			Progress:    0,
			Target:      3,
			Completed:   false,
			Reward:      50,
		},
		{
			ID:          "quest_perfect",
			Name:        "Perfekcyjna Gra",
			Description: "Odgadnij słowo bez żadnego błędu",
			Objective:   "perfect_game",
			Progress:    0,
			Target:      1,
			Completed:   false,
			Reward:      100,
		},
		{
			ID:          "quest_difficult",
			Name:        "Mistrz Trudności",
			Description: "Wygraj grę na trudnym poziomie",
			Objective:   "win_hard",
			Progress:    0,
			Target:      1,
			Completed:   false,
			Reward:      150,
		},
	}
}

// UpdateQuests aktualizuje postęp w zadaniach
func UpdateQuests(quests []RPGQuest, eventType string, value int) ([]RPGQuest, int) {
	totalXP := 0

	for i, quest := range quests {
		if quest.Completed {
			continue
		}

		// Sprawdź, czy wydarzenie pasuje do celu zadania
		if quest.Objective == eventType {
			quests[i].Progress += value

			// Sprawdź, czy zadanie zostało ukończone
			if quests[i].Progress >= quest.Target {
				quests[i].Progress = quest.Target
				quests[i].Completed = true
				totalXP += quest.Reward
			}
		}
	}

	return quests, totalXP
}
