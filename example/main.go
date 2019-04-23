package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/vinely/dids"
	"github.com/vinely/dids/ont"
	kvdb "github.com/vinely/kvdb"
)

var (
	didDB *kvdb.BoltDB
)

func init() {
	d, _ := kvdb.NewKVDataBase("bolt://did.db/did?count=50")
	didDB = d.(*kvdb.BoltDB)
}

func idget(c *gin.Context) {
	id := c.Param("id")
	info := didDB.Get(id)
	if !info.Result {
		c.JSON(http.StatusBadRequest, info.Error())
		return
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var v ont.Document
	fmt.Println(string(info.Data.([]byte)))
	err := json.Unmarshal(info.Data.([]byte), &v)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	v.ID = dids.ID(id)
	c.JSON(http.StatusOK, v)
}

func idnew(c *gin.Context) {
	v, _, err := ont.New()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if info := didDB.SetData(string(v.ID), v); !info.Result {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, v)
}

func idcreate(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "Not support now!")
}

func idlist(c *gin.Context) {
	strPage := c.Query("page")
	page, err := strconv.ParseUint(strPage, 10, 32)
	if err != nil {
		page = 0
	}
	data := didDB.List(uint(page), func(k, v []byte) *kvdb.KVResult {
		var doc ont.Document
		var json = jsoniter.ConfigCompatibleWithStandardLibrary
		err = json.Unmarshal(v, &doc)
		if err != nil {
			return &kvdb.KVResult{
				Result: false,
				Info:   err.Error(),
			}
		}
		return &kvdb.KVResult{
			Data:   doc,
			Result: true,
			Info:   "",
		}
	})
	c.JSON(http.StatusOK, gin.H{"Identities": data.Data})
}

func main() {
	engine := gin.Default()
	engine.GET("/api/v1/did/id/:id", idget)
	engine.GET("/api/v1/did/id", idlist)
	engine.GET("/api/v1/did/new", idnew)
	// engine.POST("/api/v1/did/id", idcreate)
	engine.Run(":8080")
}
