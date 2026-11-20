#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
scratch=$(mktemp -d "${TMPDIR:-/tmp}/meli-web-runtime.XXXXXX")
printf 'Test scratch (retained for parent cleanup): %s\n' "$scratch"
mkdir -p "$scratch/bin"

fail() { printf 'FAIL: %s\n' "$*" >&2; exit 1; }
[ -f "$root/frontend/docker-entrypoint.sh" ] || fail 'runtime entrypoint is missing'

# Remap fixed container paths only; exercise the production script unchanged otherwise.
sed -e "s|/etc/nginx/nginx.conf.template|$root/frontend/nginx.conf|g" \
    -e "s|/tmp/nginx.conf|$scratch/nginx.conf|g" \
    "$root/frontend/docker-entrypoint.sh" > "$scratch/entrypoint.sh"
cat > "$scratch/bin/envsubst" <<'STUB'
#!/bin/sh
set -eu
[ "$#" -eq 1 ] && [ "$1" = '${API_UPSTREAM}' ] || exit 97
printf 'envsubst\n' >> "$CALL_LOG"
sed 's|${API_UPSTREAM}|'"$API_UPSTREAM"'|g'
STUB
cat > "$scratch/bin/nginx" <<'STUB'
#!/bin/sh
set -eu
printf 'nginx\n' >> "$CALL_LOG"
[ "$#" -eq 4 ] && [ "$1" = '-c' ] && [ "$2" = "$RENDERED_CONFIG" ] &&
    [ "$3" = '-g' ] && [ "$4" = 'daemon off;' ]
STUB
chmod +x "$scratch/bin/envsubst" "$scratch/bin/nginx"
export PATH="$scratch/bin:$PATH" CALL_LOG="$scratch/calls" RENDERED_CONFIG="$scratch/nginx.conf"

run_valid() {
    : > "$CALL_LOG"
    if [ "$1" = unset ]; then
        (unset API_UPSTREAM; sh "$scratch/entrypoint.sh")
        expected=api:8080
    else
        API_UPSTREAM="$1" sh "$scratch/entrypoint.sh"
        expected=${1:-api:8080}
    fi
    [ "$(grep -Fc "proxy_pass http://$expected;" "$RENDERED_CONFIG")" -eq 2 ] || fail "upstream $expected was not rendered in both routes"
    for literal in '$host' '$proxy_add_x_forwarded_for' '$uri'; do
        grep -F "$literal" "$RENDERED_CONFIG" >/dev/null || fail "nginx variable $literal was replaced"
    done
    ! grep -F '${API_UPSTREAM}' "$RENDERED_CONFIG" >/dev/null || fail 'unrendered upstream placeholder'
    [ "$(cat "$CALL_LOG")" = "$(printf 'envsubst\nnginx')" ] || fail 'nginx did not execute after template rendering'
}

run_valid unset
run_valid ''
run_valid api:8080
run_valid backend.internal:9000
run_valid api-v2:1
run_valid 127.0.0.1:65535
run_valid API.internal:443

# Reject malformed addresses and nginx/shell injection before rendering or starting nginx.
for invalid in api api:0 api:65536 api:-1 api:abc api:123456 \
    :8080 api:8080:90 http://api:8080 api:8080/path api:8080\; \
    'api:8080; return 200 injected;' 'api:8080\n' 'api name:8080' \
    '-api:8080' 'api-:8080' 'a..b:8080' '.api:8080' 'api.:8080' \
    'a_thing:8080' 'api:8080#' '$(touch bad):8080' '[::1]:8080'; do
    : > "$CALL_LOG"
    if API_UPSTREAM="$invalid" sh "$scratch/entrypoint.sh" > "$scratch/out" 2> "$scratch/error"; then
        fail "accepted invalid API_UPSTREAM: $invalid"
    fi
    [ ! -s "$CALL_LOG" ] || fail 'invalid upstream reached envsubst or nginx'
    grep -F 'Invalid API_UPSTREAM' "$scratch/error" >/dev/null || fail 'missing validation error'
done
invalid_label=$(printf '%064d' 0)
if API_UPSTREAM="$invalid_label:8080" sh "$scratch/entrypoint.sh" > "$scratch/out" 2> "$scratch/error"; then
    fail 'accepted hostname label longer than 63 characters'
fi
valid_label=$(printf '%063d' 0)
if API_UPSTREAM="$valid_label.$valid_label.$valid_label.$valid_label:8080" sh "$scratch/entrypoint.sh" > "$scratch/out" 2> "$scratch/error"; then
    fail 'accepted hostname longer than 253 characters'
fi
printf 'PASS: defaults, overrides, nginx variables, launch arguments, invalid upstream rejection\n'
