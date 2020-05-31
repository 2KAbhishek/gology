#!/bin/bash

cd "goatsay" || return
go install
cd ..

cd "goldog" || return
go install
cd ..

cd "goldog" || return
go install
cd ..

go install
