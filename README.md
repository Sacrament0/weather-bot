# Weather telegram bot
This telegram bot can show you weather (temperature, pressure, etc.) according to your geolocation. It has single button to send geolocation to server. 
## Try It!
This app has been deployed! Here is a link t.me/weather_around_bot
@weather_around_bot

## How it works
Bot uses [Open Weather Map](https://openweathermap.org/) service to get weather forecast. If there is a location in your message, bot sends it to OWM service. OWM handles your location and returns weather data. Finally bot handles weather data, forms weather message and sends it to user.

## What do we use?

 - [Viper](https://github.com/spf13/viper) - configuration solution
 - [Telegram bot API](https://github.com/go-telegram-bot-api/telegram-bot-api) - tg bot api for golang
 - [OpenWeatherMap API](https://github.com/briandowns/openweathermap) - open weather map service api for golang
 ## CI/CD solution
 - [Docker Hub](https://hub.docker.com/)
 - [GitHub actions](https://docs.github.com/en/actions)

This app is based on [M Zhashkevych](https://github.com/zhashkevych) lessons
