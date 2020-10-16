package resp

import "errors"

var (
	FORMATERROR      = errors.New("format error")
	INSERTERROR      = errors.New("insert failed")
	FINDONEERROR     = errors.New("find one failed")
	FINDERROR        = errors.New("find failed")
	FINDDECODEERROR  = errors.New("find decode failed")
	UPDATEEERROR     = errors.New("update failed")
	DELETEERROR      = errors.New("delete failed")
	MONGOERROR       = errors.New("connection failed")
	OBJECTIDERROR    = errors.New("objectID failed")
	REPEATERROR      = errors.New("Name repeated")
	CHECKREPEATERROR = errors.New("check repeat failed")
	STRCONVERROR     = errors.New("strconv failed")
	SEARCHERROR      = errors.New("search failed")
)
