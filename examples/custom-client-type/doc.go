package customclienttype

// This is an example of how to add a prefix to the name of the generated Client struct
// See https://github.com/obouchet/oapi-codegen/issues/785 for why this might be necessary

//go:generate go run github.com/obouchet/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml api.yaml
