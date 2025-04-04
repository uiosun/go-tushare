package client

// Income 获取上市公司财务利润表数据
func (api *TuShare) Income(params map[string]string, isVip bool) (*APIResponse, error) {
	apiName := "income"
	if !isVip {
		// Check params
		_, hasTsCode := params["ts_code"]
		if !hasTsCode {
			return nil, ERR_ARGUEMENT
		}
	} else {
		apiName = "income_vip"
	}

	fields := []string{
		"ts_code",
		"ann_date",
		"f_ann_date",
		"end_date",
		"report_type",
		"comp_type",
		"end_type",
		"basic_eps",
		"diluted_eps",
		"total_revenue",
		"revenue",
		"int_income",
		"prem_earned",
		"comm_income",
		"n_commis_income",
		"n_oth_income",
		"n_oth_b_income",
		"prem_income",
		"out_prem",
		"une_prem_reser",
		"reins_income",
		"n_sec_tb_income",
		"n_sec_uw_income",
		"n_asset_mg_income",
		"oth_b_income",
		"fv_value_chg_gain",
		"invest_income",
		"ass_invest_income",
		"forex_gain",
		"total_cogs",
		"oper_cost",
		"int_exp",
		"comm_exp",
		"biz_tax_surchg",
		"sell_exp",
		"admin_exp",
		"fin_exp",
		"assets_impair_loss",
		"prem_refund",
		"compens_payout",
		"reser_insur_liab",
		"div_payt",
		"reins_exp",
		"oper_exp",
		"compens_payout_refu",
		"insur_reser_refu",
		"reins_cost_refund",
		"other_bus_cost",
		"operate_profit",
		"non_oper_income",
		"non_oper_exp",
		"nca_disploss",
		"total_profit",
		"income_tax",
		"n_income",
		"n_income_attr_p",
		"minority_gain",
		"oth_compr_income",
		"t_compr_income",
		"compr_inc_attr_p",
		"compr_inc_attr_m_s",
		"ebit",
		"ebitda",
		"insurance_exp",
		"undist_profit",
		"distable_profit",
		"rd_exp",
		"fin_exp_int_exp",
		"fin_exp_int_inc",
		"transfer_surplus_rese",
		"transfer_housing_imprest",
		"transfer_oth",
		"adj_lossgain",
		"withdra_legal_surplus",
		"withdra_legal_pubfund",
		"withdra_biz_devfund",
		"withdra_rese_fund",
		"withdra_oth_ersu",
		"workers_welfare",
		"distr_profit_shrhder",
		"prfshare_payable_dvd",
		"comshare_payable_dvd",
		"capit_comstock_div",
		"net_after_nr_lp_correct",
		"credit_impa_loss",
		"net_expo_hedging_benefits",
		"oth_impair_loss_assets",
		"total_opcost",
		"amodcost_fin_assets",
		"oth_income",
		"asset_disp_income",
		"continued_net_profit",
		"end_net_profit",
		"update_flag",
	}
	body := map[string]interface{}{
		"api_name": apiName,
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)

}

