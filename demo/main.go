package main

import (
	"flag"
	"fmt"
	gsv "github.com/SimonWaldherr/GolangSortingVisualization"
	"time"
)

func makeVisualizer(name string) gsv.Visualizer {
	if name == "stdout" {
		return gsv.FrameGen(gsv.WriteStdout)
	}
	if name == "gif" {
		return &gsv.GifVisualizer{}
	}
	return nil
}

func runSort(visName string, algo string, sortFunc gsv.Sorter) {
	visualizer := makeVisualizer(visName)
	visualizer.Setup(algo)
	arr := gsv.RandomArray(gsv.Count, gsv.Max)
	sortFunc(arr, visualizer.AddFrame)
	visualizer.Complete()
}

func main() {
	var algo string
	var visName string
	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm all/bogo/[bubble]/comb/counting/gnome/insertion/oddEven/selection/sleep")
	flag.IntVar(&gsv.Fps, "fps", 10, "frames per second")
	flag.IntVar(&gsv.Max, "max", 9, "highest value")
	flag.IntVar(&gsv.Count, "count", 30, "number of values")
	flag.IntVar(&gsv.Mode, "mode", 1, "visualization mode")
	flag.StringVar(&visName, "vis", "stdout", "Select output: [stdout]/gif")

	flag.Parse()

	sorterMap := map[string]gsv.Sorter{
		//"bogo":    gsv.BogoSort,
		"bubble":    gsv.BubbleSort,
		"comb":      gsv.CombSort,
		"counting":  gsv.CountingSort,
		"gnome":     gsv.GnomeSort,
		"insertion": gsv.InsertionSort,
		"oddEven":   gsv.OddEvenSort,
		"selection": gsv.SelectionSort,
		"sleep":     gsv.SleepSort,
		"stooge":    gsv.StoogeSort,
	}

	fmt.Printf("sorting via %v-sort\nhighest value: %v\nnumber of values: %v\n\n", algo, gsv.Max, gsv.Count)
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
