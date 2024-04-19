#!/bin/bash

mkdir $1 && touch $1/main.go && cat template.txt > $1/main.go