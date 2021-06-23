// main.go
package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/unrolled/render"
    "database/sql/driver"
    "encoding/json"
    "errors"
)

type JSONB map[string]interface{}

func (a JSONB) Value() (driver.Value, error) {
    return json.Marshal(a)
}

func (a *JSONB) Scan(value interface{}) error {
    b, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }

    return json.Unmarshal(b, &a)
}

type Item struct {
    ID    int          `json:"id"`
    Attrs JSONB        `json:"attrs"`
}

type Product struct {
    ID      int             `json:"id"`
    Name    string          `json:"name"`
    Supplier []Supplier     `json:"suppliers"`
}

type Supplier struct {
    ID      int             `json:"id"`
    Name    string          `json:"name"`   
}

func main() {
    r := render.New(render.Options{
        IndentJSON: true,
    })

    router := gin.Default()

    // struct literal initialize
    suppliers := []Supplier{
        {ID: 1, Name: "Indofood", },
        {ID: 2, Name: "Garudafood", },
    }
    
    product := new(Product)
    product.ID = 1
    product.Supplier = suppliers

    // json b initialize
    item := new(Item)
    item.Attrs = JSONB{
        "users": []interface{}{
            product,
        },
    }

    router.GET("/", func(c *gin.Context) {
        r.JSON(c.Writer, http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
    })
    router.GET("/jsonb", func(c *gin.Context) {
        r.JSON(c.Writer, http.StatusOK, item)
    })

    router.Run("127.0.0.1:8080")
}