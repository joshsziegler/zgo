// log provides opinionated, simplified defaults.
//
// - Two logging modes: prod (default), and debug.
//   See the table below for what is output for each mode.
// - To turn on debug logging, set the environment variable 'DEBUG=TRUE'
// - Date and time are not appended (we assume they are provided by the backend logger such as SystemD).
// - Log level is output as a single character to make lines shorter and more scannable.
//
//
//        Info  Error  Fatal  Debug
// Prod   Y     Y      Y
// Debug  Y     Y      Y      Y
//
//
// Example 1:
//     log.Info("Hello")
//     I: Hello
//     log.Error("Flux capacitor not fluxxing")
//     E: Flux capacitor not fluxxing
//     log.Debug("Showing other message prefixes")
//     [ Debug messages are not printed in production. ]
//     log.Infof("Using formatting for variable Foo: %+v", foo)
//     I: Using formatting for variable Foo: FooStruct{Bar: 5, Valid: True}
//
// Example 2:
//     export DEBUG=TRUE
//     log.Info("Hello")
//     I: Hello
//     log.Debug("connecting to database")
//     D: connecting to database

package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var debug bool

func init() {
	debug := strings.ToLower(strings.TrimSpace(os.Getenv("debug"))) == "true"
	log.SetFlags(0) // Disable default flags to hide datetime and level prefixes
}

// Info logs a normal, informative message.
// For information that should not be logged in production, used Debug() instead.
func Info(args ...interface{}) {
	log.Print("I: " + fmt.Sprint(args...))
}

func Infof(format string, args ...interface{}) {
	log.Print("I: " + fmt.Sprintf(format, args...))
}

// Debug logs a message that should only be shown only in debug.
func Debug(args ...interface{}) {
	if debug {
		log.Print("D: " + fmt.Sprint(args...))
	}
}

func Debugf(format string, args ...interface{}) {
	if debug {
		log.Print("D: " + fmt.Sprintf(format, args...))
	}
}

// Error logs a message that is a recoverable error.
// For unrecoverable errors, use Fatal() instead.
func Error(args ...interface{}) {
	log.Print("E: " + fmt.Sprint(args...))
}

func Errorf(format string, args ...interface{}) {
	log.Print("E: " + fmt.Sprintf(format, args...))
}

// Fatal logs an unrecoverable error.
// For recoverable errors, use Error() instead.
func Fatal(args ...interface{}) {
	log.Fatal("F: " + fmt.Sprint(args...))
}

func Fatalf(format string, args ...interface{}) {
	log.Fatal("F: " + fmt.Sprintf(format, args...))
}
