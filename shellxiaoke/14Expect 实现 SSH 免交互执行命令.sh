#!/bin/bash
#Expect是一个自动交互式应用程序的工具，如telnet，ftp，passwd等。
#需先安装expect软件包。
#方法1：EOF标准输出作为expect标准输入

#USER=root
#PASS=123.com
#IP=192.168.1.120
#expect << EOFset timeout 30spawn ssh $USER@$IP   expect { "(yes/no)" {send "yes\r"; exp_continue}    "password:" {send "$PASS\r"}
#}
#expect "$USER@*"  {send "$1\r"}
#expect "$USER@*"  {send "exit\r"}
#expect eof
#EOF

#方法二
USER=root
PASS=123.com
IP=192.168.1.120
expect -c "
    spawn ssh $USER@$IP
    expect {
        \"(yes/no)\" {send \"yes\r\"; exp_continue}
        \"password:\" {send \"$PASS\r\"; exp_continue}
        \"$USER@*\" {send \"df -h\r exit\r\"; exp_continue}
    }"

#方法三
#登录脚本：
# cat login.exp
#!/usr/bin/expect
set ip [lindex $argv 0]
set user [lindex $argv 1]
set passwd [lindex $argv 2]
set cmd [lindex $argv 3]
if { $argc != 4 } {
puts "Usage: expect login.exp ip user passwd"
exit 1
}
set timeout 30
spawn ssh $user@$ip
expect {
    "(yes/no)" {send "yes\r"; exp_continue}
    "password:" {send "$passwd\r"}
}
expect "$user@*"  {send "$cmd\r"}
expect "$user@*"  {send "exit\r"}
expect eof




#执行命令脚本：写个循环可以批量操作多台服务器
HOST_INFO=user_info.txt
for ip in $(awk '{print $1}' $HOST_INFO)
do
    user=$(awk -v I="$ip" 'I==$1{print $2}' $HOST_INFO)
    pass=$(awk -v I="$ip" 'I==$1{print $3}' $HOST_INFO)
    expect login.exp $ip $user $pass $1
done
#Linux主机SSH连接信息：
# cat user_info.txt
192.168.1.120 root 123456
#创建10个用户，并分别设置密码，密码要求10位且包含大小写字母以及数字，最后需要把每个用户的密码存在指定文件中
#```bash
#!/bin/bash
##############################################################
#创建10个用户，并分别设置密码，密码要求10位且包含大小写字母以及数字
#最后需要把每个用户的密码存在指定文件中#前提条件：安装mkpasswd命令
##############################################################
#生成10个用户的序列（00-09）
for u in `seq -w 0 09`do
 #创建用户
 useradd user_$u
 #生成密码
 p=`mkpasswd -s 0 -l 10`
 #从标准输入中读取密码进行修改（不安全）
 echo $p|passwd --stdin user_$u
 #常规修改密码
 echo -e "$p\n$p"|passwd user_$u
 #将创建的用户及对应的密码记录到日志文件中
 echo "user_$u $p" >> /tmp/userpassworddone