package main

import (
	"flip/compliter" // 替换为实际模块路径
	"fmt"
)

func main() {
	tokens, err := compliter.TokenizeFile("test.flip")
	if err != nil {
		fmt.Println("Lex error:", err)
		return
	}

	ast, err := compliter.Parse(tokens)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	fmt.Println("=== AST ===")
	fmt.Println(ast.String())
}
