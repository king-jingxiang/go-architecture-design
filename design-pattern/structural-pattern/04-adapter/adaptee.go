package _4_adapter

// ObjectAdapteeV1 被适配对象
type ObjectAdapteeV1 struct {
}

// SpecificExecute 被适配对象的方法
func (b *ObjectAdapteeV1) SpecificExecute() string {
	return "Final Execute Result V1"
}

// ObjectAdapteeV2 被适配对象
type ObjectAdapteeV2 struct {
}

// SpecificExecute 被适配对象的方法
func (b *ObjectAdapteeV2) SpecificExecute() string {
	return "Final Execute Result V2"
}
