package utils

import (
	"bytes"
	"crypto/md5"
	"dcr-gin/app/global"
	"encoding/hex"
	"fmt"
	"github.com/speps/go-hashids"
	"golang.org/x/crypto/bcrypt"
	"sort"
	"strings"
)

// Md5x
//@function: Md5x
//@description: md5加密
//@param: str string
//@return: string
func Md5x(str string) string {
	md5Hash := md5.New()
	md5Hash.Write([]byte(str))
	cipherStr := md5Hash.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// getMd5Sign
//@function: getMd5Sign
//@description: md5加密
//@param: paramMap map[string]string, signKey string
//@return: string
func getMd5Sign(paramMap map[string]string, signKey string) string {
	sortString := getSortString(paramMap)
	sign := Md5x(sortString + "&salt=" + signKey)
	return strings.ToUpper(sign)
}

// CheckMd5Sign
//@function: CheckMd5Sign
//@description: 验证签名
//@param: rspMap map[string]string, md5Key, sign string
//@return: bool
func CheckMd5Sign(rspMap map[string]string, md5Key, sign string) bool {
	calculateSign := getMd5Sign(rspMap, md5Key)
	fmt.Println(calculateSign, sign, md5Key)
	return calculateSign == sign
}

// getSortString
//@function: getSortString
//@description: 字符串排序拼接
//@param: m map[string]string
//@return: string
func getSortString(map2 map[string]string) string {
	var buf bytes.Buffer
	keysSlice := make([]string, 0, len(map2))
	for k := range map2 {
		keysSlice = append(keysSlice, k)
	}
	sort.Strings(keysSlice)
	//遍历切片
	for _, k := range keysSlice {
		vs := map2[k]
		if vs == "" {
			continue
		}
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		buf.WriteString(vs)
	}
	return buf.String()
}

// GeneratePassword 对明文密码进行加密
func GeneratePassword(password string) (string, string, error) {
	//获取随机支付串
	salt := Randx(5, KC_RAND_KIND_ALL)
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("%s%s", password, salt)), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}
	return string(hashPassword), salt, nil

}

// CheckPassword 校验密码是否正确
func CheckPassword(sqlPassword string, password string, salt string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(sqlPassword), []byte(fmt.Sprintf("%s%s", password, salt)))
	if err != nil {
		return false, err
	}
	return true, nil

}

// HashidsDecode 校验密码是否正确
func HashidsDecode(param string) (string, error) {
	hashIds := hashids.NewData()
	hashIds.Salt = global.ServerConfig.Hashids.Salt
	hashIds.MinLength = global.ServerConfig.Hashids.MinLength
	h, _ := hashids.NewWithData(hashIds)
	stationId, err := h.DecodeHex(param)
	return stationId, err
}

// HashidsEncode 校验密码是否正确
func HashidsEncode(param string) (string, error) {
	hashIds := hashids.NewData()
	hashIds.Salt = global.ServerConfig.Hashids.Salt
	hashIds.MinLength = global.ServerConfig.Hashids.MinLength
	h, _ := hashids.NewWithData(hashIds)
	stationId, err := h.EncodeHex(param)
	return stationId, err
}
