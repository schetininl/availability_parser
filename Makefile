.PHONY: .build
.build:
	$(info #Building...)
	$CGO_ENABLED=0 go build -ldflags "-X 'github.com/schetininl/availability_parser/version.Name=availability_parser' -X 'github.com/schetininl/availability_parser/version.ProjectID=' -X 'github.com/schetininl/availability_parser/version.Version=' -X 'github.com/schetininl/availability_parser/version.GoVersion=' -X 'github.com/schetininl/availability_parser/version.BuildDate=' -X 'github.com/schetininl/availability_parser/version.GitLog=' -X 'github.com/schetininl/availability_parser/version.GitHash=' -X 'github.com/schetininl/availability_parser/version.GitBranch=' -X 'github.com/schetininl/availability_parser/version.publicPortDefault=' -X 'github.com/schetininl/availability_parser/version.adminPortDefault=' -X 'github.com/schetininl/availability_parser/version.grpcPortDefault='" -o ./bin/availability_parser ./cmd/availability_parser

# build app
.PHONY: build
build: .build
