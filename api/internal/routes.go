package internal

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func registerOpenRoutes(internalApi *huma.Group) {
	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/{id}",
		OperationID: "get-player-by-id",
		Summary:     "Get a Player",
		Description: "get a player by ID",
		Tags:        []string{"Player"},
	}, HandleGetPlayer)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/all",
		OperationID: "get-all-players",
		Summary:     "Get all Players",
		Description: "get all players",
		Tags:        []string{"Player"},
	}, HandleGetAllPlayers)
}

func registerSessionRoutes(internalApi *huma.Group) {
	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players",
		OperationID: "get-player",
		Summary:     "Get the current session's Player",
		Description: "get the current session's player",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandleGetSelfPlayer)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/profile/{id}",
		OperationID: "get-player-profile-by-id",
		Summary:     "Get a Player Profile",
		Description: "get info for a player's profile by ID",
		Tags:        []string{"Player"},
	}, HandleGetPlayerProfile)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/preferredclass/{class}",
		OperationID: "set-player-preferredclass",
		Summary:     "Set Player's preferred class",
		Description: "set the current session player's preferred class by class name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfPreferredClass)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/preferredlauncher/{launcher}",
		OperationID: "set-player-preferredlauncher",
		Summary:     "Set Player's preferred rocket launcher",
		Description: "set the current session player's preferred rocket launcher by launcher name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfPreferredLauncher)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/tempusinfo/{tempus_id}",
		OperationID: "set-player-tempusinfo",
		Summary:     "Set a Player's own Tempus ID",
		Description: "set a player's own Tempus ID, country, and country code, found at https://tempus2.xyz",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfTempusInfo)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/steamtradetoken/{url}",
		OperationID: "set-player-steam-trade-token",
		Summary:     "Set a Player's own Steam trade token",
		Description: "set a player's own Steam trade token from their Steam Trade URL, found at https://steamcommunity.com/id/{steamid}/tradeoffers/privacy",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfSteamTradeToken)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/steamavatarurl",
		OperationID: "update-player-steam-avatar-url",
		Summary:     "Update a Player's own Steam avatar url",
		Description: "update a player's own Steam avatar from their Steam Profile",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfSteamAvatarUrl)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/requests/{request_type}/{request_string}",
		OperationID: "insert-player-request",
		Summary:     "Insert a player request ",
		Description: "send a request for division placement or name change",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfPlayerRequest)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/requests",
		OperationID: "get-player-requests",
		Summary:     "Get a player's own requests",
		Description: "Get all of a player's own requests",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandleGetSelfPlayerRequests)
}

func registerModeratorRoutes(moderatorApi *huma.Group) {
	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/all",
		OperationID: "get-all-full-players",
		Summary:     "Get full info on all Players",
		Description: "get full info on all players",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandleGetAllFullPlayers)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/displayname/{id}/{name}",
		OperationID: "update-player-displayname",
		Summary:     "Update a Player's display name",
		Description: "update a player's display name",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePutPlayerDisplayName)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/soldierdivision/{id}/{division}",
		OperationID: "update-player-soldierdivision",
		Summary:     "Update a Player's soldier division",
		Description: "update a player's soldier division",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePutPlayerSoldierDivision)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/demodivision/{id}/{division}",
		OperationID: "update-player-demodivision",
		Summary:     "Update a Player's demo division",
		Description: "update a player's demo division",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePutPlayerDemoDivision)
}
