package users

import (
	"fmt"
	"time"

	"github.com/Sequera86/godesde0/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Antonio", time.Now(), true)
	fmt.Println(u)
}
