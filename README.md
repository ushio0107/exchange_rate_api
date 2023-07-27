# Currency Exchange API

This is an API that provides currency exchange functionality using Golang, Gin, Cobra, and Viper.

## Features

- Currency exchange feature to convert amounts from one currency to another.
- Support for converting between multiple currencies, including Taiwanese Dollar (TWD), Japanese Yen (JPY), and US Dollar (USD).

## How to Use

1. **Install Golang**: Make sure you have Golang installed on your computer.

2. **Download the Project**: Clone or download the project to your computer.

3. **Install Dependencies**: In the root directory of the project, install the required dependencies by running the following commands:

```bash
go get -u github.com/gin-gonic/gin
go get -u github.com/spf13/cobra
go get -u github.com/spf13/viper
```

## Set up Configuration: 
Prepare a config.json file with the exchange rates data in the following format:
```json
{
  "currencies": {
    "TWD": {
      "TWD": 1,
      "JPY": 3.669,
      "USD": 0.03281
    },
    "JPY": {
      "TWD": 0.26956,
      "JPY": 1,
      "USD": 0.00885
    },
    "USD": {
      "TWD": 30.444,
      "JPY": 111.801,
      "USD": 1
    }
  }
}
```

## Build and Run: 
Build the project and run the API server using the following command:
```bash
go run main.go
```

## Access the API: 
Use your preferred API client (e.g., curl, Postman) to access the API. The endpoint for currency conversion is as follows:

```GET /convert?source=USD&target=JPY&amount=$1,525```

Replace source, target, and amount query parameters to perform different currency conversions.