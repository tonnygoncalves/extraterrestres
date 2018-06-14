package db

//Method to create system tables
func CrearTablas() bool {
	db, err := connnection()

	// Create a new table for Wheather
	var wheatherTable = `
	CREATE TABLE IF NOT EXISTS Wheather (
		day int(11) NOT NULL,
		wheather varchar(10) NOT NULL,
		maxRain tinyint(1) NOT NULL
	  );	
		`
	insert, err := db.Query(wheatherTable)

	// if there are some error, handle it
	if err != nil {
		panic(err.Error())
	}
	var cleanTable = "DELETE FROM wheather;"
	insert2, err := db.Query(cleanTable)

	defer insert.Close()
	defer insert2.Close()
	defer db.Close()
	return true
}
