# Hanged – Console Game

An advanced implementation of the classic 'Hangman' game running in the console. It's an interactive game with RPG elements, where the player has to guess a hidden word by providing letters. The game features a colorful interface, level and experience system, inventory, and an item shop.

## Screens
<img width="801" height="554" alt="1" src="https://github.com/user-attachments/assets/554a6127-109b-4cb6-a416-3dfb8a49b656" />
<img width="903" height="582" alt="2" src="https://github.com/user-attachments/assets/fa8d8e6e-af1c-44a2-94d4-fc04a562f563" />
<img width="989" height="595" alt="3" src="https://github.com/user-attachments/assets/31ec03fe-d754-42de-9cd6-13f23450de4e" />
<img width="903" height="601" alt="4" src="https://github.com/user-attachments/assets/4fe44685-c4bb-4d97-b2e8-cbe735dbedcd" />
<img width="845" height="576" alt="5" src="https://github.com/user-attachments/assets/1bfe1310-5876-4648-bbc7-97616b3a5984" />
<img width="895" height="676" alt="6" src="https://github.com/user-attachments/assets/dea4632e-1d4c-477f-884d-3ae7c80f07f9" />
<img width="847" height="673" alt="7" src="https://github.com/user-attachments/assets/bfdbf97c-7628-45d4-8edb-a35cb13556ce" />



## Features

- Colorful console interface  
- ASCII hangman visualization  
- Multiple difficulty levels (easy, medium, hard)  
- Scoring system  
- Player statistics saving and display  
- Polish characters support  
- Word database for guessing  

## Requirements

- Go 1.16 or newer  

## Installation

1. Clone the repository:
```
git clone https://github.com/R3PER/hanged-game.git
cd hanged-game
```

2. Build the project:
```
go build -o hangman ./cmd
```

## Running

Run the compiled program:
```
./hangman
```

## How to Play

1. Select "New Game" from the main menu.  
2. Choose a difficulty level (or use the default).  
3. Try to guess the word by entering single letters.  
4. Each time you enter a letter that is not in the word, a new part of the hangman is drawn.  
5. The game ends with a win if you guess the whole word, or with a loss if the hangman drawing is completed.  

## Scoring Rules

- +10 points for each correctly guessed letter  
- -5 points for each wrong guess  
- +50 bonus points for winning the game  
- +5 points for each unused attempt  

## Project Structure

```
hanged-game/
├── cmd/
│   └── main.go          # Application entry point
├── internal/
│   ├── game/            # Game logic
│   │   ├── game.go      # Main game logic
│   │   ├── drawing.go   # Hangman drawing
│   │   └── words.go     # Word management
│   ├── ui/              # User interface
│   │   └── console.go   # Console handling
│   └── storage/         # Data saving/loading
│       └── stats.go     # Statistics saving
├── data/
│   └── words.txt        # Word database file
├── go.mod               # Go module definition
└── README.md            # Instructions and documentation
```

## Expanding the Word Database

You can add your own words to the `data/words.txt` file, one word per line.

## License

This project is licensed under the MIT License.