// BalanceSheet 获取上市公司资产负债表
func (api *TuShare) BalanceSheet(params map[string]string, isVip bool) (*APIResponse, error) {
	apiName := "balancesheet"
	if !isVip {
		// Check params
		_, hasTsCode := params["ts_code"]
		if !hasTsCode {
			return nil, ERR_ARGUEMENT
		}
	} else {
		apiName = "balancesheet_vip"
	}

	fields := []string{
		"ts_code",
		"ann_date",
		"f_ann_date",
		"end_date",
		"report_type",
		"comp_type",
		"end_type",
		"total_share",
		"cap_rese",
		"undistr_porfit",
		"surplus_rese",
		"special_rese",
		"money_cap",
		"trad_asset",
		"notes_receiv",
		"accounts_receiv",
		"oth_receiv",
		"prepayment",
		"div_receiv",
		"int_receiv",
		"inventories",
		"amor_exp",
		"nca_within_1y",
		"sett_rsrv",
		"loanto_oth_bank_fi",
		"premium_receiv",
		"reinsur_receiv",
		"reinsur_res_receiv",
		"pur_resale_fa",
		"oth_cur_assets",
		"total_cur_assets",
		"fa_avail_for_sale",
		"htm_invest",
		"lt_eqt_invest",
		"invest_real_estate",
		"time_deposits",
		"oth_assets",
		"lt_rec",
		"fix_assets",
		"cip",
		"const_materials",
		"fixed_assets_disp",
		"produc_bio_assets",
		"oil_and_gas_assets",
		"intan_assets",
		"r_and_d",
		"goodwill",
		"lt_amor_exp",
		"defer_tax_assets",
		"decr_in_disbur",
		"oth_nca",
		"total_nca",
		"cash_reser_cb",
		"depos_in_oth_bfi",
		"prec_metals",
		"deriv_assets",
		"rr_reins_une_prem",
		"rr_reins_outstd_cla",
		"rr_reins_lins_liab",
		"rr_reins_lthins_liab",
		"refund_depos",
		"ph_pledge_loans",
		"refund_cap_depos",
		"indep_acct_assets",
		"client_depos",
		"client_prov",
		"transac_seat_fee",
		"invest_as_receiv",
		"total_assets",
		"lt_borr",
		"st_borr",
		"cb_borr",
		"depos_ib_deposits",
		"loan_oth_bank",
		"trading_fl",
		"notes_payable",
		"acct_payable",
		"adv_receipts",
		"sold_for_repur_fa",
		"comm_payable",
		"payroll_payable",
		"taxes_payable",
		"int_payable",
		"div_payable",
		"oth_payable",
		"acc_exp",
		"deferred_inc",
		"st_bonds_payable",
		"payable_to_reinsurer",
		"rsrv_insur_cont",
		"acting_trading_sec",
		"acting_uw_sec",
		"non_cur_liab_due_1y",
		"oth_cur_liab",
		"total_cur_liab",
		"bond_payable",
		"lt_payable",
		"specific_payables",
		"estimated_liab",
		"defer_tax_liab",
		"defer_inc_non_cur_liab",
		"oth_ncl",
		"total_ncl",
		"depos_oth_bfi",
		"deriv_liab",
		"depos",
		"agency_bus_liab",
		"oth_liab",
		"prem_receiv_adva",
		"depos_received",
		"ph_invest",
		"reser_une_prem",
		"reser_outstd_claims",
		"reser_lins_liab",
		"reser_lthins_liab",
		"indept_acc_liab",
		"pledge_borr",
		"indem_payable",
		"policy_div_payable",
		"total_liab",
		"treasury_share",
		"ordin_risk_reser",
		"forex_differ",
		"invest_loss_unconf",
		"minority_int",
		"total_hldr_eqy_exc_min_int",
		"total_hldr_eqy_inc_min_int",
		"total_liab_hldr_eqy",
		"lt_payroll_payable",
		"oth_comp_income",
		"oth_eqt_tools",
		"oth_eqt_tools_p_shr",
		"lending_funds",
		"acc_receivable",
		"st_fin_payable",
		"payables",
		"hfs_assets",
		"hfs_sales",
		"cost_fin_assets",
		"fair_value_fin_assets",
		"cip_total",
		"oth_pay_total",
		"long_pay_total",
		"debt_invest",
		"oth_debt_invest",
		"oth_eq_invest",
		"oth_illiq_fin_assets",
		"oth_eq_ppbond",
		"receiv_financing",
		"use_right_assets",
		"lease_liab",
		"contract_assets",
		"contract_liab",
		"accounts_receiv_bill",
		"accounts_pay",
		"oth_rcv_total",
		"fix_assets_total",
		"update_flag",
	}
	body := map[string]interface{}{
		"api_name": apiName,
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// CashFlow 获取上市公司现金流量表
func (api *TuShare) CashFlow(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, ERR_ARGUEMENT
	}

	body := map[string]interface{}{
		"api_name": "cashflow",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Forecast 获取业绩预告数据
func (api *TuShare) Forecast(params map[string]string, isVip bool) (*APIResponse, error) {
	apiName := "forecast"
	if !isVip {
		// Check params
		_, hasTsCode := params["ts_code"]
		if !hasTsCode {
			return nil, ERR_ARGUEMENT
		}
	} else {
		apiName = "forecast_vip"
	}

	fields := []string{
		"ts_code",
		"ann_date",
		"end_date",
		"type",
		"p_change_min",
		"p_change_max",
		"net_profit_min",
		"net_profit_max",
		"last_parent_net",
		"first_ann_date",
		"summary",
		"change_reason",
		"update_flag",
	}
	body := map[string]interface{}{
		"api_name": apiName,
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Dividend 分红送股数据
func (api *TuShare) Dividend(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "dividend",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// Express 获取上市公司业绩快报
func (api *TuShare) Express(params map[string]string, isVip bool) (*APIResponse, error) {
	apiName := "express"
	if !isVip {
		// Check params
		_, hasTsCode := params["ts_code"]
		if !hasTsCode {
			return nil, ERR_ARGUEMENT
		}
	} else {
		apiName = "express_vip"
	}

	fields := []string{
		"ts_code",
		"ann_date",
		"end_date",
		"revenue",
		"operate_profit",
		"total_profit",
		"n_income",
		"total_assets",
		"total_hldr_eqy_exc_min_int",
		"diluted_eps",
		"diluted_roe",
		"yoy_net_profit",
		"bps",
		"yoy_sales",
		"yoy_op",
		"yoy_tp",
		"yoy_dedu_np",
		"yoy_eps",
		"yoy_roe",
		"growth_assets",
		"yoy_equity",
		"growth_bps",
		"or_last_year",
		"op_last_year",
		"tp_last_year",
		"np_last_year",
		"eps_last_year",
		"open_net_assets",
		"open_bps",
		"perf_summary",
		"is_audit",
		"remark",
		"update_flag",
	}
	body := map[string]interface{}{
		"api_name": apiName,
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// FinaIndicator 获取上市公司财务指标数据
func (api *TuShare) FinaIndicator(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, ERR_ARGUEMENT
	}

	body := map[string]interface{}{
		"api_name": "fina_indicator",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// FinaAudit 获取上市公司定期财务审计意见数据
func (api *TuShare) FinaAudit(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, ERR_ARGUEMENT
	}

	body := map[string]interface{}{
		"api_name": "fina_audit",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// FinaMainbz 获得上市公司主营业务构成，分地区和产品两种方式
func (api *TuShare) FinaMainbz(params map[string]string, fields []string) (*APIResponse, error) {
	// Check params
	_, hasTsCode := params["ts_code"]
	if !hasTsCode {
		return nil, ERR_ARGUEMENT
	}

	body := map[string]interface{}{
		"api_name": "fina_mainbz",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}

// DisclosureDate 获取财报披露计划日期
func (api *TuShare) DisclosureDate(params map[string]string, fields []string) (*APIResponse, error) {

	body := map[string]interface{}{
		"api_name": "disclosure_date",
		"token":    api.token,
		"params":   params,
		"fields":   fields,
	}

	return api.postData(body)
}
