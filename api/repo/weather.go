package repo

type PixabayRepo interface {
	GetPhoto(city string) (string, error)
	GetVideo(location string) (string, error)
}
