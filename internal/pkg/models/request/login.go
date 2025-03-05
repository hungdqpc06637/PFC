package request

// REQUEST - FMEA

type LoginRequest struct {
	Userid   string `form:"USERID" json:"USERID"`
	Password string `form:"password" json:"password"`
}
