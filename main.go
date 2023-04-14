package main

import (
	"honoka-chan/config"
	"honoka-chan/handler"
	"honoka-chan/middleware"
	_ "honoka-chan/tools"
	"honoka-chan/xclog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	xclog.Init(config.Conf.Log.LogDir, "", config.Conf.Log.LogLevel, config.Conf.Log.LogSave)
}

func main() {
	// Router
	r := gin.Default()
	r.Static("/static", "static")
	r.LoadHTMLGlob("static/*.tmpl")

	// /
	r.Any("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello, world!")
	})

	// Private APIs
	v1 := r.Group("v1")
	{
		v1.GET("/basic/getcode", handler.GetCodeHandler)
		v1.POST("/account/active", handler.ActiveHandler)
		v1.POST("/basic/publickey", handler.PublicKeyHandler)
		v1.POST("/basic/handshake", handler.HandshakeHandler)
		v1.POST("/account/initialize", handler.InitializeHandler)
		v1.POST("/account/login", handler.AccountLoginHandler)
		v1.POST("/account/loginauto", handler.LoginAutoHandler)
		v1.POST("/basic/loginarea", handler.LoginAreaHandler)
		v1.POST("/account/reportRole", handler.ReportRoleHandler)
		v1.POST("/basic/getProductList", handler.GetProductListHandler)
	}
	r.GET("/agreement/all", handler.AgreementHandler)
	r.GET("/integration/appReport/initialize", handler.ReportApp)
	r.POST("/report/ge/app", handler.ReportLog)
	// Private APIs

	// Server APIs
	m := r.Group("main.php").Use(middleware.Common)
	{
		m.POST("/login/authkey", middleware.AuthKey, handler.AuthKey)
		m.POST("/login/login", middleware.Login, handler.Login)
		m.POST("/user/userInfo", handler.UserInfoHandler)
		m.POST("/gdpr/get", handler.GdprHandler)
		m.POST("/personalnotice/get", handler.PersonalNoticeHandler)
		m.POST("/tos/tosCheck", handler.TosCheckHandler)
		m.POST("/download/additional", handler.DownloadAdditionalHandler)
		m.POST("/download/batch", handler.DownloadBatchHandler)
		m.POST("/download/getUrl", handler.DownloadUrlHandler)
		m.POST("/download/update", handler.DownloadUpdateHandler)
		m.POST("/download/event", handler.DownloadEventHandler)
		m.POST("/lbonus/execute", handler.LBonusExecuteHandler)
		m.POST("/api", handler.ApiHandler)
		m.POST("/announce/checkState", handler.AnnounceCheckStateHandler)
		m.POST("/scenario/startup", handler.ScenarioStartupHandler)
		m.POST("/scenario/reward", handler.ScenarioRewardHandler)
		m.POST("/user/setNotificationToken", handler.SetNotificationTokenHandler)
		m.POST("/user/changeNavi", handler.SetNotificationTokenHandler)
		m.POST("/event/eventList", handler.EventListHandler)
		m.POST("/payment/productList", handler.ProductListHandler)
		m.POST("/live/partyList", handler.PartyListHandler)
		m.POST("/live/play", handler.PlayLiveHandler)
		m.POST("/live/preciseScore", handler.PlayScoreHandler)
		m.POST("/live/reward", handler.PlayRewardHandler)
		m.POST("/live/gameover", handler.GameOverHandler)
		m.POST("/unit/setDisplayRank", handler.SetDisplayRankHandler)
		m.POST("/unit/favorite", handler.SetDisplayRankHandler)
		m.POST("/unit/deck", handler.SetDeckHandler)
		m.POST("/unit/deckName", handler.SetDeckNameHandler)
		m.POST("/subscenario/startup", handler.SubScenarioStartupHandler)
		m.POST("/subscenario/reward", handler.SubScenarioStartupHandler)
		m.POST("/album/seriesAll", handler.AlbumSeriesAllHandler)
		m.POST("/notice/noticeFriendVariety", handler.NoticeFriendVarietyHandler)
		m.POST("/notice/noticeFriendGreeting", handler.NoticeFriendGreetingHandler)
		m.POST("/notice/noticeUserGreetingHistory", handler.NoticeUserGreetingHandler)
		m.POST("/multiunit/scenarioStartup", handler.MultiUnitStartUpHandler)
	}
	r.GET("/webview.php/announce/index", handler.AnnounceIndexHandler)
	// Server APIs

	// Web
	// Manga
	r.GET("/manga", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "manga.tmpl", gin.H{})
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
