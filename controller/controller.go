// Package controller has some methods to controll, what data send to database queries
package controller

import (
	"math/rand"
	"time"

	"github.com/Tournament/entity"
	"github.com/Tournament/errors"
)

// PlayerDB is an interface for database, that used to controll player activity methods
type PlayerDB interface {
	GetPlayer(string) (entity.Player, error)
	CreatePlayer(string, int) (entity.Player, error)
	UpdatePlayer(string, int) error
	DeletePlayer(string) error
}

// TourDB is an interface for database, that used to controll tournament activity methods
type TourDB interface {
	CreateTournament(string, int) error
	GetTournamentState(string) (bool, error)
	GetWinner(string) (entity.Winners, error)
	UpdateTourAndPlayer(string, string) error
	CloseTournament(string) error
	GetParticipants(string) ([]string, error)
	SetTournamentWinner(string, entity.Winner) error
	DeleteTournament(string) error
}

// Game methods controll game activity
type Game struct {
	PDB PlayerDB
	TDB TourDB
}

// Fund controlls funding player
func (g Game) Fund(id string, points int) (entity.Player, error) {
	if points < 0 {
		return entity.Player{}, errors.Error{Code: errors.NegativePointsNumberError, Message: "fund: cannot fund negative number of points"}
	}
	if id == "" {
		return entity.Player{}, errors.Error{Code: errors.NotFoundError, Message: "fund: id must be not nil"}
	}
	_, err := g.PDB.GetPlayer(id)
	if err != nil {
		return g.PDB.CreatePlayer(id, points)
	}
	return entity.Player{}, g.PDB.UpdatePlayer(id, points)
}

// Take controlls taking points
func (g Game) Take(id string, points int) error {
	if points < 0 {
		return errors.Error{Code: errors.NegativePointsNumberError, Message: "take: cannot take negative number of points"}
	}
	if id == "" {
		return errors.Error{Code: errors.NotFoundError, Message: "take: id must be not nil"}
	}
	return g.PDB.UpdatePlayer(id, -1*points)
}

// Balance controlls getting actual player balance
func (g Game) Balance(id string) (entity.Player, error) {
	if id == "" {
		return entity.Player{}, errors.Error{Code: errors.NotFoundError, Message: "balance: id must be not nil"}
	}
	return g.PDB.GetPlayer(id)
}

// AnnounceTournament controlls announcing tournament
func (g Game) AnnounceTournament(id string, deposit int) error {
	if deposit <= 0 {
		return errors.Error{Code: errors.NegativeDepositError, Message: "announce: cannot create tournament with non positive deposite, id: " + id}
	}
	if id == "" {
		return errors.Error{Code: errors.NotFoundError, Message: "announce: id must be not nil"}
	}
	return g.TDB.CreateTournament(id, deposit)
}

// JoinTournament controlls joining player to tournament
func (g Game) JoinTournament(tourID, playerID string) error {
	if tourID == "" {
		return errors.Error{Code: errors.NotFoundError, Message: "joining tournament: id must be not nil"}
	}
	if playerID == "" {
		return errors.Error{Code: errors.NotFoundError, Message: "joining tournament: id must be not nil"}
	}
	isOpen, err := g.TDB.GetTournamentState(tourID)
	if err != nil {
		return err
	}
	if !isOpen {
		return errors.Error{Code: errors.ClosedTournamentError, Message: "joining tournament: cannot join to closed tournament"}
	}
	p, err := g.TDB.GetParticipants(tourID)
	if err != nil {
		return err
	}
	for i := range p {
		if p[i] == playerID {
			return errors.Error{Code: errors.DuplicatedIDError, Message: "joining tournament: cannot join to one tournament twice, playerID: " + playerID}
		}
	}
	return g.TDB.UpdateTourAndPlayer(tourID, playerID)
}

// Results controls getting results from tournament
// If tournament is opened, it closes it
func (g Game) Results(tourID string) (entity.Winners, error) {
	if tourID == "" {
		return entity.Winners{}, errors.Error{Code: errors.NotFoundError, Message: "results: id must be not nil"}
	}
	isOpen, err := g.TDB.GetTournamentState(tourID)
	if err != nil {
		return entity.Winners{}, err
	}
	if isOpen {
		err = g.TDB.CloseTournament(tourID)
		if err != nil {
			return entity.Winners{}, err
		}
		winner, err := chooseWinner(g, tourID)
		if err != nil {
			return entity.Winners{}, err
		}
		err = g.TDB.SetTournamentWinner(tourID, winner)
		if err != nil {
			return entity.Winners{}, err
		}
	}
	return g.TDB.GetWinner(tourID)
}

func chooseWinner(g Game, tourID string) (entity.Winner, error) {
	p, err := g.TDB.GetParticipants(tourID)
	if err != nil {
		return entity.Winner{}, err
	}
	if len(p) == 0 {
		return entity.Winner{}, errors.Error{Code: errors.NoneParticipantsError, Message: "cannot choose winner: tournament has no participants, id: " + tourID}
	}
	rand.Seed(time.Now().UnixNano())
	win, err := g.PDB.GetPlayer(p[rand.Intn(len(p))])
	if err != nil {
		return entity.Winner{}, err
	}
	return entity.Winner{ID: win.ID, Points: win.Points}, nil
}

// DeletePlayer controls deleting player
func (g Game) DeletePlayer(playerID string) error {
	if playerID == "" {
		return errors.Error{Code: errors.NotFoundError, Message: "deleting player: cannot delete player with empty id"}
	}
	return g.PDB.DeletePlayer(playerID)
}

// DeleteTournament controls deleting tournament
func (g Game) DeleteTournament(tourID string) error {
	if tourID == "" {
		return errors.Error{Code: errors.NotFoundError, Message: "deleting tournament: cannot delete tournament with empty id"}
	}
	return g.TDB.DeleteTournament(tourID)
}
