package main

import (
	"bytes"
	cryptoRand "crypto/rand"
	"flag"
	"fmt"
	"time"
)

var max int
var fps int
var count int

func randomArray(n int, max int) []int {
	var i int
	var number float64
	arr := make([]int, n)

	for i = 0; i < n; i++ {
		b := make([]byte, 1)
		cryptoRand.Read(b)
		number = float64(b[0])
		arr[i] = int(number / 255 * float64(max))
	}
	return arr
}

func visualize(arr []int) {
	var buffer bytes.Buffer
	var x int
	var y int

	for y = 0; y < max; y++ {
		for x = 0; x < count; x++ {
			if arr[x] < y {
				buffer.WriteByte(byte('#'))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}
	time.Sleep(time.Second / time.Duration(fps))
	fmt.Print("\033[2J")
	fmt.Print(buffer.String())
}

func bubbleSort(arr []int) {
	var tmp int = 0
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				tmp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
		visualize(arr)
	}
}

func gnomeSort(arr []int) {
	var i int = 1
	var tmp int = 0

	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			tmp = arr[i]
			arr[i] = arr[i-1]
			arr[i-1] = tmp
			if i > 1 {
				i--
			}
		}
		visualize(arr)
	}
}

func selectionSort(arr []int) {
	var min int = 0
	var tmp int = 0
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		min = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		tmp = arr[i]
		arr[i] = arr[min]
		arr[min] = tmp
		visualize(arr)
	}
}

func main() {
	var algo string
	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm [bubble]/gnome/selection")
	flag.IntVar(&fps, "fps", 10, "frames per second")
	flag.IntVar(&max, "max", 9, "highest value")
	flag.IntVar(&count, "count", 30, "number of values")
	flag.Parse()
	arr := randomArray(count, max)
	fmt.Printf("sorting via %v-sort\nhighest value: %v\nnumber of values: %v\n\n", algo, max, count)
	time.Sleep(time.Second * 1)
	switch algo {
	case "bubble":
		bubbleSort(arr)
	case "gnome":
		gnomeSort(arr)
	case "selection":
		selectionSort(arr)
	}
}
