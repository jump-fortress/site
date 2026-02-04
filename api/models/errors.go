package models

import (
	"fmt"

	"github.com/danielgtaylor/huma/v2"
)

func WrapDBErr(err error) error {
	return huma.Error500InternalServerError("oops.. i messed up somewhere. please report this if it keeps happening!", err)
}

func SessionErr() error {
	return huma.Error401Unauthorized("session required")
}

func WrapTempusErr(err error) error {
	return huma.Error503ServiceUnavailable("something went wrong connecting to Tempus. if Tempus isn't down and it keeps happening, please report this!", err)
}

func PlayerClassErr(playerClass string) error {
	return huma.Error400BadRequest(fmt.Sprintf("%s isn't a player class", playerClass))
}

func DivErr(div string) error {
	return huma.Error400BadRequest(fmt.Sprintf("%s isn't a div", div))
}

func EventKindErr(kind string) error {
	return huma.Error400BadRequest(fmt.Sprintf("%s isn't an event kind", kind))
}

func InvalidDurationErr(duration float64) error {
	return huma.Error400BadRequest(fmt.Sprintf("%.3f is negative or over 10 hours. please check your time again", duration))
}
