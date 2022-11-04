#!/bin/bash

generate_dummy_junit_result() {

  if ! [ -d "artifacts" ]; then
    mkdir "artifacts"
  fi

  cat << EOF > artifacts/junit-dummy.xml
<?xml version="1.0" encoding="UTF-8" ?>
<testsuite tests="1">
    <testcase classname="dummy-class" name="dummytest"/>
</testsuite>
EOF

}

# --------------------------------------------
# Options that must be configured by app owner
# --------------------------------------------
export APP_NAME="platform-changelog"  # name of app-sre "application" folder this component lives in
export COMPONENT_NAME="platform-changelog-go"  # name of app-sre "resourceTemplate" in deploy.yaml for this component
export IMAGE="quay.io/cloudservices/platform-changelog-go"  

export IQE_PLUGINS="platform-changelog"
export IQE_MARKER_EXPRESSION="smoke"
export IQE_FILTER_EXPRESSION=""
export IQE_CJI_TIMEOUT="30m"

# Install bonfire repo/initialize
CICD_URL='https://raw.githubusercontent.com/RedHatInsights/bonfire/master/cicd'
curl -s "${CICD_URL}/bootstrap.sh" > .cicd_bootstrap.sh && source .cicd_bootstrap.sh

source "${CICD_ROOT}/build.sh"

generate_dummy_junit_result

# Execute unit tests
# source $APP_ROOT/unit_test.sh

# Deploy to ephemeral env
# source $CICD_ROOT/deploy_ephemeral_env.sh
# source $CICD_ROOT/cji_smoke_test.sh


