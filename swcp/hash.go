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
	"fmt"
	"hash"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
)

type hasher struct {
	algo bccsp.HashType
	impl func() hash.Hash
}

func (c *hasher) Hash(msg []byte, algo bccsp.HashType) ([]byte, error) {
	if algo != c.algo {
		return nil, fmt.Errorf("hasher does not implement %s", algo)
	}
	h := c.impl()
	h.Write(msg)
	return h.Sum(nil), nil
}

func (c *hasher) Hasher(algo bccsp.HashType) (hash.Hash, error) {
	if algo != c.algo {
		return nil, fmt.Errorf("hasher does not implement %s", algo)
	}
	return c.impl(), nil
}
