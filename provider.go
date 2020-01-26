package shiiba

// Provider provides activities
type Provider func(*Activities)

func NullProvider(acts *Activities) {
  // do nothing
}

func getProvider(service string) Provider {
	return NullProvider
}
