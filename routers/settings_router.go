package routers

import (
	"gvb_server/api"

	"github.com/gin-gonic/gin"
)

func SettingsRouter(router *gin.Engine) {
	settingsApi := api.ApiGroupApp.SettingsApi

	router.GET("", settingsApi.SettingsInfoView)
}
