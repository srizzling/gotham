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


function version() {
    local SERVICE=$1
    SERVICE=$(basename "$SERVICE")
    local VERSION_LOCAL
    while read -r LINE; do
        case "$LINE" in 
            $SERVICE*)
                VERSION_LOCAL=$(echo "$LINE" | cut -d ' ' -f 2);;
            *)
                echo "";;
        esac      
    done < "$VERSION_FILE"

    if [[ $TYPE == "dev" ]];
    then
        VERSION_LOCAL="$VERSION_LOCAL-dev"
    fi

    echo "$VERSION_LOCAL"
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

function build() {
    for f in services/*; do
        echo "-------------------------------------------------------"
        f=$(basename "$f")
        if [ $(contains "${IGNORE[@]}" "$f") == "n" ]; then
            printf "Building %s service binary in docker container\n" "$f"
            VERSION=$(version "$f")
            f=${f//$'\n'/}
            VERSION=${VERSION//$'\n'/}
            rocker build -var "SERVICE=$f" -var "VERSION=$VERSION"
            echo "-------------------------------------------------------"
        fi
    done
}

build