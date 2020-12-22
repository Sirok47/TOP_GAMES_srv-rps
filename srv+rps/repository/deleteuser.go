package repository

func (r *TopGamesPostgres) DeleteUser(n string) error {
	_, err := r.db.Exec("delete from users where Name = $1", n)
	return err
}
