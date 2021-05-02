package circle

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ReqClient struct {
	APIKey    string
	PublicKey string
	PublicID  string
	URL       string
	http.Client
}

func (c *ReqClient) GetPublicKey() error {
	var data publicKeyResponse

	resp, err := c.makeRequest("/encryption/public", "GET", nil)
	if err != nil {
		log.Println("Client.PublicKey() resp: ",err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("client.GetPublicKey() read: ", err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal("client.GetPublicKey() decode: ", err)
	}
	c.PublicID = data.Data.KeyID
	c.PublicKey = data.Data.PublicKey
	return err
	// fmt.Println(c.PublicKey)
	// fmt.Println(&data)
}

func (c *ReqClient) makeRequest(endpoint, rtype string, body io.Reader) (*http.Response, error) {
	var req *http.Request
	var err error
	bearer := fmt.Sprintf("Bearer %v", c.APIKey)

	switch rtype {
	case "GET":
		req, err = http.NewRequest("GET", c.URL+endpoint, nil)
		if err != nil {
			return nil, err
		}

	case "POST":
		req, err = http.NewRequest("POST", c.URL+endpoint, body)
		if err != nil {
			return nil, err
		}
	}
	req.Header.Set("Authorization", bearer)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
