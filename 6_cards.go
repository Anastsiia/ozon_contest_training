package main

import (
	"bufio"
	"fmt"
	"os"
)

func findSolution(collection []int, cards int, minCard int) []int {
	res := make([]int, len(collection))
	check := len(collection)

	for i := minCard; i <= cards; i++ {
		for a, v := range collection {
			if res[a] != 0 {
				continue
			} else if v < i && res[a] == 0 {
				res[a] = i
				check--
				if check == 0 {
					return res
				}
				break
			}
		}
	}
	return []int{-1}
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	if n >= m {
		fmt.Fprintln(out, -1)
	} else {
		collection := make([]int, n)
		minCard := m
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &collection[i])
			if collection[i] < minCard {
				minCard = collection[i]
			}
		}
		if minCard > m-n {
			fmt.Fprintln(out, -1)
		} else {
			res := findSolution(collection, m, minCard)
			for _, v := range res {
				fmt.Fprintf(out, "%d ", v)
			}
			fmt.Fprintln(out)
		}

	}
}
