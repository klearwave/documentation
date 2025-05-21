#!/usr/bin/env sh

set -e

# DESCRIPTION: performs the overlaying a set of manifests.  this is a simply script for now, but may
#              grow over time so we will keep logic out of the Makefile
# INPUTS:
#   CATEGORY: the category of components, from the parent directory in .source (e.g. certificates, ingress)
#   PROJECT: the project, within the CATEGORY, to overlay
#
# USAGE:
#   CATEGORY=secrets PROJECT=external-secrets scripts/overlay.sh

# retreive the common functions
source "$(dirname "$0")/common.sh"

# ensure inputs are provided
set_capability_component_vars

if [ -z `which yot` ]; then
    echo "yot not installed...please run 'make yot'..."
    exit 1
fi

if [ -z `which yq` ]; then
    echo "yq not installed...please manually install..."
    exit 1
fi

# ensure common values exist
if [ ! -f ${CAPABILITY_DIR_BASE}/values.yaml ]; then
    echo "missing common values file at ${CAPABILITY_DIR_BASE}/values.yaml..."
    exit 1
fi

# ensure capability values exist
if [ ! -f ${CAPABILITY_DIR_SOURCE}/values.yaml ]; then
    echo "missing capability-specific values file at ${CAPABILITY_DIR_SOURCE}/values.yaml..."
    exit 1
fi

# ensure component values exist
if [ ! -f ${COMPONENT_DIR}/config/values.yaml ]; then
    echo "missing component values file at ${COMPONENT_DIR}/config/values.yaml..."
    exit 1
fi

for OVERLAY in `ls ${COMPONENT_DIR}/config/overlays`; do \
    echo "using overlay ${OVERLAY} for ${COMPONENT}..."
    yot \
        --indent-level=2 \
        --instructions=${COMPONENT_DIR}/config/overlays/${OVERLAY} \
        --output-directory=. \
        --values-file=${CAPABILITY_DIR_BASE}/values.yaml \
        --values-file=${CAPABILITY_DIR_SOURCE}/values.yaml \
        --values-file=${COMPONENT_DIR}/config/values.yaml \
        --remove-comments \
        --stdout | yq -r eval-all 'select((. != null) or (type != "!!null"))' > ${COMPONENT_DIR}/manifests/${OVERLAY}
    # TODO: the yq logic removes blank yaml documents, but this should be handled in yot
    # itself and not part of a one-off script

    # TODO: remove duplication in overlays for common labels
    # run the stdout through our common overlays
    # yot \
    #    --path=- \
    #    --instructions=.source/overlay.yaml \
    #    --values-file=${COMPONENT_DIR}/config/values.yaml \
    #    --values-file=${CAPABILITY_DIR}/values.yaml \
    #    --remove-comments \
    #    --stdout > ${CAPABILITY}/${COMPONENT}/${OVERLAY}
done
