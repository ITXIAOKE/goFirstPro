package main



import "fmt"

func Num(a int) []int {
	//var res int
	m := []int{}

	for a > 0 {
		m = append(m, a%10) //取余数
		a = a / 10          //除数取值
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m)-i-1; j++ {
			if m[j] > m[j+1] {
				m[j], m[j+1] = m[j+1], m[j]
			}
		}
	}

	return m

}

func main() {
	var a = 1234
	fmt.Println(Num(a))
	//fmt.Println(1234/10)
	//fmt.Println(1234%10)
}
