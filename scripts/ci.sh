#!/bin/sh
# Shared local/Jenkins checks. Only disposable, explicitly named projects may run.
set -eu
cd "$(dirname "$0")/.."
case "${CI_PROJECT:-}" in
  meli-check-*) ;;
  *) echo 'Set CI_PROJECT=meli-check-<unique-run> for an isolated run.' >&2; exit 2 ;;
esac
case "$CI_PROJECT" in *[!a-z0-9-]*) echo 'Invalid CI_PROJECT' >&2; exit 2 ;; esac
export WEB_PORT="${WEB_PORT:-0}"
EVIDENCE_DIR=${EVIDENCE_DIR:-evidence/entrega2/local}
mkdir -p "$EVIDENCE_DIR"
compose() { docker compose -p "$CI_PROJECT" "$@"; }
record() {
  name=$1; shift
  if "$@" > "$EVIDENCE_DIR/$name.log" 2>&1; then
    cat "$EVIDENCE_DIR/$name.log"
  else
    result=$?
    cat "$EVIDENCE_DIR/$name.log" >&2
    return "$result"
  fi
}
case "${1:-}" in
  build)
    record build compose build api web
    ;;
  test)
    record backend-tests docker build --no-cache --progress=plain -f infra/ci/backend.Dockerfile backend
    record frontend-tests docker build --no-cache --progress=plain -f infra/ci/frontend.Dockerfile frontend
    ;;
  smoke)
    # Preserve test failure; cleanup failures must also fail an otherwise successful run.
    cleanup() {
      result=$?
      trap - EXIT
      if compose down > "$EVIDENCE_DIR/cleanup.log" 2>&1; then :; else
        cleanup_result=$?
        if [ "$result" -eq 0 ]; then result=$cleanup_result; fi
      fi
      exit "$result"
    }
    trap cleanup EXIT
    record startup compose up -d --wait --wait-timeout 120 --no-build
    record services compose ps
    record direct-api compose exec -T web wget -qO- http://api:8080/api/health
    record proxy-api compose exec -T web wget -qO- http://127.0.0.1:8080/api/health
    record web-health compose exec -T web wget -qO- http://127.0.0.1:8080/healthz
    grep -qx '{"status":"ok"}' "$EVIDENCE_DIR/direct-api.log"
    grep -qx '{"status":"ok"}' "$EVIDENCE_DIR/proxy-api.log"
    grep -qx 'ok' "$EVIDENCE_DIR/web-health.log"
    docker network inspect "${CI_PROJECT}_application" > "$EVIDENCE_DIR/network-inspect.json"
    for service in api web; do
      id=$(compose ps -q "$service")
      test -n "$id"
      test "$(docker inspect --format '{{.State.Health.Status}}' "$id")" = healthy
      grep -q "$id" "$EVIDENCE_DIR/network-inspect.json"
    done
    echo 'Two healthy services communicate by api service name and through the web proxy.'
    ;;
  cleanup) record cleanup compose down ;;
  *) echo 'Usage: CI_PROJECT=meli-check-<unique-run> sh scripts/ci.sh build|test|smoke|cleanup' >&2; exit 2 ;;
esac
