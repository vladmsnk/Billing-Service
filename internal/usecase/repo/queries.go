package repo

const (
	InsertUser = `
	INSERT INTO users (id, name, username, email, password) VALUES
		$1, $2. $3, $4, $5 ON CONFLICT (id) DO NOTHING;
	`

	CheckUserExists = `SELECT * FROM users WHERE email = $1;`

	GetWalletByUserID = `SELECT wallet_id, amount 
		FROM wallet WHERE user_id = $1 LIMIT 1;`
)
