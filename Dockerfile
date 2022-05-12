# First stage. Build choice. AS builder means, this stage is used for app building
FROM golang:1.15-alpine3.12 AS builder
# copy files from current directory to new folder with go module name
COPY . /github.com/Sacrament0/weather-bot/
# mark this directory as working. All commands will be executed there
WORKDIR /github.com/Sacrament0/weather-bot/
# download all dependences
RUN go mod download
# compile bin files to new directory "bin" and name new file "bot", then path to main.go
RUN go build -o ./bin/bot cmd/bot/main.go

# Second stage
FROM alpine:latest
# make workdir
WORKDIR /root/
# copy created bin file "bot" to this direcrory. --from=0 means copy from the previous build stage
# dot at the end means to the current directory
COPY --from=0 /github.com/Sacrament0/weather-bot/bin/bot .
# copy config folder to the new config folder, because app will not start without config
COPY --from=0 /github.com/Sacrament0/weather-bot/configs configs/
# starts this command in console. When container starting, this command runs
CMD ["./bot"]
