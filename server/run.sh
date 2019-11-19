#!/bin/bash
###
# @Author: Sy.
 # @Create: 2019-11-01 20:54:15
 # @LastTime: 2019-11-16 17:07:19
 # @LastEdit: Sy.
 # @FilePath: \server\run.sh
 # @Description: linux
 ###

case $1 in
	start)
		nohup ./server_admin 2>&1 >> info.log 2>&1 /dev/null &
		echo "服务已启动..."
		sleep 1
	;;
	stop)
		killall server_admin
		echo "服务已停止..."
		sleep 1
	;;
	restart)
		killall server_admin
		sleep 1
		nohup ./server_admin 2>&1 >> info.log 2>&1 /dev/null &
		echo "服务已重启..."
		sleep 1
	;;
	*)
		echo "$0 {start|stop|restart}"
		exit 4
	;;
esac
