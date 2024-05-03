/*
	Прежде чем вы приступите к решению задач, используйте эту, чтобы познакомиться с платформой Techpoint Techpoint, а также получить первые 5 баллов.

Напишите программу, которая выводит сумму двух целых чисел.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var num int
	fmt.Fscan(in, &num)

	for i := 0; i < num; i++ {
		var a, b int
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
