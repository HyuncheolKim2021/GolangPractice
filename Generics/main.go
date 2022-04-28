package main

import (
	"fmt"
)

// 1. ポインタを返す関数
// 引数で渡された値を持つ変数のポインタ
// 仮引数のポインタを返す
func Ptr[output any](arg output) *output {
	return &arg
}

// 2. スライスに関数を適用する
// スライスの各要素に関数を適用
// 適用する関数を引数として渡すことができる
func Apply[T any](array []T, fn func(int, T)) {
	for i, v := range array {
		fn(i, v)
	}
}

// 3. スライスのフィルター
// スライスの各要素に引数で指定した関数を適用
// 関数がtrueを返す要素だけで構成されたスライスを返す
// 並び順は元のスライスと同じ
func Filter[T any](array []T, fn func(int, T) bool) []T {
	var newArray []T
	for i, v := range array {
		if fn(i, v) {
			newArray = append(newArray, v)
		}
	}
	return newArray
}

func main() {
	ptr := Ptr[bool](false)
	fmt.Println(ptr) // 実行例) 0xc00001c072

	var sum int
	Apply[int]([]int{10, 20}, func(i, v int) {
		sum += v
	})
	fmt.Println(sum) // 30

	ns := []int{1, 2, 3, 4}
	ms := Filter[int](ns, func(i, v int) bool {
		return v%2 == 0
	})
	fmt.Println(ms) // [2 4]
}
