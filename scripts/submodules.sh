#!/usr/bin/env bash

COMMIT="${COMMIT:-false}"
SUBMODULES_DIR='docs/submodules'

mkdir -p ${SUBMODULES_DIR}

# NOTE: assumes this is run from the Make target at the root of the repo
SUBMODULES=$(grep 'url =' .gitmodules | awk '{print $NF}')
cd ${SUBMODULES_DIR}

for REPO in $SUBMODULES; do
    DIR=$(echo $REPO | awk -F'/' '{print $NF}' | awk -F'.git' '{print $1}')
    echo "ensuring submodule for ${REPO} exists in ${DIR}..."
    if [[ ! -d ${DIR} ]]; then
        # add submodule if it does not exist
        echo "adding submodule ${REPO} to ${DIR}..."
        git submodule add --force ${REPO} ${DIR}
        git add ${DIR}

        if [[ "${COMMIT}" == "true" ]]; then
            git commit -m "chore: added submodule ${DIR}"
        fi
    fi

    echo "updating submodule for ${REPO} in ${DIR}..."
    git submodule update --init --recursive
    pushd $DIR
    git pull
    popd
done

if [[ "${COMMIT}" == "true" ]]; then
    git commit -a -m "chore: updating submodules" 
fi
