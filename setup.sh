# #!/bin/bash
# mkdir -p out/pgdata
# sudo chown -R 999:999 out/pgdata
# sudo chmod -R 755 out/pgdata
# echo "✅ Права настроены"

#!/bin/bash
mkdir -p out/pgdata
# Создать общую группу
sudo groupadd postgres_local
sudo usermod -a -G postgres_local vboxuser
# Установить права для группы
sudo chown -R 999:postgres_local out/pgdata
sudo chmod -R 775 out/pgdata
sudo chmod g+s out/pgdata
echo "✅ Права настроены для общего доступа"