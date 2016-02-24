// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/algorithm"
	"github.com/concourse/atc/scheduler"
	"github.com/pivotal-golang/lager"
)

type FakeJobService struct {
	CanBuildBeScheduledStub        func(logger lager.Logger, build db.Build, buildPrep *db.BuildPreparation, versions *algorithm.VersionsDB) ([]db.BuildInput, bool, string, error)
	canBuildBeScheduledMutex       sync.RWMutex
	canBuildBeScheduledArgsForCall []struct {
		logger    lager.Logger
		build     db.Build
		buildPrep *db.BuildPreparation
		versions  *algorithm.VersionsDB
	}
	canBuildBeScheduledReturns struct {
		result1 []db.BuildInput
		result2 bool
		result3 string
		result4 error
	}
}

func (fake *FakeJobService) CanBuildBeScheduled(logger lager.Logger, build db.Build, buildPrep *db.BuildPreparation, versions *algorithm.VersionsDB) ([]db.BuildInput, bool, string, error) {
	fake.canBuildBeScheduledMutex.Lock()
	fake.canBuildBeScheduledArgsForCall = append(fake.canBuildBeScheduledArgsForCall, struct {
		logger    lager.Logger
		build     db.Build
		buildPrep *db.BuildPreparation
		versions  *algorithm.VersionsDB
	}{logger, build, buildPrep, versions})
	fake.canBuildBeScheduledMutex.Unlock()
	if fake.CanBuildBeScheduledStub != nil {
		return fake.CanBuildBeScheduledStub(logger, build, buildPrep, versions)
	} else {
		return fake.canBuildBeScheduledReturns.result1, fake.canBuildBeScheduledReturns.result2, fake.canBuildBeScheduledReturns.result3, fake.canBuildBeScheduledReturns.result4
	}
}

func (fake *FakeJobService) CanBuildBeScheduledCallCount() int {
	fake.canBuildBeScheduledMutex.RLock()
	defer fake.canBuildBeScheduledMutex.RUnlock()
	return len(fake.canBuildBeScheduledArgsForCall)
}

func (fake *FakeJobService) CanBuildBeScheduledArgsForCall(i int) (lager.Logger, db.Build, *db.BuildPreparation, *algorithm.VersionsDB) {
	fake.canBuildBeScheduledMutex.RLock()
	defer fake.canBuildBeScheduledMutex.RUnlock()
	return fake.canBuildBeScheduledArgsForCall[i].logger, fake.canBuildBeScheduledArgsForCall[i].build, fake.canBuildBeScheduledArgsForCall[i].buildPrep, fake.canBuildBeScheduledArgsForCall[i].versions
}

func (fake *FakeJobService) CanBuildBeScheduledReturns(result1 []db.BuildInput, result2 bool, result3 string, result4 error) {
	fake.CanBuildBeScheduledStub = nil
	fake.canBuildBeScheduledReturns = struct {
		result1 []db.BuildInput
		result2 bool
		result3 string
		result4 error
	}{result1, result2, result3, result4}
}

var _ scheduler.JobService = new(FakeJobService)
