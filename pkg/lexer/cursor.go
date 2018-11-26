package lexer

// Lexical Cursor Functions ===================================================
//

// Return true if the input has been fully consumed; false otherwise.
func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.input)
}

// Advance the current cursor forward one character and return
// the original underlying character.
func (l *Lexer) advance() rune {
	l.current++
	l.column++
	return l.input[l.current-1]
}

// Check if the current character matches the expected character
// and if true consume it by incrementing the cursor.
//
// If the current cursor character matched the expected character
// then advance the current cursor one step and return true; else
// false.
func (l *Lexer) matches(expected rune) bool {
	if l.isAtEnd() {
		return false
	}
	if l.input[l.current] != expected {
		return false
	}
	l.current++
	l.column++
	return true
}

// Peek at the current character without advancing the cursor.
//
// If the cursor reaches the end of the input return the nil character;
// else return the current character.
func (l *Lexer) peek() rune {
	if l.isAtEnd() {
		return nilByte
	}
	// TODO: Make safe for index
	return l.input[l.current]
}

// Peek at the next character without advancing the cursor.
//
// If the cursor reaches the end of the input return the nil character;
// else return the next character.
func (l *Lexer) peekNext() rune {
	if l.current+1 >= len(l.input) {
		return nilByte
	}
	return l.input[l.current+1]
}