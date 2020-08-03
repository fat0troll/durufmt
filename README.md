# durafmt_ru - это durafmt, но на русском

## Intro in English

[durafmt](https://github.com/hako/durafmt) is a tiny Go library that formats `time.Duration` strings (and types) into a human readable format. durafmt_ru is the same, but for Russian language.

All documentation and contributions are in Russian due to library specific use case. If you interested in English language human readable durations, consult the original library documentation. If you interested in any other language (or want to merge durafmt, durafmt_ru and maybe more into one fully internationalized library), feedback is welcome.

## Что это за библиотека?

[durafmt](https://github.com/hako/durafmt) - это маленькая Go-библиотека, которая позволяет отформатировать строки `time.Duration` (и типы тоже) в человекочитаемый формат на английском языке. durafmt_ru делает то же самое, но на русском языке.

Вся документация (в виде этого README), комментарии в коде и тесты переведены на русский язык. Заинтересованные в человекочитаемых временных интервалах на английском языке могут использовать исходную библиотеку durafmt. Заинтересованные в других языках, помимо английского и русского (а так же, возможно, в создании более универсальной библиотеки, которая будет поддерживать нормальную локализацию) могут отправлять свой фидбэк мне или автору оригинальной библиотеки.

## Установка

```
go get github.com/fat0troll/durafmt_ru
```

## Зачем нужна эта библиотека?

Автор оригинальной durafmt боролся с тем, что дефолтное строковое представление `time.Duration` выглядит так себе. Например, вот так:

```
53m28.587093086s // :)
```

Легко прочитать? А как насчёт вот такого?

```
354h22m3.24s // :S
```

Но есть ещё одна проблема, касающаяся конкретно durafmt_ru - числительные в русском языке работают несколько сложнее, чем в английском. Недостаточно просто в конце поставить одну буковку, чтобы всё заработало - нужно учитывать три склонения (а ещё помнить, что одиннадцать, двенадцать и тринадцать - волшебные числительные-исключения). Из-за невнимательности и/или нежелания заморачиваться иногда мы видим в локализованных на русский язык интерфейсах что-нибудь типа "осталось 4 минут(-ы)", или же ещё хуже - "осталось 1 минуты". durafmt_ru позволяет отформатировать временной интервал вызовом одной функции, не задумываясь о сложности и красоте русского языка.

## Использование

Библиотека старается быть drop-in replacement для оригинальной durafmt, но никто не гарантирует, что так будет всегда.

### durafmt_ru.ParseString()

```go
package main

import (
	"fmt"
	
	"github.com/fat0troll/durafmt_ru"
)

func main() {
	duration, err := durafmt_ru.ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(duration) // 2 недели 18 часов 22 минуты 3 секунды
	// duration.String() // "2 недели 18 часов 22 минуты 3 секунды"
}
```

### durafmt_ru.ParseStringShort()

Версия функции `durafmt_ru.ParseString()`, возвращающая только первый (наибольший) элемент временной продолжительности.

```go
package main

import (
	"fmt"
	
	"github.com/fat0troll/durafmt_ru"
)

func main() {
	duration, err := durafmt_ru.ParseStringShort("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(duration) // 2 недели
	// duration.String() // "2 недели"
}
```

### durafmt_ru.Parse()

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/fat0troll/durafmt_ru"
)

func main() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := durafmt.Parse(timeduration).String()

	fmt.Println(duration) // 2 недели 18 часов 22 минуты 3 секунды
}
```

#### LimitFirstN()

Функция похожа на `durafmt_ru.ParseStringShort()`, но позволяет оставить первые N частей строки вместо одной.

```go
package main

import (
	"fmt"
	"time"
	
	"github.com/fat0troll/durafmt_ru"
)

func main() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := durafmt.Parse(timeduration).LimitFirstN(2) // ограничимся двумя первыми частями

	fmt.Println(duration) // 2 недели 18 часов
}
```

## Помощь и участие в разработке библиотеки

Помощь приветствуется! Форкайте репозиторий, меняйте его, присылайте пулл-реквесты.

В разделе Issues гитхаба можно оставлять описание багов, предложение фич и впечатления от работы с библиотекой.

Тесты можно прогнать с помощью `go test`, кроме того, библиотека тестируется линтером `golangci-lint`.

Перед присыланием пулл-реквеста запуск `go test` и `golangci-lint` обязателен.

## Заключение

Отдельная благодарность Wesley Hill и всем авторам и контрибьюторам в библиотеку durafmt.

## Лицензия

MIT, см. файл LICENSE.
