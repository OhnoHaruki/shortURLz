package main

import "testing"

func Example_No_Argument() {
	goMain([]string{"./shortURLz"})
	// Output:
	// no token was given
}

func Example_Token() {
	goMain([]string{"./shortURLz", "--token"})
	// Output:
	// no token was given
}

func Example_Delete() {
	goMain([]string{"./shortURLz", "--delete"})
	// Output:
	// no token was given
}

func Example_Completion() {
	goMain([]string{"./shortURLz", "--generate-completions"})
	// Output:
	// do completion
	// no token was given
}

func Example_Help() {
	goMain([]string{"./shortURLz", "--help"})
	// Output:
	// shortURLz [OPTIONS] [URLs...]
	// OPTIONS
	//     -t, --token <TOKEN>      specify the token for the service. This option is mandatory.
	//     -c, --config <CONFIG>    specify the configuration file.
	//     -g, --group <GROUP>      specify the group name for the service. Default is "shortURLz"
	//     -d, --delete             delete the specified shorten URL.
	//     -h, --help               print this message and exit.
	//     -v, --version            print the version and exit.
	// ARGUMENT
	//     URL     specify the url for shortening. this arguments accept multiple values.
	//             if no arguments were specified, shortURLz prints the list of available shorten urls.
}

func Test_Main(t *testing.T) {
	if status := goMain([]string{"./shortURLz", "-v"}); status != 0 {
		t.Error("Expected 0, got ", status)
	}
}
