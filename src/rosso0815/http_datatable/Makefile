#------------------------------------------------------------------------------
#
# Author : PfisterA
#
# Renark : Datatable JS
#
#------------------------------------------------------------------------------

PID      = /tmp/awesome-golang-project.pid
GO_FILES = $(wildcard *.go)
APP      = ./http_serve

serve: restart
	@echo "restart"
	@fswatch -o . | xargs -n1 -I{}  make restart || make kill

clean:
	@echo "=> clean files + dirs"
	@rm -f $(APP)
	@rm -f /tmp/$(APP)
	@rm -rf node_modules
	@npm i bootstrap
	@npm i bootstrap-vue
	@npm i vue
	@npm i vue-router
	@npm i vue-resource
	
kill:
	@echo "kill"
	@kill `cat $(PID)` || true

before:
	@echo "before $(APP) $(GO_FILES)"
	#$(APP): $(GO_FILES)
	#@( cd src && go build -o /tmp/$(APP) )
	go build -o /tmp/$(APP)

restart: kill before
	@echo "\n>>===== restart ======<<"
	#@rm $(PID)
	@/tmp/$(APP) & echo $$! > $(PID)

#.PHONY: serve restart kill before # let's go to reserve rules names
