package config

type Rds struct {
	Port     string `mapstructure:"port" validate:"required"`
	User     string `mapstructure:"user" validate:"required"`
	Password string `mapstructure:"password" validate:"required"`
	Host     string `mapstructure:"host" validate:"required"`
	DB       string `mapstructure:"db" validate:"required"`
	Protocol string `mapstructure:"protocol" validate:"required"`
}

type Mq struct {
	Host string `mapstructure:"host"  validate:"required"`
	User string `mapstructure:"user"  validate:"required"`
	Pass string `mapstructure:"pass"  validate:"required"`
}

type App struct {
	Port string `mapstructure:"port" validate:"required"`
}

type Email struct {
	Sender   string `mapstructure:"sender" validate:"required"`
	TextBody string `mapstructure:"text_body" validate:"required"`
}

type Otp struct {
	Max    int   `mapstructure:"max" validate:"required"`
	Min    int   `mapstructure:"min" validate:"required"`
	Expiry int64 `mapstructure:"expiry" validate:"required"` //in seconds
}

type Session struct {
	Expiry int64 `mapstructure:"expiry" validate:"required"`
}

type Aws struct {
	Region  string `json:"region" validate:"required"`
	Profile string `json:"string" validate:"required"`
}

type Config struct {
	App     App     `mapstructure:"app" validate:"required"`
	Rds     Rds     `mapstructure:"rds" validate:"required"`
	Mq      Mq      `mapstructure:"mq" validate:"required"`
	Otp     Otp     `mapstructure:"otp" validate:"required"`
	Email   Email   `mapstructure:"email" validate:"required"`
	Aws     Aws     `mapstructure:"aws" validate:"required"`
	Session Session `mapstructure:"session" validate:"required"`
}
