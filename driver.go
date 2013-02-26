// Copyright 2012 Julien Schmidt. All rights reserved.
// http://www.julienschmidt.com
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
package mysql

import (
	"bufio"
	"database/sql"
	"database/sql/driver"
	"net"
)

type mysqlDriver struct{}

var connCount int

// Open new Connection.
// See https://github.com/Go-SQL-Driver/MySQL#dsn-data-source-name for how
// the DSN string is formated
func (d *mysqlDriver) Open(dsn string) (driver.Conn, error) {
	var err error

	// New mysqlConn
	mc := new(mysqlConn)
	mc.id = connCount
	connCount++
	mc.cfg = parseDSN(dsn)

	// Connect to Server
	mc.netConn, err = net.Dial(mc.cfg.net, mc.cfg.addr)
	if err != nil {
		return nil, err
	}
	mc.bufReader = bufio.NewReader(mc.netConn)

	// Reading Handshake Initialization Packet
	err = mc.readInitPacket()
	if err != nil {
		return nil, err
	}

	// Send Client Authentication Packet
	err = mc.writeAuthPacket()
	if err != nil {
		return nil, err
	}

	// Read Result Packet
	err = mc.readResultOK()
	if err != nil {
		return nil, err
	}

	// Handle DSN Params
	err = mc.handleParams()
	if err != nil {
		return nil, err
	}

	return mc, err
}

func init() {
	connCount = 0
	sql.Register("mysql", &mysqlDriver{})
}
