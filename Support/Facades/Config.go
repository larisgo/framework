package Facades

import (
	RepositoryContract "github.com/larisgo/framework/Contracts/Config"
)

var Config func() RepositoryContract.Repository = func() RepositoryContract.Repository {
	return NewFacade("config").Get().(RepositoryContract.Repository)
}
