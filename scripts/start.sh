#!/bin/bash
echo 'Start the program: GradeManagement-0.0.1-SNAPSHOT.jar'
chmod 777 /home/arkrypto/app/GradeManagement-0.0.1-SNAPSHOT.jar
echo '-------Starting-------'
cd /home/arkrypto/app/
nohup java -jar GradeManagement-0.0.1-SNAPSHOT.jar &
echo 'start success'