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
	requestCount int // 本分钟请求数
}

type TuShareConfig struct {
	AutoWaitRateLimit    bool          // 在流控触发时，不报错，自动等待流控
	AutoWaitRateLimitSec time.Duration // 在流控触发时，自动等待的秒数
	RateLimit            bool          // 由 SDK 负责流控
	RateLimitMinute      int           // 每分钟最大请求数
}

func New(token string, config *TuShareConfig) *TuShare {
	if config.AutoWaitRateLimit && config.AutoWaitRateLimitSec == 0 {
		config.AutoWaitRateLimitSec = 45
	}
	return NewWithClient(token, http.DefaultClient, config)
}

func NewWithClient(token string, httpClient *http.Client, config *TuShareConfig) *TuShare {
	api := &TuShare{
		token:  token,
		client: httpClient,
		config: *config,
	}
	return api
}

// ResetRateLimit 重置流控
func (api *TuShare) ResetRateLimit() {
	time.Sleep(61 * time.Second)
	api.requestCount = 0
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
	var err error

	// 读取原始body内容
	bodyBytes, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	bodyCache := ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	var resp *http.Response
	if resp, err = api.client.Do(req); err != nil {
		fmt.Println("api.client.Do has error.", err, req)
		return nil, err
	}
	if resp.StatusCode == http.StatusBadGateway {
		fmt.Println(resp)
		return nil, fmt.Errorf("bad work error: %d", resp.StatusCode)
	} else if resp.StatusCode != http.StatusOK {
		fmt.Println(resp)
		return nil, fmt.Errorf("net work error: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		fmt.Println("Body has error.")
		return nil, err
	}

	var result *APIResponse
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Println("Body to Result has error.")
		return nil, err
	}

	switch result.Code {
	case -2001:
		return result, ERR_ARGUEMENT
	case -2002:
		return result, ERR_PERMISSION
	case 40203:
		if api.config.AutoWaitRateLimit {
			time.Sleep(api.config.AutoWaitRateLimitSec * time.Second)
			req.Body = bodyCache
			result, err := api.doRequest(req)
			if err != nil {
				fmt.Println("40203 retry is dead.")
				return result, ERR_TOO_MANY_REQUESTS
			}
			return result, nil
		}
		return result, ERR_TOO_MANY_REQUESTS
	case 0:
		return result, nil
	default:
		return result, fmt.Errorf("receive flurried code: %d, this's the message: %s", result.Code, result.Msg)
	}
}

func (api *TuShare) postData(body map[string]interface{}) (*APIResponse, error) {
	if api.config.RateLimit {
		api.requestCount++
	}
	req, err := api.request(PostMethod, Domain, body)
	if err != nil {
		println("Request has error.")
		return nil, err
	}
	resp, err := api.doRequest(req)
	if err != nil {
		println("Do request has error.")
		return resp, err
	}
	return resp, nil
}
