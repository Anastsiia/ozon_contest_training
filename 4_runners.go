package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func findElement(arr []int, elem int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		midl := (right-left)/2 + left
		if arr[midl] == elem {
			return midl
		} else if arr[midl] < elem {
			left = midl + 1
		} else {
			right = midl - 1
		}
	}
	return -1
}

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
		startArr := make([]int, n)
		sortedArr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &startArr[i])
		}
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
