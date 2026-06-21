# 🐱🦞Manboster: 你的曼波虾头小助手！

[English Version](README.md)

> Manboster = Manbo(曼波) + Lobster

我们取了龙虾OpenClaw和龙虾增强版IronClaw的精华，做了一个更加安全的Lobster，并把它取名为Manboster。

## Meet Your Personal Lobster

At the beginning, it is a personal AI assistant which is able to chat with you with ease. However, equipped with wasm-based plugins, you can use it to do anything you want with security guaranteed!

Since this project is work in progress and only shows as a MVP now, there are a few chat options(Telegram and Feishu/Lark as a second supported provider in next versions) and LLM options(openrouter, kimi, DeepSeek and openai-compatible APIs)

We will sincerely appreciate if you contribute your own codes!

另外，Manboster的主要代码和架构都是由人完成的，剩下的部分实用函数补充和繁杂任务，以及让AI提供了部分架构建议，人最后拍板是否采纳。

## Why Manboster?

1. 基于Golang，单可执行文件，开箱即用。
2. 得益于Golang的语言特性，它占用内存较小，运行速度快，内置多线程，卡不死。
3. Introducing Hachimi, a guard model running normally locally, will evaluate LLM's tool call if you need. If Hachimi thinks some action is unsafe, it will stop and let you decide.
4. Default zero-trust design, gatekeeper system and TTL settings to handle all tool calls.
5. 可插拔可配置的内置工具，你完全可以用 `manboster config` 启用或禁用一个工具。
6. 一个内置的搜索工具，它内置了基于 `go-rod` 的无头浏览器，你可以用它做搜索，当然，你也可以用自己的搜索API Key来搜索。 Use it to search for the Internet or give your search API's keys
7. [开发中功能] 兼容旧的 OpenClaw skills，只需要输入 `manboster skills install SKILLS.md` 或 `manboster skills install SKILLS.zip` 即可安装！再也不用担心老的OpenClaw技能不能在manboster里面不能用啦！
8. [Work in Progress] A built-in vault tool helps you store your sensitive data using industry best practices while balancing your experience. LLM NEVER has access to you credentials.
9. [开发中功能] 内置的脚本运行器沙盒工具，支持大模型在 Wasm 沙盒中运行 JavaScript 或 Python 脚本。
10. [计划中功能] 可插拔的 RAG（检索增强生成）记忆系统及对 mem0 理论的适配。
11. [计划中功能] MCP协议支持，只需要配置连接MCP服务器，就可以让你的Manboster轻松接入其他工具！
12. [计划中功能] 模拟 UI/输入交互、屏幕截图，并提供大量内置 SDK 供你构建插件时使用。
13. [计划中功能] 基于 Wasm 和 Extism 的插件系统，它不仅轻量，还能防止恶意插件破坏你的机器。
14. [计划中功能] MamboHub，一个让你可以轻松下载并使用技能与插件的分发中心。我们当然也会支持类似于 ClawHub 的 skills 安装 Also, we will compat ClawHub and more skill hubs.
15. [计划中功能] 你可以使用 .manboplugin 文件安装插件，或者是使用 manbodev 辅助工具来开发和编写自己的插件！

## 快速上手

1. Download binary files built in releases and open it via double-click or terminal shell `./manboster`.
2. You're required to configure your manboster when it runs in the first time, just configure it.
3. Enjoy yourselves!

If you downloaded the Go programming language development kit? go install github.com/manboster/manboster/cmd/manboster@latest

更多详见 [快速开始文档](https://manboster.dev/zh-cn/docs/quickstart.html)

**Notes for the Daemon**: Simply double-click would not start the daemon, if you want to create and start the daemon, please run `manboster start`.

## Make a contribution

Manboster is now open to the community and we are looking forward to your contributions! Read [CONTRIBUTING.md](./CONTRIBUTING.md) to start!

## 遇到了问题？

如果你遇到了问题，请先阅读[文档](https://manboster.dev/zh-cn/docs/)。

Create a new issue after trying ways that documentation says or what you've faced is not recorded in documentation.

Before creating a new issue, search for the problem. If there is none, create it after reading [How To Ask Questions The Smart Way](http://www.catb.org/~esr/faqs/smart-questions.html).

## 授权协议

本软件使用Apache 2.0协议作为开源协议。

## 致谢

请阅读 [THANKING.md](./THANKING.md)
