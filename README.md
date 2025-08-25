# Wisielec - Gra konsolowa

Rozbudowana implementacja klasycznej gry "Wisielec" napisana w języku Go.

## Funkcjonalności

- Kolorowy interfejs konsolowy
- Wizualizacja wisielca w ASCII
- Różne poziomy trudności (łatwy, średni, trudny)
- System punktacji
- Zapisywanie i wyświetlanie statystyk gracza
- Obsługa polskich znaków
- Baza słów do zgadywania

## Wymagania

- Go 1.16 lub nowszy

## Instalacja

1. Sklonuj repozytorium:
```
git clone https://github.com/r3per/hanged-game.git
cd hanged-game
```

2. Zbuduj projekt:
```
go build -o wisielec ./cmd
```

## Uruchomienie

Uruchom skompilowany program:
```
./wisielec
```

## Jak grać

1. Wybierz opcję "Nowa gra" z menu głównego.
2. Wybierz poziom trudności (lub użyj domyślnego).
3. Próbuj odgadnąć słowo, podając pojedyncze litery.
4. Za każdym razem, gdy podasz literę, która nie występuje w słowie, dostajesz jedną część wisielca.
5. Gra kończy się wygraną, gdy odgadniesz całe słowo, lub przegraną, gdy rysunek wisielca zostanie ukończony.

## Zasady punktacji

- +10 punktów za każdą odgadniętą literę
- -5 punktów za każdą błędną próbę
- +50 punktów bonus za wygraną grę
- +5 punktów za każdą niewykorzystaną próbę

## Struktura projektu

```
hanged-game/
├── cmd/
│   └── main.go          # Punkt wejściowy aplikacji
├── internal/
│   ├── game/            # Logika gry
│   │   ├── game.go      # Główna logika gry
│   │   ├── drawing.go   # Rysowanie wisielca
│   │   └── words.go     # Zarządzanie słowami
│   ├── ui/              # Interfejs użytkownika
│   │   └── console.go   # Obsługa konsoli
│   └── storage/         # Zapisywanie/odczytywanie danych
│       └── stats.go     # Zapisywanie statystyk
├── data/
│   └── words.txt        # Plik z bazą słów
├── go.mod               # Definicja modułu Go
└── README.md            # Instrukcje i dokumentacja
```

## Rozszerzanie bazy słów

Możesz dodać własne słowa do pliku `data/words.txt`, każde słowo w nowej linii.

## Licencja

Ten projekt jest licencjonowany na zasadach licencji MIT.
