// Go MySQL Driver - A MySQL-Driver for Go's database/sql package
//
// Copyright 2012 Julien Schmidt. All rights reserved.
// http://www.julienschmidt.com
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

package mysql

const (
	MIN_PROTOCOL_VERSION = 10
	MAX_PACKET_SIZE      = 1<<24 - 1
	TIME_FORMAT          = "2006-01-02 15:04:05"
)

// MySQL constants documentation:
// http://dev.mysql.com/doc/internals/en/client-server-protocol.html

type ClientFlag uint32

const (
	CLIENT_LONG_PASSWORD ClientFlag = 1 << iota
	CLIENT_FOUND_ROWS
	CLIENT_LONG_FLAG
	CLIENT_CONNECT_WITH_DB
	CLIENT_NO_SCHEMA
	CLIENT_COMPRESS
	CLIENT_ODBC
	CLIENT_LOCAL_FILES
	CLIENT_IGNORE_SPACE
	CLIENT_PROTOCOL_41
	CLIENT_INTERACTIVE
	CLIENT_SSL
	CLIENT_IGNORE_SIGPIPE
	CLIENT_TRANSACTIONS
	CLIENT_RESERVED
	CLIENT_SECURE_CONN
	CLIENT_MULTI_STATEMENTS
	CLIENT_MULTI_RESULTS
)

type commandType byte

const (
	COM_QUIT commandType = iota + 1
	COM_INIT_DB
	COM_QUERY
	COM_FIELD_LIST
	COM_CREATE_DB
	COM_DROP_DB
	COM_REFRESH
	COM_SHUTDOWN
	COM_STATISTICS
	COM_PROCESS_INFO
	COM_CONNECT
	COM_PROCESS_KILL
	COM_DEBUG
	COM_PING
	COM_TIME
	COM_DELAYED_INSERT
	COM_CHANGE_USER
	COM_BINLOG_DUMP
	COM_TABLE_DUMP
	COM_CONNECT_OUT
	COM_REGISTER_SLAVE
	COM_STMT_PREPARE
	COM_STMT_EXECUTE
	COM_STMT_SEND_LONG_DATA
	COM_STMT_CLOSE
	COM_STMT_RESET
	COM_SET_OPTION
	COM_STMT_FETCH
)

type FieldType byte

const (
	FIELD_TYPE_DECIMAL FieldType = iota
	FIELD_TYPE_TINY
	FIELD_TYPE_SHORT
	FIELD_TYPE_LONG
	FIELD_TYPE_FLOAT
	FIELD_TYPE_DOUBLE
	FIELD_TYPE_NULL
	FIELD_TYPE_TIMESTAMP
	FIELD_TYPE_LONGLONG
	FIELD_TYPE_INT24
	FIELD_TYPE_DATE
	FIELD_TYPE_TIME
	FIELD_TYPE_DATETIME
	FIELD_TYPE_YEAR
	FIELD_TYPE_NEWDATE
	FIELD_TYPE_VARCHAR
	FIELD_TYPE_BIT
)
const (
	FIELD_TYPE_NEWDECIMAL FieldType = iota + 0xf6
	FIELD_TYPE_ENUM
	FIELD_TYPE_SET
	FIELD_TYPE_TINY_BLOB
	FIELD_TYPE_MEDIUM_BLOB
	FIELD_TYPE_LONG_BLOB
	FIELD_TYPE_BLOB
	FIELD_TYPE_VAR_STRING
	FIELD_TYPE_STRING
	FIELD_TYPE_GEOMETRY
)

type FieldFlag uint16

const (
	FLAG_NOT_NULL FieldFlag = 1 << iota
	FLAG_PRI_KEY
	FLAG_UNIQUE_KEY
	FLAG_MULTIPLE_KEY
	FLAG_BLOB
	FLAG_UNSIGNED
	FLAG_ZEROFILL
	FLAG_BINARY
	FLAG_ENUM
	FLAG_AUTO_INCREMENT
	FLAG_TIMESTAMP
	FLAG_SET
	FLAG_UNKNOWN_1
	FLAG_UNKNOWN_2
	FLAG_UNKNOWN_3
	FLAG_UNKNOWN_4
)
