// Поскольку у нас исторически был создан один репозиторий для всего проекта.
// Для модуля создан отдельный репозиторий,
// но в головном репозитории папка с текущей версией модуля так же будет присутствовать
// однако вызывать мы будем как модуль через зависимости в go.mod

package main

import (
	"fmt"
)

func main() {

	fmt.Println()

}
