#!/usr/bin/env bash

#   获取随机数
getRandomStr() {
    echo "md5:" $RANDOM |md5sum |cut -c 1-8;
}
getRandomStr