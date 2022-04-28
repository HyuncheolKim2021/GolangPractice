package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Q1. マップ関数
func Map[T, U constraints.Ordered](array []T, fn func(T) U) []U {
	newArray := make([]U, len(array))
	for i, v := range array {
		newArray[i] = fn(v)
	}
	return newArray
}

// Q2. Tuple型
type Tuple[T1, T2 constraints.Ordered] struct {
	X T1
	Y T2
}

func New[T1, T2 constraints.Ordered](x T1, y T2) *Tuple[T1, T2] {
	return &Tuple[T1, T2]{
		X: x,
		Y: y,
	}
}

// Q3. 任意の数値のうち大きな値を取得する関数について考察せよ
func Max[N constraints.Ordered](n1 N, n2 N) N {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func Try() {
	var ss []string = Map([]int{10, 20}, func(n int) string {
		return fmt.Sprintf("0x%x", n)
	})
	fmt.Println(ss)

	var t *Tuple[int, string] = New(10, "hoge")
	fmt.Println(t.X, t.Y)

	fmt.Println(Max(10, 20))
	fmt.Println(Max(3.5, 2.5))
	fmt.Println(Max(-2, -1))
	// 型はconstraints.Orderedで一緒だけどこういうのはエラーになる
	// fmt.Println(Max(3.5, 2))
	// fmt.Println(Max(-2, -1.5))
}
