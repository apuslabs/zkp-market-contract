#!/bin/bash

# 设置要执行的命令
command_to_repeat="python scripts/contract_lib.py"

# 设置重复执行的次数
repeat_count=1000000000000000

# 循环执行命令
for ((i=1; i<=repeat_count; i++))
do
    echo "执行命令: $command_to_repeat"
    eval $command_to_repeat
    echo "执行次数: $i"
    echo "---------------------------"
done
