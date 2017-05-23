package perms

import (
	"encoding/json"
	"equestriaunleashed.com/eclipsingr/celesticore"
)

// Rank represents a rank on the server
type Rank struct {

	// RankName is the name of the rank.
	RankName string `json:"rank_name"`

	// RankOrderID is the order of the rank in the rank list.
	RankOrderID int64 `json:"rank_order_id"`

	// RankBase is the rank this rank is based off, nil or empty if no base rank.
	BaseRank string `json:"base_rank"`
}

func DeserializeRanks(inputJson string) []Rank {
	var ranks []Rank
	err := json.Unmarshal([]byte(inputJson), ranks)
	if err != nil {
		celesticore.LogError(err.Error(), "RankDeserializer")
	}
	return ranks
}
