#!/usr/bin/env bash

GOEXEC=$(which go)

for dir in ./sets/*;
  do (
    cd "$dir"
    echo "*********"
    echo "starting ${PWD##*/}"
    echo "*********"
    for dir in ./challenge*
      do (
        cd $dir
        echo "\n----"
        echo "running ${PWD##*/}"
        exec $GOEXEC run challenge.go
      )
    done
  )
done

echo "******"
echo "finished"
echo "******"