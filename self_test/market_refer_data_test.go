package self

import (
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestMargin(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250416"
	data, err := share.Margin(params)
	if err != nil {
		t.Error(err)
	}

	// [trade_date exchange_id rzye rzmre rzche rqye rqmcl rzrqye rqyl]
	if len(data.Data.Fields) != 9 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestMarginDetail(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250416"
	data, err := share.MarginDetail(params)
	if err != nil {
		t.Error(err)
	}

	// [trade_date ts_code name rzye rqye rzmre rqyl rzche rqchl rqmcl rzrqye]
	if len(data.Data.Fields) != 11 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestMarginSecs(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20250416"
	data, err := share.MarginSecs(params)
	if err != nil {
		t.Error(err)
	}

	// [trade_date ts_code name exchange]
	if len(data.Data.Fields) != 4 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}
