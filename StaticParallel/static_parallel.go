/*
goroutine で 平行処理ができる
goroutine は　Thread, Process でもない。
実行する関数にgoのキーワードを付与するだけでよい
http://blog.excite.co.jp/exdev/27147029/

10000くらいにしないと差がでませんでした！
*/


package main

import (
    "fmt"
    "time"
)

func main() {
    max := 10000
    fmt.Printf("%v 以下の素数:", max)

    start := time.Now() //Start


	finished := make(chan bool)
/************************************************/
	go func(){
    for n := 2; n <= 3333; n++ {
    	
        flag := true
        for m := 2; m < n; m++ {
            if (n % m) == 0 { // n が m で割り切れる → 素数ではない
                flag = false
                break
            }
        }
        if flag {
            fmt.Printf(" %v", n)
        }
    }
    finished <- true
    }()//go func

/************************************************/
    go func(){
    for n1 := 3334; n1 <= 6666; n1 ++ {
    	
        flag1 := true
        for m1 := 2; m1 < n1; m1++ {
            if (n1 % m1) == 0 { // n1 が m で割り切れる → 素数ではない
                flag1 = false
                break
            }
        }
        if flag1 {
            fmt.Printf(" %v", n1)
        }
    }
    finished <- true
    }()//go func
/************************************************/
    go func(){
    for n2 := 6667; n2 <= max; n2++ {
    	
        flag2 := true
        for m2 := 2; m2 < n2; m2++ {
            if (n2 % m2) == 0 { // n2 が m2 で割り切れる → 素数ではない
                flag2 = false
                break
            }
        }
        if flag2 {
            fmt.Printf(" %v", n2)
        }
    }
    finished <- true
    }()//go func

    <-finished
    <-finished
    <-finished

    goal := time.Now() //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
