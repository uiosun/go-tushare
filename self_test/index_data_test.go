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
