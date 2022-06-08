// Copyright 2018 The goftp Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package integrations

import (
	"testing"

	"github.com/bin-work/go-example/pkg/bftp"

	"github.com/stretchr/testify/assert"
)

func runServer(t *testing.T, opt *bftp.Options, notifiers []bftp.Filter, execute func()) {
	s, err := bftp.NewServer(opt)
	assert.NoError(t, err)
	for _, notifier := range notifiers {
		s.RegisterNotifer(notifier)
	}
	go func() {
		err := s.ListenAndServe()
		assert.EqualError(t, err, bftp.ErrServerClosed.Error())
	}()

	execute()

	assert.NoError(t, s.Shutdown())
}
