package config
import (
    "fmt"
    "github.com/spf13/viper"
)
type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
}
type ServerConfig struct {
    Port int
    Mode string
}
type DatabaseConfig struct {
    Driver   string
    Host     string
    Port     int
    Username string
    Password string
    DBName   string
    Charset  string
}
func (c *DatabaseConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
        c.Username,
        c.Password,
        c.Host,
        c.Port,
        c.DBName,
        c.Charset,
    )
}

var GlobalConfig Config
func Init() error {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    
    // 只设置相对于 go.mod 的配置路径
    viper.AddConfigPath("config")
    
    if err := viper.ReadInConfig(); err != nil {
        return fmt.Errorf("failed to read config file: %v\nSearched paths: %v", err, viper.ConfigFileUsed())
    }

    if err := viper.Unmarshal(&GlobalConfig); err != nil {
        return fmt.Errorf("failed to unmarshal config: %v", err)
    }

    // 打印服务器配置
    fmt.Printf("\n=== Server Configuration ===\n")
    fmt.Printf("Port: %d\n", GlobalConfig.Server.Port)
    fmt.Printf("Mode: %s\n", GlobalConfig.Server.Mode)

    // 打印数据库配置
    fmt.Printf("\n=== Database Configuration ===\n")
    fmt.Printf("Driver: %s\n", GlobalConfig.Database.Driver)
    fmt.Printf("Host: %s\n", GlobalConfig.Database.Host)
    fmt.Printf("Port: %d\n", GlobalConfig.Database.Port)
    fmt.Printf("Username: %s\n", GlobalConfig.Database.Username)
    fmt.Printf("Database: %s\n", GlobalConfig.Database.DBName)
    fmt.Printf("Charset: %s\n", GlobalConfig.Database.Charset)
    fmt.Printf("DSN: %s\n", GlobalConfig.Database.DSN())


    fmt.Printf("\n=== Configuration End ===\n\n")
    return nil
} 