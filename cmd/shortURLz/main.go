package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/OhnoHaruki/shortURLz"
	flag "github.com/spf13/pflag"
)

const VERSION = "0.1.11"

func versionString(args []string) string {
	prog := "shortURLz"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version %s", prog, VERSION)
}

/*
helpMessage prints the help message.
This function is used in the small tests, so it may be called with a zero-length slice.
*/
func helpMessage(args []string) string {
	prog := "shortURLz"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf(`%s [OPTIONS] [URLs...]
OPTIONS
    -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
    -c, --config <CONFIG>    specify the configuration file.
    -g, --group <GROUP>      specify the group name for the service. Default is "shortURLz"
    -d, --delete             delete the specified shorten URL.
    -h, --help               print this message and exit.
    -v, --version            print the version and exit.
ARGUMENT
    URL     specify the url for shortening. this arguments accept multiple values.
            if no arguments were specified, shortURLz prints the list of available shorten urls.`, prog)
}

type ShortURLzError struct {
	statusCode int
	message    string
}

func (e ShortURLzError) Error() string {
	return e.message
}

type flags struct {
	deleteFlag    bool
	listGroupFlag bool
	helpFlag      bool
	versionFlag   bool
}

type runOpts struct {
	token  string
	config string
	group  string
}

/*
This struct holds the values of the options.
*/
type options struct {
	runOpt  *runOpts
	flagSet *flags
}

func newOptions() *options {
	return &options{runOpt: &runOpts{}, flagSet: &flags{}}
}

func (opts *options) mode(args []string) shortURLz.Mode {
	switch {
	case opts.flagSet.listGroupFlag:
		return shortURLz.ListGroup
	case len(args) == 0:
		return shortURLz.List
	case opts.flagSet.deleteFlag:
		return shortURLz.Delete
	default:
		return shortURLz.Shorten
	}
}

var completions bool

/*
Define the options and return the pointer to the options and the pointer to the flagset.
*/
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := newOptions()
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() { fmt.Println(helpMessage(args)) }
	flags.StringVarP(&opts.runOpt.token, "token", "t", "", "specify the token for the service. This option is mandatory.")
	flags.StringVarP(&opts.runOpt.config, "config", "c", "", "specify the configuration file.")
	flags.StringVarP(&opts.runOpt.group, "group", "g", "", "specify the group name for the service. Default is \"shortURLz\"")
	flags.BoolVarP(&opts.flagSet.listGroupFlag, "list-group", "L", false, "list the groups. This is hidden option.")
	flags.BoolVarP(&opts.flagSet.deleteFlag, "delete", "d", false, "delete the specified shorten URL.")
	flags.BoolVarP(&opts.flagSet.helpFlag, "help", "h", false, "print this message and exit.")
	flags.BoolVarP(&opts.flagSet.versionFlag, "version", "v", false, "print the version and exit.")
	flags.BoolVarP(&completions, "generate-completions", "", false, "generate completions")
	flags.MarkHidden("generate-completions")
	return opts, flags
}

/*
parseOptions parses options from the given command line arguments.
*/
func parseOptions(args []string) (*options, []string, *ShortURLzError) {
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])

	if completions {
		fmt.Println("do completion")
		GenerateCompletion(flags)
	}
	if opts.flagSet.helpFlag {
		fmt.Println(helpMessage(args))
		return nil, nil, &ShortURLzError{statusCode: 0, message: ""}
	}
	if opts.flagSet.versionFlag {
		fmt.Println(versionString(args))
		return nil, nil, &ShortURLzError{statusCode: 0, message: ""}
	}
	if opts.runOpt.token == "" {
		return nil, nil, &ShortURLzError{statusCode: 3, message: "no token was given"}
	}
	return opts, flags.Args(), nil
}

func shortenEach(bitly *shortURLz.Bitly, config *shortURLz.Config, url string) error {
	result, err := bitly.Shorten(config, url)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}

func deleteEach(bitly *shortURLz.Bitly, config *shortURLz.Config, url string) error {
	return bitly.Delete(config, url)
}

func listUrls(bitly *shortURLz.Bitly, config *shortURLz.Config) error {
	urls, err := bitly.List(config)
	if err != nil {
		return err
	}
	for _, url := range urls {
		fmt.Println(url)
	}
	return nil
}

func listGroups(bitly *shortURLz.Bitly, config *shortURLz.Config) error {
	groups, err := bitly.Groups(config)
	if err != nil {
		return err
	}
	for i, group := range groups {
		fmt.Printf("GUID[%d] %s\n", i, group.Guid)
	}
	return nil
}

func performImpl(args []string, executor func(url string) error) *ShortURLzError {
	for _, url := range args {
		err := executor(url)
		if err != nil {
			return makeError(err, 3)
		}
	}
	return nil
}

func perform(opts *options, args []string) *ShortURLzError {
	bitly := shortURLz.NewBitly(opts.runOpt.group)
	config := shortURLz.NewConfig(opts.runOpt.config, opts.mode(args))
	config.Token = opts.runOpt.token
	switch config.RunMode {
	case shortURLz.List:
		err := listUrls(bitly, config)
		return makeError(err, 1)
	case shortURLz.ListGroup:
		err := listGroups(bitly, config)
		return makeError(err, 2)
	case shortURLz.Delete:
		return performImpl(args, func(url string) error {
			return deleteEach(bitly, config, url)
		})
	case shortURLz.Shorten:
		return performImpl(args, func(url string) error {
			return shortenEach(bitly, config, url)
		})
	}
	return nil
}

func makeError(err error, status int) *ShortURLzError {
	if err == nil {
		return nil
	}
	ue, ok := err.(*ShortURLzError)
	if ok {
		return ue
	}
	return &ShortURLzError{statusCode: status, message: err.Error()}
}

func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err := perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
