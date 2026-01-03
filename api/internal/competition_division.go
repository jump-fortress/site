package internal

import (
	"github.com/spiritov/jump/api/db/queries"
	"github.com/spiritov/jump/api/db/responses"
)

func getCompetitionDivisionResponse(cd queries.CompetitionDivision) responses.CompetitionDivision {
	return responses.CompetitionDivision{
		ID:            cd.ID,
		CompetitionID: cd.CompetitionID,
		Division:      cd.Division,
		Map:           cd.Map,
	}
}
