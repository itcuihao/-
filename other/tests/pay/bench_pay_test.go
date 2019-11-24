package pay

import (
	"strconv"
	"testing"
	"time"

	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/mch/pay"

	chpay "pay"
	chpay_com "pay/common"
	chpay_regi "pay/register"
)

var (
	appid  = "123"
	mchid  = "123"
	apiKey = "123"

	chanClient = core.NewClient(appid, mchid, apiKey, nil)
)

func TestChanPay(t *testing.T) {
	no:=strconv.FormatInt(time.Now().UnixNano(),10)
	req := map[string]string{
		"out_trade_no": no,
		"nonce_str":    "123",
		"body":         "123",
		"total_fee":    "1",
		"trade_type":   "APP",
		"notify_url":   "http://www.123.com",
	}
	res, err := pay.UnifiedOrder(chanClient, req)
	t.Log(err)
	t.Logf("res: %+v", res)
}

func TestChPay(t *testing.T) {
	chpay_regi.PayConf().RegisterWeChat(apiKey, appid, mchid)
	no:=strconv.FormatInt(time.Now().UnixNano(),10)
	payer ,err:= chpay.NewPayer(1)
	if err!=nil{
		t.Log(err)
		return
	}
	req := chpay_com.UnifiedOrderRequest{
		SignType:    "MD5",
		OutTradeNo:  no,
		NotifyURL:   "http://www.123.com",
		TotalFee:    1,
		NonceStr:    "123",
		ProductName: "123",
		Body:        "123",
		TradeType:   "APP",
	}
	res, err := payer.UnifiedOrder(req)
	t.Log(err)
	t.Logf("res: %+v", res)
}