#!/bin/zsh

tmpFile='time.tmp'
outFile='time.txt'
prlt='~'

i=1; max=$1
while [ $i -le $max ]
do
    echo $i
    go run divisor.go 2>> $tmpFile >$prlt
    i=$((i+1))
done

tmp1=`cat $tmpFile | cut -f3 -d" "`
tmp2=`echo $tmp1 | cut -f1 -d"µ"`
echo $tmp2 > $outFile
rm $tmpFile $prlt
