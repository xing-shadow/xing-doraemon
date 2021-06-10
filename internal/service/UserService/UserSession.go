package UserService

import (
	"fmt"
	ginsession "github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"xing-doraemon/global"
	"xing-doraemon/internal/model/db"
)

const SessionName = "session_doraemon"

type userSession struct {
	option ginsession.Options
}

func initUserSession() {
	UserSession = userSession{
		ginsession.Options{
			Path:     "/",
			MaxAge:   opt.SessCfg.MaxAge,
			HttpOnly: true,
		},
	}
}

func (u *userSession) Save(c *gin.Context, user *db.User) error {
	sess := ginsession.Default(c)
	sess.Set("user", user)
	sess.Options(u.option)
	err := sess.Save()
	if err != nil {
		global.Log.Error("userSession save session err:" + err.Error())
		return err
	}
	return err
}

func (u *userSession) Read(c *gin.Context) *db.User {
	sess := ginsession.Default(c)
	userTemp := sess.(ginsession.Session).Get("user")
	user, ok := userTemp.(*db.User)
	if !ok {
		return nil
	}
	return user
}

func (u *userSession) Logout(c *gin.Context) error {
	sess := ginsession.Default(c)
	sess.Options(ginsession.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})
	sess.Delete("user")
	return sess.Save()
}

func NewSessionStore() ginsession.Store {
	storeTyp := opt.SessCfg.Type
	switch storeTyp {
	case "redis":
	case "files":
	default:
		fmt.Println("NewSessionStore", "cookie", opt.SessCfg.Secret)
		store := cookie.NewStore([]byte(opt.SessCfg.Secret))
		store.Options(ginsession.Options{
			Path:     "/",
			MaxAge:   opt.SessCfg.MaxAge,
			HttpOnly: true,
		})
		return store
	}
	return nil
}
