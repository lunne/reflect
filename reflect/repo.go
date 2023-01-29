package reflect

type Repository interface {
	AddGoal(goal *Goal) error
	GetGoal(id string) (*Goal, error)
}
