#!/bin/bash
FILESCOUNTS=$(find . -type f,d | wc -l)
NUMBER=5
TOTALFILES=$(($FILESCOUNTS * $NUMBER))
printf "\t\vTotal files * $NUMBER: $TOTALFILES\v\n"