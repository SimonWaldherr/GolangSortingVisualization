package gsv

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
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

// BogoSort is an implementation of https://en.wikipedia.org/wiki/Bogosort 
func BogoSort(arr []int, frameGen FrameGen) {
	if frameGen != nil {
		frameGen(arr)
	}
	for isSorted(arr) == false {
		arr = shuffle(arr)
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// BubbleSort is an implementation of https://en.wikipedia.org/wiki/Bubble_sort 
func BubbleSort(arr []int, frameGen FrameGen) {
	var i int
	var j int

	if frameGen != nil {
		frameGen(arr)
	}
	for i = 0; i < len(arr); i++ {
		for j = 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			if frameGen != nil {
				frameGen(arr)
			}
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// CocktailSort is an implementation of https://en.wikipedia.org/wiki/Cocktail_shaker_sort 
func CocktailSort(arr []int, frameGen FrameGen) {
	var i int
	if frameGen != nil {
		frameGen(arr)
	}
	for !isSorted(arr) {
		for i = 0; i < len(arr)-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
		for ; i > 0; i-- {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
	}
}

// CombSort is an implementation of https://en.wikipedia.org/wiki/Comb_sort 
func CombSort(arr []int, frameGen FrameGen) {
	var gap int = len(arr)
	var swapped bool = false
	var i int

	if frameGen != nil {
		frameGen(arr)
	}
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
			if frameGen != nil {
				frameGen(arr)
			}
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// CountingSort is an implementation of https://en.wikipedia.org/wiki/Counting_sort 
func CountingSort(arr []int, frameGen FrameGen) {
	count := make([]int, Max+1)

	if frameGen != nil {
		frameGen(arr)
	}
	for _, x := range arr {
		count[x-0]++
	}
	z := 0
	for i, c := range count {
		for ; c > 0; c-- {
			arr[z] = i
			z++
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// CycleSort is an implementation of https://en.wikipedia.org/wiki/Cycle_sort 
func CycleSort(arr []int, frameGen FrameGen) {
	if frameGen != nil {
		frameGen(arr)
	}
	for cycleStart, item := range arr {
		pos := cycleStart
		for _, item2 := range arr[cycleStart+1 : len(arr)] {
			if item2 < item {
				pos++
			}
		}
		if pos == cycleStart {
			continue
		}
		for item == arr[pos] {
			pos++
		}
		arr[pos], item = item, arr[pos]
		if frameGen != nil {
			frameGen(arr)
		}

		for pos != cycleStart {
			pos = cycleStart
			for _, item2 := range arr[cycleStart+1 : len(arr)] {
				if item2 < item {
					pos++
				}
			}
			for item == arr[pos] {
				pos++
			}
			arr[pos], item = item, arr[pos]
			if frameGen != nil {
				frameGen(arr)
			}
		}
	}
}

// GnomeSort is an implementation of https://en.wikipedia.org/wiki/Gnome_sort 
func GnomeSort(arr []int, frameGen FrameGen) {
	var i int = 1

	if frameGen != nil {
		frameGen(arr)
	}
	for i < len(arr) {
		if arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			if i > 1 {
				i--
			}
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// InsertionSort is an implementation of https://en.wikipedia.org/wiki/Insertion_sort 
func InsertionSort(arr []int, frameGen FrameGen) {
	var i int
	var j int

	if frameGen != nil {
		frameGen(arr)
	}
	for i = 0; i < len(arr); i++ {
		j = i
		for j > 0 && arr[j-1] > arr[j] {
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j = j - 1
			if frameGen != nil {
				frameGen(arr)
			}
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// OddEvenSort is an implementation of https://en.wikipedia.org/wiki/Odd–even_sort 
func OddEvenSort(arr []int, frameGen FrameGen) {
	var sorted bool = false
	var i int

	if frameGen != nil {
		frameGen(arr)
	}
	for !sorted {
		sorted = true
		for i = 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			if frameGen != nil {
				frameGen(arr)
			}
		}
		for i = 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			if frameGen != nil {
				frameGen(arr)
			}
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// SelectionSort is an implementation of https://en.wikipedia.org/wiki/Selection_sort 
func SelectionSort(arr []int, frameGen FrameGen) {
	var min int = 0
	var i int
	var j int

	if frameGen != nil {
		frameGen(arr)
	}
	for i = 0; i < len(arr); i++ {
		min = i
		for j = i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// SleepSort is an implementation of the SleepSort Algorithm - NOT ON WIKIPEDIA 
func SleepSort(arr []int, frameGen FrameGen) {
	var j int
	arr2 := make([]int, len(arr))
	channel := make(chan int, 1)
	if frameGen != nil {
		frameGen(arr)
	}
	for i := 0; i < len(arr); i++ {
		go func(arr []int, i int) {
			time.Sleep(time.Duration(arr[i]) * time.Second / 4)
			channel <- arr[i]
		}(arr, i)
	}

	for i := 0; i < len(arr); i++ {
		arr2[j] = <-channel
		j++
		if frameGen != nil {
			frameGen(arr2)
		}
	}
}

// StoogeSort is an implementation of https://en.wikipedia.org/wiki/Stooge_sort 
func StoogeSort(arr []int, frameGen FrameGen) {
	stoogesort(arr, 0, len(arr)-1, frameGen)
}

func stoogesort(arr []int, i int, j int, frameGen FrameGen) []int {
	var t int
	if frameGen != nil {
		frameGen(arr)
	}
	if arr[j] < arr[i] {
		arr[i], arr[j] = arr[j], arr[i]
	}
	if j-i+1 > 2 {
		t = (j - i + 1) / 3
		arr = stoogesort(arr, i, j-t, frameGen)
		arr = stoogesort(arr, i+t, j, frameGen)
		arr = stoogesort(arr, i, j-t, frameGen)
		if frameGen != nil {
			frameGen(arr)
		}
	}

	return arr
}

// QuickSort is an implementation of https://en.wikipedia.org/wiki/Quicksort
func QuickSort(arr []int, frameGen FrameGen) {
	if frameGen != nil {
		frameGen(arr)
	}
	quickSort(arr, 0, len(arr)-1, frameGen)
}

func quickSort(arr []int, l int, r int, frameGen FrameGen) {
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
	if frameGen != nil {
		frameGen(arr)
	}
}

// MergeSort is an implementation of https://en.wikipedia.org/wiki/Merge_sort
func MergeSort(arr []int, frameGen FrameGen) {
	// initial frame
	if frameGen != nil {
		frameGen(arr)
	}
	mergesort(arr, frameGen)
}

func mergesort(arr []int, frameGen FrameGen) []int {
	// base case
	if len(arr) <= 1 {
		return arr
	}

	// split the arr
	n := len(arr) / 2
	l, r := arr[:n], arr[n:]

	// sort the left and right
	l = mergesort(l, frameGen)
	r = mergesort(r, frameGen)

	return merge(l, r, frameGen)
}

func merge(l, r []int, frameGen FrameGen) []int {
	result := make([]int, 0)
	if frameGen != nil {
		frameGen(result)
	}

	for len(l) > 0 && len(r) > 0 {
		if l[0] <= r[0] {
			result = append(result, l[0])
			l = l[1:]
		} else {
			result = append(result, r[0])
			r = r[1:]
		}
		if frameGen != nil {
			frameGen(result)
		}
	}

	return append(append(result, l...), r...)
}

// ShellSort is an implementation of https://en.wikipedia.org/wiki/Shellsort
func ShellSort(arr []int, frameGen FrameGen) {
	n := len(arr)

	h := 1
	for h < n/3 {
		h = 3*h + 1
	}

	if frameGen != nil {
		frameGen(arr)
	}
	for h >= 1 {
		for i := h; i < n; i++ {
			for j := i; j >= h && arr[j] < arr[j-h]; j -= h {
				arr[j], arr[j-h] = arr[j-h], arr[j]
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
		h /= 3
	}
}

// HeapSort https://en.wikipedia.org/wiki/Heapsort
func HeapSort(arr []int, frameGen FrameGen) {
	heapsort(arr, len(arr), frameGen)
}

func heapsort(arr []int, c int, frameGen FrameGen) {
	heapify(arr, c, frameGen)

	end := c - 1
	for end > 0 {
		// move the largest value (arr[0]) to the front of the sorted values
		arr[end], arr[0] = arr[0], arr[end]

		// reduce heap size by one
		end--

		siftDown(arr, 0, end, frameGen)
	}
}

func parent(i int) int { return int(math.Floor((float64(i) - 1) / 2)) }
func child(i int) int  { return 2*i + 1 }

// heapify puts the elements of arr in heap order, in place
func heapify(arr []int, c int, frameGen FrameGen) {
	s := parent(c - 1)

	for s >= 0 {
		// sift down the node at index 'start' to the proper place
		// such that all nodes below the start index are in heap order
		siftDown(arr, s, c-1, frameGen)

		// goto the next parent node
		s--
	}
}

// siftDown repairs the heap whose root element is at index 's',
// assuming the heaps rooted at its children are valid
func siftDown(arr []int, start, end int, frameGen FrameGen) {
	root := start

	for child(root) <= end {
		child := child(root)
		swap := root

		if arr[swap] < arr[child] {
			swap = child
		}

		if child+1 <= end && arr[swap] < arr[child+1] {
			swap = child + 1
		}

		if swap == root {
			return
		}

		arr[root], arr[swap] = arr[swap], arr[root]
		root = swap

		if frameGen != nil {
			frameGen(arr)
		}
	}
}
