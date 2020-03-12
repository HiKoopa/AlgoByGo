package main

import "fmt"
import "io"

/*
//1
func main() {
	var num1 int
	fmt.Scan(&num1)
	for i := 0; i < 9; i++ {
		fmt.Printf("%d * %d = %d\n", num1, i+1, num1 * (i+1))
	}
}
*/
/*
//2
func main() {
	var n, a, b int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(a + b)
	}
}
*/
/*
//4
//FAIL
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Println(n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		writer.WriteString(strconv.Itoa(a+b))
		writer.WriteString("\n")
		writer.Flush()
	}
}
*/
/*
//7
func main() {
	var n, a, b int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Printf("Case #%d: %d\n", i+1, a+b)
	}
}
*/
/*
//8
func main() {
	var n, a, b int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d", &a, &b)
		fmt.Printf("Case #%d: %d + %d = %d\n", i+1, a, b, a+b)
	}
}
*/
//11
/*
func main() {
	var n1, n2, a int
	fmt.Scanf("%d %d", &n1,& n2)
	for i := 0; i < n1; i++ {
		fmt.Scanf("%d", &a)
		if a < n2 {
			fmt.Printf("%d ", a)
		}
	}
}
*/
//while
/*
//1
func main() {
	var n1, n2 int
	for {
		fmt.Scanf("%d %d", &n1, &n2)
		if n1 == 0 && n2 ==0 {
			return
		} else {
			fmt.Println(n1 + n2)
		}
	}
}
*/
//2
func main() {
	var n1, n2 int
	for {
		if _, err := fmt.Scanf("%d %d", &n1, &n2); err != io.EOF {
			fmt.Println(n1 + n2)
		} else {
			return
		}
	}
}