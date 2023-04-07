package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

var in *bufio.Reader
var out *bufio.Writer
var scan = fmt.Fscan
var println = func(iface interface{}) {
	fmt.Fprintln(io.Writer(out), iface)
}

func solve() string {

	var n, s1, s2 int
	scan(in, &n, &s1, &s2)

	type pair struct {
		id  int
		sec int
	}
	a := make([]pair, n)
	var sec int
	for i := range a {
		scan(in, &sec)
		a[i].id = i + 1
		a[i].sec = sec
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].sec > a[j].sec
	})

	// println(a)

	i, j := 1, 1
	b, c := []int{}, []int{}
	cost := 0
	for _, p := range a {
		if i*s1*p.sec <= j*s2*p.sec {
			b = append(b, p.id)
			cost += i * s1 * p.sec
			i++
		} else {
			c = append(c, p.id)
			cost += j * s2 * p.sec
			j++
		}
	}
	// println(cost)

	fmt.Fprintf(out, "%d ", len(b))
	for _, x := range b {
		fmt.Fprintf(out, "%d ", x)
	}
	println("")

	fmt.Fprintf(out, "%d ", len(c))
	for _, x := range c {
		fmt.Fprintf(out, "%d ", x)
	}
	println("")
	return ""

}

func min(a ...int) int {
	res := a[0]
	for _, x := range a[1:] {
		if x < res {
			res = x
		}
	}
	return res
}
func max(a ...int) int {
	res := a[0]
	for _, x := range a[1:] {
		if x > res {
			res = x
		}
	}
	return res
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

	var _r io.Reader
	var err error
	if _r, err = os.Open("./test.txt"); err != nil {
		_r = os.Stdin
	}
	var _w io.Writer = os.Stdout
	in = bufio.NewReader(_r)
	out = bufio.NewWriter(_w)
	defer func() {
		// 退出函数前清空缓存
		out.Flush()
	}()

	var t int
	scan(in, &t)
	for ; t > 0; t-- {
		// println(solve())
		solve()
	}

}
