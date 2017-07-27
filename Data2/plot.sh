#!/bin/zsh

file='Number-of-Divisions'

gnuplot<<PLOT
set xlabel "Prime Range"
set ylabel "Execution Time"
set title "Number of Divisions"
set terminal svg
set output "Number-of-Divisions.svg"

plot "DaynamParallel.dat" w lp
PLOT

svg2pdf ${file}.svg ${file}.pdf
rm *.svg
