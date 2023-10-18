package iplocation

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	fastahKey      string
	phonePrefixMap = map[string]string{}
)

type LocationPhone struct {
	CountryName string `json:"countryName"`
	CountryCode string `json:"countryCode"`
	DialingCode string `json:"dialingCode"`
}

type Location struct {
	IP              string `json:"ip"`
	IsEuropeanUnion bool   `json:"isEuropeanUnion"`
	L10N            struct {
		CurrencyName string   `json:"currencyName"`
		CurrencyCode string   `json:"currencyCode"`
		LangCodes    []string `json:"langCodes"`
	} `json:"l10n"`
	LocationData struct {
		CountryName    string  `json:"countryName"`
		CountryCode    string  `json:"countryCode"`
		CityName       string  `json:"cityName"`
		CityGeonamesID int     `json:"cityGeonamesId"`
		Lat            float64 `json:"lat"`
		Lng            float64 `json:"lng"`
		Tz             string  `json:"tz"`
		ContinentCode  string  `json:"continentCode"`
	} `json:"locationData"`
}

func Init(key string, codeFile string) (err error) {
	SetFastahKey(key)
	csvFile, err := os.Open(codeFile)

	if err != nil {
		return err
	}

	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.Comma = '\t'
	// 第0列是国家名称，第1列是国家代码，第2列手机号前缀
	reader.FieldsPerRecord = 3
	csvData, err := reader.ReadAll()
	if err != nil {
		return
	}
	for _, each := range csvData {
		phonePrefixMap[each[1]] = each[2]
	}
	return
}

func GetLocationPhone(ip string) (loc *LocationPhone, err error) {
	loc = &LocationPhone{}
	rawLoc, err := GetRawLocation(ip)
	if err != nil {
		return
	}
	loc.CountryName = rawLoc.LocationData.CountryName
	loc.CountryCode = rawLoc.LocationData.CountryCode
	loc.DialingCode = phonePrefixMap[rawLoc.LocationData.CountryCode]
	if loc.DialingCode == "" {
		loc.DialingCode = "1"
	}
	loc.DialingCode = "+" + loc.DialingCode
	return
}

func SetFastahKey(key string) {
	fastahKey = key
}

func GetRawLocation(ip string) (loc *Location, err error) {

	// curl \
	//         -X GET "https://ep.api.getfastah.com/whereis/v1/json/127.0.0.1" \
	//         -H "Fastah-Key: 1"

	req, err := http.NewRequest("GET", "https://ep.api.getfastah.com/whereis/v1/json/"+ip, nil)
	if err != nil {
		return
	}
	req.Header.Set("Fastah-Key", fastahKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	loc = &Location{}
	err = json.Unmarshal(body, loc)
	return
}
