#!/bin/sh
# Structural checks only: these do not start Jenkins or execute a Pipeline.
set -eu

root=$(CDPATH='' cd -- "$(dirname -- "$0")/.." && pwd)
controller="$root/infra/jenkins"

require_line() {
    grep -Fqx -- "$2" "$1" || {
        printf 'Missing expected configuration in %s: %s\n' "$1" "$2" >&2
        exit 1
    }
}

require_line "$controller/Dockerfile" 'FROM jenkins/jenkins:lts'
require_line "$controller/Dockerfile" 'USER jenkins'
require_line "$controller/Dockerfile" 'RUN jenkins-plugin-cli --plugin-file /usr/share/jenkins/ref/plugins.txt'
require_line "$controller/Dockerfile" "    'Jenkins.get().setNumExecutors(0)' \\"
require_line "$controller/compose.yaml" 'name: meli-jenkins-entrega2'
require_line "$controller/compose.yaml" '    user: jenkins'
require_line "$controller/compose.yaml" '      - "127.0.0.1:8081:8080"'
require_line "$controller/compose.yaml" '      - jenkins_home:/var/jenkins_home'
require_line "$controller/compose.yaml" '    cap_drop:'
require_line "$controller/compose.yaml" '      - ALL'
require_line "$controller/compose.yaml" '      - no-new-privileges:true'

for plugin in git github workflow-aggregator pipeline-stage-view docker-workflow credentials-binding; do
    require_line "$controller/plugins.txt" "$plugin"
done

if grep -Ein 'privileged:|docker\.sock|50000|user:[[:space:]]*(root|"?0)|USER[[:space:]]+root|runSetupWizard|securityRealm|authorizationStrategy|useSecurity|JENKINS_ADMIN|GITHUB_TOKEN' \
    "$controller/compose.yaml" "$controller/Dockerfile"; then
    printf 'Unsafe or unexpected controller configuration found.\n' >&2
    exit 1
fi

printf 'PASS: Jenkins controller structural defaults and plugin declarations.\n'
