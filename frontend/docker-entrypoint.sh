#!/bin/sh
set -eu

API_UPSTREAM=${API_UPSTREAM:-api:8080}
export API_UPSTREAM

invalid_upstream() {
    printf '%s\n' 'Invalid API_UPSTREAM: expected DNS hostname:port (port 1-65535).' >&2
    exit 1
}

# Only a literal DNS hostname and decimal port may enter the nginx configuration.
# Reject delimiters, whitespace, shell/nginx syntax, URLs, and IPv6 literals.
case "$API_UPSTREAM" in
    *:*) host=${API_UPSTREAM%:*}; port=${API_UPSTREAM##*:} ;;
    *) invalid_upstream ;;
esac
case "$host" in
    ''|*[!a-zA-Z0-9.-]*|.*|*.|*..*) invalid_upstream ;;
esac
[ "${#host}" -le 253 ] || invalid_upstream
old_ifs=$IFS
IFS=.
set -- $host
IFS=$old_ifs
for label do
    [ "${#label}" -le 63 ] || invalid_upstream
    case "$label" in -*|*-) invalid_upstream ;; esac
done
case "$port" in ''|*[!0-9]*) invalid_upstream ;; esac
[ "${#port}" -le 5 ] && [ "$port" -ge 1 ] && [ "$port" -le 65535 ] || invalid_upstream

# Explicit substitution allowlist preserves nginx variables such as $host and $uri.
# /tmp is writable by USER nginx; /etc/nginx remains untouched at runtime.
envsubst '${API_UPSTREAM}' < /etc/nginx/nginx.conf.template > /tmp/nginx.conf
exec nginx -c /tmp/nginx.conf -g 'daemon off;'
