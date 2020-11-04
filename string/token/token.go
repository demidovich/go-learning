import (
    "fmt"
    "math/rand"
    "strings"
)

func token(length int) string {

    chars := [36]string{
        "0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
        "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", 
        "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", 
        "u", "v", "w", "x", "y", "z"}

    var value string

    for i := 0; i < length; i++ {
        char := chars[rand.Intn(35)]
        if rand.Intn(2) == 1 {
            char = strings.ToUpper(char)
        }
        value = fmt.Sprintf("%s%s", value, char)
    }

    return value
}

func main() {
    fmt.Println(token(32))
}
