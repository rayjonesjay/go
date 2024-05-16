#!/bin/bash

for file in $(ls -d */)
do
		cd $file
		go test -v
		cd ..
done

