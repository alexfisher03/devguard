// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	core "github.com/l3montree-dev/devguard/internal/core"
	gorm "gorm.io/gorm"

	mock "github.com/stretchr/testify/mock"

	models "github.com/l3montree-dev/devguard/internal/database/models"

	uuid "github.com/google/uuid"
)

// FirstPartyVulnRepository is an autogenerated mock type for the FirstPartyVulnRepository type
type FirstPartyVulnRepository struct {
	mock.Mock
}

type FirstPartyVulnRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *FirstPartyVulnRepository) EXPECT() *FirstPartyVulnRepository_Expecter {
	return &FirstPartyVulnRepository_Expecter{mock: &_m.Mock}
}

// Activate provides a mock function with given fields: tx, id
func (_m *FirstPartyVulnRepository) Activate(tx *gorm.DB, id string) error {
	ret := _m.Called(tx, id)

	if len(ret) == 0 {
		panic("no return value specified for Activate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_Activate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Activate'
type FirstPartyVulnRepository_Activate_Call struct {
	*mock.Call
}

// Activate is a helper method to define mock.On call
//   - tx *gorm.DB
//   - id string
func (_e *FirstPartyVulnRepository_Expecter) Activate(tx interface{}, id interface{}) *FirstPartyVulnRepository_Activate_Call {
	return &FirstPartyVulnRepository_Activate_Call{Call: _e.mock.On("Activate", tx, id)}
}

func (_c *FirstPartyVulnRepository_Activate_Call) Run(run func(tx *gorm.DB, id string)) *FirstPartyVulnRepository_Activate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Activate_Call) Return(_a0 error) *FirstPartyVulnRepository_Activate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_Activate_Call) RunAndReturn(run func(*gorm.DB, string) error) *FirstPartyVulnRepository_Activate_Call {
	_c.Call.Return(run)
	return _c
}

// All provides a mock function with no fields
func (_m *FirstPartyVulnRepository) All() ([]models.FirstPartyVuln, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for All")
	}

	var r0 []models.FirstPartyVuln
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.FirstPartyVuln, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.FirstPartyVuln); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.FirstPartyVuln)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_All_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'All'
type FirstPartyVulnRepository_All_Call struct {
	*mock.Call
}

// All is a helper method to define mock.On call
func (_e *FirstPartyVulnRepository_Expecter) All() *FirstPartyVulnRepository_All_Call {
	return &FirstPartyVulnRepository_All_Call{Call: _e.mock.On("All")}
}

func (_c *FirstPartyVulnRepository_All_Call) Run(run func()) *FirstPartyVulnRepository_All_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FirstPartyVulnRepository_All_Call) Return(_a0 []models.FirstPartyVuln, _a1 error) *FirstPartyVulnRepository_All_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_All_Call) RunAndReturn(run func() ([]models.FirstPartyVuln, error)) *FirstPartyVulnRepository_All_Call {
	_c.Call.Return(run)
	return _c
}

// ApplyAndSave provides a mock function with given fields: tx, dependencyVuln, vulnEvent
func (_m *FirstPartyVulnRepository) ApplyAndSave(tx *gorm.DB, dependencyVuln *models.FirstPartyVuln, vulnEvent *models.VulnEvent) error {
	ret := _m.Called(tx, dependencyVuln, vulnEvent)

	if len(ret) == 0 {
		panic("no return value specified for ApplyAndSave")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.FirstPartyVuln, *models.VulnEvent) error); ok {
		r0 = rf(tx, dependencyVuln, vulnEvent)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_ApplyAndSave_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ApplyAndSave'
type FirstPartyVulnRepository_ApplyAndSave_Call struct {
	*mock.Call
}

// ApplyAndSave is a helper method to define mock.On call
//   - tx *gorm.DB
//   - dependencyVuln *models.FirstPartyVuln
//   - vulnEvent *models.VulnEvent
func (_e *FirstPartyVulnRepository_Expecter) ApplyAndSave(tx interface{}, dependencyVuln interface{}, vulnEvent interface{}) *FirstPartyVulnRepository_ApplyAndSave_Call {
	return &FirstPartyVulnRepository_ApplyAndSave_Call{Call: _e.mock.On("ApplyAndSave", tx, dependencyVuln, vulnEvent)}
}

func (_c *FirstPartyVulnRepository_ApplyAndSave_Call) Run(run func(tx *gorm.DB, dependencyVuln *models.FirstPartyVuln, vulnEvent *models.VulnEvent)) *FirstPartyVulnRepository_ApplyAndSave_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.FirstPartyVuln), args[2].(*models.VulnEvent))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_ApplyAndSave_Call) Return(_a0 error) *FirstPartyVulnRepository_ApplyAndSave_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_ApplyAndSave_Call) RunAndReturn(run func(*gorm.DB, *models.FirstPartyVuln, *models.VulnEvent) error) *FirstPartyVulnRepository_ApplyAndSave_Call {
	_c.Call.Return(run)
	return _c
}

