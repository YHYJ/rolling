/*
File: define_pacman.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-15 10:49:08

Description: 解析 pacman 日志文件
*/

package general

import (
	"os"
	"strings"
)

// GetSystemUpdateCount 获取系统更新次数
//
// 参数：
//   - file: 记录系统更新信息的文件的路径
//
// 返回：
//   - 系统更新次数
//   - 错误信息
func GetSystemUpdateCount(file string) (int, error) {
	// 读取日志文件内容
	content, err := os.ReadFile(file)
	if err != nil {
		return 0, err
	}

	// 初始化计数器
	updateCounter := 0

	// 标记是否在更新事务中
	inTransaction := false

	// 标记上一次出现的 'upgraded' 行的索引
	var lastUpgradedIndex int = -1

	// 将文件内容按行分割
	lines := strings.Split(string(content), "\n")
	// 逐行解析
	for index, line := range lines {
		// 查找开始系统更新的关键字
		if strings.Contains(line, "starting full system upgrade") {
			inTransaction = true
			continue
		}

		// 在更新事务中，查找事务开始的关键字
		if inTransaction && strings.Contains(line, "transaction started") {
			continue
		}

		// 在更新事务中，查找事务完成的关键字
		if inTransaction && strings.Contains(line, "transaction completed") {
			inTransaction = false

			// 如果上次出现的 'upgraded' 行的索引在事务内，增加成功更新的计数
			if lastUpgradedIndex != -1 && lastUpgradedIndex < index {
				updateCounter++
			}
			continue
		}

		// 在更新事务中，查找升级成功的关键字
		if inTransaction && strings.Contains(line, "upgraded") {
			lastUpgradedIndex = index
		}
	}

	return updateCounter, nil
}
