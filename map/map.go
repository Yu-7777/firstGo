package main

import "fmt"

func first_map() {
	// カッコ内がキーの型、外が値の型
	var nilMap map[string]int
	fmt.Println(len(nilMap))
	fmt.Println(nilMap["abc"])
	// nilMap["abc"] = 3 といきなり書くとパニックになる
	// つまりPythonのように扱うことはできない
}

func map_literal() {
	// マップリテラルを利用した宣言
	// これだと書き込みが可能になる
	totalWins := map[string]int{}
	fmt.Println(totalWins == nil)
	fmt.Println(totalWins["abc"])
	totalWins["abc"] = 3;
	fmt.Println(totalWins["abc"])

	// 宣言時に値を入れることもできる
	// linterから値の型指定は不要と言われる
	teams := map[string][]string {
		"ライターズ": []string{"夏目", "森", "国木田"},
		"ナイツ": []string{"武田", "徳川", "明智"},
		"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
	}
	fmt.Println(teams)

	// makeを利用してサイズを指定
	// 恐らくマップ本体が10, 要素は0?
	ages := make(map[string]int, 10)
	fmt.Println(len(ages))
}

func read_and_write() {
	totalWins := map[string]int{}
	totalWins["ライターズ"] = 1
	totalWins["ナイツ"] = 2
	fmt.Println(totalWins["ライターズ"])
	fmt.Println(totalWins["ミュージシャンズ"])
	totalWins["ミュージシャンズ"]++
	fmt.Println(totalWins["ミュージシャンズ"])
	totalWins["ナイツ"] = 3
	fmt.Println(totalWins["ナイツ"])

	// ゼロ値が0なので、インクリメントできる（用途不明）
}

func comma_ok_idiom() {
	// ゼロ値があるので、存在確認専用の関数がある
	// okに真偽値が入る
	m := map[string]int{
		"hello": 5,
		"world": 0,
	}

	v, ok := m["hello"]
	fmt.Println(v, ok)

	v, ok = m["world"]
	fmt.Println(v, ok)

	v, ok = m["goodbye"]
	fmt.Println(v, ok)
}

func delete_from_map() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}

	// delete(マップ, キー) で削除
	delete(m, "hello")
}

func set_by_map() {
	// キーだけを利用することでセットとして運用している？
	// キーの重複は許されないため
	intSet := map[int]bool{}
	vals := []int{15, 10, 2, 5, 8, 7, 3, 9, 1, 2 ,10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[3] {
		fmt.Println("100は含まれている")
	}
	if intSet[10] {
		fmt.Println("10は含まれている")
	}
}

func main() {
	first_map()
}