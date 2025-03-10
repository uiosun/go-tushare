package self

import (
	"fmt"
	"github.com/uiosun/go-tushare/client"
	"github.com/uiosun/go-tushare/config"
	"testing"
)

var token = ""

func getToken() string {
	if token == "" {
		token = config.GetConfig("/..")["TS_TOKEN"]
	}

	return token
}

func TestSzDailyInfo(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20200320"
	data, err := share.SzDailyInfo(params)
	if err != nil {
		t.Error(err)
	}
	if len(data.Data.Fields) != 9 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
		fmt.Println(data.Data.Fields)
	}
}

func TestKplList(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250306"
	data, err := share.KplList(params)
	if err != nil {
		t.Error(err)
	}
	// [ts_code name trade_date lu_time ld_time open_time last_time lu_desc tag theme net_change bid_amount status bid_change bid_turnover lu_bid_vol pct_chg bid_pct_chg rt_pct_chg limit_order amount turnover_rate free_float lu_limit_order]
	if len(data.Data.Fields) != 24 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
		fmt.Println(data.Data.Fields)
	}
}

func TestKplConcept(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250306"
	data, err := share.KplConcept(params)
	if err != nil {
		t.Error(err)
	}
	// [trade_date ts_code name z_t_num up_num]
	if len(data.Data.Fields) != 5 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
		fmt.Println(data.Data.Fields)
	}
}

func TestKplConceptCons(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250306"
	data, err := share.KplConceptCons(params)
	if err != nil {
		t.Error(err)
	}
	// [ts_code name con_name con_code trade_date desc hot_num]
	if len(data.Data.Fields) != 7 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
		fmt.Println(data.Data.Fields)
	}
}

func TestThsIndex(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250306"
	data, err := share.ThsIndex(params, []string{})
	if err != nil {
		t.Error(err)
	}
	// [ts_code name count exchange list_date type]
	if len(data.Data.Fields) != 6 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
		fmt.Println(data.Data.Fields)
	}
}

func TestThsMember(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250306"
	data, err := share.ThsMember(params)
	if err != nil {
		t.Error(err)
	}
	// [ts_code con_code con_name weight in_date out_date is_new]
	if len(data.Data.Fields) != 7 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
		fmt.Println(data.Data.Fields)
	}
}
