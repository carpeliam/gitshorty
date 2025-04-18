#!/usr/bin/env bash

set -u

go build -o e2e/sc .

cd e2e
eval "$(direnv export bash)"


function test() {
    type=$1
    vhs "${type}.tape" --quiet
    status=$?
    cmp --print-bytes "${type}".ascii golden/"${type}".ascii
    status=$?
    return $status
}

# help
test "help"
help_status=$?

# tasks
git checkout -b sample-story-with-tasks-sc-19
test "tasks"
tasks_status=$?
git checkout - --quiet
git branch -d sample-story-with-tasks-sc-19

# clean
git branch completed-sc-22
test "clean"
clean_status=$?
# git branch -d completed-sc-22

# mywork
git branch sample-story-with-tasks-sc-19
test "mywork"
mywork_status=$?
git branch -d sample-story-with-tasks-sc-19

rm sc
rm *.ascii

exit $(($help_status + $tasks_status * 2 + $clean_status * 4 + $mywork_status * 8))
