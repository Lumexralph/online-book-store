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

func (ts *Timestamp) Scan(value interface{}) error {
	switch t := value.(type) {
	case time.Time:
		var err error
		ts.Timestamp, err = ptypes.TimestampProto(t)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("Not a protobuf Timestamp")
	}
	return nil
}

func (ts Timestamp) Value() (driver.Value, error) {
	return ptypes.Timestamp(ts.Timestamp)
}
