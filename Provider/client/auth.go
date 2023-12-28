package air

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

// Default Azure Tenant Id
const TenantId string = "63982aff-fb6c-4c22-973b-70e4acfb63e6"

func structToValues(authStruct AuthStruct) url.Values {
	values := url.Values{}
	structValue := reflect.ValueOf(authStruct)
	structType := reflect.TypeOf(authStruct)

	for i := 0; i < structValue.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("json")
		value := structValue.Field(i).Interface()
		values.Add(tag, fmt.Sprintf("%v", value))
	}

	return values
}

// SignIn - Get a new token for user
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.ClientId == "" || c.Auth.ClientSecret == "" {
		return nil, fmt.Errorf("define clientId and clientSecret")
	}

	reqBody := structToValues(c.Auth).Encode()

	req, err := http.NewRequest("POST", fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/token", TenantId), strings.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	body, err := c.doRequest(req, nil)
	if err != nil {
		return nil, err
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
