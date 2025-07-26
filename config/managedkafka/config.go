package managedkafka

import (
	"github.com/crossplane/upjet/pkg/config"
	"github.com/upbound/provider-gcp/config/common"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("google_managed_kafka_acl", func(r *config.Resource) {
		r.ShortGroup = "managedkafka"
		r.Kind = "Acl"
		r.Version = "v1beta1"

		// Configure subnet reference
		r.References["gce_cluster.subnet"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.ExtractFolderIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_managed_kafka_cluster", func(r *config.Resource) {
		r.ShortGroup = "managedkafka"
		r.Kind = "Cluster"
		r.Version = "v1beta1"
		r.UseAsync = true

		// Configure subnet reference
		r.References["gce_cluster.subnet"] = config.Reference{
			TerraformName: "google_compute_subnetwork",
			Extractor:     common.ExtractFolderIDFuncPath,
		}
	})

	p.AddResourceConfigurator("google_managed_kafka_topic", func(r *config.Resource) {
		r.ShortGroup = "managedkafka"
		r.Kind = "Topic"
		r.Version = "v1beta1"

		// Configure cluster reference
		r.References["cluster"] = config.Reference{
			TerraformName: "google_managed_kafka_cluster",
			Extractor:     common.ExtractFolderIDFuncPath,
		}
	})
}
