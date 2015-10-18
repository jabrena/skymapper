# skymapper
A REST API developed to help robots on local navigation tasks

![ScreenShot](https://raw.githubusercontent.com/jabrena/skymapper/master/docs/webcamIdea.png)

Skymapper is a software developed for EV3 Brick which it has a ARM5 board.

To run the software execute the following statements:

``` go
export GOPATH=`pwd` 
go get github.com/blackjack/webcam
go get github.com/lucasb-eyer/go-colorful

go test
go test -bench .
```
