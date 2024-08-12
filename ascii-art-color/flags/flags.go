package flags

import (
	args2 "ascii/args"
	"ascii/fmtx"
	"ascii/help"
	"errors"
	"fmt"
	"os"
	"regexp"
)

// Args models the Text to be drawn and the defined Banner style to be used.
// Also includes commandline Flags passed to the program
type Args struct {
	Text, Banner string
	Flags        []Flag
}

// Flag models a given commandline flag, with a given Value, and any ExtraValue if it takes any
// Example, case `--color=red "kit"`, then Name=color, Value=red, ExtraValue=kit
type Flag struct {
	Name       string
	Value      string
	ExtraValue string
}

// ParseFlags checks the given program args for flags and attempts to extract the text to be
// drawn and the defined banner style
func ParseFlags(args []string) (out Args) {
	optionalValueArgs := []string{"color"}
	definedFlags := []string{"color", "output", "align"}

	var flags []Flag = nil
	var positionalArgs []string
	for i, arg := range args {
		if arg == "--" {
			continue
		} else if arg == "--help" {
			// should exit with usage
			help.PrintLongUsage()
		} else if MoreFlags(args, i) {
			previousFlag, ok := lastFlag(flags)
			if !isFlag(arg) && ok && contains(optionalValueArgs, previousFlag.Name) {
				flags[len(flags)-1].ExtraValue = arg
				continue
			} else if !isFlag(arg) && ok {
				fmtx.FatalErrorf(
					"detected more flags ahead, but flag %s doesn't take extra values\n",
					previousFlag.Name,
				)
			}
			out, err := matchFlag(arg, definedFlags)
			if err != nil {
				// fmt.Printf("error: invalid argument %q\n%v\n\n", arg, err)
				help.PrintUsage()
			}
			if !contains(definedFlags, out.Name) {
				fmt.Printf("undefined flag: %s\n", out.Name)
				os.Exit(1)
			}
			flags = append(flags, Flag{Name: out.Name, Value: out.Value})
		} else {
			// There are no more flags
			positionalArgs = args[i:]
			break
		}
	}

	previousFlag, ok := lastFlag(flags)
	if len(positionalArgs) >= 2 && ok && contains(optionalValueArgs, previousFlag.Name) {
		// one of these positional args is an extra value for the last flag
		flags[len(flags)-1].ExtraValue = positionalArgs[0]
		positionalArgs = positionalArgs[1:]
	}

	switch len(positionalArgs) {
	case 0:
		out.Text = ""
	case 1:
		out.Text = positionalArgs[0]
	case 2:
		out.Text = positionalArgs[0]
		out.Banner = positionalArgs[1]
	default:
		help.PrintUsage()
	}
	out.Flags = flags
	// default to the standard banner
	if out.Banner == "" {
		out.Banner = "standard"
	}
	if out.Text != "" {
		// interpret \ escape(s) in text
		out.Text = args2.Escape(out.Text)
	}

	return
}

// uses a series of regular expressions to match a valid command-line flag with a value from the
// input program argument, returning the matched flag if valid.
// Returns an error describing the failed flag interpretation
func matchFlag(s string, definedFlags []string) (Flag, error) {
	// Check if the given flag is value-less, such as, `--output`
	if flag, yes := isValueLessFlag(s); yes && contains(definedFlags, flag) {
		return Flag{}, errors.New(fmt.Sprintf("option: %q needs a value", flag))
	}

	flagRegexes := []struct {
		main, sub, msg, prepend string
	}{
		// Series of regex to match a valid option flag with a value, such as, --output=file.txt
		{`^(--)`, `.*`, "not a command option", ""},
		{`^([a-z]+)=`, `^([^=]+)`, "unrecognized option", "--"},
		{`^(.+)$`, `.*`, "option needs value", ""},
	}

	var flag Flag
	// attempt to match the flag regexes in order
	str := s
	for i, re := range flagRegexes {
		out, err := subRegex(re.main, re.sub, str, re.msg, re.prepend)
		if err != nil {
			return flag, err
		}

		switch i {
		case 1:
			// we have a valid flag name
			flag.Name = out.match
		case 2:
			// we have a valid flag value defined
			flag.Value = out.match
		}
		str = out.str
	}

	return flag, nil
}

// structured return type for the subRegex function; where, match is the matched regex group, and str is a
// string with the matched regex group removed
type match struct {
	str, match string
}

// tests the `main` regex on the given strings str, then returns the first matched group,
// and a string where the matched group has been removed.
// If the main regex match fails, uses the `sub` regex to generate a value, prepended by the prepend string,
// for the given error message template
func subRegex(main, sub, str, errorMsgTemplate, prepend string) (match, error) {
	// attempt to match the main regex
	var out match
	re := regexp.MustCompile(main)
	mainMatch := re.FindStringSubmatch(str)
	if mainMatch == nil {
		// main regex failed, extract error message value with the sub regex
		more := ""
		if sub != "" {
			re := regexp.MustCompile(sub)
			subMatch := re.FindStringSubmatch(str)

			if subMatch != nil && len(subMatch) > 1 {
				more = fmt.Sprintf(": %q", prepend+subMatch[1])
			}
		}

		return out, errors.New(fmt.Sprintf("%s%s", errorMsgTemplate, more))
	}
	// expect to match at least one group in the main regex
	if len(mainMatch) > 1 {
		out.match = mainMatch[1]
	} else {
		return out, errors.New(fmt.Sprintf("%s", errorMsgTemplate))
	}

	// remove the matched group
	out.str = re.ReplaceAllString(str, "")
	return out, nil
}

// MoreFlags checks whether there are more command-line flags from the ith index in args onwards
func MoreFlags(args []string, i int) bool {
	nilIndex := indexOf(args, "--")
	if nilIndex == -1 {
		nilIndex = len(args)
	}

	for i := i; i < len(args) && i >= 0 && i <= nilIndex; i++ {
		a := args[i]
		if isFlag(a) {
			return true
		}
	}

	return false
}

// checks whether the given argument might be a flag
func isFlag(arg string) bool {
	re := regexp.MustCompile(`^--`)
	return re.MatchString(arg)
}

// returns the index in flags where the given flag exists, -1 otherwise
func indexOf(flags []string, flag string) int {
	for i, f := range flags {
		if f == flag {
			return i
		}
	}
	return -1
}

// returns true iff the given flag exists in the set of flags, false otherwise
func contains(flags []string, flag string) bool {
	return indexOf(flags, flag) > -1
}

// returns the last flag in the given list of flags, with a true (okay) status if it exists, false otherwise
func lastFlag(array []Flag) (element Flag, ok bool) {
	if len(array) == 0 {
		return element, false
	}
	return array[len(array)-1], true
}

// checks whether the given argument is a flag whose value has not been specified with an `=` sign,
// returns the name of the given valueless flag with a true status; returns false otherwise
func isValueLessFlag(s string) (string, bool) {
	// valid flags only have lowercase names
	re := regexp.MustCompile(`^(--)([a-z]+)$`)
	match := re.FindStringSubmatch(s)
	if match != nil {
		if len(match) > 2 {
			return match[2], true
		}
		return "", true
	}
	return "", false
}
