// Package substring intends to add the very useful substr method we so much
// like from Perl (see: http://perldoc.perl.org/functions/substr.html)
package substring

import "log"

// Substr Extracts a substring out of EXPR and returns it. First character
// is at offset zero. If OFFSET is negative, starts that far back from the
// end of the string. If LENGTH is omitted, returns everything through the end
// of the string. If LENGTH is negative, leaves that many characters off the
// end of the string.
//
// 	substring.Substr(EXPRESSION, OFFSET, LENGTH)
//
// 	s      := "The black cat climbed the green tree";
// 	color  := substring.Substr(s, 4, 5)      # black
// 	middle := substring.Substr(s, 4, -11)    # black cat climbed the
// 	end    := substring.Substr(s, 14)        # climbed the green tree
// 	tail   := substring.Substr(s, -4)        # tree
// 	z      := substring.Substr(s, -4, 2)     # tr
func Substr(expr string, pos ...int) string {

	var offset int
	var length int
	//	var replacement string

	switch len(pos) {
	case 1:
		if pos[0] < 0 {
			offset = len(expr) + pos[0]
			length = pos[0] * -1
		} else {
			offset = pos[0]
			length = len(expr) - offset
		}
	case 2:
		if pos[0] < 0 {
			offset = len(expr) + pos[0]
			length = pos[1]
		} else {
			offset = pos[0]
			length = pos[1]
		}
	case 3:
		offset = pos[0]
		length = pos[1]
		// replacement = "ABC"
	default:
		log.Panicf("Wrong number of arguments for Substr %i", len(pos))
	}

	if offset > len(expr) {
		return "" //, &SubstringError{err: "Offset > than length"}
	}
	if offset < 0 {
		offset = len(expr) - offset
	}

	if length < 0 {
		length = (len(expr) + length - offset)
	}

	if (length + offset) > len(expr) {
		length = len(expr) - offset
	}

	if offset > len(expr) {
		return "" //, &SubstringError{err: "Offset > than length"}
	}

	return expr[offset : offset+length]

}
