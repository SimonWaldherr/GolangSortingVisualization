#GolangSortingVisualization

[![Gittip donate button](http://img.shields.io/gittip/bevry.png)](https://www.gittip.com/SimonWaldherr/ "Donate weekly to this project using Gittip") [![Flattr donate button](https://raw.github.com/balupton/flattr-buttons/master/badge-89x18.gif)](https://flattr.com/submit/auto?user_id=SimonWaldherr&url=http%3A%2F%2Fgithub.com%2FSimonWaldherr%2FGolangSortingVisualization "Donate monthly to this project using Flattr")


##Sort Algorithms

###BogoSort

[BogoSort sorts by shuffling until all is sorted](http://en.wikipedia.org/wiki/Bogosort)

###BubbleSort

![Bubble Sort Animation](http://upload.wikimedia.org/wikipedia/commons/3/37/Bubble_sort_animation.gif)

###CombSort

![Comb Sort Animation](http://upload.wikimedia.org/wikipedia/commons/4/46/Comb_sort_demo.gif)

###[CountingSort](http://en.wikipedia.org/wiki/Counting_sort)

###GnomeSort

![Gnome Sort Animation](http://upload.wikimedia.org/wikipedia/commons/3/37/Sorting_gnomesort_anim.gif)

###InsertionSort

![Insertion Sort Animation](http://upload.wikimedia.org/wikipedia/commons/4/42/Insertion_sort.gif)

###OddEvenSort

![OddEven Sort Animation](http://upload.wikimedia.org/wikipedia/commons/1/1b/Odd_even_sort_animation.gif)

###SelectionSort

![Selection Sort Animation](http://upload.wikimedia.org/wikipedia/commons/thumb/b/b0/Selection_sort_animation.gif/250px-Selection_sort_animation.gif)

###SleepSort

##HowTo

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
```

##License

MIT