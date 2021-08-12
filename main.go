package main

import (
	"fmt"
	"time"
)

//func main() {
//	cmd.Execute()
//}

type test struct {
	t time.Time
}

func (t *test) set1() {
	go func() {
		for {
			t.t = time.Now().Add(time.Hour)
			time.Sleep(time.Millisecond * 600)
		}
	}()
}
func (t *test) set2() {
	go func() {
		for {
			t.t = time.Now()
			time.Sleep(time.Second * 1)
		}
	}()
}

func (t *test) now() {
	t.t = time.Now()
}

func (t *test) print() {
	fmt.Println(t.t)
}
func main() {
	t1 := test{t: time.Now()}
	t1.set1()
	t1.set2()
	for {
		fmt.Println(t1.t)
		time.Sleep(time.Millisecond * 300)
	}
}
