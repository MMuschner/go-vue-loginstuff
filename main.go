package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"log/syslog"
	"loginstuff/database"
	"loginstuff/routes"
	"os"
)

func logInit() {
	// Setting up logger
	logLevel := "error"
	syslogger, err := syslog.New(syslog.LOG_INFO|syslog.LOG_DEBUG|syslog.LOG_WARNING|syslog.LOG_ERR, "TEKKSTATT")
	syslogwriter := zerolog.SyslogLevelWriter(syslogger)
	multi := zerolog.MultiLevelWriter(syslogwriter, os.Stdout)
	log.Logger = zerolog.New(multi).With().Timestamp().Caller().Logger()
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)
}

func main() {
	// Initialising logger
	logInit()

	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8000")
}
