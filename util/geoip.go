package util

import (
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func GetIpLocation(ip string) map[string]interface{} {
	db, err := geoip2.Open("static/geo/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *geoip2.Reader) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	ipParse := net.ParseIP(ip)
	record, err := db.City(ipParse)
	if err != nil {
		log.Fatal(err)
	}

	geoInfo := make(map[string]interface{})
	geoInfo["continent"] = record.Continent.Names["zh-CN"]      // 洲名
	geoInfo["country"] = record.Country.Names["zh-CN"]          // 国家
	geoInfo["iso_code"] = record.Country.IsoCode                // 国家简码
	geoInfo["time_zone"] = record.Location.TimeZone             // 时区
	geoInfo["province"] = record.Subdivisions[0].Names["zh-CN"] // 州/省
	geoInfo["city"] = record.City.Names["zh-CN"]                // 城市
	geoInfo["longitude"] = record.Location.Longitude            // 经度
	geoInfo["latitude"] = record.Location.Latitude              // 维度
	return geoInfo
}
