package mtmanapi

// 获取32/64位
func GetSysVersion() int {
	const intSize = 32 << (^uint(0) >> 63)
	return intSize
}
