package InfluxDB

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/spf13/viper"
)

var client influxdb2.Client

func InitDB() {
	client = influxdb2.NewClient(viper.GetString("InfluxDB.HostUrl"), viper.GetString("InfluxDB.BYZY_TOKEN"))
}

func GetFluxWriteAPI() api.WriteAPIBlocking {
	return client.WriteAPIBlocking(viper.GetString("InfluxDB.FluxOrg"), viper.GetString("InfluxDB.FluxBucket"))
}