// Begin provides a mock function with no fields
func (_m *FirstPartyVulnRepository) Begin() *gorm.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Begin")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func() *gorm.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// FirstPartyVulnRepository_Begin_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Begin'
type FirstPartyVulnRepository_Begin_Call struct {
	*mock.Call
}

// Begin is a helper method to define mock.On call
func (_e *FirstPartyVulnRepository_Expecter) Begin() *FirstPartyVulnRepository_Begin_Call {
	return &FirstPartyVulnRepository_Begin_Call{Call: _e.mock.On("Begin")}
}

func (_c *FirstPartyVulnRepository_Begin_Call) Run(run func()) *FirstPartyVulnRepository_Begin_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Begin_Call) Return(_a0 *gorm.DB) *FirstPartyVulnRepository_Begin_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_Begin_Call) RunAndReturn(run func() *gorm.DB) *FirstPartyVulnRepository_Begin_Call {
	_c.Call.Return(run)
	return _c
}

// Create provides a mock function with given fields: tx, t
func (_m *FirstPartyVulnRepository) Create(tx *gorm.DB, t *models.FirstPartyVuln) error {
	ret := _m.Called(tx, t)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.FirstPartyVuln) error); ok {
		r0 = rf(tx, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type FirstPartyVulnRepository_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - tx *gorm.DB
//   - t *models.FirstPartyVuln
func (_e *FirstPartyVulnRepository_Expecter) Create(tx interface{}, t interface{}) *FirstPartyVulnRepository_Create_Call {
	return &FirstPartyVulnRepository_Create_Call{Call: _e.mock.On("Create", tx, t)}
}

func (_c *FirstPartyVulnRepository_Create_Call) Run(run func(tx *gorm.DB, t *models.FirstPartyVuln)) *FirstPartyVulnRepository_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.FirstPartyVuln))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Create_Call) Return(_a0 error) *FirstPartyVulnRepository_Create_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_Create_Call) RunAndReturn(run func(*gorm.DB, *models.FirstPartyVuln) error) *FirstPartyVulnRepository_Create_Call {
	_c.Call.Return(run)
	return _c
}

