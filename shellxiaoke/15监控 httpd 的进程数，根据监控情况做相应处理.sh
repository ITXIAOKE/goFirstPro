#!/bin/bash
###############################################################################################################################
#需求：
#1.每隔10s监控httpd的进程数，若进程数大于等于500，则自动重启Apache服务，并检测服务是否重启成功
#2.若未成功则需要再次启动，若重启5次依旧没有成功，则向管理员发送告警邮件，并退出检测
#3.如果启动成功，则等待1分钟后再次检测httpd进程数，若进程数正常，则恢复正常检测（10s一次），否则放弃重启并向管理员发送告警邮件，并退出检测
###############################################################################################################################
#计数器函数
#seq命令可以输出连续的数字

check_service() {
  j=0
  for i in $(seq 1 5); do
    #重启Apache的命令
    /usr/local/apache2/bin/apachectl restart 2>/var/log/httpderr.log
    #判断服务是否重启成功
    # shellcheck disable=SC1073
    if [ $? -eq 0 ]; then
      break
    else
      j=$(($j + 1))
    fi
    #判断服务是否已尝试重启5次
    if [ $j -eq 5 ]; then
      mail.py exit
    fi
  done
}

while :; do
  n=$(pgrep -l httpd | wc -l)
  #判断httpd服务进程数是否超过500
  if [ $n -gt 500 ]; then
    /usr/local/apache2/bin/apachectl restart
    if [ $? -ne 0 ]; then
      check_service
    else
      sleep 60
      n2=$(pgrep -l httpd | wc -l)
      #判断重启后是否依旧超过500
      if [ $n2 -gt 500 ]; then
        mail.py exit
      fi
    fi
  fi
  #每隔10s检测一次
  sleep 10
done

#$!：：Shell最后运行的后台Process的PID

#$-：：使用Set命令设定的Flag一览

#$*：：所有参数列表。如"$*"用「"」括起来的情况、以"$1 $2 … $n"的形式输出所有参数。只有单个字符串
#$@：：所有参数列表。如"$@"用「"」括起来的情况、以"$1" "$2" … "$n" 的形式输出所有参数。每个都是一个字符串

#$# 是传给脚本的参数个数
#$0 是脚本本身的名字
#$1 是传递给该shell脚本的第一个参数
#$2 是传递给该shell脚本的第二个参数

#$$ 是脚本运行的当前进程ID号
#$? 是显示最后命令的退出状态，0表示没有错误，其他表示有错误
