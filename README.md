# Gerello

## Installation Server
- go to application directory
```sh
cd server
```
- download dependencies
```sh
go mod download
```
- move from .env.example to .env
```sh
mv .env.example .env
```
- start command
```sh
make app-up
```

- start application
```sh
make run
```

## Installation Client
- go to application directory
```sh
cd client
```
- download dependencies
```sh
npm i
```
- start application
```sh
npm run dev
```