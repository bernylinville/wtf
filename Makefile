APP=wtfutil

define HEADER
____    __    ____ .___________. _______
\   \  /  \  /   / |           ||   ____|
 \   \/    \/   /  `---|  |----`|  |__
  \            /       |  |     |   __|
   \    /\    /        |  |     |  |
    \__/  \__/         |__|     |__|

endef
export HEADER

# -------------------- Actions -------------------- #

## build: builds a local version
build:
	@echo "$$HEADER"
	@echo "Building..."
	go build -o bin/${APP}
	@echo "Done building"

## clean: removes old build cruft
clean:
	rm -rf ./dist
	rm -rf ./bin/${APP}
	@echo "Done cleaning"