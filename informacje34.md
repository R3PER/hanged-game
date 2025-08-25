# Informacje o projekcie "Wisielec"

## Opis projektu
Projekt "Wisielec" to rozbudowana implementacja klasycznej gry "Wisielec" działająca w konsoli. Jest to interaktywna gra, w której gracz musi odgadnąć ukryte słowo, podając litery. Za każdym razem, gdy gracz poda literę nieobecną w słowie, dorysowywana jest część rysunku wisielca. Gra kończy się wygraną, gdy gracz odgadnie całe słowo, lub przegraną, gdy rysunek wisielca zostanie ukończony.

## Język programowania
Projekt został napisany w języku Go (Golang) w wersji 1.24.4.

## Użyte technologie
- **Go** - nowoczesny język programowania o składni zbliżonej do C, ze świetnym wsparciem dla programowania współbieżnego
- **Interfejs konsolowy** - aplikacja działa w trybie tekstowym w konsoli/terminalu
- **System modułów Go** - projekt wykorzystuje system modułów wprowadzony w Go 1.11
- **Obsługa plików JSON** - do przechowywania statystyk graczy
- **Obsługa plików tekstowych** - do przechowywania bazy słów oraz preferencji językowych

## Funkcjonalności
- Kolorowy interfejs konsolowy
- Wizualizacja wisielca w ASCII
- Różne poziomy trudności (łatwy, średni, trudny)
- System punktacji z różnymi bonusami
- Zapisywanie i wyświetlanie statystyk gracza
- Obsługa polskich znaków
- Baza słów do zgadywania
- Wielojęzyczność (polski i angielski)
- Elementy RPG:
  - System poziomów i doświadczenia
  - Ekwipunek postaci
  - System zadań (questów)
  - Sklep z przedmiotami

## Struktura projektu
Projekt ma uporządkowaną strukturę katalogów zgodną z dobrymi praktykami w Go:

```
hanged-game/
├── cmd/                  # Punkt wejściowy aplikacji
│   └── main.go           # Główny plik wykonawczy
├── internal/             # Kod wewnętrzny projektu
│   ├── game/             # Logika gry
│   │   ├── game.go       # Główna logika gry
│   │   ├── drawing.go    # Rysowanie wisielca
│   │   ├── words.go      # Zarządzanie słowami
│   │   └── rpg.go        # System RPG
│   ├── ui/               # Interfejs użytkownika
│   │   ├── console.go    # Obsługa konsoli
│   │   ├── keyboard.go   # Obsługa klawiatury
│   │   └── rpg_ui.go     # UI dla elementów RPG
│   ├── localization/     # Obsługa wielu języków
│   │   └── translations.go # Tłumaczenia tekstów
│   └── storage/          # Zapisywanie/odczytywanie danych
│       └── stats.go      # Zapisywanie statystyk
├── data/                 # Dane aplikacji
│   ├── words.txt         # Plik z bazą słów
│   └── stats.json        # Plik ze statystykami
├── go.mod                # Definicja modułu Go
└── README.md             # Instrukcje i dokumentacja
```

## Wymagania systemowe
- Go 1.16 lub nowszy
- System operacyjny obsługujący kolorowe wyjście w terminalu (Windows, macOS, Linux)

## Jak uruchomić
1. Sklonuj repozytorium
2. Zbuduj projekt: `go build -o wisielec ./cmd`
3. Uruchom skompilowany program: `./wisielec`