// CreateBatch provides a mock function with given fields: tx, ts
func (_m *FirstPartyVulnRepository) CreateBatch(tx *gorm.DB, ts []models.FirstPartyVuln) error {
	ret := _m.Called(tx, ts)

	if len(ret) == 0 {
		panic("no return value specified for CreateBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.FirstPartyVuln) error); ok {
		r0 = rf(tx, ts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_CreateBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBatch'
type FirstPartyVulnRepository_CreateBatch_Call struct {
	*mock.Call
}

// CreateBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - ts []models.FirstPartyVuln
func (_e *FirstPartyVulnRepository_Expecter) CreateBatch(tx interface{}, ts interface{}) *FirstPartyVulnRepository_CreateBatch_Call {
	return &FirstPartyVulnRepository_CreateBatch_Call{Call: _e.mock.On("CreateBatch", tx, ts)}
}

func (_c *FirstPartyVulnRepository_CreateBatch_Call) Run(run func(tx *gorm.DB, ts []models.FirstPartyVuln)) *FirstPartyVulnRepository_CreateBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.FirstPartyVuln))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_CreateBatch_Call) Return(_a0 error) *FirstPartyVulnRepository_CreateBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_CreateBatch_Call) RunAndReturn(run func(*gorm.DB, []models.FirstPartyVuln) error) *FirstPartyVulnRepository_CreateBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: tx, id
func (_m *FirstPartyVulnRepository) Delete(tx *gorm.DB, id string) error {
	ret := _m.Called(tx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string) error); ok {
		r0 = rf(tx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type FirstPartyVulnRepository_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - tx *gorm.DB
//   - id string
func (_e *FirstPartyVulnRepository_Expecter) Delete(tx interface{}, id interface{}) *FirstPartyVulnRepository_Delete_Call {
	return &FirstPartyVulnRepository_Delete_Call{Call: _e.mock.On("Delete", tx, id)}
}

func (_c *FirstPartyVulnRepository_Delete_Call) Run(run func(tx *gorm.DB, id string)) *FirstPartyVulnRepository_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Delete_Call) Return(_a0 error) *FirstPartyVulnRepository_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_Delete_Call) RunAndReturn(run func(*gorm.DB, string) error) *FirstPartyVulnRepository_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteBatch provides a mock function with given fields: tx, ids
func (_m *FirstPartyVulnRepository) DeleteBatch(tx *gorm.DB, ids []models.FirstPartyVuln) error {
	ret := _m.Called(tx, ids)

	if len(ret) == 0 {
		panic("no return value specified for DeleteBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.FirstPartyVuln) error); ok {
		r0 = rf(tx, ids)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_DeleteBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteBatch'
type FirstPartyVulnRepository_DeleteBatch_Call struct {
	*mock.Call
}

// DeleteBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - ids []models.FirstPartyVuln
func (_e *FirstPartyVulnRepository_Expecter) DeleteBatch(tx interface{}, ids interface{}) *FirstPartyVulnRepository_DeleteBatch_Call {
	return &FirstPartyVulnRepository_DeleteBatch_Call{Call: _e.mock.On("DeleteBatch", tx, ids)}
}

func (_c *FirstPartyVulnRepository_DeleteBatch_Call) Run(run func(tx *gorm.DB, ids []models.FirstPartyVuln)) *FirstPartyVulnRepository_DeleteBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.FirstPartyVuln))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_DeleteBatch_Call) Return(_a0 error) *FirstPartyVulnRepository_DeleteBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_DeleteBatch_Call) RunAndReturn(run func(*gorm.DB, []models.FirstPartyVuln) error) *FirstPartyVulnRepository_DeleteBatch_Call {
	_c.Call.Return(run)
	return _c
}

// GetByAssetId provides a mock function with given fields: tx, assetId
func (_m *FirstPartyVulnRepository) GetByAssetId(tx *gorm.DB, assetId uuid.UUID) ([]models.FirstPartyVuln, error) {
	ret := _m.Called(tx, assetId)

	if len(ret) == 0 {
		panic("no return value specified for GetByAssetId")
	}

	var r0 []models.FirstPartyVuln
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID) ([]models.FirstPartyVuln, error)); ok {
		return rf(tx, assetId)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID) []models.FirstPartyVuln); ok {
		r0 = rf(tx, assetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.FirstPartyVuln)
		}
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, uuid.UUID) error); ok {
		r1 = rf(tx, assetId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_GetByAssetId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByAssetId'
type FirstPartyVulnRepository_GetByAssetId_Call struct {
	*mock.Call
}

// GetByAssetId is a helper method to define mock.On call
//   - tx *gorm.DB
//   - assetId uuid.UUID
func (_e *FirstPartyVulnRepository_Expecter) GetByAssetId(tx interface{}, assetId interface{}) *FirstPartyVulnRepository_GetByAssetId_Call {
	return &FirstPartyVulnRepository_GetByAssetId_Call{Call: _e.mock.On("GetByAssetId", tx, assetId)}
}

func (_c *FirstPartyVulnRepository_GetByAssetId_Call) Run(run func(tx *gorm.DB, assetId uuid.UUID)) *FirstPartyVulnRepository_GetByAssetId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_GetByAssetId_Call) Return(_a0 []models.FirstPartyVuln, _a1 error) *FirstPartyVulnRepository_GetByAssetId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_GetByAssetId_Call) RunAndReturn(run func(*gorm.DB, uuid.UUID) ([]models.FirstPartyVuln, error)) *FirstPartyVulnRepository_GetByAssetId_Call {
	_c.Call.Return(run)
	return _c
}

