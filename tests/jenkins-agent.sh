#!/bin/sh
set -eu
cd "$(dirname "$0")/.."
# Declarative security contract: controller remains unchanged and isolated.
test -f infra/jenkins/agent/Dockerfile
grep -q '^USER jenkins$' infra/jenkins/agent/Dockerfile
grep -q 'docker.sock:/var/run/docker.sock' infra/jenkins/compose.agent.yaml
grep -q 'external: true' infra/jenkins/compose.agent.yaml
grep -q 'read_only: true' infra/jenkins/compose.agent.yaml
grep -q '@' infra/jenkins/compose.agent.yaml
if grep -Eq 'privileged: true|user: root|jenkins_home' infra/jenkins/compose.agent.yaml; then
  echo 'Agent isolation contract violated' >&2; exit 1
fi
if sh infra/jenkins/agent/start.sh -help 2>/dev/null; then
  echo 'Missing agent secret must fail closed' >&2; exit 1
else
  test "$?" -eq 2
fi
sh -n infra/jenkins/agent/start.sh
echo 'PASS: non-root dedicated agent, secret volume and fail-closed startup.'
