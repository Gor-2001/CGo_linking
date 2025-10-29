#!/usr/bin/env bash
set -e
gcc -c math_utils.c -o math_utils.o
gcc -c string_utils.c -o string_utils.o
ar rcs libhelpers.a math_utils.o string_utils.o
echo "✅ Built libhelpers.a"
