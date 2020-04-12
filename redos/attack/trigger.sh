#!/bin/bash

echo Hello, which script you want to trigger?

# read script name
read script_name


if [ "$script_name" == "attack" ]

then
	
# trigger attack script
	echo triggering attack script

	while true; do
		./attack.sh 
		echo `date`
	done

else

# trigger solution script
	echo triggering solution script

	while true; do
		./solution.sh 
		echo `date`
	done

fi
