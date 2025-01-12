package main

import "fmt"

func main() {
	/*
	* NOTE: 
	* := という書き方は、関数の中で利用できる。
	* 変数定義の省略形。
	* 省略しない場合は、var helloHandler = ...
	*/
	// handler: HTTPリクエストを受け取って、それに対するHTTPレスポンスの内容をコネクションに書き込む関数のこと
	// helloHandler := func(w http.ResponseWriter, req *http.Request) {
	// 	io.WriteString(w, "Hello, world!\n")
	// 	// io.WriteString(w, count)

	// 	/* fmt.Println()とprintlnの違いは？ */
	// 	fmt.Println()
	// 	println()
	// }
	// http.HandleFunc("/", helloHandler)
	// log.Println("server start at port 8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// fmt.Printf("v is of type %T\n", "")

	// 現在時刻を数値で取得する
	// t := time.Now().UnixNano()
	// println(t)

	// 乱数のたねを設定。
	// NOTE: rand.Seedの利用は非推奨となっている（1.20以降）
	// rand.Seed(t)

	// xは0-10の間の値になる
	// s := rand.Intn(10)
	// println(s)

	// 配列
	// ns := [3]int{1,29,39}
	// ns := [...]int{1,29,39}
	// println(ns[2])

	// println(("-------------"))


	// var ns1 []int
	// ns1 = make([]int, 3, 10)
	// ns1[0] = 2
	// println(ns1[1])
	// fmt.Printf("ns is of type %T\n", ns)
	// fmt.Printf("ns1 is of type %T\n", ns1)

	// println(("-------------"))
	//  ns2 := []int{10, 20, 30, 40, 50}
	// println(ns2[0])

	// println(("-------------"))
	// ns3 := []int{5: 50, 10: 100}

	// println(ns3[0])
	// println(ns3[5])
	// println(ns3[10])




	// NOTE: 単純に配列の要素数を指定していないものをスライスと呼んでいる？
	// 配列の一部を切り出したデータ構造という表現気になる
	// println(("-------------"))
	//  var slice  [1]int
	// array := [5]int{1, 2, 3, 4, 5}
	// println(slice)
	// fmt.Printf("v is of type %T\n", slice)

	// println("-------------")
	// println("スライスの操作")
	// println("-------------")
	// lenは格納されてる要素数？
	// capは格納できる要素数？
	
	// ns4 := []int{10, 20, 30, 40, 50}
	// println(len(ns4))
	// println(cap(ns4))

	// ns4の容量が足りてないので、ns4の容量は2倍になる。たぶん別のメモリ領域に値が格納されてる
	// ns4= append(ns4, 60, 70)
	// println(len(ns4))
	// println(cap(ns4))

	// println(("-------------"))
	// println(("-------------"))
	// a := []int{10, 20}
	// a := []int{10, 20, 5: 10}
	// println(len(a))
	// println(cap(a))

	// b := append(a, 30)
	// println(len(b))
	// println(cap(b))

	// c := append(b, 40)
	// println(len(c))
	// println(cap(c))

	// for i, v := range []string{"foo", "bar", "baz"} {	
  //   println(i, v )
	// }
	// println(("-------------"))
	// println(("-------------"))
	// d := []int{10, 20, 30, 40, 50}
	// dn := d[1:5]
	// println(dn[0])
	// println(len(dn))
	// println(cap(dn))
	
	// 1番目（実質2番目）の要素が削除される
	// dna := append(d[:1], d[2:]...)
	// for _, v:= range dna {
	// 	println("dna", v)
	// }

	// NOTE: チャレンジ
	// n1 := 19
	// n2 := 86
	// n3 := 1
	// n4 := 12
	nums := []int{19, 86, 1, 12}
	fmt.Println(nums)


	// sum := 0
	// for _, v := range nums {
	// 	sum += v
	// }

	// sum := n1 + n2 + n3 + n4
	// println(sum)

	// println(("-------------"))
	// println(("-------------"))
	// var m map[string]int
	// m = make(map[string]int)

	// // 容量の指定
	// m = make(map[string]int, 10)

	// リテラルでの初期化
	// m := map[string]int{"x": 10, "y": 20}
	// println(m["x"])

	// m["z"] = 30
	// comma ok イディオム？
	// n, ok := m["z"]
	// println(n, ok)
	// delete(m, "z")
	// n2, ok := m["z"]
	// println(n2, ok)
	// println(len(m))

	// for i, vv := range m {
	// 	println(i, vv)
	// }
	// x/exp/mapsパッケージ
	// maps.Clear

	functionAndTyep()
}