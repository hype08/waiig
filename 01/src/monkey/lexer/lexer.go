package lexer

type Lexer struct {
	input        string
	position     int  // current position in input, pointing to ch
	readPosition int  // current reading position in input, after ch
	ch           byte // current character
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
