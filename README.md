# IP Locator

This Go program handles HTTP requests and performs an IP geolocation lookup on the `http://ip-api.com/json/` endpoint. The program checks if the IP address has been seen in the last 5 minutes, and if not, it appends the geolocation data to a log file named status.log.

## Requirements

- Go 1.16 or later
- Internet connection to access the `http://ip-api.com/json/` endpoint

## Usage

- Clone this repository to your local machine.
- Open a terminal and navigate to the cloned repository directory.
- To start the server, run the following command in the project directory:<br>`go run main.go`
- The server will start on port 8080 by default. To specify a different port, set the PORT environment variable: <br> `PORT=8888 go run main.go`
- Send HTTP requests to the program using any client, e.g. curl.
- The client should set the X-Real-IP and X-Service-Name headers in the request.
- Check the status.log file for the appended geolocation data.
