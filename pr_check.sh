#!/bin/bash

# --------------------------------------------
# Options that must be configured by app owner
# --------------------------------------------
APP_NAME="platform-changelog"  # name of app-sre "application" folder this component lives in
COMPONENT_NAME="platform-changelog-go"  # name of app-sre "resourceTemplate" in deploy.yaml for this component
IMAGE="quay.io/cloudservices/platform-changelog-go"  

IQE_PLUGINS="platform-changelog"
IQE_MARKER_EXPRESSION="smoke"
IQE_FILTER_EXPRESSION=""
IQE_CJI_TIMEOUT="30m"

# Install bonfire repo/initialize
CICD_URL=https://raw.githubusercontent.com/RedHatInsights/bonfire/master/cicd
curl -s $CICD_URL/bootstrap.sh > .cicd_bootstrap.sh && source .cicd_bootstrap.sh

source $CICD_ROOT/build.sh

# Execute unit tests
# source $APP_ROOT/unit_test.sh

# Deploy to ephemeral env
# source $CICD_ROOT/deploy_ephemeral_env.sh
# source $CICD_ROOT/cji_smoke_test.sh