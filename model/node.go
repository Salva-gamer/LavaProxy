package model

type Node struct {
	Host     string `yaml:"host" json:"host"`
	Port     string `yaml:"port" json:"port"`
	Password string `yaml:"password" json:"password"`
	Secure   bool   `yaml:"secure" json:"secure"`
	Alive    bool   `yaml:"-" json:"-"`
}

type Config struct {
	Puerto     string `yaml:"puerto"`
	CheckDelay int    `yaml:"check_delay"` // en segundos
	Nodos      []Node `yaml:"nodos"`
}