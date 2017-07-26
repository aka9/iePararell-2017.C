#!/bin/zsh

# Error
if [ $# -lt 3 ]
then
    echo 'Usage:' '[実行ファイル] [素数探索最大範囲] [刻み幅]'
    exit
fi

execFile=$1      
max=$2             # 素数探索最大範囲
span=$3            # 刻み幅
tmpFile='time.tmp' # 途中経過の出力先ファイル
outFile='time.txt' # 平均実行時間の出力先ファイル
prlt='~'           # 標準出力の出力先ファイル

# 素数探索最大範囲まで計算
i=1;
while [ $i -le $max ]
do
    echo 'Exec:' $i
    go run $execFile $i 2>> $tmpFile >$prlt
    i=$((i+span))
done

tmp1=`cat $tmpFile | cut -f3,4 -d" "` # 一時ファイルから実行時間のみを抽出
tmp2=`echo $tmp1 | tr 'max: ' ' '`
tmp3=`echo $tmp2 | tr 's' ' '`

echo $tmp3 > $outFile # 実行時間の書き出し
rm $tmpFile $prlt     # 一時ファイルと標準出力ファイルを削除 
