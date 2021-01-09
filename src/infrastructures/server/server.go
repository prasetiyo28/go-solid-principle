package server

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/configs"
	loggers "github.com/prasetiyo28/go-solid-principle/src/infrastructures/logger"
	"github.com/prasetiyo28/go-solid-principle/src/infrastructures/middlewares"
	orm "github.com/prasetiyo28/go-solid-principle/src/infrastructures/orm/gorm"
	redis "github.com/prasetiyo28/go-solid-principle/src/infrastructures/redis"
	routers "github.com/prasetiyo28/go-solid-principle/src/interfaces/routes/v1"
)

func Serve() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}

	datasources := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		os.Getenv("DATABASE_USERNAME"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))

	// redisSources := fmt.Sprintf("%s:%s", os.Getenv("REDIS_PORT"), os.Getenv("REDIS_HOST"))
	db, errDb := sql.Open("mysql", datasources)
	if errDb != nil {
		log.Fatal(fmt.Sprintf("Error connecting db, %s", errDb.Error()))
	}
	db.SetMaxIdleConns(10)

	db.SetMaxOpenConns(100)
	orm, errOrm := orm.New(db)
	if errDb != nil {
		log.Panic(fmt.Sprintf("Error connecting db, %s", errOrm.Error()))
	}
	db.SetConnMaxLifetime(time.Hour)

	redisdb := redis.RedisInit()
	err := redisdb.Set("key", "value", 0).Err()
	if err != nil {
		log.Panic(fmt.Sprintf("Error connecting Redis, %s", err))

	}
	loggers.Init()
	e := New()
	authMiddle := middlewares.NewAuthMiddleware(orm)
	h := routers.NewHandler(orm)
	e.Use(authMiddle.MiddleGate())
	v1 := e.Group("/api")
	h.Register(v1)
	appHost := os.Getenv("APPLICATION_PORT")

	if "" == appHost {
		log.Panic("key of APPLICATION HOST are not define.")
	}

	e.Logger.Fatal(e.Start(":" + appHost))

}

func New() *echo.Echo {
	eHandler := middlewares.NewErrorMiddleware()
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.Validator = configs.NewValidator()
	e.HTTPErrorHandler = eHandler.HttpErrorHandler

	return e
}
