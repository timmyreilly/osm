package lds

import (
	xds_discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"

	"github.com/openservicemesh/osm/pkg/catalog"
	"github.com/openservicemesh/osm/pkg/certificate"
	"github.com/openservicemesh/osm/pkg/configurator"
	"github.com/openservicemesh/osm/pkg/envoy"
	"github.com/openservicemesh/osm/pkg/envoy/registry"
	"github.com/openservicemesh/osm/pkg/identity"
)

// NewResponse creates a new Listener Discovery Response.
// The response build 3 Listeners:
// 1. Inbound listener to handle incoming traffic
// 2. Outbound listener to handle outgoing traffic
// 3. Prometheus listener for metrics
func NewResponse(meshCatalog catalog.MeshCataloger, proxy *envoy.Proxy, _ *xds_discovery.DiscoveryRequest, cfg configurator.Configurator, _ certificate.Manager, proxyRegistry *registry.ProxyRegistry) ([]types.Resource, error) {
	svcAccount, err := envoy.GetServiceAccountFromProxyCertificate(proxy.GetCertificateCommonName())
	if err != nil {
		log.Error().Err(err).Msgf("Error retrieving ServiceAccount for proxy %s", proxy.String())
		return nil, err
	}

	var ldsResources []types.Resource

	var statsHeaders map[string]string
	if featureflags := cfg.GetFeatureFlags(); featureflags.EnableWASMStats {
		statsHeaders = proxy.StatsHeaders()
	}

	lb := newListenerBuilder(meshCatalog, svcAccount.ToServiceIdentity(), cfg, statsHeaders)

	if proxy.Kind() == envoy.KindGateway {
		return lb.buildGatewayListeners(), nil
	}

	// --- OUTBOUND -------------------
	outboundListener, err := lb.newOutboundListener()
	if err != nil {
		log.Error().Err(err).Msgf("Error building outbound listener for proxy %s", proxy.String())
	} else {
		if outboundListener == nil {
			// This check is important to prevent attempting to configure a listener without a filter chain which
			// otherwise results in an error.
			log.Debug().Msgf("Not programming outbound listener for proxy %s", proxy.String())
		} else {
			ldsResources = append(ldsResources, outboundListener)
		}
	}

	// --- INBOUND -------------------
	inboundListener := newInboundListener()

	svcList, err := proxyRegistry.ListProxyServices(proxy)
	if err != nil {
		log.Error().Err(err).Msgf("Error looking up MeshService for proxy %s", proxy.String())
		return nil, err
	}
	// Create inbound filter chains per service behind proxy
	for _, proxyService := range svcList {
		// Create in-mesh filter chains
		inboundSvcFilterChains := lb.getInboundMeshFilterChains(proxyService)
		inboundListener.FilterChains = append(inboundListener.FilterChains, inboundSvcFilterChains...)

		// Create ingress filter chains if there are any ingress routes
		if ingressInboundPolicies, err := meshCatalog.GetIngressPoliciesForService(proxyService); err != nil {
			log.Error().Err(err).Msgf("Error getting ingress inbound traffic policies for service %s", proxyService)
		} else {
			thereAreIngressRoutes := len(ingressInboundPolicies) > 0
			if thereAreIngressRoutes {
				log.Info().Msgf("Found k8s Ingress for MeshService %s, applying necessary filters", proxyService)
				// This proxy is fronting a service that is a backend for an ingress, add a FilterChain for it
				ingressFilterChains := lb.getIngressFilterChains(proxyService)
				inboundListener.FilterChains = append(inboundListener.FilterChains, ingressFilterChains...)
			} else {
				log.Trace().Msgf("There is no k8s Ingress for service %s", proxyService)
			}
		}
	}

	if len(inboundListener.FilterChains) > 0 {
		// Inbound filter chains can be empty if the there both ingress and in-mesh policies are not configured.
		// Configuring a listener without a filter chain is an error.
		ldsResources = append(ldsResources, inboundListener)
	}

	if pod, err := envoy.GetPodFromCertificate(proxy.GetCertificateCommonName(), meshCatalog.GetKubeController()); err != nil {
		log.Warn().Msgf("Could not find pod for connecting proxy %s. No metadata was recorded.", proxy.GetCertificateSerialNumber())
	} else if meshCatalog.GetKubeController().IsMetricsEnabled(pod) {
		// Build Prometheus listener config
		prometheusConnManager := getPrometheusConnectionManager()
		if prometheusListener, err := buildPrometheusListener(prometheusConnManager); err != nil {
			log.Error().Err(err).Msgf("Error building Prometheus listener for proxy %s", proxy.String())
		} else {
			ldsResources = append(ldsResources, prometheusListener)
		}
	}

	return ldsResources, nil
}

// Note: ServiceIdentity must be in the format "name.namespace" [https://github.com/openservicemesh/osm/issues/3188]
func newListenerBuilder(meshCatalog catalog.MeshCataloger, svcIdentity identity.ServiceIdentity, cfg configurator.Configurator, statsHeaders map[string]string) *listenerBuilder {
	return &listenerBuilder{
		meshCatalog:     meshCatalog,
		serviceIdentity: svcIdentity,
		cfg:             cfg,
		statsHeaders:    statsHeaders,
	}
}
