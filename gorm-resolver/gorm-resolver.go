package main

// import
import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type User struct {
	// gorm.Model

	ID                   uint `gorm:"primarykey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	FirstName            string `gorm:"first_name"`
	LastName             string `gorm:"last_name"`
	Email                string `gorm:"email"`
	Type                 string `gorm:"type"`
	TripsCount           int64  `gorm:"trips_count"`
	DriversLicenseNumber string `gorm:"drivers_license_number"`
}

func (r *User) TableName() string {
	return "rideshare.users"
}

type Rider struct {
	ID                   uint `gorm:"primarykey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	FirstName            string `gorm:"first_name"`
	LastName             string `gorm:"last_name"`
	Email                string `gorm:"email"`
	Password             string `gorm:"password"`
	DriversLicenseNumber string `gorm:"drivers_license_number"`
}

func (r *Rider) TableName() string {
	return "rideshare.users"
}

type Vehicle struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"name"`
	Status    string `gorm:"status"`
}

func (r *Vehicle) TableName() string {
	return "rideshare.vehicles"
}

// main function
func main() {
	dsn_primary := fmt.Sprintf("host=localhost user=owner password=HSnDDgFtyW9fyFI dbname=rideshare_development port=54321 TimeZone=Asia/Jakarta")
	dsn_replica_1 := fmt.Sprintf("host=localhost user=owner password=HSnDDgFtyW9fyFI dbname=rideshare_development port=54322 TimeZone=Asia/Jakarta")
	dsn_replica_2 := fmt.Sprintf("host=localhost user=owner password=HSnDDgFtyW9fyFI dbname=rideshare_development port=54323 TimeZone=Asia/Jakarta")

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

	db, err := gorm.Open(postgres.Open(dsn_primary), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}
	db.Use(dbresolver.Register(dbresolver.Config{
		// use `db2` as sources, `db3`, `db4` as replicas
		Sources:  []gorm.Dialector{postgres.Open(dsn_primary)},
		Replicas: []gorm.Dialector{postgres.Open(dsn_replica_1), postgres.Open(dsn_replica_2)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
		// print sources/replicas mode in logger
		TraceResolverMode: true,
	}))

	// Read
	var user User
	var vehicle Vehicle
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(user User, wg *sync.WaitGroup) {
			defer wg.Done()
			db.First(&user, 13251893) // find product with integer primary key

		}(user, &wg)
	}

	wg.Wait()

	// db.Updates(&user)

	// delete vehicle
	db.Where("name = 'Ligier'").Delete(&vehicle)

	vehicle = Vehicle{
		Name:   "Ligier",
		Status: "draft",
	}

	db.Create(&vehicle)

	fmt.Println(vehicle)

}
