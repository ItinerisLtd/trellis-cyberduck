package lib

import (
	core "trellis-cli/trellis"
)

type Trellis struct{}

func (t *Trellis) DetectPath() (trellisProject *core.Trellis, path string, err error) {
	// Ensure we are inside trellis root and activate virtualenv.
	project := &core.Project{}
	trellisProject = core.NewTrellis(project)
	if err := trellisProject.LoadProject(); err != nil {
		return nil, "", err
	}
	return trellisProject, trellisProject.Path, nil
}

func (t *Trellis) DetectPathAndEnv(maybeEnv string) (trellisProject *core.Trellis, path string, env string, err error) {
	trellisProject, path, err = t.DetectPath()
	if err != nil {
		return nil, "", "", err
	}

	// Validate environment exist.
	if err := trellisProject.ValidateEnvironment(maybeEnv); err != nil {
		return nil, "", "", err
	}
	return trellisProject, path, maybeEnv, nil
}

func (t *Trellis) DetectPathAndEnvAndSite(maybeEnv string, maybeSite string) (trellisProject *core.Trellis,path string, env string, site string, err error) {
	// Ensure we are inside trellis root and activate virtualenv.
	trellisProject, path, env, err = t.DetectPathAndEnv(maybeEnv)
	if err != nil {
		return nil, "", "", "", err
	}

	// Validate or detect site.
	if site, err = trellisProject.FindSiteNameFromEnvironment(env, maybeSite); err != nil {
		return nil, "", "", "", err
	}
	return
}
