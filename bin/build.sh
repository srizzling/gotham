#!/bin/bash

set -o errexit
set -o pipefail
set -o nounset
IFS=$'\n'

# Build type (is build building a snapshot or a full release)
TYPE=$1

# Figure out version
VERSION=version

VERSION_FILE="VERSION.txt"

IGNORE=("base" "tests")


version() {
    local SERVICE=shift
    SERVICE=$(basename $SERVICE)
    local VERSION_LOCAL
    while read -r LINE; do
        echo "$LINE"
        # echo "$LINE"
        if  [[ $SERVICE == $LINE* ]]; then
             local VERSION_LOCAL=$LINE
             echo "$VERSION_LOCAL"
        fi
        echo "$VERSION_LOCAL"
        
    done < "$VERSION_FILE"

    # if [[ $TYPE == "dev" ]];
    # then
    #     VERSION="$VERSION-dev"
    # fi
    # echo "$VERSION"
}

function contains() {
    local n=$#
    local value=${!n}
    for ((i=1;i < $#;i++)) {
        if [ "${!i}" == "${value}" ]; then
            echo "y"
            return 0
        fi
    }
    echo "n"
    return 1
}

for f in services/*; do
    #printf "-------------------------------------------------------\n"
    f=$(basename "$f")
    if [ $(contains "${IGNORE[@]}" "$f") == "n" ]; then
        printf "Building %s service binary in docker container\n" "$f"
        version $f
        # VERSION=$(version "$f")
        # rocker build -var SERVICE="$f" -var VERSION="$VERSION"
        #printf "-------------------------------------------------------\n"
    fi
done