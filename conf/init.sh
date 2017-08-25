#!/bin/sh
set -e

echo "Setting conf"

cp /frpc /data/frpc 
cp /frps /data/frps

exec /frpc
