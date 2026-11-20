#!/bin/sh
set -eu
root=$(CDPATH= cd -- "$(dirname "$0")/.." && pwd)
scratch=$(mktemp -d)
printf 'Scratch retained for parent cleanup: %s\n' "$scratch"
mkdir "$scratch/bin"
cat > "$scratch/bin/docker" <<'MOCK'
#!/bin/sh
printf '%s\n' "$*" >> "$CALLS"
exit "${DOCKER_EXIT:-0}"
MOCK
chmod +x "$scratch/bin/docker"
export PATH="$scratch/bin:$PATH" CALLS="$scratch/calls"
cd "$root"
if CI_PROJECT=meli-ci sh scripts/ci.sh cleanup; then echo 'Unsafe project accepted' >&2; exit 1; fi
test ! -f "$CALLS"
CI_PROJECT=meli-check-contract EVIDENCE_DIR="$scratch/evidence" sh scripts/ci.sh build
grep -q 'build api web' "$CALLS"
if DOCKER_EXIT=42 CI_PROJECT=meli-check-contract EVIDENCE_DIR="$scratch/failure" sh scripts/ci.sh build; then
  echo 'Build failure hidden' >&2; exit 1
else
  test "$?" -eq 42
fi
CI_PROJECT=meli-check-contract EVIDENCE_DIR="$scratch/test" sh scripts/ci.sh test
grep -q 'backend.Dockerfile' "$CALLS"
grep -q 'frontend.Dockerfile' "$CALLS"
if DOCKER_EXIT=43 CI_PROJECT=meli-check-contract EVIDENCE_DIR="$scratch/test-failure" sh scripts/ci.sh test; then
  echo 'Test failure hidden' >&2; exit 1
else
  test "$?" -eq 43
fi
if DOCKER_EXIT=44 CI_PROJECT=meli-check-contract EVIDENCE_DIR="$scratch/smoke-failure" sh scripts/ci.sh smoke; then
  echo 'Smoke failure hidden' >&2; exit 1
else
  test "$?" -eq 44
fi
tail -n 1 "$CALLS" | grep -Fx 'compose -p meli-check-contract down'
echo 'CI contract checks passed'  
