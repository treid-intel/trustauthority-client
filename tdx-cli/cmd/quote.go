/*
 *   Copyright (c) 2022 Intel Corporation
 *   All rights reserved.
 *   SPDX-License-Identifier: BSD-3-Clause
 */

package cmd

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/treid-intel/trustauthority-client/go-tdx"
	"github.com/treid-intel/trustauthority-client/tdx-cli/constants"
)

// quoteCmd represents the quote command
var quoteCmd = &cobra.Command{
	Use:   constants.QuoteCmd,
	Short: "Fetches the TD quote",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := getQuote(cmd)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(quoteCmd)
	quoteCmd.Flags().StringP(constants.NonceOption, "n", "", "Nonce in base64 encoded format")
	quoteCmd.Flags().StringP(constants.UserDataOption, "u", "", "User Data in base64 encoded format")
}

func getQuote(cmd *cobra.Command) error {

	userData, err := cmd.Flags().GetString(constants.UserDataOption)
	if err != nil {
		return err
	}

	nonce, err := cmd.Flags().GetString(constants.NonceOption)
	if err != nil {
		return err
	}

	var userDataBytes []byte
	if userData != "" {
		userDataBytes, err = base64.StdEncoding.DecodeString(userData)
		if err != nil {
			return errors.Wrap(err, "Error while base64 decoding of userdata")
		}
	}

	var nonceBytes []byte
	if nonce != "" {
		nonceBytes, err = base64.StdEncoding.DecodeString(nonce)
		if err != nil {
			return errors.Wrap(err, "Error while base64 decoding of nonce")
		}
	}

	adapter, err := tdx.NewEvidenceAdapter(userDataBytes, nil)
	if err != nil {
		return errors.Wrap(err, "Error while creating tdx adapter")
	}
	evidence, err := adapter.CollectEvidence(nonceBytes)
	if err != nil {
		return errors.Wrap(err, "Failed to collect evidence")
	}

	fmt.Fprintln(os.Stdout, evidence.Evidence)
	return nil
}
