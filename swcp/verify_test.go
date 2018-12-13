// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swcp

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	bccsp "github.com/ipfn/go-ipfn-bccsp"
	mocks2 "github.com/ipfn/go-ipfn-bccsp/mocks"
	"github.com/ipfn/go-ipfn-bccsp/swcp/mocks"
)

func TestVerify(t *testing.T) {
	t.Parallel()

	expectedKey := &mocks2.MockKey{}
	expectetSignature := []byte{1, 2, 3, 4, 5}
	expectetDigest := []byte{1, 2, 3, 4}
	expectedOpts := &mocks2.SignerOpts{}
	expectetValue := true
	expectedErr := errors.New("Expected Error")

	verifiers := make(map[reflect.Type]bccsp.Verifier)
	verifiers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Verifier{
		KeyArg:       expectedKey,
		SignatureArg: expectetSignature,
		DigestArg:    expectetDigest,
		OptsArg:      expectedOpts,
		Value:        expectetValue,
		Err:          nil,
	}
	csp := CSP{verifiers: verifiers}
	value, err := csp.Verify(expectedKey, expectetSignature, expectetDigest, expectedOpts)
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	verifiers = make(map[reflect.Type]bccsp.Verifier)
	verifiers[reflect.TypeOf(&mocks2.MockKey{})] = &mocks.Verifier{
		KeyArg:       expectedKey,
		SignatureArg: expectetSignature,
		DigestArg:    expectetDigest,
		OptsArg:      expectedOpts,
		Value:        false,
		Err:          expectedErr,
	}
	csp = CSP{verifiers: verifiers}
	value, err = csp.Verify(expectedKey, expectetSignature, expectetDigest, expectedOpts)
	assert.False(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
}
