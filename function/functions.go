/*
File: functions.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-22 14:15:50

Description: 自定义函数
*/

package function

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

// 读取文件指定行
func ReadFileLine(file string, line int) string {
	// 打开文件
	text, err := os.Open(file)
	// 相当于Python的with语句
	defer text.Close()
	// 处理错误
	if err != nil {
		log.Println(err)
	}
	// 行计数
	count := 1
	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if line == count {
			return scanner.Text()
		}
		count++
	}
	return ""
}

// 读取文件包含指定字符串的行
func ReadFileKey(file, key string) string {
	// 打开文件
	text, err := os.Open(file)
	// 相当于Python的with语句
	defer text.Close()
	// 处理错误
	if err != nil {
		log.Println(err)
	}
	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			return scanner.Text()
		}
	}
	return ""
}

// 获取文件包含指定字符串的行的计数
func ReadFileCount(file, key string) int {
	// 打开文件
	text, err := os.Open(file)
	// 相当于Python的with语句
	defer text.Close()
	// 处理错误
	if err != nil {
		log.Println(err)
	}
	// 计数器
	count := 0
	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			count++
		}
	}
	return count
}

// 运行指定命令
func RunCommand(command, args string) string {
	// 定义命令
	cmd := exec.Command(command, args)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout // 标准输出
	cmd.Stderr = &stderr // 标准错误

	// 执行命令获取输出
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())

	if err != nil {
		log.Printf("Run failed with %s", errStr)
	}

	return strings.TrimRight(outStr, "\n")
}
