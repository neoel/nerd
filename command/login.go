package command

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/mitchellh/cli"
	"github.com/nerdalize/nerd/nerd/client"
	"github.com/nerdalize/nerd/nerd/conf"
	"github.com/pkg/errors"
)

//LoginOpts describes command options
type LoginOpts struct {
	User string `long:"user" default:"" default-mask:"" env:"NERD_USER" description:"nerd username"`
	Pass string `long:"pass" default:"" default-mask:"" env:"NERD_PASS" description:"nerd password"`
	NerdOpts
}

//Login command
type Login struct {
	*command

	opts   *LoginOpts
	parser *flags.Parser
}

//LoginFactory returns a factory method for the join command
func LoginFactory() func() (cmd cli.Command, err error) {
	cmd := &Login{
		command: &command{
			help:     "",
			synopsis: "Setup an authorized session.",
			parser:   flags.NewNamedParser("nerd login", flags.Default),
			ui: &cli.BasicUi{
				Reader: os.Stdin,
				Writer: os.Stderr,
			},
		},

		opts: &LoginOpts{},
	}

	cmd.runFunc = cmd.DoRun
	_, err := cmd.command.parser.AddGroup("options", "options", cmd.opts)
	if err != nil {
		panic(err)
	}

	return func() (cli.Command, error) {
		return cmd, nil
	}
}

//DoRun is called by run and allows an error to be returned
func (cmd *Login) DoRun(args []string) error {
	user := cmd.opts.User
	pass := cmd.opts.Pass
	if cmd.opts.User == "" || cmd.opts.Pass == "" {
		var err error
		user, pass, err = UserPassProvider(cmd.ui)()
		if err != nil {
			return errors.Wrap(err, "failed to retreive username and password")
		}
	}
	config, err := conf.Read()
	if err != nil {
		return errors.Wrap(err, "failed to read nerd config file")
	}
	cl := client.NewAuthAPI(config.Auth.APIEndpoint)
	token, err := cl.GetToken(user, pass)
	if err != nil {
		return errors.Wrap(err, "failed to get nerd token for username and password")
	}
	if token == "" {
		return errors.New("failed to get nerd token for username and password")
	}
	err = conf.WriteNerdToken(token)
	if err != nil {
		return errors.Wrap(err, "failed to write nerd token to disk")
	}
	logrus.Info("Successful login")
	return nil
}
