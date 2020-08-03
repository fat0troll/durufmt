package durafmt_ru

import (
	"fmt"
	"testing"
	"time"
)

var (
	testStrings []struct {
		test     string
		expected string
	}
	testTimes []struct {
		test     time.Duration
		expected string
	}
	testTimesWithLimitUnit []struct {
		test      time.Duration
		limitUnit string
		expected  string
	}
	testTimesWithLimit []struct {
		test     time.Duration
		limitN   int
		expected string
	}
)

// TestParse тестирует форматирование time.Duration.
func TestParse(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Microsecond, "1 микросекунда"},
		{1 * time.Millisecond, "1 миллисекунда"},
		{1 * time.Second, "1 секунда"},
		{1 * time.Hour, "1 час"},
		{1 * time.Minute, "1 минута"},
		{2 * time.Microsecond, "2 микросекунды"},
		{2 * time.Millisecond, "2 миллисекунды"},
		{2 * time.Second, "2 секунды"},
		{2 * time.Minute, "2 минуты"},
		{1 * time.Hour, "1 час"},
		{2 * time.Hour, "2 часа"},
		{10 * time.Hour, "10 часов"},
		{24 * time.Hour, "1 день"},
		{48 * time.Hour, "2 дня"},
		{120 * time.Hour, "5 дней"},
		{168 * time.Hour, "1 неделя"},
		{672 * time.Hour, "4 недели"},
		{8759 * time.Hour, "52 недели 23 часа"},
		{8760 * time.Hour, "1 год"},
		{17519 * time.Hour, "1 год 52 недели 23 часа"},
		{17520 * time.Hour, "2 года"},
		{26279 * time.Hour, "2 года 52 недели 23 часа"},
		{26280 * time.Hour, "3 года"},
		{201479 * time.Hour, "22 года 52 недели 23 часа"},
		{201480 * time.Hour, "23 года"},
		{-1 * time.Second, "-1 секунда"},
		{-10 * time.Second, "-10 секунд"},
		{-100 * time.Second, "-1 минута 40 секунд"},
		{-1 * time.Millisecond, "-1 миллисекунда"},
		{-10 * time.Millisecond, "-10 миллисекунд"},
		{-100 * time.Millisecond, "-100 миллисекунд"},
		{-1 * time.Microsecond, "-1 микросекунда"},
		{-10 * time.Microsecond, "-10 микросекунд"},
		{-100 * time.Microsecond, "-100 микросекунд"},
		{-1000 * time.Microsecond, "-1 миллисекунда"},
		{-1000000 * time.Microsecond, "-1 секунда"},
		{-1001000 * time.Microsecond, "-1 секунда 1 миллисекунда"},
		{-1010000 * time.Microsecond, "-1 секунда 10 миллисекунд"},
		{-1001001 * time.Microsecond, "-1 секунда 1 миллисекунда 1 микросекунда"},
		{-1001002 * time.Microsecond, "-1 секунда 1 миллисекунда 2 микросекунды"},
	}

	for _, table := range testTimes {
		result := Parse(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. получено %q, ожидалось %q",
				table.test, result, result, table.expected)
		}
	}
}

func TestParseWithLimitToUnit(t *testing.T) {
	testTimesWithLimitUnit = []struct {
		test      time.Duration
		limitUnit string
		expected  string
	}{
		{87593183 * time.Second, Seconds, "87593183 секунды"},
		{87593183 * time.Second, Minutes, "1459886 минут 23 секунды"},
		{87593183 * time.Second, Hours, "24331 час 26 минут 23 секунды"},
		{87593183 * time.Second, Days, "1013 дней 19 часов 26 минут 23 секунды"},
		{87593183 * time.Second, Weeks, "144 недели 5 дней 19 часов 26 минут 23 секунды"},
		{87593183 * time.Second, Years, "2 года 40 недель 3 дня 19 часов 26 минут 23 секунды"},
		{87593183 * time.Second, "", "2 года 40 недель 3 дня 19 часов 26 минут 23 секунды"},
	}

	for _, table := range testTimesWithLimitUnit {
		result := Parse(table.test).LimitToUnit(table.limitUnit).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. получено %q, ожидалось %q",
				table.test, result, result, table.expected)
		}
	}
}

