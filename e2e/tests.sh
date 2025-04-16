#!/usr/bin/env bash

set -eu

go build -o e2e/sc .

cd e2e
eval "$(direnv export bash)"

# help
vhs help.tape --quiet
help_status=0
cmp help.ascii golden/help.ascii || echo "failed help"
cmp -s help.ascii golden/help.ascii || help_status=$?

# tasks
git checkout -b sample-story-with-tasks-sc-19
vhs tasks.tape --quiet
git checkout - --quiet
git branch -d sample-story-with-tasks-sc-19
tasks_status=0
cmp tasks.ascii golden/tasks.ascii || echo "failed tasks"
cmp -s tasks.ascii golden/tasks.ascii || tasks_status=$?

# clean
git branch completed-sc-22
vhs clean.tape --quiet
# git branch -d completed-sc-22
clean_status=0
cmp clean.ascii golden/clean.ascii || echo "failed clean"
cmp -s clean.ascii golden/clean.ascii || clean_status=$?

# mywork
git branch sample-story-with-tasks-sc-19
vhs mywork.tape --quiet
git branch -d sample-story-with-tasks-sc-19
mywork_status=0
cmp mywork.ascii golden/mywork.ascii || echo "failed mywork"
cmp -s mywork.ascii golden/mywork.ascii || mywork_status=$?


rm sc
rm *.ascii

exit $(($help_status + $tasks_status + $clean_status + $mywork_status))
