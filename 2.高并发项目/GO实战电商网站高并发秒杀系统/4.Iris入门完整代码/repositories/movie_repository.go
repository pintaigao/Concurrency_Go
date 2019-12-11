package repositories

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager() MovieRepository  {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	//数据操作
	return "慕课网视频"
}