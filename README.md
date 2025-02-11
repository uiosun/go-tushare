# TuShare SDK (Go version)

> 谨慎炒股，珍惜生活！！！

## 用法

主要是对 TuShare 接口的 Go 封装。

### 流控

对于 TuShare 的流控，本 SDK 涵盖了流控时暂停的情况，但不会重置流控，推荐自行在实例外，通过协程调用 `api.ResetRateLimit()` 重置流控：

```go
token := ""
api := tushare.New(token)
go func() {
    api.ResetRateLimit() // 这将自动每分钟重置流控
}()
```

### 使用样例

如对具体接口有疑问，请参考接口注释、测试用例。

```go
import (
    "github.com/uiosun/go-tushare/client"
)

func main()  {
	token := "your tushare token"
	tushareClient := client.New(token)

	params := make(map[string]string)
	params["list_status"] = "L"
	var fields []string
	fieldStr := "ts_code,symbol,name,area,industry,market,list_date"
	fields = strings.Split(fieldStr, ",")
	
	data,err := tushareClient.StockBasic(params,fields)
	// todo fill your code
}
```

## 测试

请在本地创建 `env.json` 的副本 `env-dev.json`，然后填入自己的 TuShare token。

## 涵盖接口

部分接口未封装，如果有需要，请提交 PR（各接口的实现方式参考 Python 版源码 / 本包源码）。

- [x] 交易所交易日历
- [x] 沪股通、深股通成分
- [x] 沪股通、深股通、港股通每日资金流向
- [x] 沪股通、深股通每日前十大成交
- [x] 港股通每日成交数据
- [x] 沪深两市每日融资融券明细
- [x] 融资融券每日交易汇总
- [x] 历史名称变更记录
- [x] 上市公司基础信息
- [x] 新股上市列表
- [x] 上市公司财务利润表
- [x] 上市公司资产负债表
- [x] 上市公司现金流量表
- [x] 上市公司前十大股东数据
- [x] 上市公司前十大流通股东数据
- [x] 上市公司股东户数数据
- [x] 龙虎榜每日交易明细
- [x] 龙虎榜机构成交明细
- [x] 股权质押统计数据
- [x] 股权质押明细数据
- [x] 股权质押统计数据
- [x] 上市公司回购股票数据
- [x] 概念股分类
- [x] 概念股分类明细
- [x] 限售股解禁
- [x] 大宗交易
- [x] 业绩预告
- [x] 分红送股
- [x] 上市公司业绩快报
- [x] 上市公司财务指标
- [x] 上市公司定期财务审计意见
- [x] 上市公司主营业务构成（地区/产品）
- [x] 股票账户开户数据（周）
- [x] 股票每日重要的基本面指标
- [x] 财报披露计划日期

### 指数

- [x] 指数基础信息(`index_basic`)
- [x] 指数日线行情(`index_daily`)
- [x] 申万行业分类(`index_classify`)
- [x] 申万行业成分构成(`index_member_all`)
- [x] 申万行业日线行情(`sw_daily`)
- [x] 同花顺概念和行业指数
- [x] 同花顺概念板块成分
- [x] 同花顺板块指数行情

### 股票

- [x] 股票基础信息(`stock_basic`)
- [x] 股票行情数据，分钟线(`stk_mins`)
- [x] 股票行情数据，日线(`daily`)
- [x] 股票行情数据，周线(`weekly`)
- [x] 股票日涨跌停数据(`stk_limit`)
- [x] 股票复权因子(`adj_factor`)
- [x] 股票每日停复牌信息(`suspend`)

## 致谢

> 本项目基于 yushikuann/go-tushare-sdk，由于未来需要快速更新，单独创建项目。

- [TuShare](https://tushare.pro)
- [yushikuann/go-tushare-sdk](https://github.com/yushikuann/go-tushare-sdk)
