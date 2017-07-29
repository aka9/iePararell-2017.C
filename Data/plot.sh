#!/bin/zsh

file='CompareExecutionTime'

gnuplot<<PLOT
set xlabel "Prime Range"
set ylabel "Execution Time (s)"
set title "Cmpare Execution Time"
set terminal svg
set output "CompareExecutionTime.svg"

plot "NotParallel.dat" w lp,\
  "static_parallel.dat" w lp,\
  "DaynamParallel.dat" w lp
PLOT

svg2pdf ${file}.svg ${file}.pdf
rm *.svg
