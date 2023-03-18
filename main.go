package main

import (
	// Import this so we don't have to use qm.Limit etc.
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func main() {
	// Open handle to database like normal
	db, err := sql.Open("postgres", "dbname=fun user=abc")
	if err != nil {
		return err
	}

	// If you don't want to pass in db to all generated methods
	// you can use boil.SetDB to set it globally, and then use
	// the G variant methods like so (--add-global-variants to enable)
	boil.SetDB(db)
	users, err := models.Users().AllG(ctx)

	// Query all users
	users, err := models.Users().All(ctx, db)

	// Panic-able if you like to code that way (--add-panic-variants to enable)
	users := models.Users().AllP(db)

	// More complex query
	users, err := models.Users(Where("age > ?", 30), Limit(5), Offset(6)).All(ctx, db)

	// Ultra complex query
	users, err := models.Users(
		Select("id", "name"),
		InnerJoin("credit_cards c on c.user_id = users.id"),
		Where("age > ?", 30),
		AndIn("c.kind in ?", "visa", "mastercard"),
		Or("email like ?", `%aol.com%`),
		GroupBy("id", "name"),
		Having("count(c.id) > ?", 2),
		Limit(5),
		Offset(6),
	).All(ctx, db)

	// Use any "boil.Executor" implementation (*sql.DB, *sql.Tx, data-dog mock db)
	// for any query.
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	users, err := models.Users().All(ctx, tx)

	// Relationships
	user, err := models.Users().One(ctx, db)
	if err != nil {
		return err
	}
	movies, err := user.FavoriteMovies().All(ctx, db)

	// Eager loading
	users, err := models.Users(Load("FavoriteMovies")).All(ctx, db)
	if err != nil {
		return err
	}
	fmt.Println(len(users.R.FavoriteMovies))
}
