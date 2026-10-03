package domain

type ProjectID string

type Project struct {
	id          ProjectID
	name        string
	description string
}

func NewProject(name, description string) (*Project, error) {
	id, err := newProjectID()
	if err != nil {
		return nil, err
	}

	return &Project{
		id,
		name,
		description,
	}, nil
}

func NewProjectWithID(id ProjectID, name, description string) *Project {
	return &Project{
		id,
		name,
		description,
	}
}

func (p *Project) Name() string { return p.name }

func (p *Project) Description() string { return p.description }

func (p *Project) ID() ProjectID { return p.id }

func (p *Project) Rename(name string) {
	p.name = name
}

func newProjectID() (ProjectID, error) {
	id, err := newID()
	if err != nil {
		return "", nil
	}

	return ProjectID(id), nil
}
