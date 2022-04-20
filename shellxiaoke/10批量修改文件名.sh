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





#选项：
#-n： 一般 sed 命令会把所有数据都输出到屏幕，如果加入此选择，则只会
#把经过 sed 命令处理的行输出到屏幕。
#-e： 允许对输入数据应用多条 sed 命令编辑。
#-f 脚本文件名： 从 sed 脚本中读入 sed 操作。和 awk 命令的-f 非常类似。
#-r： 在 sed 中支持扩展正则表达式。
#-i： 用 sed 的修改结果直接修改读取数据的文件，而不是由屏幕输出

#动作：
#a \： 追加，在当前行后添加一行或多行。添加多行时，除最后 一行外，
#每行末尾需要用“\”代表数据未完结。

#c \： 行替换，用 c 后面的字符串替换原数据行，替换多行时，除最后一行
#外，每行末尾需用“\”代表数据未完结。

#i \： 插入，在当期行前插入一行或多行。插入多行时，除最后 一行外，
#每行末尾需要用“\”代表数据未完结。

#d： 删除，删除指定的行。
#p： 打印，输出指定的行。

#s： 字串替换，用一个字符串替换另外一个字符串。格式为“行范围 s/
#旧字串/新字串/g”（和 vim 中的替换格式类似）。
