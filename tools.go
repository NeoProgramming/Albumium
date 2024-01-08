package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

func WhereAnd(prev string) string {
	if prev != "" {
		return " AND"
	}
	return " WHERE"
}

func Atoi(s string) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func Atodi(s string, d int) int {
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return d
}

func Ttoa(t time.Time) string {
	if t == time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC) {
		return "---"
	}
	return t.Format("06-01-02 15:04")
}

func Tmtoa(t int64) string {
	if t == 0 {
		return "---"
	}
	return time.Unix(t, 0).Format("06-01-02 15:04")
}

func Btoa(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

func arr(args ...interface{}) []interface{} {
	return args
}

func minTime(t ...int64) int64 {
	var mt int64 = math.MaxInt64
	for _, t := range t {
		if t < mt {
			mt = t
		}
	}
	return mt
}

func maxTime(t ...int64) int64 {
	var mt int64 = 0
	for _, t := range t {
		if mt < t {
			mt = t
		}
	}
	return mt
}

func decodeFilterMasks(filters string) (int, int) {
	m := 0
	im := 0

	for i := 0; i < len(filters); i++ {
		if filters[i] == '1' {
			m |= 1 << i
		}
		if filters[i] == '2' {
			im |= 1 << i
		}
	}
	fmt.Printf("decodeFilterMasks: %s  => %02X %02X\n", filters, m, im)
	return m, im
}
