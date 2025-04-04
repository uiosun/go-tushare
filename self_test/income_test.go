package self

import (
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestIncome(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20231229"

	if _, err := share.Income(params, false); err != client.ERR_ARGUEMENT {
		t.Fatal(err)
	}

	params["ts_code"] = "000001.SZ"
	data, err := share.Income(params, false)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date f_ann_date end_date report_type comp_type end_type basic_eps diluted_eps total_revenue revenue int_income prem_earned comm_income n_commis_income n_oth_income n_oth_b_income prem_income out_prem une_prem_reser reins_income n_sec_tb_income n_sec_uw_income n_asset_mg_income oth_b_income fv_value_chg_gain invest_income ass_invest_income forex_gain total_cogs oper_cost int_exp comm_exp biz_tax_surchg sell_exp admin_exp fin_exp assets_impair_loss prem_refund compens_payout reser_insur_liab div_payt reins_exp oper_exp compens_payout_refu insur_reser_refu reins_cost_refund other_bus_cost operate_profit non_oper_income non_oper_exp nca_disploss total_profit income_tax n_income n_income_attr_p minority_gain oth_compr_income t_compr_income compr_inc_attr_p compr_inc_attr_m_s ebit ebitda insurance_exp undist_profit distable_profit rd_exp fin_exp_int_exp fin_exp_int_inc transfer_surplus_rese transfer_housing_imprest transfer_oth adj_lossgain withdra_legal_surplus withdra_legal_pubfund withdra_biz_devfund withdra_rese_fund withdra_oth_ersu workers_welfare distr_profit_shrhder prfshare_payable_dvd comshare_payable_dvd capit_comstock_div net_after_nr_lp_correct credit_impa_loss net_expo_hedging_benefits oth_impair_loss_assets total_opcost amodcost_fin_assets oth_income asset_disp_income continued_net_profit end_net_profit update_flag]
	if len(data.Data.Fields) != 94 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestIncomeVip(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20231229"
	data, err := share.Income(params, true)
	if err != nil {
		t.Fatal(err)
	}

	// [ts_code ann_date f_ann_date end_date report_type comp_type end_type basic_eps diluted_eps total_revenue revenue int_income prem_earned comm_income n_commis_income n_oth_income n_oth_b_income prem_income out_prem une_prem_reser reins_income n_sec_tb_income n_sec_uw_income n_asset_mg_income oth_b_income fv_value_chg_gain invest_income ass_invest_income forex_gain total_cogs oper_cost int_exp comm_exp biz_tax_surchg sell_exp admin_exp fin_exp assets_impair_loss prem_refund compens_payout reser_insur_liab div_payt reins_exp oper_exp compens_payout_refu insur_reser_refu reins_cost_refund other_bus_cost operate_profit non_oper_income non_oper_exp nca_disploss total_profit income_tax n_income n_income_attr_p minority_gain oth_compr_income t_compr_income compr_inc_attr_p compr_inc_attr_m_s ebit ebitda insurance_exp undist_profit distable_profit rd_exp fin_exp_int_exp fin_exp_int_inc transfer_surplus_rese transfer_housing_imprest transfer_oth adj_lossgain withdra_legal_surplus withdra_legal_pubfund withdra_biz_devfund withdra_rese_fund withdra_oth_ersu workers_welfare distr_profit_shrhder prfshare_payable_dvd comshare_payable_dvd capit_comstock_div net_after_nr_lp_correct credit_impa_loss net_expo_hedging_benefits oth_impair_loss_assets total_opcost amodcost_fin_assets oth_income asset_disp_income continued_net_profit end_net_profit update_flag]
	if len(data.Data.Fields) != 94 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}
