// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package state

import (
	"context"
	"database/sql"
	"time"

	"github.com/juju/names/v5"
	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	corepermission "github.com/juju/juju/core/permission"
	"github.com/juju/juju/core/user"
	accesserrors "github.com/juju/juju/domain/access/errors"
	schematesting "github.com/juju/juju/domain/schema/testing"
	"github.com/juju/juju/internal/uuid"
	jujutesting "github.com/juju/juju/testing"
)

type permissionStateSuite struct {
	schematesting.ControllerSuite

	debug bool
}

var _ = gc.Suite(&permissionStateSuite{})

func (s *permissionStateSuite) SetUpTest(c *gc.C) {
	s.ControllerSuite.SetUpTest(c)
	s.ensureUser(c, "42", "admin", "42") // model owner
	s.ensureUser(c, "123", "bob", "42")
	s.ensureUser(c, "456", "sue", "42")
	s.ensureCloud(c, "987", "test-cloud")
	s.ensureCloud(c, "654", "another-cloud")
	s.ensureModel(c, "default-model")
	s.ensureModel(c, "model-uuid")
}

func (s *permissionStateSuite) TestCreatePermissionModel(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	userAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Check(userAccess.UserID, gc.Equals, "123")
	c.Check(userAccess.UserTag, gc.Equals, names.NewUserTag("bob"))
	c.Check(userAccess.Object.Id(), gc.Equals, "model-uuid")
	c.Check(userAccess.Access, gc.Equals, corepermission.WriteAccess)
	c.Check(userAccess.DisplayName, gc.Equals, "bob")
	c.Check(userAccess.UserName, gc.Equals, "bob")
	c.Check(userAccess.CreatedBy, gc.Equals, names.NewUserTag("admin"))

	s.checkPermissionRow(c, corepermission.WriteAccess, "123", "model-uuid")
}

func (s *permissionStateSuite) TestCreatePermissionCloud(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	userAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "test-cloud",
				ObjectType: corepermission.Cloud,
			},
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Check(userAccess.UserID, gc.Equals, "123")
	c.Check(userAccess.UserTag, gc.Equals, names.NewUserTag("bob"))
	c.Check(userAccess.Object.Id(), gc.Equals, "test-cloud")
	c.Check(userAccess.Access, gc.Equals, corepermission.AddModelAccess)
	c.Check(userAccess.DisplayName, gc.Equals, "bob")
	c.Check(userAccess.UserName, gc.Equals, "bob")
	c.Check(userAccess.CreatedBy, gc.Equals, names.NewUserTag("admin"))

	s.checkPermissionRow(c, corepermission.AddModelAccess, "123", "test-cloud")
}

func (s *permissionStateSuite) TestCreatePermissionController(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	userAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "controller",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	c.Check(userAccess.UserID, gc.Equals, "123")
	c.Check(userAccess.UserTag, gc.Equals, names.NewUserTag("bob"))
	c.Check(userAccess.Object.Id(), gc.Equals, "controller")
	c.Check(userAccess.Access, gc.Equals, corepermission.SuperuserAccess)
	c.Check(userAccess.DisplayName, gc.Equals, "bob")
	c.Check(userAccess.UserName, gc.Equals, "bob")
	c.Check(userAccess.CreatedBy, gc.Equals, names.NewUserTag("admin"))

	s.checkPermissionRow(c, corepermission.SuperuserAccess, "123", "controller")
}

func (s *permissionStateSuite) TestCreatePermissionForModelWithBadInfo(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "foo-bar",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.ReadAccess,
		},
	})
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionTargetInvalid)
}

func (s *permissionStateSuite) TestCreatePermissionForControllerWithBadInfo(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "foo-bar",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionTargetInvalid)
}

