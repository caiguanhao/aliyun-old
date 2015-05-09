if test -n "${ZSH_VERSION-}"; then
  autoload bashcompinit
  bashcompinit
fi

complete -F __aliyun_main aliyun

function __aliyun_has_action {
  for ((i=1; i<$COMP_CWORD; i++)); do
    case "${COMP_WORDS[$i]}" in
    -*)
      continue
      ;;
    *)
      return 0
      ;;
    esac
  done
  return 1
}

function __aliyun_main {
  if __aliyun_has_action; then
    local prev=${COMP_WORDS[COMP_CWORD-1]}
    case "$prev" in
    list|list-instances)
      COMPREPLY=($(aliyun --quiet list-instances))
      ;;
    esac
    return 0
  fi

  COMPREPLY=($(aliyun --quiet --help))
}