func TestParseWithLimitN(t *testing.T) {
	testTimesWithLimit = []struct {
		test     time.Duration
		limitN   int
		expected string
	}{
		{1 * time.Millisecond, 0, "1 миллисекунда"},
		{8759 * time.Hour, 0, "52 недели 23 часа"},
		{17519 * time.Hour, 0, "1 год 52 недели 23 часа"},
		{-1 * time.Second, 0, "-1 секунда"},
		{-100 * time.Second, 0, "-1 минута 40 секунд"},
		{1 * time.Millisecond, 1, "1 миллисекунда"},
		{8759 * time.Hour, 1, "52 недели"},
		{17519 * time.Hour, 1, "1 год"},
		{-1 * time.Second, 1, "-1 секунда"},
		{-100 * time.Second, 1, "-1 минута"},
		{1 * time.Millisecond, 2, "1 миллисекунда"},
		{8759 * time.Hour, 2, "52 недели 23 часа"},
		{17519 * time.Hour, 2, "1 год 52 недели"},
		{-1 * time.Second, 2, "-1 секунда"},
		{-100 * time.Second, 2, "-1 минута 40 секунд"},
		{1 * time.Millisecond, 3, "1 миллисекунда"},
		{8759 * time.Hour, 3, "52 недели 23 часа"},
		{17519 * time.Hour, 3, "1 год 52 недели 23 часа"},
		{-1 * time.Second, 3, "-1 секунда"},
		{-100 * time.Second, 3, "-1 минута 40 секунд"},
	}

	for _, table := range testTimesWithLimit {
		result := Parse(table.test).LimitFirstN(table.limitN).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. получено %q, ожидалось %q",
				table.test, result, result, table.expected)
		}
	}
}

// TestParseShort тестирует форматирование time.Duration, краткая версия.
func TestParseShort(t *testing.T) {
	testTimes = []struct {
		test     time.Duration
		expected string
	}{
		{1 * time.Microsecond, "1 микросекунда"},
		{1 * time.Millisecond, "1 миллисекунда"},
		{1 * time.Second, "1 секунда"},
		{1 * time.Hour, "1 час"},
		{1 * time.Minute, "1 минута"},
		{2 * time.Microsecond, "2 микросекунды"},
		{2 * time.Millisecond, "2 миллисекунды"},
		{2 * time.Second, "2 секунды"},
		{2 * time.Minute, "2 минуты"},
		{1 * time.Hour, "1 час"},
		{2 * time.Hour, "2 часа"},
		{10 * time.Hour, "10 часов"},
		{24 * time.Hour, "1 день"},
		{48 * time.Hour, "2 дня"},
		{120 * time.Hour, "5 дней"},
		{168 * time.Hour, "1 неделя"},
		{672 * time.Hour, "4 недели"},
		{8759 * time.Hour, "52 недели"},
		{8760 * time.Hour, "1 год"},
		{17519 * time.Hour, "1 год"},
		{17520 * time.Hour, "2 года"},
		{26279 * time.Hour, "2 года"},
		{26280 * time.Hour, "3 года"},
		{201479 * time.Hour, "22 года"},
		{201480 * time.Hour, "23 года"},
		{-1 * time.Second, "-1 секунда"},
		{-10 * time.Second, "-10 секунд"},
		{-100 * time.Second, "-1 минута"},
		{-1 * time.Millisecond, "-1 миллисекунда"},
		{-10 * time.Millisecond, "-10 миллисекунд"},
		{-100 * time.Millisecond, "-100 миллисекунд"},
		{-1 * time.Microsecond, "-1 микросекунда"},
		{-10 * time.Microsecond, "-10 микросекунд"},
		{-100 * time.Microsecond, "-100 микросекунд"},
		{-1000 * time.Microsecond, "-1 миллисекунда"},
		{-1000000 * time.Microsecond, "-1 секунда"},
		{-1001000 * time.Microsecond, "-1 секунда"},
		{-1010000 * time.Microsecond, "-1 секунда"},
		{-1001001 * time.Microsecond, "-1 секунда"},
		{-1001002 * time.Microsecond, "-1 секунда"},
	}

	for _, table := range testTimes {
		result := ParseShort(table.test).String()
		if result != table.expected {
			t.Errorf("Parse(%q).String() = %q. получено %q, ожидалось %q",
				table.test, result, result, table.expected)
		}
	}
}

