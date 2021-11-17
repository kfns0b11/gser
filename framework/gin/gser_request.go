package gin

import (
	"mime/multipart"

	"github.com/spf13/cast"
)

type IRequest interface {
	// parse parameters from uri query
	DefaultQueryInt(key string, defaultValue int) (int, bool)
	DefaultQueryInt64(key string, defaultValue int64) (int64, bool)
	DefaultQueryFloat32(key string, defaultValue float32) (float32, bool)
	DefaultQueryFloat64(key string, defaultValue float64) (float64, bool)
	DefaultQueryBool(key string, defaultValue bool) (bool, bool)
	DefaultQueryString(key string, defaultValue string) (string, bool)
	DefaultQueryStringSlice(key string, defaultValue []string) ([]string, bool)

	// parse parameters from url
	DefaultParamInt(key string, defaultValue int) (int, bool)
	DefaultParamInt64(key string, defaultValue int64) (int64, bool)
	DefaultParamFloat32(key string, defaultValue float32) (float32, bool)
	DefaultParamFloat64(key string, defaultValue float64) (float64, bool)
	DefaultParamBool(key string, defaultValue bool) (bool, bool)
	DefaultParamString(key string, defaultValue string) (string, bool)
	DefaultParam(key string) interface{}

	// parse parameters from form
	DefaultFormInt(key string, defaultValue int) (int, bool)
	DefaultFormInt64(key string, defaultValue int64) (int, bool)
	DefaultFormFloat64(key string, defaultValue float64) (float64, bool)
	DefaultFormFloat32(key string, defaultValue float32) (float32, bool)
	DefaultFormBool(key string, defaultValue bool) (bool, bool)
	DefaultFormString(key string, defaultValue string) (string, bool)
	DefaultFormStringSlice(key string, defaultValue []string) ([]string, bool)
	DefaultFormFile(key string) (*multipart.FileHeader, error)
	DefaultForm(key string) interface{}

	BindJson(obj interface{}) error

	BindXml(obj interface{}) error

	GetRawData() ([]byte, error)

	// basic information
	Uri() string
	Method() string
	Host() string
	ClientIp() string

	// header
	Headers() map[string]string
	Header(key string) (string, bool)

	// cookie
	Cookies() map[string]string
	Cookie(key string) (string, bool)
}

func (ctx *Context) QueryAll() map[string][]string {
	ctx.initQueryCache()
	return map[string][]string(ctx.queryCache)
}

func (ctx *Context) DefaultQueryInt(key string, defaultValue int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultQueryInt64(key string, defaultValue int64) (int64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultQueryFloat64(key string, defaultValue float64) (float64, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultQueryFloat32(key string, defaultValue float32) (float32, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultQueryBool(key string, defaultValue bool) (bool, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultQueryString(key string, defaultValue string) (string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultQueryStringSlice(key string, defaultValue []string) ([]string, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return defaultValue, false
}

func (ctx *Context) GetParam(key string) interface{} {
	if val, ok := ctx.Params.Get(key); ok {
		return val
	}
	return nil
}

func (ctx *Context) DefaultParamInt(key string, defaultValue int) (int, bool) {
	if val := ctx.GetParam(key); val != nil {
		// 通过cast进行类型转换
		return cast.ToInt(val), true
	}
	return defaultValue, false
}

func (ctx *Context) DefaultParamInt64(key string, defaultValue int64) (int64, bool) {
	if val := ctx.GetParam(key); val != nil {
		return cast.ToInt64(val), true
	}
	return defaultValue, false
}

func (ctx *Context) DefaultParamFloat64(key string, defaultValue float64) (float64, bool) {
	if val := ctx.GetParam(key); val != nil {
		return cast.ToFloat64(val), true
	}
	return defaultValue, false
}

func (ctx *Context) DefaultParamFloat32(key string, defaultValue float32) (float32, bool) {
	if val := ctx.GetParam(key); val != nil {
		return cast.ToFloat32(val), true
	}
	return defaultValue, false
}

func (ctx *Context) DefaultParamBool(key string, defaultValue bool) (bool, bool) {
	if val := ctx.GetParam(key); val != nil {
		return cast.ToBool(val), true
	}
	return defaultValue, false
}

func (ctx *Context) DefaultParamString(key string, defaultValue string) (string, bool) {
	if val := ctx.GetParam(key); val != nil {
		return cast.ToString(val), true
	}
	return defaultValue, false
}

func (ctx *Context) FormAll() map[string][]string {
	ctx.initFormCache()
	return map[string][]string(ctx.formCache)
}

func (ctx *Context) DefaultFormInt64(key string, defaultValue int64) (int64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToInt64(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultFormFloat64(key string, defaultValue float64) (float64, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat64(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultFormFloat32(key string, defaultValue float32) (float32, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToFloat32(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultFormBool(key string, defaultValue bool) (bool, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return cast.ToBool(vals[0]), true
		}
	}
	return defaultValue, false
}

func (ctx *Context) DefaultFormStringSlice(key string, defaultValue []string) ([]string, bool) {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		return vals, true
	}
	return defaultValue, false
}

func (ctx *Context) DefaultForm(key string) interface{} {
	params := ctx.FormAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			return vals[0]
		}
	}
	return nil
}
