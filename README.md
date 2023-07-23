## BETTINGS MANAGEMENT API 📈
# API to manage your bookings.
Made this simples CRUD API with GO and it's own http library.


## FEATURES
- List all bookings
- Get details of a specific booking
- Add a new booking to your list
- Update the status, odds or stake of a booking
- Delete a booking


## INSTALATION
1. Clone this repository: `git clone https://github.com/matheusjost/go-betting-api.git`
2. Navigate to the project directory: `cd go-betting-api`


## CONFIGURATION
1. Create a config.toml file on base dir:
```
[api]
port = httpserver_port

[database]
host = db_host
user = db_user
pass = pass
name = db_name
```
2. Install dependencies:
```
go mod tidy
```


## USAGE
1. Run main file: 
```
go run main.go
```
2. Access the API @ `http://localhost:9000`


## API ENDPOINTS
- GET `/`: Returns all bookings.
- GET `/:id`: Returns details of a specific booking.
- POST `/`: Adds a new booking.
- PUT `/:id`: Updates the status, odds or stake of a booking.
- DELETE `/:id`: Deletes an existing booking.

## CONTRIBUTION
I'm still learning Golang basics so feel free to help me or give any advice!
You can fork the repository and make a pull request!