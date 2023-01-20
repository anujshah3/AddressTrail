package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))

func GetSession(req *http.Request, sessionName string) (*sessions.Session, error) {
	session, err := store.Get(req, sessionName)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func SetAuthenticated(session *sessions.Session, userInfo map[string]interface{}) {
	session.Values["authenticated"] = true
	session.Values["userData"] = userInfo
	session.Options.MaxAge = 5 * 60
}

func IsAuthenticated(session *sessions.Session) bool {
	if auth, ok := session.Values["authenticated"].(bool); ok && auth {
		return true
	}

	return false
}

func GetUserInfo(session *sessions.Session) map[string]interface{} {
	if userInfo, ok := session.Values["userData"].(map[string]interface{}); ok {
		return userInfo
	}

	return nil
}

func ClearSession(session *sessions.Session) {
	session.Options.MaxAge = -1
	session.Save(nil, nil)
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		session, _ := GetSession(req, "user-session")

		if !IsAuthenticated(session) {
			http.Redirect(res, req, "/login", http.StatusSeeOther)
			return
		}

		next(res, req)
	}
}
