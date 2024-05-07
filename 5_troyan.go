package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var x int
	fmt.Fscan(in, &x)

	for i := 0; i < x; i++ {
		var n int
		fmt.Fscan(in, &n)

		copy(sortedArr, startArr)
		sort.Ints(sortedArr)

		for _, v := range startArr {
			index := findElement(sortedArr, v)
			for index > 0 && sortedArr[index]-sortedArr[index-1] <= 1 {
				index--
			}
			fmt.Fprintf(out, "%d ", index+1)
		}
		fmt.Fprintln(out)
	}
}
