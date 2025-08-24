# GoEnv

A small Go package that makes working with environment variables simple, type-safe, and convenient.

goenv provides helper functions for retrieving environment variables with fallbacks, strict parsing, and common types like int, float, bool, time.Time, and time.Duration.

## Features

- [x] Simple GetEnv with fallback values
- [x] Type-safe parsing (int, float32, float64, bool, time.Time, time.Duration)
- [x] TryGetEnv functions that return (value, error)
- [x] MustGetEnv functions that panic if the variable is missing or invalid
- [x] Clean and minimal API

## Installation

```bash
go get github.com/J-Me-2307/goenv
```

## Usage

### Basic

```go
import (
    "fmt"
    "time"

    "github.com/J-Me-2307/goenv"
)

func main() {
    // String with fallback
    host := goenv.GetEnv("HOST", "localhost")
    fmt.Println("Host:", host)

    // Int with fallback
    port := goenv.GetEnvInt("PORT", 8080)
    fmt.Println("Port:", port)

    // Bool with fallback
    debug := goenv.GetEnvBool("DEBUG", false)
    fmt.Println("Debug mode:", debug)

    // Time with fallback
    defaultTime := time.Now()
    start := goenv.GetEnvTime("START_TIME", defaultTime)
    fmt.Println("Start:", start)

    // Duration with fallback
    timeout := goenv.GetEnvDuration("TIMEOUT", 5*time.Second)
    fmt.Println("Timeout:", timeout)
}
```

### TryGet (with error handling)

```go
value, err := goenv.TryGetEnvInt("MAX_CONNECTIONS")
if err != nil {
    fmt.Println("Could not parse:", err)
} else {
    fmt.Println("Max connections:", value)
}

```

### MustGet (panics on error)

```go
secret := goenv.MustGetEnv("SECRET_KEY")
fmt.Println("Secret key:", secret)

```

## License

GPL-3.0 license
