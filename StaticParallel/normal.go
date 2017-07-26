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
    for n := 2; n <= max; n++ {
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
	end := time.Now() // 計測終了
	// 実行時間を出力
	log.Printf("max:%d %fs", max, (end.Sub(start)).Seconds())
}
