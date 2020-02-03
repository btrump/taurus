package client

func Hello() string {
	return "clientlib hello"
}

type Client struct {
	ID   string
	Name string
}

var Clients = map[string]Client{
	"1": {
		ID:   "1",
		Name: "Blair Trump",
	},
	"2": {
		ID:   "2",
		Name: "Mina Hu",
	},
}
