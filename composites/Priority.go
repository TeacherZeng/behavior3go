package composites

import (
	b3 "github.com/magicsea/behavior3go"
	. "github.com/magicsea/behavior3go/core"
)

type Priority struct {
	Composite
}

/**
 * Tick method.遍历子节点，子节点FAILURE就执行下一个，最后返回FAILURE，否则返回非失败的子节点状态
类似 or 语义，只会选择一个成功节点
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Priority) OnTick(tick *Tick) b3.Status {
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.FAILURE {
			return status
		}
	}
	return b3.FAILURE
}
