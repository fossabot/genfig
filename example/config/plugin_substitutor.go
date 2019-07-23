// Code generated by genfig plugin 'substitutor' on 2019-07-23T22:25:20+02:00; DO NOT EDIT.

package config

import (
	"strings"
)

const (
	maxSubstitutionIteraions = 5
)

var (
	raw Config
)

// Substitute replaces all.
// The return value informs, whether all substitutions could be
// applied within {maxSubstitutionIteraions} or not
func (c *Config) Substitute() bool {
	c.ResetSubstitution()

	// backup the "raw" configuration
	raw = *c

	run := 0
	for {
		if run == maxSubstitutionIteraions {
			return false
		}
		if c.substitute() == 0 {
			return true
		}
	}
}

// ResetSubstitution resets the configuration to the state,
// before the substitution was applied
func (c *Config) ResetSubstitution() {
	c = &raw
}

// substitute tries to replace all substitutions in strings
func (c *Config) substitute() int {
	cnt := 0

	r := strings.NewReplacer(
		"${apis.google.uri}", c.Apis.Google.Uri,

		"${db.pass}", c.Db.Pass,

		"${db.uri}", c.Db.Uri,

		"${db.user}", c.Db.User,

		"${longdesc.de}", c.LongDesc.De,

		"${longdesc.en}", c.LongDesc.En,

		"${project}", c.Project,

		"${server.host}", c.Server.Host,

		"${version}", c.Version,
	)

	if strings.Contains(c.Apis.Google.Uri, "${") {
		cnt += 1
		c.Apis.Google.Uri = r.Replace(c.Apis.Google.Uri)
		if !strings.Contains(c.Apis.Google.Uri, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.Db.Pass, "${") {
		cnt += 1
		c.Db.Pass = r.Replace(c.Db.Pass)
		if !strings.Contains(c.Db.Pass, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.Db.Uri, "${") {
		cnt += 1
		c.Db.Uri = r.Replace(c.Db.Uri)
		if !strings.Contains(c.Db.Uri, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.Db.User, "${") {
		cnt += 1
		c.Db.User = r.Replace(c.Db.User)
		if !strings.Contains(c.Db.User, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.LongDesc.De, "${") {
		cnt += 1
		c.LongDesc.De = r.Replace(c.LongDesc.De)
		if !strings.Contains(c.LongDesc.De, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.LongDesc.En, "${") {
		cnt += 1
		c.LongDesc.En = r.Replace(c.LongDesc.En)
		if !strings.Contains(c.LongDesc.En, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.Project, "${") {
		cnt += 1
		c.Project = r.Replace(c.Project)
		if !strings.Contains(c.Project, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.Server.Host, "${") {
		cnt += 1
		c.Server.Host = r.Replace(c.Server.Host)
		if !strings.Contains(c.Server.Host, "${") {
			cnt -= 1
		}
	}

	if strings.Contains(c.Version, "${") {
		cnt += 1
		c.Version = r.Replace(c.Version)
		if !strings.Contains(c.Version, "${") {
			cnt -= 1
		}
	}

	return cnt
}
