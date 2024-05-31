#!/bin/bash

go build -o file2go main.go

strip file2go

mv file2go $HOME/go/bin/
