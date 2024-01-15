package config

type Jaeger struct {
	// Log
	Log bool `yaml:"log"`
	// Agent Host Port
	AgentHostPort string `yaml:"agentHostPort"`
	// Collector Host Port
	CollectorHostPort string `yaml:"collectorHostPort"`
}
