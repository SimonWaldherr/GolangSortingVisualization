package gsv

import "testing"

var visName string = "gif"

func makeVisualizer(name string) Visualizer {
	if name == "gif" {
		return &GifVisualizer{}
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

func Test_Main(t *testing.T) {
	Max = 9
	Count = 30
	Mode = 1

	sorterMap := map[string]Sorter{
		//"bogo":    BogoSort,
		"bubble":    BubbleSort,
		"comb":      CombSort,
		"counting":  CountingSort,
		"gnome":     GnomeSort,
		"insertion": InsertionSort,
		"oddEven":   OddEvenSort,
		"selection": SelectionSort,
		"sleep":     SleepSort,
	}

	for k, v := range sorterMap {
		t.Log(k)
		runSort(visName, k, v)
	}
	t.Log("finish")
}
