#!/bin/bash

dir=$1
export XCHAIN_ROOT=`pwd`

# install docker in precondition
if ! command -v docker &>/dev/null; then
    echo "missing docker command, please install docker first."
    exit 1
fi

# check if xdev available
if ! command -v ./bin/xdev &>/dev/null; then
    echo "missing xdev command, please cd ${XDEV_ROOT} && make"
    exit 1
fi

# build examples
mkdir -p build
for elem in `ls $dir`; do
    cc=$dir/$elem

    # build single cc file
    if [[ -f $cc ]]; then
        out=build/$(basename $elem .cc).wasm
        echo "build $cc"
        ./bin/xdev build -o $out $cc
    fi

    # build package
    if [[ -d $cc ]]; then
        echo "build $cc"
        bash -c "cd $cc && ./bin/xdev build && mv -v $elem.wasm ../../build/"
    fi
    echo
done