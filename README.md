# test-esri
Simple REST API for Complaining App on Water Company

## How to Run
1. Rename `db.sqlite3.example` to `db.sqlite3`
2. Install go in your machine for building the executable
3. Run `go build main.go` from this directory
4. Run `./main` from this directory and make sure your port 5000 is free

## Default Username and Password
I have created some dummy data in `db.sqlite3.example` for you to test out. In any case you want to start fresh, you can create a ner `db.sqlite3` file and when the application starts it will automatically migrate required tables and data.
```
Default Username & Password
Username: admin
Password: admin
```

## Endpoints

| Endpoint      | HTTP   | Description  | Request Body/Params | Response Body|
| ------------- | ------ | ---------------------- | ---------------------------------------------------------------- |----|
| `/`    | GET    | Will give a simple response.|  |`{"message": "pong"}`
|
