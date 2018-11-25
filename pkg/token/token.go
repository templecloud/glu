package token

// Source represents the position of a token in a source file or stream.
type Source struct {
	Origin string
	Line   int
	Column int
	Length int
}

// Tokens =====================================================================
//

// Type represents the type of the token.
type Type string

// Token represents a lexeme in a source file or stream.
type Token struct {
	Type
	Lexeme string
	Source
}

// New creates a new token.
func New(
	tokenType Type,
	lexeme string,
	origin string,
	line int,
	column int,
	length int,
) *Token {
	return &Token{
		tokenType,
		lexeme,
		Source{
			Origin: origin,
			Line:   line,
			Column: column,
			Length: length,
		},
	}
}

// Structural tokens.
const (
	LeftParen  = "LeftParen"  // "("
	RightParen = "RightParen" // ")"
	LeftBrace  = "LeftBrace"  // "{"
	RightBrace = "RightBrace" // "}"
	Comma      = "Comma"      // ","
	Dot        = "Dot"        // "."
	Semicolon  = "Semicolon"  // ";"
)

// Arithmetic operators.
const (
	Minus        = "Minus"        // "-"
	Plus         = "Plus"         // "+"
	ForwardSlash = "ForwardSlash" // "/"
	Star         = "Star"         // "*"
)

// Comparison operators.
const (
	Not                = "Not"                // "!"
	NotEqual           = "NotEqual"           // "!="
	Equal              = "Equal"              // "="
	EqualEqual         = "EqualEqual"         // ""=="
	GreaterThan        = "GreaterThan"        // ">"
	GreaterThanOrEqual = "GreaterThanOrEqual" // ">="
	LessThan           = "LessThan"           // "<"
	LessThanOrEqual    = "LessThanOrEqual"    // "<="
)

// Literals.
const (
	Identifier = "Identifier"
	String     = "String"
	Number     = "Number"
)

// Keywords.
const (
	Nil    = "nil"
	True   = "true"
	False  = "false"
	And    = "and" // "&&"
	Or     = "or"  // "||"
	If     = "if"
	Else   = "else"
	While  = "while"
	For    = "for"
	Return = "return"
	Var    = "var"
	Func   = "func"
	Log    = "log"
)

// Special.
const (
	EOF = "EOF"
)
