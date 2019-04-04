# Go 1.11 Modules

Go 1.11 includes preliminary support for versioned modules as proposed [here](https://golang.org/design/24301-versioned-go). Modules are an experimental opt-in feature in Go 1.11, with the [plan](https://blog.golang.org/modules2019) of incorporating feedback and finalizing the feature for Go 1.13. Even though some details may change, future releases will support modules defined using Go 1.11 or 1.12.

The initial prototype (`vgo`) was [announced](https://research.swtch.com/vgo) in February 2018. In July 2018, support for versioned modules [landed](https://groups.google.com/d/msg/golang-dev/a5PqQuBljF4/61QK4JdtBgAJ) in the main repository. Go 1.11 was released in August 2018.

Please provide feedback on modules via [existing or new issues](https://github.com/golang/go/wiki/Modules#github-issues) and via [experience reports](https://github.com/golang/go/wiki/ExperienceReports).

## Table of Contents

The "Quick Start" and "New Concepts" sections are particularly important for someone who is starting to work with modules. The "How to..." sections cover more details on mechanics. The largest quantity of content on this page is in the FAQs answering more specific questions; it can be worthwhile to at least skim the FAQ one-liners listed here.

* [Quick Start](https://github.com/golang/go/wiki/Modules#quick-start)
   * [Example](https://github.com/golang/go/wiki/Modules#example)
   * [Daily Workflow](https://github.com/golang/go/wiki/Modules#daily-workflow)
* [New Concepts](https://github.com/golang/go/wiki/Modules#new-concepts)
   * [Modules](https://github.com/golang/go/wiki/Modules#modules)
   * [go.mod](https://github.com/golang/go/wiki/Modules#gomod)
   * [Version Selection](https://github.com/golang/go/wiki/Modules#version-selection)
   * [Semantic Import Versioning](https://github.com/golang/go/wiki/Modules#semantic-import-versioning)
* [How to Use Modules](https://github.com/golang/go/wiki/Modules#how-to-use-modules)
   * [How to Install and Activate Module Support](https://github.com/golang/go/wiki/Modules#how-to-install-and-activate-module-support)
   * [How to Define a Module](https://github.com/golang/go/wiki/Modules#how-to-define-a-module)
   * [How to Upgrade and Downgrade Dependencies](https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies)
   * [How to Prepare for a Release (All Versions)](https://github.com/golang/go/wiki/Modules#how-to-prepare-for-a-release)
   * [How to Prepare for a Release (v2 or Higher)](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher)
   * [Publishing a Release](https://github.com/golang/go/wiki/Modules#publishing-a-release)
* [Migrating to Modules](https://github.com/golang/go/wiki/Modules#migrating-to-modules)
* [Additional Resources](https://github.com/golang/go/wiki/Modules#additional-resources)
* [Changes Since the Initial Vgo Proposal](https://github.com/golang/go/wiki/Modules#changes-since-the-initial-vgo-proposal)
* [GitHub Issues](https://github.com/golang/go/wiki/Modules#github-issues)
* [FAQs](https://github.com/golang/go/wiki/Modules#faqs)
  * [How are versions marked as incompatible?](https://github.com/golang/go/wiki/Modules#how-are-versions-marked-as-incompatible)
  * [Can two modules depend on each other?](https://github.com/golang/go/wiki/Modules#can-two-modules-depend-on-each-other-cyclical-import)
  * [Can a module depend on a different version of itself?](https://github.com/golang/go/wiki/Modules#can-a-module-depend-on-a-different-version-of-itself)
  * [When do I get old behavior vs. new module-based behavior?](https://github.com/golang/go/wiki/Modules#when-do-i-get-old-behavior-vs-new-module-based-behavior)
  * [Why does installing a tool via 'go get' fail with error 'cannot find main module'?](https://github.com/golang/go/wiki/Modules#why-does-installing-a-tool-via-go-get-fail-with-error-cannot-find-main-module)
  * [How can I track tool dependencies for a module?](https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module)
  * [What is the status of module support in IDEs, editors and standard tools like goimports, gorename, etc.?](https://github.com/golang/go/wiki/Modules#what-is-the-status-of-module-support-in-ides-editors-and-standard-tools-like-goimports-gorename-etc)
* [FAQs — Additional Control](https://github.com/golang/go/wiki/Modules#faqs--additional-control)
  * [What community tooling exists for working with modules?](https://github.com/golang/go/wiki/Modules#what-community-tooling-exists-for-working-with-modules)
  * [When should I use the 'replace' directive?](https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive)
  * [Can I work entirely outside of VCS on my local filesystem?](https://github.com/golang/go/wiki/Modules#can-i-work-entirely-outside-of-vcs-on-my-local-filesystem)
  * [How do I use vendoring with modules? Is vendoring going away?](https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away)
  * [Are there "always on" module repositories and enterprise proxies?](https://github.com/golang/go/wiki/Modules#are-there-always-on-module-repositories-and-enterprise-proxies)
  * [Can I control when go.mod gets updated and when the go tools use the network to satisfy dependencies?](https://github.com/golang/go/wiki/Modules#can-i-control-when-gomod-gets-updated-and-when-the-go-tools-use-the-network-to-satisfy-dependencies)
  * [How do I use modules with CI systems such as Travis or CircleCI?](https://github.com/golang/go/wiki/Modules#how-do-i-use-modules-with-ci-systems-such-as-travis-or-circleci)
* [FAQs — go.mod and go.sum](https://github.com/golang/go/wiki/Modules#faqs--gomod-and-gosum)
  * [Why does 'go mod tidy' record indirect and test dependencies in my 'go.mod'?](https://github.com/golang/go/wiki/Modules#why-does-go-mod-tidy-record-indirect-and-test-dependencies-in-my-gomod)
  * [Is 'go.sum' a lock file? Why does 'go.sum' include information for module versions I am no longer using?](https://github.com/golang/go/wiki/Modules#is-gosum-a-lock-file-why-does-gosum-include-information-for-module-versions-i-am-no-longer-using)
  * [Should I still add a 'go.mod' file if I do not have any dependencies?](https://github.com/golang/go/wiki/Modules#should-i-still-add-a-gomod-file-if-i-do-not-have-any-dependencies)
  * [Should I commit my 'go.sum' file as well as my 'go.mod' file?](https://github.com/golang/go/wiki/Modules#should-i-commit-my-gosum-file-as-well-as-my-gomod-file)
* [FAQs — Semantic Import Versioning](https://github.com/golang/go/wiki/Modules#faqs--semantic-import-versioning)
  * [Why must major version numbers appear in import paths?](https://github.com/golang/go/wiki/Modules#why-must-major-version-numbers-appear-in-import-paths)
  * [Why are major versions v0, v1 omitted from import paths?](https://github.com/golang/go/wiki/Modules#why-are-major-versions-v0-v1-omitted-from-import-paths)
  * [What are some implications of tagging my project with major version v0, v1, or making breaking changes with v2+?](https://github.com/golang/go/wiki/Modules#what-are-some-implications-of-tagging-my-project-with-major-version-v0-v1-or-making-breaking-changes-with-v2)
  * [Can a module consume a package that has not opted in to modules?](https://github.com/golang/go/wiki/Modules#can-a-module-consume-a-package-that-has-not-opted-in-to-modules)
  * [Can a module consume a v2+ package that has not opted into modules? What does '+incompatible' mean?](https://github.com/golang/go/wiki/Modules#can-a-module-consume-a-v2-package-that-has-not-opted-into-modules-what-does-incompatible-mean)
  * [How are v2+ modules treated in a build if modules support is not enabled? How does "minimal module compatibility" work in 1.9.7+, 1.10.3+, and 1.11?](https://github.com/golang/go/wiki/Modules#how-are-v2-modules-treated-in-a-build-if-modules-support-is-not-enabled-how-does-minimal-module-compatibility-work-in-197-1103-and-111)
* [FAQs — Multi-Module Repositories](https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories)
  * [What are multi-module repositories?](https://github.com/golang/go/wiki/Modules#what-are-multi-module-repositories)
  * [Is it possible to add a module to a multi-module repository?](https://github.com/golang/go/wiki/Modules#is-it-possible-to-add-a-module-to-a-multi-module-repository)
  * [Is it possible to remove a module from a multi-module repository?](https://github.com/golang/go/wiki/Modules#is-it-possible-to-remove-a-module-from-a-multi-module-repository)
  * [Can a module depend on an internal/ in another?](https://github.com/golang/go/wiki/Modules#can-a-module-depend-on-an-internal-in-another)
* [FAQs — Minimal Version Selection](https://github.com/golang/go/wiki/Modules#faqs--minimal-version-selection)
  * [Won't minimal version selection keep developers from getting important updates?](https://github.com/golang/go/wiki/Modules#wont-minimal-version-selection-keep-developers-from-getting-important-updates)
* [FAQs — Possible Problems](https://github.com/golang/go/wiki/Modules#faqs-possible-problems)
  * [What are some general things I can spot check if I am seeing a problem?](https://github.com/golang/go/wiki/Modules#what-are-some-general-things-i-can-spot-check-if-i-am-seeing-a-problem)
  * [What can I check if I am not seeing the expected version of a dependency?](https://github.com/golang/go/wiki/Modules#what-can-i-check-if-i-am-not-seeing-the-expected-version-of-a-dependency)
  * [Why am I getting an error 'cannot find module providing package foo'?](https://github.com/golang/go/wiki/Modules#why-am-i-getting-an-error-cannot-find-module-providing-package-foo)
  * [Why does 'go mod init' give the error 'cannot determine module path for source directory'?](https://github.com/golang/go/wiki/Modules#why-does-go-mod-init-give-the-error-cannot-determine-module-path-for-source-directory)
  * [I have a problem with a complex dependency that has not opted in to modules. Can I use information from its current dependency manager?](https://github.com/golang/go/wiki/Modules#i-have-a-problem-with-a-complex-dependency-that-has-not-opted-in-to-modules-can-i-use-information-from-its-current-dependency-manager)
  * [Why does 'go build' require gcc, and why are prebuilt packages such as net/http not used?](https://github.com/golang/go/wiki/Modules#why-does-go-build-require-gcc-and-why-are-prebuilt-packages-such-as-nethttp-not-used)
  * [Do modules work with relative imports like `import "./subdir"`?](https://github.com/golang/go/wiki/Modules#do-modules-work-with-relative-imports-like-import-subdir)
  * [Some needed files may not be present in populated vendor directory](https://github.com/golang/go/wiki/Modules#some-needed-files-may-not-be-present-in-populated-vendor-directory)

## Quick Start

#### Example

The details are covered in the remainder of this page, but here is a simple example of creating a module from scratch.

Create a directory outside of your GOPATH:
```
$ mkdir -p /tmp/scratchpad/hello
$ cd /tmp/scratchpad/hello
```

Initialize a new module:
```
$ go mod init github.com/you/hello

go: creating new go.mod: module github.com/you/hello
```

Write your code:
```
$ cat <<EOF > hello.go
package main

import (
    "fmt"
    "rsc.io/quote"
)

func main() {
    fmt.Println(quote.Hello())
}
EOF
```

Build and run:
```
$ go build
$ ./hello

Hello, world.
```

The `go.mod` file was updated to include explicit versions for your dependencies, where `v1.5.2` here is a [semver](https://semver.org) tag:
```
$ cat go.mod

module github.com/you/hello

require rsc.io/quote v1.5.2
```

#### 日常流程

注意：上面的例子中，`go get` 不是必需的。

典型的工作流程是：

* import你依赖的包到`.go`文件。
* 标准命令像`go build` or `go test`将自动添加你import所需要的依赖（更新`go.mod`文件并下载这些依赖）
* 可以像这样`go get foo@v1.2.3`, `go get foo@master`, `go get foo@e3702bed2`选择添加依赖项的版本，或者直接将依赖项及其版本编辑到`go.mod`文件中

可能常用到的其他功能:

* `go list -m all` — 查看将被用于构建的，所有直接或者间接引用的依赖项的最终版本
* `go list -u -m all` — 查看所有直接或者间接引用的依赖项的可用可更新的minor和patch版本
* `go get -u` or `go get -u=patch` — 更新所有直接或者间接引用的依赖项到最新的minor或patch版本
* `go build ./...` or `go test ./...` — 当从模块根目录运行命令时，将会build和test模块中的所有包
* `go mod tidy` — 分析`.go`文件，添加或删除所有（使用的/未使用的）直接或者间接引用的依赖项到`go.mod`文件中
* `replace` — 用fork、本地源文件的版本替换所依赖的版本

在阅读接下来的四个章节的“新概念”后，你将有足够的信息开始并创建大多数基于modules的项目。

## 新概念

这个部分介绍了新的主要概念和高级介绍

### Modules

**模块**是作为单个单元一起进行版本化的相关go包的集合。记录了完整的依赖列表便于重复构建。

通常，一个版本控制库都包含一个明确模块在其根目录下（单模块库）。当然一个库中也可以有多个模块（多模块库），相比较而言，多模块库会多做更多的工作，另行介绍。

总结一下repositories, modules, 和 packages 之间的关系:
* 一个 repository 可以包含一个或者多个Go modules.
* 每一个 module 可以包含一个或者多个Go packages.
* 每一个包可以包含一个或者多`.go`文件在一个单一的目录下.

Modules 必须参照 [semver](https://semver.org/) 进行语义版本化，通常形式为`v(major).(minor).(patch)`，例如：`v0.1.0`, `v1.2.3`, or `v1.5.0-rc.1`，前缀`v`是必需的。
如果使用Git，发布的[tag](https://git-scm.com/book/en/v2/Git-Basics-Tagging)必须随着版本一起提交。公共或者私有的模块库以及代理都将变得可用。

### go.mod

一个树形结构的源文件目录及其根目录下`go.mod`文件构成了一个模块。模块源文件可能位于`GOPATH`之外。他包含4个指令：`module`, `require`, `replace`, `exclude`

下面是`github.com/my/thing`模块的`go.mod`文件示例：

```
module github.com/my/thing

require (
    github.com/some/dependency v1.2.3
    github.com/another/dependency/v4 v4.0.0
)
```

前面说到，`go.mod`文件包含了4个指令，其中`module`指令定义了模块的唯一标识，这就规定了**module path**。模块中所有包的import路径都共享这一**module path**作为访问他们的公共前缀。**module path**和从`go.mod`文件路径到包的相对路径一起，决定了这些包的**import path**。

例如，如果要为存储库github.com/my/repo创建一个Module，该Module的package中包含两个import path为github.com/my/repo/foo和github.com/my/repo/bar的包，则go.mod文件中的第一行通常会将**module path**声明为module github.com/my/repo，相应的目录结构可以是：

```
repo/
├── go.mod
├── bar
│   └── bar.go
└── foo
    └── foo.go
```

在一个`.go`的源代码中，包是通过包含了**module path**的完整路径通过import导入的。例如，一个模块在`go.mod`中声明了他的唯一标识`module example.com/my/module`作为**module path**，那么使用起来如下：

```
import "example.com/my/module/mypkg"
```

表示从`example.com/my/module` import `mypkg`包。

`exclude` 和 `replace`指令仅在当前（“主”）模块上有效。在构建主模块时，主模块以外的模块（依赖模块）中的`exclude` 和 `replace`指令将被忽略。因此，replace和exclude语句允许主模块完全控制其自身的构建，而不受依赖项的影响。

### Version Selection

如果你新增一个新的import到你的源代码中并且还没有作为一个`require`到`go.mod`，大多数的go命令如`go build` 和 `go test`将自动寻找模块的*highest*版本，并作为直接依赖，通过`require`指令将其添加到`go.mod`中。

例如，如果你新import的依赖项为M，并且M的`latest tagged release version` 为 `v1.2.3`，上述中那些go命令执行后，你的`go.mod`文件将会以`require M v1.2.3`结尾，这意味着模块M是一个依赖项，并且允许的版本是 `>= v1.2.3` （且小于`v2`，假设`v2`与`v1`不兼容）

最小版本选择算法用于选择构建中使用的所有模块的版本。对于构建中的每个模块，最小版本选择总是由主模块或其依赖项中的require指令显式列出的语义最高的版本作为依据。
例如，如果您的模块依赖于模块A，而模块A又依赖D的v1.0.0，并且您的模块也依赖于模块B，而模块B却D的v1.1.1，那么在构建中的最小版本选择将会是D的1.1.1版本（假定D的v1.1.1是D目前的最高版本），使得D的版本在构建中选择保持一致，即使以后有v1.2.0的D版本可用。
这个示例告诉我们模块系统是如何100%保证构建的可复用性的。准备好后，模块作者或用户可以选择升级到D的最新可用版本，或者选择D的显式版本。

有关最小版本选择算法的基本原理和概述，请参阅官方提案的“[High Fidelity Builds](https://github.com/golang/proposal/blob/master/design/24301-versioned-go.md#update-timing--high-fidelity-builds) ”部分，或参阅更详细的[vgo博客系列](https://research.swtch.com/vgo)。

使用`go list -m all`查看模块选择的版本列表（包括间接依赖）

另请参阅下面的“[如何升级和降级依赖项](https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies)”部分以及下面“[版本如何标记为不兼容？](https://github.com/golang/go/wiki/Modules#how-are-versions-marked-as-incompatible)”的常见问答。

### <span id="Semantic_Import_Versioning">语义导入版本控制</span>

多年来，官方的Go常见问题解答都包含了关于包版本控制的建议：

> "Packages intended for public use should try to maintain backwards compatibility as they evolve. The Go 1 compatibility guidelines are a good reference here: don't remove exported names, encourage tagged composite literals, and so on. If different functionality is required, add a new name instead of changing an old one. If a complete break is required, create a new package with a new import path."

注：
> “面向公共用途的软件包应该在发展过程中尽量保持向后兼容性。Go1兼容性准则在这里是一个很好的参考：不要删除导出的名称，鼓励标记的复合字面量，等等。如果需要不同的功能，请添加新名称，而不是更改旧名称。如果需要彻底打破兼容性，请使用新的导入路径创建一个新包。“

最后一句话特别重要-如果您破坏了兼容性，您应该更改包的导入路径。伴随着Go 1.11 对于模块的支持，该建议将正式化被纳入导入的兼容性原则：

> "If an old package and a new package have the same import path,
> the new package must be backwards compatible with the old package."

> "如果一个旧的包和一个新的包，他们拥有相同的导入路径,那么这个新包必须向后兼容旧的包。"

当v1或更高版本的包做了向后不兼容的变更时，我们应该想到遵循Semver语义版本化原则对不兼容的包做**major**版本的更改。遵循导入兼容性和语义导入版本控制的原则，将**major**版本号包含在`import path`中-这可确保`import path`，在由于**major**版本增加致兼容性被破坏时，能随时变化。

反模式：

如果模块依赖模块A v1.1.0，那么import path为 `import github.com/A/pkg`，假设A v2.0.0不兼容v1.1.0，而v2.0.0还有相同的import path的话，后续在构建时很可能就会出现问题。

较好的方式是如果版本不兼容，v2.0.0最好是新给一个`module path`，如`module github/A/v2`，即将**major**版本号放到**module path**的末尾，导入代码为`import github/A/v2/pkg`，由于import path不同，能很好的避免因兼容性问题而带来的冲突。

作为导入版本控制的结果，选择使用Go modules **必须遵循以下原则**：

* 遵循[semver](https://semver.org/) 原则（使用例如`v1.2.3`这样的版本提交，即`v major.minor.patch`这样的格式）
* 如果一个模块的版本是v2或者更高（v1不需要），那么模块的主要版本 ***必须*** 包含一个`/vN`在`go.mod`文件内的模块路径末尾(例如, `module github.com/my/mod/v2`, `require github.com/my/mod/v2 v2.0.0`) 并且在包内import (例如, `import "github.com/my/mod/v2/mypkg"`)这样的路径。
* 如果模块版本为v0或者v1，则不需要将**major**版本号包含在`module path`中。

通常，具有不同导入路径的包是不同的包。例如，`Math/Rand`与`Crypto/Rand`是不同的包。同样，`example.com/my/mod/mypkg`与`example.com/my/mod/v2/mypkg`也是不同的包，两者都可以在单个构建中导入，这不仅有助于解决依赖性冲突性问题，还允许根据其v2替换来实现v1模块，反之亦然。

有关语义导入版本控制的详细信息，请参阅go命令文档的“["Module compatibility and semantic versioning"](https://golang.org/cmd/go/#hdr-Module_compatibility_and_semantic_versioning) ”部分，有关语义版本控制的详细信息，请参阅https://semver.org。

到目前为止，本节的重点是选择模块并导入其他模块的代码。但是，将主要版本放在v2+模块的导入路径中可能会导致与旧版本的go或尚未选择模块的代码不兼容。为了解决这一点，有三个重要的过渡性例外情况与上述规则的并存。当然，随着越来越多的包选择加入模块，这些过渡的例外将变得越来越不那么重要。

**三种过渡性的例外**

1. **gopkg.in**

    那些使用`gopkg.in` (如 `gopkg.in/yaml.v1` and `gopkg.in/yaml.v2`)开头作为`import path`的代码，依然能够沿用现有的导入方式，即便使用`module`之后也是如此。

2. **'+incompatible' when importing non-module v2+ packages**

    模块可以import那些没有使用模块化v2+版本的包，即使版本库有遵循semver的v2+ tag 提交（如v2.0.0），也会视其为v1系列的扩展（不兼容）处理，并通过'+incompatible'在`go.mod`文件中标记。

    '+incompatible'意味着：

    * v2+的包没有主动选择加入模块

    * 或者没有将`major`版本号作为`module path`的一部分

3. **"Minimal module compatibility" when module mode is not enabled（未启用模块模式时的“最小模块兼容性”）**

    在引入module模式之前（即Go 1.11版本之前），代码中是不存在模块化的概念的。为了让这些未启用module模式的代码（其中Go 1.11版本须设置Go111Module=off，表示关闭module模式），能够import陆续使用了module模式包含**major**版本号的包（如：`import github.com/A/v2/pkg`）,Go对版本1.9.7+、1.10.3+和1.11进行了更新,使得在使用这些Go版本构建的代码，能够在不需要修改现有代码的情况下，同时也能正确处理v2+模块，这种行为被称为“**未启用模块模式时的最小模块兼容性**”（仅当Go工具禁用了完整模块模式时才会生效）。

有关required V2+模块的确切机制，请参阅下面的["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher)部分。

## 如何使用模块

### 如何安装和激活模块模式支持

使用模块模式，有两种安装方法可选:
* 安装最新的Go 1.11 release版本
* 从`master`分支通过源码安装Go toolchain

你可以使用下面任意一种方法激活模块支持：
* 在`$GOPATH/src`目录之外，一个包含有效的`go.mod`文件的目录（或其父目录中包含该文件）下，执行`go`命令，并且没有设置环境变量`GO111MODULE=off`（默认值是`auto`）
* 设置`GO111MODULE=on`，执行`go`命令。

### How to Define a Module

为已存在的项目创建`go.mod`文件：

1. 进入`GOPATH`之外的模块代码的根目录：

   ```
   $ cd <project path outside $GOPATH/src>         # e.g., cd ~/projects/hello
   ```
   在`GOPATH`之外，你不需要设置`GO111MODULE`来激活module模式。

   另一种方式, 或者你希望在`GOPATH`内来做这项工作:

   ```
   $ export GO111MODULE=on                         # manually active module mode
   $ cd $GOPATH/src/<project path>                 # e.g., cd $GOPATH/src/you/hello
   ```

2. 创建并初始化module定义，生成`go.mod`文件：

   ```
   $ go mod init
   ```
  此步骤从现有存在的[`dep`](https://github.com/golang/dep) `Gopkg.lock`文件或者九种支持的依赖格式进行转换，添加与现有配置匹配的`require`声明。

  九种支持的依赖格式：

  ```
    var Converters = map[string]func(string, []byte) (*modfile.File, error){
        "GLOCKFILE":          ParseGLOCKFILE,
        "Godeps/Godeps.json": ParseGodepsJSON,
        "Gopkg.lock":         ParseGopkgLock,
        "dependencies.tsv":   ParseDependenciesTSV,
        "glide.lock":         ParseGlideLock,
        "vendor.conf":        ParseVendorConf,
        "vendor.yml":         ParseVendorYML,
        "vendor/manifest":    ParseVendorManifest,
        "vendor/vendor.json": ParseVendorJSON,
    }
  ```

   `go mod init` 通常能够使用辅助数据（如VCS元数据）自动确定适当的模块路径，但如果`go mod init`声明它不能自动确定`module path`，或者如果您需要重写该路径，则可以将`module path`作为`go mod init`的可选参数提供，例如：

   ```
   $ go mod init github.com/my/repo
   ```

   请注意，如果您的依赖项包括v2+模块，或者如果您正在初始化v2+模块，那么在运行`go mod init`之后，您可能还需要编辑`go.mod`和`.go`代码以添加`/vn`到`import path`和`module path`，如上文<a href="#Semantic_Import_Versioning" target="_self">语义导入版本控制</a>部分所述。即使go mod init自动从`dep`或9种支持的依赖格式转换依赖关系信息，这也适用。（因此，在运行`go mod init`之后，在成功运行`go build`之前，通常不应该立即运行`go mod tidy`，更应该是按照本节介绍的顺序）。

3. 构建模块，当从模块的根目录执行`build`命令时，`./..`模式匹配当前模块中的所有包。`go build`将根据需要自动添加缺少的或未转换的依赖项，以满足构建需要。
   ```
   $ go build ./...
   ```

4. 按配置测试模块，以确保其与所选版本一起工作：
   ```
   $ go test ./...
   ```

5. （可选）运行模块的测试以及所有直接和间接依赖项的测试，以检查兼容性：
   ```
   $ go test all
   ```

在标记发布之前，请参阅下面的“["How to Prepare for a Release"](https://github.com/golang/go/wiki/Modules#how-to-prepare-for-a-release) ”部分。

更多主题相关信息，都可以在[官方modules文档](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)中的找到入口。

## How to Upgrade and Downgrade Dependencies（如何对依赖项做版本升级和降级处理）

使用“go get”完成依赖项的日常升级和降级，它将自动更新`go.mod`文件。或者，您可以直接编辑go.mod，指定依赖项的版本。

此外，“go build”、“go test”甚至“go list”等go命令将根据需要自动添加新的依赖项以满足导入需要（更新go.mod并下载新的依赖项）。

要查看所有直接和间接依赖项的可用可升级的`minor`和`patch`版本，请运行`go list -u -m all`。

将当前模块的所有直接和间接依赖项升级到最新版本：

 * 运行 `go get -u` 使用最新发布的*minor 或 patch*版本
 * 运行 `go get -u=patch` 使用最新发布的*patch* 版本

要升级或降级到更具体的版本，“go get”允许通过向package参数添加@version后缀或“module query”的方式来覆盖版本的自动选择，例如：`go get foo@v1.6.2、go` 或 `get foo@e3702bed2` 或 `go get foo@'<v1.6.2'`。

`go get foo` 获取使用foo semver tag的最新版本来更新，如果没有semver tag，则获取已知的最新提交。`go get foo`相当于`go get foo@latest` - 换句话说，如果未指定@version，则默认为@latest。

注意，`go get -u foo`和`go get -u foo@latest` 同时还升级了foo的所有直接和间接依赖。一种常见的做法是不带`-u`的`go get foo`或`go get foo@latest`（并且在此之后，考虑选择性的使用`go get -u=patch foo`、`go get -u foo`、`go get -u` 或 `go get -u=patch`）。

使用诸如`go get foo@master`之类的分支名称是获取最新提交的一种方法，不管它是否有semver tag。

通常，无法解析为semver tag的模块查询，将作为`pseudo-versions`(伪版本)记录在`go.mod`文件中。

**[pseudo-versions](https://tip.golang.org/cmd/go/#hdr-Pseudo_versions) （伪版本）**

`go.mod`文件和go命令通常使用语义版本作为描述模块版本的标准形式，这样就可以比较决定版本间的先后。像v1.2.3这样的模块版本是通过版本库的提交而来的，对于一直未做tag的提交，可以使用“伪版本”如`v0.0.0-yyyymmddhhmmss-abcdefabcdef`，其中时间是以UTC表示的提交时间，最后的后缀是提交哈希的前缀。时间部分确定可以将两个伪版本的先后，哈希值代表提交标识，前缀（在本例中为v0.0.0）是从最新的tagged version产生的。

有三种伪版本形式：

* vX.0.0-yymmddhhmmss-abcdefabcdefabcdef - 被用于在目标提交之前没有具有适合更早的`major`版本的时候。

* vX.Y.Z-pre.0.yyyymmddhhmmss-abcdefabcdef - 被用于在目标提交之前的最新提交版本为vX.Y.Z-pre的时候。

* vX.Y.(Z+1)-0.yyyymmddhhmmss-abcdefabcdef - 被用于在目标提交之前的最新提交版本为vX.Y.Z的时候。

伪版本永远不需要手工键入：go命令将接受纯提交哈希，并自动将其转换为伪版本（或标记版本，如果可用）。

参考 ["Module-aware go get"](https://golang.org/cmd/go/#hdr-Module_aware_go_get) 和 ["Module queries"](https://golang.org/cmd/go/#hdr-Module_queries) 以查看与主题相关的更多信息。

模块能够使用尚未进行模块化的包，还可以使用还没有任何合适semver tag的包（在这种情况下，它们将使用go.mod中的伪版本进行记录），并对他们做升级和降级处理。

在升级或降级任何依赖项之后，您可能需要对构建中的所有包（包括直接和间接依赖项）再次运行测试，以检查不兼容性：

   ```
   $ go test all
   ```

## How to Prepare for a Release

### Releasing Modules (All Versions)

Best practices for creating a release of a module are expected to emerge as part of the initial modules experiment. Many of these might end up being automated by a [future 'go release' tool](https://github.com/golang/go/issues/26420).

Some current suggested best practices to consider prior to tagging a release:

* Run `go mod tidy` to possibly prune any extraneous requirements (as described [here](https://tip.golang.org/cmd/go/#hdr-Maintaining_module_requirements)) and also ensure your current go.mod reflects all possible build tags/OS/architecture combinations (as described [here](https://github.com/golang/go/issues/25971#issuecomment-399091682)).
  * In contrast, other commands like `go build` and `go test` will not remove dependencies from `go.mod` that are no longer required and only update `go.mod` based on the current build invocation's tags/OS/architecture.

* Run `go test all` to test your module (including running the tests for your direct and indirect dependencies) as a way of validating that the currently selected packages versions are compatible.
  * The number of possible version combinations is exponential in the number of modules, so in general you cannot expect your dependencies to have tested against all possible combinations of their dependencies.
  * As part of the modules work, `go test all` has been [re-defined to be more useful](https://research.swtch.com/vgo-cmd) to include all the packages in the current module, plus all the packages they depend on through a sequence of one or more imports, while excluding packages that don't matter in the current module.

* Ensure your `go.sum` file is committed along with your `go.mod` file. See [FAQ below](https://github.com/golang/go/wiki/Modules#should-i-commit-my-gosum-file-as-well-as-my-gomod-file) for more details and rationale.

### Releasing Modules (v2 or Higher)

If you are releasing a v2 or higher module, please first review the discussion in the ["Semantic Import Versioning" ](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) section above, which includes why major versions are included in the module path and import path for v2+ modules, as well as how Go versions 1.9.7+ and 1.10.3+ have been updated to simplify that transition.

Note that if you are adopting modules for the first time for a pre-existing repository or set of packages that have already been tagged `v2.0.0` or higher before adopting modules, then the [recommended best practice](https://github.com/golang/go/issues/25967#issuecomment-422828770) is to increment the major version when first adopting modules. For example, if you are the author of `foo`, and the latest tag for the `foo` repository is `v2.2.2`, and `foo` has not yet adopted modules, then the best practice would be to use `v3.0.0` for the first release of `foo` to adopt modules (and hence the first release of `foo` to contain a `go.mod` file). Incrementing the major version in this case provides greater clarity to consumers of `foo`, allows for additional non-module patches or minor releases on the v2 series of `foo` if needed, and provides a strong signal for a module-based consumer of `foo` that different major versions result if you do `import "foo"` and a corresponding `require foo v2.2.2+incompatible`, vs. `import "foo/v3"` and a corresponding `require foo/v3 v3.0.0`. (Note that this advice regarding incrementing the major version when first adopting modules does _not_ apply to pre-existing repos or packages whose latest versions are v0.x.x or v1.x.x).

There are two alternative mechanisms to release a v2 or higher module. Note that with both techniques, the new module release becomes available to consumers when the module author pushes the new tags. Using the example of creating a `v3.0.0` release, the two options are:

1. **Major branch**: Update the `go.mod` file to include a `/v3` at the end of the module path in the `module` directive (e.g., `module github.com/my/module/v3`). Update import statements within the module to also use `/v3` (e.g., `import "github.com/my/module/v3/mypkg"`). Tag the release with `v3.0.0`.
   * Go versions 1.9.7+, 1.10.3+, and 1.11 are able to properly consume and build a v2+ module created using this approach without requiring updates to consumer code that has not yet opted in to modules (as described in the the ["Semantic Import Versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) section above).
   * A community tool [github.com/marwan-at-work/mod](https://github.com/marwan-at-work/mod) helps automate this procedure. See the [repository](https://github.com/marwan-at-work/mod) or the [community tooling FAQ](https://github.com/golang/go/wiki/Modules#what-community-tooling-exists-for-working-with-modules) below for an overview.
   * To avoid confusion with this approach, consider putting the `v3.*.*` commits for the module on a separate v3 branch.
   * **Note:** creating a new branch is _not_ required. If instead you have been previously releasing on master and would prefer to tag `v3.0.0` on master, that is a viable option. (However, be aware that introducing an incompatible API change in `master` can cause issues for non-modules users who issue a `go get -u` given the `go` tool is not aware of [semver](https://semver.org) prior to Go 1.11 or when [module mode](https://github.com/golang/go/wiki/Modules#when-do-i-get-old-behavior-vs-new-module-based-behavior) is not enabled in Go 1.11+).

2. **Major subdirectory**: Create a new `v3` subdirectory (e.g., `my/module/v3`) and place a new `go.mod` file in that subdirectory. The module path must end with `/v3`. Copy or move the code into the `v3` subdirectory. Update import statements within the module to also use `/v3` (e.g., `import "github.com/my/module/v3/mypkg"`). Tag the release with `v3.0.0`.
   * This provides greater backwards compatibility. In particular, Go versions older than 1.9.7 and 1.10.3 are also able to properly consume and build a v2+ module created using this approach.
   * A more sophisticated approach here could exploit type aliases (introduced in Go 1.9) and forwarding shims between major versions residing in different subdirectories.  This can provide additional compatibility and allow one major version to be implemented in terms of another major version, but would entail more work for a module author. An in-progress tool to automate this is `goforward`. Please see [here](https://golang.org/cl/137076) for more details and rationale, along with a functioning initial version of `goforward`.

See https://research.swtch.com/vgo-module for a more in-depth discussion of these alternatives.

### Publishing a release

A new module version may be published by pushing a tag to the repository that contains the module source code. The tag is formed by concatenating two strings: a *prefix* and a *version*.

The *version* is the semantic import version for the release. It should be chosen by following the rules of [semantic import versioning](#semantic-import-versioning).

The *prefix* indicates where a module is defined within a repository. If the module is defined at the root of the repository, the prefix is empty, and the tag is just the version. However, in [multi-module repositories](#faqs--multi-module-repositories), the prefix distinguishes versions for different modules. The prefix is the directory within the repository where the module is defined. If the repository follows the major subdirectory pattern described above, the prefix does not include the major version suffix.

For example, suppose we have a module `example.com/repo/sub/v2`, and we want to publish version `v2.1.6`. The repository root corresponds to `example.com/repo`, and the module is defined in `sub/v2/go.mod` within the repository. The prefix for this module is `sub/`. The full tag for this release should be `sub/v2.1.6`.

## Migrating to Modules

This section attempts to briefly enumerate the major decisions to be made when migrating to modules as well as list other migration-related topics. References are generally provided to other sections for more details.

This material is primarily based on best practices that have emerged from the community as part of the modules experiment; this is therefore a work-in-progress section that will improve over time as the community gains more experience.

Summary:

* The modules system is designed to allow different packages in the overall Go ecosystem to opt in at different rates.
* Packages that are already on version v2 or higher have more migration considerations, primarily due to the implications of [Semantic Import versioning](https://github.com/golang/go/wiki/Modules#semantic-import-versioning).
* New packages and packages on v0 or v1 have substantially fewer considerations when adopting modules.
* Modules defined with Go 1.11 can be used by older Go versions (although the exact Go versions depends on the strategy used by the main module and its dependencies, as outlined below).

Migration topics:

#### Automatic Migration from Prior Dependency Managers

  * `go mod init` automatically translates the required information from [dep, glide, govendor, godep and 5 other pre-existing dependency managers](https://tip.golang.org/pkg/cmd/go/internal/modconv/?m=all#pkg-variables) into a `go.mod `file that produces the equivalent build.
  * If you are creating a v2+ module, be sure your `module` directive in the converted `go.mod` includes the appropriate `/vN` (e.g., `module foo/v3`).
  * Note that if you are importing v2+ modules, you might need to do some manual adjustments after an initial conversion in order to add `/vN` to the `require` statements that `go mod init` generates after translating from a prior dependency manager. See the ["How to Define a Module"](https://github.com/golang/go/wiki/Modules#how-to-define-a-module) section above for more details.
  * In addition, `go mod init` will not edit your `.go` code to add any required `/vN` to import statements. See the ["Semantic Import versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) and ["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) sections above for the required steps, including some options around community tools to automate the conversion.

#### Providing Dependency Information to Older Versions of Go

  * Older versions of Go understand how to consume a vendor directory created by `go mod vendor`. Therefore, vendoring is one way for a module to provide dependencies to older versions of Go that do not fully understand modules. See the [vendoring FAQ](https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away) and the `go` command [documentation](https://tip.golang.org/cmd/go/#hdr-Modules_and_vendoring) for more details.

#### Updating Pre-Existing Install Instructions

  * Pre-modules, it is common for install instructions to include `go get -u foo`. If you are publishing a module `foo`, consider dropping the `-u` in instructions for modules-based consumers.
     * `-u` asks the `go` tool to upgrade all the direct and indirect dependencies of `foo`.
	 * A module consumer might choose to run `go get -u foo` later, but there are more benefits of ["High Fidelity Builds"](https://github.com/golang/proposal/blob/master/design/24301-versioned-go.md#update-timing--high-fidelity-builds) if `-u` is not part of the initial install instructions. See ["How to Upgrade and Downgrade Dependencies"](https://github.com/golang/go/wiki/Modules#how-to-upgrade-and-downgrade-dependencies) for more details.
     * `go get -u foo` does still work, and can still be a valid choice for install instructions.
  * In addition, `go get foo` is not strictly needed for a module-based consumer.
     * Simply adding an import statement `import "foo"` is sufficient. (Subsequent commands like `go build` or `go test` will automatically download `foo` and update `go.mod` as needed).
  * Module-based consumers will not use a `vendor` directory by default.
     * When module mode is enabled in the `go` tool, `vendor` is not strictly required when consuming a module (given the information contained in `go.mod` and the cryptographic checksums in `go.sum`), but some pre-existing install instructions assume the `go` tool will use `vendor` by default. See the [vendoring FAQ](https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away) for more details.
  * Install instructions that include `go get foo/...` might have issues in some cases (see discussion in [#27215](https://github.com/golang/go/issues/27215#issuecomment-427672781)).

#### Incrementing the Major Version When First Adopting Modules with v2+ Packages

* If you have packages that have already been tagged v2.0.0 or higher before adopting modules, then the recommended best practice is to increment the major version when first adopting modules. For example, if you are on `v2.0.1` and have not yet adopted modules, then you would use `v3.0.0` for the first release that adopts modules. See the ["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) section above for more details.

#### v2+ Modules Allow Multiple Major Versions Within a Single Build

* If a module is on v2 or higher, an implication is that multiple major versions can be in a single build (e.g., `foo` and `foo/v3` might end up in a single build).
  * This flows naturally from the rule that "packages with different import paths are different packages".
  * When this happens, there will be multiple copies of package-level state (e.g., package-level state for `foo` and package-level state for `foo/v3`) as well as each major version will run its own `init` function.
  * This approach helps with multiple aspects of the modules system, including helping with diamond dependency problems, gradual migration to new versions within large code bases, and allowing a major version to be implemented as a shim around a different major version.
* See the "Avoiding Singleton Problems" section of https://research.swtch.com/vgo-import or [#27514](https://github.com/golang/go/issues/27514) for some related discussion.

#### Modules Consuming Non-Module Code

  * Modules are capable of consuming packages that have not yet opted into modules, with the appropriate package version information recorded in the importing module's `go.mod`.  Modules can consume packages that do not yet have any proper semver tags. See FAQ [below](https://github.com/golang/go/wiki/Modules#can-a-module-consume-a-package-that-has-not-opted-in-to-modules) for more details.
  * Modules can also import a v2+ package that has not opted into modules. It will be recorded with an `+incompatible` suffix if the imported v2+ package has valid semver tags. See FAQ [below](https://github.com/golang/go/wiki/Modules#can-a-module-consume-a-v2-package-that-has-not-opted-into-modules-what-does-incompatible-mean) for more details.

#### Non-Module Code Consuming Modules

  * **Non-module code consuming v0 and v1 modules**:
     * Code that has not yet opted in to modules can consume and build v0 and v1 modules (without any requirement related to the Go version used).

  * **Non-module code consuming v2+ modules**:

    * Go versions 1.9.7+, 1.10.3+ and 1.11 have been updated so that code built with those releases can properly consume v2+ modules without requiring modification of pre-existing code as described in the ["Semantic Import versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) and ["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) sections above.

    * Go versions prior to 1.9.7 and 1.10.3 can consume v2+ modules if the v2+ module was created following the "Major subdirectory" approach outlined in the ["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) section.

#### Strategies for Authors of Pre-Existing v2+ Packages

For authors of pre-existing v2+ packages considering opting in to modules, one way to summarize the alternative approaches is as a choice between three top-level strategies . Each choice then has follow-on decisions and variations (as outlined above). These alternative top-level strategies are:

1. **Require clients to use Go versions 1.9.7+, 1.10.3+, or 1.11+**.

    The approach uses the "Major Branch" approach and relies on the "minimal module awareness" that was backported to 1.9.7 and 1.10.3. See the ["Semantic Import versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) and ["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) sections above for more details.

2. **Allow clients to use even older Go versions like Go 1.8**.

    This approach uses the "Major Subdirectory" approach and involves creating a subdirectory such as `/v2` or `/v3`. See the ["Semantic Import versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) and ["Releasing Modules (v2 or Higher)"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) sections above for more details.

3. **Wait on opting in to modules**.

    In this strategy, things continue to work with client code that has opted in to modules as well as with client code that has not opted in to modules. As time goes by, Go versions 1.9.7+, 1.10.3+, and 1.11+ will be out for an increasingly longer time period, and at some point in the future, it becomes more natural or client-friendly to require Go versions 1.9.7+/1.10.3+/1.11+, and at that point in time, you can implement strategy 1 above (requiring Go versions 1.9.7+, 1.10.3+, or 1.11+) or even strategy 2 above (though if you are ultimately going to go with strategy 2 above in order to support older Go versions like 1.8, then that is something you can do now).

## Additional Resources

### Documentation and Proposal

* Official documentation:
  * Latest [HTML documentation for modules on golang.org](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more)
  * Run `go help modules` for more about modules. (This is the main entry point for modules topics via `go help`)
  * Run `go help mod` for more about the `go mod` command.
  * Run `go help module-get` for more about the behavior of `go get` when in module-aware mode.
  * Run `go help goproxy` for more about the module proxy, including a pure file-based option via a `file:///` URL.
* The initial ["Go & Versioning"](https://research.swtch.com/vgo) series of blog posts on `vgo` by Russ Cox (first posted February 20, 2018)
* Official [golang.org blog post introducing the proposal](https://blog.golang.org/versioning-proposal) (March 26, 2018)
  * This provides a more succinct overview of the proposal than the full `vgo` blog series, along with some of the history and process behind the proposal
* Official [Versioned Go Modules Proposal](https://golang.org/design/24301-versioned-go) (last updated March 20, 2018)

### Introductory Material

* Introductory 40 minute video ["The Principles of Versions in Go"](https://www.youtube.com/watch?v=F8nrpe0XWRg&list=PLq2Nv-Sh8EbbIjQgDzapOFeVfv5bGOoPE&index=3&t=0s) from GopherCon Singapore by Russ Cox (May 2, 2018)
  * Succinctly covers the philosophy behind the design of versioned Go modules, including the three core principles of "Compatibility", "Repeatability", and "Cooperation"
* Example based 35 minute introductory video ["What are Go modules and how do I use them?"](https://www.youtube.com/watch?v=6MbIzJmLz6Q&list=PL8QGElREVyDA2iDrPNeCe8B1u7li5S6ep&index=5&t=0s) ([slides](https://talks.godoc.org/github.com/myitcv/talks/2018-08-15-glug-modules/main.slide#1)) by Paul Jolly (August 15, 2018)
* Introductory blog post ["Taking Go Modules for a Spin"](https://dave.cheney.net/2018/07/14/taking-go-modules-for-a-spin) by Dave Cheney (July 14, 2018)
* Introductory [Go Meetup slides on modules](https://docs.google.com/presentation/d/1ansfXN8a_aVL-QuvQNY7xywnS78HE8aG7fPiITNQWMM/edit#slide=id.g3d87f3177d_0_0) by Chris Hines (July 16, 2018)
* Introductory 30 minute video ["Intro to Go Modules and SemVer"](https://www.youtube.com/watch?v=aeF3l-zmPsY) by Francesc Campoy (Nov 15, 2018)

### Additional Material

* Blog post ["Using Go modules with vendor support on Travis CI"](https://arslan.io/2018/08/26/using-go-modules-with-vendor-support-on-travis-ci/) by Fatih Arslan (August 26, 2018)
* Blog post ["Go Modules and CircleCI"](https://medium.com/@toddkeech/go-modules-and-circleci-c0d6fac0b000) by Todd Keech (July 30, 2018)
* Blog post ["The vgo proposal is accepted. Now what?"](https://research.swtch.com/vgo-accepted) by Russ Cox (May 29, 2018)
  * Includes summary of what it means that versioned modules are currently an experimental opt-in feature
* Blog post on [how to build go from tip and start using go modules](https://carolynvanslyck.com/blog/2018/07/building-go-from-source/) by Carolyn Van Slyck (July 16, 2018)

## Changes Since the Initial Vgo Proposal

As part of the proposal, prototype, and beta processes, there have been over 400 issues created by the overall community. Please continue to supply feedback.

Here is a partial list of some of the larger changes and improvements, almost all of which were primarily based on feedback from the community:

* Top-level vendor support was retained rather than vgo-based builds ignoring vendor directories entirely ([discussion](https://groups.google.com/d/msg/golang-dev/FTMScX1fsYk/uEUSjBAHAwAJ), [CL](https://go-review.googlesource.com/c/vgo/+/118316))
* Backported minimal module-awareness to allow older Go versions 1.9.7+ and 1.10.3+ to more easily consume modules for v2+ projects ([discussion](https://github.com/golang/go/issues/24301#issuecomment-371228742),  [CL](https://golang.org/cl/109340))
* Allowed vgo to use v2+ tags by default for pre-existing packages did not yet have a go.mod (recent update in related behavior described [here](https://github.com/golang/go/issues/25967#issuecomment-407567904))
* Added support via command `go get -u=patch` to update all transitive dependencies to the latest available patch-level versions on the same minor version ([discussion](https://research.swtch.com/vgo-cmd), [documentation](https://tip.golang.org/cmd/go/#hdr-Module_aware_go_get))
* Additional control via environmental variables (e.g., GOFLAGS in [#26585](https://github.com/golang/go/issues/26585), [CL](https://go-review.googlesource.com/c/go/+/126656))
* Finer grain control on whether or not go.mod is allowed to be updated, how vendor directory is used, and whether or not network access is allowed (e.g., -mod=readonly, -mod=vendor, GOPROXY=off; related [CL](https://go-review.googlesource.com/c/go/+/126696) for recent change)
* Added more flexible replace directives ([CL](https://go-review.googlesource.com/c/vgo/+/122400))
* Added additional ways to interrogate modules (for human consumption, as well as for better editor / IDE integration)
* The UX of the go CLI has continued to be refined based on experiences so far (e.g., [#26581](https://github.com/golang/go/issues/26581), [CL](https://go-review.googlesource.com/c/go/+/126655))
* Additional support for warming caches for use cases such as CI or docker builds via `go mod download` ([#26610](https://github.com/golang/go/issues/26610#issuecomment-408654653))
* **Most likely**: better support for installing specific versions of programs to GOBIN ([#24250](https://github.com/golang/go/issues/24250#issuecomment-377553022))

## GitHub Issues

* [Currently open module issues](https://golang.org/issues?q=is%3Aopen+is%3Aissue+label:modules)
* [Closed module issues](https://github.com/golang/go/issues?q=is%3Aclosed+is%3Aissue+label%3Amodules+sort%3Aupdated-desc)
* [Closed vgo issues](https://github.com/golang/go/issues?q=-label%3Amodules+vgo+is%3Aclosed+sort%3Aupdated-desc)
* Submit a [new module issue](https://github.com/golang/go/issues/new?title=cmd%2Fgo%3A%20%3Cfill%20this%20in%3E) using 'cmd/go:' as the prefix

## FAQs

### How are versions marked as incompatible?

The `require` directive allows any module to declare that it should be built with version >= x.y.z of a dependency D (which may be specified due to  incompatibilities with version < x.y.z of module D). Empirical data suggests [this is the dominant form of constraints used in `dep` and `cargo`](https://twitter.com/_rsc/status/1022590868967116800). In addition, the top-level module in the build can `exclude` specific versions of dependencies or `replace` other modules with different code. See the full proposal for [more details and rationale](https://github.com/golang/proposal/blob/master/design/24301-versioned-go.md).

One of the key goals of the versioned modules proposal is to add a common vocabulary and semantics around versions of Go code for both tools and developers. This lays a foundation for future capabilities to declare additional forms of incompatibilities, such as possibly:
* declaring deprecated versions as [described](https://research.swtch.com/vgo-module) in the initial `vgo` blog series
* declaring pair-wise incompatibility between modules in an external system as discussed for example [here](https://github.com/golang/go/issues/24301#issuecomment-392111327) during the proposal process
* declaring pair-wise incompatible versions or insecure versions of a module after a release has been published. See for example the on-going discussion in [#24031](https://github.com/golang/go/issues/24031#issuecomment-407798552) and [#26829](https://github.com/golang/go/issues/26829)

### When do I get old behavior vs. new module-based behavior?

In general, modules are opt-in for Go 1.11, so by design old behavior is preserved by default.

Summarizing when you get the old 1.10 status quo behavior vs. the new opt-in modules-based behavior:

* Inside GOPATH — defaults to old 1.10 behavior (ignoring modules)
* Outside GOPATH while inside a file tree with a `go.mod` — defaults to modules behavior
* GO111MODULE environment variable:
  * unset or `auto` —  default behavior above
  * `on` —  force module support on regardless of directory location
  * `off` — force module support off regardless of directory location

### Can two modules depend on each other (cyclical import)?
Yes. Two packages however may not depend on each other (this is a build constraint).

### Can a module depend on a different version of itself?
A module can depend on a different major version of itself: by-and-large, this is comparable to depending on a different module.

### Why does installing a tool via `go get` fail with error `cannot find main module`?

This occurs when you have set `GO111MODULE=on`, but are not inside of a file tree with a `go.mod` when you run `go get`.

The simplest solution is to leave `GO111MODULE` unset (or equivalently explicitly set to `GO111MODULE=auto`), which avoids this error.

Recall one of the primary reason modules exist is to record precise dependency information. This dependency information is written to your current `go.mod`.  If you are not inside of a file tree with a `go.mod` but you have told the `go get` command to operate in module mode by setting `GO111MODULE=on`, then running `go get` will result in the error `cannot find main module` because there is no `go.mod` available to record dependency information.

Solution alternatives include:

1. Leave `GO111MODULE` unset (the default, or explicitly set `GO111MODULE=auto`), which results in friendlier behavior. This will give you Go 1.10 behavior when you are outside of a module and hence will avoid `go get` reporting `cannot find main module`.

2. Leave `GO111MODULE=on`, but as needed disable modules temporarily and enable Go 1.10 behavior during `go get`, such as via `GO111MODULE=off go get example.com/cmd`. This can be turned into a simple script or shell alias such as `alias oldget='GO111MODULE=off go get'`

3. Create a temporary `go.mod` file that is then discarded. This has been automated by a [simple shell script](https://gist.github.com/rogpeppe/7de05eef4dd774056e9cf175d8e6a168) by [@rogpeppe](https://github.com/rogpeppe). This script allows version information to optionally be supplied via `vgoget example.com/cmd[@version]`. (This can be a solution for avoiding the error `cannot use path@version syntax in GOPATH mode`).

4. `gobin` is a module-aware command to install and run main packages. By default, `gobin` installs/runs main packages without first needing to manually create a module, but with the `-m` flag it can be told to use an existing module to resolve dependenci
es. Please see the `gobin` [README](https://github.com/myitcv/gobin#usage) and [FAQ](https://github.com/myitcv/gobin/wiki/FAQ) for details and additional use cases.

5. Create a `go.mod` you use to track your globally installed tools, such as in `~/global-tools/go.mod`, and `cd` to that directory prior to running `go get` or `go install` for any globally installed tools.

6. Create a `go.mod` for each tool in separate directories, such as `~/tools/gorename/go.mod` and `~/tools/goimports/go.mod`, and `cd` to that appropriate directory prior to running `go get` or `go install` for the tool.

This current limitation will be resolved. However, the primary issue is that modules are currently opt-in, and a full solution will likely wait until GO111MODULE=on becomes the default behavior. See [#24250](https://github.com/golang/go/issues/24250#issuecomment-377553022) for more discussion, including this comment:

> This clearly must work eventually. The thing I'm not sure about is exactly what this does as far as the version is concerned: does it create a temporary module root and go.mod, do the install, and then throw it away? Probably. But I'm not completely sure, and for now I didn't want to confuse people by making vgo do things outside go.mod trees. Certainly the eventual go command integration has to support this.

This FAQ has been discussing tracking _globally_ installed tools.

If instead you want to track the tools required by a _specific_ module, see the next FAQ.

### How can I track tool dependencies for a module?

If you:
 *  want to use a go-based tool (e.g. `stringer`) while working on a module, and
 *  want to ensure that everyone is using the same version of that tool while tracking the tool's version in your module's `go.mod` file

then one currently recommended approach is to add a `tools.go` file to your module that includes import statements for the tools of interest (such as `import _ "golang.org/x/tools/cmd/stringer"`), along with a `// +build tools` build constraint. The import statements allow the `go` command to precisely record the version information for your tools in your module's `go.mod`, while the `// +build tools` build constraint prevents your normal builds from actually importing your tools.

For a concrete example of how to do this, please see this ["Go Modules by Example" walkthrough](https://github.com/go-modules-by-example/index/blob/master/010_tools/README.md).

A discussion of the approach along with an earlier concrete example of how to do this is in [this comment in #25922](https://github.com/golang/go/issues/25922#issuecomment-412992431).

The brief rationale (also from [#25922](https://github.com/golang/go/issues/25922#issuecomment-402918061)):

> I think the tools.go file is in fact the best practice for tool dependencies, certainly for Go 1.11.
>
> I like it because it does not introduce new mechanisms.
>
> It simply reuses existing ones.

### What is the status of module support in IDEs, editors and standard tools like goimports, gorename, etc?

Support for modules is starting to land in editors and IDEs.

For example:
* **GoLand**: currently has full support for modules outside and inside GOPATH, including completion, syntax analysis, refactoring, navigation as described [here](https://blog.jetbrains.com/go/2018/08/24/goland-2018-2-2-is-here/).
* **VS Code**: work is in progress and looking for contributors to help. Tracking issue is [#1532](https://github.com/Microsoft/vscode-go/issues/1532). An initial beta is described in the [VS Code module status wiki page](https://github.com/Microsoft/vscode-go/wiki/Go-modules-support-in-Visual-Studio-Code).
* **Atom with go-plus**: tracking issue is [#761](https://github.com/joefitzgerald/go-plus/issues/761).
* **vim with vim-go**: initial support for syntax highlighting and formatting `go.mod` has [landed](https://github.com/fatih/vim-go/pull/1931). Broader support tracked in [#1906](https://github.com/fatih/vim-go/issues/1906).
* **emacs with go-mode.el**: tracking issue in [#237](https://github.com/dominikh/go-mode.el/issues/237).

The status of other tools such as goimports, guru, gorename and similar tools is being tracked in an umbrella issue [#24661]( https://github.com/golang/go/issues/24661). Please see that umbrella issue for latest status.

Some tracking issues for particular tools includes:
* **gocode**: tracking issue in [mdempsky/gocode/#46](https://github.com/mdempsky/gocode/issues/46). Note that `nsf/gocode` is recommending people migrate off of `nsf/gocode` to `mdempsky/gocode`.
* **go-tools** (tools by dominikh such as staticcheck, megacheck, gosimple): sample tracking issue [dominikh/go-tools#328](https://github.com/dominikh/go-tools/issues/328).

In general, even if your editor, IDE or other tools have not yet been made module aware, much of their functionality should work with modules if you are using modules inside GOPATH and do `go mod vendor` (because then the proper dependencies should be picked up via GOPATH).

The full fix is to move programs that load packages off of `go/build` and onto `golang.org/x/tools/go/packages`, which understands how to locate packages in a module-aware manner. This will likely eventually become `go/packages`.

## FAQs — Additional Control

### What community tooling exists for working with modules?

The community is starting to build tooling on top of modules. For example:

* [github.com/rogpeppe/gohack](https://github.com/rogpeppe/gohack)
  * A new community tool to automate and greatly simplify `replace` and multi-module workflows, including allowing you to easily modify one of your dependencies
  * For example, `gohack example.com/some/dependency` automatically clones the appropriate repository and adds the necessary `replace` directives to your `go.mod`
  * Remove all gohack replace statements with `gohack undo`
  * The project is continuing to expand to make other module-related workflows easier
* [github.com/marwan-at-work/mod](https://github.com/marwan-at-work/mod)
  * Command line tool to automatically upgrade/downgrade major versions for modules
  * Automatically adjusts `go.mod` files and related import statements in go source code
  * Helps with upgrades, or when first opting in to modules with a v2+ package
* [github.com/goware/modvendor](https://github.com/goware/modvendor)
  * Helps copy additional files into the `vendor/` folder, such as shell scripts, .cpp files, .proto files, etc.

### When should I use the replace directive?

* As described in the ['go.mod' concepts section above](https://github.com/golang/go/wiki/Modules#gomod), `replace` directives provide additional control in the top-level `go.mod` for what is actually used to satisfy a dependency found in the Go source or go.mod files, while `replace` directives in modules other than the main module are ignored when building the main module.
* The `replace` directive allows you to supply another import path that might be another module located in VCS (GitHub or elsewhere), or on your local filesystem with a relative or absolute file path. The new import path from the `replace` directive is used without needing to update the import paths in the actual source code.
* One sample use case is if you need to fix or investigate something in a dependency, you can have a local fork and add the something like the following in your top-level `go.mod`:
   * `replace example.com/original/import/path => /your/forked/import/path`
* `replace` also allows the top-level module control over the exact version used for a dependency, such as:
   * `replace example.com/some/dependency => example.com/some/dependency v1.2.3`
* `replace` also can be used to inform the go tooling of the relative or absolute on-disk location of modules in a multi-module project, such as:
   * `replace example.com/project/foo => ../foo`
* In general, you have the option of specifying a version to the left of the `=>` in a replace directive, but typically it is less sensitive to change if you omit that (e.g., as done in all of the `replace` examples above).
* **Note**: for direct dependencies, a `require` directive is needed even when doing a `replace`. For example, if `foo` is a direct dependency, you cannot do `replace foo => ../foo` without a corresponding `require` for `foo`. (If you are not sure what version to use in the `require` directive, you can often use `v0.0.0` such as `require foo v0.0.0`; see [#26241](https://golang.org/issue/26241)).
* You can confirm you are getting your expected versions by running `go list -m all`, which shows you the actual final versions that will be used in your build including taking into account `replace` statements.
* See the ['go mod edit' documentation](https://golang.org/cmd/go/#hdr-Edit_go_mod_from_tools_or_scripts) for more details.
* [github.com/rogpeppe/gohack](https://github.com/rogpeppe/gohack) makes these types of workflows much easier, especially if your goal is to have mutable checkouts of dependencies of a module.  See the [repository](https://github.com/rogpeppe/gohack) or the immediately prior FAQ for an overview.
* See the next FAQ for the details of using `replace` to work entirely outside of VCS.

### Can I work entirely outside of VCS on my local filesystem?

Yes. VCS is not required.

This is very simple if you have a single module you want to edit at a time outside of VCS (and you either have only one module in total, or if the other modules reside in VCS). In this case, you can place the file tree containing the single `go.mod` in a convenient location. Your `go build`, `go test` and similar commands will work even if your single module is outside of VCS (without requiring any use of `replace` in your `go.mod`).

If you want to have multiple inter-related modules on your local disk that you want to edit at the same time, then `replace` directives are one approach. Here is a sample `go.mod` that uses a `replace` with a relative path to point the `hello` module at the on-disk location of the `goodbye` module (without relying on any VCS):

```
module example.com/me/hello

require (
  example.com/me/goodbye v0.0.0
)

replace example.com/me/goodbye => ../goodbye
```
As shown in this example, if outside of VCS you can use `v0.0.0` as the version in the `require` directive. Note that as mentioned in the prior FAQ, the `require` directive is needed here. (`replace example.com/me/goodbye => ../goodbye` does not yet work without a corresponding `require example.com/me/goodbye v0.0.0`; this might change in the future with [#26241](https://golang.org/issue/26241)).

A small runnable example is shown in this [thread](https://groups.google.com/d/msg/golang-nuts/1nYoAMFZVVM/eppaRW2rCAAJ).

### How do I use vendoring with modules? Is vendoring going away?

The initial series of `vgo` blog posts did propose dropping vendoring entirely, but [feedback](https://groups.google.com/d/msg/golang-dev/FTMScX1fsYk/uEUSjBAHAwAJ) from the community resulted in retaining support for vendoring.

In brief, to use vendoring with modules:
* `go mod vendor` resets the main module's vendor directory to include all packages needed to build and test all of the module's packages based on the state of the go.mod files and Go source code.
* By default, go commands like `go build` ignore the vendor directory when in module mode.
* The `-mod=vendor` flag (e.g., `go build -mod=vendor`) instructs the go commands to use the main module's top-level vendor directory to satisfy dependencies. The go commands in this mode therefore ignore the dependency descriptions in go.mod and assume that the vendor directory holds the correct copies of dependencies. Note that only the main module's top-level vendor directory is used; vendor directories in other locations are still ignored.
* Some people will want to routinely opt-in to vendoring by setting a `GOFLAGS=-mod=vendor` environment variable.

Older versions of Go such as 1.10 understand how to consume a vendor directory created by `go mod vendor`, so vendoring is one way to provide dependencies to older versions of Go that do not fully understand modules.

If you are considering using vendoring, it is worthwhile to read the ["Modules and vendoring"](https://tip.golang.org/cmd/go/#hdr-Modules_and_vendoring) and ["Make vendored copy of dependencies"](https://tip.golang.org/cmd/go/#hdr-Make_vendored_copy_of_dependencies) sections of the tip documentation.

### Are there "always on" module repositories and enterprise proxies?

Publicly hosted "always on" immutable module repositories and optional privately hosted proxies and repositories are becoming available.

For example:
* [Project Athens](https://github.com/gomods/athens): Open source project in the works and looking for contributors.
* [JFrog Artifactory](https://jfrog.com/artifactory/): Commercial offering. Support for Go 1.11 modules started with release 5.11 as described [here](https://jfrog.com/blog/goproxy-artifactory-go-registries/) and [here](https://www.jfrog.com/confluence/display/RTF/Go+Registry). From Artifactory version 6.2.0, please use [JFrog CLI 1.20.2](https://www.jfrog.com/confluence/display/CLI/CLI+for+JFrog+Artifactory#CLIforJFrogArtifactory-BuildingGoPackages) and above.
* [THUMBAI](https://thumbai.app) - Open source project - Go Mod Proxy server and Go Vanity Import Path server
* [Goproxy China](https://github.com/aofei/goproxy.cn) - Open source project - A trusted Go module proxy located in China.

Note that you are not required to run a proxy. Rather, the go tooling in 1.11 has added optional proxy support via [GOPROXY](https://tip.golang.org/cmd/go/#hdr-Module_proxy_protocol) to enable more enterprise use cases (such as greater control), and also to better handle situations such as "GitHub is down" or people deleting GitHub repositories.

### Can I control when go.mod gets updated and when the go tools use the network to satisfy dependencies?

By default, a command like `go build` will reach out to the network as needed to satisfy imports.

Some teams will want to disallow the go tooling from touching the network at certain points, or will want greater control regarding when the go tooling updates `go.mod`, how dependencies are obtained, and how vendoring is used.

The go tooling provides a fair amount of flexibility to adjust or disable these default behaviors, including via `-mod=readonly`, `-mod=vendor`, `GOFLAGS`, `GOPROXY=off`, `GOPROXY=file:///filesystem/path`, `go mod vendor`, and `go mod download`.

The details on these options are spread throughout the official documentation. One community attempt at a consolidated overview of knobs related to these behaviors is [here](https://github.com/thepudds/go-module-knobs/blob/master/README.md), which includes links to the official documentation for more information.

### How do I use modules with CI systems such as Travis or CircleCI?

The simplest approach is likely just setting the environment variable `GO111MODULE=on`, which should work with most CI systems.

However, it can be valuable to run tests in CI on Go 1.11 with modules enabled as well as disabled, given some of your users will not have yet opted in to modules themselves. Vendoring is also a topic to consider.

The following two blog posts cover these topics more concretely:

* ["Using Go modules with vendor support on Travis CI"](https://arslan.io/2018/08/26/using-go-modules-with-vendor-support-on-travis-ci/) by Fatih Arslan
* ["Go Modules and CircleCI"](https://medium.com/@toddkeech/go-modules-and-circleci-c0d6fac0b000) by Todd Keech

## FAQs — go.mod and go.sum

### Why does 'go mod tidy' record indirect and test dependencies in my 'go.mod'?

The modules system records precise dependency requirements in your `go.mod`. (For more details, see the [go.mod concepts](https://github.com/golang/go/wiki/Modules#gomod) section above or the [go.mod tip documentation](https://tip.golang.org/cmd/go/#hdr-The_go_mod_file)).

`go mod tidy` updates your current `go.mod` to include the dependencies needed for tests in your module — if a test fails, we must know which dependencies were used in order to reproduce the failure.

`go mod tidy` also ensures your current `go.mod` reflects the dependency requirements for all possible combinations of OS, architecture, and build tags (as described [here](https://github.com/golang/go/issues/25971#issuecomment-399091682)). In contrast, other commands like `go build` and `go test` only update `go.mod` to provide the packages imported by the requested packages under the current `GOOS`, `GOARCH`, and build tags (which is one reason `go mod tidy` might add requirements that were not added by `go build` or similar).

If a dependency of your module does not itself have a `go.mod` (e.g., because the dependency has not yet opted in to modules itself), or if its `go.mod` file is missing one or more of its dependencies (e.g., because the module author did not run `go mod tidy`), then the missing transitive dependencies will be added to _your_ module's requirements, along with an `// indirect` comment to indicate that the dependency is not from a direct import within your module.

Note that this also means that any missing test dependencies from your direct or indirect dependencies will also be recorded in your `go.mod`. (An example of when this is important: `go test all` runs the tests of _all_ direct and indirect dependencies of your module, which is one way to validate that your current combination of versions work together. If a test fails in one of your dependencies when you run `go test all`, it is important to have a complete set of test dependency information recorded so that you have reproducible `go test all` behavior).

Another reason you might have `// indirect` dependencies in your `go.mod` file is if you have upgraded (or downgraded) one of your indirect dependencies beyond what is required by your direct dependencies, such as if you ran `go get -u` or `go get foo@1.2.3`. The go tooling needs a place to record those new versions, and it does so in your `go.mod` file (and it does not reach down into your dependencies to modify _their_ `go.mod` files).

In general, the behaviors described above are part of how modules provide 100% reproducible builds and tests by recording precise dependency information.

If you are curious as to why a particular module is showing up in your `go.mod`, you can run `go mod why -m <module>` to [answer](https://tip.golang.org/cmd/go/#hdr-Explain_why_packages_or_modules_are_needed) that question.  Other useful tools for inspecting requirements and versions include `go mod graph` and `go list -m all`.

### Is 'go.sum' a lock file? Why does 'go.sum' include information for module versions I am no longer using?

No, `go.sum` is not a lock file. The `go.mod` files in a build provide enough information for 100% reproducible builds.

For validation purposes, `go.sum` contains the expected cryptographic checksums of the content of specific module versions. See the [FAQ below](https://github.com/golang/go/wiki/Modules#should-i-commit-my-gosum-file-as-well-as-my-gomod-file) for more details on `go.sum` (including why you typically should check in `go.sum`) as well as the ["Module downloading and verification"](https://tip.golang.org/cmd/go/#hdr-Module_downloading_and_verification) section in the tip documentation.

In part because `go.sum` is not a lock file, it retains cryptographic checksums for module versions even after you stop using a module or particular module version. This allows validation of the checksums if you later resume using something, which provides additional safety.

In addition, your module's `go.sum` records checksums for all direct and indirect dependencies used in a build (and hence your `go.sum` will frequently have more modules listed than your `go.mod`).

### Should I commit my 'go.sum' file as well as my 'go.mod' file?

Typically your module's `go.sum` file should be committed along with your `go.mod` file.

* `go.sum` contains the expected cryptographic checksums of the content of specific module versions.
* If someone clones your repository and downloads your dependencies using the go command, they will receive an error if there is any mismatch between their downloaded copies of your dependencies and the corresponding entries in your `go.sum`.
* In addition, `go mod verify` checks that the on-disk cached copies of module downloads still match the entries in `go.sum`.
* Note that `go.sum` is not a lock file as used in some alternative dependency management systems. (`go.mod` provides enough information for reproducible builds).
* See very brief [rationale here](https://twitter.com/FiloSottile/status/1029404663358087173) from Filippo Valsorda on why you should check in your `go.sum`. See the ["Module downloading and verification"](https://tip.golang.org/cmd/go/#hdr-Module_downloading_and_verification) section of the tip documentation for more details. See possible future extensions being discussed for example in [#24117](https://github.com/golang/go/issues/24117) and [#25530](https://github.com/golang/go/issues/25530).

### Should I still add a 'go.mod' file if I do not have any dependencies?

Yes. This supports working outside of GOPATH, helps communicate to the ecosystem that you are opting in to modules, and in addition the `module` directive in your `go.mod` serves as a definitive declaration of the identity of your code (which is one reason why import comments might eventually be deprecated). Of course, modules are purely an opt-in capability in Go 1.11.

## FAQs — Semantic Import Versioning

### Why must major version numbers appear in import paths?

Please see the discussion on the Semantic Import Versioning and the import compatibility rule in the ["Semantic Import Versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) concepts section above. See also the [blog post announcing the proposal](https://blog.golang.org/versioning-proposal), which talks more about the motivation and justification for the import compatibility rule.

### Why are major versions v0, v1 omitted from import paths?"

Please see the question "Why are major versions v0, v1 omitted from import paths?" in the earlier [FAQ from the official proposal discussion](https://github.com/golang/go/issues/24301#issuecomment-371228664).

### What are some implications of tagging my project with major version v0, v1, or making breaking changes with v2+?

In response to a comment about *"k8s does minor releases but changes the Go API in each minor release"*, Russ Cox made the following [response](https://github.com/kubernetes/kubernetes/pull/65683#issuecomment-403705882) that highlights some implications for picking v0, v1, vs. frequently making breaking changes with v2, v3, v4, etc. with your project:

>  I don't fully understand the k8s dev cycle etc, but I think generally the k8s team needs to decide/confirm what they intend to guarantee to users about stability and then apply version numbers accordingly to express that.
>
> * To make a promise about API compatibility (which seems like the best user experience!) then start doing that and use 1.X.Y.
> * To have the flexibility to make backwards-incompatible changes in every release but allow different parts of a large program to upgrade their code on different schedules, meaning different parts can use different major versions of the API in one program, then use X.Y.0, along with import paths like k8s.io/client/vX/foo.
> * To make no promises about API compatible and also require every build to have only one copy of the k8s libraries no matter what, with the implied forcing of all parts of a build to use the same version even if not all of them are ready for it, then use 0.X.Y.

On a related note, Kubernetes has some atypical build approaches (currently including custom wrapper scripts on top of godep), and hence Kubernetes is an imperfect example for many other projects, but it will likely be an interesting example as [Kubernetes moves towards adopting Go 1.11 modules](https://github.com/kubernetes/kubernetes/pull/64731#issuecomment-407345841).

### Can a module consume a package that has not opted in to modules?

Yes.

If a repository has not opted in to modules but has been tagged with valid [semver](https://semver.org) tags (including the required leading `v`), then those semver tags can be used in a `go get`, and a corresponding semver version will be record in the importing module's `go.mod` file. If the repository does not have any valid semver tags, then the repository's version will be recorded with a ["pseudo-version"](https://golang.org/cmd/go/#hdr-Pseudo_versions) such as ` v0.0.0-20171006230638-a6e239ea1c69` (which includes a timestamp and a commit hash, and which are designed to allow a total ordering across versions recored in `go.mod` and to make it easier to reason about which recorded versions are "later" than another recorded version).

For example, if the latest version of package `foo` is tagged `v1.2.3` but `foo` has not itself opted in to modules, then running `go get foo` or `go get foo@v1.2.3` from inside module M will be recorded in module M's `go.mod` file as:

```
require  foo  v1.2.3
```

The `go` tool will also use available semver tags for a non-module package in additional workflows (such as `go list -u=patch`, which upgrades the dependencies of a module to available patch releases, or `go list -u -m all`, which shows available upgrades, etc.).

Please see the next FAQs for additional details related to v2+ packages that have not opted in to modules.

### Can a module consume a v2+ package that has not opted into modules? What does '+incompatible' mean?

Yes, a module can import a v2+ package that has not opted into modules, and if the imported v2+ package has a valid [semver](https://semver.org) tag, it will be recorded with an `+incompatible` suffix.

**Additional Details**

Please be familiar with the material in the ["Semantic Import Versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) section above.

It is helpful to first review some core principles that are generally useful but particularly important to keep in mind when thinking about the behavior described in this FAQ.

The following core principles are _always_ true when the `go` tool is operating in module mode (e.g., `GO111MODULE=on`):

1. A package's import path defines the identity of the package.
   * Packages with _different_ import paths are treated as _different_ packages.
   * Packages with the _same_ import path are treated as the _same_ package (and this is true _even if_ the VCS tags say the packages have different major versions).
2. An import path without a `/vN` is treated as a v1 or v0 module (and this is true _even if_ the imported package has not opted in to modules and has VCS tags that say the major version is greater than 1).
3. The module path (such as `module foo/v2`) declared at the start of a module's `go.mod` file is both:
   * the definitive declaration of that module's identity
   * the definitive declaration of how that module must be imported by consuming code

As we will see in the next FAQ, these principles are not always true when the `go` tool is _not_ in module mode, but these principles are always true when the `go` tool _is_ in module mode.

In short, the `+incompatible` suffix indicates that principle 2 above is in effect when the following are true:
* an imported package has not opted in to modules, and
* its VCS tags say the major version is greater than 1, and
* principle 2 is overriding the VCS tags – the import path without a `/vN` is treated as a v1 or v0 module (even though the VCS tags say otherwise)

When the `go` tool is in module mode, it will assume a non-module v2+ package has no awareness of Semantic Import Versioning and treat it as an (incompatible) extension of the v1 version series of the package (and the `+incompatible` suffix is an indication that the `go` tool is doing so).

**Example**

Suppose:
* `oldpackage` is a package that predates the introduction of modules
* `oldpackage` has never opted in to modules (and hence does not have a `go.mod` itself)
* `oldpackage` has a valid semver tag `v3.0.1`, which is its latest tag

In this case, running for example `go get oldpackage@latest` from inside module M will record the following in module M's `go.mod` file:

```
require  oldpackage  v3.0.1+incompatible
```

Note that there is no `/v3` used at the end of `oldpackage` in the `go get` command above or in the recorded `require` directive – using `/vN` in module paths and import paths is a feature of [Semantic Import Versioning](https://github.com/golang/go/wiki/Modules#semantic-import-versioning), and `oldpackage` has not signaled its acceptance and understanding of Semantic Import Versioning given `oldpackage` has not opted into modules by having a `go.mod` file within `oldpackage` itself. In other words, even though `oldpackage` has a [semver](https://semver.org) tag of `v3.0.1`, `oldpackage` is not granted the rights and responsibilities of [Semantic Import Versioning](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) (such as using `/vN` in import paths) because `oldpackage` has not yet stated its desire to do so.

The `+incompatible` suffix indicates that the `v3.0.1` version of `oldpackage` has not actively opted in to modules, and hence the `v3.0.1` version of `oldpackage` is assumed to _not_ understand Semantic Import Versioning or how to use major versions in import paths. Therefore, when operating in [module mode](https://github.com/golang/go/wiki/Modules#when-do-i-get-old-behavior-vs-new-module-based-behavior), the `go` tool will treat the non-module `v3.0.1` version of `oldpackage` as an (incompatible) extension of the v1 version series of `oldpackage` and assume that the `v3.0.1` version of `oldpackage` has no awareness of Semantic Import Versioning, and the `+incompatible` suffix is an indication that the `go` tool is doing so.

The fact that the the `v3.0.1` version of `oldpackage` is considered to be part of the v1 release series according to Semantic Import Versioning means for example that versions `v1.0.0`, `v2.0.0`, and `v3.0.1` are all always imported using the same import path:

```
import  "oldpackage"
```

Note again that there is no `/v3` used at the end of `oldpackage`.

In general, packages with different import paths are different packages. In this example, given versions `v1.0.0`, `v2.0.0`, and `v3.0.1` of `oldpackage` would all be imported using the same import path, they are therefore treated by a build as the same package (again because `oldpackage` has not yet opted in to Semantic Import Versioning), with a single copy of `oldpackage` ending up in any given build. (The version used will be the semantically highest of the versions listed in any `require` directives; see ["Version Selection"](https://github.com/golang/go/wiki/Modules#version-selection)).

If we suppose that later a new `v4.0.0` release of `oldpackage` is created that adopts modules and hence contains a `go.mod` file, that is the signal that `oldpackage` now understands the rights and responsibilities of Semantic Import Versioning, and hence a module-based consumer would now import using `/v4` in the import path:

```
import  "oldpackage/v4"
```

and the version would be recorded as:

```
require  oldpackage/v4  v4.0.0
```

`oldpackage/v4` is now a different import path than `oldpackage`, and hence a different package.  Two copies (one for each import path) would end up in a module-aware build if some consumers in the build have `import "oldpackage/v4"` while other consumers in the same build have `import "oldpackage"`. This is desirable as part of the strategy to allow gradual adoption of modules. In addition, even after modules are out of their current transitional phase, this behavior is also desirable to allow gradual code evolution over time with different consumers upgrading at different rates to newer versions (e.g., allowing different consumers in a large build to choose to upgrade at different rates from `oldpackage/v4` to some future `oldpackage/v5`).

### How are v2+ modules treated in a build if modules support is not enabled? How does "minimal module compatibility" work in 1.9.7+, 1.10.3+, and 1.11?

When considering older Go versions or Go code that has not yet opted in to modules, Semantic Import Versioning has significant backwards compatibility implications related to v2+ modules.

As described in the ["Semantic Import Versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) section above:
* a module that is version v2 or higher must include a `/vN` in its own module path declared in its `go.mod`.
* a module-based consumer (that is, code that has opted in to modules) must include a `/vN` in the import path to import a v2+ module.

However, the ecosystem is expected to proceed at varying paces of adoption for modules and Semantic Import Versioning.

As described in more detail in the ["How to Release a v2+ Module"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) section, in the "Major Subdirectory" approach, the author of a v2+ module creates subdirectories such as `mymodule/v2` or `mymodule/v3` and moves or copies the approriate packages underneath those subdirectories. This means the traditional import path logic (even in older Go releases such as Go 1.8 or 1.7) will find the appropriate packages upon seeing an import statement such as `import "mymodule/v2/mypkg"`. Hence, packages residing in a "Major Subdirectory" v2+ module will be found and used even if modules support is not enabled (whether that is because you are running Go 1.11 and have not enabled modules, or because you are running a older version like Go 1.7, 1.8, 1.9 or 1.10 that does not have full module support).  Please see the ["How to Release a v2+ Module"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) section for more details on the "Major Subdirectory" approach.

The remainder of this FAQ is focused on the "Major Branch" approach described in the ["How to Release a v2+ Module"](https://github.com/golang/go/wiki/Modules#releasing-modules-v2-or-higher) section. In the "Major Branch" approach, no `/vN` subdirectories are created and instead the module version information is communicated by the `go.mod` file and by applying semver tags to commits (which often will be on `master`, but could be on different branches).

In order to help during the current transitional period, "minimal module compatibility" was [introduced](https://go-review.googlesource.com/c/go/+/109340) to Go 1.11 to provide greater compatibility for Go code that has not yet opted in to modules, and that "minimal module compatibility" was also backported to Go 1.9.7 and 1.10.3 (where those versions are effectively always operating with full module mode disabled given those older Go versions do not have full module support).

The primary goals of "minimal module compatibility" are:

1. Allow older Go versions 1.9.7+ and 1.10.3+ to be able to more easily compile modules that are using Semantic Import Versioning with `/vN` in import paths, and provide that same behavior when [module mode](https://github.com/golang/go/wiki/Modules#when-do-i-get-old-behavior-vs-new-module-based-behavior) is disabled in Go 1.11.

2. Allow old code to be able to consume a v2+ module without requiring that old consumer code to immediately change to using a new `/vN` import path when consuming a v2+ module.

3. Do so without relying on the module author to create `/vN` subdirectories.

**Additional Details – "Minimal Module Compatibility"**

"Minimal module compatibility" only takes effect when full [module mode](https://github.com/golang/go/wiki/Modules#when-do-i-get-old-behavior-vs-new-module-based-behavior) is disabled for the `go` tool, such as if you have set `GO111MODULE=off` in Go 1.11, or are using Go versions 1.9.7+ or 1.10.3+.

When a v2+ module author has _not_ created `/v2` or `/vN` subdirectories and you are instead relying on the "minimal module compatibility" mechanism in Go 1.9.7+, 1.10.3+ and 1.11:

* A package that has _not_ opted in to modules would _not_ include the major version in the import path for any imported v2+ modules.
* In contrast, a package that _has_ opted in to modules _must_ include the major version in the import path to import any v2+ modules.
  * If a package has opted in to modules, but does not include the major version in the import path when importing a v2+ modules, it will not import a v2+ version of that module when the `go` tool is operating in full module mode. (A package that has opted in to modules is assumed to "speak" Semantic Import Versioning. If `foo` is a module with v2+ versions, then under Semantic Import Versioning saying `import "foo"` means import the v1 Semantic Import Versioning series of `foo`).
* The mechanism used to implement "minimal module compatibility" is intentionally very narrow:
  * The entirety of the logic is – when operating in GOPATH mode, an unresolvable import statement containing a `/vN` will be tried again after removing the `/vN` if the import statement is inside code that has opted in to modules (that is, import statements in `.go` files within a tree with a valid `go.mod` file).
  * The net effect is that an import statement such as `import "foo/v2"` within code that lives inside of a module will still compile correctly in GOPATH mode in 1.9.7+, 1.10.3+ and 1.11, and it will resolve as if it said `import "foo"` (without the `/v2`), which means it will use the version of `foo` that resides in your GOPATH without being confused by the extra `/v2`.
  * "Minimal module compatibility" does not affect anything else, including it does not the affect paths used in the `go` command line (such as arguments to `go get` or `go list`).
* This transitional "minimal module awareness" mechanism purposefully breaks the rule of "packages with different import paths are treated as different packages" in pursuit a very specific backwards compatibility goal – to allow old code to compile unmodified when it is consuming a v2+ module. In slightly more detail:
  * It would be a more burdensome for the overall ecosystem if the only way for old code to consume a v2+ module was to first change the old code.
  * If we are not modifying old code, then that old code must work with pre-module import paths for v2+ modules.
  * On the other hand, new or updated code opting in to modules must use the new `/vN` import for v2+ modules.
  * The new import path is not equal to old import path, yet both are allowed to work in a single build, and therefore we have two different functioning import paths that resolve to the same package.
  * For example, when operating in GOPATH mode, `import "foo/v2"` appearing in module-based code resolves to the same code residing in your GOPATH as `import "foo"`, and the build ends up with one copy of `foo` – in particular, whatever version is on disk in GOPATH. This allows module-based code with  `import "foo/v2"` to compile even in GOPATH mode in 1.9.7+, 1.10.3+ and 1.11.
* In contrast, when the `go` tool is operating in full module mode:
   * There are no exceptions to the rule "packages with different import paths are different packages" (including vendoring has been refined in full module mode to also adhere to this rule).
   * For example, if the `go` tool is in full module mode and `foo` is a v2+ module, then `import "foo"` is asking for a v1 version of `foo` vs. `import "foo/v2"` is asking for a v2 version of `foo`.

## FAQS — Multi-Module Repositories

### What are multi-module repositories?

A multi-module repository is a repository that contains multiple modules, each with its own go.mod file. Each module starts at the directory containing its go.mod file, and contains all packages from that directory and its subdirectories recursively, excluding any subtree that contains another go.mod file.

Each module has its own version information. Version tags for modules below the root of the repository must include the relative directory as a prefix. For example, consider the following repository:

```
my-repo
|____foo
| |____rop
| | |____go.mod
```

The tag for version 1.2.3 of module "my-repo/foo/rop" is "foo/rop/v1.2.3".

Typically, the path for one module in the repository will be a prefix of the others. For example, consider this repository:

```
my-repo
|____go.mod
|____bar
|____foo
| |____rop
| |____yut
|____mig
| |____go.mod
| |____vub
```
![Fig. A top-level module's path is a prefix of another module's path.](https://github.com/jadekler/module-testing/blob/master/imagery/multi_module_repo.png)

_Fig. A top-level module's path is a prefix of another module's path._

This repository contains two modules. However, the module "my-repo" is a prefix of the path of the module "my-repo/mig".

Adding modules, removing modules, and versioning modules in such a configuration require considerable care and deliberation, so it is almost always easier and simpler to manage a single-module repository rather than multiple modules in an existing repository.

### Is it possible to add a module to a multi-module repository?

Yes. However, there are two classes of this problem:

The first class: the package to which the module is being added to is not in version control yet (a new package). This case is straightforward: add the package and the go.mod in the same commit, tag the commit, and push.

The second class: the path at which the module is being added is in version control and contains one or more existing packages. This case requires a considerable amount of care. To illustrate, consider again the following repository (now in a github.com location to simulate the real-world better):

```
github.com/my-repo
|____go.mod
|____bar
|____foo
| |____rop
| |____yut
|____mig
| |____vub
```

Consider adding module "github.com/my-repo/mig". If one were to follow the same approach as above, the package /my-repo/mig could be provided by two different modules: the old version of "github.com/my-repo", and the new, standalone module "github.com/my-repo/mig. If both modules are active, importing "github.com/my-repo/mig" would cause an “ambiguous import” error at compile time.

The way to get around this is to make the newly-added module depend on the module it was "carved out" from, at a version after which it was carved out.

Let's step through this with the above repository, assuming that "github.com/my-repo" is currently at v1.2.3:

1. Add github.com/my-repo/mig/go.mod:

    ```
    cd path-to/github.com/my-repo/mig
    go mod init github.com/my-repo/mig

    # Note: if "my-repo/mig" does not actually depend on "my-repo", add a blank
    # import.
    # Note: version must be at or after the carve-out.
    go mod edit -require github.com/myrepo@v1.3
    ```

1. `git commit`
1. `git tag v1.3.0`
1. `git tag mig/v1.0.0`
1. Next, let's test these. We can't `go build` or `go test` naively, since the go commands would try to fetch each dependent module from the module cache. So, we need to use replace rules to cause `go` commands to use the local copies:

    ```
    cd path-to/github.com/my-repo/mig
    go mod edit -replace github.com/my-repo@v1.3.0=../
    go test ./...
    go mod edit -dropreplace github.com/my-repo@v1.3.0
    ```

1. `git push origin master v1.2.4 mig/v1.0.0` push the commit and both tags

Note that in the future [golang.org/issue/28835](https://github.com/golang/go/issues/28835) should make the testing step a more straightforward experience.

Note also that code was removed from module "github.com/my-repo" between minor versions. It may seem strange to not consider this a major change, but in this instance the transitive dependencies continue to provide compatible implementations of the removed packages at their original import paths.

### Is it possible to remove a module from a multi-module repository?

Yes, with the same two cases and similar steps as above.

### Can a module depend on an internal/ in another?

Yes. Packages in one module are allowed to import internal packages from another module as long as they share the same path prefix up to the internal/ path component. For example, consider the following repository:

```
my-repo
|____go.mod
|____internal
|____foo
| |____go.mod
```

Here, package foo can import /my-repo/internal as long as module "my-repo/foo" depends on module "my-repo". Similarly, in the following repository:

```
my-repo
|____internal
| |____go.mod
|____foo
| |____go.mod
```

Here, package foo can import my-repo/internal as long as module "my-repo/foo" depends on module "my-repo/internal". The semantics are the same in both: since my-repo is a shared path prefix between my-repo/internal and my-repo/foo, package foo is allowed to import package internal.

## FAQs — Minimal Version Selection

### Won't minimal version selection keep developers from getting important updates?

Please see the question "Won't minimal version selection keep developers from getting important updates?" in the earlier [FAQ from the official proposal discussion](https://github.com/golang/go/issues/24301#issuecomment-371228664).

## FAQs — Possible Problems

### What are some general things I can spot check if I am seeing a problem?

* Double-check that modules are enabled by running `go env` to confirm it does not show an empty value for the read-only `GOMOD` variable.
   * Note: you never set `GOMOD` as a variable because it is effectively read-only debug output that `go env` outputs.
   * If you are setting `GO111MODULE=on` to enable modules, double-check that it is not accidentally the plural `GO111MODULES=on`. (People sometimes naturally include the `S` because the feature is often called "modules").
* If vendoring is expected to be used, double-check check that the `-mod=vendor` flag is being passed to `go build `or similar, or that `GOFLAGS=-mod=vendor` is set.
   * Modules by default ignore the `vendor` directory unless you ask the `go` tool to use `vendor`.
* It is frequently helpful to check `go list -m all` to see the list of actual versions selected for your build
  * `go list -m all` usually gives you more detail compared to if you were to instead just look a `go.mod` file.
* If running `go get foo` fails in some way, or if `go build` is failing on a particular package `foo`, it often can be helpful to check the output from `go get -v foo` or `go get -v -x foo`:
  * In general, `go get` will often provide more a detailed error message than `go build`.
  * The `-v` flag to `go get` asks to print more verbose details, though be mindful that certain "errors" such as 404 errors _might_ be expected based on how a remote repository was configured.
  * If the nature of the problem is still not clear, you can also try the more verbose `go get -v -x foo`, which also shows the git or other VCS commands being issued.  (If warranted, you can often execute the same git commands outside of the context of the `go` tool for troubleshooting purposes).
* You can check to see if you are using a particularly old git version
  * Older versions of git were a common source of problems for the `vgo` prototype and Go 1.11 beta, but much less frequently in the GA 1.11.
* The module cache in Go 1.11 can sometimes cause various errors, primarily if there were previously network issues or multiple `go` commands executing in parallel (see [#26794](https://github.com/golang/go/issues/26794), which is addressed for Go 1.12).  As a troubleshooting step, you can copy $GOPATH/pkg/mod to a backup directory (in case further investigation is warranted later), run `go clean -modcache`, and then see whether the original problem persists.
* If you are using Docker, it can be helpful to check if you can reproduce the behavior outside of Docker (and if the behavior only occurs in Docker, the list of bullets above can be used as a starting point to compare results between inside Docker vs. outside).

The error you are currently examining might be a secondary issue caused by not having the expected version of a particular module or package in your build. Therefore, if the cause of a particular error is not obvious, it can be helpful to spot check your versions as described in the next FAQ.

### What can I check if I am not seeing the expected version of a dependency?

1. A good first step is to run `go mod tidy`. There is some chance this might resolve the issue, but it will also help put your `go.mod` file into a consistent state with respect to your `.go` source code, which will help make any subsequent investigation easier.

2. The second step usually should be to check `go list -m all` to see the list of actual versions selected for your build.  `go list -m all` shows you the final selected versions, including for indirect dependencies and after resolving versions for any shared dependencies. It also shows the outcome of any `replace` and `exclude` directives.

3. A good next step can be to examine the output of `go mod graph` or `go mod graph | grep <module-of-interest>`.  `go mod graph` prints the module requirement graph (including taking into account replacements). Each line in the output has two fields: the first column is a consuming module, and the second column is one of that module's requirements (including the version required by that consuming module).  This can be a quick way to see which modules are requiring a particular dependency, including when your build has a dependency that has different required versions from different consumers in your build (and if that is the case, it is important to be familiar with the behavior described in the ["Version Selection"](https://github.com/golang/go/wiki/Modules#version-selection) section above).

`go mod why -m <module>` can also be useful here, although it is typically more useful for seeing why a dependency is included at all (rather than why a dependency ends up with a particular version).

`go list` provides many more variations of queries that can be useful to interrogate your modules if needed. One example is the following, which will show the exact versions used in your build excluding test-only dependencies:
```
go list -deps -f '{{with .Module}}{{.Path}} {{.Version}}{{end}}' ./... | sort -u
```

A more detailed set of commands and examples for interrogating your modules can be seen in a runnable "Go Modules by Example" [walkthough](https://github.com/go-modules-by-example/index/tree/master/018_go_list_mod_graph_why).

One cause of unexpected versions can be due to someone having created an invalid or unexpected `go.mod` file that was not intended, or a related mistake (for example: a `v2.0.1` version of module might have incorrectly declared itself to be `module foo` in its `go.mod` without the required `/v2`; an import statement in `.go` code intended to import a v3 module might be be missing the required `/v3`; a `require` statement in a `go.mod` for a v4 module might be be missing the required `/v4`). Therefore, if the cause of a particular issue you are seeing is not obvious, it can be worthwhile to first re-read the material in the ["go.mod"](https://github.com/golang/go/wiki/Modules#gomod) and ["Semantic Import Versioning"](https://github.com/golang/go/wiki/Modules#semantic-import-versioning) sections above (given these include important rules that modules must follow) and then take a few minutes to spot check the most relevant `go.mod` files and import statements.

### Why am I getting an error 'cannot find module providing package foo'?

This is a general error message that can occur for several different underlying causes.

In some cases, this error is simply due to a mistyped path, so the first step likely should be to double-check for incorrect paths based on the details listed in the error message.

If you have not already done so, a good next step is often to try `go get -v foo` or `go get -v -x foo`:
* In general, `go get` will often provide more a detailed error message than `go build`.
* See the first troubleshooting FAQ in this section [above](https://github.com/golang/go/wiki/Modules#what-are-some-general-things-i-can-spot-check-if-i-am-seeing-a-problem) for more details.

Some other possible causes:

* You might see the error `cannot find module providing package foo` if you have issued `go build` or `go build .` but do not have any `.go` source files in the current directory. If this is what you are encountering, the solution might be an alternative invocation such as `go build ./...` (where the `./...` expands out to match all the packages within the current module). See [#27122](https://github.com/golang/go/issues/27122).

* The module cache in Go 1.11 can cause this error, including in the face of network issues or multiple `go` commands executing in parallel. This is resolved in Go 1.12. See the first troubleshooting FAQ in this section [above](https://github.com/golang/go/wiki/Modules#what-are-some-general-things-i-can-spot-check-if-i-am-seeing-a-problem) for more details and possible corrective steps.

### Why does 'go mod init' give the error 'cannot determine module path for source directory'?

`go mod init` without any arguments will attempt to guess the proper module path based on different hints such as VCS meta data. However, it is not expected that `go mod init` will always be able to guess the proper module path.

If `go mod init` gives you this error, those heuristics were not able to guess, and you must supply the module path yourself (such as `go mod init github.com/you/hello`).

### I have a problem with a complex dependency that has not opted in to modules. Can I use information from its current dependency manager?

Yes. This requires some manual steps, but can be helpful in some more complex cases.

When you run `go mod init` when initializing your own module, it will automatically convert from a prior dependency manager by translating configuration files like `Gopkg.lock`, `glide.lock`, or `vendor.json` into a `go.mod` file that contains corresponding `require` directives. The information in a pre-existing `Gopkg.lock` file for example usually describes version information for all of your direct and indirect dependencies.

However, if instead you are adding a new dependency that has not yet opted in to modules itself, there is not a similar automatic conversion process from any prior dependency manager that your new dependency might have been using. If that new dependency itself has non-module dependencies that have had breaking changes, then in some cases that can cause incompatibility problems. In other words, a prior dependency manager of your new dependency is not automatically used, and that can cause problems with your indirect dependencies in some cases.

One approach is to run `go mod init` on your problematic non-module direct dependency to convert from its current dependency manager, and then use the `require` directives from the resulting temporary `go.mod` to populate or update the `go.mod` in your module.

For example, if `github.com/some/nonmodule` is a problematic direct dependency of your module that is currently using another dependency manager, you can do something similar to:

```
$ git clone -b v1.2.3 https://github.com/some/nonmodule /tmp/scratchpad/nonmodule
$ cd /tmp/scratchpad/nonmodule
$ go mod init
$ cat go.mod
```

The resulting `require` information from the temporary `go.mod` can be manually moved into the actual `go.mod` for your module, or you can consider using https://github.com/rogpeppe/gomodmerge, which is a community tool targeting this use case. In addition, you will want to add a `require github.com/some/nonmodule v1.2.3` to your actual `go.mod` to match the version that you manually cloned.

A concrete example of following this technique for docker is in this [#28489 comment](https://github.com/golang/go/issues/28489#issuecomment-454795390), which illustrates getting a consistent set of versions
of docker dependencies to avoid case sensitive issues between `github.com/sirupsen/logrus` vs. `github.com/Sirupsen/logrus`.

### Why does 'go build' require gcc, and why are prebuilt packages such as net/http not used?

In short:

> Because the pre-built packages are non-module builds and can’t be reused. Sorry. Disable cgo for now or install gcc.

This is only an issue when opting in to modules (e.g., via `GO111MODULE=on`). See [#26988](https://github.com/golang/go/issues/26988#issuecomment-417886417) for additional discussion.

### Do modules work with relative imports like `import "./subdir"`?

No. See [#26645](https://github.com/golang/go/issues/26645#issuecomment-408572701), which includes:

> In modules, there finally is a name for the subdirectory. If the parent directory says "module m" then the subdirectory is imported as "m/subdir", no longer "./subdir".

### Some needed files may not be present in populated vendor directory

Directories without `.go` files are not copied inside the `vendor` directory by `go mod vendor`. This is by design.

In short, setting aside any particular vendoring behavior – the overall model for go builds is that the files needed to build a package should be in the directory with the `.go` files.

Using the example of cgo – modifying C source code in other directories will not trigger a rebuild, and instead your build will use stale cache entries. The cgo documentation now [includes](https://go-review.googlesource.com/c/go/+/125297/5/src/cmd/cgo/doc.go):

> Note that changes to files in other directories do not cause the package
to be recompiled, so _all non-Go source code for the package should be
stored in the package directory_, not in subdirectories.

A community tool https://github.com/goware/modvendor allows you to easily copy a complete set of .c, .h, .s, .proto or other files from a module into the `vendor` director. Although this can be helpful, some care must be taken to make sure your go build is being handled properly in general (regardless of vendoring) if you have files needed to build a package that are outside of the directory with the `.go` files.

See additional discussion in [#26366](https://github.com/golang/go/issues/26366#issuecomment-405683150).

An alternative approach to traditional vendoring is to check in the module cache. It can end up with similar benefits as traditional vendoring and in some ways ends up with a higher fidelity copy. This approach is explained as a "Go Modules by Example" [walkthrough](https://github.com/go-modules-by-example/index/blob/master/012_modvendor/README.md).