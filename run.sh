#!/bin/bash

# 执行source .zshrc 
# 加载环境变量使用
source ~/.zshrc

script_dir="$(cd "$(dirname "$0")" && pwd)"
if [ -e "${script_dir}/go_script" ]; then
  $script_dir/go_script tx-dns-update > ok.log
else
  go build
  $script_dir/go_script tx-dns-update > ok.log
fi