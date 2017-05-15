package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	mplock sync.RWMutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	fmt.Println(URLshorten(`https://mail.qq.com/cgi-bin/frame_html?sid=gTdMx7z4dBwFEAVc&t=newwin_frame&url=%2fcgi-bin%2freadmail%3fmailid%3dZC4615-8EY35p3Z1r6nXksSm~FPK75%26need_textcontent%3dtrue%26s%3dnotify%26newwin%3Dtrue%26t%3dreadmail&r=3bef5eda5ccbc419dacb7cc4687d62ca`))
}

//小写hex
func Md5HexFromString(data, salt string) []byte {
	hash := md5.New()
	hash.Write([]byte(data))
	if len(salt) > 0 {
		hash.Write([]byte(salt))
	}
	return []byte(hex.EncodeToString(hash.Sum(nil)))
}

//生成一个长URL的短链接
func URLshorten(longURL string) string {
	//加密字符串
	key := "URLshorten"
	//URL字符表，共62个，下标为0~61
	text := "abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	md5text := Md5HexFromString(longURL, key)
	shortURLs := make([]string, 0, 4)
	for i := 0; i < 4; i++ {
		//取出MD5中第i个字节，并忽略超过30位的部分
		str := md5text[i*8 : (i+1)*8]
		//把str里的十六进制表示转化成int
		num, err := strconv.Atoi(fmt.Sprintf("%x", string(str)))
		if err != nil {
			panic(fmt.Errorf("invalid longurl:%v", longURL))
		}
		num &= 0x3FFFFFFF //选择低30位

		//取30位的后6位与0x0000003D进行逻辑与操作，结果范围是0~61，作为text的下标选择字符
		//把num左移5位重复进行，得到6个字符组成短URL
		shortURL := make([]byte, 0, 6)
		for j := 0; j < 6; j++ {
			shortURL = append(shortURL, text[num&0x0000003D])
			num >>= 5
		}
		shortURLs = append(shortURLs, string(shortURL))
	}
	//	fmt.Printf("%+v\n", shortURLs)
	return shortURLs[rand.Intn(len(shortURLs))]
	//	return shortURLs
}
