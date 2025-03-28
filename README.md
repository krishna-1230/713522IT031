# Average Calculator HTTP Microservice

A Go-based microservice that calculates the average of numbers fetched from external APIs.

## Requirements

- Go 1.21 or later
- Internet connection to access the external APIs

## Configuration

The service is configured with the following constants in `app/config.go`:

- `WindowSize`: Maximum number of unique numbers to store (default: 10)
- `RequestTimeout`: Maximum time for HTTP requests (default: 500ms)
- `BaseAPIURL`: Base URL for the external APIs
- `AuthToken`: Authorization token for API requests

**Important**: Before running the service, make sure to update the `AuthToken` constant in `app/config.go` with your actual bearer token.

## External APIs

The service uses the following external APIs:

- Prime numbers: `http://20.244.56.144/test/primes`
- Fibonacci numbers: `http://20.244.56.144/test/fibo`
- Even numbers: `http://20.244.56.144/test/even`
- Random numbers: `http://20.244.56.144/test/rand`

All API requests include an authorization header with a bearer token.

## Running the Service

```bash
# Install dependencies
go mod tidy

# Run the service
go run main.go
```

The service will start on port 9876.

## API Endpoints

### GET /numbers/{numberid}

Fetches numbers from the external APIs based on the provided number ID.

Supported number IDs:
- `p`: Prime numbers
- `f`: Fibonacci numbers
- `e`: Even numbers
- `r`: Random numbers

#### Response Format

```json
{
    "windowPrevState": [2, 4, 6, 8],        // Previous state of the window
    "windowCurrState": [2, 4, 6, 8, 10],    // Current state of the window after adding unique numbers
    "numbers": [6, 8, 10],                  // Numbers fetched from the external API
    "avg": 6.0                              // Average of numbers in the window
}
```

## Example

```
GET http://localhost:9876/numbers/e
```

First response example:
```json
{
    "windowPrevState": [],
    "windowCurrState": [2, 4, 6, 8],
    "numbers": [2, 4, 6, 8],
    "avg": 5.0
}
``` 