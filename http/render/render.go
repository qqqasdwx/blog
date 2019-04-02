package render

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/qqqasdwx/blog/config"
	"github.com/qqqasdwx/blog/http/cookie"
	"github.com/unrolled/render"
)

var Render *render.Render

func Init() {
	debug := config.Config().Debug
	Render = render.New(render.Options{
		Directory:  "views",
		Extensions: []string{".html"},
		Delims: render.Delims{
			Left:  "{{",
			Right: "}}",
		},
		Funcs:         []template.FuncMap{},
		IndentJSON:    false,
		IsDevelopment: debug,
	})
}

func Data(c *gin.Context, key string, val interface{}) {
	m, ok := c.Get("DATA_MAP")
	if ok {
		mm := m.(map[string]interface{})
		mm[key] = val
		c.Set("DATA_MAP", mm)
	} else {
		c.Set("DATA_MAP", map[string]interface{}{key: val})
	}
}

func HTML(c *gin.Context, name string, htmlOpt ...render.HTMLOptions) {
	userid, username, found := cookie.ReadUser(c.Request)
	Data(c, "ID", userid)
	Data(c, "NAME", username)
	Data(c, "LOGIN", found)
	Data(c, "TEST", "1")
	Render.HTML(c.Writer, http.StatusOK, name, c.MustGet("DATA_MAP"), htmlOpt...)
}

func JSON(c *gin.Context, v interface{}, statusCode ...int) {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	Render.JSON(c.Writer, code, v)

	// bs, _ := json.Marshal(v)
	// w.WriteHeader(code)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.Write(bs)
}

func Auto(c *gin.Context, err error, v ...interface{}) {
	if err != nil {
		JSON(c, map[string]interface{}{"msg": err.Error()})
		return
	}

	if len(v) > 0 {
		JSON(c, map[string]interface{}{"msg": "", "data": v[0]})
	} else {
		JSON(c, map[string]interface{}{"msg": ""})
	}
}

func Text(c *gin.Context, v string, codes ...int) {
	code := http.StatusOK
	if len(codes) > 0 {
		code = codes[0]
	}
	Render.Text(c.Writer, code, v)
}
