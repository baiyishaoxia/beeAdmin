package geoip

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/oschwald/geoip2-golang"
	"net"
)

var (
	cityIpReader *geoip2.Reader
	err          error
	GeoIP        Geo
)

type Geo struct {
	IP          string
	City        string
	Subdivision string
	Country     string
	CountryCode string
	TimeZone    string
	Latitude    float64
	Longitude   float64
}

func InitGeoIP(ctx *context.Context) Geo {
	ip := ctx.Input.IP()
	if beego.BConfig.RunMode != beego.PROD {
		ip = "183.15.88.228"
	}
	return NewGeoIP(ip)
}

func NewGeoIP(ip string) Geo {
	GeoIP.IP = ip
	ipNet := net.ParseIP(ip)
	if cityRecord, err := cityIpReader.City(ipNet); err == nil {
		city := cityRecord.City.Names
		subdivision := cityRecord.Subdivisions
		country := cityRecord.Country.Names

		if city != nil && len(city["en"]) > 0 {
			GeoIP.City = city["en"]
		}

		if subdivision != nil {
			subNames := subdivision[0].Names
			if len(subNames) > 0 && len(subNames["en"]) > 0 {
				GeoIP.Subdivision = subNames["en"]
			}
		}

		if country != nil && len(country["en"]) > 0 {
			GeoIP.Country = country["en"]
		}

		GeoIP.CountryCode = cityRecord.Country.IsoCode
		GeoIP.TimeZone = cityRecord.Location.TimeZone
		GeoIP.Latitude = cityRecord.Location.Latitude
		GeoIP.Longitude = cityRecord.Location.Longitude
	}

	return GeoIP
}

func RegisterCityIpReader() {
	cityPath := beego.AppConfig.String("maxminddbCity")
	cityIpReader, err = geoip2.Open(cityPath)

	if err != nil {
		panic(err)
	}
}
