package hago

import (
	"log"
	"net/http"
	"path"
)

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		parent: group,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

var anyMethods = []string{
	http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch,
	http.MethodHead, http.MethodOptions, http.MethodDelete, http.MethodConnect,
	http.MethodTrace,
}

func (group *RouterGroup) ANY(pattern string, handler HandlerFunc) {
	for _, method := range anyMethods {
		group.addRoute(method, pattern, handler)
	}
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodGet, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPost, pattern, handler)
}

func (group *RouterGroup) PUT(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPut, pattern, handler)
}

func (group *RouterGroup) DELETE(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodDelete, pattern, handler)
}
func (group *RouterGroup) OPTIONS(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodOptions, pattern, handler)
}

func (group *RouterGroup) PATCH(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodPatch, pattern, handler)
}

func (group *RouterGroup) HEAD(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodHead, pattern, handler)
}

func (group *RouterGroup) TRACE(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodTrace, pattern, handler)
}

func (group *RouterGroup) CONNECT(pattern string, handler HandlerFunc) {
	group.addRoute(http.MethodConnect, pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) createStaticHandler(relativePath string, fs http.FileSystem) HandlerFunc {
	absolutePath := path.Join(group.prefix, relativePath)
	fileServer := http.StripPrefix(absolutePath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(http.StatusNotFound)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouterGroup) Static(relativePath string, root string) {
	handler := group.createStaticHandler(relativePath, http.Dir(root))
	urlPattern := path.Join(relativePath, "/*filepath")
	group.GET(urlPattern, handler)
}
