package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {
		fmt.Println("_____________________________________________")
		fmt.Println(titlewords("Hello, Welcome to Choice Making Zone! ", "title"))
		fmt.Println("_____________________________________________")
		read := bufio.NewReader(os.Stdin)
		fmt.Println(titlewords("Please Enter Your Name: ", "head"))
		name, _ := read.ReadString('\n')
		fmt.Println(titlewords(name, "head"))

		fmt.Println(titlewords("Please Select Option:\n 1. Convert Hex to Decimal.\n 2. convert binnery to decimal.\n 3. digit calculator.\n 4. help.\n 5. contact us.\n 6. quite", "head"))
		choose, _ := read.ReadString('\n')
		choose = strings.TrimSpace(choose)

		if choose == "1" {
			fmt.Println()

		}

	}
}
