// Package domain contains the generated code from the protoc
// compiler using the proto files. It also has extra methods to
// handle the serializing of Timestamp from the protobuf to
// sql timestamp format.
//
// Copyright 2020 Olumide Ogundele. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package domain

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/golang/protobuf/ptypes"

	// register driver needed for postgreSQL
	_ "github.com/lib/pq"
)

// Scan method is needed to implement the Scanner interface used by
// by the sql package to read data from a database to a Go type.
// ref: https://golang.org/pkg/database/sql/#Rows.Scan
func (ts *Timestamp) Scan(value interface{}) error {
	switch t := value.(type) {
	case time.Time:
		var err error
		ts.Timestamp, err = ptypes.TimestampProto(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("not a protobuf Timestamp")
	}
	return nil
}

// Value method helps implement the Valuer interface to be able to
// convert the protobuf Timestamp type to an sql driver Value because
// Gorm also support the interface, it makes it possible to create this
// field in the postgres database.
// ref: https://golang.org/pkg/database/sql/driver/#Valuer
//     https://gorm.io/docs/models.html#Declaring-Models
func (ts Timestamp) Value() (driver.Value, error) {
	return ptypes.Timestamp(ts.Timestamp)
}
