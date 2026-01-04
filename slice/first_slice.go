package main

import "fmt"

func main() {
	// [n]、[...]は配列になる
	// 大きさを指定しない
	var x = []int{10, 20, 30}
	fmt.Println(x)

	// 添字利用
	var y = []int{1, 5: 4, 6, 10: 100, 15}
	fmt.Println(y)

	x[0] = 40
	fmt.Print(x[2])

	// 多次元が可能
	var z [][]int
	fmt.Println(z)

	// スライスのゼロ値は 0 ではなく nil
	// nil は異なる方への代入や比較が可能
	var a []int
	fmt.Println(a == nil)

	// スライス自体は比較不能

	first_len()
	first_append()
	first_cap()
}

func first_len() {
	var x []int
	fmt.Println(len(x))
}

// スライスへの要素追加
func first_append() {
	var x []int
	x = append(x, 10)
	fmt.Println(x)

	var y = []int{1, 2, 3}
	y = append(y, 4)
	fmt.Println(y)

	// 複数追加が可能
	x = append(x, 5, 6, 7)
	fmt.Println(x)

	// yを展開して個別に追加
	x = append(x, y...)
	fmt.Println(x)

	// appendした結果キャパを超えるとより大きい容量を確保したスライスにコピー＋追加され、それが返される
	// つまり、内部的には拡張されずに生まれ変わる？
	// そのキャパを確認するために cap 関数が存在する
}

func first_cap() {
	var x []int
	fmt.Println(x, len(x), cap(x))

	x = append(x, 10)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 20)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 30)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 40)
	fmt.Println(x, len(x), cap(x))

	x = append(x, 50)
	fmt.Println(x, len(x), cap(x))

	// cap は実際のデータ長よりも余裕を持って確保される
	// だが、分かってるなら最初から決定した方がいい → make 関数
}

func first_make() {
	// 長さ、キャパ5のスライス
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x))

	// 長さ=キャパのスライスにappendを利用するとキャパオーバーが必至で無駄
	y := make([]int, 5)
	y = append(y, 10)
	fmt.Println(y, len(y), cap(y))

	// 長さとキャパを個別に指定
	z := make([]int, 3, 5)
	fmt.Println(z, len(z), cap(z))

	// これならappendを利用してもキャパオーバーしない
	z = append(z, 10)
	fmt.Println(z, len(z), cap(z))
}