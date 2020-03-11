package main

import "fmt"
/*
//1
func main() {
	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)
	if num1 > num2 {
		fmt.Println(">")
	} else if num1 < num2 {
		fmt.Println("<")
	} else {
		fmt.Println("==")
	}
}
*/

/*
//2
func main() {
	var num1 int
	fmt.Scanf("%d", &num1)
	switch {
	case num1 >= 90 && num1 <= 100:
		fmt.Println("A")
	case num1 >= 80 && num1 <= 89:
		fmt.Println("B")
	case num1 >= 70 && num1 <= 79:
		fmt.Println("C")
	case num1 >= 60 && num1 <= 69:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
*/
/*
//3
func main() {
	var num1 int
	fmt.Scanf("%d", &num1)
	switch {
	case num1 % 400 == 0:
		fmt.Println("1")
	case num1 % 4 == 0 && num1 % 100 != 0:
		fmt.Println("1")
	default:
		fmt.Println("0")
	}
}
*/
/*
//4
func main() {
	var num1, num2, num3, num4, num5, min1, min2 int
	fmt.Scanf("%d", &num1)
	fmt.Scanf("%d", &num2)
	fmt.Scanf("%d", &num3)
	fmt.Scanf("%d", &num4)
	fmt.Scanf("%d", &num5)

	if num1 < num2 {
		if num1 < num3 {
			min1 = num1
		} else {
			min1 = num3
		}
	} else {
		if num2 < num3 {
			min1 = num2
		} else {
			min1 = num3
		}
	}

	if num4 < num5 {
		min2 = num4
	} else {
		min2 = num5
	}

	fmt.Println(min1 + min2 - 50)
}
*/
//5
func main() {
	var num1, num2, num3, mid int
	fmt.Scanf("%d %d %d", &num1, &num2, &num3)

	if num1 < num2 {
		if num1 < num3 {
			if num2 < num3 {
				mid = num2
			} else {
				mid = num3
			}
		} else {
			mid = num1
		}
	} else {
		if num2 < num3 {
			if num1 < num3 {
				mid = num1
			} else {
				mid = num3
			}
		} else {
			mid = num2
		}
	}

	fmt.Println(mid)
}