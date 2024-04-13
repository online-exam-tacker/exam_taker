package controller

import (
	"hamideh/data/request"
	"hamideh/data/response"
	"hamideh/helper"
	"hamideh/service"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type ExamsController struct {
	examsService service.ExamsService
}

func NewExamsController(service service.ExamsService) *ExamsController {
	return &ExamsController{
		examsService: service,
	}
}

// CreateTags		godoc
// @Summary			Create tags
// @Description		Save tags data in Db.
// @Param			tags body request.CreateTagsRequest true "Create tags"
// @Produce			application/json
// @Tags			tags
// @Success			200 {object} response.Response{}
// @Router			/tags [post]
func (controller *ExamsController) Create(ctx *gin.Context) {
	log.Info().Msg("create exams")
	createExamRequest := request.CreateExamRequest{}
	err := ctx.ShouldBindJSON(&createExamRequest)
	helper.ErrorPanic(err)

	controller.examsService.Create(createExamRequest)
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// // UpdateTags		godoc
// // @Summary			Update tags
// // @Description		Update tags data.
// // @Param			tagId path string true "update tags by id"
// // @Param			tags body request.CreateTagsRequest true  "Update tags"
// // @Tags			tags
// // @Produce			application/json
// // @Success			200 {object} response.Response{}
// // @Router			/tags/{tagId} [patch]
func (controller *ExamsController) Update(ctx *gin.Context) {
	log.Info().Msg("update exam")
	updateExamRequest := request.UpdateExamRequest{}
	err := ctx.ShouldBindJSON(&updateExamRequest)
	helper.ErrorPanic(err)

	examId := ctx.Param("examId")
	id, err := strconv.Atoi(examId)
	helper.ErrorPanic(err)
	updateExamRequest.Id = id

	controller.examsService.Update(updateExamRequest)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// // DeleteTags		godoc
// // @Summary			Delete tags
// // @Description		Remove tags data by id.
// // @Produce			application/json
// // @Tags			tags
// // @Success			200 {object} response.Response{}
// // @Router			/tags/{tagID} [delete]
func (controller *ExamsController) Delete(ctx *gin.Context) {
	log.Info().Msg("delete exam")
	examId := ctx.Param("examId")
	id, err := strconv.Atoi(examId)
	helper.ErrorPanic(err)
	controller.examsService.Delete(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   nil,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// // FindByIdTags 		godoc
// // @Summary				Get Single tags by id.
// // @Param				tagId path string true "update tags by id"
// // @Description			Return the tahs whoes tagId valu mathes id.
// // @Produce				application/json
// // @Tags				tags
// // @Success				200 {object} response.Response{}
// // @Router				/tags/{tagId} [get]
func (controller *ExamsController) FindById(ctx *gin.Context) {
	log.Info().Msg("findbyid tags")
	examId := ctx.Param("examId")
	id, err := strconv.Atoi(examId)
	helper.ErrorPanic(err)

	examResponse := controller.examsService.FindById(id)

	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   examResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

// // FindAllTags 		godoc
// // @Summary			Get All tags.
// // @Description		Return list of tags.
// // @Tags			tags
// // @Success			200 {obejct} response.Response{}
// // @Router			/tags [get]
func (controller *ExamsController) FindAll(ctx *gin.Context) {
	log.Info().Msg("findAll tags")
	examResponse := controller.examsService.FindAll()
	webResponse := response.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   examResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
