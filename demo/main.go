package main

import (
	"github.com/h4ckm03d/GolangSortingVisualization"
	cryptoRand "crypto/rand"
	"flag"
	"fmt"
	gsv "simonwaldherr.de/go/GolangSortingVisualization"
	"time"
)

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
	arr := randomArray(gsv.Count, gsv.Max)
	sortFunc(arr, visualizer.AddFrame)
	visualizer.Complete()
}

func main() {
	var algo string
	var visName string
	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm all/bogo/[bubble]/cocktail/comb/counting/gnome/insertion/oddEven/selection/sleep/quick")
	flag.IntVar(&gsv.Fps, "fps", 10, "frames per second")
	flag.IntVar(&gsv.Max, "max", 9, "highest value")
	flag.IntVar(&gsv.Count, "count", 30, "number of values")
	flag.IntVar(&gsv.Mode, "mode", 1, "visualization mode")
	flag.StringVar(&visName, "vis", "stdout", "Select output: [stdout]/gif")

	flag.Parse()

	sorterMap := map[string]gsv.Sorter{
		//"bogo":    gsv.BogoSort,
		"bubble":    gsv.BubbleSort,
		"cocktail":  gsv.CocktailSort,
		"comb":      gsv.CombSort,
		"counting":  gsv.CountingSort,
		"gnome":     gsv.GnomeSort,
		"insertion": gsv.InsertionSort,
		"oddEven":   gsv.OddEvenSort,
		"selection": gsv.SelectionSort,
		"sleep":     gsv.SleepSort,
		"stooge":    gsv.StoogeSort,
		"quick": 		 gsv.QuickSort,
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
