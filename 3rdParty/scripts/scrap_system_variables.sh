#!/bin/bash

set -e

DOCKERFILE="docker/dockerfile"

if [ ! -f "$DOCKERFILE" ]; then
	echo "File $DOCKERFILE not exist."
  	exit 1
fi

if [ -f .env ]; then
	rm -f .env
fi

grep -E '^ENV ' "$DOCKERFILE" | while read -r line; do

	env_line="${line#ENV }"

	while [[ "$env_line" =~ ([^= ]+)=([^= ]+) ]]; do
		name="${BASH_REMATCH[1]}"
		value="${BASH_REMATCH[2]}"

		if [ $name != "PATH" ]; then
			echo -e "$name=$value" >> .env
		fi

		env_line="${env_line#*$name=$value}"
		env_line="${env_line#" "}"
	done
done

export ZMIENNA="test"
