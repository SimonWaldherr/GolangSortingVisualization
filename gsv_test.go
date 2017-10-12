package gsv

import (
	cryptoRand "crypto/rand"
	"testing"
)

var visName string
var sorterMap map[string]Sorter

func init() {
	test = true

	sorterMap = map[string]Sorter{
		"bogo":      BogoSort,
		"bubble":    BubbleSort,
		"cocktail":  CocktailSort,
		"comb":      CombSort,
		"counting":  CountingSort,
		"gnome":     GnomeSort,
		"insertion": InsertionSort,
		"oddEven":   OddEvenSort,
		"selection": SelectionSort,
		"sleep":     SleepSort,
		"stooge":    StoogeSort,
		"quick":     QuickSort,
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

func makeVisualizer(name string) Visualizer {
	if name == "gif" {
		return &GifVisualizer{}
	}
	if name == "stdout" {
		return FrameGen(WriteStdout)
	}
	return nil
}

func runSort(visName string, arr []int, algo string, sortFunc Sorter) {
	visualizer := makeVisualizer(visName)
	visualizer.Setup(algo)

	sortFunc(arr, visualizer.AddFrame)
	visualizer.Complete()
}

func Test_GIF(t *testing.T) {
	Max = 9
	Count = 9
	Mode = 2
	
	runSort("gif", randomArray(Count, Max), "selection", SelectionSort)

	Mode = 1

	for k, v := range sorterMap {
		t.Log(k)
		runSort("gif", randomArray(Count, Max), k, v)
	}

	t.Log("finish")
}

func Test_STDOUT(t *testing.T) {
	Max = 9
	Count = 9
	Mode = 1

	for k, v := range sorterMap {
		t.Log(k)
		runSort("stdout", randomArray(Count, Max), k, v)
	}

	t.Log("finish")
}
