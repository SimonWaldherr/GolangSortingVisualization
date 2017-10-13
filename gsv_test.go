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
// go test -bench=.
func Benchmark_bogo_sort(b *testing.B) { benchmarkSort("bogo", b)}
func Benchmark_bubble_sort(b *testing.B) { benchmarkSort("bubble", b)}
func Benchmark_cocktail_sort(b *testing.B) { benchmarkSort("cocktail", b)}
func Benchmark_comb_sort(b *testing.B) { benchmarkSort("comb", b)}
func Benchmark_counting_sort(b *testing.B) { benchmarkSort("counting", b)}
func Benchmark_gnome_sort(b *testing.B) { benchmarkSort("gnome", b)}
func Benchmark_insertion_sort(b *testing.B) { benchmarkSort("insertion", b)}
func Benchmark_oddEven_sort(b *testing.B) { benchmarkSort("oddEven", b)}
func Benchmark_sleep_sort(b *testing.B) { benchmarkSort("sleep", b)}
func Benchmark_stooge_sort(b *testing.B) { benchmarkSort("stooge", b)}
func Benchmark_quick_sort(b *testing.B) { benchmarkSort("quick", b)}

func benchmarkSort(sort string, b *testing.B){
	arr:=randomArray(Count, Max)
	frameGen:=FrameGen(WriteStdout)
	if sort, found := sorterMap[sort]; found {
		for n := 0; n < b.N; n++ {
			sort(arr, frameGen)
		}	
	} 
}