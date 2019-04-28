#!/bin/sh
find . -type d | xargs -L 1 golint
