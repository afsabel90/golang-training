package air

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetApplications() ([]Application, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/api/applications", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	applications := []Application{}
	err = json.Unmarshal(body, &applications)
	if err != nil {
		return nil, err
	}

	return applications, nil
}