// GetByAssetVersionPaged provides a mock function with given fields: tx, assetVersionName, assetID, pageInfo, search, filter, sort
func (_m *FirstPartyVulnRepository) GetByAssetVersionPaged(tx *gorm.DB, assetVersionName string, assetID uuid.UUID, pageInfo core.PageInfo, search string, filter []core.FilterQuery, sort []core.SortQuery) (core.Paged[models.FirstPartyVuln], map[string]int, error) {
	ret := _m.Called(tx, assetVersionName, assetID, pageInfo, search, filter, sort)

	if len(ret) == 0 {
		panic("no return value specified for GetByAssetVersionPaged")
	}

	var r0 core.Paged[models.FirstPartyVuln]
	var r1 map[string]int
	var r2 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) (core.Paged[models.FirstPartyVuln], map[string]int, error)); ok {
		return rf(tx, assetVersionName, assetID, pageInfo, search, filter, sort)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, string, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) core.Paged[models.FirstPartyVuln]); ok {
		r0 = rf(tx, assetVersionName, assetID, pageInfo, search, filter, sort)
	} else {
		r0 = ret.Get(0).(core.Paged[models.FirstPartyVuln])
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, string, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) map[string]int); ok {
		r1 = rf(tx, assetVersionName, assetID, pageInfo, search, filter, sort)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(map[string]int)
		}
	}

	if rf, ok := ret.Get(2).(func(*gorm.DB, string, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) error); ok {
		r2 = rf(tx, assetVersionName, assetID, pageInfo, search, filter, sort)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// FirstPartyVulnRepository_GetByAssetVersionPaged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByAssetVersionPaged'
type FirstPartyVulnRepository_GetByAssetVersionPaged_Call struct {
	*mock.Call
}

// GetByAssetVersionPaged is a helper method to define mock.On call
//   - tx *gorm.DB
//   - assetVersionName string
//   - assetID uuid.UUID
//   - pageInfo core.PageInfo
//   - search string
//   - filter []core.FilterQuery
//   - sort []core.SortQuery
func (_e *FirstPartyVulnRepository_Expecter) GetByAssetVersionPaged(tx interface{}, assetVersionName interface{}, assetID interface{}, pageInfo interface{}, search interface{}, filter interface{}, sort interface{}) *FirstPartyVulnRepository_GetByAssetVersionPaged_Call {
	return &FirstPartyVulnRepository_GetByAssetVersionPaged_Call{Call: _e.mock.On("GetByAssetVersionPaged", tx, assetVersionName, assetID, pageInfo, search, filter, sort)}
}

func (_c *FirstPartyVulnRepository_GetByAssetVersionPaged_Call) Run(run func(tx *gorm.DB, assetVersionName string, assetID uuid.UUID, pageInfo core.PageInfo, search string, filter []core.FilterQuery, sort []core.SortQuery)) *FirstPartyVulnRepository_GetByAssetVersionPaged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(string), args[2].(uuid.UUID), args[3].(core.PageInfo), args[4].(string), args[5].([]core.FilterQuery), args[6].([]core.SortQuery))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_GetByAssetVersionPaged_Call) Return(_a0 core.Paged[models.FirstPartyVuln], _a1 map[string]int, _a2 error) *FirstPartyVulnRepository_GetByAssetVersionPaged_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *FirstPartyVulnRepository_GetByAssetVersionPaged_Call) RunAndReturn(run func(*gorm.DB, string, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) (core.Paged[models.FirstPartyVuln], map[string]int, error)) *FirstPartyVulnRepository_GetByAssetVersionPaged_Call {
	_c.Call.Return(run)
	return _c
}

