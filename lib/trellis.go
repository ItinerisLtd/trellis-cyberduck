package lib

import core "trellis-cli/trellis"

type Trellis struct{}

func (t *Trellis) Detect(maybeEnv string, maybeSite string) (path string, env string, site string, err error) {
	// Ensure we are inside trellis root and activate virtualenv.
	project := &core.Project{}
	trellis := core.NewTrellis(project)
	if err := trellis.LoadProject(); err != nil {
		return "", "", "", err
	}
	path = trellis.Path

	// Validate environment exist.
	if err := trellis.ValidateEnvironment(maybeEnv); err != nil {
		return "", "", "", err
	}
	env = maybeEnv

	// Validate or detect site.
	if site, err = trellis.FindSiteNameFromEnvironment(env, maybeSite); err != nil {
		return "", "", "", err
	}

	return
}
