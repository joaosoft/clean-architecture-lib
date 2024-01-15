package context

import (
	"context"
	"encoding/json"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/joaosoft/clean-infrastructure/domain"

	msg "github.com/joaosoft/clean-infrastructure/utils/pagination"

	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
	meta       interface{}
	pagination *msg.Pagination
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		Context: ctx,
	}
}

func (c *Context) Values(key any) any {
	return c.Value(key)
}

func (c *Context) Params() gin.Params {
	return c.Context.Params
}

func (c *Context) Keys() map[string]any {
	if c == nil || c.Context == nil {
		return nil
	}

	return c.Context.Keys
}

func (c *Context) Request() *http.Request {
	return c.Context.Request
}

func (c *Context) Response() gin.ResponseWriter {
	return c.Context.Writer
}

func (c *Context) SetIdMarket(idMarket int) {
	c.Context.Set(CtxIdMarket, idMarket)
}

func (c *Context) SetIdBu(idBu int) {
	c.Context.Set(CtxIdBu, idBu)
}

func (c *Context) SetIdShop(idShop int) {
	c.Context.Set(CtxIdShop, idShop)
}

func (c *Context) SetIdFascia(idFascia int) {
	c.Context.Set(CtxIdFascia, idFascia)
}

func (c *Context) SetIdUserExternal(idUserExternal int) {
	c.Context.Set(CtxIdUserExternal, idUserExternal)
}

func (c *Context) SetUsername(username string) {
	c.Context.Set(CtxUsername, username)
}

func (c *Context) SetLanguageCode(languageCode string) {
	c.Context.Set(CtxLanguageCode, languageCode)
}

func (c *Context) SetBody(body []byte) {
	c.Context.Set(CtxBody, body)
}

func (c *Context) SetAuthorizations(authorizations []string) {
	c.Context.Set(CtxAuthorizations, authorizations)
}

func (c *Context) GetIdMarket() int {
	return c.Context.GetInt(CtxIdMarket)
}

func (c *Context) GetIdBu() int {
	return c.Context.GetInt(CtxIdBu)
}

func (c *Context) GetIdShop() int {
	return c.Context.GetInt(CtxIdShop)
}

func (c *Context) GetIdFascia() int {
	return c.Context.GetInt(CtxIdFascia)
}

func (c *Context) GetIdUserExternal() int {
	return c.Context.GetInt(CtxIdUserExternal)
}

func (c *Context) GetUsername() string {
	return c.Context.GetString(CtxUsername)
}

func (c *Context) GetLanguageCode() string {
	return c.Context.GetString(CtxLanguageCode)
}

func (c *Context) GetBody() []byte {
	body, exists := c.Context.Get(CtxBody)
	if exists {
		if b, ok := body.([]byte); ok {
			return b
		}
	}
	return nil
}

func (c *Context) GetAuthorizations() []string {
	return c.Context.GetStringSlice(CtxAuthorizations)
}

func (c *Context) AddMeta(meta interface{}) domain.IContext {
	c.meta = meta
	return c
}

func (c *Context) AddPagination(pagination *msg.Pagination) domain.IContext {
	c.pagination = pagination
	return c
}

func (c *Context) GetMeta() any {
	return c.meta
}

func (c *Context) GetPagination() *msg.Pagination {
	return c.pagination
}

func (c *Context) FromGrpc(ctx context.Context) domain.IContext {
	gCtx := &gin.Context{}
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		md, _ = metadata.FromOutgoingContext(ctx)
	}

	if len(md.Get(ContextGrpcKeys)) > 0 {
		_ = json.Unmarshal([]byte(md.Get(ContextGrpcKeys)[0]), &gCtx.Keys)
	}

	return NewContext(gCtx)
}

func (c *Context) ToGrpc() context.Context {
	if c.Keys() == nil {
		return c
	}
	res := make(map[string]string)
	b, _ := json.Marshal(c.Keys())
	res[ContextGrpcKeys] = string(b)
	return metadata.NewOutgoingContext(context.Background(), metadata.New(res))
}
