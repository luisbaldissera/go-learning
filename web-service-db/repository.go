package main

import (
	"database/sql"
)

type Repo struct {
    DB  *sql.DB
}

func NewRepo(db *sql.DB) Repo {
    return Repo{db}
}

func (r *Repo) SelectAlbums() ([]Album, error) {
    albums := []Album{}
    rows, err := r.DB.Query("SELECT id, title, artist, price FROM album")
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var a Album
        if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
            return nil, err
        }
        albums = append(albums, a)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return albums, nil
}

func (r *Repo) SelectAlbumByID(id int64) (Album, error) {
    a := Album{}
    row := r.DB.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id)
    if err := row.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
        return a, err
    }
    return a, nil
}

func (r *Repo) SelectAlbumsByArtist(artist string) ([]Album, error) {
    albums := []Album{}
    rows, err := r.DB.Query("SELECT id, title, artist, price FROM album WHERE artist = ?", artist)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var a Album
        if err := rows.Scan(&a.ID, &a.Title, &a.Artist, &a.Price); err != nil {
            return nil, err
        }
        albums = append(albums, a)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return albums, nil
    
}

func (r *Repo) SelectArtists() ([]string, error) {
    artists := []string{}
    rows, err := r.DB.Query("SELECT DISTINCT artist FROM albums")
    if err != nil {
        return nil, err
    }
    for rows.Next() {
        var artist string
        if err := rows.Scan(&artist); err != nil {
            return nil, err
        }
        artists = append(artists, artist)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return artists, nil
}

func (r *Repo) InsertAlbum(album Album) (int64, error) {
    result, err := r.DB.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", album.Title, album.Artist, album.Price)
    if err != nil {
        return 0, err
    }
    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }
    return id, nil
}

func (r *Repo) DeleteAlbum(id int64) (Album, error) {
    album := Album{}
    row := r.DB.QueryRow("SELECT id, title, artist, price FROM album WHERE id = ?", id)
    if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
        return album, err
    }
    _, err := r.DB.Exec("DELETE FROM album WHERE id = ?", id)
    if err != nil {
        return album, err
    }
    return album, nil
}
// vim: fdl=0
