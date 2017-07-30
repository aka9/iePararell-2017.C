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

	// 素数を求める範囲の最大値を引数として受け取る
	max, _ := strconv.Atoi(os.Args[1])

    // 等分を受け取る
    divided, _ := strconv.Atoi(os.Args[2])
    task := max / divided

    // 最小数
    min := 0

    // goroutineをグループで管理する変数
    var wg sync.WaitGroup

    // 分割
    split := []int{min}

    // 割り算
    quotient := max / task
    remainder := max % task
    if remainder != 0 {
		quotient = quotient + 1
    }

    //割り当て
    for i := 1; i <= quotient-1; i++ {
        split = append(split , i*task)
    }

    //要素の最後にmaxを加える
    split = append(split , max)

    //表示
    //fmt.Println(split)

    // 実行環境のcpu数を取得し, 表示
    maxcpu := runtime.NumCPU()
    //fmt.Printf("maxcpu = %d\n", maxcpu)

    // cpuの使うコア数を指定
    runtime.GOMAXPROCS(4)

    // cpu数に応じて作業を割り当てるための変数
    semaphore := make(chan int, maxcpu)

	// 計測開始
	start := time.Now()


    for i := 1; i <= quotient; i++ {
        wg.Add(1)
        go func(i int) {
			defer wg.Done()
			semaphore <- 1 // semaphoreを1加える
			Worker(split[i-1] , split[i])	 // Workerプロセスの代わりの関数
			<- semaphore   // semaphoreを出す.
        }(i)
    }

    wg.Wait() // グループ数が0になくなるまで待つ.
	end := time.Now() // 計測終了
	// 実行時間を出力
	//	log.Printf("max:%d %fs", max, (end.Sub(start)).Seconds())
	log.Printf("divided:%d %fs", divided, (end.Sub(start)).Seconds())

}

// Workerプロセスの代わりの関数.
func Worker(i int,j int){
	// iが素数かどうか求める. 素数の場合, 標準出力する.
	Prime(i+1, j)
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
