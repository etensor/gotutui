#!/bin/zsh
filepi="./seqpi/pitxt.txt"

function access_pi () {
    if [ ! -r "$filepi" ]; then
        echo "Error: file not found."
        exit 1
    fi

    head -c $1 $filepi | grep --color=always $2
    #grep --color=always $number <(head -c $1 $filepi)
}
access_pi $1 $2

