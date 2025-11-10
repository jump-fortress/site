package responses

import database "github.com/spiritov/jump/api/db/queries"

type PlayerIDInput struct {
	ID int64 `path:"id" minimum:"1" doc:"player ID"`
}

type PlayerResponse struct {
	Body database.Player
}
