/*
Для отслеживания посылок компания NOZO использует наклейки с надписями. Иногда надпись (или её часть) на наклейке нужно исправить, и тогда поверх старой наклейки лепят новую.
На очередной посылке появилось слишком много наклеек и теперь невозможно прочитать наклеенную надпись целиком.
Помогите это сделать по истории этих наклеек.
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

	var s string
	fmt.Fscan(in, &s)

	var n int
	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var start, end int
		var r string
		fmt.Fscan(in, &start, &end, &r)
		s = s[:start-1] + r + s[end:]

	}
	fmt.Fprintln(out, s)
}
