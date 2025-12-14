	huma.Register(internalApi, huma.Operation{
		Method:      http.MethodPut,
		Path:        "/players/request/{request_type}/{body}",
		OperationID: "insert-player-request",
		Summary:     "Insert a player request ",
		Description: "send a request for division placement or name change",
		Tags:        []string{"Player"},
		Security:    sessionCookieSecurityMap,
		Middlewares: requireUserSessionMiddlewares,
	}, HandlePutSelfPlayerRequest)
}
