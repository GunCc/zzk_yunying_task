package utils

var (
	RegisterVerify = Rules{
		"NickName": {
			NotEmpty(),
		},
		"Password": {
			NotEmpty(),
		},
	}

	LoginVerify = Rules{
		"NickName": {
			NotEmpty(),
		},
		"Password": {
			NotEmpty(),
		},
	}

	VideoVerity = Rules{
		"StartTime": {
			"regexp=([0-5][0-9]):([0-5][0-9])",
		},
		"EndTime": {
			"regexp=([0-5][0-9]):([0-5][0-9])",
		},
	}
)
