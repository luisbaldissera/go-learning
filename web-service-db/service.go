package main

type Service struct {
    Repo *Repo
}

func (s *Service) ReadAlbums() ([]Album, error) {
    return s.Repo.SelectAlbums()
}

func (s *Service) ReadAlbum(id int64) (Album, error) {
    return s.Repo.SelectAlbumByID(id)
}

func (s *Service) Artists() ([]string, error) {
    return s.Repo.SelectArtists();
}

func (s *Service) AddAlbum(album Album) (Album, error) {
    id, err := s.Repo.InsertAlbum(album)
    if err != nil {
        return album, err
    }
    album.ID = id
    return album, nil
}

func (s *Service) DeleteAlbum(id int64) (Album, error) {
    album, err := s.Repo.DeleteAlbum(id)
    if err != nil {
        return album, err
    }
    return album, nil
}

// vim: fdl=0
