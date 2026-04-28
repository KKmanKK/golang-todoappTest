### Goalng todoApp


# Найти процесс, использующий порт 5432
sudo lsof -i :5432
# Если это локальный PostgreSQL, остановите его:
sudo systemctl stop postgresql  # для Linux

sudo chmod -R 755 /home/vboxuser/Desktop/projects/golang_todoapp/out/