#!/bin/bash
echo "Stop Procedure : demo2-0.0.1-SNAPSHOT.jar"
pid=`ps -ef |grep java|grep GradeManagement-0.0.1-SNAPSHOT.jar|awk '{print $2}'`
echo 'old Procedure pid:'$pid
if [ -n "$pid" ]
then
kill -9 $pid
fi