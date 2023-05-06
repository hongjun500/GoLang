package module

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/GoLang-master/database-example/config"
	"log"
	"time"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  string
	config.DbConn
}

/*
根据主键Id查询一条数据
*/

func AlbumByID(id int64, album *Album) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := album.Db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

/*
根据 artist 查询多条数据
*/

func AlbumsByArtist(name string, album *Album) ([]Album, error) {
	var albums []Album

	rows, err := album.Db.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumsByArtist %q, %v", name, err)
	}
	defer rows.Close()
	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumsByArtist %q, %v", name, err)
		}

		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}
	return albums, nil
}

/*
查询多行数据
*/

func AlbumAll(album *Album) ([]Album, error) {
	var albumList []Album
	// 返回多行
	rows, err := album.Db.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("albumAll fail%v", err)
	}
	defer rows.Close()

	// 对多行结果循环
	for rows.Next() {
		var alb Album
		var t sql.NullString
		// 赋值
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist,
			&alb.Price); err != nil {
			return albumList, err
		}
		// 空字符串处理
		if t.Valid {
			alb.Title = t.String
		}
		albumList = append(albumList, alb)
	}
	return albumList, nil
}

/*
添加数据
*/
func SaveAlbum(album *Album) (Album, error) {
	var alb Album
	result, err := album.Db.Exec("INSERT INTO album (title, artist, price) values (?, ?, ?)", album.Title, album.Artist, album.Price)
	if err != nil {
		return alb, fmt.Errorf("insert album fail %v", err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return alb, fmt.Errorf("lastInsertId fail%v", err)
	}
	album.ID = lastInsertId
	fmt.Printf("insert album success ID=【%v】\n", lastInsertId)
	return *album, nil
}

func listResult(album *Album) (any, error) {
	// todo 这里有语法错误，官网写法如此
	query, err := album.Db.Query("SELECT * from album; SELECT * from song;")
	if err != nil {
		return nil, fmt.Errorf("query fail%v", err)
	}
	defer album.Db.Close()
	resultSet := query.NextResultSet()
	fmt.Println("resultSet = ", resultSet)
	return nil, err
}

/*
预处理语句
*/

func PreAlbumById(id int64, album *Album) (Album, error) {
	stmt, err := album.Db.Prepare("select * from album where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	var alb Album
	err = stmt.QueryRow(id).Scan(&alb.ID, &alb.Artist, &alb.Title, &alb.Price)
	if err != nil {
		if err == sql.ErrNoRows {

		}
		return alb, nil
	}
	return alb, nil
}

/*
事务操作
*/

func CreateTag(ctx context.Context, album *Album, albumId int, title string) (tagId int64, err error) {
	// 定义匿名函数
	fail := func(err error) (int64, error) {
		return 0, err
	}
	// 开启事务
	tx, err := album.Db.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	// 回滚
	defer tx.Rollback()
	var enough bool
	// 查询
	if tx.QueryRowContext(ctx, "select (title like ?) from album where id = ?", title, albumId).Scan(&enough); err != nil {
		if err == sql.ErrNoRows {
			return fail(fmt.Errorf("no such album"))
		}
		return fail(err)
	}
	if !enough {
		return fail(fmt.Errorf("not enough inventory"))
	}
	// 修改
	_, err = tx.ExecContext(ctx, "update album set title = ? where id = ?", title+title, albumId)
	if err != nil {
		return fail(err)
	}
	// 添加
	result, err := tx.ExecContext(ctx, "insert into tag (name, created_by, state, created_on) values (?, ?, ?, ?)", "album", "create_album", 0, time.Now())
	if err != nil {
		return fail(err)
	}
	tagId, err = result.LastInsertId()
	if err != nil {
		return fail(err)
	}
	// 提交事务
	if err = tx.Commit(); err != nil {
		return fail(err)
	}
	return tagId, nil
}