// TestParseString тестирует конвертирование строковых входных данных.
func TestParseString(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1µs", "1 микросекунда"},
		{"2µs", "2 микросекунды"},
		{"1ms", "1 миллисекунда"},
		{"2ms", "2 миллисекунды"},
		{"1s", "1 секунда"},
		{"2s", "2 секунды"},
		{"1m", "1 минута"},
		{"2m", "2 минуты"},
		{"1h", "1 час"},
		{"2h", "2 часа"},
		{"10h", "10 часов"},
		{"24h", "1 день"},
		{"48h", "2 дня"},
		{"120h", "5 дней"},
		{"168h", "1 неделя"},
		{"672h", "4 недели"},
		{"8759h", "52 недели 23 часа"},
		{"8760h", "1 год"},
		{"17519h", "1 год 52 недели 23 часа"},
		{"17520h", "2 года"},
		{"26279h", "2 года 52 недели 23 часа"},
		{"26280h", "3 года"},
		{"201479h", "22 года 52 недели 23 часа"},
		{"201480h", "23 года"},
		{"1m0s", "1 минута"},
		{"1m2s", "1 минута 2 секунды"},
		{"3h4m5s", "3 часа 4 минуты 5 секунд"},
		{"6h7m8s9ms", "6 часов 7 минут 8 секунд 9 миллисекунд"},
		{"0µs", "0 микросекунд"},
		{"0ms", "0 миллисекунд"},
		{"0s", "0 секунд"},
		{"0m", "0 минут"},
		{"0h", "0 часов"},
		{"0m1µs", "1 микросекунда"},
		{"0m1ms1µs", "1 миллисекунда 1 микросекунда"},
		{"0m1ms", "1 миллисекунда"},
		{"0m1s", "1 секунда"},
		{"0m1m", "1 минута"},
		{"0m2ms", "2 миллисекунды"},
		{"0m2s", "2 секунды"},
		{"0m2m", "2 минуты"},
		{"0m2m3h", "3 часа 2 минуты"},
		{"0m2m34h", "1 день 10 часов 2 минуты"},
		{"0m56h7m8ms", "2 дня 8 часов 7 минут 8 миллисекунд"},
		{"-1µs", "-1 микросекунда"},
		{"-1ms", "-1 миллисекунда"},
		{"-1s", "-1 секунда"},
		{"-1m", "-1 минута"},
		{"-1h", "-1 час"},
		{"-2µs", "-2 микросекунды"},
		{"-2ms", "-2 миллисекунды"},
		{"-2s", "-2 секунды"},
		{"-2m", "-2 минуты"},
		{"-2h", "-2 часа"},
		{"-10h", "-10 часов"},
		{"-24h", "-1 день"},
		{"-48h", "-2 дня"},
		{"-120h", "-5 дней"},
		{"-168h", "-1 неделя"},
		{"-672h", "-4 недели"},
		{"-8760h", "-1 год"},
		{"-1m0s", "-1 минута"},
		{"-0m2s", "-2 секунды"},
		{"-0m2m", "-2 минуты"},
		{"-0m2m3h", "-3 часа 2 минуты"},
		{"-0m2m34h", "-1 день 10 часов 2 минуты"},
		{"-0µs", "-0 микросекунд"},
		{"-0ms", "-0 миллисекунд"},
		{"-0s", "-0 секунд"},
		{"-0m", "-0 минут"},
		{"-0h", "-0 часов"},
	}

	for _, table := range testStrings {
		d, err := ParseString(table.test)
		if err != nil {
			t.Errorf("%q", err)
		}
		result := d.String()
		if result != table.expected {
			t.Errorf("d.String() = %q. получено %q, ожидалось %q",
				table.test, result, table.expected)
		}
	}
}

