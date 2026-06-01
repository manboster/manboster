package skill

import "sync"

var lock sync.RWMutex

var registryMap = make(map[string]*Skill)

func Register(name string, skill *Skill) {
	lock.Lock()
	defer lock.Unlock()
	_, avail := registryMap[name]
	if avail {
		panic("[Manboster Skills] skill already registered!!! Name: " + name)
	}
	registryMap[name] = skill
}

func List() []*Skill {
	lock.RLock()
	defer lock.RUnlock()
	
	var skills []*Skill
	for _, skill := range registryMap {
		skills = append(skills, skill)
	}

	return skills
}

func Get(name string) (*Skill, bool) {
	lock.RLock()
	defer lock.RUnlock()

	s, ok := registryMap[name]
	return s, ok
}
