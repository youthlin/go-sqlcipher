## go-sqlcipher

### Description 简介
Fork from https://github.com/mutecomm/go-sqlcipher

本项目基于以下仓库：
- Go sqlite3 driver: https://github.com/mattn/go-sqlite3
- SQLite extension with AES-256 codec: https://github.com/sqlcipher/sqlcipher
- AES-256 implementation from: https://github.com/libtom/libtomcrypt

### Installation 安装

    go get github.com/youthlin/go-sqlcipher


### Documentation 文档

示例代码
```go
package main

import (
	"database/sql"
	"fmt"
	"strings"

	sqlite3 "github.com/youthlin/go-sqlcipher"
)

func main() {
	var (
		db   *sql.DB
		dsn = "EnMicroMsg.db"
		key  = "xxxXxxx"// 密码
	)
	encrypted, err := sqlite3.IsEncrypted(dsn)
	if err != nil {
		panic(err) // 文件不存在or空文件
	}
	if encrypted { // 已加密
		params := []string{
			fmt.Sprintf(`_pragma_key=%q`, key),
			// 注意顺序 _pragma_key 最先；注意引号
			`_pragma_cipher_page_size=1024`,
			`_pragma_kdf_iter=4000`,
			`_pragma_cipher_hmac_algorithm="HMAC_SHA1"`,
			`_pragma_cipher_kdf_algorithm="PBKDF2_HMAC_SHA1"`,
			`_pragma_cipher_use_hmac="OFF"`,
		}
		// 带参数打开加密数据库
		dsn = fmt.Sprintf("%s?%s", dsn, strings.Join(params, "&"))
	}
	db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("SELECT count(*) FROM sqlite_master where type='table' ")
	var count int
	err = row.Scan(&count)
	if err != nil {
		panic(err) // may invalid key 可能是密码错误
	}
	fmt.Printf("table count = %d\n", count)
}

```

- See also [PRAGMA key](https://www.zetetic.net/sqlcipher/sqlcipher-api/#PRAGMA_key)
- 文档 https://pkg.go.dev/github.com/youthlin/go-sqlcipher
- 可以使用 `sqlite3.IsEncrypted()` 判断一个数据库文件是否已加密。
- `_example` 目录下有更多示例用法。

### Maintenance 从上游更新

参见 [Maintenance](MAINTENANCE) 文件.

### License

参见上游项目各自的开源协议。

The code of the originating packages is covered by their respective licenses.
See [LICENSE](LICENSE) file for details.
