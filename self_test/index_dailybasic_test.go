package self

import (
	"fmt"
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestIndexDailyBasic(t *testing.T) {
	share := client.New(getToken(), &client.TuShareConfig{})
	params := make(map[string]string)
	params["trade_date"] = "20231229"
	data, err := share.IndexDailyBasic(params)
	if err != nil {
		t.Error(err)
	}
	if len(data.Data.Fields) != 12 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
		fmt.Println(data.Data.Fields)
	}
}
