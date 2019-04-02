package cookie

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/qqqasdwx/blog/config"
)

const USER_COOKIE_NAME = "v2af-blog"

var SecureCookie *securecookie.SecureCookie

func Init() {
	var hashKey = []byte(config.Config().Http.Secret)
	var blockKey = []byte(nil)
	SecureCookie = securecookie.New(hashKey, blockKey)
}

func ReadUser(r *http.Request) (id int64, name string, found bool) {
	if cookie, err := r.Cookie(USER_COOKIE_NAME); err == nil {
		value := make(map[string]interface{})
		if err = SecureCookie.Decode(USER_COOKIE_NAME, cookie.Value, &value); err == nil {
			id = value["id"].(int64)
			name = value["name"].(string)
			if id == 0 || name == "" {
				return
			} else {
				found = true
				return
			}
		}
	}
	return
}

func WriteUser(w http.ResponseWriter, id int64, name string) error {
	value := make(map[string]interface{})
	value["id"] = id
	value["name"] = name
	encoded, err := SecureCookie.Encode(USER_COOKIE_NAME, value)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:   USER_COOKIE_NAME,
		Value:  encoded,
		Path:   "/",
		MaxAge: 3600 * 24 * 7,
	}
	http.SetCookie(w, cookie)
	return nil
}

func RemoveUser(w http.ResponseWriter) error {
	value := make(map[string]interface{})
	value["id"] = 0
	value["name"] = ""
	encoded, err := SecureCookie.Encode(USER_COOKIE_NAME, value)
	if err != nil {
		return err
	}
	cookie := &http.Cookie{
		Name:   USER_COOKIE_NAME,
		Value:  encoded,
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	return nil
}
