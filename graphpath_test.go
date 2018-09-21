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

package ag

import "testing"

func TestBasicPathScanNil(t *testing.T) {
	var p BasicPath
	err := p.Scan(nil)
	if err != nil {
		t.Error(err)
	} else if p.Valid {
		t.Errorf("got %s, want NULL", p)
	}
}

func TestBasicPathScanType(t *testing.T) {
	src := 0
	var p BasicPath
	err := p.Scan(src)
	if err == nil {
		t.Errorf("error expected for %T", src)
	}
}

func TestBasicPathScanZero(t *testing.T) {
	var src interface{} = []byte(nil)
	var p BasicPath
	err := p.Scan(src)
	if err == nil {
		t.Errorf("error expected for %v", src)
	}
}

func TestBasicPathScan(t *testing.T) {
	tests := []struct {
		b  []byte
		nv int
		ne int
	}{
		{[]byte("[]"), 0, 0},
		{[]byte(`[v[3.1]{},e[4.1][3.1,3.2]{},v[3.2]{},NULL,NULL]`), 3, 2},
	}
	for _, c := range tests {
		var p BasicPath
		err := p.Scan(c.b)
		if err != nil {
			t.Error(err)
		} else if !p.Valid {
			t.Errorf("got NULL, want Valid %T", p)
		} else {
			if nv := len(p.Vertices); nv != c.nv {
				t.Errorf("got len(p.Vertices) == %d, want %d", nv, c.nv)
			} else if ne := len(p.Edges); ne != c.ne {
				t.Errorf("got len(p.Edges) == %d, want %d", ne, c.ne)
			}
		}
	}
}

func TestServerGraphpath(t *testing.T) {
	skipUnlessServerTest(t)

	db := mustOpenAndSetGraph(t)
	defer db.Close()

	_, err := db.Exec(`CREATE (:pv)-[:pe]->(:pv)-[:pe]->(:pv)`)
	if err != nil {
		t.Fatal(err)
	}

	var p BasicPath
	q := `MATCH p=(:pv)-[:pe]->(:pv)-[:pe]->(:pv) RETURN p LIMIT 1`
	err = db.QueryRow(q).Scan(&p)
	if err != nil {
		t.Error(err)
	} else if !p.Valid {
		t.Errorf("got NULL, want Valid %T", p)
	} else {
		if nv := len(p.Vertices); nv != 3 {
			t.Errorf("got len(p.Vertices) == %d, want %d", nv, 3)
		} else if ne := len(p.Edges); ne != 2 {
			t.Errorf("got len(p.Edges) == %d, want %d", ne, 2)
		}
	}
}
