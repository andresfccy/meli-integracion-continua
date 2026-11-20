#!/bin/sh
set -eu
# The controller-issued secret is provisioned separately, never in the image/env.
if [ ! -r /run/jenkins-agent/secret ] || [ ! -s /run/jenkins-agent/secret ]; then
    echo 'Agent secret is unavailable. Complete Jenkins setup and provision the dedicated secret volume.' >&2
    exit 2
fi
exec /usr/local/bin/jenkins-agent "$@"
