package main

// import 
import (
  "gorm.io/gorm"
  "fmt"
  "time"
  "database/sql"
  "gorm.io/driver/postgres"
)


// Struct
type Product struct {
  gorm.Model
  Code  string
  Price uint
  UserID uint
}

func (p *Product) TableName() string {
  return "products"
}
// ---------------------
type User struct {
  gorm.Model
  Name         string       `gorm:"default:Musa"`
  Email        *string
  Age          uint8        `gorm:"default:20"`
  Birthday     *time.Time
  MemberNumber sql.NullString
  ActivedAt    sql.NullTime
  Products      []Product
}

func (c *User) TableName() string {
  return "users"
}
// ---------------------


// main function
func main() {
  dsn := "host=localhost user=musaabdillah password=123456 dbname=dbmate_development port=5432 sslmode=disable TimeZone=Asia/Shanghai"
  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
  db.AutoMigrate(&User{}, &Product{})

  // set data 
  Email         := "musa.abdilla@ottodigital.id"
  Email2        := "musa@clappingape.com"
  Birthday      := time.Now()
  MemberNumber  := sql.NullString{String: "1234567", Valid: true}
  ActivedAt     := sql.NullTime{Time: time.Now(), Valid: true}

  // Create
  db.Create(&Product{Code: "D42", Price: 100, UserID: 2})

  // skip hook using &gorm.Session
  db.Session(&gorm.Session{SkipHooks: true}).Create(&User{
    Name: "Musa", 
    Email: &Email, 
    Age: 24, 
    Birthday: &Birthday, 
    MemberNumber: MemberNumber,
    ActivedAt: ActivedAt,
  })


  // Read
  var product Product
  db.First(&product, 1) // find product with integer primary key
  db.First(&product, "code = ?", "D42") // find product with code D42

  fmt.Println("========= begin product =========")
  fmt.Println("========= ID =========", product.ID)
  fmt.Println("========= CreatedAt =========", product.CreatedAt)
  fmt.Println("========= UpdatedAt =========", product.UpdatedAt)
  fmt.Println("========= DeletedAt =========", product.DeletedAt)
  fmt.Println("========= end product =========")

  var user User
  userResult := db.Find(&user, 2)
  fmt.Println("========= Result Error %v =========", userResult.Error)
  fmt.Println("========= Result RowsAffected %v =========", userResult.RowsAffected)
  fmt.Println("========= User %v =========", user)
  fmt.Println("========= User Relation With Product %v =========", user.Products)

  user2       := User{Name: "Musa", Email: &Email2, Age: 21}
  user2Result := db.Select("Name", "Email", "Age").Create(&user2)
  fmt.Println("========= user2Result Result Error %v =========", user2Result.Error)
  fmt.Println("========= user2Result Result RowsAffected %v =========", user2Result.RowsAffected)

  user3         := User{Email: &Email, Birthday: &Birthday, MemberNumber: MemberNumber}
  user3Result   := db.Omit("Name", "Age").Create(&user3)
  fmt.Println("========= user3Result Result Error %v =========", user3Result.Error)
  fmt.Println("========= user3Result Result RowsAffected %v =========", user3Result.RowsAffected)


  // Batch Insert 
  var users = []User{
    {Name: "Musa", Email: &Email, Age: 20, Birthday: &Birthday}, 
    {Name: "Musa Abdillah", Email: &Email, Age: 20, Birthday: &Birthday}, 
  }

  usersResult := db.Create(&users)
  fmt.Println("========= usersResult %v =========", usersResult) //result is pointer &{0xc0001f8120 <nil> 2 0xc000234340 0}
  fmt.Println("========= usersResult Result Error %v =========", usersResult.Error)
  fmt.Println("========= usersResult Result RowsAffected %v =========", usersResult.RowsAffected)

  for _, user := range users {
    fmt.Println("========= user ID %v =========", user.ID)
  }

   // Batch Insert With Size 
  var usersInBatches = []User{
    {Name: "Musa", Email: &Email, Age: 20, Birthday: &Birthday}, 
    {Name: "Musa Abdillah", Email: &Email, Age: 20, Birthday: &Birthday}, 
  }

  usersInBatchesResult := db.CreateInBatches(&usersInBatches, 1)
  fmt.Println("========= usersInBatches %v =========", usersInBatches) //result is pointer &{0xc0001f8120 <nil> 2 0xc000234340 0}
  fmt.Println("========= usersInBatches Result Error %v =========", usersInBatchesResult.Error)
  fmt.Println("========= usersInBatches Result RowsAffected %v =========", usersInBatchesResult.RowsAffected)

  for _, user := range usersInBatches {
    fmt.Println("========= user ID %v =========", user.ID)
  }

  // Create using map
  userMap := db.Model(&User{}).Create(map[string]interface{}{
    "Name": "Musa", "Email": &Email,
  })
  fmt.Println("========= userMap %v =========", userMap) //result is pointer &{0xc0001f8120 <nil> 2 0xc000234340 0}
  fmt.Println("========= userMap Result Error %v =========", userMap.Error)
  fmt.Println("========= userMap Result RowsAffected %v =========", userMap.RowsAffected)


  // Create using with default value
  userWithDefault := db.Create(&User{Email: &Email})
  fmt.Println("========= userWithDefault %v =========", userWithDefault) //result is pointer &{0xc0001f8120 <nil> 2 0xc000234340 0}
  fmt.Println("========= userWithDefault Result Error %v =========", userWithDefault.Error)
  fmt.Println("========= userWithDefault Result RowsAffected %v =========", userWithDefault.RowsAffected)







  // Update - update product's price to 200
  db.Model(&product).Update("Price", 200)

  // Update - update multiple fields
  db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
  db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

  // Delete - delete product
  db.Delete(&product, 9)

}


// hook
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  fmt.Println("========= User BeforeCreate =========", u)
  return nil
}