package handler

import (
	"github.com/ayang818/adx/dal/model"
	"github.com/ayang818/adx/dal/query"
	"github.com/ayang818/adx/model/dto"
	"github.com/ayang818/adx/pkg/logger"
	"github.com/ayang818/adx/pkg/mysql"
	"github.com/ayang818/adx/utils"
	"github.com/gin-gonic/gin"
)

func CreateDemander(c *gin.Context) {
	req := dto.CreateDemanderReq{}
	utils.ParsePostReq(c, req)
	logger.CtxInfo(c, "username: %v, password: %v", req.Username, req.Password)
	demander := model.DemanderInfo{Name: req.Name, Username: req.Username, Password: req.Password}
	db := mysql.GetDB()
	err := query.Use(db).DemanderInfo.WithContext(c).Create(&demander)
	if err != nil {
		logger.CtxError(c, "create demander failed")
		utils.FailResp(c, "create demander failed")
		return
	}
	logger.CtxInfo(c, "create demander suc")
	resp := dto.BaseSucResp{Status: "suc"}
	utils.SucResp(c, &resp)
}
