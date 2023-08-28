package gwebz

import "go.uber.org/fx"

func RunContainerIOC(fc func(), provideConstructors ...interface{}) {
	// provide中应该包括我们需要初始化的东西 注意 provide中是懒加载，fx框架会自动帮我们处理依赖关系
	// Invoke 中应该放入我们的server监听对象，
	// 这样就通过把依赖关系和启动方式包起来 使用依赖注入的方式进行初始化启动
	fx.New(fx.Provide(provideConstructors), fx.Invoke(fc)).Run()
}