// GetDB provides a mock function with given fields: tx
func (_m *FirstPartyVulnRepository) GetDB(tx *gorm.DB) *gorm.DB {
	ret := _m.Called(tx)

	if len(ret) == 0 {
		panic("no return value specified for GetDB")
	}

	var r0 *gorm.DB
	if rf, ok := ret.Get(0).(func(*gorm.DB) *gorm.DB); ok {
		r0 = rf(tx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gorm.DB)
		}
	}

	return r0
}

// FirstPartyVulnRepository_GetDB_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDB'
type FirstPartyVulnRepository_GetDB_Call struct {
	*mock.Call
}

// GetDB is a helper method to define mock.On call
//   - tx *gorm.DB
func (_e *FirstPartyVulnRepository_Expecter) GetDB(tx interface{}) *FirstPartyVulnRepository_GetDB_Call {
	return &FirstPartyVulnRepository_GetDB_Call{Call: _e.mock.On("GetDB", tx)}
}

func (_c *FirstPartyVulnRepository_GetDB_Call) Run(run func(tx *gorm.DB)) *FirstPartyVulnRepository_GetDB_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_GetDB_Call) Return(_a0 *gorm.DB) *FirstPartyVulnRepository_GetDB_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_GetDB_Call) RunAndReturn(run func(*gorm.DB) *gorm.DB) *FirstPartyVulnRepository_GetDB_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultFirstPartyVulnsByOrgIdPaged provides a mock function with given fields: tx, userAllowedProjectIds, pageInfo, search, filter, sort
func (_m *FirstPartyVulnRepository) GetDefaultFirstPartyVulnsByOrgIdPaged(tx *gorm.DB, userAllowedProjectIds []string, pageInfo core.PageInfo, search string, filter []core.FilterQuery, sort []core.SortQuery) (core.Paged[models.FirstPartyVuln], error) {
	ret := _m.Called(tx, userAllowedProjectIds, pageInfo, search, filter, sort)

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultFirstPartyVulnsByOrgIdPaged")
	}

	var r0 core.Paged[models.FirstPartyVuln]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []string, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) (core.Paged[models.FirstPartyVuln], error)); ok {
		return rf(tx, userAllowedProjectIds, pageInfo, search, filter, sort)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, []string, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) core.Paged[models.FirstPartyVuln]); ok {
		r0 = rf(tx, userAllowedProjectIds, pageInfo, search, filter, sort)
	} else {
		r0 = ret.Get(0).(core.Paged[models.FirstPartyVuln])
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, []string, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) error); ok {
		r1 = rf(tx, userAllowedProjectIds, pageInfo, search, filter, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultFirstPartyVulnsByOrgIdPaged'
type FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call struct {
	*mock.Call
}

// GetDefaultFirstPartyVulnsByOrgIdPaged is a helper method to define mock.On call
//   - tx *gorm.DB
//   - userAllowedProjectIds []string
//   - pageInfo core.PageInfo
//   - search string
//   - filter []core.FilterQuery
//   - sort []core.SortQuery
func (_e *FirstPartyVulnRepository_Expecter) GetDefaultFirstPartyVulnsByOrgIdPaged(tx interface{}, userAllowedProjectIds interface{}, pageInfo interface{}, search interface{}, filter interface{}, sort interface{}) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call {
	return &FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call{Call: _e.mock.On("GetDefaultFirstPartyVulnsByOrgIdPaged", tx, userAllowedProjectIds, pageInfo, search, filter, sort)}
}

func (_c *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call) Run(run func(tx *gorm.DB, userAllowedProjectIds []string, pageInfo core.PageInfo, search string, filter []core.FilterQuery, sort []core.SortQuery)) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]string), args[2].(core.PageInfo), args[3].(string), args[4].([]core.FilterQuery), args[5].([]core.SortQuery))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call) Return(_a0 core.Paged[models.FirstPartyVuln], _a1 error) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call) RunAndReturn(run func(*gorm.DB, []string, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) (core.Paged[models.FirstPartyVuln], error)) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByOrgIdPaged_Call {
	_c.Call.Return(run)
	return _c
}

