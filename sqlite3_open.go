package sqlite3

// diff to go-sqlite3

import (
	"fmt"
	"strings"
)

type ExecFun func(string) error

type OnOpenHookFun func(dsn string, exec ExecFun) error

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
			// fmt.Println(cmd) // print for debug
			err := exec(cmd)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
