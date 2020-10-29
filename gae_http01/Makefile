default: run ;
	
clean:
	@echo "clean"
	rm -rf external
	rm -rf static/external

update:
	@echo "update"
	bower install jquery --config.directory=static/external/js
	bower install angularjs --config.directory=static/external/js
	bower install bootstrap --config.directory=static/external/js
	bower install bootstrap-css-only --config.directory=static/external/css
	
run:
	@echo "start ..."
	goapp serve --port 9001
