package self

import (
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestBalanceSheet(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	data, err := share.BalanceSheet(params, true)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date f_ann_date end_date report_type comp_type end_type total_share cap_rese undistr_porfit surplus_rese special_rese money_cap trad_asset notes_receiv accounts_receiv oth_receiv prepayment div_receiv int_receiv inventories amor_exp nca_within_1y sett_rsrv loanto_oth_bank_fi premium_receiv reinsur_receiv reinsur_res_receiv pur_resale_fa oth_cur_assets total_cur_assets fa_avail_for_sale htm_invest lt_eqt_invest invest_real_estate time_deposits oth_assets lt_rec fix_assets cip const_materials fixed_assets_disp produc_bio_assets oil_and_gas_assets intan_assets r_and_d goodwill lt_amor_exp defer_tax_assets decr_in_disbur oth_nca total_nca cash_reser_cb depos_in_oth_bfi prec_metals deriv_assets rr_reins_une_prem rr_reins_outstd_cla rr_reins_lins_liab rr_reins_lthins_liab refund_depos ph_pledge_loans refund_cap_depos indep_acct_assets client_depos client_prov transac_seat_fee invest_as_receiv total_assets lt_borr st_borr cb_borr depos_ib_deposits loan_oth_bank trading_fl notes_payable acct_payable adv_receipts sold_for_repur_fa comm_payable payroll_payable taxes_payable int_payable div_payable oth_payable acc_exp deferred_inc st_bonds_payable payable_to_reinsurer rsrv_insur_cont acting_trading_sec acting_uw_sec non_cur_liab_due_1y oth_cur_liab total_cur_liab bond_payable lt_payable specific_payables estimated_liab defer_tax_liab defer_inc_non_cur_liab oth_ncl total_ncl depos_oth_bfi deriv_liab depos agency_bus_liab oth_liab prem_receiv_adva depos_received ph_invest reser_une_prem reser_outstd_claims reser_lins_liab reser_lthins_liab indept_acc_liab pledge_borr indem_payable policy_div_payable total_liab treasury_share ordin_risk_reser forex_differ invest_loss_unconf minority_int total_hldr_eqy_exc_min_int total_hldr_eqy_inc_min_int total_liab_hldr_eqy lt_payroll_payable oth_comp_income oth_eqt_tools oth_eqt_tools_p_shr lending_funds acc_receivable st_fin_payable payables hfs_assets hfs_sales cost_fin_assets fair_value_fin_assets cip_total oth_pay_total long_pay_total debt_invest oth_debt_invest oth_eq_invest oth_illiq_fin_assets oth_eq_ppbond receiv_financing use_right_assets lease_liab contract_assets contract_liab accounts_receiv_bill accounts_pay oth_rcv_total fix_assets_total update_flag]
	if len(data.Data.Fields) != 158 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestForecast(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["period"] = "20231231"
	if _, err := share.Forecast(params, false); err != client.ERR_ARGUEMENT {
		t.Fatal(err)
	}

	params["ts_code"] = "000001.SZ"
	data, err := share.Forecast(params, false)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date end_date type p_change_min p_change_max net_profit_min net_profit_max last_parent_net first_ann_date summary change_reason update_flag]
	if len(data.Data.Fields) != 13 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestForecastVip(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["period"] = "20231231"
	data, err := share.Forecast(params, true)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date end_date type p_change_min p_change_max net_profit_min net_profit_max last_parent_net first_ann_date summary change_reason update_flag]
	if len(data.Data.Fields) != 13 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestExpress(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["period"] = "20231231"
	if _, err := share.Express(params, false); err != client.ERR_ARGUEMENT {
		t.Fatal(err)
	}

	params["ts_code"] = "000001.SZ"
	data, err := share.Express(params, false)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date end_date revenue operate_profit total_profit n_income total_assets total_hldr_eqy_exc_min_int diluted_eps diluted_roe yoy_net_profit bps yoy_sales yoy_op yoy_tp yoy_dedu_np yoy_eps yoy_roe growth_assets yoy_equity growth_bps or_last_year op_last_year tp_last_year np_last_year eps_last_year open_net_assets open_bps perf_summary is_audit remark update_flag]
	if len(data.Data.Fields) != 33 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestExpressVip(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["period"] = "20231231"
	data, err := share.Express(params, true)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date end_date revenue operate_profit total_profit n_income total_assets total_hldr_eqy_exc_min_int diluted_eps diluted_roe yoy_net_profit bps yoy_sales yoy_op yoy_tp yoy_dedu_np yoy_eps yoy_roe growth_assets yoy_equity growth_bps or_last_year op_last_year tp_last_year np_last_year eps_last_year open_net_assets open_bps perf_summary is_audit remark update_flag]
	if len(data.Data.Fields) != 33 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}
