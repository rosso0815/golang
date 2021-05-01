
go get -u -x github.com/hairyhenderson/gomplate/v3/cmd/gomplate


echo "Hello, {{.Env.USER}}" | gomplate

gomplate --input-dir=templates --output-dir=generated --datasource config=config.yaml

gomplate -d config=./config.yaml -i 'the value we want is: {{ (datasource "config").foo.bar.baz }}'