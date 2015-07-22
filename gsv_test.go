package gsv

import "testing"

var visName string
var sorterMap map[string]Sorter

func init() {
	test = true

	sorterMap = map[string]Sorter{
		//"bogo":    BogoSort,
		"bubble":    BubbleSort,
		"comb":      CombSort,
		"counting":  CountingSort,
		"gnome":     GnomeSort,
		"insertion": InsertionSort,
		"oddEven":   OddEvenSort,
		"selection": SelectionSort,
		"sleep":     SleepSort,
		"stooge":    StoogeSort,
	}
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

func runSort(visName string, algo string, sortFunc Sorter) {
	visualizer := makeVisualizer(visName)
	visualizer.Setup(algo)
	arr := RandomArray(Count, Max)
	sortFunc(arr, visualizer.AddFrame)
	visualizer.Complete()
}

func Test_GIF(t *testing.T) {
	Max = 9
	Count = 30
	Mode = 1

	for k, v := range sorterMap {
		t.Log(k)
		runSort("gif", k, v)
	}

	t.Log("finish")
}

func Test_STDOUT(t *testing.T) {
	Max = 9
	Count = 30
	Mode = 1

	for k, v := range sorterMap {
		t.Log(k)
		runSort("stdout", k, v)
	}

	t.Log("finish")
}
