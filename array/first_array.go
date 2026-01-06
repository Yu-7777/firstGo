package main

import "fmt"

func main() {

	// 空の配列の宣言
	var x [3]int

	// 初期値を設定し宣言
	var y = [3]int{10, 20, 30}

	// 添字利用
	var z = [12]int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(z)

	// 配列リテラル
	var a = [...]int{10, 20, 30}
	fmt.Println(a)

	// 比較
	fmt.Println(x == y)

	// 二次元配列(厳密にGO言語には2次元配列は存在しない)
	var b[2][3]int

	a[0] = 10
	fmt.Println(b[0])

	// 長さ len
	fmt.Println(len(x))

	// 配列の大きさ決定に変数を利用できず、異なる長さへの変換も不可
	// そのためスライスがよく使われる
}
