package main

import (
    "fmt"
    "time"
)

func main() {
    max := 10000
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
    goal := time.Now() //Goal
    fmt.Printf("\n%v 経過", goal.Sub(start)) //経過時間を表示
}
