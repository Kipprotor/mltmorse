package morse

import "fmt"

// ErrNoEncoding is the error used when there is no representation
// Its primary use is inside Handlers
type ErrNoEncoding struct{ Text string }

// Error implements the error interface
func (e ErrNoEncoding) Error() string { return fmt.Sprintf("No encoding for: %q", e.Text) }

// averageSize is the average size of a morse char
const averageSize = 4.53 //Magic

/*
DefaultMorse is the default map used to convert between morse and text
The map contains all the standard codes defined as costants but doesn't include commands like Understood and Error
This map may remain constant.
*/
var DefaultMorse = MergeEncMap(LatinMorse, NumSymbolMorse)

var reverseDefaultMorse = reverseEncodingMap(DefaultMorse)

// DefaultConverter is the default converter, it uses the exported morse set and has an IgnoreHandler, the separation character is a space
// Lowercase letter are encoded as upper ones. DefaultConverter uses explicitly IgnoreHandler and adds the trailing separator
var DefaultConverter = NewConverter(
	DefaultMorse,

	WithCharSeparator(" "),
	WithWordSeparator("   "),
	WithLowercaseHandling(true),
	WithHandler(IgnoreHandler),
	WithTrailingSeparator(false),
)
