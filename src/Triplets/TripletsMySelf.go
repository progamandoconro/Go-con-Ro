package main
import "fmt"

var a1, a2, a3, b1, b2, b3 int
var ra1, ra2, ra3, rb1, rb2, rb3 int

func main() {

	fmt.Scanln(&a1, &a2, &a3)
	fmt.Scanln(&b1, &b2, &b3)

        // Round 1
        if a1 == b1 {
                ra1 , rb1 = 0 , 0
        }

	if a1 > b1 {
                ra1 , rb1 = 1 , 0
        }

	if a1 < b1 {
                ra1 , rb1 = 0,  1
        }
	//Round 2
        if a2 == b2 {
                ra2, rb2 = 0 , 0
        }

	if a2 > b2 {
                ra2 ,rb2 = 1 ,  0
        }

	if a2 < b2 {
                ra2 , rb2 =  0 , 1
        }
        //Round 3
        if a3 == b3 {

                ra3 , rb3 = 0 , 0
        }

	if a3 > b3 {
                ra3 , rb3 = 1 , 0
        }

        if a3 < b3 {
                ra3 , rb3 = 0 ,  1
        }

        fmt.Println(ra1 + ra2 + ra3 , rb1 + rb2 + rb3)
}
