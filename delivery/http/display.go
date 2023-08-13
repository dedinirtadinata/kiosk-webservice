package http

import (
	"github.com/dedinirtadinata/kiosk-webservice/display"
	"github.com/dedinirtadinata/kiosk-webservice/display/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/**
 * Created by S DEDI NIRTADINATA on 10/08/23
 */

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
}

type displayHandler struct {
	usecase display.Usecase
}

func (h displayHandler) Create(context *gin.Context) {
	var request model.RequestCreate
	//var data models.ResponseModel
	err := context.ShouldBind(&request)

	if err != nil {
		context.JSON(http.StatusOK, ResponseSuccess{Status: "NOT_OK", Message: err.Error()})
		return
	}

	_, err = h.usecase.CreateData(request.Title, request.Description, request.BackgroundURL, request.CardImageURL, request.Url)

	if err != nil {
		context.JSON(http.StatusOK, ResponseSuccess{Status: "NOT_OK", Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, ResponseSuccess{Status: "OK"})
}

func (h displayHandler) GetDetail(context *gin.Context) {
	paramID := context.Params.ByName("id")
	id, _ := strconv.Atoi(paramID)

	data, err := h.usecase.GetDataById(id)

	if err != nil {
		context.JSON(http.StatusOK, ResponseSuccess{Status: "NOT_OK", Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, ResponseSuccess{Status: "OK", Data: data})
}

func (h displayHandler) GetAall(context *gin.Context) {
	result, err := h.usecase.GetAllData()

	if err != nil {
		context.JSON(http.StatusOK, ResponseSuccess{Status: "NOT_OK", Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, ResponseSuccess{Status: "OK", Data: result})
}

func (h displayHandler) Update(context *gin.Context) {
	var request model.RequestCreate
	//var data models.ResponseModel
	err := context.ShouldBind(&request)

	if err != nil {
		context.JSON(http.StatusOK, ResponseSuccess{Status: "NOT_OK", Message: err.Error()})
		return
	}

	paramID := context.Params.ByName("id")
	id, _ := strconv.Atoi(paramID)

	err = h.usecase.Update(id, request.Title, request.Description, request.BackgroundURL, request.CardImageURL, request.Url)

	if err != nil {
		context.JSON(http.StatusOK, ResponseSuccess{Status: "NOT_OK", Message: err.Error()})
		return
	}

	context.JSON(http.StatusOK, ResponseSuccess{Status: "OK", Message: ""})
}

func (h displayHandler) Delete(context *gin.Context) {
	paramID := context.Params.ByName("id")
	id, _ := strconv.Atoi(paramID)

	h.usecase.Delete(id)
}

func NewdisplayHandler(g *gin.Engine, uc display.Usecase) {
	handler := &displayHandler{
		usecase: uc,
	}

	ver := g.Group("v1")

	ver.POST("/display", handler.Create)
	ver.GET("/display/:id", handler.GetDetail)
	ver.GET("/display/available", handler.GetAall)
	ver.GET("/display", handler.GetAall)
	ver.PUT("/display/:id", handler.Update)
	ver.DELETE("/display/:id", handler.Delete)
}
