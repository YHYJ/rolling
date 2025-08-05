/*
File: define_math.go
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-24 13:35:22

Description: 处理数学计算
*/

package general

import (
	"math"

	"github.com/gookit/color"
)

// RoundFloat32 返回保留指定位小数的 float32 类型数字
//
// 参数：
//   - number: 数字，类型为 float32
//
// 返回：
//   - 保留指定位小数的数字
func RoundFloat32(number float32, precision int) float32 {
	pow10_n := math.Pow10(precision) // 10 的 n 次方
	return float32(math.Round(float64(number)*pow10_n) / pow10_n)
}

// FindFakeMaxLength 计算多个 int, float 类型数字直接转为 string 后最长者长度
//
// 参数：
//   - numbers: 数字列表
//
// 返回：
//   - 最长长度
func FindFakeMaxLength(numbers []any) int {
	maxLength := 0

	for _, num := range numbers {
		str := color.Sprintf("%v", num)
		length := len(str)

		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
