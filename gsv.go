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

type Sorter func([]int, FrameGen)

type FrameGen func([]int)

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
	writeGif(gv.name, gv.g)
}

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

func writeStdout(arr []int) {
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

/* SORTING ALGORITHMS BEGIN HERE */

/* http://en.wikipedia.org/wiki/Bogosort */
func bogoSort(arr []int, frameGen FrameGen) {
	for isSorted(arr) == false {
		arr = shuffle(arr)
		frameGen(arr)
	}
}

/* http://en.wikipedia.org/wiki/Bubble_sort */
func bubbleSort(arr []int, frameGen FrameGen) {
	var i int
	var j int

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

/* http://en.wikipedia.org/wiki/Comb_sort */
func combSort(arr []int, frameGen FrameGen) {
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
			frameGen(arr)
		}
		frameGen(arr)
	}
}

/* http://en.wikipedia.org/wiki/Counting_sort */
func countingSort(arr []int, frameGen FrameGen) {
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
		frameGen(arr)
	}
}

/* http://en.wikipedia.org/wiki/Gnome_sort */
func gnomeSort(arr []int, frameGen FrameGen) {
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
		frameGen(arr)
	}
}

/* http://en.wikipedia.org/wiki/Insertion_sort */
func insertionSort(arr []int, frameGen FrameGen) {
	var i int
	var j int

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

/* http://en.wikipedia.org/wiki/Odd–even_sort */
func oddEvenSort(arr []int, frameGen FrameGen) {
	var sorted bool = false
	var i int

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

/* http://en.wikipedia.org/wiki/Selection_sort */
func selectionSort(arr []int, frameGen FrameGen) {
	var min int = 0
	var i int
	var j int

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
func sleepSort(arr []int, frameGen FrameGen) {
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

/* SORTING ALGORITHMS END HERE */

func makeVisualizer(name string) Visualizer {
	if name == "console" {
		return FrameGen(writeStdout)
	}
	if name == "gif" {
		return &GifVisualizer{}
	}
	return nil
}

func runSort(visName string, algo string, sortFunc Sorter) {
	visualizer := makeVisualizer(visName)
	visualizer.Setup(algo)
	arr := randomArray(count, max)
	sortFunc(arr, visualizer.AddFrame)
	visualizer.Complete()
}

func main() {
	var algo string
	var visName string
	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm all/bogo/[bubble]/comb/counting/gnome/insertion/oddEven/selection/sleep")
	flag.IntVar(&fps, "fps", 10, "frames per second")
	flag.IntVar(&max, "max", 9, "highest value")
	flag.IntVar(&count, "count", 30, "number of values")
	flag.IntVar(&mode, "mode", 1, "visualization mode")
	flag.StringVar(&visName, "vis", "console", "Select output: [console]/gif")

	flag.Parse()

	sorterMap := map[string]Sorter{
		//	"bogo":    bogoSort,
		"bubble":    bubbleSort,
		"comb":      combSort,
		"counting":  countingSort,
		"gnome":     gnomeSort,
		"insertion": insertionSort,
		"oddEven":   oddEvenSort,
		"selection": selectionSort,
		"sleep":     sleepSort,
	}

	fmt.Printf("sorting via %v-sort\nhighest value: %v\nnumber of values: %v\n\n", algo, max, count)
	time.Sleep(time.Second * 1)
	if algo == "all" {
		for k, v := range sorterMap {
			runSort(visName, k, v)
		}
	} else {
		sortFunc := sorterMap[algo]
		if sortFunc != nil {
			runSort(visName, algo, sortFunc)
		}
	}
}
