#!/bin/bash
HOST=$1
PORT="22 25 80 8080"
for PORT in $PORT; do
  # shellcheck disable=SC2261
  if echo &>/dev/null >/dev/tcp/$HOST/$PORT; then
    echo "$PORT open"
  else
    echo "$PORT close"
  fi
done
#用 shell 打印示例语句中字母数小于6的单词

#示例语句：
#Bash also interprets a number of multi-character options.
#!/bin/bash
##############################################################
#shell打印示例语句中字母数小于6的单词
##############################################################
for s in Bash also interprets a number of multi-character options.; do
  n=$(echo $s | wc -c)
  if [ $n -lt 6 ]; then
    echo $s
  fi
done
