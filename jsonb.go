package main

import (
  "fmt"
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
    ID    int
    Attrs JSONB
}

func main() {
fmt.Println("Hello, playground")
item := new(Item)
item.Attrs = JSONB{
 "mars_response": map[string]interface{}{
    "response_body": map[string]interface{}{
       "count":1,
       "error_detail":
                          []interface{}{
                          map[string]interface{}{
                           "dsc":"2110",
                           "error_message":"Terdapat perbedaan harga",
                           "item": map[string]interface{}{
                              "amount":2616.6667,
                              "buom_rate":62800,
                              "bypass_status":"",
                              "convert_to":24,
                              "created_at":"2021-01-22T16:53:14.549+07:00",
                              "credit_limit_id":96746,
                              "date":"2021-01-22T16:53:14.549+07:00",
                              "discount_amount":0,
                              "discount_id":nil,
                              "gross_amount":389883.3383,
                              "id":3492,
                              "item_type":"0",
                              "net_amount":389883.3383,
                              "no_invoice":"",
                              "old_buom_rate":0,
                              "old_gross_amount":0,
                              "old_price_level":"",
                              "order_id":2494,
                              "order_item_id":nil,
                              "price_type":"R",
                              "product_id":4288,
                              "qty":6,
                              "qty_b":5,
                              "selected":true,
                              "status":"unpaid",
                              "uom":"CAR",
                              "uom_b":"PAC",
                              "updated_at":"2021-01-22T17:47:06.514+07:00",
                              "voucher_amount":0,
                              "voucher_id":nil,
                           },
                           "mars_value":"65000.0000",
                           "product": map[string]interface{}{
                              "active_status":"0",
                              "barcode":nil,
                              "created_at":"2018-11-27T20:01:05.040+07:00",
                              "description":"MI INSTAN SAR KARI AYAM DUA",
                              "id":4288,
                              "item_sn":"RKRD",
                              "item_type":"NPL",
                              "principal_id":"ISM",
                              "product_code":"124393",
                              "seq_id":13681,
                              "status":nil,
                              "sub_brand":"Sarimi Isi Dua",
                              "title":"MI INSTAN SAR KARI AYAM DUA",
                              "updated_at":"2019-12-11T10:22:44.295+07:00",
                           },
                           "sales": map[string]interface{}{},
                           "value":"ITEMCD=124393",
                          },
                          map[string]interface{}{
                           "dsc":"2111",
                           "error_message":"Terdapat perbedaan harga",
                           "item": map[string]interface{}{
                              "amount":2616.6667,
                              "buom_rate":62800,
                              "bypass_status":"",
                              "convert_to":24,
                              "created_at":"2021-01-22T16:53:14.549+07:00",
                              "credit_limit_id":96746,
                              "date":"2021-01-22T16:53:14.549+07:00",
                              "discount_amount":0,
                              "discount_id":nil,
                              "gross_amount":389883.3383,
                              "id":3492,
                              "item_type":"0",
                              "net_amount":389883.3383,
                              "no_invoice":"",
                              "old_buom_rate":0,
                              "old_gross_amount":0,
                              "old_price_level":"",
                              "order_id":2494,
                              "order_item_id":nil,
                              "price_type":"R",
                              "product_id":4288,
                              "qty":6,
                              "qty_b":5,
                              "selected":true,
                              "status":"unpaid",
                              "uom":"CAR",
                              "uom_b":"PAC",
                              "updated_at":"2021-01-22T17:47:06.514+07:00",
                              "voucher_amount":0,
                              "voucher_id":nil,
                           },
                           "mars_value":"65000.0000",
                           "product": map[string]interface{}{
                              "active_status":"0",
                              "barcode":nil,
                              "created_at":"2018-11-27T20:01:05.040+07:00",
                              "description":"MI INSTAN SAR KARI AYAM DUA",
                              "id":4288,
                              "item_sn":"RKRD",
                              "item_type":"NPL",
                              "principal_id":"ISM",
                              "product_code":"124393",
                              "seq_id":13681,
                              "status":nil,
                              "sub_brand":"Sarimi Isi Dua",
                              "title":"MI INSTAN SAR KARI AYAM DUA",
                              "updated_at":"2019-12-11T10:22:44.295+07:00",
                           },
                           "sales": map[string]interface{}{},
                           "value":"ITEMCD=124393",
                          },
                          },
          
       "execution_datetime":"2021-01-22 17:47:18",
       "logid":1,
       "status_code":"109",
    },
    "response_code":200,
 },
 "message":"Failed",
 "order_id":2494,
 "order_no":"235981081150229",
}

  fmt.Println("------------- BEGIN OF mars_response")
  mars_response, _ := item.Attrs["mars_response"].(map[string]interface{})
  fmt.Println(mars_response)
  fmt.Println("------------- END OF mars_response")

  fmt.Println("------------- BEGIN OF response_body")
  response_body, _ := mars_response["response_body"].(map[string]interface{})
  fmt.Println(response_body)
  fmt.Println("------------- END OF response_body")

  fmt.Println("------------- BEGIN OF error_detail")
  error_detail, _ := response_body["error_detail"].([]interface{})
  fmt.Println(error_detail)
  fmt.Println("------------- END OF error_detail")


  Map := make(map[string]interface{})
  Map["error_details"] = error_detail

  fmt.Println("Map")
  fmt.Println(Map["error_details"].([]interface{})[1])
  fmt.Println("Map")

  for k, ed := range Map["error_details"].([]interface{}) {
    fmt.Println("------------- BEGIN OF ed -------------")
    // mars_value, _ := ed["mars_value"].(map[string]interface{})
    fmt.Println(k)
    mars_value := ed.(map[string]interface{})["mars_value"]
    fmt.Println(ed.(map[string]interface{})["mars_value"])
    fmt.Println("------------- END OF ed -------------")

    if mars_value != "" {
      fmt.Println("FOUND mars_value !!")
    }
    // for k, v := range ed.(map[string]interface{}) {
    //   fmt.Println("------------- BEGIN OF mars_value")
    //   fmt.Println("== key ==")
    //   fmt.Println(k)
    //   fmt.Println("== value ==")
    //   fmt.Println(v)
    //   fmt.Println("------------- END OF mars_value")
    // }
  }

}
