package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	read := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("------------------------------------------------")
		fmt.Println(titlewords("Hello, Welcome to Choice Making Zone! ", "title"))
		fmt.Println("_____________________________________________")

		fmt.Println(titlewords("Please Enter Your Name: ", "head"))
		name, _ := read.ReadString('\n')
		fmt.Println(titlewords(name, "head"))

		fmt.Println(titlewords("Please Select Option:\n 1. Convert Hex to Decimal.\n 2. convert binery to decimal.\n 3. convert octal to decimal.\n 4. digit calculator.\n 5. helps.\n 6. contact us\n 7. quit.", "head"))
		choose, _ := read.ReadString('\n')
		choose = strings.TrimSpace(choose)

		//			hexadecimal convertion
		if choose == "1" {
			fmt.Println(titlewords("Enter hexadecimal number to convert: ", "head"))
			h, _ := read.ReadString('\n')
			h = strings.TrimSpace(h)
			fmt.Println(convert(h, choose))
			break

		}
		// binary convertion
		if choose == "2" {
			fmt.Println(titlewords("enter binery number to convert: ", "head"))
			b, _ := read.ReadString('\n')
			b = strings.TrimSpace(b)
			fmt.Println(convert(b, choose))
			break
		}
		// octal convertion
		if choose == "3" {
			fmt.Println(titlewords("please enter octal number to convert: ", "head"))
			c, _ := read.ReadString('\n')
			c = strings.TrimSpace(c)
			fmt.Println(convert(c, choose))
			break
		}
		// digit calculator
		if choose == "4" {

			var num1, num2 int
			fmt.Println(titlewords("please enter your first number: ", "head"))
			fmt.Scanln(&num1)
			// signs
			var sign string
			fmt.Println(titlewords("Enter your signs.\n 1. add.\n 2. minus.\n 3. multiply.\n 4. divide.\n 5. return.", "head"))
			fmt.Scanln(&sign)
			// num2
			fmt.Println(titlewords("Enter you Last number: ", "head"))
			fmt.Scanln(&num2)
			fmt.Println(digit(num1, sign, num2))
			break

		}
	}
}
