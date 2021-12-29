/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package notify

import (
	"encoding/json"
	"testing"

	. "github.com/tevid/gohamcrest"
	"mango/third_party/agollo/cluster/roundrobin"
	_ "mango/third_party/agollo/cluster/roundrobin"
	"mango/third_party/agollo/env"
	"mango/third_party/agollo/env/config"
	jsonConfig "mango/third_party/agollo/env/config/json"
	_ "mango/third_party/agollo/env/file/json"
	jsonFile "mango/third_party/agollo/env/file/json"
	"mango/third_party/agollo/extension"
)

func init() {
	extension.SetLoadBalance(&roundrobin.RoundRobin{})
	extension.SetFileHandler(&jsonFile.FileHandler{})
}

var (
	jsonConfigFile = &jsonConfig.ConfigFile{}
)

func initNotifications() *config.AppConfig {
	appConfig := env.InitFileConfig()
	appConfig.NamespaceName = "application,abc1"
	appConfig.Init()
	return appConfig
}

func TestUpdateAllNotifications(t *testing.T) {
	//clear
	c := initNotifications()

	notifyJson := `[
  {
    "namespaceName": "application",
    "notificationId": 101
  }
]`
	notifies := make([]*config.Notification, 0)

	err := json.Unmarshal([]byte(notifyJson), &notifies)

	Assert(t, err, NilVal())
	Assert(t, true, Equal(len(notifies) > 0))

	c.GetNotificationsMap().UpdateAllNotifications(notifies)

	Assert(t, true, Equal(c.GetNotificationsMap().GetNotifyLen() > 0))
	Assert(t, int64(101), Equal(c.GetNotificationsMap().GetNotify("application")))
}

func TestUpdateAllNotificationsError(t *testing.T) {
	//clear
	appConfig := initNotifications()

	notifyJson := `ffffff`
	notifies := make([]*config.Notification, 0)

	err := json.Unmarshal([]byte(notifyJson), &notifies)

	Assert(t, err, NotNilVal())
	Assert(t, true, Equal(len(notifies) == 0))

	appConfig.GetNotificationsMap().UpdateAllNotifications(notifies)

	Assert(t, appConfig.GetNotificationsMap().GetNotifyLen(), Equal(2))
}

func getTestAppConfig() *config.AppConfig {
	jsonStr := `{
    "appId": "test",
    "cluster": "dev",
    "namespaceName": "application",
    "ip": "localhost:8888",
    "releaseKey": "1"
	}`
	c, _ := env.Unmarshal([]byte(jsonStr))

	appConfig := c.(*config.AppConfig)
	appConfig.Init()
	return appConfig
}
