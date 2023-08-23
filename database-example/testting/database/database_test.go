package testting

import (
	"context"
	"testing"
	"time"

	"github.com/hongjun500/GoLang-master/database-example/config"
	"github.com/hongjun500/GoLang-master/database-example/module"
)

func TestDatabase(t *testing.T) {
	t.Log("TestDatabase,start")
}

// 测试事务方法
func TestCreateTagTX(t *testing.T) {
	t.Logf("现在时间【%v】", time.Now())
	// 获取数据库连接
	dbConn, _ := config.GetMySQLConn()
	alb := &module.Album{
		DbConn: dbConn,
	}
	ctx := context.Background()
	tagId, err := module.CreateTag(ctx, alb, 3, "Jeru")
	if err != nil {
		t.Fatalf("事务方法创建错误【%v】", err)
	}
	t.Logf("返回的tagId【%v】", tagId)
}

// 测试通过 id 查询
func TestAlbumById(t *testing.T) {
	dbConn, _ := config.GetMySQLConn()
	alb := &module.Album{DbConn: dbConn}
	album, _ := module.AlbumByID(1, alb)
	t.Logf("AlbumId = 1 ->【%v】", album)
}
