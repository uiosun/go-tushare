package client

import (
	"fmt"
)

// StockBasic 股票基础信息，包括股票代码、名称、上市日期、退市日期等
func (api *TuShare) StockBasic(params map[string]string) (*APIResponse, error) {
	fields := []string{
		"ts_code",
		"symbol",
		"name",
		"area",
		"industry",
		"fullname",
		"enname",
		"cnspell",
		"market",
		"exchange",
		"curr_type",
		"list_status",
		"list_date",
		"delist_date",
		"is_hs",
		"act_name",
		"act_ent_type",
	}
	body := map[string]interface{}{
		"api_name": "stock_basic",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// IndexBasic 指数基本信息
func (api *TuShare) IndexBasic(params map[string]string) (*APIResponse, error) {
	fields := []string{
		"ts_code",
		"name",
		"fullname",
		"market",
		"publisher",
		"index_type",
		"category",
		"base_date",
		"base_point",
		"list_date",
		"weight_rule",
		"desc",
		"exp_date",
	}
	body := map[string]interface{}{
		"api_name": "index_basic",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// TradeCal 获取各大交易所交易日历数据,默认提取的是上交所
func (api *TuShare) TradeCal(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "trade_cal",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// HSConst 获取沪股通、深股通成分数据
func (api *TuShare) HSConst(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "hs_const",
		"token":    api.token,
		"fields":   fields,
	}
	//bodyParam := make(map[string]string)
	// Add params
	if _, ok := params["hs_type"]; !ok {
		return nil, fmt.Errorf("hs_type required")
	}

	body["params"] = params

	return api.postData(body)
}

// NameChange 历史名称变更记录
func (api *TuShare) NameChange(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "namechange",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// StockCompany 获取上市公司基础信息
func (api *TuShare) StockCompany(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "stock_company",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// NewShare 获取新股上市列表数据
func (api *TuShare) NewShare(params map[string]string, fields []string) (*APIResponse, error) {
	body := map[string]interface{}{
		"api_name": "new_share",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
