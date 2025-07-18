package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func read_flip() {
	fmt.Println("in read compliter")
	// 打开同目录下的文件
	file, err := os.Open("test.compliter")
	if err != nil {
		log.Fatal(err) // 文件不存在或权限错误时终止
	}
	defer file.Close() // 确保文件关闭

	// 创建带缓冲的读取器（高效处理大文件）
	scanner := bufio.NewScanner(file)

	// 逐行读取
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // 输出每一行
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("读取错误:", err)
	}
}
