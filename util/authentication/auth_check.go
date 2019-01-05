package authentication

import (
	"github.com/NEUOJ-NG/NEUOJ-NG-backend/model"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	authMap = make(map[string]int)
	once    sync.Once
)

func AuthCheck(controllerName string, user *model.User) bool {
	log.Debugf("controller %v requested by user %v with privilege %v",
		controllerName, user.Username, user.Privilege)

	once.Do(func() {
		// init auth map only once
		initAuthMap(authMap)
		log.Debug("auth map initialized")
		log.Debug(authMap)
	})

	if v, ok := authMap[controllerName]; ok {
		return user.Privilege >= v
	}

	// default: auth failed
	return false
}
