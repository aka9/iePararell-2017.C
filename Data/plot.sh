#!/bin/zsh

file='CmpareExecutionTime'

gnuplot<<PLOT
set xlabel "Prime Range"
set ylabel "Execution Time"
set title "Cmpare Execution Time"
set terminal svg
set output "CmpareExecutionTime.svg"

plot "NotParallel.dat" w lp,\
  "StaticParallel.dat" w lp,\
  “DaynamParallel.dat" w lp
PLOT

svg2pdf ${file}.svg ${file}.pdf
rm *.svg
