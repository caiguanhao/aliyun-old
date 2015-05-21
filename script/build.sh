#!/bin/bash

set -e

function str_to_array {
  eval "local input=\"\$$1\""
  input="$(echo "$input" | awk '
  {
    split($0, chars, "")
    for (i = 1; i <= length($0); i++) {
      if (i > 1) {
        printf(", ")
      }
      printf("\\\\\\\"%s\\\\\\\"", chars[i])
    }
  }
  ')"
  eval "$1=\"$input\""
}

function update_access_key {
  str_to_array KEY
  str_to_array SECRET
  awk "
  /KEY/ {
    print \"var KEY = []string{${KEY}}\"
    next
  }
  /SECRET/ {
    print \"var SECRET = []string{${SECRET}}\"
    next
  }
  {
    print
  }
  " access_key.go > _access_key.go

  mv _access_key.go access_key.go
}

echo -n "Please paste your access key ID: (will not be echoed) "
read -s KEY
echo
echo -n "Please paste your access key SECRET: (will not be echoed) "
read -s SECRET
echo
update_access_key

go build

KEY="key"
SECRET="secret"
update_access_key
