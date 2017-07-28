#!/bin/zsh

execFile='DaynamParallel.go'
outFile='DaynamParallel.dat'
max=100000                                 # 素数探索最大範囲
span=1000                               # 刻み幅
tmpFile='time.tmp'                     # 途中経過の出力先ファイル
prlt='~'                               # 標準出力の出力先ファイル

if [ -e $tmpFile ]
then
    rm -f $tmpFile $prlt
fi

# 素数探索最大範囲まで計算
i=1000;
while [ $i -le $max ]
do
    echo 'Exec:' $i
    go run $execFile $max $i 2>> $tmpFile >$prlt
    #    i=$((i*span))
    i=$((i+span))
done

tmp1=`cat $tmpFile | cut -f3,4 -d" "`           # 一時ファイルから実行データを抽出
tmp2=`echo $tmp1 | tr 'divided: ' ' ' | tr 's' ' '` # 実行データからmax: と単位記号を除去

echo $tmp2 > ${outFile}  # 実行時間の書き出し
rm $tmpFile $prlt        # 一時ファイルと標準出力ファイルを削除
