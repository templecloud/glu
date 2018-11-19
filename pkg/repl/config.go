package repl


// Debug ======================================================================
//

type debug struct {
	tokenHeader    bool
	token          bool
	tokenErrHeader bool
	tokenErr       bool
	parseErrHeader bool
	parseErr       bool
	exprHeader     bool
	expr           bool
	resultHeader   bool
	result         bool
}

func noDebug() debug {
	return debug{
		tokenHeader:    false,
		token:          false,
		tokenErrHeader: false,
		tokenErr:       false,
		parseErrHeader: false,
		parseErr:       false,
		exprHeader:     false,
		expr:           false,
		resultHeader:   false,
		result:         true,
	}
}

func defaultDebug() debug {
	return debug{
		tokenHeader:    false,
		token:          false,
		tokenErrHeader: false,
		tokenErr:       false,
		parseErrHeader: true,
		parseErr:       true,
		exprHeader:     false,
		expr:           false,
		resultHeader:   false,
		result:         true,
	}
}

func fullDebug() debug {
	return debug{
		tokenHeader:    true,
		token:          true,
		tokenErrHeader: true,
		tokenErr:       true,
		parseErrHeader: true,
		parseErr:       true,
		exprHeader:     true,
		expr:           true,
		resultHeader:   true,
		result:         true,
	}
}

// Config =====================================================================
//

// Config represents configuration option for the Lexer.
type config struct {
	// The debug configuration. Enables and disables debug output.
	debug
}

func defaultConfig() config {
	return config{debug: defaultDebug()}
}
