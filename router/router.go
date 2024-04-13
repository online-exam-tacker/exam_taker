package router

import (
	"hamideh/controller"

	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ExamRouter(examsController *controller.ExamsController) *gin.Engine {
	router := gin.Default()
	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})
	baseRouter := router.Group("/api")
	examsRouter := baseRouter.Group("/exams")
	examsRouter.POST("", examsController.Create)
	examsRouter.DELETE("/:examId", examsController.Delete)
	examsRouter.GET("", examsController.FindAll)
	examsRouter.GET("/:examId", examsController.FindById)
	examsRouter.PATCH("/:examId", examsController.Update)

	return router
}

func UserRouter(userController *controller.UsersController) *gin.Engine {
	router := gin.Default()

	baseRouter := router.Group("/api")
	usersRouter := baseRouter.Group("/users")
	usersRouter.POST("", userController.Create)

	return router
}
