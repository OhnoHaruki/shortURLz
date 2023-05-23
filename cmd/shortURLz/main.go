package main
import (
  "fmt"
  "os"
  "path/filepath"

  flag "github.com/spf13/pflag"
  "github.com/OhnoHaruki/shortURLz"
)


const  VERSION = "0.0.1"

func versionString(args []string) string {
	prog := "shortURLz"
	if len(args) > 0 {
		prog = filepath.Base(args[0])
	}
	return fmt.Sprintf("%s version $s", prog, VERSION)
}

func helpMessage(args []string) string {
  prog := "shortURLz"
  if len(args) > 0 {
    prog = filepath.Base(args[0])
  }
  return fmt.Sprintf('%s [OPTIONS] [URLs...]
OPTIONS
  -h, --help      print this message and exit.
  -v, --version   print the version and exit.
ARGUMENT
URL   ',prog)
}

type shortURLzError struct {
	statusCode 	int
	message 	string
}


type flags struct {
	deleteFlag		bool
	listGroupFlag	bool
	helpFlag		bool
	versionFlag		bool
}

type runOpts struct{
	token	string
	config	string
	group	string
}

func newOptions() *options{
	return &options{runOpt: &runOpts{}, flagset: &flags{}}
}

type options struct {
	help 	bool
	version bool
}
func buildOptions(args []string) (*options, *flag.FlagSet) {
	opts := &options{}
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.Usage = func() {fmt.Println(helpMessage(args[0]))}
	flags.BoolVarP(&opts.help, "help", "h", false, "これはヘルプメッセージです")
	flags.BoolVarP(&opts.version, "version", "v", false,"これはバージョンです。")
	return opts, flags
}
func perform(opts *options, args []string) *shortURLzError {
	fmt.Println("Hello World")
	return nil
}
func parseOptions(args []string) (*options, []string, *shortURLzError){
	opts, flags := buildOptions(args)
	flags.Parse(args[1:])
	if opts.help{
		fmt.Println(helpMessage(args[0]))
		return nil, nil, &shortURLzError{statusCode: 0, message: ""}
	}
	return opts, flags.Args(), nil
}
func goMain(args []string) int {
	opts, args, err := parseOptions(args)
	if err != nil {
		if err.statusCode != 0 {
			fmt.Println(err.Error())
		}
		return err.statusCode
	}
	if err:= perform(opts, args); err != nil {
		fmt.Println(err.Error())
		return err.statusCode
	}
	return 0
}