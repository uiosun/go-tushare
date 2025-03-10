package self

import (
	"fmt"
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
		fmt.Println(data.Data.Fields)
	}
}
