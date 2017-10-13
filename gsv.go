package gsv

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math/rand"
	"os"
	"time"
)

type Sorter func([]int, FrameGen)

type FrameGen func([]int)

var test bool = false

func (fg FrameGen) Setup(name string) {
}

func (fg FrameGen) AddFrame(arr []int) {
	fg(arr)
}

func (fg FrameGen) Complete() {
}

type Visualizer interface {
	Setup(string)
	AddFrame([]int)
	Complete()
}

type GifVisualizer struct {
	name string
	g    *gif.GIF
}

func (gv *GifVisualizer) Setup(name string) {
	gv.g = &gif.GIF{
		LoopCount: 1,
	}
	gv.name = name
}

func (gv *GifVisualizer) AddFrame(arr []int) {
	frame := buildImage(arr)
	gv.g.Image = append(gv.g.Image, frame)
	gv.g.Delay = append(gv.g.Delay, 2)
}

func (gv *GifVisualizer) Complete() {
	WriteGif(gv.name, gv.g)
}

var Max int
var Fps int
var Count int
var Mode int

func buildImage(arr []int) *image.Paletted {
	var frame = image.NewPaletted(
		image.Rectangle{
			image.Point{0, 0},
			image.Point{len(arr), Max},
		},
		color.Palette{
			color.Gray{uint8(255)},
			color.Gray{uint8(0)},
		},
	)
	for k, v := range arr {
		frame.SetColorIndex(k, Max-v, uint8(1))
		if Mode == 2 {
			for y := Max - v + 1; y < Max; y++ {
				frame.SetColorIndex(k, y, uint8(1))
			}
		}
	}
	return frame
}

func WriteGif(name string, g *gif.GIF) {
	w, err := os.Create(name + ".gif")
	if err != nil {
		fmt.Println("os.Create")
		panic(err)
	}
	defer func() {
		if err := w.Close(); err != nil {
			fmt.Println("w.Close")
			panic(err)
		}
	}()
	err = gif.EncodeAll(w, g)
	if err != nil {
		fmt.Println("gif.EncodeAll")
		panic(err)
	}
}

func WriteStdout(arr []int) {
	var buffer bytes.Buffer
	var x int
	var y int

	for y = 0; y < Max; y++ {
		for x = 0; x < len(arr); x++ {
			if arr[x] == y {
				buffer.WriteByte(byte('#'))
			} else if arr[x] < y && Mode == 1 {
				buffer.WriteByte(byte('#'))
			} else if arr[x] > y && Mode == 2 {
				buffer.WriteByte(byte('#'))
			} else {
				buffer.WriteByte(byte(' '))
			}
		}
		buffer.WriteByte('\n')
	}

	if !test {
		time.Sleep(time.Second / time.Duration(Fps))
		fmt.Print("\033[2J")
		fmt.Print(buffer.String())
	}
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
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

/* SORTING ALGORITHMS BEGIN HERE */

/* https://en.wikipedia.org/wiki/Bogosort */
func BogoSort(arr []int, frameGen FrameGen) {
	frameGen(arr)
	for isSorted(arr) == false {
		arr = shuffle(arr)
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Bubble_sort */
func BubbleSort(arr []int, frameGen FrameGen) {
	var i int
	var j int

	frameGen(arr)
	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			frameGen(arr)
		}
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Cocktail_shaker_sort */
func CocktailSort(arr []int, frameGen FrameGen) {
	var i int
	frameGen(arr)
	for !isSorted(arr) {
		for i = 0; i < len(arr)-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				frameGen(arr)
			}
		}
		for ; i > 0; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				frameGen(arr)
			}
		}
	}
}

/* https://en.wikipedia.org/wiki/Comb_sort */
func CombSort(arr []int, frameGen FrameGen) {
	var gap int = len(arr)
	var swapped bool = false
	var i int

	frameGen(arr)
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
			frameGen(arr)
		}
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Counting_sort */
func CountingSort(arr []int, frameGen FrameGen) {
	count := make([]int, Max+1)

	frameGen(arr)
	for _, x := range arr {
		count[x-0]++
	}
	z := 0
	for i, c := range count {
		for ; c > 0; c-- {
			arr[z] = i
			z++
		}
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Gnome_sort */
func GnomeSort(arr []int, frameGen FrameGen) {
	var i int = 1

	frameGen(arr)
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Insertion_sort */
func InsertionSort(arr []int, frameGen FrameGen) {
	var i int
	var j int

	frameGen(arr)
	for i = 0; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j = j - 1
			frameGen(arr)
		}
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Odd–even_sort */
func OddEvenSort(arr []int, frameGen FrameGen) {
	var sorted bool = false
	var i int

	frameGen(arr)
	for !sorted {
		sorted = true
		for i = 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			frameGen(arr)
		}
		for i = 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			frameGen(arr)
		}
		frameGen(arr)
	}
}

/* https://en.wikipedia.org/wiki/Selection_sort */
func SelectionSort(arr []int, frameGen FrameGen) {
	var min int = 0
	var i int
	var j int

	frameGen(arr)
	for i = 0; i < len(arr); i++ {
		min = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
				frameGen(arr)
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		frameGen(arr)
	}
}

/* NOT ON WIKIPEDIA */
func SleepSort(arr []int, frameGen FrameGen) {
	var j int
	arr2 := make([]int, len(arr))
	channel := make(chan int, 1)
	frameGen(arr)
	for i := 0; i < len(arr); i++ {
		go func(arr []int, i int) {
			time.Sleep(time.Duration(arr[i]) * time.Second / 4)
			channel <- arr[i]
		}(arr, i)
	}

	for i := 0; i < len(arr); i++ {
		arr2[j] = <-channel
		j++
		frameGen(arr2)
	}
}

/* https://en.wikipedia.org/wiki/Stooge_sort */
func StoogeSort(arr []int, frameGen FrameGen) {
	stoogesort(arr, 0, len(arr)-1, frameGen)
}

func stoogesort(arr []int, i int, j int, frameGen FrameGen) []int {
	var t int
	frameGen(arr)
	if arr[j] < arr[i] {
		arr[i], arr[j] = arr[j], arr[i]
	}
	if j-i+1 > 2 {
		t = (j - i + 1) / 3
		arr = stoogesort(arr, i, j-t, frameGen)
		arr = stoogesort(arr, i+t, j, frameGen)
		arr = stoogesort(arr, i, j-t, frameGen)
		frameGen(arr)
	}

	return arr
}

// QuickSort https://en.wikipedia.org/wiki/Quicksort
func QuickSort(arr []int, frameGen FrameGen) {
	frameGen(arr)
	quickSort(arr, 0, len(arr)-1, frameGen)
}

func quickSort(arr []int, l int, r int, frameGen FrameGen){
	if l >= r {
			return 
	}
	
	pivot := arr[l]
	i := l + 1

	for j := l; j <= r; j++ {
			if pivot > arr[j] {
					arr[i], arr[j] = arr[j], arr[i]
					i++
			}
	}

	arr[l], arr[i-1] = arr[i-1], pivot

	quickSort(arr, l, i-2, frameGen)
	quickSort(arr, i, r, frameGen)
	frameGen(arr)	
}