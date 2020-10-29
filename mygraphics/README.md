# mygraphics

we use imagemagick@6

```
export IM_PATH="/usr/local/opt/imagemagick@6"
export PATH="${IM_PATH}/bin:$PATH"
export LDFLAGS="-L${IM_PATH}/lib"
export CPPFLAGS="-I${IM_PATH}/include"
export PKG_CONFIG_PATH="/${IM_PATH}/lib/pkgconfig"
export CGO_CFLAGS_ALLOW='-Xpreprocessor'
pkg-config --cflags --libs MagickWand
go clean --modcache && go clean -i -r -cache -testcache -modcache
GODEBUG=gocacheverify=1 go get -u gopkg.in/gographics/imagick.v2/imagick
GOFLAGS="-count=1"  go test -race .
```
