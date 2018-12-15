# Todo

Things to do next.

---

#### Global
* Refactor tests to share set-up code.
* Refactor Expr type names.
* Collapse separate expr/stmt files into one file.
* Add file reading.
* Add function attributes - bash exec mode.
* Rename fn to func.
* Mechanism for casting string to and from their natural types.
* Harmonise printer visitor.
* Replace log statement with NIFs.
* Implement a debugger.
* Tidy lexer, parser, interpreter errors.
* Pretty print parser/interpreter errors.

---

#### Lexer
* Improve representation.
* Review current Lexemes.

---

#### AST + Parser
* Add recovery' function.
* Add more tests.
* Add 'numeral' node type.
* Remove log statement and add printing functions.

---

#### Interpreter
* Allow a runtime error to be returned as a result.
* Handle 'division by 0'. 
* Add 'numeral' node type.


---

#### Repl
* Add command line options to configure token and expression output.
* Add functions to the repl to allow the options to be reconfigured from within
  the repl.
* Create separate repl and