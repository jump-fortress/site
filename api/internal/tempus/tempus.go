package tempus

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/jump-fortress/site/models"
)

func GetPR(map_name string, tempusID int64, tempusClassID int64) (*models.TempusTime, error) {
	url := fmt.Sprintf("https://tempus2.xyz/api/v0/maps/name/%s/zones/typeindex/map/1/records/player/%d/%d", map_name, tempusID, tempusClassID)

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusTimeInfo := &models.TempusTimeInfo{}
	if err := json.Unmarshal(body, &tempusTimeInfo); err != nil {
		return nil, err
	}

	return &tempusTimeInfo.PR, nil
}

func GetPlayerInfo(tempusID int64) (*models.TempusPlayer, error) {
	url := fmt.Sprintf("https://tempus2.xyz/api/v0/players/id/%d/stats", tempusID)

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusPlayerInfo := &models.TempusPlayerInfo{}
	if err := json.Unmarshal(body, &tempusPlayerInfo); err != nil {
		return nil, err
	}

	return &tempusPlayerInfo.PlayerInfo, nil
}

func GetMaps() (*[]models.TempusMap, error) {
	url := "https://tempus2.xyz/api/v0/maps/detailedList"

	response, err := retryablehttp.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	tempusMaps := &[]models.TempusMap{}
	if err := json.Unmarshal(body, &tempusMaps); err != nil {
		return nil, err
	}

	return tempusMaps, nil
}
