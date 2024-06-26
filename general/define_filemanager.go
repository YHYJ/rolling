/*
File: define_filemanager.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-02-22 14:15:50

Description: 文件管理
*/

package general

import (
	"bufio"
	"os"
	"strings"

	"github.com/gookit/color"
)

// ReadFileLine 读取文件指定行
//
// 参数：
//   - file: 文件路径
//   - line: 行号
//
// 返回：
//   - 指定行的内容
func ReadFileLine(file string, line int) string {
	// 打开文件
	text, err := os.Open(file)
	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s -> Unable to open file: %s\n", DangerText("Error:"), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
	defer text.Close()

	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 行计数
	count := 1
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if line == count {
			return scanner.Text()
		}
		count++
	}
	return ""
}

// ReadFileKey 读取文件包含关键字的行
//
// 参数：
//   - file: 文件路径
//   - key: 关键字
//
// 返回：
//   - 包含关键字的行的内容
func ReadFileKey(file, key string) string {
	// 打开文件
	text, err := os.Open(file)
	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s -> Unable to open file: %s\n", DangerText("Error:"), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
	defer text.Close()

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

// ReadFileCount 获取文件包含关键字的行的计数
//
// 参数：
//   - file: 文件路径
//   - key: 关键字
//
// 返回：
//   - 包含关键字的行的数量
func ReadFileCount(file, key string) int {
	// 打开文件
	text, err := os.Open(file)
	if err != nil {
		fileName, lineNo := GetCallerInfo()
		color.Printf("%s %s -> Unable to open file: %s\n", DangerText("Error:"), SecondaryText("[", fileName, ":", lineNo+1, "]"), err)
	}
	defer text.Close()

	// 创建一个扫描器对象按行遍历
	scanner := bufio.NewScanner(text)
	// 计数器
	count := 0
	// 逐行读取，输出指定行
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), key) {
			count++
		}
	}
	return count
}

// FileExist 判断文件是否存在
//
// 参数：
//   - filePath: 文件路径
//
// 返回：
//   - 文件存在返回 true，否则返回 false
func FileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
