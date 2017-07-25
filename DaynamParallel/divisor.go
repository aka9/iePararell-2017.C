package main

import (
	"fmt"     // fmtはデータを整形して出力する
	"log"     // 実行時間の出力先変更 (メソッドPrintf)
	"os"      // 引数の受け取り (配列Args)
	"strconv" // 文字列の数値変換 （メソッドAtoi）
	"time"    // 実行時間計測のため (メソッドNow，Since)
	"runtime" // 実行環境のcpu数を表示する. (メソッド NumCPU)
	"sync" // 排他的処理のために必要.
)

// 調査の対象となる数字をnとすると1~nの数字で割って割り切れたらdivisorは+1され、
// 「divisorが2のとき」なのは、1を取り除くため
func main() {

	// goroutineをグループで管理する変数.
	var wg sync.WaitGroup

	// 素数を求める範囲の最大値を引数として受け取る
	max, _ := strconv.Atoi(os.Args[1])
	min := 0

	// 実行環境のcpu数を取得し, 表示
	maxcpu := runtime.NumCPU()
	fmt.Printf("maxcpu = %d\n", maxcpu)

	// cpuの使うコア数を指定
	runtime.GOMAXPROCS(4)

	// cpu数に応じて作業を割り当てるための変数
	semaphore := make(chan int, maxcpu)

	// 計測開始
	start := time.Now()

	// 動的割当の場合, 計算量が多い順から割り当てた方が計算量が良い.
	// そのため, 最大値からデクリメントしながら関数を使用
	for i := max; i > min; i-- {
		// 待機グループ数を1つ増やす.
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // この関数を終えた際, グループ数を1つ減らす
			semaphore <- 1 // semaphoreを1加える
			Worker(i)			  // Workerプロセスの代わりの関数
			<- semaphore   // semaphoreを出す.
		}(i)
	}

	wg.Wait() // グループ数が0になくなるまで待つ.
	end := time.Now() // 計測終了
	// 実行時間を出力
	log.Printf("%fs", (end.Sub(start)).Seconds())

}

//func MiniMaster(){}

// Workerプロセスの代わりの関数.
func Worker(i int){
	// iが素数かどうか求める. 素数の場合, 標準出力する.
	Prime(i, i)
}

// Prime 素数を求める関数.
// @min この値から素数を求めていく.
// @max この値まで素数を求めていく.
func Prime(min int, max int) int {
	var divisor int
	for i := min; i <= max; i++ {
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
