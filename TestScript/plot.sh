#!/bin/zsh

# Error
if [ $# -eq 0 ]
then
    echo 'Usage:' $0 '[素数探索最大値]'
    exit
fi

span=10            # 刻み幅
maxPrime=10000     # 素数最大値
tmpFile='time.tmp' # 途中経過の出力先ファイル
outFile='time.txt' # 平均実行時間の出力先ファイル
prlt='~'           # 標準出力の出力先ファイル

# 任意の回数プログラムを実行
i=1; max=$1
while [ $i -le $max ]
do
    echo 'Exec:' $i
    go run divisor.go $i 2>> $tmpFile >$prlt
    i=$((i+span))
done

tmp1=`cat $tmpFile | cut -f3,4 -d" "` # 一時ファイルから実行時間のみを抽出
#tmp2=`echo $tmp1 | cut -f1 -d"s"` # 実行時間から単位を削除

echo $tmp1 > $outFile # 実行時間の書き出し
rm $tmpFile $prlt     # 一時ファイルと標準出力ファイルを削除 
