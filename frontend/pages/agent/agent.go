package pages

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func formatTotalWasteWeighed(f float64) string {
	switch {
	case f < 10_000:
		return formatFloatWithComma(f, 1)
	case f < 100_000:
		return fmt.Sprintf("%.1fk", math.Round(f/1_000*10)/10)
	case f < 1_000_000:
		return fmt.Sprintf("%.1fk", math.Round(f/1_000*10)/10)
	default:
		return fmt.Sprintf("%.1fM", math.Round(f/1_000_000*10)/10)
	}
}

func formatFloatWithComma[T float32 | float64](f T, dp int) string {
	if dp < 0 {
		dp = 1
	}

	factor := math.Pow(10, float64(dp))
	ff := math.Round(float64(f)*factor) / factor

	formatStr := fmt.Sprintf("%%.%df", dp)
	formatted := fmt.Sprintf(formatStr, ff)

	parts := strings.Split(formatted, ".")

	wholeWithCommas := addCommasToNumber(parts[0])

	if dp > 0 {
		return wholeWithCommas + "." + parts[1]
	}
	return wholeWithCommas
}

// addCommasToNumber inserts commas into a whole number string
func addCommasToNumber(num string) string {
	n := len(num)
	if n <= 3 {
		return num
	}

	var result []string
	start := n % 3
	if start > 0 {
		result = append(result, num[:start])
	}

	for i := start; i < n; i += 3 {
		if i > 0 {
			result = append(result, ",")
		}
		result = append(result, num[i:i+3])
	}

	return strings.Join(result, "")
}

func formatSparklineValues(v []float64) string {
	s := make([]string, len(v))
	for i, val := range v {
		s[i] = strconv.Itoa(int(math.Round(val)))
	}
	return strings.Join(s, ",")
}

func colorForSparkline(v []float64) string {
	if len(v) < 1 || math.Round(v[len(v)-2]) > math.Round(v[len(v)-1]) {
		return "red"
	}
	return "green"
}
