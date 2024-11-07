<p align="center">
  <img src="https://i.makeagif.com/media/12-24-2015/f63iKm.gif" alt="Christmas GIF"/>
</p>
# PapitoDDoS

A simple HTTP stress testing tool written in Go. This tool is designed for educational purposes and testing the resilience of your own web applications.

⚠️ **WARNING**: This tool should only be used for legitimate testing purposes on systems you own or have explicit permission to test. Unauthorized DDoS attacks are illegal.

## Features

- Concurrent request handling using Go routines
- Configurable number of workers
- Support for both GET and POST requests
- Custom headers configuration
- Form data submission with random data generation
- Real-time request counter and status monitoring

## Installation

```bash
git clone https://github.com/aviorp/papitoddos
cd papitoddos
go mod download
```

## Usage

### Basic GET Request

```go
go
config := &ddos.RequestConfig{
URL: "https://your-target-url.com",
Method: "GET",
WorkerCount: 1000,
}
d := ddos.New(config)
d.Start()
```

### Advanced POST Request with Headers

```go
config := &ddos.RequestConfig{
URL: "https://your-target-url.com",
Method: "POST",
WorkerCount: 1000,
Headers: map[string]string{
"Content-Type": "application/x-www-form-urlencoded",
"User-Agent": "Mozilla/5.0...",
"Accept": "/",
// Add more headers as needed
},
}
// Add form data
data := url.Values{}
data.Set("name", utils.GenerateRandomName())
data.Set("email", utils.GenerateRandomEmail())
data.Set("phone", utils.GenerateRandomPhone())
config.Body = strings.NewReader(data.Encode())
d := ddos.New(config)
d.Start()

```

## Configuration Options

| Option      | Description                      |
| ----------- | -------------------------------- |
| URL         | Target URL for the requests      |
| Method      | HTTP method (GET, POST, etc.)    |
| Headers     | Map of HTTP headers              |
| Body        | Request body (for POST requests) |
| WorkerCount | Number of concurrent workers     |

## Utility Functions

The tool includes utility functions for generating random data:

- `GenerateRandomName()`: Generates random full names
- `GenerateRandomEmail()`: Generates random email addresses
- `GenerateRandomPhone()`: Generates random phone numbers (Israeli format)

## Notes

1. The tool uses a shared counter to track the total number of requests made
2. Each worker runs in its own goroutine for maximum performance
3. Failed requests are logged but don't stop the execution
4. The program runs indefinitely until manually stopped
5. Consider rate limiting and your system's capabilities when setting WorkerCount

## Legal Disclaimer

This tool is provided for educational and testing purposes only. The authors are not responsible for any misuse or damage caused by this program. Users must ensure they have proper authorization before testing any systems.

## License

MIT License - See LICENSE file for details
