// Библиотека durafmt_ru форматирует строку из типа time.Duration в человекочитаемый формат
// на русском языке.
// В качестве основы взята подобная библиотека durafmt (https://github.com/hako/durafmt/),
// существующая только для английского языка.
package durufmt

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	// Константы каноничных имён единиц времени.
	Microseconds = "microseconds"
	Milliseconds = "milliseconds"
	Seconds      = "seconds"
	Minutes      = "minutes"
	Hours        = "hours"
	Days         = "days"
	Weeks        = "weeks"
	Years        = "years"

	// Типы единственного и множественного чисел для выражения числительных на русском языке.
	Singular = "one"  // 1, 21, 31... (но не 11)
	Some     = "some" // 2, 3, 4, 22, 23, 24... (но не 12, 13 и 14)
	Many     = "many" // 5, 15, 25, 35... (а так же 11, 12, 13 и 14)
)

var (
	units      = []string{Years, Weeks, Days, Hours, Minutes, Seconds, Milliseconds, Microseconds}
	unitsShort = []string{"y", "w", "d", "h", "m", "s", "ms", "µs"}
	unitNames  = map[string]map[string]string{
		Years: {
			Singular: "год",
			Some:     "года",
			Many:     "лет",
		},
		Weeks: {
			Singular: "неделя",
			Some:     "недели",
			Many:     "недель",
		},
		Days: {
			Singular: "день",
			Some:     "дня",
			Many:     "дней",
		},
		Hours: {
			Singular: "час",
			Some:     "часа",
			Many:     "часов",
		},
		Minutes: {
			Singular: "минута",
			Some:     "минуты",
			Many:     "минут",
		},
		Seconds: {
			Singular: "секунда",
			Some:     "секунды",
			Many:     "секунд",
		},
		Milliseconds: {
			Singular: "миллисекунда",
			Some:     "миллисекунды",
			Many:     "миллисекунд",
		},
		Microseconds: {
			Singular: "микросекунда",
			Some:     "микросекунды",
			Many:     "микросекунд",
		},
	}
)

// Durafmt хранит в себе спарсированный интервал времени и оригинальный ввод пользователя.
type Durafmt struct {
	duration  time.Duration
	input     string // Справочная информация.
	limitN    int    // В случае ненулевого значения ограничивает количество выдаваемых элементов в результате.
	limitUnit string // Непустое значение лимитирует максимальную единицу времени для выдачи.
}

// LimitToUnit устанавливает формат вывода, вы не получите в итоговой строке единицу времени больше заданной.
// unit = "" означает отсутствие ограничений.
func (d *Durafmt) LimitToUnit(unit string) *Durafmt {
	d.limitUnit = unit

	return d
}

// LimitFirstN устанавливает формат вывода, ограничивая в нём количество элементов до n.
// n == 0 означает отсутствие лимита.
func (d *Durafmt) LimitFirstN(n int) *Durafmt {
	d.limitN = n

	return d
}

func (d *Durafmt) Duration() time.Duration {
	return d.duration
}

// Parse создаёт новую структуру *Durafmt. Возвращает ошибку в случае неправильных входных параметров.
func Parse(dinput time.Duration) *Durafmt {
	input := dinput.String()

	return &Durafmt{dinput, input, 0, ""}
}

// ParseShort создаёт новую структуру *Durafmt, краткой формы. Возвращает ошибку в случае неправильных
// входных параметров. Синоним `Parse(dur).LimitFirstN(1)`.
func ParseShort(dinput time.Duration) *Durafmt {
	input := dinput.String()

	return &Durafmt{dinput, input, 1, ""}
}

// ParseString создаёт структуру *Durafmt из строки. Формат строки аналогичен используемому в durafmt.
// Возвращает ошибку в случае неправильных входных данных.
func ParseString(input string) (*Durafmt, error) {
	if input == "0" || input == "-0" {
		return nil, errors.New("durafmt_ru: не указана единица времени во входном параметре " + input)
	}

	duration, err := time.ParseDuration(input)
	if err != nil {
		return nil, err
	}

	return &Durafmt{duration, input, 0, ""}, nil
}

// ParseStringShort создаёт структуру *Durafmt из строки, краткой формы. Формат строки аналогичен
// используемому в durafmt. Возвращает ошибку в случае неправильных входных данных.
// Синоним вызова `ParseString(durStr)` и следующего за ним `LimitFirstN(1)`.
func ParseStringShort(input string) (*Durafmt, error) {
	if input == "0" || input == "-0" {
		return nil, errors.New("durafmt_ru: не указана единица времени во входном параметре " + input)
	}

	duration, err := time.ParseDuration(input)
	if err != nil {
		return nil, err
	}

	return &Durafmt{duration, input, 1, ""}, nil
}

