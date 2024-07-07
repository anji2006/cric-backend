package services

import (
	"context"
	"errors"

	"cric.com/backend/models/request"
	"cric.com/backend/models/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TeamService interface {
	CreateTeam(*request.Team) error
	UpdateTeam(*request.Team) error
	GetAllTeams() (*response.TeamAllRes, error)
	GetTeam(string) (*response.Team, error)
}

type TeamServiceImpl struct {
	teamcollection *mongo.Collection
	ctx            context.Context
}

func NewTeamService(teamcollection *mongo.Collection, ctx context.Context) TeamService {
	return &TeamServiceImpl{
		teamcollection: teamcollection,
		ctx:            ctx,
	}
}

func (ts *TeamServiceImpl) CreateTeam(team *request.Team) error {
	_, err := ts.teamcollection.InsertOne(ts.ctx, team)
	return err
}

func (ts *TeamServiceImpl) GetTeam(id string) (*response.Team, error) {
	var team *response.Team

	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	query := bson.D{
		bson.E{Key: "_id", Value: _id},
	}

	err = ts.teamcollection.FindOne(ts.ctx, query).Decode(&team)

	return team, err
}

func (ts *TeamServiceImpl) GetAllTeams() (*response.TeamAllRes, error) {
	var teams []response.Team
	cursor, err := ts.teamcollection.Find(ts.ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ts.ctx) {
		var team response.Team
		err := cursor.Decode(&team)

		if err != nil {
			return nil, err
		}

		teams = append(teams, team)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	cursor.Close(ts.ctx)

	if len(teams) == 0 {
		return nil, errors.New("Documents not found")
	}

	teamsRes := response.TeamAllRes{
		Data:  teams,
		Total: len(teams),
	}
	// teamsRes.Data = *teams[]

	return &teamsRes, nil
}

func (ts *TeamServiceImpl) UpdateTeam(team *request.Team) error {
	filter := bson.D{{Key: "name", Value: team.Name}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "captain", Value: team.Captain},
		}},
	}

	result, _ := ts.teamcollection.UpdateOne(ts.ctx, filter, update)

	if result.MatchedCount != 1 {
		return errors.New("No matched document found for update..")
	}

	return nil
}
