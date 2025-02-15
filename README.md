Краткая выжимка: Есть 2 структурыЖ up и down, в up записываются изменения, в down записывается откат. По командам раскидаю ниже

Установка через терминал:

	go install github.com/pressly/goose/v3/cmd/goose@latest

1. Создание миграции:

	goose create add_team_table sql

	1.1 Создаетя файл с двумя структурами: одна для применения миграции (up), и другой для отката (down).
	1.2 Создать таблицу в up	

2. Скачать с докерхаб контейнер с Postgres (в моем случае)

	docker pull postgres
	docker run --name goose_db -e POSTGRES_PASSWORD=secret -p 5432:5432:5432 -d postgres

3. Применение изменений
	3.1 	goose postgres postgresql://postgres:secret@localhost:5432/postgres up

	Должен создаться какой-то лог, измениения применены

	3.2 Подключиться к БД (в VSC проще всего скачать расширение PostgreSQL и создать коннект)

4. Откат
	4.1	goose postgres postgresql://postgres:secret@localhost:5432/postgres down
	Должны появиться изменения 

дальше я создал (goose create add_second_table sql) second_table.sql и на ней уже более явно показано что нужно записать в структуру up, а что в down

5. Применение конктретной миграции
	goose postgres postgresql://postgres:secret@localhost:5432/postgres up-to /ИЛИ/ down-to<some id> 

6. Статус
   	goose postgres postgresql://postgres:secret@localhost:5432/postgres status
