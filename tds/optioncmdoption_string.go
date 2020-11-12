// SPDX-FileCopyrightText: 2020 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=OptionCmdOption"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_OPT_UNUSED-0]
	_ = x[TDS_OPT_DATEFIRST-1]
	_ = x[TDS_OPT_TEXTSIZE-2]
	_ = x[TDS_OPT_STAT_TIME-3]
	_ = x[TDS_OPT_STAT_IO-4]
	_ = x[TDS_OPT_ROWCOUNT-5]
	_ = x[TDS_OPT_NATLANG-6]
	_ = x[TDS_OPT_DATEFORMAT-7]
	_ = x[TDS_OPT_ISOLATION-8]
	_ = x[TDS_OPT_AUTHON-9]
	_ = x[TDS_OPT_CHARSET-10]
	_ = x[TDS_OPT_PLAN-11]
	_ = x[TDS_OPT_ERRLVL-12]
	_ = x[TDS_OPT_SHOWPLAN-13]
	_ = x[TDS_OPT_NOEXEC-14]
	_ = x[TDS_OPT_ARITHIGNOREON-15]
	_ = x[TDS_OPT_ARITHABORTON-16]
	_ = x[TDS_OPT_PARSEONLY-17]
	_ = x[TDS_OPT_ESTIMATE-18]
	_ = x[TDS_OPT_GETDATA-19]
	_ = x[TDS_OPT_NOCOUNT-20]
	_ = x[TDS_OPT_FORCEPLAN-21]
	_ = x[TDS_OPT_FORMATONLY-22]
	_ = x[TDS_OPT_CHAINXACTS-23]
	_ = x[TDS_OPT_CURCLOSEONXACT-24]
	_ = x[TDS_OPT_FIPSFLAG-25]
	_ = x[TDS_OPT_RESTREES-26]
	_ = x[TDS_OPT_IDENTITYON-27]
	_ = x[TDS_OPT_CURREAD-28]
	_ = x[TDS_OPT_CURWRITE-29]
	_ = x[TDS_OPT_IDENTITYOFF-30]
	_ = x[TDS_OPT_AUTHOFF-31]
	_ = x[TDS_OPT_ANSINULL-32]
	_ = x[TDS_OPT_QUOTED_IDENT-33]
	_ = x[TDS_OPT_ANSIPERM-34]
	_ = x[TDS_OPT_STR_RTRUNC-35]
	_ = x[TDS_OPT_SORTMERGE-36]
	_ = x[TDS_OPT_JTC-37]
	_ = x[TDS_OPT_CLIENTREALNAME-38]
	_ = x[TDS_OPT_CLIENTHOSTNAME-39]
	_ = x[TDS_OPT_CLIENTAPPLNAME-40]
	_ = x[TDS_OPT_IDENTITYUPD_ON-41]
	_ = x[TDS_OPT_IDENTITYUPD_OFF-42]
	_ = x[TDS_OPT_NODATA-43]
	_ = x[TDS_OPT_CIPHERTEXT-44]
	_ = x[TDS_OPT_SHOW_FI-45]
	_ = x[TDS_OPT_HIDE_VCC-46]
	_ = x[TDS_OPT_LOBLOCATOR-47]
	_ = x[TDS_REQ_LOBLOCATOR-48]
	_ = x[TDS_OPT_LOBLOCATORFETCHSIZE-49]
	_ = x[TDS_OPT_ISOLATION_MODE-102]
}

const (
	_OptionCmdOption_name_0 = "TDS_OPT_UNUSEDTDS_OPT_DATEFIRSTTDS_OPT_TEXTSIZETDS_OPT_STAT_TIMETDS_OPT_STAT_IOTDS_OPT_ROWCOUNTTDS_OPT_NATLANGTDS_OPT_DATEFORMATTDS_OPT_ISOLATIONTDS_OPT_AUTHONTDS_OPT_CHARSETTDS_OPT_PLANTDS_OPT_ERRLVLTDS_OPT_SHOWPLANTDS_OPT_NOEXECTDS_OPT_ARITHIGNOREONTDS_OPT_ARITHABORTONTDS_OPT_PARSEONLYTDS_OPT_ESTIMATETDS_OPT_GETDATATDS_OPT_NOCOUNTTDS_OPT_FORCEPLANTDS_OPT_FORMATONLYTDS_OPT_CHAINXACTSTDS_OPT_CURCLOSEONXACTTDS_OPT_FIPSFLAGTDS_OPT_RESTREESTDS_OPT_IDENTITYONTDS_OPT_CURREADTDS_OPT_CURWRITETDS_OPT_IDENTITYOFFTDS_OPT_AUTHOFFTDS_OPT_ANSINULLTDS_OPT_QUOTED_IDENTTDS_OPT_ANSIPERMTDS_OPT_STR_RTRUNCTDS_OPT_SORTMERGETDS_OPT_JTCTDS_OPT_CLIENTREALNAMETDS_OPT_CLIENTHOSTNAMETDS_OPT_CLIENTAPPLNAMETDS_OPT_IDENTITYUPD_ONTDS_OPT_IDENTITYUPD_OFFTDS_OPT_NODATATDS_OPT_CIPHERTEXTTDS_OPT_SHOW_FITDS_OPT_HIDE_VCCTDS_OPT_LOBLOCATORTDS_REQ_LOBLOCATORTDS_OPT_LOBLOCATORFETCHSIZE"
	_OptionCmdOption_name_1 = "TDS_OPT_ISOLATION_MODE"
)

var (
	_OptionCmdOption_index_0 = [...]uint16{0, 14, 31, 47, 64, 79, 95, 110, 128, 145, 159, 174, 186, 200, 216, 230, 251, 271, 288, 304, 319, 334, 351, 369, 387, 409, 425, 441, 459, 474, 490, 509, 524, 540, 560, 576, 594, 611, 622, 644, 666, 688, 710, 733, 747, 765, 780, 796, 814, 832, 859}
)

func (i OptionCmdOption) String() string {
	switch {
	case i <= 49:
		return _OptionCmdOption_name_0[_OptionCmdOption_index_0[i]:_OptionCmdOption_index_0[i+1]]
	case i == 102:
		return _OptionCmdOption_name_1
	default:
		return "OptionCmdOption(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}