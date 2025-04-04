package self

import (
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestStockBasic(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	data, err := share.StockBasic(params)
	if err != nil {
		t.Error(err)
	}

	// [ts_code symbol name area industry fullname enname cnspell market exchange curr_type list_status list_date delist_date is_hs act_name act_ent_type]
	if len(data.Data.Fields) != 17 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestIndexBasic(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	data, err := share.IndexBasic(params)
	if err != nil {
		t.Error(err)
	}

	// [ts_code name fullname market publisher index_type category base_date base_point list_date weight_rule desc exp_date]
	if len(data.Data.Fields) != 13 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}
