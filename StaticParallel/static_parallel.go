/*AA
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
	"log"     // 実行時間の出力先変更 (メソッドPrintf)
	"strconv" // 文字列の数値変換 （メソッドAtoi）
	"os"      // 引数の受け取り (配列Args)
)

func main() {
	// 素数を求める範囲の最大値を引数として受け取る
	max, _ := strconv.Atoi(os.Args[1])
    fmt.Printf("%v 以下の素数:", max)

    start := time.Now() //Start


	finished := make(chan bool)
	/************************************************/
	go func(){
        Prime(1, max/4)
		finished <- true
    }()//go func

	/************************************************/
    go func(){
        Prime((max/4)+1, max/2)
		finished <- true
    }()//go func
	/************************************************/
    go func(){
        Prime((max/2)+1, (max*3)/4)
		finished <- true
    }()//go func

	go func(){
        Prime(((max*3)/4)+1, max)
		finished <- true
    }()//go func


    <-finished
    <-finished
    <-finished
    <-finished

	end := time.Now() // 計測終了
	// 実行時間を出力
	log.Printf("max:%d %fs", max, (end.Sub(start)).Seconds())
}


func Prime(start int, max int) int {
	var divisor int
	for i := start; i <= max; i++ {
		divisor = 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				divisor++
			}
		}
		if divisor == 2 {
			fmt.Println(i)
		}
	}
	return 0
}