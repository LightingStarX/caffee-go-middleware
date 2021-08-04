package stringx

import (
	"errors"
	"github.com/oxcaffee/caffee-go-middleware/common/global"
	"github.com/oxcaffee/caffee-go-middleware/common/utils/mathx"
	"time"
)

var (
	ErrInvalidStartPosition = errors.New("[common/utils/stringx]: Invalid start position error")
	ErrInvalidEndPosition   = errors.New("[common/utils/stringx]: Invalid end position error")
)

const (
	defaultRandLen = 8
	letterIdxBits  = 6                    // 6比特来表示一个字符的索引位
	letterIdxMax   = 1<<letterIdxBits - 1 // 在该索引位下的索引边界
	letterIdxMask  = 1<<letterIdxBits - 1 // 字符掩码
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var src = mathx.NewLockedRandSource(time.Now().UnixNano())

// RandString 返回一个默认长度为8的字符串
func RandString() string {
	return RandStringWithLen(defaultRandLen)
}

// RandStringWithLen 返回一个给定长度的随机字符串
func RandStringWithLen(n int) string {
	b := make([]byte, n)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

// Contains 判断string数组中是否含有给定的string类型
func Contains(list []string, str string) bool {
	for _, each := range list {
		if each == str {
			return true
		}
	}
	return false
}

// Filter 将给定的字符串按照特定的规则滤除特定字符，很好的一个算法，原地过滤字符串
func Filter(s string, filterFunc func(r rune) bool) string {
	idx := 0
	chars := []rune(s)

	for i, x := range chars {
		if idx < i {
			chars[idx] = x
		}
		if !filterFunc(x) {
			idx++
		}
	}
	return string(chars[:idx])
}

// HasEmptyStr 判断给定的字符串数组是否含有空字符串
func HasEmptyStr(ss ...string) bool {
	for _, s := range ss {
		if len(s) == 0 {
			return true
		}
	}
	return false
}

// NotEmpty 判断给定的字符串数组是否全是非空字符串
func NotEmpty(ss ...string) bool {
	return !HasEmptyStr(ss...)
}

// Remove 从strings中移除ss中包含的字符串
func Remove(strings []string, ss ...string) []string {
	res := append([]string(nil), strings...)

	for _, target := range ss {
		idx := 0
		for _, tmp := range res {
			if target != tmp {
				res[idx] = tmp
				idx++
			}
		}
		res = res[:idx]
	}
	return res
}

// Reverse 反转字符串
func Reverse(s string) string {
	chars := []rune(s)
	n := len(s)

	for i := 0; i < n/2; i++ {
		ct := chars[i]
		chars[i] = chars[n-1-i]
		chars[n-1-i] = ct
	}
	return string(chars)
}

// SubStr 取子字符串
func SubStr(s string, start int, end int) (string, error) {
	if start < 0 || start > len(s) || start >= end {
		return "", ErrInvalidStartPosition
	}

	if end < 0 || end > len(s) {
		return "", ErrInvalidEndPosition
	}

	return string(s[start:end]), nil
}

// FirstNotEmptyStr 返回第一个不是空字符串的字符串，如果全部都是空，返回""
func FirstNotEmptyStr(ss ...string) string {
	for _, str := range ss {
		if len(str) > 0 {
			return str
		}
	}
	return ""
}

// FirstNotEmptyFuncRes 返回第一个结果不是空字符串的函数结果
func FirstNotEmptyFuncRes(fs ...func() string) string {
	for _, f := range fs {
		res := f()
		if len(res) > 0 {
			return res
		}
	}
	return ""
}

// Union 返回两个字符串数组的合并结果，去重
func Union(ss1 []string, ss2 []string) []string {
	set := make(map[string]global.UniversalType)

	for _, s1 := range ss1 {
		set[s1] = global.UniversalType{}
	}

	for _, s2 := range ss2 {
		set[s2] = global.UniversalType{}
	}

	merged := make([]string, 0, len(set))
	for str, _ := range set {
		merged = append(merged, str)
	}
	return merged
}
