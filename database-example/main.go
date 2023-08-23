package main

import (
	"context"
	"fmt"
	"log"

	"github.com/hongjun500/GoLang-master/database-example/config"
	"github.com/hongjun500/GoLang-master/database-example/module"
)

/*// 声明一个数据库句柄
var db *sql.DB
var Txh *sql.Tx*/

func main() {

	dbconn, _ := config.GetMySQLConn()

	alb := &module.Album{
		DbConn: dbconn,
	}

	albums, err := module.AlbumsByArtist("John Coltrane", alb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Albums found: %v\n", albums)

	resultAlb, err := module.AlbumByID(2, alb)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Album found: %v\n", resultAlb)

	newAlbum := &module.Album{
		Title:  "hongjun500",
		Artist: "a",
		Price:  "1",
		DbConn: dbconn,
	}
	album, err := module.SaveAlbum(newAlbum)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album insert ok%v", album)

	all, _ := module.AlbumAll(alb)
	fmt.Println(all)

	/*list_result, _ := listResult()
	fmt.Println(list_result)*/

	preAlbum, _ := module.PreAlbumById(2, alb)
	log.Println(preAlbum)

	ctx := context.Background()
	// 执行事务的函数
	tagId, err := module.CreateTag(ctx, alb, 1, "Blue Train")
	log.Println("CreateTag -> tagId = ", tagId)
}
