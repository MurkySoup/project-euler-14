package main

// project euler 14
// see: https://projecteuler.net/problem=14

import (
        "fmt"
        "runtime"
        "time"
)

// accept an int 'n', return length of collatz sequence for that integer
func collatz(number int) int {
        count := 1

        for {
                if number > 1 {
                        if number & 1 == 0 {
                                // number = number / 2
                                number /= 2
                        } else {
                                number = (3 * number) + 1
                        }

                } else {
                        break // collatz sequence end at '1'
                }
                count++
        }

        return count
}

func main() {
        limit := 1000000

        runtime.GOMAXPROCS(4)
        start := time.Now()

        hi_val := 0
        hi_seq := 0

        for number := 2; number <= limit; number++ {
                go func(number int) {
                        seq_length := collatz(number)
                        if seq_length > hi_seq {
                                hi_seq = seq_length
                                hi_val = number
                        }
                }(number)
        }

        elapsed := time.Since(start)

        fmt.Printf("\nnumber = %d, seq len = %d\n", hi_val, hi_seq)
        fmt.Printf("\nExecution time: %s\n\n", elapsed)
}

// end of script
