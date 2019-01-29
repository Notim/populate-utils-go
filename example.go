package main

import (
    "./populate"
    "./model"
    "encoding/json"
    "fmt"
    "math"
    "math/rand"
)

func main()  {

    var listPersons []model.Person

    for index := 0; index < 1e4; index++ {
        var person = model.Person{
            ID:            index,
            Age:           rand.Intn(60),
            Name:          populate.GenerateName(2),
            Money :        math.Round(rand.Float64() * float64(rand.Intn(150000))),
            Birth :        "null",
            CardNumber:    rand.Intn(9999),
            MaritalStatus: populate.GenerateMaritalStatus(),
            Country:       populate.GenerateCountry(),
            Gender:        "null",
        }
        listPersons = append(listPersons, person)
    }
    j, _ := json.Marshal(listPersons)

    fmt.Printf("%s", j)
}