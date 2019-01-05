package authentication

const ControllerBaseName = "github.com/NEUOJ-NG/NEUOJ-NG-backend/controller."

func initAuthMap(m map[string]int) {
	m[ControllerBaseName + "RefreshToken"] = StandardUser
	m[ControllerBaseName + "Ping"] = Admin
}
