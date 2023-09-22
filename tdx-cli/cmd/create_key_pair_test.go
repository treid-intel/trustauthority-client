/*
 *   Copyright (c) 2022 Intel Corporation
 *   All rights reserved.
 *   SPDX-License-Identifier: BSD-3-Clause
 */

package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/treid-intel/trustauthority-client/tdx-cli/constants"
)

const (
	privateKeyPath = "privatekey.pem"
	publicKeyPath  = "publickey.pem"
	confFilePath   = "config.json"
)

func TestCreateKeyPairCmd(t *testing.T) {

	defer func() {
		os.Remove(privateKeyPath)
		os.Remove(publicKeyPath)
	}()

	tt := []struct {
		args        []string
		wantErr     bool
		description string
	}{
		{
			args: []string{
				constants.CreateKeyPairCmd,
				"--" + constants.PublicKeyPathOption,
				publicKeyPath,
			},
			wantErr:     false,
			description: "Test with all valid inputs",
		},
		{
			args: []string{
				constants.CreateKeyPairCmd,
				"--" + constants.PublicKeyPathOption,
				"../cmd",
			},
			wantErr:     true,
			description: "Test with directory as public key path",
		},
	}

	for _, tc := range tt {
		_, err := execute(t, rootCmd, tc.args...)

		if tc.wantErr == true {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func execute(t *testing.T, c *cobra.Command, args ...string) (string, error) {
	t.Helper()

	buf := new(bytes.Buffer)
	c.SetOut(buf)
	c.SetErr(buf)
	c.SetArgs(args)

	err := c.Execute()
	return strings.TrimSpace(buf.String()), err
}
