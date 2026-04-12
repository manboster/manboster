# 🐱🦞Manboster: 你的曼波虾头小助手！ 

[English Version](README.md)

> Manboster = Manbo(曼波) + Lobster

我们取了龙虾OpenClaw和龙虾增强版IronClaw的精华，修复了其安全性差的缺点，做成了现在你所看到的Manboster。

在刚开始，它可能只是一个普通的ai聊天软件。但是，在用技能和插件把它武装后，它可以成为你的可靠助手！

虽然现在只有telegram接入和openai兼容api的大语言模型接入，但是我们仍然在给项目添加新功能！欢迎来为社区做出自己的一份贡献！

## 特色功能

1. 基于Golang，单可执行文件，开箱即用。
2. 运行速度快，内置多线程，不会聊天卡死。
3. 主控端执行无论是插件还是技能还是命令时，本地运行的llm小模型(哈基米)都会对执行的命令进行打分，分数高需要用户明确授权。
4. 同时在兼容openclaw提供的md的技能的情况下，我们推出了基于wasm和extism的插件配置。wasm的特性可以在快速启动时还能保证插件的执行几乎不会对你的电脑做出任何破坏。
5. 内置基于extism的控制电脑的sdk可供插件沙盒调用，比如说屏幕截图，模拟点击，互联网搜索等等。
6. 互联网搜索可选api搜索或无头浏览器搜索。
7. manbohub可以直接安装和配置技能和插件。你可以直接安装 `.manboskill` 或 `.manboplugin` ，也可以来贡献自己写的技能和插件！

## 开发进度

- [ ] 聊天api
- [ ] SDK Agent API

## 快速上手

去releases下载匹配自己系统的可执行文件，直接打开即可。

第一次开启时，它会要求你配置信息，根据要求配置就好了。

更多详见命令行帮助文档： `manboster help`

直接启动不会开启守护进程，输入以下内容启动守护进程。

```shell
manboster start
```

## 加入我们！

欢迎来为社区做出自己的一份贡献！

在写代码，提出PR前，请先阅读[CONTRIBUTING.md](./CONTRIBUTING.md)来了解项目的贡献要求

## 遇到了问题？

如果你遇到了问题，请先阅读[文档](https://manboster.dev/docs/)。

文档里面没有给出解决方案或者是解决方案无法解决问题时再去issue搜索这个问题，如果还没有，请开issue。

在开issue之前，请先充分阅读并理解[提问的艺术](https://github.com/ryanhanwu/How-To-Ask-Questions-The-Smart-Way/blob/main/README-zh_CN.md)以给你我减少不必要的时间，冲突和麻烦。

## 授权协议

本软件使用Apache 2.0协议作为开源协议。
