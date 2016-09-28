package logrus

import "testing"

// TestLogrus exercises the customizable logging package named logrus
// FIXME: Need assertion test to make sure output is as expected
func TestLogrus(t *testing.T) {
	t.Logf("Start log format is %s", GetFormat())
	Println("This should be in text format")

	FormatText()
	Println("This should still be in text format")

	FormatJSON()

	Println("This should be in json format")

	FormatJSON()
	Println("This should still be in json format")

	FormatText()

	Println("This should be back in text format")

	FormatNone()

	Println("This should not print at all")

	FormatNone()

	Println("This should not be printed either")

	// WARNING: These will not go to log with no formatter
	Error("Simulate fail. Will it print with no formatter?")

	FormatText()
	Println("This also should be in text")

	t.Logf("End log format is %s", GetFormat())
}

// TestLogrusPanic tests recoverable panic messages with no formatter.
// NOTE: this does not print a stack trace like panic() function
func TestLogrusPanic(t *testing.T) {
	t.Logf("Testing panic log with no formatter")
	defer func() {
		if gotFail := recover(); gotFail != nil {
			t.Logf("Test panic recovery: %v", gotFail)
		}
	}()

	FormatText()
	Println("Testing panic log")
	FormatNone()
	// Panic logs will still print
	Panic("Simulate panic. Will it print with no formatter?")
}

// TestLogrusFatal tests unrecoverable panic messages with no formatter
// Commented out since test purposely ends in a test fault
/*
func TestLogrusFatal(t *testing.T) {
	t.Logf("Testing fatal log with no formatter")
	defer func() {
		t.Logf("End testing fatal log with no formatter")
                if gotFail := recover(); gotFail != nil {
                        t.Logf("Test fatal recovery: %v", gotFail)
                }
        }()

	FormatText();
	Println("Testing fatal log")
	FormatNone();
	// Fatal logs will still print
	Fatal("Simulate fatal. Will it print with no formatter?")
}
*/
