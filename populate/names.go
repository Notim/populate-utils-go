package populate

import (
    "bufio"
    "math/rand"
    "os"
)
var (
    scanner *bufio.Scanner
    arrayNames    []string
    arraySurNames []string
    name    map[int]string
    surname map[int]string
    err              error
)

func init() {
    loadNames()    // read name list archive
    loadSurNames() // read sur name list archive
}

// unexported function
func loadNames(){
    var lines []string

    names, err := os.Open("./data/names")
    if err != nil {
        panic(err)
    }
    defer names.Close() // to close after run function

    scanner = bufio.NewScanner(names)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    arrayNames = make([]string,0)
    for _, name := range (lines){
        arrayNames = append(arrayNames, name)
    }

    name = map[int]string{}

    for in, na := range(arrayNames) {
        name[in] = na
    }
}

// unexported function
func loadSurNames() {
    names, err := os.Open("./data/surnames")
    if err != nil {
        panic(err)
    }

    defer names.Close()

    var lines []string
    scanner := bufio.NewScanner(names)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    arraySurNames = make([]string,0)
    for _, name := range(lines){

        arraySurNames = append(arraySurNames, name)
    }

    surname = map[int]string{}

    for in, na := range(arraySurNames) {
        surname[in] = na
    }
}

// size of name with one name and (size*surnames)
func GenerateName(size int) string{
    var str = ""
    str += name[rand.Intn(len(name))]
    for i:=0; i < size; i++ {
        str += " " + surname[rand.Intn(len(surname))]
    }

    return str
}