// GetDefaultFirstPartyVulnsByProjectIdPaged provides a mock function with given fields: tx, projectID, pageInfo, search, filter, sort
func (_m *FirstPartyVulnRepository) GetDefaultFirstPartyVulnsByProjectIdPaged(tx *gorm.DB, projectID uuid.UUID, pageInfo core.PageInfo, search string, filter []core.FilterQuery, sort []core.SortQuery) (core.Paged[models.FirstPartyVuln], error) {
	ret := _m.Called(tx, projectID, pageInfo, search, filter, sort)

	if len(ret) == 0 {
		panic("no return value specified for GetDefaultFirstPartyVulnsByProjectIdPaged")
	}

	var r0 core.Paged[models.FirstPartyVuln]
	var r1 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) (core.Paged[models.FirstPartyVuln], error)); ok {
		return rf(tx, projectID, pageInfo, search, filter, sort)
	}
	if rf, ok := ret.Get(0).(func(*gorm.DB, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) core.Paged[models.FirstPartyVuln]); ok {
		r0 = rf(tx, projectID, pageInfo, search, filter, sort)
	} else {
		r0 = ret.Get(0).(core.Paged[models.FirstPartyVuln])
	}

	if rf, ok := ret.Get(1).(func(*gorm.DB, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) error); ok {
		r1 = rf(tx, projectID, pageInfo, search, filter, sort)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetDefaultFirstPartyVulnsByProjectIdPaged'
type FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call struct {
	*mock.Call
}

// GetDefaultFirstPartyVulnsByProjectIdPaged is a helper method to define mock.On call
//   - tx *gorm.DB
//   - projectID uuid.UUID
//   - pageInfo core.PageInfo
//   - search string
//   - filter []core.FilterQuery
//   - sort []core.SortQuery
func (_e *FirstPartyVulnRepository_Expecter) GetDefaultFirstPartyVulnsByProjectIdPaged(tx interface{}, projectID interface{}, pageInfo interface{}, search interface{}, filter interface{}, sort interface{}) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call {
	return &FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call{Call: _e.mock.On("GetDefaultFirstPartyVulnsByProjectIdPaged", tx, projectID, pageInfo, search, filter, sort)}
}

func (_c *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call) Run(run func(tx *gorm.DB, projectID uuid.UUID, pageInfo core.PageInfo, search string, filter []core.FilterQuery, sort []core.SortQuery)) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(uuid.UUID), args[2].(core.PageInfo), args[3].(string), args[4].([]core.FilterQuery), args[5].([]core.SortQuery))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call) Return(_a0 core.Paged[models.FirstPartyVuln], _a1 error) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call) RunAndReturn(run func(*gorm.DB, uuid.UUID, core.PageInfo, string, []core.FilterQuery, []core.SortQuery) (core.Paged[models.FirstPartyVuln], error)) *FirstPartyVulnRepository_GetDefaultFirstPartyVulnsByProjectIdPaged_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ids
func (_m *FirstPartyVulnRepository) List(ids []string) ([]models.FirstPartyVuln, error) {
	ret := _m.Called(ids)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []models.FirstPartyVuln
	var r1 error
	if rf, ok := ret.Get(0).(func([]string) ([]models.FirstPartyVuln, error)); ok {
		return rf(ids)
	}
	if rf, ok := ret.Get(0).(func([]string) []models.FirstPartyVuln); ok {
		r0 = rf(ids)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.FirstPartyVuln)
		}
	}

	if rf, ok := ret.Get(1).(func([]string) error); ok {
		r1 = rf(ids)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type FirstPartyVulnRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ids []string
func (_e *FirstPartyVulnRepository_Expecter) List(ids interface{}) *FirstPartyVulnRepository_List_Call {
	return &FirstPartyVulnRepository_List_Call{Call: _e.mock.On("List", ids)}
}

func (_c *FirstPartyVulnRepository_List_Call) Run(run func(ids []string)) *FirstPartyVulnRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]string))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_List_Call) Return(_a0 []models.FirstPartyVuln, _a1 error) *FirstPartyVulnRepository_List_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_List_Call) RunAndReturn(run func([]string) ([]models.FirstPartyVuln, error)) *FirstPartyVulnRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListByScanner provides a mock function with given fields: assetVersionName, assetID, scannerID
func (_m *FirstPartyVulnRepository) ListByScanner(assetVersionName string, assetID uuid.UUID, scannerID string) ([]models.FirstPartyVuln, error) {
	ret := _m.Called(assetVersionName, assetID, scannerID)

	if len(ret) == 0 {
		panic("no return value specified for ListByScanner")
	}

	var r0 []models.FirstPartyVuln
	var r1 error
	if rf, ok := ret.Get(0).(func(string, uuid.UUID, string) ([]models.FirstPartyVuln, error)); ok {
		return rf(assetVersionName, assetID, scannerID)
	}
	if rf, ok := ret.Get(0).(func(string, uuid.UUID, string) []models.FirstPartyVuln); ok {
		r0 = rf(assetVersionName, assetID, scannerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.FirstPartyVuln)
		}
	}

	if rf, ok := ret.Get(1).(func(string, uuid.UUID, string) error); ok {
		r1 = rf(assetVersionName, assetID, scannerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_ListByScanner_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListByScanner'
type FirstPartyVulnRepository_ListByScanner_Call struct {
	*mock.Call
}

// ListByScanner is a helper method to define mock.On call
//   - assetVersionName string
//   - assetID uuid.UUID
//   - scannerID string
func (_e *FirstPartyVulnRepository_Expecter) ListByScanner(assetVersionName interface{}, assetID interface{}, scannerID interface{}) *FirstPartyVulnRepository_ListByScanner_Call {
	return &FirstPartyVulnRepository_ListByScanner_Call{Call: _e.mock.On("ListByScanner", assetVersionName, assetID, scannerID)}
}

func (_c *FirstPartyVulnRepository_ListByScanner_Call) Run(run func(assetVersionName string, assetID uuid.UUID, scannerID string)) *FirstPartyVulnRepository_ListByScanner_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(uuid.UUID), args[2].(string))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_ListByScanner_Call) Return(_a0 []models.FirstPartyVuln, _a1 error) *FirstPartyVulnRepository_ListByScanner_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_ListByScanner_Call) RunAndReturn(run func(string, uuid.UUID, string) ([]models.FirstPartyVuln, error)) *FirstPartyVulnRepository_ListByScanner_Call {
	_c.Call.Return(run)
	return _c
}

