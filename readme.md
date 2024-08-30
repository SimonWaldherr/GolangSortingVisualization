# GolangSortingVisualization

[![Go Report Card](https://goreportcard.com/badge/simonwaldherr.de/go/golangsortingvisualization)](https://goreportcard.com/report/simonwaldherr.de/go/golangsortingvisualization)
[![codebeat badge](https://codebeat.co/badges/c175babc-9113-40ab-8802-1cdb4b14d250)](https://codebeat.co/projects/github-com-simonwaldherr-golangsortingvisualization-master)

this sorting visualization is not intended to recommend any algorithm, if you need a recommendation go [somewhere else](https://en.wikipedia.org/wiki/Sorting_algorithm#Comparison_of_algorithms).  

if you like, feel free to add more Sorting Algorithm examples. Many thanks to all [contributors](https://github.com/SimonWaldherr/GolangSortingVisualization/graphs/contributors).

## Sorting Algorithms

### BogoSort

[![Bogo Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_bogo.gif)](https://en.wikipedia.org/wiki/Bogosort) 

### BubbleSort

[![Bubble Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_bubble.gif)](https://en.wikipedia.org/wiki/Bubble_sort) 

### CocktailSort

[![Cocktail Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_cocktail.gif)](https://en.wikipedia.org/wiki/Cocktail_shaker_sort) 

### CombSort

[![Comb Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_comb.gif)](https://en.wikipedia.org/wiki/Comb_sort) 

### CountingSort

[![Counting Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_counting.gif)](https://en.wikipedia.org/wiki/Counting_sort)

### CycleSort

[![Cycle Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_cycle.gif)](https://en.wikipedia.org/wiki/Cycle_sort)

### GnomeSort

[![Gnome Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_gnome.gif)](https://en.wikipedia.org/wiki/Gnome_sort)

### HeapSort

[![Heap Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_heap.gif)](https://en.wikipedia.org/wiki/Heapsort)

### InsertionSort

[![Insertion Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_insertion.gif)](https://en.wikipedia.org/wiki/Insertion_sort)

### MergeSort

[![Merge Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_merge.gif)](https://en.wikipedia.org/wiki/Merge_sort)

### OddEvenSort

[![OddEven Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_oddEven.gif)](https://en.wikipedia.org/wiki/Odd–even_sort)

### PancakeSort

[![Pancake Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_pancake.gif)](https://en.wikipedia.org/wiki/Pancake_sorting)

### QuickSort

[![Quick Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_quick.gif)](https://en.wikipedia.org/wiki/Quicksort)

### ShellSort

[![Shell Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_shell.gif)](https://en.wikipedia.org/wiki/Shellsort)

### SelectionSort

[![Selection Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_selection.gif)](https://en.wikipedia.org/wiki/Selection_sort)

### StoogeSort

[![Stooge Sort Animation](https://simonwaldherr.github.io/GolangSortingVisualization/sort_stooge.gif)](https://en.wikipedia.org/wiki/Stooge_sort)

## HowTo

```sh
./start.sh
```

```sh
$ go run gsv.go --help
Usage of gsv:
  -algo="bubble": Select sorting algorithm all/bogo/[bubble]/comb/counting/gnome/insertion/oddEven/selection/sleep
  -count=30: number of values
  -fps=10: frames per second
  -max=9: highest value
  -mode=1: visualization mode
  -vis="stdout": Select output: [stdout]/gif
```

## License

[MIT](https://github.com/SimonWaldherr/GolangSortingVisualization/blob/master/LICENSE)
