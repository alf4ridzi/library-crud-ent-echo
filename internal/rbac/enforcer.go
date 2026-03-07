package rbac

import (
	"sync"

	"github.com/casbin/casbin/v3"
)

var (
	enforcer *casbin.Enforcer
	once     sync.Once
)

func NewEnforcer(modelPath, policyPath string) (*casbin.Enforcer, error) {
	var err error
	once.Do(func() {
		enforcer, err = casbin.NewEnforcer(modelPath, policyPath)
	})

	return enforcer, err
}

func GetEnforcer() *casbin.Enforcer {
	return enforcer
}
