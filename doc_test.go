/*
 * Copyright (c) 2014-2018 Bitnine, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package ag_test

import (
	"database/sql"

	"github.com/bitnine-oss/agensgraph-golang"
)

var db *sql.DB

func ExampleGraphId_Scan() {
	var gid ag.GraphId
	err := db.QueryRow(`MATCH (n) RETURN id(n) LIMIT 1`).Scan(&gid)
	if err == nil && gid.Valid {
		// Valid graphid
	} else {
		// An error occurred or the graphid is NULL
	}
}

func ExampleGraphId_Value() {
	gid, _ := ag.NewGraphId("1.1")
	db.QueryRow(`MATCH (n) WHERE id(n) = $1 RETURN n`, gid)
}

func ExampleArray_graphId() {
	var gids []ag.GraphId

	// Scan()
	db.QueryRow(`MATCH ()-[es*]->() RETURN [e IN es | id(e)] LIMIT 1`).Scan(ag.Array(&gids))

	// Value()
	gid0, _ := ag.NewGraphId("1.1")
	gid1, _ := ag.NewGraphId("65535.281474976710655")
	gids = []ag.GraphId{gid0, gid1}
	db.QueryRow(`MATCH (n) WHERE id(n) IN $1 RETURN n LIMIT 1`, ag.Array(gids))
}

func ExampleBasicVertex_Scan() {
	var v ag.BasicVertex
	err := db.QueryRow(`MATCH (n) RETURN n LIMIT 1`).Scan(&v)
	if err == nil && v.Valid {
		// Valid vertex
	} else {
		// An error occurred or the vertex is NULL
	}
}

func ExampleArray_basicVertex() {
	var vs []ag.BasicVertex
	err := db.QueryRow(`MATCH p=(:v)-[:e]->(:v) RETURN nodes(p) LIMIT 1`).Scan(ag.Array(&vs))
	if err == nil && vs != nil {
		// Valid _vertex
	} else {
		// An error occurred or the _vertex is NULL
	}
}

func ExampleBasicEdge_Scan() {
	var e ag.BasicEdge
	err := db.QueryRow(`MATCH ()-[e]->() RETURN e LIMIT 1`).Scan(&e)
	if err == nil && e.Valid {
		// Valid edge
	} else {
		// An error occurred or the edge is NULL
	}
}

func ExampleArray_basicEdge() {
	var es []ag.BasicEdge
	err := db.QueryRow(`MATCH p=(:v)-[:e]->(:v)-[:e]->(:v) RETURN relationships(p) LIMIT 1`).Scan(ag.Array(&es))
	if err == nil && es != nil {
		// Valid _edge
	} else {
		// An error occurred or the _edge is NULL
	}
}

func ExampleBasicPath_Scan() {
	var p ag.BasicPath
	err := db.QueryRow(`MATCH p=(:v)-[:e]->(:v)-[:e]->(:v) RETURN p LIMIT 1`).Scan(&p)
	if err == nil && p.Valid {
		// Valid graphpath
	} else {
		// An error occurred or the graphpath is NULL
	}
}
