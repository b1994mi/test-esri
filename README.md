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
| `/`    | GET    | Will give a simple response.|  |`{"message": "pong"}`|
| `/register`    | POST    | Register a new user.| `{"full_name:"user biasa", ""username": "userbiasa", "password": "123123"}` |`{"data": {"id":1, "username": "userbiasa"}}`|

| `/login`    | POST    | Get the token needed for restricted endpoints. Have limited expiry time of 3 min.| `{"username": "admin", "password": "admin"}` |`{"data": "eyJhbGciOiJIUzI1NiIsI"}`|
| `/complaint`    | GET    | List of complaints accessible for that user.| `?page=1&size=10` |`{"data": [{"id": 1, "user_id": 1, "meteran_id": 2}]}`|
| `/complaint`    | POST    | Save complaint to db and store images. Request must use Formdata. Will return the saved object as response.| Formdata: `issue: {"meteran_id": 1,"category_id": 1,"complaint_name": "Sebuah keluhan","short_description": "","priority_level": 1}`; Formdata: `image: file` |`{"data": {"id": 1, "user_id": 1, "meteran_id": 2}}`|

