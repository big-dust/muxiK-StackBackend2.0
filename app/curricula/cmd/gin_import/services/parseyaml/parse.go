package parseyaml

import "github.com/spf13/viper"

func GetYaml() *viper.Viper {
	v := viper.New()
	v.AddConfigPath("D:\\muxiK-StackBackend2.0\\app\\curricula\\cmd\\gin_import\\config") // 路径
	v.SetConfigName("config")                                                             // 名称
	v.SetConfigType("yaml")                                                               // 类型
	err := v.ReadInConfig()                                                               // 读取配置
	if err != nil {
		panic(err)
		return nil
	}
	return v
}
