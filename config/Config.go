package config

var configInstance *config

type config struct {
	MySql *DbDetails
}

func GetConfig() *config {
	if configInstance == nil {
		/* this connects to local host,
		you can configure it to connect to
		remove instances
		 */
		configInstance = &config{
			MySql: &DbDetails{
				Username: "nsa",
				Password: "letmein",
				DbName:   "funsa", /* schema */
			},
		}
	}
	return configInstance
}
