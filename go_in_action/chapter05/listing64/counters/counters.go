package counters

type alertCounter int

//New 创建并返回一个未公开的 alertCounter 类型的值
func New(value int) alertCounter {
	return alertCounter(value)
}
