package main

import "fmt"

/**
少于1024双倍扩容，大于1024是1.25倍扩容 内存增长

1，一个一个的append
1024 * 1.25 = 1280
1024 * 1.25 * 4 = 5120 查表5376/4=1344   【表是sizeclasses.go文件   使用int32	占4个字节】
1024 * 1.25 * 8 = 10240 查表10240/8=1280   【表是sizeclasses.go文件   使用int64	占8个字节】

2，批量的append
1025*8=8200  【查表sizeclasses.go文件   9472/8=1184】
*/
func main() {
	var arr1 []int64
	var arr2 []int64
	for i := 0; i < 1025; i++ {
		arr1 = append(arr1, int64(i))
	}
	fmt.Println(len(arr1), cap(arr1))

	arr2 = append(arr2, arr1...)
	fmt.Println(len(arr2), cap(arr2))

}
