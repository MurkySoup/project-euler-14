package main

import (
        "fmt"
)

func CollatzSequence(n int) int {
        result := 0

        for n != 1 {
                result ++
                if n % 2 == 0 {
                        n /= 2
                } else {
                        n = (3 * n) + 1
                }
        }

        return result
}

func main() {
        number := int(1)
        max := 1

        for n := int(1); n < 1000000; n++ {
                length := CollatzSequence(n)
                if length > max {
                        number = n
                        max = length
                }
        }

        fmt.Printf("\nNumber : %d\n", number)
        fmt.Printf("Collatz: %d\n\n", max + 1)
}
