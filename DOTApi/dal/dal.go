package dal

import (
	"DOTApi/models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var dataSourceName = "root:@tcp(127.0.0.1:3306)/dot_scanner?parseTime=true"

func GetAllScans() []models.Scan{
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("SELECT `scan`.`id`, `scan`.`scan_type_id`, `scan`.`latitude`, `scan`.`longitude`, `scan`.`expires_on` FROM `dot_scanner`.`scan` WHERE `scan`.`deleted` = 0;")

	if err != nil {
		panic(err.Error())
	}

	var scans []models.Scan

	for results.Next() {
		var scan models.Scan

		err = results.Scan(&scan.Id, &scan.ScanTypeId, &scan.Latitude, &scan.Longitude, &scan.ExpiresOn)
		if err != nil {
			return []models.Scan{}
		}

		scans = append(scans, scan)
	}

	err = db.Close()
	if err != nil {
		panic(err.Error())
	}

	return scans
}

func GetAllScansByUserId(userId int64) []models.Scan{
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("SELECT `scan`.`id`, `scan`.`scan_type_id`, `scan`.`latitude`, `scan`.`longitude`, `scan`.`expires_on` FROM `dot_scanner`.`scan` WHERE `scan`.`created_by_user_id` = ?;", userId)

	if err != nil {
		panic(err.Error())
	}

	var scans []models.Scan

	for results.Next() {
		var scan models.Scan

		err = results.Scan(&scan.Id, &scan.ScanTypeId, &scan.Latitude, &scan.Longitude, &scan.ExpiresOn)
		if err != nil {
			return []models.Scan{}
		}

		scans = append(scans, scan)
	}

	err = db.Close()
	if err != nil {
		panic(err.Error())
	}

	return scans
}


func GetAllScanTypes () []models.ScanType {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	results, err := db.Query("SELECT `scan_type`.`id`, `scan_type`.`name`, `scan_type`.`is_paid_version` = b'1', `scan_type`.`default_expiration_time` FROM `dot_scanner`.`scan_type` WHERE `scan_type`.`deleted` = 0;")

	if err != nil {
		panic(err.Error())
	}

	var scanTypes []models.ScanType

	for results.Next() {
		var scanType models.ScanType

		err = results.Scan(&scanType.Id, &scanType.Name, &scanType.IsPaidVersion, &scanType.DefaultExpirationTime)
		if err != nil {
			panic(err.Error())
		}

		scanTypes = append(scanTypes, scanType)
	}

	err = db.Close()
	if err != nil {
		panic(err.Error())
	}

	return scanTypes
}

func GetScanById(scanId int64) models.Scan {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	var scan models.Scan

	err = db.QueryRow("SELECT `scan`.`id`, `scan`.`scan_type_id`, `scan`.`latitude`, `scan`.`longitude`, `scan`.`expires_on` FROM `dot_scanner`.`scan` WHERE `scan`.`id` = ?;", scanId).Scan(&scan.Id, &scan.ScanTypeId, &scan.Latitude, &scan.Longitude, &scan.ExpiresOn)
	if err != nil {
		return models.Scan{}
	}

	err = db.Close()
	if err != nil {
		return models.Scan{}
	}

	return scan
}

func GetScanTypeById(scanTypeId int64) models.ScanType {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err.Error())
	}

	var scanType models.ScanType

	err = db.QueryRow("SELECT `scan_type`.`id`, `scan_type`.`name`, `scan_type`.`is_paid_version` = b'1', `scan_type`.`default_expiration_time` FROM `dot_scanner`.`scan_type` WHERE `scan_type`.`id` = ?;", scanTypeId).Scan(&scanType.Id, &scanType.Name, &scanType.IsPaidVersion, &scanType.DefaultExpirationTime)
	if err != nil {
		return models.ScanType{}
	}

	err = db.Close()
	if err != nil {
		return models.ScanType{}
	}

	return scanType
}

func InsertScan(scan models.Scan) int64 {
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	stmt, err := db.Prepare("INSERT INTO `dot_scanner`.`scan` (`scan_type_id`, `latitude`, `longitude`, `expires_on`, `created_by_user_id`) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(scan.Id, scan.ScanTypeId, scan.Latitude, scan.Longitude, scan.ExpiresOn, scan.CreatedByUserId)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)

	return id
}

func InsertUser(user models.User) int64 {
	db, err := sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db)

	stmt, err := db.Prepare("INSERT INTO `dot_scanner`.`user`(`email`, `password`, `phone_number`, `paid_member`, `created_by_user_id`) VALUES (?, ?, ?, ?, ?);")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(user.Email, user.Password, user.PhoneNumber, user.PaidMember, user.CreatedByUserId)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()

	if err != nil {
		panic(err.Error())
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			panic(err.Error())
		}
	}(stmt)

	return id
}