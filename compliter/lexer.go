package compliter

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

type TokenType int

const (
	TokenIndent TokenType = iota
	TokenDedent
	TokenText
	TokenNewline
	TokenEOF
)

type Token struct {
	Type  TokenType
	Value string
	Line  int
}

type LexerState struct {
	Tokens []Token
	Stack  []int
	Line   int
}

func NewLexerState() LexerState {
	return LexerState{
		Stack: []int{0},
		Line:  0,
	}
}

func TokenizeFile(filename string) ([]Token, error) {
	file, err := os.Open(filepath.Join(".", filename))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	state := NewLexerState()

	for scanner.Scan() {
		state.Line++
		line := scanner.Text()
		state = processLine(state, line)
	}

	// 关闭所有缩进
	for len(state.Stack) > 1 {
		state.Tokens = append(state.Tokens, Token{TokenDedent, "", state.Line})
		state.Stack = state.Stack[:len(state.Stack)-1]
	}

	state.Tokens = append(state.Tokens, Token{TokenEOF, "", state.Line})
	return state.Tokens, nil
}

func processLine(state LexerState, line string) LexerState {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return state
	}

	indent := 0
	for _, r := range line {
		if !unicode.IsSpace(r) {
			break
		}
		indent++
	}

	newState := state
	top := newState.Stack[len(newState.Stack)-1]

	switch {
	case indent > top:
		newState.Tokens = append(newState.Tokens, Token{TokenIndent, "", newState.Line})
		newState.Stack = append(newState.Stack, indent)
	case indent < top:
		for indent < top {
			newState.Tokens = append(newState.Tokens, Token{TokenDedent, "", newState.Line})
			newState.Stack = newState.Stack[:len(newState.Stack)-1]
			top = newState.Stack[len(newState.Stack)-1]
		}
	}

	newState.Tokens = append(newState.Tokens, Token{TokenText, trimmed, newState.Line})
	return newState
}
