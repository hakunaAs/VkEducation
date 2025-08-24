#### Прохождение курса по алгоритмам от VK Education
https://education.vk.company/

Выполняются мной используя язык Go

Корректность задания проверяется инструментами платформы _AllCups_
https://cups.online/ru

У данной платформы есть недостаток: 

не интуитивно понятный способ обработки ввода данных

_Примерно в таком виде воспринимает платформа All Cups ввод данных:_

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n int
	_, _ = fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Fscan(in, &arr[i])
	}
}
```
