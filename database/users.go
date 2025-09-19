package database

type Users struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsOwner  bool   `json:"is_owner"`
}

var usersList []Users

func (u Users) Store() Users {
	if u.ID != 0 {
		return u
	}
	u.ID = len(usersList) + 1
	usersList = append(usersList, u)
	return u
}

func Find(email string, password string) *Users {
	for i := range usersList {
		if usersList[i].Email == email && usersList[i].Password == password {
			return &usersList[i]
		}
	}
	return nil
}
