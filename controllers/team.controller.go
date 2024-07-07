package controllers

import (
	"net/http"

	"cric.com/backend/models/request"
	"cric.com/backend/services"
	"cric.com/backend/utils"
	"github.com/gin-gonic/gin"
)

type TeamController struct {
	TeamService services.TeamService
}

func New(teamservice services.TeamService) TeamController {
	return TeamController{
		TeamService: teamservice,
	}
}

func (tc *TeamController) CreateTeam(ctx *gin.Context) {

	var user request.Team
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {

		if ve := utils.FormValidations(err); ve != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": ve})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := tc.TeamService.CreateTeam(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully created team."})
}

func (tc *TeamController) GetTeam(ctx *gin.Context) {
	id := ctx.Param("id")
	team, err := tc.TeamService.GetTeam(id)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, team)
}

func (tc *TeamController) GetAllTeams(ctx *gin.Context) {
	teams, err := tc.TeamService.GetAllTeams()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, teams)

}

func (tc *TeamController) UpdateTeam(ctx *gin.Context) {

	var user request.Team
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {

		if ve := utils.FormValidations(err); ve != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": ve})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := tc.TeamService.UpdateTeam(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully Updated team."})
}

func (tc *TeamController) RegisterTeamRoutes(rg *gin.RouterGroup) {
	teamroute := rg.Group("/team")

	teamroute.POST("", tc.CreateTeam)
	teamroute.PUT("", tc.UpdateTeam)
	teamroute.GET("/all", tc.GetAllTeams)
	teamroute.GET("/:id", tc.GetTeam)
}
