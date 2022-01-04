package updater

import "github.com/goncharovnikita/wallpaperize/app/api"

type Unsplash struct {
	api api.UnsplashAPI
}

func NewUnsplash(accessToken string) *Unsplash {
	return &Unsplash{
		api: api.NewUnsplashApi(staticTokenGetter{
			token: accessToken,
		}),
	}
}

func (u *Unsplash) Run() error {
	return nil
}

func (u *Unsplash) Stop() error {
	return nil
}

type staticTokenGetter struct {
	token string
}

func (s *staticTokenGetter) GetToken() (string, error) {
	return s.token, nil
}
