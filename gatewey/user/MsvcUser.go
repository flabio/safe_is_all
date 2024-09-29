package user

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/safe/utils"
)

// Handler for the  service
func MsvcUser(c *fiber.Ctx) error {
	pageParam := c.Query("page")
	log.Println(pageParam)
	id := c.Params(utils.ID)
	url := utils.MSVC_USER_URL

	if len(id) != 0 && url != "" {
		url = utils.MSVC_USER_URL + "/" + id
	} else {
		url = utils.MSVC_USER_URL + "?page=" + pageParam
	}

	req, err := http.NewRequest(c.Method(), url, bytes.NewBuffer(c.Body()))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(utils.FAILED_CREATE)
	}
	req.Header.Set(utils.AUTHORIZATION, c.Get(utils.AUTHORIZATION))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(utils.SERVICE_NOT_AVAILALE)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)

	return c.Send(respBody)
}
