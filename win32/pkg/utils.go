package pkg

import (
	"errors"
	"fmt"
	"github.com/asaka1234/go-mtmanapi/win32/mtmanapi"
	"github.com/shopspring/decimal"
	"math"
)

// 获取32/64位
func GetSysVersion() int {
	const intSize = 32 << (^uint(0) >> 63)
	return intSize
}

// 获取api version
func GetApiVersion() int {
	//(a << 16) | b
	return (mtmanapi.ManAPIProgramVersion << 16) | mtmanapi.ManAPIProgramBuild
}

// 获取组点
func GetGroupSpreadDiff(manager mtmanapi.CManagerInterface, req mtmanapi.RequestInfo) (*GroupSpreadValue, error) {
	symbol := req.GetTrade().GetSymbol()
	group := req.GetGroup()
	return GetGroupSpreadDiffBySymbol(manager, group, symbol)
}

type GroupSpreadValue struct {
	Bid   decimal.Decimal
	Ask   decimal.Decimal
	Digit int
}

// 获取组点(only can be used in pumping mode)
func GetGroupSpreadDiffBySymbol(manager mtmanapi.CManagerInterface, group string, symbol string) (*GroupSpreadValue, error) {
	//增加组点
	var symbolInfo mtmanapi.ConSymbol
	code := manager.SymbolGet(symbol, symbolInfo)
	if code != mtmanapi.RET_OK {
		return nil, errors.New(fmt.Sprintf("SymbolGet err, symbol:%s, errCode:%d", symbol, code))
	}
	xtype := symbolInfo.GetXtype()
	digit := symbolInfo.GetDigits()

	var groupInfo mtmanapi.ConGroup
	code = manager.GroupRecordGet(group, groupInfo)
	if code != mtmanapi.RET_OK {
		return nil, errors.New(fmt.Sprintf("GroupRecordGet err, symbol:%s, errCode:%d", symbol, code))
	}

	secGroups := groupInfo.GetSecgroups()
	singleGroup := mtmanapi.ConGroupSecArray_getitem(secGroups, int64(xtype))
	spreadDiff := singleGroup.GetSpread_diff() //获取组点

	//数量
	spreadBid := spreadDiff / 2
	spreadAsk := spreadDiff - spreadBid

	//基本单位
	denominator := math.Pow(0.1, float64(digit))

	//两个方向各自的组点值
	bidVal := decimal.NewFromInt(int64(spreadBid)).Mul(decimal.NewFromFloat(denominator))
	askVal := decimal.NewFromInt(int64(spreadAsk)).Mul(decimal.NewFromFloat(denominator))

	return &GroupSpreadValue{
		bidVal,
		askVal,
		digit,
	}, nil
}
