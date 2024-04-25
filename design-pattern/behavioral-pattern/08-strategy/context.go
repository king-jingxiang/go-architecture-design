package _8_strategy

// Context 上下文
type Context struct {
	strategy Strategy
}

// SetStrategy 配置策略
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

// Execute 上下文类方法
func (c *Context) Execute() string {
	return c.strategy.Execute()
}

// NewContext 创建一个上下文
func NewContext() *Context {
	return &Context{}
}
