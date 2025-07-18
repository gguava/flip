package compliter

import "strings"

func Parse(tokens []Token) (Program, error) {
	program, _, err := parseProgram(tokens, 0)
	return program, err
}

func parseProgram(tokens []Token, pos int) (Program, int, error) {
	var stmts []ASTNode
	var err error

	for pos < len(tokens) && tokens[pos].Type != TokenEOF {
		var node ASTNode
		node, pos, err = parseStatement(tokens, pos)
		if err != nil {
			return Program{}, pos, err
		}
		if node != nil { // 过滤空节点
			stmts = append(stmts, node)
		}
	}

	return Program{Statements: stmts}, pos, nil
}

func parseStatement(tokens []Token, pos int) (ASTNode, int, error) {
	if pos >= len(tokens) {
		return nil, pos, nil
	}

	tok := tokens[pos]
	switch {
	case tok.Value == "if":
		return parseIfStmt(tokens, pos)
	case strings.HasPrefix(tok.Value, "print"):
		return PrintStmt{
			Text: tok.Value,
			Line: tok.Line,
		}, pos + 1, nil
	default:
		return nil, pos + 1, nil // 跳过无法识别的Token
	}
}

func parseIfStmt(tokens []Token, pos int) (IfStmt, int, error) {
	startLine := tokens[pos].Line
	pos++ // 跳过 'if'

	// 解析条件
	var condParts []string
	for pos < len(tokens) && tokens[pos].Value != ":" {
		condParts = append(condParts, tokens[pos].Value)
		pos++
	}
	cond := strings.Join(condParts, " ")
	pos++ // 跳过 ':'

	// 检查缩进
	if pos >= len(tokens) || tokens[pos].Type != TokenIndent {
		return IfStmt{}, pos, nil
	}
	pos++

	// 解析语句块
	var body []ASTNode
	for pos < len(tokens) && tokens[pos].Type != TokenDedent {
		node, newPos, err := parseStatement(tokens, pos)
		if err != nil {
			return IfStmt{}, pos, err
		}
		if node != nil {
			body = append(body, node)
		}
		pos = newPos
	}

	if pos < len(tokens) && tokens[pos].Type == TokenDedent {
		pos++
	}

	return IfStmt{
		Condition: cond,
		Body:      body,
		Line:      startLine,
	}, pos, nil
}
