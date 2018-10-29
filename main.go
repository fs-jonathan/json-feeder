package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Record struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Cost     int    `json:"cost"`
	Compare  int    `json:"compare"`
	Rate     int    `json:"rate"`
}

type Message struct {
	Message int `json:"message"`
}

type LineUser struct {
	UserId string `json:"lineUserId"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Debug mode
	e.Debug = true

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// dump middleware captures the request and response payload
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		log.Println(string(reqBody))
		log.Println(string(resBody))
	}))

	// Route => handler
	e.Static("/", "html")
	e.POST("/loginLiff", liffLogin)
	e.POST("/getJson", jsonWriter)

	// Start server from PORT number
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}

func jsonWriter(c echo.Context) error {
	// Receive message
	message := new(Message)

	if reqErr := c.Bind(message); reqErr != nil {
		return echo.NewHTTPError(http.StatusBadRequest, reqErr.Error())
	}

	// Return message
	s0 := []Record{Record{1, "本日（現時点まで）", "", 9, 8, 7}}
	s1 := append(s0, Record{2, "昨日", "先週の同じ曜日との比較", 344, 43243, 43})
	s2 := append(s1, Record{3, "今月（現時点まで）", "先週の同じ曜日との比較", 1, 0, 0})
	s3 := append(s2, Record{4, "先月", "先々月との比較", 93, 83, 72})
	records := append(s3, Record{5, "全期間", "", 20, 4, 434})

	return c.JSON(http.StatusOK, records)
}

func liffLogin(c echo.Context) error {
	// Received Line UserId
	user := new(LineUser)

	if err := c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	log.Println(user.UserId)

	message := Message{0}
	return c.JSON(http.StatusOK, message)
}
