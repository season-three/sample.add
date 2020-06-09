package main

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	//testsにtest1, 2, 3を配列する
	tests := []struct {
		//構造化
		Result, Expected int
	}{
		//test1
		{Add(3, 5), 8},

		//test2
		{Add(2, 4), 6},

		//test3
		{Add(1, 3), 4},
	}

	//testsをループ処理する
	for _, t := range tests {

		if t.Result != t.Expected {
			fmt.Printf("failed test:result was %d, but expected %d", t.Result, t.Expected)
		}
	}

}
