#!/bin/bash

# Copyright (C) 2016, 2018 Yoshiki Shibata. All rights reserved.

go build -o clock2 clock.go
go build -o clockwall clockwall.go

trap ctrl_c INT

function ctrl_c() {
	killall clock2
}

killall clock2

TZ=US/Eastern    ./clock2 -port 8010 &
TZ=Asia/Tokyo    ./clock2 -port 8020 &
TZ=Europe/London ./clock2 -port 8030 &
TZ=UTC ./clock2 -port 8040&

echo ""
./clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030 UTC=localhost:8040
