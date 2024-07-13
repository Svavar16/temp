package main

type Users struct {
	name string
}

// update: this take the second user and changed it's name
func (u Users) update(user *Users) {
	user.name = u.name
}

// update2: then does this change the first user?
func (u Users) update2(user *Users) {
	u.name = user.name
}
