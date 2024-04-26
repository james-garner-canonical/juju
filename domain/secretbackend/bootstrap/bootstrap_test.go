// Copyright 2024 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package bootstrap

import (
	"context"

	jc "github.com/juju/testing/checkers"
	gc "gopkg.in/check.v1"

	coremodel "github.com/juju/juju/core/model"
	schematesting "github.com/juju/juju/domain/schema/testing"
)

type bootstrapSuite struct {
	schematesting.ControllerSuite
}

var _ = gc.Suite(&bootstrapSuite{})

func (s *bootstrapSuite) TestCreateDefaultBackendsIAAS(c *gc.C) {
	err := CreateDefaultBackends(coremodel.IAAS)(context.Background(), s.TxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)

	var name string
	row := s.DB().QueryRow("SELECT name FROM secret_backend where backend_type_id = ?", 0) // 0 = controller
	c.Assert(row.Scan(&name), jc.ErrorIsNil)
	c.Assert(name, gc.Equals, "internal")
}

func (s *bootstrapSuite) TestCreateDefaultBackendsCAAS(c *gc.C) {
	err := CreateDefaultBackends(coremodel.CAAS)(context.Background(), s.TxnRunner(), s.NoopTxnRunner())
	c.Assert(err, jc.ErrorIsNil)

	var name string
	row := s.DB().QueryRow("SELECT name FROM secret_backend where backend_type_id = ?", 1) // 1 = kubernetes
	c.Assert(row.Scan(&name), jc.ErrorIsNil)
	c.Assert(name, gc.Equals, "kubernetes")
}
