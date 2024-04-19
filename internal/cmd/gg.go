package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var Gg = cGg{}

type cGg struct {
	g.Meta `name:"gg" ad:"gf 生成器"`
}
type cGgInput struct {
	g.Meta `name:"gg"`
}

type cGgOutput struct {
}

func (receiver *cGg) Index(ctx context.Context, in cGgInput) (out *cGgOutput, err error) {
	gcmd.CommandFromCtx(ctx).Print()
	return
}
