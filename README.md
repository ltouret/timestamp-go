timestamp-go
============

timestamp-go is a simple API project that allows you to convert Unix timestamps to UTC strings and vice versa. It's built using Go and can handle a wide range of date inputs.

API Endpoints
-------------

### 1\. Convert a Date to Unix Timestamp and UTC String

*   **Endpoint**: `/api/:date?`
    
*   **Method**: GET
    
*   **Description**: Converts a valid date to a Unix timestamp in milliseconds and its corresponding UTC string.
    
*   **Response**: JSON object with `unix` (Number) and `utc` (String) keys.
    
*   **Example**:
    
    Request: `GET /api/1451001600000`
    
    Response:
    
    json
    
    ```json
    { "unix": 1451001600000, "utc": "Fri, 25 Dec 2015 00:00:00 GMT" }
    ```
    

### 2\. Handling Invalid Date Input

*   **Description**: If the input date string is invalid, the API returns an object with the structure `{ "error": "Invalid Date" }`.
    
*   **Example**:
    
    Request: `GET /api/invalid-date`
    
    Response:
    
    json
    
    ```json
    { "error": "Invalid Date" }
    ```
    

### 3\. Get Current Time

*   **Description**: An empty date parameter returns the current time in a JSON object with `unix` and `utc` keys.
    
*   **Example**:
    
    Request: `GET /api/`
    
    Response:
    
    json
    
    ```json
    { "unix": [current Unix timestamp in milliseconds], "utc": "[current UTC time in the format: Thu, 01 Jan 1970 00:00:00 GMT]" }
    ```
    

Usage
-----

1.  Clone the repository:
    
    bash
    
    ```bash
    git clone [repository_url]
    ```
    
2.  Navigate to the project directory:
    
    bash
    
    ```bash
    cd timestamp-go
    ```
    
3.  Build and run the application:
    
    bash
    
    ```bash
    go build
    ./timestamp-go
    ```
    
4.  Access the API using your preferred web browser or API testing tool.
    

Dependencies
------------

This project uses Go 1.20, so make sure you have Go installed on your system.

License
-------

This project is licensed under the MIT License.