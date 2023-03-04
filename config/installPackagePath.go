package config

/*
	 	作用是优化我们安装包的体验 可以直接使用名字安装包 而不是全部的路径
		使用方式 gt install gin == gt i gin == go get -u github.com/gin-gonic/ginq
		所以我们每次有新的安装包的时候都需要在这里添加一个键值对
*/
var (
	InstallPackagePathDict = map[string]string{
		"qrcode": "github.com/yeqown/go-qrcode/v2", // 二维码生成器
		"gin":    "github.com/gin-gonic/gin",       // web框架
		"cli":    "github.com/urfave/cli/v2",       // cli框架
		"gorm":   "gorm.io/gorm",                   // orm框架
		"mysql":  "gorm.io/driver/mysql",           // mysql驱动
		"sqlite": "gorm.io/driver/sqlite",          // sqlite驱动
		"fiber":  "github.com/gofiber/fiber/v2",    // fiber框架

	}
)
