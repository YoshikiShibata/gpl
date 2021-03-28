#!/bin/bash

# Copyright (C) 2016, 2018, 2019 Yoshiki Shibata. All rights reserved.

go build -o clock2 clock.go
go build -o clockwall clockwall.go

trap ctrl_c INT

function ctrl_c() {
	killall clock2
}

killall clock2

TZ=US/Pacific    ./clock2 -port 8010 &
TZ=US/Eastern    ./clock2 -port 8020 &
TZ=Asia/Tokyo    ./clock2 -port 8030 &
TZ=Europe/London ./clock2 -port 8040 &
TZ=UTC ./clock2 -port 8050&
TZ=CEST ./clock2 -port 8060 & 

sleep 3
echo ""
./clockwall "Palo Alto"=localhost:8010 NewYork=localhost:8020 Tokyo=localhost:8030 London=localhost:8040 UTC=localhost:8050 Germany=localhost:8060
