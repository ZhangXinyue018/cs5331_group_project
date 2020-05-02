#!/bin/bash

echo Hello, which script you want to trigger?

# read script name
read script_name


if [ "$script_name" == "a" ]

then
	
# trigger attack script
	echo triggering attack script

	while true; do
		./attack2.sh 
		echo `date`
		echo ""
	done

else

# trigger solution script
	echo triggering solution script

	while true; do
		./solution2.sh 
		echo `date`
		echo ""
	done

fi
