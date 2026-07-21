package mapper

import (
	"IOTA_ZZ_server/model"
	"IOTA_ZZ_server/pkg/InfluxDB"
	"context"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func SaveASDatas(datas model.ASDatas) {
	// 定义一个数据的映射，用于存储键值对，表示数据的实际值
	FluxData := map[string]interface{}{
		"T":   datas.T,
		"H":   datas.H,
		"CO2": datas.CO2,
		"O2":  datas.O2,
	}
	// 定义一个标签的映射，用于存储键值对，表示标签的值
	tags := map[string]string{
		"Device": strconv.FormatInt(int64(datas.ASID), 10),
	}
	// 创建一个点对象，用于存储一个时间序列数据，指定测量值，标签，字段和时间戳
	point := write.NewPoint(viper.GetString("InfluxDB.FluxAS"), tags, FluxData, time.Now())
	// 调用写入API的WritePoint方法，将点对象写入InfluxDB，如果发生错误，打印错误并终止程序
	if err := InfluxDB.GetFluxWriteAPI().WritePoint(context.Background(), point); err != nil {
		log.Println("SaveASDatas WritePoint方法写入InfluxDB出错: ", err)
	}
}
