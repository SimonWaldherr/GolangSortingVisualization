#!/bin/bash

go run demo/main.go -count=80 -max=80 -algo=all -vis=gif

for f in *.gif ; do 
  echo $f
  gifsicle --resize 320x320 -O --careful -d 5 -o sort_$f $f
  rm $f
done
