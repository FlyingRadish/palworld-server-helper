package pal

import (
	"fmt"
	"pal-server-helper/pal/rcn"
	"strings"
	"time"
)

type PalPlayer struct {
	Name    string `json:"name"`
	Uid     string `json:"uid"`
	SteamId string `json:"steamId"`
}

var OnlinePlayers []PalPlayer

func MonitorPlayers(checkInternal int) {
	OnlinePlayers = nil
	client := rcn.GetRCNClient()
	ticker := time.NewTicker(time.Duration(checkInternal) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			res, err := client.ShowPalyers()
			if err != nil {
				fmt.Println("ShowPlayers failed", err)
			}
			OnlinePlayers = parsePlayerInfos(res)
		}
	}
}

func parsePlayerInfos(res string) []PalPlayer {
	lines := strings.Split(res, "\n")
	players := make([]PalPlayer, 0)

	if len(lines) <= 1 {
		return players
	}

	for i := 1; i < len(lines); i++ {
		if len(lines[i]) == 0 {
			continue
		}
		playerLine := strings.Split(lines[i], ",")
		player := PalPlayer{
			Name:    removeExtraTerminator(playerLine[0]),
			Uid:     removeExtraTerminator(playerLine[1]),
			SteamId: removeExtraTerminator(playerLine[2]),
		}
		players = append(players, player)
	}
	return players
}

func removeExtraTerminator(data string) string {
	return strings.ReplaceAll(data, "\u0000", "")
}
