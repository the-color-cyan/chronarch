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

func newProjectID() (ProjectID, error) {
	id, err := newID()
	if err != nil {
		return "", nil
	}

	return ProjectID(id), nil
}
