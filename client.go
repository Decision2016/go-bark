/**
  @author: decision
  @date: 2024/4/8
  @note:
**/

package bark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	Url     string `json:"url"`     // API url with http/https
	PushKey string `json:"pushKey"` // string push key
	// todo: encrypt with AES
	Encrypt bool `json:"encrypt"`

	defaultTitle string
	defaultGroup string
	defaultLevel MsgLevel
}

func (c *Client) baseURL() string {
	return fmt.Sprintf("%s/%s", c.Url, c.PushKey)
}

func (c *Client) send(r *RequestBody) error {
	payloadBytes, err := json.Marshal(r)
	if err != nil {
		return err
	}

	payload := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", c.baseURL(), payload)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	result, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if result.StatusCode != http.StatusOK {
		return fmt.Errorf("Send post request to server failed.")
	}
	return nil
}

func NewClient(url, pushKeys string) *Client {
	return &Client{
		Url:     url,
		PushKey: pushKeys,
	}
}

var dc *Client = &Client{}

// SetAPIUrl
//
//	@Description: Set url for default client
//	@param url - bark api url
func SetAPIUrl(url string) {
	dc.Url = url
}

// SetPushKey
//
//	@Description: Set push key for default client
//	@param pushKey - bark api push key
func SetPushKey(pushKey string) {
	dc.PushKey = pushKey
}

func SetDefaultTitle(title string) {
	dc.defaultTitle = title
}

func SetDefaultGroup(group string) {
	dc.defaultGroup = group
}

func Active(r *RequestBody) {
	r.Level = MsgLevelActive

}
