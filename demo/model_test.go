package main

import (
	"os"
	"testing"

	"github.com/gocraft/dbr"
	"github.com/mh-cbon/jedi/builder"
	"github.com/mh-cbon/jedi/drivers"
	jedi "github.com/mh-cbon/jedi/runtime"
)

func getConnFromEnv(t *testing.T, forceSetup bool) *dbr.Connection {
	dsn := os.Getenv("JDSN")
	driver := os.Getenv("JDRIVER")
	if driver == "" {
		driver = "sqlite3"
	}
	if dsn == "" {
		dsn = "db.db"
	}
	t.Logf("Driver:%q dsn:%v", driver, dsn)
	conn, err := dbr.Open(driver, dsn, nil)
	if err != nil {
		t.Fatalf("Connection setup failed: %v", err)
	}
	if err := jedi.Setup(conn, forceSetup); err != nil {
		t.Fatal(err)
	}
	return conn
}

func TestModel(t *testing.T) {

	conn := getConnFromEnv(t, false)
	defer conn.Close()
	defer func() {
		if jedi.Runs(drivers.Sqlite) {
			os.Remove(os.Getenv("JDSN"))
		}
	}()

	t.Run("can create the schema", func(t *testing.T) {
		if err := JBasicTypesSetup().Create(conn); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("model properties are set", func(t *testing.T) {
		props := JBasicTypesModel.Properties()
		pName := "ID"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "int64", "id", "INTEGER", true, true)
		}
		pName = "String"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "string", "string", "TEXT", false, false)
		}
		pName = "StringP"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*string", "string_p", "TEXT", false, false)
		}
		pName = "Int"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "int", "int", "INTEGER", false, false)
		}
		pName = "Int32"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "int32", "int32", "INTEGER", false, false)
		}
		pName = "Int64"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "int64", "int64", "INTEGER", false, false)
		}
		pName = "UInt"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "uint", "u_int", "INTEGER", false, false)
		}
		pName = "UInt32"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "uint32", "u_int32", "INTEGER", false, false)
		}
		pName = "UInt64"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "uint64", "u_int64", "INTEGER", false, false)
		}
		pName = "IntP"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*int", "int_p", "INTEGER", false, false)
		}
		pName = "Int32P"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*int32", "int32_p", "INTEGER", false, false)
		}
		pName = "Int64P"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*int64", "int64_p", "INTEGER", false, false)
		}
		pName = "UIntP"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*uint", "u_int_p", "INTEGER", false, false)
		}
		pName = "UInt32P"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*uint32", "u_int32_p", "INTEGER", false, false)
		}
		pName = "UInt64P"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*uint64", "u_int64_p", "INTEGER", false, false)
		}
		pName = "Bool"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "bool", "bool", "INTEGER", false, false)
		}
		pName = "BoolP"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*bool", "bool_p", "INTEGER", false, false)
		}
		pName = "Float32"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "float32", "float32", "FLOAT", false, false)
		}
		pName = "Float32P"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*float32", "float32_p", "FLOAT", false, false)
		}
		pName = "Float64"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "float64", "float64", "FLOAT", false, false)
		}
		pName = "Float64P"
		if f, ok := props[pName]; ok == false {
			t.Fatalf("property %v must exist", pName)
		} else {
			checkProperty(t, f, pName, "*float64", "float64_p", "FLOAT", false, false)
		}
	})

	t.Run("can drop the schema", func(t *testing.T) {
		if err := JBasicTypesSetup().Drop(conn); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("can alias", func(t *testing.T) {
		original := JBasicTypesModel
		aliased := original.As("r")
		if original.Alias() != original.Table() {
			t.Fatal("Invalid alias wanted=", original.Table(), " got", original.Alias())
		}
		if JBasicTypesModel.Alias() != JBasicTypesModel.Table() {
			t.Fatal("Invalid alias wanted=", JBasicTypesModel.Table(), " got", JBasicTypesModel.Alias())
		}
		if aliased.Alias() != "r" {
			t.Fatal("Invalid alias wanted='r' got", aliased.Alias())
		}
	})
}

func checkProperty(t *testing.T, f builder.MetaProvider, goName, goType, sqlName, sqlType string, isPk, isAi bool) {
	if isPk && !f.IsPk() {
		t.Fatalf("property %v must be PK", f.Name())
	}
	if !isPk && f.IsPk() {
		t.Fatalf("property %v must not be PK", f.Name())
	}
	if isAi && !f.IsAI() {
		t.Fatalf("property %v must be AUTOINREMENT", f.Name())
	}
	if !isAi && f.IsAI() {
		t.Fatalf("property %v must not be AUTOINREMENT", f.Name())
	}
	if f.Name() != sqlName {
		t.Fatalf("property sql name must eq %v got=%v ", sqlName, f.Name())
	}
	if f.Type() != sqlType {
		t.Fatalf("property sql type must eq %v got=%v ", sqlType, f.Type())
	}
	if f.GoName() != goName {
		t.Fatalf("property go name must eq %v got=%v ", goName, f.GoName())
	}
	if f.GoType() != goType {
		t.Fatalf("property go name must eq %v got=%v ", goType, f.GoType())
	}
}
