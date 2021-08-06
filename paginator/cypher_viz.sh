#!/usr/bin/zsh
textfile="./seqpi/$1txt.txt"

function access_cyphers () {
	if [ ! -r "$textfile" ]; then
		echo "Error: file not found."
		exit 1
	fi

	head -c $2 $textfile | grep --color=always $3

}

access_cyphers $1 $2 $3


