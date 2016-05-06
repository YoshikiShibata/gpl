// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package eval

import "fmt"

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) String() string {
	var result = fmt.Sprintf("%s(", c.fn)
	for i, arg := range c.args {
		if i > 0 {
			result += ", "
		}
		result += arg.String()
	}
	return result + ")"
}

func (l list) String() string {
	var result = fmt.Sprintf("%s[", l.fn)
	for i, arg := range l.args {
		if i > 0 {
			result += ", "
		}
		result += arg.String()
	}
	return result + "]"
}
