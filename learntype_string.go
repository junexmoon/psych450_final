// // Code generated by "stringer -type=LearnType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Hebbian-0]
	_ = x[ErrorDriven-1]
	_ = x[ErrorHebbIn-2]
	_ = x[LearnTypeN-3]
}

const _LearnType_name = "HebbianErrorDrivenErrorHebbInLearnTypeN"

var _LearnType_index = [...]uint8{0, 7, 18, 29, 39}

func (i LearnType) String() string {
	if i < 0 || i >= LearnType(len(_LearnType_index)-1) {
		return "LearnType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _LearnType_name[_LearnType_index[i]:_LearnType_index[i+1]]
}
