package main

import "fmt"

var a1, a2, a3, b1, b2, b3 int

func main() {
	var a1, a2, a3, b1, b2, b3 int
        fmt.Scanln(&a1)
        fmt.Scanln(&a2)
        fmt.Scanln(&a3)
        fmt.Scanln(&b1)
        fmt.Scanln(&b2)
        fmt.Scanln(&b3)
        if a1 == b1 {
                a1, b1 = 0, 0
        }
	if a2 == b2 {
                a2, b2 = 0, 0
        }
	if a3 == b3 {

                a3, b3 = 0, 0
        }

	if a1 > b1 {
                a1, b1 = 1, 0
        }

	if a2 > b2 {
                a2, b2 = 1, 0
        }

	if a3 > b3 {
                a3, b3 = 1, 0
        }

	if a1 < b1 {
                a1, b1 = 0, 1
        }

	if a2 > b2 {
                a2, b2 = 0, 1
        }

	if a3 > b3 {
                a3, b3 = 0, 1
        }

	fmt.Println((a1 + a2 + a3), (b1 + b2 + b3))


// not working yet
