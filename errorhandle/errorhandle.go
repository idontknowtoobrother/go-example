package errorhandle

import (
	"errors"
)

type GameIdentifier struct {
	SteamHex  string
	XboxID    string
	DiscordID string
}

type XPlayer struct {
	Source     int
	Identifier GameIdentifier
	Name       string
	Cash       float64
}

var Players = []XPlayer{
	{
		Source: 1,
		Identifier: GameIdentifier{
			SteamHex:  "steam_hex_1",
			XboxID:    "xbox_id_1",
			DiscordID: "discord_id_1",
		},
		Name: "Player 1",
		Cash: 100.00,
	},
	{
		Source: 2,
		Identifier: GameIdentifier{
			SteamHex:  "steam_hex_2",
			DiscordID: "discord_id_2",
		},
		Name: "Player 2",
		Cash: 200.00,
	},
}

func GetPlayerById(source int) (*XPlayer, error) {

	if source == 0 {
		return nil, errors.New("WAT DA FOK SERVER CAN'T BE PLAYER!!")
	}

	if source >= 1 && source < len(Players) {
		return &Players[source-1], nil
	}

	return nil, errors.New("Not found player")
}
