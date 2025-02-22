basic:
	go run 1_hello_world/first.go
value:
	@go run 2_values/index.go
var:
	@go run 3_simple_variables/index.go
const:
	@go run 4_constants/index.go
loop:
	@go run 5_forloop/index.go
cond:
	@go run 6_if-else/index.go
switch:
	@go run 7_switch/index.go
arr:
	@go run 8_Arrays/index.go
slice:
	@if [ -f 9_slice/index.go ]; then \
		go run 9_slice/index.go; \
	else \
		echo "Error: 9_slice/index.go does not exist."; \
	fi
map: 
   
	@go run 10_map/index.go

range:
	@go run 11_range/index.go

func:
	@go run 12_func/index.go

import:
	@go run 13_import_stuff/index.go	

