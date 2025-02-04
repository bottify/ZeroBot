package repeat

import (
	zero "github.com/wdvxdr1123/ZeroBot"
)

func init() { // 插件主体
	engine := zero.New()
	engine.OnCommand("开启复读", zero.OnlyToMe).SetBlock(true).SetPriority(10).
		Handle(func(ctx *zero.Ctx) {
			stop := zero.NewFutureEvent("message", 8, true,
				zero.CommandRule("关闭复读"), // 关闭复读指令
				ctx.CheckSession()).      // 只有开启者可以关闭复读模式
				Next()                    // 关闭需要一次

			echo, cancel := ctx.FutureEvent("message",
				ctx.CheckSession()). // 只复读开启复读模式的人的消息
				Repeat()             // 不断监听复读
			ctx.Send("已开启复读模式!")
			for {
				select {
				case e := <-echo: // 接收到需要复读的消息
					ctx.Send(e.RawMessage)
				case <-stop: // 收到关闭复读指令
					cancel() // 取消复读监听
					return   // 返回
				}
			}
		})
}
