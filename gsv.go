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

// Sorter defines a function type for sorting algorithms
type Sorter func([]int, FrameGen)

// FrameGen defines a function type for generating frames
type FrameGen func([]int)

//var test bool = false

func (fg FrameGen) Setup(name string) {
}

func (fg FrameGen) AddFrame(arr []int) {
	fg(arr)
}

func (fg FrameGen) Complete() {
}

// Visualizer interface for visualizing sorting steps
type Visualizer interface {
	Setup(string)
	AddFrame([]int)
	Complete()
}

// GifVisualizer is a visualizer that outputs a GIF
type GifVisualizer struct {
	name string
	g    *gif.GIF
}

var Max int
var Fps int
var Count int
var Mode int
var test bool = false

// Setup initializes the GIF visualizer
func (gv *GifVisualizer) Setup(name string) {
	gv.g = &gif.GIF{
		LoopCount: 1,
	}
	gv.name = name
}

// AddFrame adds a frame to the GIF
func (gv *GifVisualizer) AddFrame(arr []int) {
	frame := buildImage(arr)
	gv.g.Image = append(gv.g.Image, frame)
	gv.g.Delay = append(gv.g.Delay, 2)
}

// Complete writes the GIF to disk
func (gv *GifVisualizer) Complete() {
	WriteGif(gv.name, gv.g)
}

// buildImage creates an image from the array state
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

// WriteGif writes the GIF file to disk
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

