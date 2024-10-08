// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

// declare main package

package main

// import cli starter and fmt
import (
	"github.com/hidevopsio/hiboot/pkg/app/cli"
	_ "hicrypto/cmd"
)

// main function
func main() {
	// create new cli application and run it
	cli.NewApplication().Run()
}
