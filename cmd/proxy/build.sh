#!/bin/bash

filedir=$(cd "$(dirname $0)";pwd)

outdir=$filedir/../../bin
cd $filedir
go build -o $outdir/proxy -i -a