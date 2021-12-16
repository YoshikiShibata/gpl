#!/bin/bash

go build -o a.out
#
./a.out
./a.out -temp 18C
./a.out -temp 212°F
./a.out -temp 273.15K

echo "// Output:"
echo "// 20°C"
echo "// -18°C"
echo "// 100°C"
echo "// 0°C"

rm a.out
