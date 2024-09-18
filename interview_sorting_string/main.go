package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := []string{"v2.4.3", "v3.1.0", "v3.1.1", "v3.1.10", "v3.1.2", "v3.1.3", "v3.1.4", "v3.1.5", "v3.1.6", "v3.1.7", "v3.1.8", "v3.1.9"}
	fmt.Println(s)
	sort.Slice(s, func(i, j int) bool {
		return sort2(s[i], s[j])
	})
	fmt.Println(s)
}
func sort2(a, b string) bool {
	x := strings.Split(strings.TrimLeft(a, "v"), ".")
	y := strings.Split(strings.TrimLeft(b, "v"), ".")

	for i := 0; i < len(x) && i < len(y); i++ {
		xInt, _ := strconv.Atoi(x[i])
		yInt, _ := strconv.Atoi(y[i])

		if xInt != yInt {
			return xInt < yInt
		}
	}
	return len(x) < len(y)
}
