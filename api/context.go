package api

type Context struct {
	groups   []string
	response *Response
}

func NewContext() *Context {
	return &Context{
		groups:   []string{},
		response: &Response{},
	}
}

func (c *Context) Groups() []string {
	return c.groups
}

func (c *Context) AddGroups(groups ...string) {
	c.groups = append(c.groups, groups...)
}

func (c *Context) Response() *Response {
	return c.response
}
