#!/usr/bin/env bash

action=$1

if [[ "$action" == "setup" ]]; then
    git branch completed-sc-22
elif [[ "$action" == "teardown" ]]; then
    : # git branch -d completed-sc-22
else
    echo "USAGE: $0 <setup | teardown>"
    exit 1
fi
