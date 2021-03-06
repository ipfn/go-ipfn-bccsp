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

package bccsp

// KeyStore represents a storage system for cryptographic keys.
// It allows to store and retrieve bccsp.Key objects.
// The KeyStore can be read only, in that case StoreKey will return
// an error.
type KeyStore interface {
	// Key returns the key this CSP associates to
	// the Subject Key Identifier ski.
	Key(ski []byte) (k Key, err error)

	// StoreKey stores the key k in this KeyStore.
	// If this KeyStore is read only then the method will fail.
	StoreKey(k Key) (err error)

	// ReadOnly returns true if this KeyStore is read only, false otherwise.
	// If ReadOnly is true then StoreKey will fail.
	ReadOnly() bool
}
