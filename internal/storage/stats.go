package storage

import (
	"encoding/json"
	"os"
	"time"
)

// GameStats reprezentuje statystyki pojedynczej gry
type GameStats struct {
	Word       string    `json:"word"`
	Result     string    `json:"result"` // "win" lub "lose"
	Points     int       `json:"points"`
	Difficulty int       `json:"difficulty"`
	Date       time.Time `json:"date"`
}

// PlayerStats reprezentuje statystyki gracza
type PlayerStats struct {
	GamesPlayed  int         `json:"games_played"`
	GamesWon     int         `json:"games_won"`
	TotalPoints  int         `json:"total_points"`
	HighestScore int         `json:"highest_score"`
	GameHistory  []GameStats `json:"game_history"`
}

// StatsManager zarządza statystykami gracza
type StatsManager struct {
	stats    PlayerStats
	filePath string
}

// NewStatsManager tworzy nowy manager statystyk
func NewStatsManager(filePath string) (*StatsManager, error) {
	sm := &StatsManager{
		filePath: filePath,
		stats: PlayerStats{
			GameHistory: []GameStats{},
		},
	}

	// Spróbuj odczytać istniejące statystyki
	_, err := os.Stat(filePath)
	if err == nil {
		err = sm.loadStats()
		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	return sm, nil
}

// loadStats wczytuje statystyki z pliku
func (sm *StatsManager) loadStats() error {
	data, err := os.ReadFile(sm.filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &sm.stats)
	if err != nil {
		return err
	}

	return nil
}

// saveStats zapisuje statystyki do pliku
func (sm *StatsManager) saveStats() error {
	data, err := json.MarshalIndent(sm.stats, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(sm.filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// AddGameResult dodaje wynik gry do statystyk
func (sm *StatsManager) AddGameResult(word string, result string, points int, difficulty int) error {
	// Utwórz nowe statystyki gry
	gameStats := GameStats{
		Word:       word,
		Result:     result,
		Points:     points,
		Difficulty: difficulty,
		Date:       time.Now(),
	}

	// Aktualizuj statystyki gracza
	sm.stats.GamesPlayed++
	sm.stats.TotalPoints += points

	if result == "win" {
		sm.stats.GamesWon++
	}

	if points > sm.stats.HighestScore {
		sm.stats.HighestScore = points
	}

	// Dodaj statystyki gry do historii
	sm.stats.GameHistory = append(sm.stats.GameHistory, gameStats)

	// Zapisz statystyki do pliku
	return sm.saveStats()
}

// GetStats zwraca statystyki gracza
func (sm *StatsManager) GetStats() PlayerStats {
	return sm.stats
}

// GetWinRate zwraca współczynnik wygranych
func (sm *StatsManager) GetWinRate() float64 {
	if sm.stats.GamesPlayed == 0 {
		return 0
	}
	return float64(sm.stats.GamesWon) / float64(sm.stats.GamesPlayed) * 100
}

// GetAverageScore zwraca średni wynik
func (sm *StatsManager) GetAverageScore() float64 {
	if sm.stats.GamesPlayed == 0 {
		return 0
	}
	return float64(sm.stats.TotalPoints) / float64(sm.stats.GamesPlayed)
}

// GetLastGames zwraca ostatnie n gier
func (sm *StatsManager) GetLastGames(n int) []GameStats {
	history := sm.stats.GameHistory
	historyLen := len(history)

	if historyLen == 0 || n <= 0 {
		return []GameStats{}
	}

	if n > historyLen {
		n = historyLen
	}

	return history[historyLen-n:]
}

// ResetStats resetuje statystyki gracza
func (sm *StatsManager) ResetStats() error {
	sm.stats = PlayerStats{
		GameHistory: []GameStats{},
	}

	return sm.saveStats()
}
