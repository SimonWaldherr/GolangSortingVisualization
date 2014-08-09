#!/bin/sh

echo "INSERT A FRAMERATE [30]"
read -t 10 FPS

echo "INPUT A NAME OF A SORTING ALGORITHM"
echo "[all]/bubble/comb/gnome/insertion/oddEven/selection"
read -t 30 ALGO

if [ "x$FPS" == "x" ]
  then
  FPS=30
fi

if [ "x$ALGO" == "x" ]
  then
  ALGO="all"
fi

go run gsv.go -count=$(tput cols) -max=$(tput lines) -fps=$FPS -algo=$ALGO
