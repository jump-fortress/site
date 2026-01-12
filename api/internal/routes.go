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

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/competitions/all/monthly",
		OperationID: "get-all-monthly",
		Summary:     "Get all monthlies",
		Description: "get all monthlies that are visible",
		Tags:        []string{"Player"},
	}, HandleGetAllMonthlies)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/competitions/prizepool/{id}",
		OperationID: "get-prizepool",
		Summary:     "get prizepool",
		Description: "get all division prizepools for a competition",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostGetPrizepool)
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
		Method:      http.MethodPost,
		Path:        "/players/preferredclass/{class}",
		OperationID: "set-player-preferredclass",
		Summary:     "Set Player's preferred class",
		Description: "set the current session player's preferred class by class name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfPreferredClass)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/preferredlauncher/{launcher}",
		OperationID: "set-player-preferredlauncher",
		Summary:     "Set Player's preferred rocket launcher",
		Description: "set the current session player's preferred rocket launcher by launcher name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfPreferredLauncher)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/preferredmap/{map}",
		OperationID: "set-player-preferredmap",
		Summary:     "Set Player's preferred map",
		Description: "set the current session player's preferred map by map name",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfPreferredMap)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/tempusinfo/{tempus_id}",
		OperationID: "set-player-tempusinfo",
		Summary:     "Set a Player's own Tempus ID",
		Description: "set a player's own Tempus ID, country, and country code, found at https://tempus2.xyz",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfTempusInfo)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/steamtradetoken/{url}",
		OperationID: "set-player-steam-trade-token",
		Summary:     "Set a Player's own Steam trade token",
		Description: "set a player's own Steam trade token from their Steam Trade URL, found at https://steamcommunity.com/id/{steamid}/tradeoffers/privacy",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfSteamTradeToken)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/steamavatarurl",
		OperationID: "update-player-steam-avatar-url",
		Summary:     "Update a Player's own Steam avatar url",
		Description: "update a player's own Steam avatar from their Steam Profile",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfSteamAvatarUrl)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/requests/{request_type}/{request_string}",
		OperationID: "insert-player-request",
		Summary:     "Insert a player request ",
		Description: "send a request for division placement or name change",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSelfPlayerRequest)

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

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/maps",
		OperationID: "get-all-maps",
		Summary:     "Get all maps",
		Description: "get all maps",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandleGetAllMaps)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/maps/names",
		OperationID: "get-all-map-names",
		Summary:     "Get all map names",
		Description: "get all map names",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandleGetAllMapNames)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/times/{id}",
		OperationID: "submit-player-time",
		Summary:     "submit player time",
		Description: "submit a player time for a competition",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSubmitPlayerTime)

	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/times/{id}/{time}",
		OperationID: "submit-unverified-player-time",
		Summary:     "submit unverified player time",
		Description: "submit an unverified player time for a competition",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePostSubmitUnverifiedPlayerTime)
}

func registerConsultantRoutes(consultantApi *huma.Group) {
	huma.Register(consultantApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/all",
		OperationID: "get-all-full-players",
		Summary:     "Get full info on all Players",
		Description: "get full info on all players",
		Tags:        []string{"Consultant"},
		Security:    sessionCookieSecurityMap,
	}, HandleGetAllFullPlayers)

	huma.Register(consultantApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/players/requests/pending",
		OperationID: "get-all-pending-player-requests",
		Summary:     "Get all pending player requests",
		Description: "get all pending player requests",
		Tags:        []string{"Consultant"},
		Security:    sessionCookieSecurityMap,
	}, HandleGetAllPendingPlayerRequests)
}

func registerModeratorRoutes(moderatorApi *huma.Group) {

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/displayname/{id}/{name}",
		OperationID: "update-player-displayname",
		Summary:     "Update a Player's display name",
		Description: "update a player's display name",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostPlayerDisplayName)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/soldierdivision/{id}/{division}",
		OperationID: "update-player-soldierdivision",
		Summary:     "Update a Player's soldier division",
		Description: "update a player's soldier division",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostPlayerSoldierDivision)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/demodivision/{id}/{division}",
		OperationID: "update-player-demodivision",
		Summary:     "Update a Player's demo division",
		Description: "update a player's demo division",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostPlayerDemoDivision)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/requests/resolve/{id}",
		OperationID: "resolve-player-request",
		Summary:     "Resolve a player's request",
		Description: "resolve a player's request, marking it as not pending",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostResolvePlayerRequest)

	huma.Register(moderatorApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/times/{id}",
		OperationID: "verify-player-time",
		Summary:     "verify player time",
		Description: "verify a player time id",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostVerifyPlayerTime)
}

func registerAdminRoutes(adminApi *huma.Group) {
	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/maps",
		OperationID: "update-maps",
		Summary:     "Update map list",
		Description: "update the database's map list from Tempus data",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostUpdateMaps)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/competitions/create/monthly",
		OperationID: "create-monthly",
		Summary:     "Create monthly",
		Description: "Create a monthly competition",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostCreateMonthly)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/competitions/update/monthly",
		OperationID: "update-monthly",
		Summary:     "Update monthly",
		Description: "Update a monthly competition",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostUpdateMonthly)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodGet,
		Path:        "/competitions/all/monthly",
		OperationID: "get-all-full-monthly",
		Summary:     "Get all monthlies",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandleGetAllFullMonthlies)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/competitions/cancel/{id}",
		OperationID: "cancel-competition",
		Summary:     "cancel competition",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostCancelCompetition)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/competitions/prizepool/create",
		OperationID: "create-prizepool",
		Summary:     "create prizepool",
		Description: "create a competition division's prizepool",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostCreatePrizepool)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/competitions/prizepool/{id}",
		OperationID: "delete-prizepool",
		Summary:     "delete prizepool",
		Description: "reset a competition division's prizepool",
		Tags:        []string{"Admin"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostDeletePrizepool)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/times/{competition_id}/{player_id}/{run_time}",
		OperationID: "create-player-time",
		Summary:     "create player time",
		Description: "create a time for a competition",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostCreatePlayerTime)

	huma.Register(adminApi, huma.Operation{
		Method:      http.MethodPost,
		Path:        "/players/times/{id}",
		OperationID: "delete-player-time",
		Summary:     "delete player time",
		Description: "delete a time id",
		Tags:        []string{"Moderator"},
		Security:    sessionCookieSecurityMap,
	}, HandlePostDeletePlayerTime)
}
