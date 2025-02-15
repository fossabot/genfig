// Code generated by genfig (config built by merging 'default.yml' and 'development.local.toml') on 2019-07-24T22:16:50+02:00; DO NOT EDIT.

package config

func init() {
	Envs.DevelopmentLocal = Config{
		Version: "1-test",
		Project: "genfig",
		Db: ConfigDb{
			User: "chuck",
			Pass: "norris",
			Uri:  "mongdb://localhost:27017/${project}",
		},
		Randomizer: ConfigRandomizer{
			Threshold: 0.12345,
		},
		Server: ConfigServer{
			Port: 8080,
			Host: "localhost",
		},
		Secrets: []string{"ChuckNorriscanwinagameofConnectFourinonlythreemoves"},
		Apis: ConfigApis{
			Google: ConfigApisGoogle{
				Uri: "google.com",
			},
		},
		LongDesc: ConfigLongDesc{
			En: "Long description",
			De: "Lange Beschreibung",
		},
		Wip: true,
	}
	Envs.DevelopmentLocal._map = map[string]interface{}{"apis": map[interface{}]interface{}{"google": map[interface{}]interface{}{"uri": "google.com"}}, "db": map[interface{}]interface{}{"pass": "norris", "uri": "mongdb://localhost:27017/${project}", "user": "chuck"}, "longDesc": map[interface{}]interface{}{"de": "Lange Beschreibung", "en": "Long description"}, "project": "genfig", "randomizer": map[interface{}]interface{}{"threshold": 0.12345}, "secrets": []interface{}{"ChuckNorriscanwinagameofConnectFourinonlythreemoves"}, "server": map[interface{}]interface{}{"host": "localhost", "port": 8080}, "version": "1-test", "wip": true}
}