// TestParseStringShort тестирует конвертирование строковых входных данных, краткая версия.
func TestParseStringShort(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1µs", "1 микросекунда"},
		{"1ms", "1 миллисекунда"},
		{"2µs", "2 микросекунды"},
		{"2ms", "2 миллисекунды"},
		{"1s", "1 секунда"},
		{"2s", "2 секунды"},
		{"1m", "1 минута"},
		{"2m", "2 минуты"},
		{"1h", "1 час"},
		{"2h", "2 часа"},
		{"10h", "10 часов"},
		{"24h", "1 день"},
		{"48h", "2 дня"},
		{"120h", "5 дней"},
		{"168h", "1 неделя"},
		{"672h", "4 недели"},
		{"8759h", "52 недели"},
		{"8760h", "1 год"},
		{"17519h", "1 год"},
		{"17520h", "2 года"},
		{"26279h", "2 года"},
		{"26280h", "3 года"},
		{"201479h", "22 года"},
		{"201480h", "23 года"},
		{"1m0s", "1 минута"},
		{"1m2s", "1 минута"},
		{"3h4m5s", "3 часа"},
		{"6h7m8s9ms", "6 часов"},
		{"0µs", "0 микросекунд"},
		{"0ms", "0 миллисекунд"},
		{"0s", "0 секунд"},
		{"0m", "0 минут"},
		{"0h", "0 часов"},
		{"0m1µs", "1 микросекунда"},
		{"0m1ms1µs", "1 миллисекунда"},
		{"0m1ms", "1 миллисекунда"},
		{"0m1s", "1 секунда"},
		{"0m1m", "1 минута"},
		{"0m2ms", "2 миллисекунды"},
		{"0m2s", "2 секунды"},
		{"0m2m", "2 минуты"},
		{"0m2m3h", "3 часа"},
		{"0m2m34h", "1 день"},
		{"0m56h7m8ms", "2 дня"},
		{"-1µs", "-1 микросекунда"},
		{"-1ms", "-1 миллисекунда"},
		{"-1s", "-1 секунда"},
		{"-1m", "-1 минута"},
		{"-1h", "-1 час"},
		{"-2µs", "-2 микросекунды"},
		{"-2ms", "-2 миллисекунды"},
		{"-2s", "-2 секунды"},
		{"-2m", "-2 минуты"},
		{"-2h", "-2 часа"},
		{"-10h", "-10 часов"},
		{"-24h", "-1 день"},
		{"-48h", "-2 дня"},
		{"-120h", "-5 дней"},
		{"-168h", "-1 неделя"},
		{"-672h", "-4 недели"},
		{"-8760h", "-1 год"},
		{"-1m0s", "-1 минута"},
		{"-0m2s", "-2 секунды"},
		{"-0m2m", "-2 минуты"},
		{"-0m2m3h", "-3 часа"},
		{"-0m2m34h", "-1 день"},
		{"-0µs", "-0 микросекунд"},
		{"-0ms", "-0 миллисекунд"},
		{"-0s", "-0 секунд"},
		{"-0m", "-0 минут"},
		{"-0h", "-0 часов"},
	}

	for _, table := range testStrings {
		d, err := ParseStringShort(table.test)
		if err != nil {
			t.Errorf("%q", err)
		}
		result := d.String()
		if result != table.expected {
			t.Errorf("d.String() = %q. получено %q, ожидалось %q",
				table.test, result, table.expected)
		}
	}
}

// TestInvalidDuration for invalid inputs.
func TestInvalidDuration(t *testing.T) {
	testStrings = []struct {
		test     string
		expected string
	}{
		{"1", ""},
		{"1d", ""},
		{"1w", ""},
		{"1wk", ""},
		{"1y", ""},
		{"", ""},
		{"m1", ""},
		{"1nmd", ""},
		{"0", ""},
		{"-0", ""},
	}

	for _, table := range testStrings {
		_, err := ParseString(table.test)
		if err == nil {
			t.Errorf("ParseString(%q). получено %q, ожидалось %q",
				table.test, err, table.expected)
		}
	}

	for _, table := range testStrings {
		_, err := ParseStringShort(table.test)
		if err == nil {
			t.Errorf("ParseString(%q). получено %q, ожидалось %q",
				table.test, err, table.expected)
		}
	}
}

// Benchmarks

func BenchmarkParse(b *testing.B) {
	for n := 1; n < b.N; n++ {
		Parse(time.Duration(n) * time.Hour)
	}
}

func BenchmarkParseStringShort(b *testing.B) {
	for n := 1; n < b.N; n++ {
		ParseStringShort(fmt.Sprintf("%dh", n))
	}
}

func BenchmarkParseString(b *testing.B) {
	for n := 1; n < b.N; n++ {
		ParseString(fmt.Sprintf("%dh", n))
	}
}
