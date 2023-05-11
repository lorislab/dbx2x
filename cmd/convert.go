package cmd

import (
	"github.com/lorislab/dbx2x/tools"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type convertFlags struct {
	File   string `mapstructure:"file"`
	Output string `mapstructure:"output"`
}

func (f convertFlags) log() log.Fields {
	return log.Fields{
		"file":   f.File,
		"output": f.Output,
	}
}

func init() {
	rootCmd.AddCommand(convertCmd)
	addFlag(convertCmd, "output", "o", "", "xml output file")
	addFlag(convertCmd, "file", "f", "", "excel input file")

}

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert excel to db-unit xml",
	Long:  `Convert excel to db-unit xml`,
	Run: func(cmd *cobra.Command, args []string) {
		options := readConvertFlags()
		log.WithFields(options.log()).Info("Converting excel file to db-unit xml")
		tools.ConvertExcel2Xml(options.File, options.Output)
		log.WithFields(options.log()).Info("The output xml file has been created")
	},
}

func readConvertFlags() convertFlags {
	options := convertFlags{}
	err := viper.Unmarshal(&options)
	if err != nil {
		log.Fatal(err)
	}

	log.WithFields(options.log()).Debug("Load configuration")
	return options
}

func addFlag(command *cobra.Command, name, shorthand, value, usage string) *pflag.Flag {
	return addFlagExt(command, name, shorthand, value, usage, false)
}

func addFlagExt(command *cobra.Command, name, shorthand, value, usage string, required bool) *pflag.Flag {
	command.Flags().StringP(name, shorthand, value, usage)
	if required {
		err := command.MarkFlagRequired(name)
		if err != nil {
			log.Panic(err)
		}
	}
	return addViper(command, name)
}

func addViper(command *cobra.Command, name string) *pflag.Flag {
	f := command.Flags().Lookup(name)
	err := viper.BindPFlag(name, f)
	if err != nil {
		log.Panic(err)
	}
	return f
}
