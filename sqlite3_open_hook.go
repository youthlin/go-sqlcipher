package sqlite3

// diff to go-sqlite3

import (
	"fmt"
	"strings"
)

// ExecFun is a wrap of sqlite3_exec.
type ExecFun func(string) error

// OnOpenHookFun is the type of the SQLiteDriver's OnOpenHook field.
// The function will be invoked once open a database.
type OnOpenHookFun func(dsn string, exec ExecFun) error

// SimpleOpenHook is a simple OnOpenHook,
// which accept any `_pragma_xxx=yyy` params,
// and pass the `PRAGMA xxx=yyy` to `sqlite3_exec` once open.
func SimpleOpenHook(dsn string, exec ExecFun) error {
	pos := strings.IndexRune(dsn, '?')
	if pos < 1 {
		return nil
	}
	kvStrList := strings.Split(dsn[pos+1:], "&")
	for _, kvStr := range kvStrList {
		kvPair := strings.Split(kvStr, "=")
		if len(kvPair) != 2 {
			return fmt.Errorf("invalid params")
		}
		key := kvPair[0]
		value := kvPair[1]
		if strings.HasPrefix(key, "_pragma_") {
			key = strings.TrimPrefix(key, "_pragma_")
			cmd := fmt.Sprintf("PRAGMA %s = %s;", key, value)
			err := exec(cmd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
