#!/usr/bin/env bash

action=$1

if [[ "$action" == "setup" ]]; then
    git branch sample-story-with-tasks-sc-19
elif [[ "$action" == "teardown" ]]; then
    git branch -d sample-story-with-tasks-sc-19
else
    echo "USAGE: $0 <setup | teardown>"
    exit 1
fi
