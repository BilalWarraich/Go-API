package db

type UserNotFoundErr struct {
}

func (u *UserNotFoundErr) Error() string {
	return "user not found!"
}
