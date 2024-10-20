package main

import (
	"fmt"
	"math/rand"
	"time"
)

//можно использовать https://go.dev/play/p/w0s82mkTK_s

func main() {
	td := testData{}
	generate(100, &td)
	time.Sleep(time.Millisecond * 200)

	fmt.Println(len(td.phones))
}

type testData struct {
	phones []int
}

func (td *testData) add() {
	time.Sleep(time.Millisecond * 200)
	td.phones = append(td.phones, randPhone())
}

func generate(n int, td *testData) *testData {
	for i := 0; i < n; i++ {
		go td.add()
	}
	return td
}

func randPhone() int {
	return 89000000000 + rand.Intn(999999999-100000000) + 100000000
}
