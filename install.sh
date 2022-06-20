#!/bin/bash

export DEBIAN_FRONTEND=noninteractive

if systemctl --all | grep -Fq 'torwrapper'; then
	echo "Torwrapper is enabled. Stopping..."
	systemctl stop torwrapper
fi

apt update
apt install -yq golang
apt install -yq iptables
apt install -yq tor
apt install -yq obfs4proxy

systemctl tor start

go build -o torwrapper torwrapper.go
cp ./torwrapper /usr/bin/torwrapper

# Service for tool monitoring and maintaining
cp ./torwrapper.service /etc/systemd/system/torwrapper.service
chmod 644 /etc/systemd/system/torwrapper.service
systemctl daemon-reload

# Bridges for TOR
cp ./bridges.txt /etc/tor/bridges.txt

echo 'Torwrapper has been installed successfully. Type \"Torwrapper help\" to see supported commands.'