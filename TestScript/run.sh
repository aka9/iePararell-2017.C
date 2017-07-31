#!/bin/zsh

# Error
if [ $# -lt 3 ]
then
    echo 'Usage:' '[実行ファイル] [素数探索最大範囲] [刻み幅]'
    exit
fi

execFile=$1
max=$2                                 # 素数探索最大範囲
span=$3                                # 刻み幅
tmpFile='time.tmp'                     # 途中経過の出力先ファイル
prlt='~'                               # 標準出力の出力先ファイル

if [ -e $tmpFile ]
then
    rm -f $tmpFile $prlt
fi

# 実行ファイル名から，実行時間出力先ファイル名を取得
outFile=`echo $execFile | rev | sed -e 's/og//' | rev`

# 素数探索最大範囲まで計算
i=1;
if [ $# -eq 4 ]
then
    max=$2; split=$3; span=$4
    while [ $i -le $max ]
    do
        echo 'Exec:' $i
        go run $execFile $i $i 2>> $tmpFile >$prlt
        i=$((i+span))
    done
else
    while [ $i -le $max ]
    do
        echo 'Exec:' $i
        go run $execFile $i 2>> $tmpFile >$prlt
        i=$((i+span))
    done
fi

tmp1=`cat $tmpFile | cut -f3,4 -d" "`           # 一時ファイルから実行データを抽出
tmp2=`echo $tmp1 | tr 'max: ' ' ' | tr 's' ' '` # 実行データからmax: と単位記号を除去

echo $tmp2 > ${outFile}dat # 実行時間の書き出し
rm $tmpFile $prlt          # 一時ファイルと標準出力ファイルを削除
mv ${outFile}dat ../exp01
