package mock

import "github.com/stretchr/testify/mock" // for testing

type MockShortURLStore struct {
	mock.Mock //  inherits all of mock.Mock's methods
}
