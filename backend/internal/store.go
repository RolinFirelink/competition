package internal

type Store interface {
	CreateUser(User) (User, error)
	GetUser(streetID, password string) (User, bool)
	ListBanners() []Banner
	ListStudentsPublic() []StudentPublic
	ListStudentsAdmin() []User
	ListRoutesByStudent(studentID string) []GrowthRoute
	CreateRoute(GrowthRoute) (GrowthRoute, error)
	AddReply(routeID string, reply Reply) (Reply, error)
	ListEvents() []Event
	GetEvent(id string) Event
	Penalize(id, reason string, show bool) error
	Stats() Stats
}
