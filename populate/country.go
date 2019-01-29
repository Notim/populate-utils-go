package populate

import "math/rand"

func GenerateCountry() string{
    var Country = map[int]string{}
    Country[0] = "United Kingdom"
    Country[1] = "United States"
    Country[2] = "Brazil"
    Country[3] = "French"
    Country[4] = "Canada"
    Country[3] = "Germany"
    Country[5] = "French"
    Country[6] = "Italy"
    Country[7] = "Russia"
    Country[8] = "Portugal"
    Country[9] = "Australia"
    Country[10] = "India"
    Country[11] = "Russia"

    return Country[rand.Intn(11)]
}