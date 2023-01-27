// Demonstration of custom logging package.
package main

// Import logging package as log.  Care should be used since log is a standard
// library package.
import log "logging"

// Entry point.  Demonstrates using the different loggers.
// 
// Sample output:
//
// INFO: 2023/01/27 22:19:48.730449 logger.go:16: Sample
// ERROR: 2023/01/27 22:19:48.730498 logger.go:20: Sample
// WARNING: 2023/01/27 22:19:48.730509 logger.go:24: Sample
// VERBOSE: 2023/01/27 22:19:48.730519 logger.go:28: Sample
func main() {
	log.Info("Sample")
	log.Error("Sample")
	log.Warning("Sample")
	log.Verbose("Sample")
}
