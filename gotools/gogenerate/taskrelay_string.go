// Code generated by "stringer -type=TaskRelay,Animal"; DO NOT EDIT.

package gogenerate

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[InvalidRelay-0]
	_ = x[MoveAwayRobotTaskRelay-1]
}

const _TaskRelay_name = "InvalidRelayMoveAwayRobotTaskRelay"

var _TaskRelay_index = [...]uint8{0, 12, 34}

func (i TaskRelay) String() string {
	if i >= TaskRelay(len(_TaskRelay_index)-1) {
		return "TaskRelay(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TaskRelay_name[_TaskRelay_index[i]:_TaskRelay_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Dog-0]
	_ = x[Pig-1]
	_ = x[Cat-2]
	_ = x[Tiger-3]
	_ = x[Lion-4]
}

const _Animal_name = "DogPigCatTigerLion"

var _Animal_index = [...]uint8{0, 3, 6, 9, 14, 18}

func (i Animal) String() string {
	if i < 0 || i >= Animal(len(_Animal_index)-1) {
		return "Animal(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Animal_name[_Animal_index[i]:_Animal_index[i+1]]
}
