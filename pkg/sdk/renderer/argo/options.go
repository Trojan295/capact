package argo

import (
	gqlpublicapi "projectvoltron.dev/voltron/pkg/och/api/graphql/public"
	"projectvoltron.dev/voltron/pkg/och/client/public"
	"projectvoltron.dev/voltron/pkg/sdk/apis/0.0.1/types"
)

type RendererOption func(*dedicatedRenderer)

func WithPlainTextUserInput(data map[string]interface{}) RendererOption {
	return func(r *dedicatedRenderer) {
		r.plainTextUserInput = data
	}
}

func WithTypeInstances(typeInstances []types.InputTypeInstanceRef) RendererOption {
	return func(r *dedicatedRenderer) {
		r.inputTypeInstances = typeInstances
	}
}

func WithImplementationRevisionFilter(filter gqlpublicapi.ImplementationRevisionFilter) RendererOption {
	return func(r *dedicatedRenderer) {
		r.ochImplementationFilters = append(r.ochImplementationFilters, public.WithImplementationFilter(filter))
	}
}