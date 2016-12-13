package setup

import (
	"bufio"
	"encoding/csv"
	"github.com/svarlamov/bintrad/models"
	"os"
	"strconv"
)

func setupUsers() error {
	in, err := os.Open("./users.csv")
	if err != nil {
		return err
	}
	r := csv.NewReader(bufio.NewReader(in))

	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for i := 1; i < len(records); i++ {
		idInt, err := strconv.ParseInt(records[i][0], 10, 64)
		if err != nil {
			return err
		}
		grInt, err := strconv.ParseInt(records[i][3], 10, 64)
		if err != nil {
			return err
		}

		x := models.User{
			Username:        records[i][0],
			FirstName:       records[i][2],
			LastName:        records[i][1],
			Grade:           grInt,
			Gender:          records[i][4],
			StudentId:       idInt,
			Passkey:         records[i][5],
			Email:           records[i][0] + "@hkis.edu.hk",
			StartingBalance: 1000000,
		}
		err = x.Create()
		if err != nil {
			return err
		}

	}
	return nil
}
