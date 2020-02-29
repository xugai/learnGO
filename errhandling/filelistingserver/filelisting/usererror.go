package filelisting

type UserError string

func (u UserError) Error() string {
	return u.Message()
}

func (u UserError) Message() string {
	return string(u)
}
