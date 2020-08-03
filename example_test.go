package durafmt_ru

import (
	"fmt"
	"math"
	"time"
)

func ExampleParseString() {
	duration, err := ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(duration) // 2 недели 18 часов 22 минуты 3 секунды
	// duration.String() // "2 недели 18 часов 22 минуты 3 секунды"
}

func ExampleDurafmt_LimitFirstN() {
	duration, err := ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}

	duration = duration.LimitFirstN(2)

	fmt.Println(duration) // 2 недели 18 часов
	// duration.String() // "2 недели 18 часов"
}

func ExampleDurafmt_LimitToUnit() {
	duration, err := ParseString("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}

	duration = duration.LimitToUnit(Days)

	fmt.Println(duration) // 14 дней 18 часов 22 минуты 3 секунды
	// duration.String() // "14 дней 18 часов 22 минуты 3 секунды"
}

func ExampleParseString_sequence() {
	for hours := 1.0; hours < 12.0; hours++ {
		hour := fmt.Sprintf("%fh", math.Pow(2, hours))
		duration, err := ParseString(hour)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(duration) // 2 часа, 4 часа, 6 часов...
	}
}

// Версия durafmt_ru.ParseString(), возвращающая только первую часть строки с продолжительностью времени.
func ExampleParseStringShort() {
	duration, err := ParseStringShort("354h22m3.24s")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(duration) // 2 недели
	// duration.String() // "2 недели"
}

func ExampleParse() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := Parse(timeduration).String()

	fmt.Println(duration) // 2 недели 18 часов 22 минуты 3 секунды
}

// Версия durafmt_ru.Parse(), возвращающая только первую часть строки с продолжительностью времени.
func ExampleParseShort() {
	timeduration := (354 * time.Hour) + (22 * time.Minute) + (3 * time.Second)
	duration := ParseShort(timeduration).String()

	fmt.Println(duration) // 2 недели
}
