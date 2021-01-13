/*
Copyright 2020 The Nocalhost Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clientgoutils

import (
	"fmt"
	"testing"
)

func TestNewClientGoUtils(t *testing.T) {

}

func TestClientGoUtils_Create(t *testing.T) {
	client, err := NewClientGoUtils("", "nh6ihig")
	if err != nil {
		panic(err)
	}
	secret, err := client.GetSecret("aaa")
	if err != nil {
		fmt.Printf("err:%s", err.Error())
		fmt.Printf("%v", secret)
		fmt.Println(secret.Name)
	} else {
		fmt.Printf("%v\n", secret)
	}
}
