package plugin

import "context"

// Args provides plugin execution arguments.
type Args struct {
	Pipeline

	// Secrets represents the key/vals to interpolate
	Secrets map[string]string `envconfig:"PLUGIN_SECRETS" required:"true"`

	// Directory is the optinal target; defaults to .
	Directory string `envconfig:"PLUGIN_DIRECTORY"`

	// FileDenyList is an optional list of files to ignore
	FileDenyList []string `envconfig:"PLUGIN_FILE_DENY_LIST"`
}

// Exec executes the plugin
func Exec(ctx context.Context, args Args) error {
	files, err := getFiles(args.Directory)
	if err != nil {
		return err
	}

	if len(args.FileDenyList) > 0 {
		files = filter(files, isAllowed(args.FileDenyList))
	}

	for find, replace := range args.Secrets {
		if err := findAndReplace(files, find, replace); err != nil {
			return err
		}
	}
	return nil
}
