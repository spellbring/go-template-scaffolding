package rest

type config struct {
	timeout int
}

func restClientConfig() *config {
	return &config{
		timeout: 30,
	}
}
