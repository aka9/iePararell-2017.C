package main

import "fmt"     // fmtはデータを整形して出力する
import "time"    // 実行時間計測のため (メソッドNow，Since)
import "log"     // 実行時間の出力先変更 (メソッドPrintf)
import "strconv" // 文字列の数値変換 （メソッドAtoi）
import "os"      // 引数の受け取り (配列Args)

// 調査の対象となる数字をnとすると1~nの数字で割って割り切れたらdivisorは+1され、
// 「divisorが2のとき」なのは、1を取り除くため
func main() {
	// 素数を求める範囲の最大値を引数として受け取る
	max, _ := strconv.Atoi(os.Args[1])

	// 計測開始
	start := time.Now()
	for i := 1; i <= max; i++ {
		var divisor int = 0
		for j := 1; j <= i; j++ {
			if i % j == 0 {
				divisor += 1
			}
		}
		if divisor == 2{
			fmt.Println(i)
		}
    }
	end := time.Now() // 計測終了
	// 実行時間を出力
	log.Printf("%fs", (end.Sub(start)).Seconds())
}
