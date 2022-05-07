package handler

import (
	"encoding/json"
	"github.com/ayang818/adx/dal/model"
	"github.com/ayang818/adx/dal/query"
	"github.com/ayang818/adx/model/dto"
	"github.com/ayang818/adx/pkg/mysql"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func CreateDemander(c *gin.Context) {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("%+v", string(bytes))
	req := dto.CreateDemanderReq{}
	err := json.Unmarshal(bytes, &req)
	if err != nil {
		log.Printf("failed to unmarshal req")
	}
	log.Printf("username: %v, password: %v", req.Username, req.Password)
	demander := model.DemanderInfo{Name: req.Name, Username: req.Username, Password: req.Password}
	db := mysql.GetDB()
	err = query.Use(db).DemanderInfo.WithContext(c).Create(&demander)
	if err != nil {
		log.Printf("create demander failed")
		c.JSON(200, gin.H{
			"message": "fail",
		})
		return
	}
	log.Printf("create demander suc")
	c.JSON(200, gin.H{
		"message": "suc",
	})
}
