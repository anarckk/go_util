package go_number

import "fmt"

func GetPercentStr(p float64) string {
	return fmt.Sprintf("%.2f%%", p*100)
}
