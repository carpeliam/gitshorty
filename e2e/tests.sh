#!/usr/bin/env bash

set -u

go build -o e2e/sc .

cd e2e
eval "$(direnv export bash)"


function test() {
    local type=$1
    if [ -f setup/"${type}".sh ]; then
        ./setup/"${type}".sh setup
    fi
    vhs "${type}.tape" --quiet
    local status=$?
    cmp --print-bytes "${type}".ascii golden/"${type}".ascii
    local status=$?
    if [ -f setup/"${type}".sh ]; then
        ./setup/"${type}".sh teardown
    fi
    if [ "$status" -eq 0 ]; then
        rm "${type}".ascii
    fi
    return $status
}

status=0
if [ $# -eq 1 ]; then
    tape=$1
    echo "testing tape $tape"
    test "$tape"
    status=$?
else
    echo "testing all tapes"
    # tapes=(*.tape) # this is currently flaky :(
    tapes=(help.tape tasks.tape clean.tape mywork.tape)
    for i in "${!tapes[@]}"; do
        tape="${tapes[$i]%.*}"
        test "${tape}"
        tape_status=$?
        status=$((status + tape_status * 2**i))
    done
fi

rm sc

exit "$status"
