#!/usr/bin/env bash

action=$1

if [[ "$action" == "setup" ]]; then
    git checkout -b sample-story-with-tasks-sc-19
elif [[ "$action" == "teardown" ]]; then
    git checkout - --quiet
    git branch -d sample-story-with-tasks-sc-19
else
    echo "USAGE: $0 <setup | teardown>"
    exit 1
fi
