# Robotgo

<!--<img align="right" src="https://raw.githubusercontent.com/cauefcr/robotgo/master/logo.jpg">-->
<!--[![Build Status](https://travis-ci.org/cauefcr/robotgo.svg)](https://travis-ci.org/cauefcr/robotgo)
[![codecov](https://codecov.io/gh/cauefcr/robotgo/branch/master/graph/badge.svg)](https://codecov.io/gh/cauefcr/robotgo)-->
<!--<a href="https://circleci.com/gh/cauefcr/robotgo/tree/dev"><img src="https://img.shields.io/circleci/project/cauefcr/robotgo/dev.svg" alt="Build Status"></a>-->
[![Build Status](https://github.com/cauefcr/robotgo/workflows/Go/badge.svg)](https://github.com/cauefcr/robotgo/commits/master)
[![CircleCI Status](https://circleci.com/gh/cauefcr/robotgo.svg?style=shield)](https://circleci.com/gh/cauefcr/robotgo)
[![Build Status](https://travis-ci.org/cauefcr/robotgo.svg)](https://travis-ci.org/cauefcr/robotgo)
![Appveyor](https://ci.appveyor.com/api/projects/status/github/cauefcr/robotgo?branch=master&svg=true)
[![Go Report Card](https://goreportcard.com/badge/github.com/cauefcr/robotgo)](https://goreportcard.com/report/github.com/cauefcr/robotgo)
[![GoDoc](https://godoc.org/github.com/cauefcr/robotgo?status.svg)](https://godoc.org/github.com/cauefcr/robotgo)
[![GitHub release](https://img.shields.io/github/release/cauefcr/robotgo.svg)](https://github.com/cauefcr/robotgo/releases/latest)
[![Join the chat at https://gitter.im/cauefcr/robotgo](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/cauefcr/robotgo?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
<!-- [![Release](https://github-release-version.herokuapp.com/github/cauefcr/robotgo/release.svg?style=flat)](https://github.com/cauefcr/robotgo/releases/latest) -->
<!-- <a href="https://github.com/cauefcr/robotgo/releases"><img src="https://img.shields.io/badge/%20version%20-%206.0.0%20-blue.svg?style=flat-square" alt="Releases"></a> -->

  >Golang Desktop Automation. Control the mouse, keyboard, bitmap, read the screen,   Window Handle and global event listener.

RobotGo supports Mac, Windows, and Linux(X11).

[Chinese Simplified](https://github.com/cauefcr/robotgo/blob/master/README_zh.md)

## Contents
- [Docs](#docs)
- [Binding](#binding)
- [Requirements](#requirements)
- [Installation](#installation)
- [Update](#update)
- [Examples](#examples)
- [Cross-Compiling](#crosscompiling)
- [Authors](#authors)
- [Plans](#plans)
- [Donate](#donate)
- [Contributors](#contributors)
- [License](#license)

## Docs
  - [GoDoc](https://godoc.org/github.com/cauefcr/robotgo)
  - [API Docs](https://github.com/cauefcr/robotgo/blob/master/docs/doc.md) &nbsp;&nbsp;&nbsp;
  - [Chinese Docs](https://github.com/cauefcr/robotgo/blob/master/docs/doc_zh.md)

## Binding:

[Robotn](https://github.com/vcaesar/robotn), binding JavaScript and other, support more language.

## Requirements:

Now, Please make sure `Golang, GCC` is installed correctly before installing RobotGo.

### ALL:
```
Golang

GCC
```

#### For Mac OS X:
```
Xcode Command Line Tools
```

#### For Windows:
```
MinGW-w64 (Use recommended) or other GCC
```

#### For everything else:

```
GCC, libpng

X11 with the XTest extension (also known as the Xtst library)

Event:

xcb, xkb, libxkbcommon
```

##### Ubuntu:

```yml
sudo apt-get install gcc libc6-dev

sudo apt-get install libx11-dev xorg-dev libxtst-dev libpng++-dev

sudo apt-get install xcb libxcb-xkb-dev x11-xkb-utils libx11-xcb-dev libxkbcommon-x11-dev
sudo apt-get install libxkbcommon-dev

sudo apt-get install xsel xclip
```

#### Fedora:

```yml
sudo dnf install libxkbcommon-devel libXtst-devel libxkbcommon-x11-devel xorg-x11-xkb-utils-devel

sudo dnf install libpng-devel

sudo dnf install xsel xclip
```

## Installation:
```
go get github.com/cauefcr/robotgo
```
  It's that easy!

png.h: No such file or directory? Please see [issues/47](https://github.com/cauefcr/robotgo/issues/47).

## Update:
```
go get -u github.com/cauefcr/robotgo
```

Note go1.10.x C file compilation cache problem, [golang #24355](https://github.com/golang/go/issues/24355).
`go mod vendor` problem, [golang #26366](https://github.com/golang/go/issues/26366).


## [Examples:](https://github.com/cauefcr/robotgo/blob/master/examples)

#### [Mouse](https://github.com/cauefcr/robotgo/blob/master/examples/mouse/main.go)

```Go
package main

import (
	"github.com/cauefcr/robotgo"
)

func main() {
  robotgo.ScrollMouse(10, "up")
  robotgo.MouseClick("left", true)
  robotgo.MoveMouseSmooth(100, 200, 1.0, 100.0)
}
```

#### [Keyboard](https://github.com/cauefcr/robotgo/blob/master/examples/key/main.go)

```Go
package main

import (
  "fmt"

  "github.com/cauefcr/robotgo"
)

func main() {
  robotgo.TypeStr("Hello World")
  robotgo.TypeStr("だんしゃり", 1.0)
  // robotgo.TypeString("テストする")

  robotgo.TypeStr("Hi galaxy. こんにちは世界.")
  robotgo.Sleep(1)

  // ustr := uint32(robotgo.CharCodeAt("Test", 0))
  // robotgo.UnicodeType(ustr)

  robotgo.KeyTap("enter")
  // robotgo.TypeString("en")
  robotgo.KeyTap("i", "alt", "command")

  arr := []string{"alt", "command"}
  robotgo.KeyTap("i", arr)

  robotgo.WriteAll("Test")
  text, err := robotgo.ReadAll()
  if err == nil {
    fmt.Println(text)
  }
}
```

#### [Screen](https://github.com/cauefcr/robotgo/blob/master/examples/screen/main.go)

```Go
package main

import (
	"fmt"

	"github.com/cauefcr/robotgo"
)

func main() {
  x, y := robotgo.GetMousePos()
  fmt.Println("pos: ", x, y)

  color := robotgo.GetPixelColor(100, 200)
  fmt.Println("color---- ", color)
}
```

#### [Bitmap](https://github.com/cauefcr/robotgo/blob/master/examples/bitmap/main.go)

```Go
package main

import (
	"fmt"

	"github.com/cauefcr/robotgo"
)

func main() {
  bitmap := robotgo.CaptureScreen(10, 20, 30, 40)
  // use `defer robotgo.FreeBitmap(bit)` to free the bitmap
  defer robotgo.FreeBitmap(bitmap)

  fmt.Println("...", bitmap)

  fx, fy := robotgo.FindBitmap(bitmap)
  fmt.Println("FindBitmap------ ", fx, fy)

  robotgo.SaveBitmap(bitmap, "test.png")
}
```

#### [Event](https://github.com/cauefcr/robotgo/blob/master/examples/event/main.go)

```Go
package main

import (
	"fmt"

	"github.com/cauefcr/robotgo"
)

func main() {
  ok := robotgo.AddEvents("q", "ctrl", "shift")
  if ok {
    fmt.Println("add events...")
  }

  keve := robotgo.AddEvent("k")
  if keve {
    fmt.Println("you press... ", "k")
  }

  mleft := robotgo.AddEvent("mleft")
  if mleft {
    fmt.Println("you press... ", "mouse left button")
  }
}
```

#### [Window](https://github.com/cauefcr/robotgo/blob/master/examples/window/main.go)

```Go
package main

import (
	"fmt"

	"github.com/cauefcr/robotgo"
)

func main() {
  fpid, err := robotgo.FindIds("Google")
  if err == nil {
    fmt.Println("pids... ", fpid)

    if len(fpid) > 0 {
      robotgo.ActivePID(fpid[0])

      robotgo.Kill(fpid[0])
    }
  }

  robotgo.ActiveName("chrome")

  isExist, err := robotgo.PidExists(100)
  if err == nil && isExist {
    fmt.Println("pid exists is", isExist)

    robotgo.Kill(100)
  }

  abool := robotgo.ShowAlert("test", "robotgo")
  if abool == 0 {
 	  fmt.Println("ok@@@ ", "ok")
  }

  title := robotgo.GetTitle()
  fmt.Println("title@@@ ", title)
}
```

## CrossCompiling

##### Windows64 to win32
```Go
SET CGO_ENABLED=1
SET GOARCH=386
go build main.go
```

#### Other to windows
```Go
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -x ./
```
```
// CC=mingw-w64\x86_64-7.2.0-win32-seh-rt_v5-rev1\mingw64\bin\gcc.exe
// CXX=mingw-w64\x86_64-7.2.0-win32-seh-rt_v5-rev1\mingw64\bin\g++.exe
```

Some discussions and questions, please see [issues/228](https://github.com/cauefcr/robotgo/issues/228), [issues/143](https://github.com/cauefcr/robotgo/issues/143).

## Authors
* [The author is vz](https://github.com/vcaesar)
* [Maintainers](https://github.com/orgs/cauefcr/people)
* [Contributors](https://github.com/cauefcr/robotgo/graphs/contributors)

## Plans
- Update Find an image on screen, read pixels from an image
- Update Window Handle
- Try support Android, maybe support IOS

## Donate

Supporting robotgo, [buy me a coffee](https://github.com/cauefcr/buy-me-a-coffee).

#### Paypal

Donate money by [paypal](https://www.paypal.me/veni0/25) to my account [vzvway@gmail.com](vzvway@gmail.com)


## Contributors

- See [contributors page](https://github.com/cauefcr/robotgo/graphs/contributors) for full list of contributors.
- See [Contribution Guidelines](https://github.com/cauefcr/robotgo/blob/master/CONTRIBUTING.md).

## License

Robotgo is primarily distributed under the terms of both the MIT license and the Apache License (Version 2.0), with portions covered by various BSD-like licenses.

See [LICENSE-APACHE](http://www.apache.org/licenses/LICENSE-2.0), [LICENSE-MIT](https://github.com/cauefcr/robotgo/blob/master/LICENSE).
