package main

import (
	"fmt"
	"strings"
	// "regexp"
	// "strings"
)

func main() {
	prima()
	deret_angka()
	count_string()
}

func prima() {
	var i, loop int

	fmt.Print("masukkan bilangan prima : ")
	fmt.Scan(&i)

	if !primaCheck(i) {
		fmt.Printf("%d bukan bilangan prima \n", i)
		return
	}

	fmt.Print("masukkan jumlah perulangan : ")
	fmt.Scan(&loop)

	countX := 1
	countY := 1
	temp := i

	for countX <= loop {
		i = temp
		countY = 1
		for i < 30 && countY <= loop {
			if primaCheck(i) && countY <= countX {
				fmt.Printf("%d ", i)
				countY++
			}
			i++
		}
		fmt.Print("\n")
		countX++
	}
}

func deret_angka() {
	var loop int
	fmt.Printf("masukan jumlah perulangan : ")
	fmt.Scan(&loop)
	var result int
	result = 1
	for i := 1; i <= loop; i++ {
		result = i * result
		fmt.Printf("%d \n", result)
	}
}

func count_string() {
	var sample string
	abjad := "abcdefghijklmnopqrstuvwxyz"

	fmt.Print("masukkan string : ")
	fmt.Scan(&sample)

	for i := 0; i < len(abjad); i++ {
		count := strings.Count(sample, string(abjad[i]))
		if count > 0 {
			fmt.Print(string(abjad[i]))
			fmt.Print(" : ")
			fmt.Printf("%d ", count)
		}
	}
}

func primaCheck(i int) bool {
	for y := 2; y < i; y++ {
		if i%y == 0 {
			return false
		}
	}
	return true
}
