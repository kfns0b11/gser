package gin

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
)

type IResponse interface {
	IJson(obj interface{}) IResponse

	IJsonp(obj interface{}) IResponse

	IXml(obj interface{}) IResponse

	IHtml(template string, obj interface{}) IResponse

	IText(format string, values ...interface{}) IResponse

	IRedirect(path string) IResponse

	ISetHeader(key string, val string) IResponse

	ISetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	ISetStatus(code int) IResponse

	ISetOkStatus() IResponse
}

func (ctx *Context) IJsonp(obj interface{}) IResponse {
	callbackFunc := ctx.Query("callback")
	ctx.ISetHeader("Content-Type", "application/javascript")
	callback := template.JSEscapeString(callbackFunc)

	_, err := ctx.Writer.Write([]byte(callback))
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write([]byte("("))
	if err != nil {
		return ctx
	}
	ret, err := json.Marshal(obj)
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write(ret)
	if err != nil {
		return ctx
	}
	_, err = ctx.Writer.Write([]byte(")"))
	if err != nil {
		return ctx
	}
	return ctx
}

func (ctx *Context) IXml(obj interface{}) IResponse {
	byt, err := xml.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/html")
	ctx.Writer.Write(byt) //nolint: errcheck
	return ctx
}

// html输出
func (ctx *Context) IHtml(file string, obj interface{}) IResponse {
	t, err := template.New("output").ParseFiles(file)
	if err != nil {
		return ctx
	}
	if err := t.Execute(ctx.Writer, obj); err != nil {
		return ctx
	}

	ctx.ISetHeader("Content-Type", "application/html")
	return ctx
}

func (ctx *Context) IText(format string, values ...interface{}) IResponse {
	out := fmt.Sprintf(format, values...)
	ctx.ISetHeader("Content-Type", "application/text")
	ctx.Writer.Write([]byte(out)) //nolint: errcheck
	return ctx
}

func (ctx *Context) IRedirect(path string) IResponse {
	http.Redirect(ctx.Writer, ctx.Request, path, http.StatusMovedPermanently)
	return ctx
}

func (ctx *Context) ISetHeader(key string, val string) IResponse {
	ctx.Writer.Header().Add(key, val)
	return ctx
}

func (ctx *Context) ISetCookie(key string, val string, maxAge int, path string, domain string, secure bool, httpOnly bool) IResponse {
	if path == "" {
		path = "/"
	}
	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     key,
		Value:    url.QueryEscape(val),
		MaxAge:   maxAge,
		Path:     path,
		Domain:   domain,
		SameSite: 1,
		Secure:   secure,
		HttpOnly: httpOnly,
	})
	return ctx
}

func (ctx *Context) ISetStatus(code int) IResponse {
	ctx.Writer.WriteHeader(code)
	return ctx
}

func (ctx *Context) ISetOkStatus() IResponse {
	ctx.Writer.WriteHeader(http.StatusOK)
	return ctx
}

func (ctx *Context) IJson(obj interface{}) IResponse {
	byt, err := json.Marshal(obj)
	if err != nil {
		return ctx.ISetStatus(http.StatusInternalServerError)
	}
	ctx.ISetHeader("Content-Type", "application/json")
	ctx.Writer.Write(byt) //nolint: errcheck
	return ctx
}
