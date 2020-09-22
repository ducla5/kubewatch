/*
Copyright 2016 Skippbox, Ltd.

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

package chatwork

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bitnami-labs/kubewatch/config"
)

func TestChatworkInit(t *testing.T) {
	s := &Chatwork{}
	expectedError := fmt.Errorf(chatworkErrMsg, "Missing chatwork token or room")

	var Tests = []struct {
		chatwork config.Chatwork
		err      error
	}{
		{config.Chatwork{Token: "foo", Room: "bar"}, nil},
		{config.Chatwork{Token: "foo"}, expectedError},
		{config.Chatwork{Room: "bar"}, expectedError},
		{config.Chatwork{}, expectedError},
	}

	for _, tt := range Tests {
		c := &config.Config{}
		c.Handler.Chatwork = tt.chatwork
		if err := s.Init(c); !reflect.DeepEqual(err, tt.err) {
			t.Fatalf("Init(): %v", err)
		}
	}
}
