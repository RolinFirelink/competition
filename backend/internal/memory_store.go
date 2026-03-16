package internal

import (
	"errors"
	"strconv"
	"time"
)

type MemoryStore struct {
	users   map[string]User
	routes  map[string][]GrowthRoute
	events  []Event
	banners []Banner
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		users: map[string]User{
			"sf6demo": {StreetID: "sf6demo", Nick: "演示学员", Tier: "钻石", Role: "隆", PlanSlots: 1, PlanUsed: 0},
		},
		routes:  map[string][]GrowthRoute{},
		events:  []Event{{ID: "1", Title: "周末赛", Announce: "报名开启", Participants: []Participant{{StreetID: "sf6demo", Nick: "演示学员", Tier: "钻石"}}}},
		banners: []Banner{{ID: "1", Title: "教学横幅", Link: "https://example.com", Image: "/banner1.png"}},
	}
}

func (m *MemoryStore) CreateUser(u User) (User, error) {
	if _, ok := m.users[u.StreetID]; ok {
		return User{}, errors.New("exists")
	}
	m.users[u.StreetID] = u
	return u, nil
}

func (m *MemoryStore) GetUser(streetID, password string) (User, bool) {
	u, ok := m.users[streetID]
	if !ok {
		return User{}, false
	}
	// demo: no password check for now
	return u, true
}

func (m *MemoryStore) ListBanners() []Banner { return m.banners }

func (m *MemoryStore) ListStudentsPublic() []StudentPublic {
	out := []StudentPublic{}
	for _, u := range m.users {
		out = append(out, StudentPublic{
			StreetID: u.StreetID, Nick: u.Nick, Tier: u.Tier, Role: u.Role, Level: u.Level,
		})
	}
	return out
}

func (m *MemoryStore) ListStudentsAdmin() []User {
	out := []User{}
	for _, u := range m.users {
		out = append(out, u)
	}
	return out
}

func (m *MemoryStore) ListRoutesByStudent(studentID string) []GrowthRoute {
	return m.routes[studentID]
}

func (m *MemoryStore) CreateRoute(g GrowthRoute) (GrowthRoute, error) {
	g.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	m.routes[g.StudentID] = append(m.routes[g.StudentID], g)
	return g, nil
}

func (m *MemoryStore) AddReply(routeID string, reply Reply) (Reply, error) {
	reply.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	reply.CreatedAt = time.Now()
	for sid, routes := range m.routes {
		for i := range routes {
			if routes[i].ID == routeID {
				m.routes[sid][i].Replies = append(m.routes[sid][i].Replies, reply)
				return reply, nil
			}
		}
	}
	return Reply{}, errors.New("route not found")
}

func (m *MemoryStore) ListEvents() []Event { return m.events }

func (m *MemoryStore) GetEvent(id string) Event {
	for _, e := range m.events {
		if e.ID == id {
			return e
		}
	}
	return Event{}
}

func (m *MemoryStore) Penalize(id, reason string, show bool) error {
	u, ok := m.users[id]
	if !ok {
		return errors.New("not found")
	}
	u.Penalized = show
	u.PenaltyReason = reason
	m.users[id] = u
	return nil
}

func (m *MemoryStore) Stats() Stats {
	return Stats{Students: len(m.users), Routes: len(m.routes), Events: len(m.events), Banners: len(m.banners)}
}
