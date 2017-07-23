package main

import (
	"fmt"     // fmtはデータを整形して出力する
	"log"     // 実行時間の出力先変更 (メソッドPrintf)
	"os"      // 引数の受け取り (配列Args)
	"strconv" // 文字列の数値変換 （メソッドAtoi）
	"time"    // 実行時間計測のため (メソッドNow，Since)
)

// 調査の対象となる数字をnとすると1~nの数字で割って割り切れたらdivisorは+1され、
// 「divisorが2のとき」なのは、1を取り除くため
func main() {
	// 素数を求める範囲の最大値を引数として受け取る
	max, _ := strconv.Atoi(os.Args[1])

	// 計測開始
	start := time.Now()
	Prime(10, max)
	end := time.Since(start) // 計測終了
	log.Printf("%s", end)    // 実行時間を出力

}

// Prime 素数を求める関数.
// @start この値から素数を求めていく.
// @max この値まで素数を求めていく.
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
