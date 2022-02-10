package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(printNumWith2(12.3456789))
	fmt.Println(printBytes("小可"))
}

// 输出两位小数
func printNumWith2(value float64) float64 {
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return val

}

//将[]byte输出为16进制
func printBytes(data string) string {
	// 转换的用的 byte数据
	byteData := []byte(data)
	// 将 byte 装换为 16进制的字符串
	hexStringData := hex.EncodeToString(byteData)

	// 将 16进制的字符串 转换 byte
	//hex_data, _ := hex.DecodeString(hex_string_data)
	// 将 byte 转换 为字符串 输出结果
	//println(string(hex_data))

	return hexStringData

}
