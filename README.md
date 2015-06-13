aliyun
======

Command-line tool to operate Aliyun ECS instances easily.

You may need to create a new Access Key on https://ak-console.aliyun.com/,
and enter the key ID and secret when building this software (`make build`).

BUILD
-----

```
read -s ALIYUN_ACCESS_KEY && export ALIYUN_ACCESS_KEY
read -s ALIYUN_ACCESS_SECRET && export ALIYUN_ACCESS_SECRET
make build
```

Or in a Docker container:

```
read -s ALIYUN_ACCESS_KEY && export ALIYUN_ACCESS_KEY
read -s ALIYUN_ACCESS_SECRET && export ALIYUN_ACCESS_SECRET
docker run --rm \
  -e ALIYUN_ACCESS_KEY -e ALIYUN_ACCESS_SECRET \
  -v="$(pwd):/go/src/github.com/caiguanhao/aliyun" \
  -w="/go/src/github.com/caiguanhao/aliyun" golang:1.4.2 make build
```

USAGE
-----

```
# define your frequently used region
alias aliyun='aliyun --region cn-hangzhou'

# add bash completion
source bash_completion.d/aliyun.sh

# type aliyun --group then press <TAB> to list available group and so on
aliyun --group <TAB> --image <TAB> --type <TAB> --name <TAB> create
```

LICENSE: MIT
