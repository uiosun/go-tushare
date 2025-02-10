package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type TuShare struct {
	token        string
	client       *http.Client
	config       TuShareConfig
	minute       int // 本分钟数
	requestCount int // 本分钟请求数
}

type TuShareConfig struct {
	RateLimit       bool // 由 SDK 负责流控
	RateLimitMinute int  // 每分钟最大请求数
}

func New(token string, config *TuShareConfig) *TuShare {
	return NewWithClient(token, http.DefaultClient, config)
}

func NewWithClient(token string, httpClient *http.Client, config *TuShareConfig) *TuShare {
	return &TuShare{
		token:  token,
		client: httpClient,
		config: *config,
	}
}

func (api *TuShare) request(method, path string, body interface{}) (*http.Request, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, path, bytes.NewBuffer(bodyJSON))
	if err != nil {
		return nil, err
	}
	return req, nil
}

func (api *TuShare) doRequest(req *http.Request) (*APIResponse, error) {
	req.Header.Set("Content-Type", "application/json")

	var resp *http.Response
	var err error
	if resp, err = api.client.Do(req); err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("net work error: %d", resp.StatusCode)
	}
	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return nil, err
	}

	var result *APIResponse
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	switch result.Code {
	case -2001:
		return result, ERR_ARGUEMENT
	case -2002:
		return result, ERR_PERMISSION
	case 40203:
		return result, ERR_TOO_MANY_REQUESTS
	case 0:
		return result, nil
	default:
		return result, fmt.Errorf("receive flurried code: %d, this's the message: %s", result.Code, result.Msg)
	}
}

func (api *TuShare) postData(body map[string]interface{}) (*APIResponse, error) {
	if api.config.RateLimit {
		if api.minute != time.Now().Minute() {
			api.minute = time.Now().Minute()
			api.requestCount = 0
		}
		if api.requestCount >= api.config.RateLimitMinute {
			time.Sleep(time.Duration(60-time.Now().Second()+1) * time.Second)
		} else {
			api.requestCount++
		}
	}
	req, err := api.request(PostMethod, Domain, body)
	if err != nil {
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
