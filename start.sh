#!/bin/sh

echo "INSERT A FRAMERATE [10]"
read -t 10 FPS

echo "INPUT A NAME OF A SORTING ALGORITHM"
echo "[bubble]/gnome/selection"
read -t 30 ALGO

if [ "x$FPS" == "x" ]
  then
  FPS=10
fi

if [ "x$ALGO" == "x" ]
  then
  ALGO="bubble"
fi

go run gsv.go -count=$(tput cols) -max=$(tput lines) -fps=$FPS -algo=$ALGO