// Read provides a mock function with given fields: id
func (_m *FirstPartyVulnRepository) Read(id string) (models.FirstPartyVuln, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Read")
	}

	var r0 models.FirstPartyVuln
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.FirstPartyVuln, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) models.FirstPartyVuln); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.FirstPartyVuln)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FirstPartyVulnRepository_Read_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Read'
type FirstPartyVulnRepository_Read_Call struct {
	*mock.Call
}

// Read is a helper method to define mock.On call
//   - id string
func (_e *FirstPartyVulnRepository_Expecter) Read(id interface{}) *FirstPartyVulnRepository_Read_Call {
	return &FirstPartyVulnRepository_Read_Call{Call: _e.mock.On("Read", id)}
}

func (_c *FirstPartyVulnRepository_Read_Call) Run(run func(id string)) *FirstPartyVulnRepository_Read_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Read_Call) Return(_a0 models.FirstPartyVuln, _a1 error) *FirstPartyVulnRepository_Read_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FirstPartyVulnRepository_Read_Call) RunAndReturn(run func(string) (models.FirstPartyVuln, error)) *FirstPartyVulnRepository_Read_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: tx, vuln
func (_m *FirstPartyVulnRepository) Save(tx *gorm.DB, vuln *models.FirstPartyVuln) error {
	ret := _m.Called(tx, vuln)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, *models.FirstPartyVuln) error); ok {
		r0 = rf(tx, vuln)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type FirstPartyVulnRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - tx *gorm.DB
//   - vuln *models.FirstPartyVuln
func (_e *FirstPartyVulnRepository_Expecter) Save(tx interface{}, vuln interface{}) *FirstPartyVulnRepository_Save_Call {
	return &FirstPartyVulnRepository_Save_Call{Call: _e.mock.On("Save", tx, vuln)}
}

func (_c *FirstPartyVulnRepository_Save_Call) Run(run func(tx *gorm.DB, vuln *models.FirstPartyVuln)) *FirstPartyVulnRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].(*models.FirstPartyVuln))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Save_Call) Return(_a0 error) *FirstPartyVulnRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_Save_Call) RunAndReturn(run func(*gorm.DB, *models.FirstPartyVuln) error) *FirstPartyVulnRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// SaveBatch provides a mock function with given fields: tx, vulns
func (_m *FirstPartyVulnRepository) SaveBatch(tx *gorm.DB, vulns []models.FirstPartyVuln) error {
	ret := _m.Called(tx, vulns)

	if len(ret) == 0 {
		panic("no return value specified for SaveBatch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*gorm.DB, []models.FirstPartyVuln) error); ok {
		r0 = rf(tx, vulns)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_SaveBatch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveBatch'
type FirstPartyVulnRepository_SaveBatch_Call struct {
	*mock.Call
}

// SaveBatch is a helper method to define mock.On call
//   - tx *gorm.DB
//   - vulns []models.FirstPartyVuln
func (_e *FirstPartyVulnRepository_Expecter) SaveBatch(tx interface{}, vulns interface{}) *FirstPartyVulnRepository_SaveBatch_Call {
	return &FirstPartyVulnRepository_SaveBatch_Call{Call: _e.mock.On("SaveBatch", tx, vulns)}
}

func (_c *FirstPartyVulnRepository_SaveBatch_Call) Run(run func(tx *gorm.DB, vulns []models.FirstPartyVuln)) *FirstPartyVulnRepository_SaveBatch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*gorm.DB), args[1].([]models.FirstPartyVuln))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_SaveBatch_Call) Return(_a0 error) *FirstPartyVulnRepository_SaveBatch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_SaveBatch_Call) RunAndReturn(run func(*gorm.DB, []models.FirstPartyVuln) error) *FirstPartyVulnRepository_SaveBatch_Call {
	_c.Call.Return(run)
	return _c
}

