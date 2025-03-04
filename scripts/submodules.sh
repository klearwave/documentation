#!/usr/bin/env bash

COMMIT="${COMMIT:-false}"

cd docs/submodules
. ./config

for REPO in $SUBMODULES; do
    DIR=$(echo $REPO | awk -F'/' '{print $NF}' | awk -F'.git' '{print $1}')
    if [[ ! -d ${DIR} ]]; then
        # add submodule if it does not exist
        git submodule add ${REPO} ${DIR}
        git add ${DIR}

        if [[ "${COMMIT}" == "true" ]]; then
            git commit -m "chore: added submodule ${DIR}"
        fi
    fi

    git submodule update --init --recursive
done

if [[ "${COMMIT}" == "true" ]]; then
    git commit -a -m "chore: updating submodules" 
fi
