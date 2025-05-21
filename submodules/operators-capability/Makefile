VERSION ?= latest
OS ?= darwin
ARCH ?= amd64

#
# developer tools
#
VENDIR_VERSION ?= v0.30.0
vendir:
        @curl -L https://github.com/vmware-tanzu/carvel-vendir/releases/download/$(VENDIR_VERSION)/vendir-$(OS)-$(ARCH) -o /usr/local/bin/vendir && chmod +x /usr/local/bin/vendir

YOT_VERSION ?= v0.6.5
yot:
        @curl -L https://github.com/vmware-tanzu-labs/yaml-overlay-tool/releases/download/$(YOT_VERSION)/yot_$(YOT_VERSION)_Darwin_x86_64.tar.gz -o /usr/local/bin/yot.tar.gz && \
                tar zxvf /usr/local/bin/yot.tar.gz -C /usr/local/bin && \
                chmod +x /usr/local/bin/yot

#
# code generation
#
WORKLOAD_CONFIG ?= .source/workload-cert-manager.yaml
define operator_builder_init
	operator-builder init \
		--workload-config $(WORKLOAD_CONFIG) \
		--repo github.com/tbd-paas/capabilities-certificates-operator \
		--controller-image quay.io/tbd-paas/certificates-operator:$(VERSION) \
		--skip-go-version-check
endef

#
# project management
#
CAPABILITY ?=
COMPONENT ?=

# create component scaffold
scaffold:
	@scripts/scaffold.sh

# download from upstream
download:
	@scripts/download.sh

# overlay changes from upstream
overlay:
	@scripts/overlay.sh
