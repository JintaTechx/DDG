package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const (
	defaultDatabaseURI string = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var (
	uri      string
	format   string
	schema   string
	database string
	path     string
	lang     string
)

var exit func(code int) = os.Exit

// setLang function to initialize the output language. case it has no set, the default selected
// will be the system language.
func setLang(lang string) {
	if lang == "" {
		translate.InitLanguage()
	} else {
		if !translate.SetLanguage(lang) {
			fmt.Printf(
				"Sorry, the language %s is not registered, try to use another: %v\n",
				lang,
				translate.GetKeys(),
			)
			exit(1)
		}
	}
}

// getFormat function to obtain the format based on received parameter
func getFormat(format string) (services.Printer, error) {
	oFormat := options.Options[strings.ToUpper(format)]
	if oFormat == nil {
		return nil, fmt.Errorf(
			"the format %s is not acceptable, please select one of: %v",
			string(format),
			options.GetKeys(),
		)
	}

	return oFormat, nil
}

// describeCmd represents the describe command
var describeCmd = &cobra.Command{
	SilenceErrors: false,
	SilenceUsage:  false,
	Use:           "describe",
	Short:         "Generate the data dictionary output",
	Long:          "Connect to the database and generate the data dictionary output in the selected format that by default is a txt expressed at the standard output.",
	Run: func(cmd *cobra.Command, args []string) {
		setLang(lang)

		oFormat, err := getFormat(format)
		if err != nil {
			cmd.PrintErr(err.Error())
			return
		}

		//output path definition
		if path == "" {
			path = "output." + strings.ToLower(format)
		}

		//execute extractions
		desc, err := services.Describe(uri, database, schema)
		if err != nil {
			cmd.PrintErr(err.Error())
			return
		}

		//set path for output and execute print
		err = oFormat.SetWriter(path)
		if err != nil {
			cmd.PrintErr(err.Error())
			return
		}

		services.PrintDocument(oFormat, *desc)

		if strings.ToUpper(format) != "DEFAULT" {
			cmd.Printf("%s document created.\n", strings.ToUpper(format))
		}
	},
}

// init function to initialize de commanda lide
func init() {
	rootCmd.AddCommand(describeCmd)

	describeCmd.PersistentFlags().StringVarP(&format, "format", "f", "DEFAULT", options.Message())
	describeCmd.PersistentFlags().StringVarP(&lang, "language", "l", "", "The language selected to the output file")

	describeCmd.PersistentFlags().StringVarP(&uri, "uri", "u", defaultDatabaseURI, "The database connection uri")

	describeCmd.PersistentFlags().StringVarP(&database, "database", "d", "postgres", "The database to be described")
	describeCmd.PersistentFlags().StringVarP(&schema, "schema", "s", "public", "The schema to be described")

	describeCmd.PersistentFlags().StringVarP(&path, "output", "o", "", "The output file path")
}