// String форматирует *Durafmt в человекочитаемый вид.
func (d *Durafmt) String() string {
	var duration string

	// Check for minus durations.
	if string(d.input[0]) == "-" {
		duration += "-"
		d.duration = -d.duration
	}

	var microseconds int64
	var milliseconds int64
	var seconds int64
	var minutes int64
	var hours int64
	var days int64
	var weeks int64
	var years int64

	shouldConvert := false
	remainingSecondsToConvert := int64(d.duration / time.Microsecond)

	// Convert duration.
	if d.limitUnit == "" {
		shouldConvert = true
	}

	if d.limitUnit == Years || shouldConvert {
		years = remainingSecondsToConvert / (365 * 24 * 3600 * 1000000)
		remainingSecondsToConvert -= years * 365 * 24 * 3600 * 1000000
		shouldConvert = true
	}

	if d.limitUnit == Weeks || shouldConvert {
		weeks = remainingSecondsToConvert / (7 * 24 * 3600 * 1000000)
		remainingSecondsToConvert -= weeks * 7 * 24 * 3600 * 1000000
		shouldConvert = true
	}

	if d.limitUnit == Days || shouldConvert {
		days = remainingSecondsToConvert / (24 * 3600 * 1000000)
		remainingSecondsToConvert -= days * 24 * 3600 * 1000000
		shouldConvert = true
	}

	if d.limitUnit == Hours || shouldConvert {
		hours = remainingSecondsToConvert / (3600 * 1000000)
		remainingSecondsToConvert -= hours * 3600 * 1000000
		shouldConvert = true
	}

	if d.limitUnit == Minutes || shouldConvert {
		minutes = remainingSecondsToConvert / (60 * 1000000)
		remainingSecondsToConvert -= minutes * 60 * 1000000
		shouldConvert = true
	}

	if d.limitUnit == Seconds || shouldConvert {
		seconds = remainingSecondsToConvert / 1000000
		remainingSecondsToConvert -= seconds * 1000000
		shouldConvert = true
	}

	if d.limitUnit == Milliseconds || shouldConvert {
		milliseconds = remainingSecondsToConvert / 1000
		remainingSecondsToConvert -= milliseconds * 1000
	}

	microseconds = remainingSecondsToConvert

	// Create a map of the converted duration time.
	durationMap := map[string]int64{
		Microseconds: microseconds,
		Milliseconds: milliseconds,
		Seconds:      seconds,
		Minutes:      minutes,
		Hours:        hours,
		Days:         days,
		Weeks:        weeks,
		Years:        years,
	}

	duration += d.buildDuration(durationMap)

	// Если запрошена краткая версия и в строке более limitN*2 пробелов, возвращаем
	// первые limitN*2 подстроки.
	if d.limitN > 0 {
		parts := strings.Split(duration, " ")

		if len(parts) > d.limitN*2 {
			duration = strings.Join(parts[:d.limitN*2], " ")
		}
	}

	return duration
}

func (d *Durafmt) buildDuration(durationMap map[string]int64) string {
	duration := ""

	// Construct duration string.
	for idx := range units {
		uKey := units[idx]
		u := unitNames[uKey]
		v := durationMap[uKey]
		strval := strconv.FormatInt(v, 10)

		if d.duration.String() == "0" || d.duration.String() == "0s" {
			pattern := fmt.Sprintf("^-?0%s$", unitsShort[idx])

			isMatch, err := regexp.MatchString(pattern, d.input)
			if err != nil {
				return ""
			}

			if isMatch {
				duration += strval + " " + u[Many]
			}
		}

		if v == 0 {
			// Пропускаем любой элемент со значением 0.
			continue
		}

		switch v % 10 {
		case 1:
			if v%100 == 11 {
				duration += strval + " " + u[Many] + " "
			} else {
				duration += strval + " " + u[Singular] + " "
			}

		case 2, 3, 4:
			if v%100 == 12 || v%100 == 13 || v%100 == 14 {
				duration += strval + " " + u[Many] + " "
			} else {
				duration += strval + " " + u[Some] + " "
			}

		default:
			duration += strval + " " + u[Many] + " "
		}
	}

	// Удаляем лишние пробелы.
	return strings.TrimSpace(duration)
}
