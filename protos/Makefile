GENERATED_PATH := generated

PROTO_FILES := attributes command job nil pipeline queueinfo runnerinfo runnerstatus
PROTO_SERVICE_FILES := runner queue server

GENERATED_FILES := $(foreach file, \
	$(PROTO_SERVICE_FILES), \
	$(GENERATED_PATH)/$(file).pb.go $(GENERATED_PATH)/$(file)_grpc.pb.go \
) $(foreach file, $(PROTO_FILES), $(GENERATED_PATH)/$(file).pb.go)

clean-proto:
	-rm $(GENERATED_FILES)

# clean: clean-proto

$(GENERATED_FILES) &: $(PROTO_PATH)/*.proto
	protoc \
		--proto_path=$(PROTO_PATH) \
		--go_out=$(GENERATED_PATH) \
    --go-grpc_out=$(GENERATED_PATH) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_PATH)/*.proto

proto: $(GENERATED_FILES)

.PHONY: clean-proto proto
