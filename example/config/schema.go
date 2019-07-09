// Code generated by genfig (schema built from 'default.yml') on 2019-07-09T23:31:30+02:00; DO NOT EDIT.

package config

type Config struct {
	Randomizer ConfigRandomizer
	LongDesc ConfigLongDesc
	Wip bool
	Version string
	Db ConfigDb
	Secrets []string
	Project string
	Server ConfigServer
	Apis ConfigApis
}

type ConfigApis struct {
	Google ConfigApisGoogle
}

type ConfigApisGoogle struct {
	Uri string
}

type ConfigDb struct {
	User string
	Pass string
	Uri string
}

type ConfigLongDesc struct {
	En string
	De string
}

type ConfigRandomizer struct {
	Threshold float64
}

type ConfigServer struct {
	Port int64
	Host string
}

