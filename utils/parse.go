package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func ParsePostReq(c *gin.Context, res interface{}) error {
	bytes, _ := ioutil.ReadAll(c.Request.Body)
	log.Printf("%+v", string(bytes))
	err := json.Unmarshal(bytes, &res)
	if err != nil {
		log.Printf("failed to unmarshal req")
		return err
	}
	return nil
}
