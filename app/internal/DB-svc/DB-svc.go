package dbsvc

import (
	"crud/internal/models"
	"database/sql"
	"log"

	"github.com/google/uuid"
)

type (
	Storage interface {
		Create(u models.User) (string, error)
		List() ([]models.User, error)
		GetUser(ID string) ([]models.User, error)
		Update(ID, Email, Name string) error
		Delete(ID string) error
	}
	dbsvc interface {
		GetClient() (con *sql.DB)
	}
)

type store struct {
	db dbsvc
}

func NewDB(db dbsvc) Storage {
	return &store{db: db}
}

func (s *store) Create(u models.User) (string, error) {
	con := s.db.GetClient()
	u.ID = uuid.New().String()
	_, err := con.Exec("INSERT INTO Users (ID, Email, Name) VALUES (?, ?, ?)", u.ID, u.Email, u.Name)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return u.ID, nil
}

func (s *store) List() ([]models.User, error) {
	con := s.db.GetClient()
	var list []models.User
	rows, err := con.Query("SELECT * FROM Users")
	if err != nil {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	var user models.User
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			log.Println(err)
			return list, err
		}
		list = append(list, user)
	}
	return list, err
}

func (s *store) GetUser(ID string) ([]models.User, error) {
	con := s.db.GetClient()
	var list []models.User
	rows, err := con.Query("SELECT * FROM Users WHERE ID = ?", ID)
	if err != nil {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	var user models.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Email, &user.Name)
		if err != nil {
			log.Println(err)
			return list, err
		}
		list = append(list, user)
	}

	return list, err
}

func (s *store) Delete(ID string) error {
	con := s.db.GetClient()
	_, err := con.Exec("DELETE FROM Users WHERE ID = ?", ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func (s *store) Update(ID, Email, Name string) error {
	con := s.db.GetClient()
	_, err := con.Exec("UPDATE Users SET Email = ?, Name = ? WHERE ID = ?", Email, Name, ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}
