package campain

type Repository interface {
	Save(campain *Campain) error
}
