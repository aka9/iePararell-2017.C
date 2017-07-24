#!/bin/zsh

# Error
if [ $# -ne 2 ]
then
    echo 'Usage:' $0 '[実行ファイル名] [実行回数]'
    exit
fi

execFile=$1        # 実行ファイル名
tmpFile='time.tmp' # 途中経過の出力先ファイル
outFile='time.txt' # 平均実行時間の出力先ファイル
prlt='~'           # 標準出力の出力先ファイル

# 任意の回数プログラムを実行
i=1; max=$2
while [ $i -le $max ]
do
    echo 'Exec:' $i
    go run $execFile 2>> $tmpFile >$prlt
    i=$((i+1))
done

tmp1=`cat $tmpFile | cut -f3 -d" "` # 一時ファイルから実行時間のみを抽出
tmp2=(`echo $tmp1 | cut -f1 -d"s"`) # 実行時間から単位を削除

# 実行時間の合計値を計算
sum=$tmp2[1]; i=2; max=${#tmp2[*]}
while [ $i -le $max ]
do
    sum=$((sum+tmp2[i]))
    i=$((i+1))
done

ave=$((sum/max))     # 平均実行時間を計算
echo $ave > $outFile # 平均実行時間の書き出し
rm $tmpFile $prlt    # 一時ファイルと標準出力ファイルを削除
