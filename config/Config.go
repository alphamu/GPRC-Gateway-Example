package config

var configInstance *config

type config struct {
	MySql *DbDetails
}

func GetConfig() *config {
	if configInstance == nil {
		configInstance = &config{
			MySql: &DbDetails{
				Username: "nsa",
				Password: "letmein",
				DbName:   "funsa",
			},
		}
	}
	return configInstance
}
