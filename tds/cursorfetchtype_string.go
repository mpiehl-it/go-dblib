// SPDX-FileCopyrightText: 2021 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=CursorFetchType"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_CUR_NEXT-1]
	_ = x[TDS_CUR_PREV-2]
	_ = x[TDS_CUR_FIRST-3]
	_ = x[TDS_CUR_LAST-4]
	_ = x[TDS_CUR_ABS-5]
	_ = x[TDS_CUR_REL-6]
}

const _CursorFetchType_name = "TDS_CUR_NEXTTDS_CUR_PREVTDS_CUR_FIRSTTDS_CUR_LASTTDS_CUR_ABSTDS_CUR_REL"

var _CursorFetchType_index = [...]uint8{0, 12, 24, 37, 49, 60, 71}

func (i CursorFetchType) String() string {
	i -= 1
	if i >= CursorFetchType(len(_CursorFetchType_index)-1) {
		return "CursorFetchType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _CursorFetchType_name[_CursorFetchType_index[i]:_CursorFetchType_index[i+1]]
}