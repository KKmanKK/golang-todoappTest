#!/bin/bash
mkdir -p out/pgdata
sudo chown -R 999:999 out/pgdata
sudo chmod -R 755 out/pgdata
echo "✅ Права настроены"
