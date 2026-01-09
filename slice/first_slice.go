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
	first_make()
	how_generate_slice()
	slice_of_slice()
	share_memory()
	full_slice()
	convert_array_to_slice()
	slice_copy()
	string_rune_byte()
	convert_to_string()
	convert_from_string()
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

func how_generate_slice() {
	// 「スライスを大きくする回数を減らす」を目的にする

	// nilのまま大きくならない可能性があるなら
	var x []int
	fmt.Println(x == nil)

	// 固定長 or 初期値あり
	data := []int{2, 4, 6, 8}
	fmt.Println(data)

	// 実行すればわかるが、描いてる時はわからない
	// 1. バッファとしてスライスを使うときは長さ1以上
	// 2. サイズが正確にわかる時はそのサイズを指定
	// 3. それ以外の時は長さ0、キャパ指定 → appendによる拡張回数を減らす
}

func slice_of_slice() {
	// スライスのスライスと言いながら切り出しているだけ？
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Println("x:", x)
	fmt.Println("y", y)
	fmt.Println("z", z)
	fmt.Println("d", d)
	fmt.Println("e", e)
}

func share_memory() {
	// スライスからサブスライスを切り出す時、メモリは共有される
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	x[1] = 20
	y[0] = 10
	z[1] = 30

	fmt.Println("x:", x)
	fmt.Println("y", y)
	fmt.Println("z", z)

	// appendを利用しても同様
	// 追加位置は対象のインデックス+1なので、コピー元のスライスは途中の値がいきなり変わるように見える
	a := []int{1, 2, 3, 4}
	b := a[:2]
	fmt.Println(cap(a), cap(b))

	b = append(b, 30)
	fmt.Println("a:", a)
	fmt.Println("b", b)

	c := make([]int, 0, 5)
	c = append(c, 1, 2, 3, 4)
	d := c[:2]
	e := c[2:]
	// キャパシティは元のスライスのキャパ-開始地点
	// なので、同じ長さでもdとeでキャパが異なる
	fmt.Println(cap(c), cap(d), cap(e))

	d = append(d, 30, 40, 50)
	c = append(c, 50)
	e = append(e, 70)
	fmt.Println("c:", c)
	fmt.Println("d", d)
	fmt.Println("e", e)
}

func full_slice() {
	// フルスライスはサブクラスのキャパ=長さに設定できる
	// その状態でappendを利用すると、appendの仕様上新しいスライスができる
	// appendの仕様とは、キャパオーバーした際に返すスライスが全く新しいスライスであるというもの
	x := make([]int, 5, 10)
	x = append(x, 1, 2, 3, 4)
	y := x[:2:2]
	z := x[2:5:7]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, 30, 40, 50)
	x = append(x, 60)
	z = append(z, 70)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func convert_array_to_slice() {
	x := [...]int{1, 2, 3, 4}
	y := x[:2]
	z := x[:]

	// 別に元が配列でもメモリは共有される
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}

func slice_copy() {
	// copy(コピー先のスライス, コピー元のスライス) 戻り値はコピーできた要素の数
	// つまり、コピー先のスライスが小さくともエラーにはならない
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num)

	z := make([]int, 2)
	// 切り出してコピーも可能
	// 戻り値は必ずしも利用する必要はない
	copy(z, x[1:])
	fmt.Println(z)

	// セルフでの上書きも可能
	// これが可能なのはスライスのスライスがメモリを共有しているため（メモリ共有のメリット）
	num = copy(x[:3], x[1:])
	fmt.Println(x, num)

	// 配列のスライスを取ることで配列に対しても可能
	// 反対を言えば配列をそのまま利用することはできない
	x = []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y = make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
}

func string_rune_byte() {
	// 文字列はバイト列
	// 原則文字コードを仮定していないが、ライブラリの中にはUTF-8を前提にしているものもある
	var s string = "Hello there"
	var b byte = s[6]
	fmt.Println(b)

	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]
	fmt.Println(s2, s3, s4)

	// UTF-8は1~4バイトなので、スライスした結果一つの文字の途中で分割してしまう可能性がある
	s = "Hello ☀️"
	s2 = s[4:7]
	s3 = s[:5]
	s4 = s[6:]
	// s2 は太陽絵文字の最初のバイトだけ取るため、 ? になる
	fmt.Println(s2, s3, s4)
	// lenを利用するとバイト長を取得できる
	// 太陽の絵文字は3バイト
	fmt.Println(len(s))
}

func convert_to_string() {
	// rune, byteはstringに変換できるが、intを変換すると文字コードとなるため変換できない
	// go vet の警告対象となる
	var a rune = 'x'
	var s string = string(a)
	var b byte = 'y'
	var s2 string = string(b)
	fmt.Println(s)
	fmt.Println(s2)
}

func convert_from_string() {
	// 反対の変換も可能
	var s string = "Hello, ☀"
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
}