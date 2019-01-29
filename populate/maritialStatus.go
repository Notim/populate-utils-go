package populate

import "math/rand"

func GenerateMaritalStatus() string{
    var status = map[int]string{}
    status[0] = "Single"
    status[1] = "Married"
    status[2] = "Widow"
    status[3] = "divorced"

    return status[rand.Intn(4)]
}

