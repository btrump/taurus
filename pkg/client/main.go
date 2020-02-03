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
		ID:   "bosa3f4",
		Name: "client1",
	},
	"2": {
		ID:   "oitnc0d",
		Name: "client2",
	},
	"3": {
		ID:   "8eexmm0",
		Name: "client3",
	},
}
