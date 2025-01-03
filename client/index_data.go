package client

// IndexDaily 指数日线行情
func (api *TuShare) IndexDaily(params map[string]string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_daily",
		"token":    api.token,
		"params":   params,
	}

	return api.postData(body)
}

// IndexWeight 指数成分及权重
func (api *TuShare) IndexWeight(params map[string]string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_weight",
		"token":    api.token,
		"params":   params,
	}

	return api.postData(body)
}

// SzDailyInfo 深圳市场每日交易概况
func (api *TuShare) SzDailyInfo(params map[string]string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "sz_daily_info",
		"token":    api.token,
		"params":   params,
	}

	return api.postData(body)
}

// IndexClassify 获取申万行业分类
func (api *TuShare) IndexClassify(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_classify",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// IndexMember 获取申万行业成分构成
func (api *TuShare) IndexMember(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "index_member_all",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// SwDaily 获取申万行业行情（日线）数据
func (api *TuShare) SwDaily(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "sw_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// ThsIndex 同花顺概念和行业指数
func (api *TuShare) ThsIndex(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ths_index",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// ThsMember 同花顺概念板块成分
func (api *TuShare) ThsMember(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ths_member",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// ThsDaily 同花顺板块指数行情
func (api *TuShare) ThsDaily(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "ths_daily",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