func (s *permissionStateSuite) checkPermissionRow(c *gc.C, access corepermission.Access, expectedGrantTo, expectedGrantON string) {
	db := s.DB()
	// Find the id for access
	accessRow := db.QueryRow(`
SELECT id
FROM permission_access_type
WHERE type = ?
`, access)
	c.Assert(accessRow.Err(), jc.ErrorIsNil)
	var accessTypeID int
	err := accessRow.Scan(&accessTypeID)
	c.Assert(err, jc.ErrorIsNil)

	// Find the permission
	row := db.QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on
FROM permission
`)
	c.Assert(row.Err(), jc.ErrorIsNil)
	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)
	err = row.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIsNil)

	// Verify the permission as expected.
	c.Check(userUuid, gc.Not(gc.Equals), "")
	c.Check(permissionTypeId, gc.Equals, accessTypeID)
	c.Check(grantTo, gc.Equals, expectedGrantTo)
	c.Check(grantOn, gc.Equals, expectedGrantON)
}

func (s *permissionStateSuite) TestCreatePermissionErrorNoUser(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "testme",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIs, accesserrors.UserNotFound)
}

func (s *permissionStateSuite) TestCreatePermissionErrorDuplicate(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	spec := corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.ReadAccess,
		},
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), spec)
	c.Assert(err, jc.ErrorIsNil)

	// Find the permission
	row := s.DB().QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
WHERE permission_type_id = 0
`)
	c.Assert(row.Err(), jc.ErrorIsNil)

	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)
	err = row.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIsNil)

	// Ensure each combination of grant_on and grant_two
	// is unique
	spec.Access = corepermission.WriteAccess
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), spec)
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionAlreadyExists)
	row2 := s.DB().QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
WHERE permission_type_id = 1
`)
	c.Assert(row2.Err(), jc.ErrorIsNil)
	err = row2.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIs, sql.ErrNoRows)
}

func (s *permissionStateSuite) TestDeletePermission(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "model-uuid",
		ObjectType: corepermission.Model,
	}
	spec := corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.ReadAccess,
		},
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), spec)
	c.Assert(err, jc.ErrorIsNil)

	err = st.DeletePermission(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)

	db := s.DB()

	var num int
	err = db.QueryRowContext(context.Background(), "SELECT count(*) FROM permission").Scan(&num)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(num, gc.Equals, 0)
}

func (s *permissionStateSuite) TestDeletePermissionFailUserNotFound(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "model-uuid",
		ObjectType: corepermission.Model,
	}
	err := st.DeletePermission(context.Background(), "testme", target)
	c.Assert(err, jc.ErrorIs, accesserrors.UserNotFound)
}

func (s *permissionStateSuite) TestDeletePermissionDoesNotExist(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "model-uuid",
		ObjectType: corepermission.Model,
	}

	// Don't fail if the permission does not exist.
	err := st.DeletePermission(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) TestReadUserAccessForTarget(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "controller",
		ObjectType: corepermission.Controller,
	}
	createUserAccess, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)

	row2 := s.DB().QueryRow(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
WHERE grant_to = 123
`)
	c.Assert(row2.Err(), jc.ErrorIsNil)
	err = row2.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
	c.Assert(err, jc.ErrorIsNil)
	c.Logf("%q, %d, to %q, on %q", userUuid, permissionTypeId, grantTo, grantOn)

	readUserAccess, err := st.ReadUserAccessForTarget(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(createUserAccess, gc.DeepEquals, readUserAccess)
}

func (s *permissionStateSuite) TestReadUserAccessLevelForTarget(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	target := corepermission.ID{
		Key:        "test-cloud",
		ObjectType: corepermission.Cloud,
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: target,
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)

	readUserAccessType, err := st.ReadUserAccessLevelForTarget(context.Background(), "bob", target)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(readUserAccessType, gc.Equals, corepermission.AddModelAccess)
}

func (s *permissionStateSuite) TestReadAllUserAccessForUser(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	userAccesses, err := st.ReadAllUserAccessForUser(context.Background(), "bob")
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(userAccesses, gc.HasLen, 4)
	for _, access := range userAccesses {
		c.Assert(access.UserName, gc.Equals, "bob")
		c.Assert(access.CreatedBy.Id(), gc.Equals, "admin")
	}
	accessOne := userAccesses[0]
	c.Assert(accessOne.Access, gc.Equals, corepermission.AddModelAccess)
}

