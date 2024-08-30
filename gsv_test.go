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
		"cycle":     CycleSort,
		"gnome":     GnomeSort,
		"insertion": InsertionSort,
		"oddEven":   OddEvenSort,
		"selection": SelectionSort,
		"sleep":     SleepSort,
		"stooge":    StoogeSort,
		"pancake":   PancakeSort,
		"quick":     QuickSort,
		"merge":     MergeSort,
		"shell":     ShellSort,
		"heap":      HeapSort,
		"radix":     RadixSort,
		"bitonic":   BitonicSort,
	}
}

// StdoutVisualizer implements the Visualizer interface for stdout output
type StdoutVisualizer struct{}

func (sv *StdoutVisualizer) Setup(name string) {
	// No setup required for stdout
}

func (sv *StdoutVisualizer) AddFrame(arr []int) {
	WriteStdout(arr)
}

func (sv *StdoutVisualizer) Complete() {
	// No completion step required for stdout
}

func randomArray(n int, max int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		b := make([]byte, 1)
		cryptoRand.Read(b)
		number := float64(b[0])
		arr[i] = int(number / 255 * float64(max))
	}
	return arr
}

func makeVisualizer(name string) Visualizer {
	if name == "gif" {
		return &GifVisualizer{}
	}
	if name == "stdout" {
		return &StdoutVisualizer{}
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
func Benchmark_bogo_sort(b *testing.B)      { benchmarkSort("bogo", b) }
func Benchmark_bubble_sort(b *testing.B)    { benchmarkSort("bubble", b) }
func Benchmark_cocktail_sort(b *testing.B)  { benchmarkSort("cocktail", b) }
func Benchmark_comb_sort(b *testing.B)      { benchmarkSort("comb", b) }
func Benchmark_counting_sort(b *testing.B)  { benchmarkSort("counting", b) }
func Benchmark_cycle_sort(b *testing.B)     { benchmarkSort("cycle", b) }
func Benchmark_gnome_sort(b *testing.B)     { benchmarkSort("gnome", b) }
func Benchmark_insertion_sort(b *testing.B) { benchmarkSort("insertion", b) }
func Benchmark_oddEven_sort(b *testing.B)   { benchmarkSort("oddEven", b) }
func Benchmark_selection_sort(b *testing.B) { benchmarkSort("selection", b) }
func Benchmark_sleep_sort(b *testing.B)     { benchmarkSort("sleep", b) }
func Benchmark_stooge_sort(b *testing.B)    { benchmarkSort("stooge", b) }
func Benchmark_pancake_sort(b *testing.B)   { benchmarkSort("pancake", b) }
func Benchmark_quick_sort(b *testing.B)     { benchmarkSort("quick", b) }
func Benchmark_shell_sort(b *testing.B)     { benchmarkSort("shell", b) }
func Benchmark_heap_sort(b *testing.B)      { benchmarkSort("heap", b) }
func Benchmark_merge_sort(b *testing.B)     { benchmarkSort("merge", b) }
func Benchmark_radix_sort(b *testing.B)     { benchmarkSort("radix", b) }
func Benchmark_bitonic_sort(b *testing.B)   { benchmarkSort("bitonic", b) }

// WriteNop is a writer for FrameGen that does nothing.
// Ensures we only benchmark algorithms.
func WriteNop(_ []int) {}

func benchmarkSort(sort string, b *testing.B) {
	arr := randomArray(Count, Max)
	frameGen := FrameGen(WriteNop)
	if sortFunc, found := sorterMap[sort]; found {
		for n := 0; n < b.N; n++ {
			sortFunc(arr, frameGen)
		}
	}
}

// cloneArray Clones an array so source and the result are not backed by the same slice.
func cloneArray(source []int) []int {
	destination := make([]int, len(source))
	copy(destination, source)
	return destination
}

// TestCloneArray checks that cloneArray creates a separate copy and not a slice backed by the same array.
func TestCloneArray(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	d := cloneArray(s)

	if &s == &d {
		t.Error("Source and Destination address should not be equal")
	}
	for i := range s {
		if d[i] != s[i] {
			t.Errorf("Expected index [%d] to be the same", i)
		}
		if &d[i] == &s[i] {
			t.Errorf("Expected address of index [%d] to be different", i)
		}
	}
}

// BenchmarkConsistentArrayNoFramegen times the sort algorithms for the same array of random data
// for each algorithm without the overhead of Frame generation.
func BenchmarkConsistentArrayNoFramegen(b *testing.B) {
	arr := randomArray(1000, 750)
	for method, sortFn := range sorterMap {
		b.Run(method, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				arrCopy := cloneArray(arr)
				sortFn(arrCopy, nil)
			}
		})
	}
}
