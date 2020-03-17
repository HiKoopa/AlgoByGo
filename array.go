package main
import (
	"bufio"
	"strconv"
	"os"
	"fmt"
)
/*
//1
func main() {
	max, min := -1000000, 1000000
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	r.Scan()
	n, _ := strconv.Atoi(r.Text())
	for i := 0; i < n; i++ {
		r.Scan()
		a, _:= strconv.Atoi(r.Text())
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}

	fmt.Fprintf(w, "%d %d\n", min, max)
	w.Flush()
}
*/
/*
//2
func main() {
	var max, index int
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	for i := 0; i < 9; i++ {
		r.Scan()
		a, _:= strconv.Atoi(r.Text())
		if a > max {
			max = a
			index = i
		}
	}

	fmt.Fprintf(w, "%d\n", max)
	fmt.Fprintf(w, "%d\n", index+1)
	w.Flush()
}
*/
/*
//4
func main() {
	mymap := make(map[int]int)
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	for i := 0; i < 10; i++ {
		r.Scan()
		a, _:= strconv.Atoi(r.Text())
		mymap[a % 42] = 1
	}

	fmt.Fprintf(w, "%d\n", len(mymap))
	w.Flush()
}
*/
//5
func main() {
	var sum, max int
	//mymap := make(map[int]int)
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	r.Split(bufio.ScanWords)
	r.Scan()
	len, _:= strconv.Atoi(r.Text())
	for i:=0; i < len; i++ {
		r.Scan()
		a, _:= strconv.Atoi(r.Text())
		if max < a {
			max = a
		}
		sum = sum + a
	}
	fmt.Fprintf(w, "%f\n", float32(sum) / float32(len) / float32(max) * 100.0)
	w.Flush()
}