#------------------------------------------------------------------------------
PID      = run.pid
GO_FILES = $(wildcard *.go)
#APP      = ./http_serve
#------------------------------------------------------------------------------

serve: restart
	@echo "=== restart"
	#@fswatch -o . | xargs -n1 -I{}  make restart || make kill

clean:
	@echo "=== clean files + dirs"
#	@rm -f $(APP)
#	@rm -f /tmp/$(APP)
#	@echo "=> update javascript"
#	@bower install angularjs
#	@bower install bootstrap
	
kill:
	@echo "=== kill"
	#@kill `cat $(PID)` || true
	#@kill `cat $(PID)` || true
	#@pkill -P `cat $(PID)` || true

before:
	@echo "=== before $(APP) $(GO_FILES)"
#	#$(APP): $(GO_FILES)
#	@( cd src && go build -o /tmp/$(APP) )

restart: kill before
	@echo "=== restart"
#	@go run main.go & echo $$! > $(PID)
#	@/tmp/$(APP) & echo $$! > $(PID)

#.PHONY: serve restart kill before # let's go to reserve rules names
