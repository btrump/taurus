# taurus-server
## Usage
## Quick Links
|Verb|Link|Description|
|-|-|-|
|GET|http://localhost:8081/client/1/connect |Add the first client to the list of connected clients|
|GET|http://localhost:8081/api/status| API status|
|GET|http://localhost:8081/server/status |Server status|
|POST|http://localhost:8081|Send a command|

## Requests
```go
type Request struct {
	ID        string
	Timestamp time.Time
	UserID    string
	Command   string
	Message   string
}
```
