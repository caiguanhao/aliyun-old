if test -n "${ZSH_VERSION-}"; then
  autoload bashcompinit
  bashcompinit
fi
complete -F __aliyun_main aliyun

ALL=($(aliyun --verbose --quiet --help))

function __aliyun_main {
  local region=""
  local i=0
  for word in "${COMP_WORDS[@]}"; do
    if test "$i" -lt "$COMP_CWORD"; then
      local j=0
      for all in "${ALL[@]}"; do
        if test "$all" == "$word"; then
          ALL[$j]=""
        fi
        let "j=j+1"
      done
    fi

    let "i=i+1"
    case "${word}" in
    --region)
      region="${COMP_WORDS[$i]}"
      ;;
    esac
  done

  local prever=${COMP_WORDS[COMP_CWORD-2]}
  case "$prever" in
  list|list-instances)
    COMPREPLY=()
    return 0
    ;;
  esac

  local prev=${COMP_WORDS[COMP_CWORD-1]}
  case "$prev" in
  list|list-instances)
    COMPREPLY=($(aliyun --quiet --region "$region" list-instances))
    ;;
  --name)
    COMPREPLY=($(aliyun --quiet --region "$region" --print-name list-instances))
    ;;
  --type)
    COMPREPLY=($(aliyun --quiet list-instance-types))
    ;;
  --image)
    COMPREPLY=($(aliyun --quiet --region "$region" list-images))
    ;;
  --region)
    COMPREPLY=($(aliyun --quiet list-regions))
    ;;
  --group)
    COMPREPLY=($(aliyun --quiet --region "$region" list-security-groups))
    ;;
  images|regions|types|groups|create|allocate|start|stop|restart|remove)
    COMPREPLY=()
    ;;
  list-images|list-regions|list-instance-types|list-security-groups|create-instance)
    COMPREPLY=()
    ;;
  allocate-public-ip|start-instance|stop-instance|restart-instance|remove-instance)
    COMPREPLY=()
    ;;
  *)
    COMPREPLY=(${ALL[@]})
    ;;
  esac
}
