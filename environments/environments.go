package environments

import (
	"strings"
	"time"

	"github.com/spf13/viper"
)

func EnvironmentInit() {
	// ไฟล์ที่จะจัดเก็บตัว Connection string Database
	viper.SetConfigName("config")
	// ภาษาที่จะใช้ในการ Config
	viper.SetConfigType("yaml")
	// แล้วเข้ามาที่ environment folder
	viper.AddConfigPath("./environments")
	// ที่อยู่ของ file config เริ่มค้นหาจาก root ด้านนอกสุด
	viper.AddConfigPath(".")

	// รับ env จากภายนอกได้ เช่นจาก docker file หรือ docker compose file
	viper.AutomaticEnv()
	// ทำการเปลี่ยน (.) ใน config.yaml เป็น (_) เพื่ออ้างการแก้ไข env จากภายนอก
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// เรียก file config.yaml มาใช้
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func TimeZoneInit() {
	ict := time.Now().Local().Location()
	time.Local = ict
}
