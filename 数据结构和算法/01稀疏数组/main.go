package main

import "fmt"

type valNode struct {
	row int
	col int
	val int //interface{}
}

func main() {
	//1,创建原始数组
	var chessMap [11][11]int

	//2,输出原始的数组
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3，转成稀疏数组
	var sparseArr []valNode //切片

	valNode2 := valNode{ //初始值
		row: 11,
		col: 11,
		val: 0}

	sparseArr = append(sparseArr, valNode2)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				valNode3 := valNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode3)
			}
		}
	}

	//4。输出稀疏数组
	fmt.Println("当前的稀疏数组是：：：：：：")
	for i, valNode4 := range sparseArr {
		fmt.Printf("%d:%d %d %d \n", i, valNode4.row, valNode4.col, valNode4.val)
	}

	/**
	恢复稀疏数组数据
	*/
	var chessMap2 [11][11]int
	for i, valNode5 := range sparseArr {
		if i != 0 { //跳过第一行记录，11，11，0，这是初始值
			chessMap2[valNode5.row][valNode5.col] = valNode5.val
		}
	}
	fmt.Println("恢复数据是：：：：：：")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

}
