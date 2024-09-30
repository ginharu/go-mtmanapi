package pkg

import (
	"github.com/asaka1234/go-mtmanapi/win64/mtmanapi"
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
func GetGroupSpreadDiff(manager mtmanapi.CManagerInterface, req mtmanapi.RequestInfo) (decimal.Decimal, decimal.Decimal) {
	//增加组点
	var symbolInfo mtmanapi.ConSymbol
	symbol := req.GetTrade().GetSymbol()
	group := req.GetGroup()
	manager.SymbolGet(symbol, symbolInfo)
	xtype := symbolInfo.GetXtype()
	digit := symbolInfo.GetDigits()

	var groupInfo mtmanapi.ConGroup
	manager.GroupRecordGet(group, groupInfo)

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

	return bidVal, askVal
}

// 获取组点
func GetGroupSpreadDiffBySymbol(manager mtmanapi.CManagerInterface, group string, symbol string) (decimal.Decimal, decimal.Decimal) {
	//增加组点
	var symbolInfo mtmanapi.ConSymbol
	manager.SymbolGet(symbol, symbolInfo)
	xtype := symbolInfo.GetXtype()
	digit := symbolInfo.GetDigits()

	var groupInfo mtmanapi.ConGroup
	manager.GroupRecordGet(group, groupInfo)

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

	return bidVal, askVal
}