// WriteStdout writes the array to stdout as an ASCII visualization
func WriteStdout(arr []int) {
	var buffer bytes.Buffer

	for y := 0; y < Max; y++ {
		for x := 0; x < len(arr); x++ {
			if arr[x] == y || (arr[x] < y && Mode == 1) || (arr[x] > y && Mode == 2) {
				buffer.WriteByte('#')
			} else {
				buffer.WriteByte(' ')
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

// shuffle randomizes the order of elements in the array
func shuffle(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		if j := rand.Intn(i + 1); i != j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	return arr
}

// isSorted checks if the array is sorted
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
	for !isSorted(arr) {
		arr = shuffle(arr)
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// BubbleSort is an implementation of https://en.wikipedia.org/wiki/Bubble_sort
func BubbleSort(arr []int, frameGen FrameGen) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
	}
}

// CocktailSort is an implementation of https://en.wikipedia.org/wiki/Cocktail_shaker_sort
func CocktailSort(arr []int, frameGen FrameGen) {
	for !isSorted(arr) {
		for i := 0; i < len(arr)-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
		for i := len(arr) - 2; i > 0; i-- {
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
	gap := len(arr)
	swapped := true

	for gap > 1 || swapped {
		swapped = false
		if gap > 1 {
			gap = int(float64(gap) / 1.3)
		}
		for i := 0; i < len(arr)-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				swapped = true
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
	}
}

// CountingSort is an implementation of https://en.wikipedia.org/wiki/Counting_sort
func CountingSort(arr []int, frameGen FrameGen) {
	count := make([]int, Max+1)
	for _, x := range arr {
		count[x]++
	}
	z := 0
	for i, c := range count {
		for c > 0 {
			arr[z] = i
			z++
			c--
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// CycleSort is an implementation of https://en.wikipedia.org/wiki/Cycle_sort
func CycleSort(arr []int, frameGen FrameGen) {
	for cycleStart := 0; cycleStart < len(arr)-1; cycleStart++ {
		item := arr[cycleStart]
		pos := cycleStart
		for i := cycleStart + 1; i < len(arr); i++ {
			if arr[i] < item {
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
		for pos != cycleStart {
			pos = cycleStart
			for i := cycleStart + 1; i < len(arr); i++ {
				if arr[i] < item {
					pos++
				}
			}
			for item == arr[pos] {
				pos++
			}
			arr[pos], item = item, arr[pos]
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// GnomeSort is an implementation of https://en.wikipedia.org/wiki/Gnome_sort
func GnomeSort(arr []int, frameGen FrameGen) {
	i := 0
	for i < len(arr) {
		if i == 0 || arr[i] >= arr[i-1] {
			i++
		} else {
			arr[i], arr[i-1] = arr[i-1], arr[i]
			i--
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// InsertionSort is an implementation of https://en.wikipedia.org/wiki/Insertion_sort
func InsertionSort(arr []int, frameGen FrameGen) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// OddEvenSort is an implementation of https://en.wikipedia.org/wiki/Odd–even_sort
func OddEvenSort(arr []int, frameGen FrameGen) {
	sorted := false
	for !sorted {
		sorted = true
		for i := 1; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			if frameGen != nil {
				frameGen(arr)
			}
		}
		for i := 0; i < len(arr)-1; i += 2 {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
			if frameGen != nil {
				frameGen(arr)
			}
		}
	}
}

// SelectionSort is an implementation of https://en.wikipedia.org/wiki/Selection_sort
func SelectionSort(arr []int, frameGen FrameGen) {
	for i := 0; i < len(arr); i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// SleepSort is a non-standard sorting algorithm
func SleepSort(arr []int, frameGen FrameGen) {
	arr2 := make([]int, len(arr))
	channel := make(chan int, 1)
	for i := 0; i < len(arr); i++ {
		go func(i int) {
			time.Sleep(time.Duration(arr[i]) * time.Millisecond)
			channel <- arr[i]
		}(i)
	}

	for i := 0; i < len(arr); i++ {
		arr2[i] = <-channel
		if frameGen != nil {
			frameGen(arr2)
		}
	}
	copy(arr, arr2)
}

// StoogeSort is an implementation of https://en.wikipedia.org/wiki/Stooge_sort
func StoogeSort(arr []int, frameGen FrameGen) {
	stoogesort(arr, 0, len(arr)-1, frameGen)
}

func stoogesort(arr []int, l, h int, frameGen FrameGen) {
	if arr[l] > arr[h] {
		arr[l], arr[h] = arr[h], arr[l]
		if frameGen != nil {
			frameGen(arr)
		}
	}
	if h-l+1 > 2 {
		t := (h - l + 1) / 3
		stoogesort(arr, l, h-t, frameGen)
		stoogesort(arr, l+t, h, frameGen)
		stoogesort(arr, l, h-t, frameGen)
	}
}

// PancakeSort is an implementation of https://en.wikipedia.org/wiki/Pancake_sorting
func PancakeSort(arr []int, frameGen FrameGen) {
	for uns := len(arr) - 1; uns > 0; uns-- {
		maxIndex := 0
		for i := 1; i <= uns; i++ {
			if arr[i] > arr[maxIndex] {
				maxIndex = i
			}
		}
		pancakeFlip(arr, maxIndex, frameGen)
		pancakeFlip(arr, uns, frameGen)
	}
}

func pancakeFlip(arr []int, r int, frameGen FrameGen) {
	for l := 0; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
		if frameGen != nil {
			frameGen(arr)
		}
	}
}

// QuickSort is an implementation of https://en.wikipedia.org/wiki/Quicksort
func QuickSort(arr []int, frameGen FrameGen) {
	quickSort(arr, 0, len(arr)-1, frameGen)
}

func quickSort(arr []int, l, r int, frameGen FrameGen) {
	if l >= r {
		return
	}
	pivot := partition(arr, l, r, frameGen)
	quickSort(arr, l, pivot-1, frameGen)
	quickSort(arr, pivot+1, r, frameGen)
}

func partition(arr []int, l, r int, frameGen FrameGen) int {
	pivot := arr[r]
	i := l
	for j := l; j < r; j++ {
		if arr[j] <= pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		if frameGen != nil {
			frameGen(arr)
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	if frameGen != nil {
		frameGen(arr)
	}
	return i
}

// MergeSort is an implementation of https://en.wikipedia.org/wiki/Merge_sort
func MergeSort(arr []int, frameGen FrameGen) {
	mergesort(arr, frameGen)
}

func mergesort(arr []int, frameGen FrameGen) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	left := mergesort(arr[:mid], frameGen)
	right := mergesort(arr[mid:], frameGen)
	return merge(left, right, frameGen)
}

func merge(left, right []int, frameGen FrameGen) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
		if frameGen != nil {
			frameGen(result)
		}
	}
	result = append(result, left...)
	result = append(result, right...)
	if frameGen != nil {
		frameGen(result)
	}
	return result
}

// ShellSort is an implementation of https://en.wikipedia.org/wiki/Shellsort
func ShellSort(arr []int, frameGen FrameGen) {
	n := len(arr)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i
			for ; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
				if frameGen != nil {
					frameGen(arr)
				}
			}
			arr[j] = temp
			if frameGen != nil {
				frameGen(arr)
			}
		}
	}
}

// HeapSort is an implementation of https://en.wikipedia.org/wiki/Heapsort
func HeapSort(arr []int, frameGen FrameGen) {
	buildMaxHeap(arr, frameGen)
	for i := len(arr) - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		maxHeapify(arr, 0, i, frameGen)
	}
}

func buildMaxHeap(arr []int, frameGen FrameGen) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		maxHeapify(arr, i, len(arr), frameGen)
	}
}

func maxHeapify(arr []int, i, n int, frameGen FrameGen) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		if frameGen != nil {
			frameGen(arr)
		}
		maxHeapify(arr, largest, n, frameGen)
	}
}

// RadixSort is an implementation of https://en.wikipedia.org/wiki/Radix_sort
func RadixSort(arr []int, frameGen FrameGen) {
	maxValue := getMax(arr)
	for exp := 1; maxValue/exp > 0; exp *= 10 {
		countingSortByDigit(arr, exp, frameGen)
	}
}

func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}

func countingSortByDigit(arr []int, exp int, frameGen FrameGen) {
	output := make([]int, len(arr))
	count := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		index := (arr[i] / exp) % 10
		count[index]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := len(arr) - 1; i >= 0; i-- {
		index := (arr[i] / exp) % 10
		output[count[index]-1] = arr[i]
		count[index]--
	}

	copy(arr, output)
	if frameGen != nil {
		frameGen(arr)
	}
}

// BitonicSort is an implementation of https://en.wikipedia.org/wiki/Bitonic_sorter
func BitonicSort(arr []int, frameGen FrameGen) {
	bitonicSort(arr, 0, len(arr), 1, frameGen)
}

func bitonicSort(arr []int, low, cnt, dir int, frameGen FrameGen) {
	if cnt > 1 {
		k := cnt / 2
		bitonicSort(arr, low, k, 1, frameGen)
		bitonicSort(arr, low+k, k, 0, frameGen)
		bitonicMerge(arr, low, cnt, dir, frameGen)
	}
}

func bitonicMerge(arr []int, low, cnt, dir int, frameGen FrameGen) {
	if cnt > 1 {
		k := cnt / 2
		for i := low; i < low+k; i++ {
			if (arr[i] > arr[i+k]) == (dir == 1) {
				arr[i], arr[i+k] = arr[i+k], arr[i]
				if frameGen != nil {
					frameGen(arr)
				}
			}
		}
		bitonicMerge(arr, low, k, dir, frameGen)
		bitonicMerge(arr, low+k, k, dir, frameGen)
	}
}
