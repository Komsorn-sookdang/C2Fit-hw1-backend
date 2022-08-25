package session

import (
	"c2fit-hw-backend/models"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

const (
	OneDay = 86400
)

var store sessions.Store

func CreateStore() {
	newStore := cookie.NewStore([]byte("secret"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	newStore.Options(sessions.Options{MaxAge: OneDay})
	store = newStore
}

func GetStore() sessions.Store {
	return store
}

func LoginSession(ctx *gin.Context, user *models.User) error {
	session := sessions.Default(ctx)
	session.Clear()
	session.Options(sessions.Options{
		MaxAge: OneDay,
	})
	session.Set("online", true)
	session.Set("userId", user.ID.Hex())
	return session.Save()
}

func HasLoggedIn(ctx *gin.Context) bool {
	session := sessions.Default(ctx)
	online := session.Get("online")
	return online != nil
}

func GetUserId(ctx *gin.Context) string {
	session := sessions.Default(ctx)
	id := session.Get("userId")
	return fmt.Sprintf("%v", id)
}

func LogoutSession(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	return session.Save()
}
