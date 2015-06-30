package user
import (
	"log"
)

// Entity
// Only for aggregator. Not export.
type clientList []*Client
type ClientKey string //key type
type Client struct {
	ClientId ClientKey
	Name     string
}


// Factory
func NewClient(clientId ClientKey, name string) *Client {
	log.Printf("create client: %s\n", clientId)
	if clientId == "" {
		return nil
	}
	return &Client{
		ClientId: clientId,
		Name:     name,
	}
}

func (this *clientList) add(client *Client){
	*this = append(*this, client)
}