// Transaction provides a mock function with given fields: txFunc
func (_m *FirstPartyVulnRepository) Transaction(txFunc func(*gorm.DB) error) error {
	ret := _m.Called(txFunc)

	if len(ret) == 0 {
		panic("no return value specified for Transaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(func(*gorm.DB) error) error); ok {
		r0 = rf(txFunc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FirstPartyVulnRepository_Transaction_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Transaction'
type FirstPartyVulnRepository_Transaction_Call struct {
	*mock.Call
}

// Transaction is a helper method to define mock.On call
//   - txFunc func(*gorm.DB) error
func (_e *FirstPartyVulnRepository_Expecter) Transaction(txFunc interface{}) *FirstPartyVulnRepository_Transaction_Call {
	return &FirstPartyVulnRepository_Transaction_Call{Call: _e.mock.On("Transaction", txFunc)}
}

func (_c *FirstPartyVulnRepository_Transaction_Call) Run(run func(txFunc func(*gorm.DB) error)) *FirstPartyVulnRepository_Transaction_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(func(*gorm.DB) error))
	})
	return _c
}

func (_c *FirstPartyVulnRepository_Transaction_Call) Return(_a0 error) *FirstPartyVulnRepository_Transaction_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *FirstPartyVulnRepository_Transaction_Call) RunAndReturn(run func(func(*gorm.DB) error) error) *FirstPartyVulnRepository_Transaction_Call {
	_c.Call.Return(run)
	return _c
}

// NewFirstPartyVulnRepository creates a new instance of FirstPartyVulnRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFirstPartyVulnRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *FirstPartyVulnRepository {
	mock := &FirstPartyVulnRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
