package lexer

// Config =====================================================================
//

// Config represents configuration option for the Lexer.
type Config struct {
	// If true, then add an EOF token to the end of the stream if one does not
	// exist already.
	autoEOFToken bool
}

func defaultConfig() Config {
	return Config{autoEOFToken: true}
}
