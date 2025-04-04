package self_test

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

type FiledMapping struct {
	theType string
	trans   string
}

func TestTooManyRequest(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{
		RateLimit:       true,
		RateLimitMinute: 500,
	})
	for i := 0; i < 1000; i++ {
		if i%10 == 0 {
			fmt.Println(i)
		}
		params := make(map[string]string)
		params["ts_code"] = "000300.SH"
		params["limit"] = "1"
		_, err := share.IndexDaily(params)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestIndexDaily(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["ts_code"] = "000300.SH"
	data, err := share.IndexDaily(params)
	if err != nil {
		t.Error(err)
	}
	if len(data.Data.Fields) != 11 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
	}
}

func TestDaily(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["ts_code"] = "000300.SH"
	data, err := share.Daily(params, []string{})
	if err != nil {
		t.Error(err)
	}
	// [ts_code trade_date open high low close pre_close change pct_chg vol amount]
	if len(data.Data.Fields) != 11 {
		t.Errorf("fields count not has %d pieces", len(data.Data.Fields))
	}
}

func TestStkMins(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["ts_code"] = "000001.SZ"
	params["start_date"] = "2004-03-12 00:27:11"
	params["end_date"] = "2009-04-11 00:27:11"
	params["freq"] = "1min"
	params["limit"] = "5"
	params["offset"] = "5"
	data, err := share.StkMins(params)
	if err != nil {
		t.Error(err)
	}
	// [ts_code trade_time close open high low vol amount]
	if len(data.Data.Fields) != 8 {
		t.Errorf("fields count not has %d pieces, %s", len(data.Data.Fields), data.Data.Fields)
	}
	if len(data.Data.Items) != 5 {
		t.Errorf("data count not has %d pieces", len(data.Data.Items))
	}

	if len(data.Data.Items) > 0 {
		mapping := []FiledMapping{
			{"string", "symbol"},
			{"string", "tradeDate"},
			{"float64", "open"},
			{"float64", "sClose"},
			{"float64", "high"},
			{"float64", "low"},
			{"float64", "vol"},
			{"float64", "amount"},
		}
		for _, item := range data.Data.Items {
			for i2, slot := range mapping {
				var ok bool

				switch slot.theType {
				case "float64":
					_, ok = item[i2].(float64)
				case "string":
					_, ok = item[i2].(string)
				default:
					t.Errorf("Noun type: %s", slot.theType)
				}
				if !ok {
					fmt.Println(item)
					t.Errorf("No.%d: %s format is failed, the symbol %s", i2, slot.trans, item[0])
				}
			}
		}
	}
}
