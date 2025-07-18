package compliter

import (
	"fmt"
	"strings"
)

// ASTNode 接口定义
type ASTNode interface {
	String() string
}

// Program 程序根节点
type Program struct {
	Statements []ASTNode
}

func (p Program) String() string {
	var sb strings.Builder
	for _, stmt := range p.Statements {
		sb.WriteString(stmt.String() + "\n")
	}
	return sb.String()
}

// IfStmt if语句节点
type IfStmt struct {
	Condition string
	Body      []ASTNode
	Line      int
}

func (i IfStmt) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[Line %d] IF %s:\n", i.Line, i.Condition))
	for _, node := range i.Body {
		sb.WriteString("  " + node.String() + "\n")
	}
	return strings.TrimSpace(sb.String())
}

// PrintStmt print语句节点
type PrintStmt struct {
	Text string
	Line int
}

func (p PrintStmt) String() string {
	return fmt.Sprintf("[Line %d] PRINT %s", p.Line, p.Text)
}
