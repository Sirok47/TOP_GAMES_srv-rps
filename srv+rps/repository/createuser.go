package repository

func (r *TopGamesPostgres) CreateUser(n string,p string) error {
	_, err := r.db.Exec("insert into users (Name,Password) values ($1,$2)",n,p)
	return err
}
