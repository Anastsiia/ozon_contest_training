/*
Вам поручено разработать систему оповещений для мессенджера. Мессенджером пользуются пользователей; им присвоены идентификационные номера от. Оповещения, которые должны приходить пользователям, могут быть двух типов:
глобальное оповещение, приходящее всем пользователям (например, при обновлении версии мессенджера);
персональное оповещение, приходящее одному пользователю (например, когда ему приходит сообщение в диалоге).
Оповещения получают номера, начиная с, в порядке их отправления. То есть первое отправленное оповещение (неважно, глобальное или персональное) получает номер, второе — номер, и так далее.
Когда пользователь загружает мессенджер, на экране загрузки сразу же должно отображаться последнее из оповещений, пришедших ему. Вам необходимо реализовать программу, которая будет выводить эту информацию.
Формально, ваша программа должна обрабатывать последовательность запросов двух типов:
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var globalNotifID int

func sendNotif(idUser int, notifID int, notifMap []int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	mu.Lock()
	defer mu.Unlock()
	if notifID > notifMap[idUser] {
		notifMap[idUser] = notifID
	}

}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	notifID := 0
	var notifMap = make([]int, n+1)

	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < q; i++ {
		var t, id int
		fmt.Fscan(in, &t, &id)

		if t == 1 {
			notifID++
			wg.Add(1)
			go sendNotif(id, notifID, notifMap, &wg, &mu)
		} else {
			wg.Wait()
			mu.Lock()
			if notifMap[0] > notifMap[id] {
				notifMap[id] = notifMap[0]
			}
			fmt.Fprintln(out, notifMap[id])
			mu.Unlock()
		}
	}
}
