if test -n "${ZSH_VERSION-}"; then
  autoload bashcompinit
  bashcompinit
fi
complete -F __aliyun_main aliyun

ALL=($(aliyun --verbose --quiet --help))

InstanceActions=(
  list     list-instances
  allocate allocate-public-ip
  start    start-instance
  stop     stop-instance
  restart  restart-instance
  remove   remove-instance
)
InstanceActions=$(printf "|%s" "${InstanceActions[@]}")
InstanceActions=${InstanceActions:1}

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

  if test "$COMP_CWORD" -gt 2; then
    local prever=${COMP_WORDS[COMP_CWORD-2]}
    case "$prever" in
    $InstanceActions)
      COMPREPLY=()
      return 0
      ;;
    esac
  fi

  local prev=${COMP_WORDS[COMP_CWORD-1]}
  case "$prev" in
  $InstanceActions)
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
  images|regions|types|groups|create)
    COMPREPLY=()
    ;;
  list-images|list-regions|list-instance-types|list-security-groups|create-instance)
    COMPREPLY=()
    ;;
  *)
    COMPREPLY=(${ALL[@]})
    ;;
  esac
}
