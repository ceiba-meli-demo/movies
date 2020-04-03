package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ceiba-meli-demo/movies/application/commands"
	"github.com/ceiba-meli-demo/movies/infrastructure/utils/logger"
	"github.com/ceiba-meli-demo/movies/infrastructure/utils/rest_errors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

const (
	paramUser        = "dni"
	portUsersService = "USERS_PORT"
	hostUsersService = "USERS_HOST"
)

type user struct {
	DNI            int64  `json:"dni"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	QuantityMovies int8   `json:"quantity_movies"`
}

type Resp struct {
	Msg  string `json:"msg"`
	Done bool   `json:"done"`
}

var (
	usersRestClient *resty.Client
)

func UserRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := verifyRequest(c.Request); err != nil {
			c.JSON(err.Status(), err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func verifyRequest(r *http.Request) rest_errors.RestErr {
	if r == nil {
		return nil
	}
	var movieCommand commands.MovieCommand
	bodyBytes, _ := ioutil.ReadAll(r.Body)

	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bodyBytes, &movieCommand); err != nil {
		logger.Error("error trying to decode body request", err)
		return rest_errors.NewBadRequestError("invalid json")
	}
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	userID := strings.TrimSpace(strconv.Itoa(movieCommand.Dni))
	if userID == "0" {
		return rest_errors.NewBadRequestError("invalid json: includes user dni")
	}
	_, errRequest := getUser(userID)
	if errRequest != nil {
		if errRequest.Status() == http.StatusNotFound {
			return nil
		}
		return errRequest
	}
	return nil
}

func getUser(userID string) (*user, rest_errors.RestErr) {
	var u user

	validateClient()
	resp, err := usersRestClient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&u).
		Get(fmt.Sprintf("/users/%s", userID))
	if err != nil {
		logger.Error(err.Error(), err)
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}
	if resp.StatusCode() > 299 {
		logger.Error(fmt.Sprintf("error when trying to get user with status code: %d", resp.StatusCode()), err)
		body := string(resp.Body())
		return nil, rest_errors.SToE(body)
	}
	return &u, nil
}

func UpdateUser(userID int) (*Resp, rest_errors.RestErr) {
	var r Resp

	validateClient()
	resp, err := usersRestClient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&r).
		Put(fmt.Sprintf("/users/%d", userID))
	if err != nil {
		logger.Error(err.Error(), err)
		return nil, rest_errors.NewInternalServerError(err.Error(), err)
	}
	if resp.StatusCode() > 299 {
		logger.Error(fmt.Sprintf("error when trying to update user with status code: %d", resp.StatusCode()), err)
		body := string(resp.Body())
		return nil, rest_errors.SToE(body)
	}
	return &r, nil
}

func validateClient() {
	if usersRestClient == nil {
		port := fmt.Sprintf(":%s", os.Getenv(portUsersService))
		host := os.Getenv(hostUsersService)
		usersRestClient = resty.New().SetHostURL(fmt.Sprintf("http://%s%s", host, port)).SetTimeout(200 * time.Millisecond)
	}
}
