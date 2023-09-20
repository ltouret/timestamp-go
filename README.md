timestamp-go
============

timestamp-go is a simple API project that allows you to convert Unix timestamps to UTC strings and vice versa. It's built using Go and can handle a wide range of date inputs.

API Endpoints
-------------

### 1\. Convert a Date to Unix Timestamp and UTC String

*   **Endpoint**: `/api/v1/:date?`
    
*   **Method**: GET
    
*   **Description**: Converts a valid date to a Unix timestamp in milliseconds and its corresponding UTC string.
    
*   **Response**: JSON object with `unix` (Number) and `utc` (String) keys.
    
*   **Example**:
    
    Request: `GET /api/v1/792201600000`
    
    Response:
    
    ```json
    { "unix": 792201600000, "utc": "Wed, 08 Feb 1995 00:00:00 GMT" }
    ```
    

### 2\. Handling Invalid Date Input

*   **Description**: If the input date string is invalid, the API returns an object with the structure `{ "error": "Invalid Date" }`.
    
*   **Example**:
    
    Request: `GET /api/v1/invalid-date`
    
    Response:
    
    ```json
    { "error": "Invalid Date" }
    ```
    

### 3\. Get Current Time

*   **Description**: An empty date parameter returns the current time in a JSON object with `unix` and `utc` keys.
    
*   **Example**:
    
    Request: `GET /api/v1/`
    
    Response:
    
    ```json
    { "unix": [current Unix timestamp in milliseconds], "utc": "[current UTC time in the format: Thu, 01 Jan 1970 00:00:00 GMT]" }
    ```
    

Usage
-----

1.  Clone the repository:
    
    ```bash
    git clone git@github.com:ltouret/timestamp-go.git
    ```
    
2.  Navigate to the project directory:
    
    ```bash
    cd timestamp-go
    ```
    
3.  Build and run the application:
    
    ```bash
    go build
    ./timestamp-go
    ```
    
4.  Access the API using your preferred web browser or API testing tool.
    

Dependencies
------------

To run this project, you'll need to have the following dependencies installed:

- **Go 1.20**: The Go programming language.
- **MySQL Driver**: The MySQL driver for Go. You can install it using `go get github.com/go-sql-driver/mysql`.
- **Gin**: The Gin web framework for Go. You can install it using `go get github.com/gin-gonic/gin`.
- **godotenv**: A Go package for loading environment variables from a .env file. You can install it using `go get github.com/joho/godotenv`.

License
-------

This project is licensed under the MIT License.