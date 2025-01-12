package main

import (
	"fmt"
	"net/http"
)

func functionAndTyep() {
	println("userTypeの呼び出し成功")

	// var intint =  [...][...]int{{0}}
	// intint[0][0] = 1
	// intint[0][1] = 1
	// fmt.Println(intint)
	// map[string][]int
	
	// type Person struct {
	// 	Name string
	// }

	// Underlying type
	

	// TRY: ユーザ定義型の利用 https://docs.google.com/presentation/d/1DtWB-8FcnNb9asxSpIaOLYbAEc9OjBAwMLNxKnPA8pc/edit#slide=id.g4cbe4d134e_0_254
	// 集計を行うことを目的としてみる
	// type User struct {
	// 	Id int
	// 	Name string
	// }

	// type UserResult struct {
	// 	User User
	// 	Point int
	// }

	// type GameResult struct {
	// 	GameNum int // 何回目のゲームか
	// 	UserResult []UserResult
	// }

	type Applicant = http.Client
	fmt.Printf("%T", Applicant{})

	// fmt.Println("add(1, 2)")
	addition := add(1, 2)
	println(addition) // ←これ出力されないのだが、なぜ??
	// fmt.Println(3)

	println(("-------------"))
	x, y := swap(10, 20)
	println(x, y)

	z, zz := swap2(100, 200)
	println(z, zz)

	// makeはスライス、mapなどを生成するやつ。第二引数は要素数
	// Goにおいて、関数はファーストクラスオブジェクト。変数への代入、引数に渡す、戻り値で返すなどが可能
	// fs := make([]func() string, 2)
	// println(len(fs)) // 要素数
	// println(cap(fs)) // 容量

	fs := make([]func(), 3)
	// iは自由変数?
	for i := range fs {
		fs[i] = func() { fmt.Println(i) }
	}
	for _, f := range fs { f() }

	println(("-------------"))

	// 値のコピー
	type User struct {age int; name string}
	person := User{age:10,name: "Gopher"}
	person2 := person
	person2.age = 30
	fmt.Println(person)
	fmt.Println(person2)

	ns := []int{10, 20, 30}
	ns2 := ns // これはコピーではない。ポインタということ?
	ns[1] = 200
	println(ns[0], ns[1], ns[2])
	println(ns2[0], ns2[1], ns2[2])

	println(("-------------"))
	np, mp := 10, 20
	println(np, mp)
	swap3(&np, &mp)
	println(np, mp)

	println(("-------------"))
	npp, mpp := 100, 200
	println(&npp, &mpp)
	npp, mpp = swap(npp, mpp)
	println(&npp, &mpp)
	println(npp, mpp)

	println(("-------------"))
	var v T
	// 以下2つは同じ意味？
	(&v).f()
	v.f()

	println(("-------------"))
	// *T型はTのメソッドも自身のメソッドとして扱われる（意味不明）
	// (T{}).f()   // T
	// (&T{}).f()  // *T
	// (*&T{}).f() // T

	// (T{}).g() // <- できない
	// (&T{}).g()
	// (*&T{}).g()

	println(("-------------"))
	var n MyInt
	println(n)
	n.Inc()
	println(n)




}

func add(x int, y int) int {
	return x + y
}

// これって返り血の実態なんなの？変数が2つ？それらにどうやってアクセスするのか
func swap(x, y int) (int, int) {
	return y, x
}

// 名前付き戻り値。関数内でreturnする値を明示していないので、戻り値用変数であるx2, y2が返却される
func swap2(x, y int) (x2, y2 int) {
	x2, y2 = x, y
	return
}

type T int
func (t *T) fa() { println("hi") } // これはポインタレシーバ??
func (t T) f()  {}
func (t *T) g() {}

type MyInt int
func (n *MyInt) Inc() { *n++ }

// 単純にx, yを逆にしてreturnするだけ（swapのような感じ）だと順番を入れ替えて返しているだけで、ポインタの入れ替えにはなってない
// a はポインタ → メモリアドレスを保持
// *a はポインタ a が指す先の「実際の int の値」
func swap3(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}