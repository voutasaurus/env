package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

var parseLog = func(error) {}

// SetParseLog sets up a logger for parse errors on environment variables with
// default values configured. The default value will be returned on a parse
// failure but the provided log function will also be called.
//
// To set the log function, as the first var in your var block put
//
// 	_ = env.SetParseLog(func(err error) {
// 		// custom logging code goes here
// 	})
//
// This way the log function will be initialized before calls to env functions
// that parse, such as Int and Duration.
//
// The default log function is a nop (does not log anything).
func SetParseLog(logFunc func(error)) struct{} {
	parseLog = logFunc
	return struct{}{}
}

// String retrieves the string value of the environment variable named by the
// key. If the variable is present in the environment the value (which may be
// empty) is returned.
// Otherwise the returned value will be the default value specified by the
// value argument.
func String(key, value, description string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return value
}

// MustString retrieves the string value of the environment variable named by
// the key. If the variable is present in the environment the value (which may
// be empty) is returned.
// Otherwise this function logs an error to the default log and causes an
// os.Exit.
func MustString(key, description string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env: %v not set", key)
	}
	return v
}

// Int retrieves the int value of the environment variable named by the key. If
// the variable is present in the environment and the value successfully parses
// into an int then that int is returned.
// Otherwise the returned value will be the default value specified by the
// value argument.
// If the parse fails the logger set in SetParseLog will be called.
func Int(key string, value int, description string) int {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.Atoi(s)
		if err == nil {
			return v
		}
		parseLog(err)
	}
	return value
}

// MustInt retrieves the int value of the environment variable named by the
// key. If the variable is present in the environment and the value
// successfully parses into an int then that int is returned.
// Otherwise this function logs an error to the default log and causes an
// os.Exit.
func MustInt(key, description string) int {
	s, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env: %v not set", key)
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("env: %v parse error: %v", key, err)
	}
	return v
}

// Int64 retrieves the int64 value of the environment variable named by the
// key. If the variable is present in the environment and the value
// successfully parses into an int64 then that int64 is returned.
// Otherwise the returned value will be the default value specified by the
// value argument.
// If the parse fails the logger set in SetParseLog will be called.
func Int64(key string, value int64, description string) int64 {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseInt(s, 10, 0)
		if err == nil {
			return v
		}
		parseLog(err)
	}
	return value
}

// MustInt64 retrieves the int64 value of the environment variable named by the
// key. If the variable is present in the environment and the value
// successfully parses into an int64 then that int64 is returned.
// Otherwise this function logs an error to the default log and causes an
// os.Exit.
func MustInt64(key, description string) int64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env: %v not set", key)
	}
	v, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		log.Fatalf("env: %v parse error: %v", key, err)
	}
	return v
}

// Bool retrieves the bool value of the environment variable named by the key.
// If the variable is not present in the environment then the returned bool is
// false. Also, if the variable is present in the environment and is set to any
// of the strings "0", "false", "False", "f", "F", "n", or "N", then the
// returned bool will be false.
// If the variable is present and set to anything not in the false values
// listed above, then the returned bool will be true.
func Bool(key string, value bool, description string) bool {
	if s, ok := os.LookupEnv(key); ok {
		switch s {
		case "0", "false", "False", "f", "F", "n", "N":
			return false
		}
		return true
	}
	return value
}

// Duration retrieves the time.Duration value of the environment variable named
// by the key. If the variable is present in the environment and the value
// successfully parses into a time.Duration then that time.Duration value is
// returned.
// Otherwise the returned value will be the default value specified by the
// value argument.
// If the parse fails the logger set in SetParseLog will be called.
func Duration(key string, value time.Duration, description string) time.Duration {
	if s, ok := os.LookupEnv(key); ok {
		v, err := time.ParseDuration(s)
		if err == nil {
			return v
		}
		parseLog(err)
	}
	return value
}

// MustDuration retrieves the time.Duration value of the environment variable
// named by the key. If the variable is present in the environment and the
// value successfully parses into an time.Duration then that time.Duration is
// returned.
// Otherwise this function logs an error to the default log and causes an
// os.Exit.
func MustDuration(key, description string) time.Duration {
	s, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env: %v not set", key)
	}
	v, err := time.ParseDuration(s)
	if err != nil {
		log.Fatalf("env: %v parse error: %v", key, err)
	}
	return v
}

// Uint retrieves the uint value of the environment variable named by the key.
// If the variable is present in the environment and the value successfully
// parses into an uint then that uint is returned.
// Otherwise the returned value will be the default value specified by the
// value argument.
// If the parse fails the logger set in SetParseLog will be called.
func Uint(key string, value uint, description string) uint {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseUint(s, 10, 0)
		if err == nil {
			return uint(v)
		}
		parseLog(err)
	}
	return value
}

// MustUint retrieves the uint value of the environment variable named by the
// key. If the variable is present in the environment and the value
// successfully parses into an uint then that uint is returned.
// Otherwise this function logs an error to the default log and causes an
// os.Exit.
func MustUint(key, description string) uint {
	s, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env: %v not set")
	}
	v, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		log.Fatalf("env: %v parse error: %v", key, err)
	}
	return uint(v)
}

// Uint64 retrieves the uint64 value of the environment variable named by the
// key. If the variable is present in the environment and the value
// successfully parses into an uint64 then that uint64 is returned.
// Otherwise the returned value will be the default value specified by the
// value argument.
// If the parse fails the logger set in SetParseLog will be called.
func Uint64(key string, value uint64, description string) uint64 {
	if s, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseUint(s, 10, 0)
		if err == nil {
			return v
		}
		parseLog(err)
	}
	return value
}

// MustUint64 retrieves the uint64 value of the environment variable named by
// the key. If the variable is present in the environment and the value
// successfully parses into an uint64 then that uint64 is returned.
// Otherwise this function logs an error to the default log and causes an
// os.Exit.
func MustUint64(key, description string) uint64 {
	s, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("env: %v not set")
	}
	v, err := strconv.ParseUint(s, 10, 0)
	if err != nil {
		log.Fatalf("env: %v parse error: %v", key, err)
	}
	return v
}
