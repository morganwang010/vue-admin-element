package main
import(
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	password, err := bcrypt.GenerateFromPassword([]byte("admin"), 10)
	if err != nil {
		fmt.Printf("dddddddd")
	}
	fmt.Println(string(password))
}

