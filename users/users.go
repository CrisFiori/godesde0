package users

import (
	"fmt"
	"home/cfiori/go/src/github.com/CrisFiori/godesde0/modelos"
	"time"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Cristian", time.Now(), true)

	fmt.Println(u)
}
