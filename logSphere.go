package logrus

//
//	Package logSphere is used with the package logrus to set up standard
//	compatible logging customized for use with Sphere. It currently allows
//	text, json, or none output, with text as default. The package logrus
//	itself uses mutexes to prevent message contension/clobbering. It has
//	log level outputs as well as hooks to send output to central loggers.
//

import "os"

var outFormat = "text"  // default of plain text output
var outFile = os.Stdout // default output to stdout

// LogSphereInit is Sphere specific log initialization
func LogSphereInit() {
	var envLogFormat string
	if envLogFormat = os.Getenv("SPHERE_LOG_FORMAT"); envLogFormat == "json" {
		outFormat = "json"
	} else if envLogFormat == "none" {
		outFormat = "none"
	}
	if outFormat == "json" {
		SetFormatter(&JSONFormatter{})
	} else if outFormat == "none" {
		SetFormatter(&NoneFormatter{})
	} else {
		SetFormatter(&TextFormatter{})
	}
	SetOutput(outFile)
	SetLevel(InfoLevel)
}

// GetFormat retrieves the current state of the output format
func GetFormat() string {
	return outFormat
}

// FormatJSON turns on logging output and output log as json
func FormatJSON() {
	if outFormat != "json" {
		if outFormat != "none" {
			Info("About to enable json logging.")
		}
		SetFormatter(&JSONFormatter{})
		outFormat = "json"
	}
}

// FormatText turns on logging output and output log as text
func FormatText() {
	if outFormat != "text" {
		if outFormat != "none" {
			Info("About to enable text logging.")
		}
		SetFormatter(&TextFormatter{})
		outFormat = "text"
	}
}

// FormatNone turns off logging output
// WARNING: This turns off all log output, including errors!
func FormatNone() {
	if outFormat != "none" {
		Info("Turning off all logging.")
		SetFormatter(&NoneFormatter{})
		outFormat = "none"
	}
}
