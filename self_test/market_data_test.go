package self

import (
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestSuspend(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	data, err := share.Suspend(params)
	if err != nil {
		t.Error(err)
	}

	// [ts_code trade_date suspend_timing suspend_type]
	if len(data.Data.Fields) != 4 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}

func TestStkLimit(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20190522"
	data, err := share.StkLimit(params)
	if err != nil {
		t.Error(err)
	}
	// [trade_date ts_code pre_close up_limit down_limit]
	if len(data.Data.Fields) != 5 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
	}
}
