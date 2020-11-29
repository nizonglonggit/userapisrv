package define

type RegisterUserReq struct {
	NickName        string `json:"nick_name"`
	PassWord        string `json:"password"`
	Gender          uint8  `json:"gender"`
	Email           string `json:"email"`
	HeadPortraitUrl string `json:"head_portrait_url"`
}
