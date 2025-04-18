#!/usr/bin/env bash

set -u

go build -o e2e/sc .

cd e2e
eval "$(direnv export bash)"

# help
vhs help.tape --quiet
help_status=$?
cmp help.ascii golden/help.ascii
help_status=$?

# tasks
git checkout -b sample-story-with-tasks-sc-19
vhs tasks.tape --quiet
tasks_status=$?
git checkout - --quiet
git branch -d sample-story-with-tasks-sc-19
cmp tasks.ascii golden/tasks.ascii
tasks_status=$?

# # clean
git branch completed-sc-22
vhs clean.tape --quiet
clean_status=$?
# git branch -d completed-sc-22
cmp clean.ascii golden/clean.ascii
clean_status=$?

# mywork
git branch sample-story-with-tasks-sc-19
vhs mywork.tape --quiet
mywork_status=$?
git branch -d sample-story-with-tasks-sc-19
cmp mywork.ascii golden/mywork.ascii
mywork_status=$?


rm sc
rm *.ascii

exit $(($help_status + $tasks_status * 2 + $clean_status * 4 + $mywork_status * 8))
