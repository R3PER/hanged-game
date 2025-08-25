# Hanged - Console Game
An advanced implementation of the classic 'Hangman' game running in the console. It's an interactive game with RPG elements, where the player has to guess a hidden word by providing letters. The game features a colorful interface, level and experience system, inventory, and an item shop.


## Screens

<img width="801" height="554" alt="1" src="https://github.com/user-attachments/assets/ee6057eb-adce-4015-8adc-388878809b4b" />
<img width="903" height="582" alt="2" src="https://github.com/user-attachments/assets/a66f77c9-2512-444d-9a4f-33008e655409" />
<img width="989" height="595" alt="3" src="https://github.com/user-attachments/assets/b9cef3f5-59cc-440f-8ea3-311a9a13a91a" />
<img width="903" height="601" alt="4" src="https://github.com/user-attachments/assets/31ee01b0-3374-4ee7-86da-4af62bb81523" />
<img width="845" height="576" alt="5" src="https://github.com/user-attachments/assets/0ad1a781-b601-48f9-b072-2c7054c641c0" />
<img width="895" height="676" alt="6" src="https://github.com/user-attachments/assets/3c3b5dec-ba59-46e6-bd58-88346ce3b0da" />
<img width="847" height="673" alt="7" src="https://github.com/user-attachments/assets/c653d657-3c77-44b4-8f10-f922e589cbdd" />

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
