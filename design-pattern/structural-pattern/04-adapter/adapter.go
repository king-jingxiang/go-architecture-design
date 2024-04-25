package _4_adapter

// ObjectAdapterV1 适配器
type ObjectAdapterV1 struct {
	Adaptee *ObjectAdapteeV1
}

// Execute 适配器执行方法
func (p *ObjectAdapterV1) Execute() string {
	return p.Adaptee.SpecificExecute()
}

// ObjectAdapterV2 适配器
type ObjectAdapterV2 struct {
	Adaptee *ObjectAdapteeV2
}

// Execute 适配器执行方法
func (p *ObjectAdapterV2) Execute() string {
	return p.Adaptee.SpecificExecute()
}
