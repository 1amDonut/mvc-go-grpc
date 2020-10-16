package resp

import "time"

// resp -
type resp struct {
	Data   interface{} `json:"data" bson:"data"`
	Status status      `json:"status" bson:"status"`
}

type status struct {
	Code int    `json:"code" bson:"code"`
	Msg  string `json:"msg" bson:"msg"`
	Time int64  `json:"time" bson:"time"`
}

func R(data interface{}) resp {
	return resp{
		Data: data,
		Status: status{
			Code: 0,
			Msg:  "",
			Time: time.Now().Unix(),
		},
	}
}

func E(e error, code int) resp {
	return resp{
		Data: nil,
		Status: status{
			Code: code,
			Msg:  e.Error(),
			Time: time.Now().Unix(),
		},
	}
}
