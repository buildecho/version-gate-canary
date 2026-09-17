// Package canary is a canary artifact for validating registry policy
// enforcement. Not for production use.
package canary

const (
	Name    = "version-gate-canary"
	Version = "v1.0.0"
	Message = "If you can install this version, the policy is NOT blocking it."
)
