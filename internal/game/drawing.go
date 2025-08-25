package game

// HangmanDrawing zawiera rysunki wisielca dla różnych etapów
type HangmanDrawing struct {
	Frames []string
}

// NewHangmanDrawing tworzy nowy obiekt do rysowania wisielca
func NewHangmanDrawing() *HangmanDrawing {
	frames := []string{
		// 0 błędów - tylko szubienica
		`
    ╔════════╗
    ║        ║
    ║        
    ║        
    ║        
    ║        
    ║        
 ╔══╩══╗     
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 1 błąd - głowa
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚╤╝
    ║        
    ║        
    ║        
 ╔══╩══╗     
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 2 błędy - tułów
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚╤╝
    ║        ┃
    ║        ┃
    ║        
 ╔══╩══╗     
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 3 błędy - lewa ręka
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚╤╝
    ║       ┏┃
    ║       ┃┃
    ║        
 ╔══╩══╗     
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 4 błędy - prawa ręka
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚╤╝
    ║       ┏┃┓
    ║       ┃┃┃
    ║        
 ╔══╩══╗     
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 5 błędów - lewa noga
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚╤╝
    ║       ┏┃┓
    ║       ┃┃┃
    ║       ┏┻┓
 ╔══╩══╗     
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 6 błędów - prawa noga
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚╤╝
    ║       ┏┃┓
    ║       ┃┃┃
    ║       ┏┻┓
 ╔══╩══╗    ┃ ┃
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 7 błędów - dodatkowe detale, widać strach (dla poziomu łatwego)
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚O╝
    ║       ┏┃┓
    ║       ┃┃┃
    ║       ┏┻┓
 ╔══╩══╗    ┃ ┃
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
		// 8 błędów - wisielec umiera (dla poziomu łatwego)
		`
    ╔════════╗
    ║        ║
    ║       ╔╧╗
    ║       ╚X╝
    ║       ┏┃┓
    ║       ┃┃┃
    ║       ┏┻┓
 ╔══╩══╗    ┃ ┃
 ║     ║     
╔╩═════╩╗    
╚═══════╝    `,
	}

	return &HangmanDrawing{Frames: frames}
}

// GetDrawing zwraca rysunek wisielca dla określonej liczby błędów
func (hd *HangmanDrawing) GetDrawing(wrongAttempts int) string {
	if wrongAttempts < 0 {
		wrongAttempts = 0
	}

	if wrongAttempts >= len(hd.Frames) {
		wrongAttempts = len(hd.Frames) - 1
	}

	return hd.Frames[wrongAttempts]
}

// GetAllDrawings zwraca wszystkie rysunki wisielca
func (hd *HangmanDrawing) GetAllDrawings() []string {
	return hd.Frames
}
