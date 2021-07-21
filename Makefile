GENDIR := generated

PROTO_FILES := ./proto/*.proto

# Function to execute a command. Note the empty line before endef to make sure each command
# gets executed separately instead of concatenated with previous one.
# Accepts command to execute as first parameter.
define exec-command
$(1)

endef


PROTO_GEN_GO_DIR ?= $(GENDIR)/go

