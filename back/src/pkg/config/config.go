package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"strconv"
	"time"
)

var conf *ini.File

func Init(configFileName string) {
	var err error
	conf, err = ini.Load(configFileName)
	if err != nil {
		panic("config Init error :" + err.Error())
	}
}

func GetStringSevere(section string, key string) string {
	if !conf.HasSection(section) {
		panic(fmt.Errorf("No section named %s ", section))
	}
	sec := conf.Section(section)
	if !sec.HasKey(key) {
		panic(fmt.Errorf("No key named %s in section %s ", key, section))
	}
	return sec.Key(key).String()
}
func GetIntSevere(section string, key string) int {
	str := GetStringSevere(section, key)
	intVal, err := strconv.Atoi(str)
	if err != nil {
		panic("config get error :" + err.Error())
	}
	return intVal
}

/*
从config中读取时间
时间解析的规则如下:
解析器会遍历时间字符串,一次读取一个数字+单位的组合,例如:
1year2day  会被认为是两个组 1year 和 2day
随后,解析器会将所有组的时间加起来,最为最终的结果,例如:
1y1y 解析为 2年时间,即duration=2*365*24*60*60*10^9
1M2y 解析为2年零一个月...
所有的单位的定义如下:
n	nano		纳秒
micor			微秒
milli			毫秒
s 	second 		秒
m 	minute 		分
h 	hour		时
d	day			日
M	month		月
y	year		年

此外,每个月默认30天,这会引入误差,请注意
当出现无法理解的情况时,会panic
*/

var multiple = [9]int64{1, 1000, 1000, 1000, 60, 60, 24, 30, 12}
var names = [9][]string{
	{"n", "ns", "nano"},
	{"micor"},
	{"ms", "milli"},
	{"s", "second"},
	{"m", "minute"},
	{"h", "hour"},
	{"d", "day"},
	{"M", "month"},
	{"y", "year"},
}

func GetTimeSevere(section string, key string) time.Duration {
	str := GetStringSevere(section, key)
	var timeAll int64 = 0
	n := len(str)
	for i := 0; i < n; {
		if str[i] >= '0' && str[i] <= '9' {
			j := i
			var num int64 = 0
			for ; j < n && str[j] >= '0' && str[j] <= '9'; j++ {
				num = num*10 + int64(str[j]-'0')
			}
			i = j
			for ; j < n && (str[j] >= 'A' && str[j] <= 'Z' || str[j] >= 'a' && str[j] <= 'z'); j++ {
			}
			unit := str[i:j]
			found := false
			for k := 0; k < len(names) && !found; k++ {
				num = num * multiple[k]
				for _, l := range names[k] {
					if l == unit {
						timeAll += num
						found = true
						break
					}
				}
			}
			if !found {
				panic("Unknow unit " + unit)
			}
			i = j
		} else {
			panic("Unknow char " + string(str[i]))
		}
	}
	return time.Duration(timeAll)
}
