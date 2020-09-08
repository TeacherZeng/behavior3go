package composites

import (
	_ "fmt"

	b3 "github.com/magicsea/behavior3go"
	_ "github.com/magicsea/behavior3go/config"
	. "github.com/magicsea/behavior3go/core"
)

type Sequence struct {
	Composite
}

/**
 * Tick method.遍历子节点，子节点SUCCESS就执行下一个，最后返回SUCCESS，否则返回不成功的子节点状态。
类似and 语义
 * @method tick
 * @param {b3.Tick} tick A tick instance.
 * @return {Constant} A state constant.
**/
func (this *Sequence) OnTick(tick *Tick) b3.Status {
	//fmt.Println("tick Sequence :", this.GetTitle())
	for i := 0; i < this.GetChildCount(); i++ {
		var status = this.GetChild(i).Execute(tick)
		if status != b3.SUCCESS {
			return status
		}
	}
	return b3.SUCCESS
}
