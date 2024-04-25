package _8_strategy

type Strategy interface {
	Execute() string
}

// StrategyA 策略
type StrategyA struct {
}

// Execute 策略方法
func (s *StrategyA) Execute() string {
	return "Strategy A"
}

// StrategyB 策略
type StrategyB struct {
}

// Execute 策略方法
func (s *StrategyB) Execute() string {
	return "Strategy B"
}

func NewStrategyA() Strategy {
	return &StrategyA{}
}

func NewStrategyB() Strategy {
	return &StrategyB{}
}
