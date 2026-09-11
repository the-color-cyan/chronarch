package domain

type ProjectID string

type Project struct {
	name        string
	description string
	id          ProjectID
}

func NewProject(name, description string) (*Project, error) {
	id, err := newProjectID()
	if err != nil {
		return nil, err
	}

	return &Project{
		name,
		description,
		id,
	}, nil
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
