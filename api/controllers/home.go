package controllers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/oneplanet/corona-backend/api/utils"
	"net/http"
	"net/url"
)

func GetDataForHomepage(c *gin.Context) {
	apiUrl, err := url.Parse(DataApiUrl)
	if err != nil {
		utils.Respond(c, utils.Message(false, "Data Api Error"))
		return
	}
	httpClient := &http.Client{}

	apiClient := Client{
		BaseURL:    apiUrl,
		UserAgent:  "",
		httpClient: httpClient,
	}

	allData, err := apiClient.AllCountries()
	if err != nil {
		utils.Respond(c, utils.Message(false, "Data Api Parsing Error"))
		return
	}

	resp := utils.Message(true, "success")
	resp["data"] = allData
	utils.Respond(c, resp)
}
