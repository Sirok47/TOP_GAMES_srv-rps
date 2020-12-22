package repository

func (r *TopGamesPostgres) UpdateUser(n string,p string) error {
	_, err := r.db.Exec("UPDATE users SET Password = $2 WHERE Name = $1",n,p)
	return err
}
