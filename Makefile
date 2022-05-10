.PHONY:
.SOLENT:

build:
	go build -o ./.bin/bot cmd/bot/main.go

run: build
	./.bin/bot

# создание образа
build-image:
	docker build -t weather-bot:v0.1 .
# запуск контейнера
start-container:
# имя weather-bot, файл в котором описаны переменные окружения, 
# чтобы прокидывать их при запуске контейнера, в конце название образа
	docker run --name weather-bot --env-file .env weather-bot:v0.1