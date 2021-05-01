
https://github.com/flosch/pongo2

export GO111MODULE=on

go get -v -x github.com/wrouesnel/p2cli

p2cli -t template.p2 -i input.yml -f yaml --enable-filters write_file
