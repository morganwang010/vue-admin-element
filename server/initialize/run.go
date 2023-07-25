package initialize

func Run() {
	LoadConfig()
	Pgsql()
	Redis()
	Alipay()
	Router()
}
