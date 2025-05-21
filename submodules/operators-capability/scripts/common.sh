# defines usage when both a capability and component are expected
capability_component_usage() {
    echo "Usage: $0 <capability> <component>"
    echo "  capability: the capability name (directory under capabilities/)"
    echo "  component:  the component name (the project name, e.g. cert-manager)"
}

# defines function to ensure both capability and component are provided
ensure_capability_component_input() {
    : ${CAPABILITY?missing environment variable CAPABILITY}
    : ${COMPONENT?missing environment variable COMPONENT}

    echo "capability: ${CAPABILITY}"
    echo "component: ${COMPONENT}"
}

# defines function to set commmon variables for capability and component
set_capability_component_vars() {
    ensure_capability_component_input

    CAPABILITY_DIR_BASE=capabilities
    CAPABILITY_DIR=${CAPABILITY_DIR_BASE}/${CAPABILITY}
    CAPABILITY_DIR_SOURCE=${CAPABILITY_DIR}/.source
    COMPONENT_DIR=${CAPABILITY_DIR_SOURCE}/${COMPONENT}
}