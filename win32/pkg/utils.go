package pkg

import "github.com/asaka1234/go-mtmanapi/win32/mtmanapi"

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
