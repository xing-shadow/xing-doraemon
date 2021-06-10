package configs

var Cfg AutoGenerated

func InitConfig(path string) error {
	config, err := newSetting(path)
	if err != nil {
		return err
	}
	if err := config.ReadStruct(&Cfg); err != nil {
		return err
	}
	return nil
}

type AutoGenerated struct {
	App        App        `yaml:"App"`
	Logger     Logger     `yaml:"Logger"`
	Session    Session    `yaml:"Session"`
	Mysql      Mysql      `yaml:"Mysql"`
	Send       Send       `yaml:"Send"`
	RuleEngine RuleEngine `yaml:"RuleEngine"`
}

type App struct {
	Httpport int    `json:"Httpport" yaml:"Httpport"`
	Runmode  string `json:"Runmode" yaml:"Runmode"`
}

type Session struct {
	MaxAge int    `yaml:"MaxAge"`
	Secret string `yaml:"Secret"`
	Type   string `yaml:"Type"`
}

type Logger struct {
	LogDir      string `yaml:"LogDir"`
	FiLeName    string `yaml:"FiLeName"`
	Level       string `yaml:"Level"`
	MaxSize     int    `yaml:"MaxSize"`
	MaxBackups  int    `yaml:"MaxBackups"`
	MaxAge      int    `yaml:"MaxAge"`
	Development bool   `yaml:"Development"`
}

type Mysql struct {
	DBType    string `json:"DBType" yaml:"DBType"`
	DBUser    string `json:"DBUser" yaml:"DBUser"`
	DBPasswd  string `json:"DBPasswd" yaml:"DBPasswd"`
	DBLoc     string `json:"DBLoc" yaml:"DBLoc"`
	DBConnTTL int    `json:"DBConnTTL" yaml:"DBConnTTL"`
	DBName    string `json:"DBName" yaml:"DBName"`
	DBTns     string `json:"DBTns" yaml:"DBTns"`
}

type Send struct {
	WebHook string `yaml:"WebHook"`
}

type RuleEngine struct {
	NotifyRetries      int
	EvaluationInterval string
	ReloadInterval     string
}
