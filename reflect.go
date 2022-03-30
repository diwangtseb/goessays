package main

import (
        "fmt"
        "reflect"
)

type User struct {
        UserName string `json:"username"`
        Grade    int
        Sex      bool
}

func main() {
        u := User{
                UserName: "tangsan",
                Grade:    1,
                Sex:      true,
        }
        t := reflect.TypeOf(u)
        v := reflect.ValueOf(&u)
        fmt.Println(t, v)
        v = v.Elem()
        f := v.FieldByName("UserName")
        fmt.Println()
        if f.Kind() == reflect.String {
                f.SetString("caonima")
        }
        fmt.Println(t, v)
        for i := 0; i < t.NumField(); i++ {
                f := t.Field(i)
                fmt.Println(f.Name, f.Type, f.Tag.Get("json"))
        }
}
