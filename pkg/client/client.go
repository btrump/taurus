package client

func Hello() string {
	return "clientlib hello"
}

type Client struct {
	ID   string
	Name string
}
