#!/bin/bash
# touch article_{1..3}.html
# lsarticle_1.html  article_2.html  article_3.html
#目的：把article改为bbs
# 方法一
# shellcheck disable=SC2045
for file in $(ls *html); do
  mv $file bbs_${file#*_}
  # mv $file $(echo $file |sed -r 's/.*(_.*)/bbs\1/')
  # mv $file $(echo $file |echo bbs_$(cut -d_ -f2)
done

#方法二
for file in $(find . -maxdepth 1 -name "*html"); do
  mv $file bbs_${file#*_}
done

#方法三
# rename article bbs *.html
#把一个文档前五行中包含字母的行删掉，同时删除6到10行包含的所有字母
#
#1）准备测试文件，文件名为2.txt
#
#第1行1234567不包含字母
#第2行56789BBBBBB
#第3行67890CCCCCCCC
#第4行78asdfDDDDDDDDD
#第5行123456EEEEEEEE
#第6行1234567ASDF
#第7行56789ASDF
#第8行67890ASDF
#第9行78asdfADSF
#第10行123456AAAA
#第11行67890ASDF
#第12行78asdfADSF
#第13行123456AAAA
###############################################################
#把一个文档前五行中包含字母的行删掉，同时删除6到10行包含的所有字母
##############################################################
sed -n '1,5'p 2.txt | sed '/[a-zA-Z]/'d
sed -n '6,10'p 2.txt | sed s'/[a-zA-Z]//'g
sed -n '11,$'p 2.txt
#最终结果只是在屏幕上打印结果，如果想直接更改文件，可将输出结果写入临时文件中，再替换2.txt或者使用-i选项
