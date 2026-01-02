// Copyright 2025 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package setting

// UnityCoco settings for AI-powered semantic search
var UnityCoco = struct {
	Enabled bool
	WebURL  string // URL of Unity Coco Web service (e.g., http://localhost:3001)
}{
	Enabled: false,
	WebURL:  "http://localhost:3001",
}

func loadUnityCocoFrom(rootCfg ConfigProvider) {
	sec := rootCfg.Section("unity_coco")
	UnityCoco.Enabled = sec.Key("ENABLED").MustBool(false)
	UnityCoco.WebURL = sec.Key("WEB_URL").MustString("http://localhost:3001")
}
