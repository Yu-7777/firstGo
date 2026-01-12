package main

import "fmt"

func first_struct() {
	type person struct {
		name string
		age int
		pet string
	}

	// マップと異なってvarとリテラルの差異は無い
	// どちらもゼロ値で初期化される
	var fred person
	bob := person{}
	fmt.Println(fred)
	fmt.Println(bob)

	// フィールド名も指定可能
	julia := person{
		"ジュリア",
		40,
		"cat",
	}

	beth := person{
		age: 30,
		name: "ベス",
	}

	fmt.Println(julia)
	fmt.Println(beth)
}

func noname_struct() {
	// 外部データからの受取や出力に利用される（具体的に想像はできないが、型を定義するほどでもないって時にコードを短縮することが目的？）
	var person struct {
		name string
		age int
		pet string
	}

	person.name = "ボブ"
	person.age = 50
	person.pet = "dog"

	pet := struct {
		name string
		kind string
	}{
		name: "ポチ",
		kind: "dog",
	}
	fmt.Println(person)
	fmt.Println(pet)
}

func convert_and_compare() {
	// 変換は同じ型かつフィールドが比較可能のときのみ
	// 同じフィールドを持つ場合に型変換が可能（名前も順も型も同じ場合のみ）
	type firstPerson struct {
		name string
		age int
	}

	type secondPerson struct {
		name string
		age int
	}

	f := firstPerson{
		name: "ボブ",
		age: 50,
	}
	var g struct {
		name string
		age int
	}

	// どちらかが無名構造体の場合は型変換なし（無名構造体のまま比較が可能）
	g = f
	fmt.Println(f == g)
}

func main() {
	first_struct()
	noname_struct()
	convert_and_compare()
}