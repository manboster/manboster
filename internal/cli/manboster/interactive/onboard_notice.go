package interactive

import (
	"context"

	"github.com/manboster/manboster/internal/cli/helper"
)

// OnboardWarningForm provides a warning notice
func OnboardWarningForm(ctx context.Context) error {
	return helper.ConfirmForm(ctx, `
# RISK DISCLOSURE & DISCLAIMER
**PLEASE READ THESE WORDS CAREFULLY:**

Manboster is an AI agent able to chat and control your computers like OpenClaw and IronClaw and currently in MVP stage. By proceeding, you acknowledge:
1. WIP means this project is **Work in Progress**, and **it is expected to encounter bugs, crashes, and breaking changes.**
2. If you run 'manboster start', you open a daemon running in your computer. **The background process has persistent resource access to your computer.**
3. WASM sandboxing plugins is strong, but **3rd-party code still carries risks**.
4. **Hachimi scoring reduces decision fatigue, but cannot fully prevent advanced prompt injections or unsafe LLM behaviors.**
5. **Granting access enables data transmission to LLMs and allows device control. We are not liable for any issues arising from these interactions.**
6. This software is provided "AS IS" under Apache 2.0. **You are strictly prohibited from using this application for any criminal or illegal purposes. We disclaim all liability and responsibility for any unlawful activities conducted using this software.**
`, "Do you understand the risks and wish to proceed?", "I Understand & Continue")
}

func OnboardVersionWarningForm(ctx context.Context) error {

	return helper.ConfirmForm(ctx, `
# UNSTABLE VERSION WARNING
**PLEASE READ THESE WORDS CAREFULLY:**
It seems that you're going to use an unstable version of Manboster. Please note that:
1. It's normal to encounter bugs, crashes, and breaking changes in unstable versions.
2. As this is not a stable version, it's not contain ANY security patches and fixes.
3. This version's configuration may be incompatible with older versions and please aware the configuration changes.
4. If you encounter bugs, we appreciate you to commit to issues and we will fix it as soon as possible.
5. PLEASE DO NOT STORE ANY SENSITIVE AND IMPORTANT DATA IN THIS VERSION! As it's unstable and we are unsure that this application will work as is.
`, "Do you understand the risks and wish to proceed?", "I Understand & Continue")

}
