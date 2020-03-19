package main
import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	r.Scan()
	n, _:= strconv.Atoi(r.Text())
	var c int
	for i := 0; i < n; i++ {
		if isHan(i+1) == true {
			c++
		}
	}
	fmt.Fprintf(w, "%d\n", c)
	w.Flush()
}

func isHan(n int) bool {
	if n < 10 {
		return true
	}
	d := n % 10 - (n / 10) % 10
	for i := 0; i < 3; i++ {
		if n < 10 {
			break
		}
		n1 := n % 10
		n2 := (n / 10) % 10
		if d != n1 - n2 {
			return false
		}
		
		n = n / 10
	}

	return true
}

/*
//1 - 15596
func sum(a []int) int {
	var r int
	for i := 0; i < len(a); i++ {
		r += a[i]
	}
	return r 
}
*/

