# База склада фирмы. Web-приложение.

Использованы следующие концепции:
- Запуск сервера: **net/http**; + HTML.
- Работа с базой данных Postgres в Docker-контейнере: **database/sql**, <a href="http://github.com/lib/pq">lib/pq</a>.
- Graceful Shutdown

### Для запуска приложения:
Install Docker: <a href="https://www.docker.com/get-started/">Link</a>


1) Первоначально для работы файлов миграции базы данных:  
<a href="https://www.freecodecamp.org/news/database-migration-golang-migrate/">Information</a>  
Windows Powershell:
```
scoop install migrate
```
Если не установлен <a href="https://scoop.sh/">scoop</a>:

```
Set-ExecutionPolicy RemoteSigned -Scope CurrentUser # Optional: Needed to run a remote script the first time
irm get.scoop.sh | iex
```

Linux:
```
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey| apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

2) Для Makefile (<a href="https://www.gnu.org/software/make/#download">make</a>):
```
scoop install main/make
```
3) Запуск впервые:
```
make build
make migrate
make run
```
Обычный запуск:
```
make run
```

---
Примечание к Docker:
- Остановка контейнера: 
```
docker stop name_of_container or container_ID
```
Список
```
docker ps -a
```
- Подключение: 
```
docker exec -it container_ID /bin/bash
$ psql -U postgres 
$ \l (выход \q) или \c имя или \d и др.
```