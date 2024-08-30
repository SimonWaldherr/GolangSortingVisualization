package main

import (
	cryptoRand "crypto/rand"
	"flag"
	"fmt"
	gsv "simonwaldherr.de/go/GolangSortingVisualization"
	"strings"
	"time"
)

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

func makeVisualizer(name string) gsv.Visualizer {
	name = "gif"
	switch name {
	case "stdout":
		//return &gsv.WriteStdout{}
	case "gif":
		return &gsv.GifVisualizer{}
	default:
		return nil
	}
	return nil
}

func runSort(visName string, algo string, sortFunc gsv.Sorter) {
	visualizer := makeVisualizer(visName)
	if visualizer == nil {
		fmt.Println("Invalid visualizer name")
		return
	}
	visualizer.Setup(algo)
	arr := randomArray(gsv.Count, gsv.Max)
	sortFunc(arr, visualizer.AddFrame)
	visualizer.Complete()
}

func keysString(m map[string]gsv.Sorter) string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return strings.Join(keys, "/")
}

func main() {
	var algo string
	var visName string

	sorterMap := map[string]gsv.Sorter{
		"bubble":    gsv.BubbleSort,
		"cocktail":  gsv.CocktailSort,
		"comb":      gsv.CombSort,
		"counting":  gsv.CountingSort,
		"cycle":     gsv.CycleSort,
		"gnome":     gsv.GnomeSort,
		"insertion": gsv.InsertionSort,
		"oddEven":   gsv.OddEvenSort,
		"selection": gsv.SelectionSort,
		"sleep":     gsv.SleepSort,
		"stooge":    gsv.StoogeSort,
		"pancake":   gsv.PancakeSort,
		"quick":     gsv.QuickSort,
		"merge":     gsv.MergeSort,
		"shell":     gsv.ShellSort,
		"heap":      gsv.HeapSort,
		"radix":     gsv.RadixSort,
		"bitonic":   gsv.BitonicSort,
	}

	flag.StringVar(&algo, "algo", "bubble", "Select sorting algorithm all/"+strings.Replace(keysString(sorterMap), "bubble", "[bubble]", 1))
	flag.IntVar(&gsv.Fps, "fps", 10, "frames per second")
	flag.IntVar(&gsv.Max, "max", 9, "highest value")
	flag.IntVar(&gsv.Count, "count", 30, "number of values")
	flag.IntVar(&gsv.Mode, "mode", 1, "visualization mode")
	flag.StringVar(&visName, "vis", "stdout", "Select output: [stdout]/gif")

	flag.Parse()

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
		} else {
			fmt.Printf("Algorithm %v not found.\n", algo)
		}
	}
}
