package models

import "github.com/matheusjost/go-bettings-api/db"

func Insert(bet Bet) (id int64, err error) {
	conn, err := db.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `INSERT INTO bets (event, selection, stake, odds, status) VALUES ($1, $2, $3, $4, $5) RETURNING ID`

	err = conn.QueryRow(query, bet.Event, bet.Selection, bet.Stake, bet.Odds, bet.Status).Scan(&id)

	return id, err
}

func Get(id int64) (bet Bet, err error) {
	conn, err := db.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT * FROM bets WHERE id = $1`

	row := conn.QueryRow(query, id)

	err = row.Scan(&bet.ID, &bet.Event, &bet.Selection, &bet.Stake, &bet.Odds, &bet.Status)

	return
}

func GetAll() (bets []Bet, err error) {
	conn, err := db.OpenConn()
	if err != nil {
		return
	}
	defer conn.Close()

	query := `SELECT * FROM bets`

	rows, err := conn.Query(query)
	if err != nil {
		return
	}

	for rows.Next() {
		var bet Bet

		err = rows.Scan(&bet.ID, &bet.Event, &bet.Selection, &bet.Stake, &bet.Odds, &bet.Status)

		if err != nil {
			continue
		}

		bets = append(bets, bet)
	}

	return
}

func Update(id int64, bet Bet) (int64, error) {
	conn, err := db.OpenConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	query := `UPDATE bets SET stake = $2, odds = $3, status = $4 WHERE id = $1`

	res, err := conn.Exec(query, id, bet.Stake, bet.Odds, bet.Status)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}

func Delete(id int64) (int64, error) {
	conn, err := db.OpenConn()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	query := `DELETE FROM bets WHERE id = $1`

	res, err := conn.Exec(query, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()

}
