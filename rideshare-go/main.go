package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type User struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

func (t *User) TableName() string {
	return "rideshare.users"
}

func (t *User) FullName() string {
	return t.FirstName + " " + t.LastName
}

type Location struct {
	Id      int64  `json:"id"`
	Address string `json:"address"`
}

func (t *Location) TableName() string {
	return "rideshare.locations"
}

type Trip struct {
	Id            int64     `json:"id"`
	TripRequestIdequestId int64     `json:"trip_request_id"`
	DriverId      int64     `json:"driver_id"`
	CompletedAt   time.Time `json:"completed_at"`
	Rating        int64     `json:"rating"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (t *Trip) TableName() string {
	return "rideshare.trips"
}

type TripRequest struct {
	Id              int64     `json:"id"`
	RiderId         int64     `json:"rider_id"`
	StartLocationId int64     `json:"start_location_id"`
	EndLocationId   int64     `json:"end_location_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (t *TripRequest) TableName() string {
	return "rideshare.trip_requests"
}

// parameter
type TripParameter struct {
	StartLocation string `json:"start_location"`
	EndLocation   string `json:"end_location"`
	RiderID       int64  `json:"rider_id"`
}

// constants
const (
	USER_RIDER  = "Rider"
	USER_DRIVER = "Driver"
)

func main() {

	dsn_primary := "host=localhost user=owner password=HSnDDgFtyW9fyFI dbname=rideshare_development port=5432 TimeZone=Asia/Jakarta"
	dsn_replica := "host=localhost user=owner password=HSnDDgFtyW9fyFI dbname=rideshare_development port=54321 TimeZone=Asia/Jakarta"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn_replica), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources: []gorm.Dialector{postgres.Open(dsn_primary)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
		// print sources/replicas mode in logger
		TraceResolverMode: true,
	}))

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/api/trips/:id", func(c *gin.Context) {

		var (
			trip Trip
		)

		db.Table(trip.TableName()).Where("id = ?", c.Param("id")).Find(&trip)
		c.JSON(200, gin.H{
			"data": trip,
		})
	})

	r.POST("/api/trips", func(c *gin.Context) {

		var (
			request       TripParameter
			tripRequest   TripRequest
			trip          Trip
			startLocation Location
			endLocation   Location
			user          User
			luckDriver    User
			rider         User
			driverIds     []int64
		)

		err = c.ShouldBindJSON(&request)
		if err != nil {
			logs.Info("Cannot bind json")
			c.JSON(http.StatusBadRequest, gin.H{
				"data": "-",
			})
		}

		// find or create start location
		db.Table(startLocation.TableName()).Where(Location{Address: request.StartLocation}).FirstOrCreate(&startLocation)

		// find or create end location
		db.Table(endLocation.TableName()).Where(Location{Address: request.EndLocation}).FirstOrCreate(&endLocation)

		// find or fail rider
		rider = User{
			Id: request.RiderID, Type: USER_RIDER,
		}
		db.Table(user.TableName()).Find(&rider)

		// if three condition are matched then insert
		if startLocation.Id != 0 && endLocation.Id != 0 && rider.Id != 0 {
			tripRequest = TripRequest{
				RiderId:         rider.Id,
				StartLocationId: startLocation.Id,
				EndLocationId:   endLocation.Id,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			}
			db.Table(tripRequest.TableName()).Create(&tripRequest)
		}

		// insert trip, search for best driver and assign as trip request
		// driver sample
		db.Table(user.TableName()).Select("id").Where("type = ?", USER_DRIVER).Scan(&driverIds)

		var (
			luckDriverId int64 = int64(rand.Intn(len(driverIds)))
		)

		db.Table(user.TableName()).Select("id, first_name, last_name").Where("id = ?", luckDriverId).Scan(&luckDriver)

		logs.Info("================LUCK DRIVER(HE/SHE) IS GO TO ================")
		logs.Info(luckDriver.FullName())

		trip = Trip{
			TripRequestId: tripRequest.Id,
			DriverId:      luckDriver.Id,
			CompletedAt:   time.Now().AddDate(0, 0, 1),
			Rating:        1,
			CreatedAt:     time.Now().Local(),
			UpdatedAt:     time.Now().Local(),
		}

		db.Table(trip.TableName()).Create(&trip)

		c.JSON(200, gin.H{
			"data": trip,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
