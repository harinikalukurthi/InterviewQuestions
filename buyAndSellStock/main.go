package main

import "fmt"

func main() {

	a := []int{7, 6, 2, 4, 3, 9}
	fmt.Println(buyandsell(a))
	fmt.Println(efficient(a))

}

func buyandsell(a []int) (int, int) {
	max := 0
	x, y := 0, 0
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				if max < a[j]-a[i] {
					max = a[j] - a[i]
					x, y = i, j
				}
			}
		}
	}
	fmt.Println(max)
	return x, y
}

func efficient(x []int) (int) {
	buy_price := x[0]
	cur := 0
	max_profit := 0
	for i := 0; i < len(x); i++ {
		if x[i] < buy_price {
			buy_price = x[i]
		}else{
			cur = x[i] - buy_price
			if cur>max_profit{
			max_profit=cur
			}
		}
	}
	return max_profit
}
