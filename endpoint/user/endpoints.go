package user

import (
	"encoding/json"
	"github.com/satori/go.uuid"
	"github.com/valyala/fasthttp"
	"github.com/gospodinzerkalo/todo_app_golang/redis"
	"fmt"
	"time"
)

// init endpointFactory for user
func NewEndpointsFactory(userinfo UserInfo) *endpointsFactory {
	return &endpointsFactory{userInfo: userinfo}
}


type endpointsFactory struct {
	userInfo UserInfo
}


//GetUser
func (ef *endpointsFactory) SignIn() func (ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		user := &User{}
		body := ctx.PostBody()
		if err := json.Unmarshal(body,user);err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: invalid input"))
			return
		}
		res,err := ef.userInfo.GetUser(user.Email)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusNotFound,[]byte("Not found s"))
			return
		}
		// check password
		if res.Password != user.Password {
			writeResponse(ctx,fasthttp.StatusUnauthorized,[]byte("incorrect email or password"))
			return
		}
		//create session token for user
		sessionToken := uuid.NewV4()

		_,err = redis.Cache.Do("SETEX",sessionToken.String(),"60",user.Email)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error!"))
			return
		}
		cookie := &fasthttp.Cookie{}
		cookie.SetKey("session_token")
		cookie.SetValue(sessionToken.String())
		cookie.SetExpire(time.Now().Add(60*time.Second))
		ctx.Response.Header.SetCookie(cookie)


	}
}
//CreateUser...
func (ef *endpointsFactory) CreateUser() func (ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		data := ctx.PostBody()
		user := &User{}
		if err := json.Unmarshal(data,user);err != nil {
			writeResponse(ctx,fasthttp.StatusBadRequest,[]byte("Error: invalid input"))
			return
		}
		res, err := ef.userInfo.CreateUser(user)
		if err !=nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: try again"))
			return
		}
		response,err := json.Marshal(res)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error!"))
			return
		}
		writeResponse(ctx,fasthttp.StatusCreated,response)
	}
}

//GetListTask ...
func (ef *endpointsFactory) GetListUsers() func(ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		list,err := ef.userInfo.ListUsers()
		if err!=nil{
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		data, err := json.Marshal(list)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte("Error: "+err.Error()))
			return
		}
		writeResponse(ctx,fasthttp.StatusOK,data)

	}
}

func (ef *endpointsFactory) Welcome() func (ctx *fasthttp.RequestCtx){
	return func(ctx *fasthttp.RequestCtx) {
		c  := string(ctx.Request.Header.Cookie("session_token"))
		if c == ""{
			writeResponse(ctx,fasthttp.StatusUnauthorized,[]byte("Error: Please signin again"))
			return
		}
		session_token := c
		response,err := redis.Cache.Do("GET",session_token)
		if err != nil {
			writeResponse(ctx,fasthttp.StatusInternalServerError,[]byte(err.Error()))
			return
		}
		if response==nil {
			writeResponse(ctx,fasthttp.StatusUnauthorized,[]byte(""))
			return
		}

		writeResponse(ctx,fasthttp.StatusOK,[]byte(fmt.Sprintf("Welcome %v",response)))

	}
}

func writeResponse(ctx *fasthttp.RequestCtx,status int,msg []byte) {
	ctx.SetStatusCode(status)
	ctx.Write(msg)
}
