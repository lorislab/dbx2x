package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	goVersion "go.hein.dev/go-version"
)

var (
	// Used for flags.
	shortened = false
	output    = "json"
	bv        BuildVersion
	cfgFile   string
	v         string
	rootCmd   = &cobra.Command{
		Use:   "dbx2x",
		Short: "Convert excel to db-unit xml",
		Long:  `Convert excel to db-unit xml`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := setUpLogs(os.Stdout, v); err != nil {
				return err
			}
			return nil
		},
	}
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version will output the current build information",
		Long:  ``,
		Run: func(_ *cobra.Command, _ []string) {
			resp := goVersion.FuncWithOutput(shortened, bv.Version, bv.Commit, bv.Date, output)
			fmt.Print(resp)
		},
	}
)

type BuildVersion struct {
	Version string
	Commit  string
	Date    string
}

// Execute executes the root command.
func Execute(version BuildVersion) {
	bv = version

	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
	})
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	versionCmd.Flags().BoolVarP(&shortened, "short", "s", false, "Print just the version number.")
	versionCmd.Flags().StringVarP(&output, "output", "o", "json", "Output format. One of 'yaml' or 'json'.")
	rootCmd.AddCommand(versionCmd)

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dbx2x.yaml)")
	rootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", log.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".dbx2x")
	}

	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.SetEnvPrefix("DBX2X")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.WithField("file", viper.ConfigFileUsed()).Debug("Using config")
	}
}

func setUpLogs(out io.Writer, level string) error {
	log.SetOutput(out)
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return err
	}
	log.SetLevel(lvl)
	return nil
}
