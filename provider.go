package shiiba

// ActivityProvider represents service api what provides user activies
type ActivityProvider interface {
	FillData(*Activities) error
}

type doNothing struct {
}

func (p *doNothing) FillData(acts *Activities) error {
	// Do nothing
	return nil
}

func getProvider(name string) (ActivityProvider, error) {
	return &doNothing{}, nil
}
