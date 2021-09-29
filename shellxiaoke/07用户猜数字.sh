#!/bin/bash

# 脚本生成一个 100 以内的随机数,提示用户猜数字,根据用户的输入,提示用户猜对了,
# 猜小了或猜大了,直至用户猜对脚本结束。
# RANDOM 为系统自带的系统变量,值为 0‐32767的随机数
# 使用取余算法将随机数变为 1‐100 的随机数num=$[RANDOM%100+1]echo "$num"
# 使用 read 提示用户猜数字
# 使用 if 判断用户猜数字的大小关系:‐eq(等于),‐ne(不等于),‐gt(大于),‐ge(大于等于),
# ‐lt(小于),‐le(小于等于)

while :
  do
    read -p "计算机生成了一个 1‐100 的随机数,你猜: " cai
    if [ $cai -eq $num ]
    then
        echo "恭喜,猜对了"
        exit
    elif [ $cai -gt $num ]
    then
        echo "Oops,猜大了"
    else
        echo "Oops,猜小了"
    fi
  done