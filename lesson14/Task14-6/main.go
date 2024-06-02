// простой ответ: main + две дочерние = 3. но дочерние несколько раз обмениваются ch

// сложный ответ:
// 1. main - нормально завершилась
// 2. первая горутина. ждет пока вторая горутина отдаст ch. столько раз сколько успеет за 5 секунд
//	(наверное 4 раза, так как задержка во второй горутине 1 секунда)
// 3. main кладет дважды значения в stop
// 4. когда первая горутина получила stop и прошла ещё раз в OUT она получает наконец default и завершается
// 5. когда вторая горутина получила stop и далее прошла OUT она получает default и делает задержку в 1 секунду
// 6. в main так же запускается задержка в 1 секунду,
//		но так как она запустилась сразу после отправки stop - main успевает завершится,
//		вторая горутина не успевает завершиться и убивается вместе с main

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	stop := make(chan struct{}, 2)
	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)

			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()
	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()
	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}
