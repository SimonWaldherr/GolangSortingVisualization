#!/bin/sh

echo "INSERT A FRAMERATE [30]"
read -t 10 FPS

echo "INPUT A NAME OF A SORTING ALGORITHM"
echo "[all]/bogo/bubble/comb/counting/gnome/insertion/oddEven/selection/sleep"
read -t 30 ALGO

echo "SELECT OUTPUT MODE [stdout]/gif"
read -t 15 OUTPUT

if [ "x$FPS" == "x" ]
  then
  FPS=30
fi

if [ "x$ALGO" == "x" ]
  then
  ALGO="all"
fi

if [ "x$OUTPUT" == "x" ]
  then
  OUTPUT="stdout"
fi

go run demo/main.go -count=$(tput cols) -max=$(tput lines) -fps=$FPS -algo=$ALGO -vis=$OUTPUT
