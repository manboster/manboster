# Manboster Contributing Guideline

*Rev 3; Apr 23, 2026*

## TL;DR

This guideline explains how to contribute to Manboster. You **MUST** read this before writing or refactoring any code.

When contributing to this project (via `git push` or Pull Request), these guidelines apply, and you are held responsible for the code you commit.

If you are a first-time contributor, please look for the `good-first-issue` label in the opened issues.

## 1. About the AI code

1.1 We **DO NOT** accept **ANY** AI Agent's independent contribution, unless the pull request is opened by its human contributor and is responsible for the code.

1.2 We accept AI generated codes. But the creator of the pull request is **FULLY RESPONSIBLE** for the code AI generated.

1.3 AI generated code **SHOULD NOT** be used to fix or refactor `good-first-issue`, as these are for human beginner developers.

## 2. About the Versioning

2.1 The Versioning follows the principle of [SemVer](https://semver.org/).

2.2 Version `0.0.1` represents minor bug fixes.

2.3 Version `0.1.0` represents new features.

2.4 Before a `0.1.0` release, Beta and Release Candidate (RC) versions should be released, tagged as `0.x.0-beta` and `0.x.0-rc`.

2.5 It's **NOT ALLOWED** to add new features in RC versions.

2.6 Core contributors may release RC, Stable, and Beta versions for `0.x.0`.

2.7 Version `1.0.0` represents a significant leap in features and architecture. The configuration version shall be upgraded accordingly.

2.8 Before a `1.0.0` release, Alpha, Beta, and Release Candidate (RC) versions should be released, tagged as `x.0.0-alpha`, `x.0.0-beta`, and `x.0.0-rc`.

2.9 Core contributors may release RC and Beta versions for `1.0.0`. Only the owner can release stable `1.0.0` versions.

2.10 Core contributors shall create new git branches for 0.1.0 and 1.0.0 versions (e.g., `dev-v0.x`, `dev-v1.x`) and merge bug fixes from older versions.

## 3. About the commit 

3.1 Commit messages must follow [the Conventional Commits](https://www.conventionalcommits.org/) standard. 

## 4. About the features

4.1 Features should be discussed and documented as [issues](https://github.com/manboster/manboster/issues).

4.2 Please open an issue of feature before you open your pull request.

4.3 New features that introduce breaking changes or modify the core architecture **MUST BE** approved by the owner.

4.4 Other features, such as new chat or LLM providers, should be approved by at least one core contributor.

## 5. About the style

5.1 You should run `go fmt` before you commit your code.

5.2 Please follow standard Go naming conventions: use `MixedCaps` (PascalCase) for exported identifiers and `camelCase` for unexported variables.

5.3 Please use human-friendly function, variable, type names. Meaningless or ambiguous words **ARE NOT** allowed in any identifiers.

5.4 Comments in code **MUST BE** English, not other languages.

## 6. About the testing

6.1 If you contribute to the `engine` package, you **MUST** provide corresponding unit tests for your changes.

6.2 Before you commit, you should run `go test ./...` to ensure all tests are passed.

## 7. About the Contribution Roles

7.1 A Contributor is any individual who contributes to the project, ranging from minor typo fixes to significant bug resolutions.

7.2 A Core Contributor is dedicated member who actively participates in technical discussions, provides roadmap suggestions, and maintains the codebase consistently.

7.3 The Owner is the creator and primary maintainer of the Manboster Project.

7.4 To be nominated as a Core Contributor, you must:

7.4.1 Propose and implement at least one meaningful feature. A meaningful feature means the feature should align with this project's roadmap and approved by the owner.

7.4.2 Demonstrate frequent and consistent contributions to the project.

7.4.3 Or, provide exceptional help in building documentation and managing the community.

7.5 Once you become a Core Contributor, you will have privileges below:

7.5.1 Your name will be featured in the Core Contributors section of THANKS.md.

7.5.2 You will be invited to the Manboster GitHub Organization and our internal developer group.

7.5.3 You will be granted extended permissions (e.g., branch management, issue labeling, and PR reviews).

7.6 Core Contributors who have not contributed code or participated in maintenance for more than 6 months will have their repository write permissions suspended for security reasons. If you wish to resume active contribution, please contact the owner to restore your permissions.

7.7 All contributors **MUST** adhere to fundamental professional ethics and the open-source code of conduct. Any behavior that intentionally harms the project, its users, or the community is strictly prohibited.

7.8 While you are welcome to use Manboster for academic projects, hackathons, or competitions, you **MUST** comply with our open-source license and provide proper attribution. Plagiarism or claiming Manboster's code as your own independent work is **STRICTLY PROHIBITED**. We reserve the right to publicly disavow such actions.

7.9 We have a **ZERO-TOLERANCE** policy for malicious contributions. Intentionally introducing backdoors, malware, spyware, or obfuscated malicious code into the repository will result in an **IMMEDIATE AND PERMANENT BAN** from the project. Furthermore, we reserve the right to report such malicious activities to GitHub Trust & Safety and relevant cybersecurity authorities.

## 8. About the Issues and Discussions

8.1 Issues is the place where people discuss features and bug fixes. 

8.2 Discussions is our community space for casual chat, questions, and networking. All participants must adhere to [GitHub's Terms of Service](https://docs.github.com/en/site-policy/github-terms/github-terms-of-service).

8.3 We **DO NOT** welcome spam, including "check-ins" (such as sending "留名" "打卡" and more) or meaningless comments in Issues. Such content will be closed or deleted **WITHOUT NOTICE**. While we are more lenient in Discussions, please keep the content constructive.

8.4 If you are a security researcher and have found a vulnerability in Manboster, please **DO NOT** open a public issue. Use the GitHub Security Advisory feature or send a detailed report to `security@manboster.dev`. Once verified, we will coordinate a fix as a priority.

## 9. About Updating this Guideline

9.1 We welcome community members to propose improvements to this guideline. If you see a need for an edit or update, please first open an issue to initiate a discussion.

9.2 Once discussed, you may submit a Pull Request. The proposed draft must be reviewed by the Core Contributors and the Owner.

9.3 Guideline updates are subject to a voting process. A proposed change will be accepted if it receives approval from over 50% of the Core Contributors or over 70% of all Contributors.

9.4 The Owner must be notified of any proposed changes and retains the ultimate right of veto to reject the edit, regardless of the voting outcome.

9.5 Any accepted updates must be properly recorded via Git. The author of the Pull Request **MUST** update the revision version (`Rev`) at the top of the file and append a brief description of the changes to the `Changelogs` section at the bottom.

9.6 This section is also in effect with modifying `SECURITY.md`.

## Changelogs

This guideline may be updated as needed. Please review it frequently to stay informed of changes.

Rev 3: Update morals and updates of this guideline, Updated on 2026.4.23

Rev 2.2: Fix typo, Updated on 2026.4.13

Rev 2.1: Fix ambiguous words may occur misunderstanding, Updated on 2026.4.13

Rev 2: Add contribution roles and issues, discussions, Updated on 2026.4.13

Rev 1.1, 1.2: Minor Typo fixes, Updated on 2026.4.13

Rev 1: CONTRIBUTING.md created, Updated on 2026.4.13