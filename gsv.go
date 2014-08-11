package main

import (
	"bytes"
	cryptoRand "crypto/rand"
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
	"time"
)

var max int
var fps int
var count int
var mode int

func buildImage(arr []int) *image.Paletted {
	var frame = image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(arr), max},
		},
		color.Palette{
			color.Gray{uint8(255)},
			color.Gray{uint8(0)},
		},
	)
	//fmt.Println(arr)
	for k, v := range arr {
		frame.SetColorIndex(k, max-v, uint8(1))
		if mode == 2 {
			for y := max - v + 1; y < max; y++ {
				frame.SetColorIndex(k, y, uint8(1))
			}
		}
	}
	return frame
}

func writeGif(name string, g *gif.GIF) {
	//fmt.Println("size of g == ", len(g.Image))
	w, err := os.Create(name + ".gif")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Println(err)
			panic(err)
		}
	}()
	err = gif.EncodeAll(w, g)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

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
		for x = 0; x < len(arr); x++ {
			if arr[x] == y {
				buffer.WriteByte(byte('#'))
			} else if arr[x] < y && mode == 1 {
				buffer.WriteByte(byte('#'))
			} else if arr[x] > y && mode == 2 {
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

func shuffle(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

func isSorted(arr []int) bool {
	for i := len(arr); i > 1; i-- {
		if arr[i-1] < arr[i-2] {
			return false
		}
	}
	return true
}

func bogoSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	for isSorted(arr) == false {
		arr = shuffle(arr)
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("bogo", g)
}

func bubbleSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			//visualize(arr)
			g.Image = append(g.Image, buildImage(arr))
			g.Delay = append(g.Delay, 2)
			/*
				if len(g.Image) > 1 {
					last := g.Image[len(g.Image)-1]
					prev := g.Image[len(g.Image)-2]
					//fmt.Println(last)
					fmt.Println(prev)
				}
			*/
		}
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("bubble", g)
}

func combSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var gap int = len(arr)
	var swapped bool = false
	var i int

	for gap > 1 || swapped == true {
		swapped = false
		if gap > 1 {
			gap = int(float64(gap) / 1.3)
		}
		for i = 0; i < len(arr)-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
			}
			//visualize(arr)
			g.Image = append(g.Image, buildImage(arr))
			g.Delay = append(g.Delay, 2)
		}
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("comb", g)
}

func countingSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	count := make([]int, max+1)
	for _, x := range arr {
		count[x-0]++
	}
	z := 0
	for i, c := range count {
		for ; c > 0; c-- {
			arr[z] = i
			z++
		}
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("counting", g)
}

func gnomeSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var i int = 1

	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("gnome", g)
}

func insertionSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j = j - 1
			//visualize(arr)
			g.Image = append(g.Image, buildImage(arr))
			g.Delay = append(g.Delay, 2)
		}
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("insertion", g)
}

func oddEvenSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var sorted bool = false
	var i int

	for !sorted {
		sorted = true
		for i = 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			//visualize(arr)
			g.Image = append(g.Image, buildImage(arr))
			g.Delay = append(g.Delay, 2)
		}
		for i = 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			//visualize(arr)
			g.Image = append(g.Image, buildImage(arr))
			g.Delay = append(g.Delay, 2)
		}
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("oddEven", g)
}

func selectionSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var min int = 0
	var i int
	var j int

	for i = 0; i < len(arr); i++ {
		min = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
				//visualize(arr)
				g.Image = append(g.Image, buildImage(arr))
				g.Delay = append(g.Delay, 2)
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		//visualize(arr)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("selection", g)
}

func sleepSort(arr []int) {
	g := &gif.GIF{
		LoopCount: 1,
	}
	var j int
	arr2 := make([]int, len(arr))
	channel := make(chan int, 1)
	//visualize(arr)
	g.Image = append(g.Image, buildImage(arr))
	g.Delay = append(g.Delay, 2)
	for i := 0; i < len(arr); i++ {
		go func(arr []int, i int) {
			time.Sleep(time.Duration(arr[i]) * time.Second / 4)
			channel <- arr[i]
		}(arr, i)
	}

	for i := 0; i < len(arr); i++ {
		arr2[j] = <-channel
		j++
		//visualize(arr2)
		g.Image = append(g.Image, buildImage(arr))
		g.Delay = append(g.Delay, 2)
	}
	writeGif("sleep", g)
}

func main() {
	var algo string
	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm all/bogo/[bubble]/comb/counting/gnome/insertion/oddEven/selection/sleep")
	flag.IntVar(&fps, "fps", 10, "frames per second")
	flag.IntVar(&max, "max", 9, "highest value")
	flag.IntVar(&count, "count", 30, "number of values")
	flag.IntVar(&mode, "mode", 1, "visualization mode")
	flag.Parse()
	arr := randomArray(count, max)
	fmt.Printf("sorting via %v-sort\nhighest value: %v\nnumber of values: %v\n\n", algo, max, count)
	time.Sleep(time.Second * 1)
	switch algo {
	case "bogo":
		bogoSort(arr)
	case "bubble":
		bubbleSort(arr)
	case "comb":
		combSort(arr)
	case "counting":
		countingSort(arr)
	case "gnome":
		gnomeSort(arr)
	case "insertion":
		insertionSort(arr)
	case "oddEven":
		oddEvenSort(arr)
	case "selection":
		selectionSort(arr)
	case "sleep":
		sleepSort(arr)
	case "all":
		arr = randomArray(count, max)
		bogoSort(arr)
		arr = randomArray(count, max)
		bubbleSort(arr)
		arr = randomArray(count, max)
		combSort(arr)
		arr = randomArray(count, max)
		countingSort(arr)
		arr = randomArray(count, max)
		gnomeSort(arr)
		arr = randomArray(count, max)
		insertionSort(arr)
		arr = randomArray(count, max)
		oddEvenSort(arr)
		arr = randomArray(count, max)
		selectionSort(arr)
		arr = randomArray(count, max)
		sleepSort(arr)
	}
}
