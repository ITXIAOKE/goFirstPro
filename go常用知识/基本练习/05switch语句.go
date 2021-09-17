package main

import "fmt"

func main() {
	fmt.Println(
	grade(0),
	grade(50),
	grade(66),
	grade(77),
	grade(88),
	grade(99),
	grade(100),
	grade(-10),
	grade(106),
	grade(105),
	)
}

func grade(score int ) string {
	g:=""
	switch {
	case score<60:
		g="不及格"
	case score<70:
		g="及格"
	case score<80:
		g="中等"
	case score<90:
		g="良好"
	case score<=100:
		g="优秀"
	case score<0||score>100:
		panic(fmt.Sprintf("错误分数：%d",score))

	}
	return g
}
