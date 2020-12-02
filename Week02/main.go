package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/sirupsen/logrus"
)

var (
	db *sql.DB
)

// User string
type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func init() {
	var err error
	db, err = sql.Open("mysql", "root:****@(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatalf("sql.Open err: %v", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db.Ping  err: %v", err)
		return
	}

	log.Info("success init mysql")
	return
}

func getUserDao(id int) (*User, error) {
	user := &User{}
	getUserSQL := `
			SELECT
				id,
				name
			FROM user
			WHERE id=?
		`
	row := db.QueryRow(getUserSQL, id)
	err := row.Scan(&user.ID, &user.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return user, nil

}

func getUserService(id int) (*User, error) {
	return getUserDao(id)
}

func getUserAPI(w http.ResponseWriter, r *http.Request) {
	// 参数校验
	r.ParseForm()
	idS, ok := r.Form["id"]
	if !ok || len(idS) < 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	value := idS[0]
	id, err := strconv.Atoi(value)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 查询
	user, err := getUserService(id)
	if err != nil {
		log.Errorf("%+v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("!"))
		return
	}
	if user == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("not found user"))
		return
	}

	// 写入
	buf, err := json.Marshal(user)
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, "json marshal failed"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(buf)
	if err != nil {
		log.Errorf("%+v", errors.Wrap(err, "write body failed"))
		w.WriteHeader(http.StatusInternalServerError)
	}
	return
}

func server() {
	addr := "127.0.0.1:9090"
	mux := http.NewServeMux()
	mux.HandleFunc("/api/user", getUserAPI)
	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}
	log.Infof("server listen: %s", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	server()
}