func (s *permissionStateSuite) TestReadAllUserAccessForTarget(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)
	targetCloud := corepermission.ID{
		Key:        "test-cloud",
		ObjectType: corepermission.Cloud,
	}
	userAccesses, err := st.ReadAllUserAccessForTarget(context.Background(), targetCloud)
	c.Assert(err, jc.ErrorIsNil)

	c.Assert(userAccesses, gc.HasLen, 2)
	accessZero := userAccesses[0]
	c.Check(accessZero.Access, gc.Equals, corepermission.AddModelAccess)
	c.Check(accessZero.Object, gc.Equals, names.NewCloudTag("test-cloud"))
	accessOne := userAccesses[1]
	c.Check(accessOne.Access, gc.Equals, corepermission.AddModelAccess)
	c.Check(accessOne.Object, gc.Equals, names.NewCloudTag("test-cloud"))

	c.Check(accessZero.UserID, gc.Not(gc.Equals), accessOne.UserID)
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeCloud(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	users, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "bob", corepermission.Cloud)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(users, gc.HasLen, 2)

	var foundTestCloud, foundAnotherCloud bool
	for _, userAccess := range users {
		c.Check(userAccess.UserTag.Id(), gc.Equals, "bob")
		c.Check(userAccess.UserName, gc.Equals, "bob")
		c.Check(userAccess.CreatedBy.Id(), gc.Equals, "admin")
		c.Check(userAccess.UserID, gc.Equals, "123")
		c.Check(userAccess.Access, gc.Equals, corepermission.AddModelAccess)
		if userAccess.Object.Id() == "test-cloud" {
			foundTestCloud = true
		}
		if userAccess.Object.Id() == "another-cloud" {
			foundAnotherCloud = true
		}
	}
	c.Check(foundTestCloud && foundAnotherCloud, jc.IsTrue, gc.Commentf("%+v", users))
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeModel(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	users, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "bob", corepermission.Model)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(users, gc.HasLen, 2)

	var admin, write bool
	for _, userAccess := range users {
		c.Check(userAccess.UserTag.Id(), gc.Equals, "bob")
		c.Check(userAccess.UserName, gc.Equals, "bob")
		c.Check(userAccess.CreatedBy.Id(), gc.Equals, "admin")
		c.Check(userAccess.UserID, gc.Equals, "123")
		if userAccess.Access == corepermission.WriteAccess {
			write = true
			c.Check(userAccess.Object.Id(), gc.Equals, "default-model")
		}
		if userAccess.Access == corepermission.AdminAccess {
			admin = true
			c.Check(userAccess.Object.Id(), gc.Equals, "model-uuid")
		}
	}
	c.Assert(admin && write, jc.IsTrue, gc.Commentf("%+v", users))
}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeController(c *gc.C) {
	st := NewPermissionState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	s.setupForRead(c, st)

	users, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "admin", corepermission.Controller)
	c.Assert(err, jc.ErrorIsNil)
	c.Assert(users, gc.HasLen, 1)
	userAccess := users[0]
	c.Check(userAccess.UserTag.Id(), gc.Equals, "admin", gc.Commentf("%+v", users))
	c.Check(userAccess.UserName, gc.Equals, "admin", gc.Commentf("%+v", users))
	c.Check(userAccess.CreatedBy.Id(), gc.Equals, "admin", gc.Commentf("%+v", users))
	c.Check(userAccess.UserID, gc.Equals, "42", gc.Commentf("%+v", users))
	c.Check(userAccess.Access, gc.Equals, corepermission.SuperuserAccess, gc.Commentf("%+v", users))
}

