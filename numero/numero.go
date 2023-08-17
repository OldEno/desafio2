package main

import "fmt"


func main() {

	for N := 1; N <= 100; N++ {
		if N%3 ==0 {
			fmt.Println(N, "Pin")
		}

		if N%3 != 0 && N%5 != 0 {
			fmt.Println(N)
		}

		if N%5 == 0 {
		fmt.Println(N, "Pan")
		}

		if N%3 == 0 && N%5 == 0 {
			fmt.Println(N, "Pin, Pan")
		}
			
	}

}