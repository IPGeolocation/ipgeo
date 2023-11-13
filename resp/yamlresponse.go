package resp

import (
	"encoding/json"
	"fmt"

	"github.com/IPGeolocation/ipgeo/models"
	"gopkg.in/yaml.v2"
)

func GetResponseInYAML(resp []byte) string {
	var geoResp models.GeolocationResponse
	if err := json.Unmarshal(resp, &geoResp); err != nil {
		fmt.Println("An error occured while converting response to yaml", err)
		return ""
	}
	return formatYAML(geoResp)
}

func formatYAML(resp models.GeolocationResponse) string {
	yamlStr, err := yaml.Marshal(resp)
	if err != nil {
		fmt.Println("An error occured while converting response to yaml", err)
		return ""
	}
	return string(yamlStr)
}