func (s *permissionStateSuite) printPermissions(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT uuid, permission_type_id, grant_to, grant_on 
FROM permission
--WHERE grant_to = "bob"
`)
	defer func() { _ = rows.Close() }()
	var (
		userUuid, grantTo, grantOn string
		permissionTypeId           int
	)

	for rows.Next() {
		err := rows.Scan(&userUuid, &permissionTypeId, &grantTo, &grantOn)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("%q, %d, %q, %q", userUuid, permissionTypeId, grantTo, grantOn)
	}

}

func (s *permissionStateSuite) printUsers(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT uuid, name, created_by_uuid
FROM user
`)
	defer func() { _ = rows.Close() }()
	var (
		rowUUID, name string
		creatorUUID   user.UUID
	)

	for rows.Next() {
		err := rows.Scan(&rowUUID, &name, &creatorUUID)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("%q, %q, %q", rowUUID, name, creatorUUID)
	}

}

func (s *permissionStateSuite) printClouds(c *gc.C) {
	rows, _ := s.DB().Query(`
SELECT uuid, name
FROM cloud
`)
	defer func() { _ = rows.Close() }()
	var (
		rowUUID, name string
	)

	for rows.Next() {
		err := rows.Scan(&rowUUID, &name)
		c.Assert(err, jc.ErrorIsNil)
		c.Logf("%q, %q", rowUUID, name)
	}

}

func (s *permissionStateSuite) TestReadAllAccessForUserAndObjectTypeNotFound(c *gc.C) {
	st := NewState(s.TxnRunnerFactory(), jujutesting.NewCheckLogger(c))

	_, err := st.ReadAllAccessForUserAndObjectType(context.Background(), "bob", corepermission.Cloud)
	c.Assert(err, jc.ErrorIs, accesserrors.PermissionNotFound)
}

func (s *permissionStateSuite) setupForRead(c *gc.C, st *PermissionState) {
	targetCloud := corepermission.ID{
		Key:        "test-cloud",
		ObjectType: corepermission.Cloud,
	}
	_, err := st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: targetCloud,
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "sue",
		AccessSpec: corepermission.AccessSpec{
			Target: targetCloud,
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "model-uuid",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.AdminAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "another-cloud",
				ObjectType: corepermission.Cloud,
			},
			Access: corepermission.AddModelAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "bob",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "default-model",
				ObjectType: corepermission.Model,
			},
			Access: corepermission.WriteAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	_, err = st.CreatePermission(context.Background(), uuid.MustNewUUID(), corepermission.UserAccessSpec{
		User: "admin",
		AccessSpec: corepermission.AccessSpec{
			Target: corepermission.ID{
				Key:        "controller",
				ObjectType: corepermission.Controller,
			},
			Access: corepermission.SuperuserAccess,
		},
	})
	c.Assert(err, jc.ErrorIsNil)
	if s.debug {
		s.printUsers(c)
		s.printClouds(c)
		s.printPermissions(c)
	}
}

func (s *permissionStateSuite) ensureUser(c *gc.C, userUUID, name, createdByUUID string) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user (uuid, name, display_name, removed, created_by_uuid, created_at)
			VALUES (?, ?, ?, ?, ?, ?)
		`, userUUID, name, name, false, createdByUUID, time.Now())
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
	err = s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO user_authentication ( user_uuid, disabled)
			VALUES (?, ?)
		`, userUUID, false)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) ensureModel(c *gc.C, uuid string) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO model_list (uuid)
			VALUES (?)
		`, uuid)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
}

func (s *permissionStateSuite) ensureCloud(c *gc.C, uuid, name string) {
	err := s.TxnRunner().StdTxn(context.Background(), func(ctx context.Context, tx *sql.Tx) error {
		_, err := tx.ExecContext(ctx, `
			INSERT INTO cloud (uuid, name, cloud_type_id, endpoint, skip_tls_verify)
			VALUES (?, ?, 7, "test-endpoint", true)
		`, uuid, name)
		return err
	})
	c.Assert(err, jc.ErrorIsNil)
}
