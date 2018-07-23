cat > /app/run.sh << EOF
#!/bin/bash
export GOPATH=/app
cd /app
go get github.com/rathodc/todo
go get github.com/gorilla/mux
go get github.com/xeipuuv/gojsonschema
go get github.com/stretchr/testify/assert
cd src/github.com/rathodc/todo/app
go install
export PATH=$GOPATH/bin:$PATH
/app/bin/app
tail -f /dev/null
EOF

chmod 755 /app/run.sh