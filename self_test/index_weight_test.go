package self

import (
	"fmt"
	"github.com/uiosun/go-tushare/client"
	"testing"
)

func TestIndexWeight(t *testing.T) {
	share := client.New(getToken())
	params := make(map[string]string)
	params["index_code"] = "000903.SH"
	params["trade_date"] = "20231229"
	data, err := share.IndexWeight(params)
	if err != nil {
		t.Error(err)
	}
	if len(data.Data.Fields) != 4 {
		t.Errorf("fields count not has %d pieces: %s", len(data.Data.Fields), data.Data.Fields)
		fmt.Println(data.Data.Fields)
	}
}
