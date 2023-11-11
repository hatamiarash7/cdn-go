# Go API client for arvancloud

Use this documentation to learn how to use the ArvanCloud SDK.

**API version**: 4.114.0

## Dependencies

Install the following packages:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

## Usage

Get the package:

```bash
go get git.arvancloud.ir/arvancloud/cdn-go
```

Put the package under your project folder and add the following in import:

```golang
import arvancloud "git.arvancloud.ir/arvancloud/cdn-go-sdk"
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), arvancloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), arvancloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), arvancloud.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), arvancloud.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://napi.arvancloud.ir/cdn/4.0*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccelerationApi* | [**AccelerationShow**](AccelerationApi.md#accelerationshow) | **Get** /domains/{domain}/acceleration | Get the content of acceleration settings
*AccelerationApi* | [**AccelerationUpdate**](AccelerationApi.md#accelerationupdate) | **Patch** /domains/{domain}/acceleration | Update the content of acceleration settings
*AccelerationApi* | [**ImageResizeShow**](AccelerationApi.md#imageresizeshow) | **Get** /domains/{domain}/image-resize | Get the content of image resize settings
*AccelerationApi* | [**ImageResizeUpdate**](AccelerationApi.md#imageresizeupdate) | **Patch** /domains/{domain}/image-resize | Update the content of image resize settings
*ActiveHealthCheckApi* | [**ActiveHealthCheckDestroy**](ActiveHealthCheckApi.md#activehealthcheckdestroy) | **Delete** /domains/{domain}/health-checks/{healthcheck} | Delete healthcheck
*ActiveHealthCheckApi* | [**ActiveHealthCheckIndex**](ActiveHealthCheckApi.md#activehealthcheckindex) | **Get** /domains/{domain}/health-checks | Get Defined HealthCheck
*ActiveHealthCheckApi* | [**ActiveHealthCheckReportsDetails**](ActiveHealthCheckApi.md#activehealthcheckreportsdetails) | **Get** /domains/{domain}/health-checks/reports/details | Get reports of a single healthcheck monitoring
*ActiveHealthCheckApi* | [**ActiveHealthCheckReportsSummary**](ActiveHealthCheckApi.md#activehealthcheckreportssummary) | **Get** /domains/{domain}/health-checks/reports/summary | Get summary reports of a single healthcheck monitoring
*ActiveHealthCheckApi* | [**ActiveHealthCheckShow**](ActiveHealthCheckApi.md#activehealthcheckshow) | **Get** /domains/{domain}/health-checks/{healthcheck} | Get a single healthcheck
*ActiveHealthCheckApi* | [**ActiveHealthCheckStore**](ActiveHealthCheckApi.md#activehealthcheckstore) | **Post** /domains/{domain}/health-checks | Store a new HealthCheck
*ActiveHealthCheckApi* | [**ActiveHealthCheckUpdate**](ActiveHealthCheckApi.md#activehealthcheckupdate) | **Patch** /domains/{domain}/health-checks/{healthcheck} | Update Health check
*ActiveHealthCheckApi* | [**HealthChecksZonesIndex**](ActiveHealthCheckApi.md#healthcheckszonesindex) | **Get** /health-checks/zones | Get list of all health-check zones
*CDNAppsApi* | [**AppsCategoryIndex**](CDNAppsApi.md#appscategoryindex) | **Get** /apps/category | Get the list of application categories
*CDNAppsApi* | [**AppsCategoryShow**](CDNAppsApi.md#appscategoryshow) | **Get** /apps/category/{application-category} | Get an existing application category
*CDNAppsApi* | [**AppsIndex**](CDNAppsApi.md#appsindex) | **Get** /apps | Get list of all available cdn-apps
*CDNAppsApi* | [**AppsLike**](CDNAppsApi.md#appslike) | **Post** /apps/{id} | Expressing like and dislike about a single cdn-app
*CDNAppsApi* | [**AppsShow**](CDNAppsApi.md#appsshow) | **Get** /apps/{id} | Get a single cdn-app
*CDNAppsApi* | [**DomainsAppsDestroy**](CDNAppsApi.md#domainsappsdestroy) | **Delete** /domains/{domain}/apps/{id} | Uninstall the application from domain
*CDNAppsApi* | [**DomainsAppsIndex**](CDNAppsApi.md#domainsappsindex) | **Get** /domains/{domain}/apps | Get list of all applications installed on a domain
*CDNAppsApi* | [**DomainsAppsInstalled**](CDNAppsApi.md#domainsappsinstalled) | **Get** /domains/{domain}/apps/{id} | Check the application is installed on the domain
*CDNAppsApi* | [**DomainsAppsStore**](CDNAppsApi.md#domainsappsstore) | **Post** /domains/{domain}/apps/{id} | Install the application on the domain
*CDNAppsApi* | [**DomainsAppsTriggerWebhook**](CDNAppsApi.md#domainsappstriggerwebhook) | **Post** /domains/{domain}/apps/{id}/actions/trigger_webhook | trigger webhook event
*CachingApi* | [**CachingDeprecatedPurge**](CachingApi.md#cachingdeprecatedpurge) | **Delete** /domains/{domain}/caching | Purge CDN Cache
*CachingApi* | [**CachingIndex**](CachingApi.md#cachingindex) | **Get** /domains/{domain}/caching | Get caching settings
*CachingApi* | [**CachingPurge**](CachingApi.md#cachingpurge) | **Post** /domains/{domain}/caching/purge | Purge CDN Cache
*CachingApi* | [**CachingUpdate**](CachingApi.md#cachingupdate) | **Patch** /domains/{domain}/caching | Update caching settings
*CachingApi* | [**PurgeTagsDestroy**](CachingApi.md#purgetagsdestroy) | **Delete** /domains/{domain}/purge-tags | Delete a Domain&#39;s Purge tag
*CachingApi* | [**PurgeTagsIndex**](CachingApi.md#purgetagsindex) | **Get** /domains/{domain}/purge-tags | Get domain&#39;s Purge tags
*CustomPagesApi* | [**CustomPagesShow**](CustomPagesApi.md#custompagesshow) | **Get** /domains/{domain}/custom-pages | Get list of custom pages
*CustomPagesApi* | [**CustomPagesUpdate**](CustomPagesApi.md#custompagesupdate) | **Post** /domains/{domain}/custom-pages | Update custom page
*DDoSApi* | [**DdosReprioritize**](DDoSApi.md#ddosreprioritize) | **Post** /domains/{domain}/ddos/actions/reprioritize | Change priority of ddos rules
*DDoSApi* | [**DdosRulesDestroy**](DDoSApi.md#ddosrulesdestroy) | **Delete** /domains/{domain}/ddos/rules/{id} | Delete DDoS protection rule
*DDoSApi* | [**DdosRulesIndex**](DDoSApi.md#ddosrulesindex) | **Get** /domains/{domain}/ddos/rules | Get DDoS Protection Rules
*DDoSApi* | [**DdosRulesShow**](DDoSApi.md#ddosrulesshow) | **Get** /domains/{domain}/ddos/rules/{id} | Get DDoS protection&#39;s rule information
*DDoSApi* | [**DdosRulesStore**](DDoSApi.md#ddosrulesstore) | **Post** /domains/{domain}/ddos/rules | Create new DDoS protection rule
*DDoSApi* | [**DdosRulesUpdate**](DDoSApi.md#ddosrulesupdate) | **Patch** /domains/{domain}/ddos/rules/{id} | Update the DDoS protection rule
*DDoSApi* | [**DdosSettingsIndex**](DDoSApi.md#ddossettingsindex) | **Get** /domains/{domain}/ddos/settings | Get DDoS protection settings
*DDoSApi* | [**DdosSettingsUpdate**](DDoSApi.md#ddossettingsupdate) | **Patch** /domains/{domain}/ddos/settings | Update domain&#39;s DDoS protection configuration
*DNSManagementApi* | [**DnsRecordsCloud**](DNSManagementApi.md#dnsrecordscloud) | **Put** /domains/{domain}/dns-records/{id}/cloud | Toggle cloud status (To proxy or not proxy, that&#39;s the question!)
*DNSManagementApi* | [**DnsRecordsDestroy**](DNSManagementApi.md#dnsrecordsdestroy) | **Delete** /domains/{domain}/dns-records/{id} | Remove a DNS record
*DNSManagementApi* | [**DnsRecordsDnsSecShow**](DNSManagementApi.md#dnsrecordsdnssecshow) | **Get** /domains/{domain}/dns-records/dnssec | Get status of DNSSEC
*DNSManagementApi* | [**DnsRecordsDnsSecUpdate**](DNSManagementApi.md#dnsrecordsdnssecupdate) | **Put** /domains/{domain}/dns-records/dnssec/actions | Update DNSSEC status
*DNSManagementApi* | [**DnsRecordsImport**](DNSManagementApi.md#dnsrecordsimport) | **Post** /domains/{domain}/dns-records/import | Import DNS records using BIND file
*DNSManagementApi* | [**DnsRecordsIndex**](DNSManagementApi.md#dnsrecordsindex) | **Get** /domains/{domain}/dns-records | Get list of DNS records
*DNSManagementApi* | [**DnsRecordsShow**](DNSManagementApi.md#dnsrecordsshow) | **Get** /domains/{domain}/dns-records/{id} | Get information of a record
*DNSManagementApi* | [**DnsRecordsStore**](DNSManagementApi.md#dnsrecordsstore) | **Post** /domains/{domain}/dns-records | Create new DNS record
*DNSManagementApi* | [**DnsRecordsUpdate**](DNSManagementApi.md#dnsrecordsupdate) | **Put** /domains/{domain}/dns-records/{id} | Update a DNS record
*DomainApi* | [**DomainsClone**](DomainApi.md#domainsclone) | **Post** /domains/{domain}/clone | Clone a domain config from another one
*DomainApi* | [**DomainsCnameSetupCheck**](DomainApi.md#domainscnamesetupcheck) | **Get** /domains/{domain}/cname-setup/check | Check Cname Setup to find whether domain is activated
*DomainApi* | [**DomainsCnameSetupConvert**](DomainApi.md#domainscnamesetupconvert) | **Post** /domains/{domain}/cname-setup/convert | Convert domain setup to cname
*DomainApi* | [**DomainsCnameSetupReset**](DomainApi.md#domainscnamesetupreset) | **Delete** /domains/{domain}/cname-setup/custom | Reset the custom record of CNAME Setup to the default value
*DomainApi* | [**DomainsCnameSetupSet**](DomainApi.md#domainscnamesetupset) | **Put** /domains/{domain}/cname-setup/custom | Set a custom record for using CNAME Setup
*DomainApi* | [**DomainsDestroy**](DomainApi.md#domainsdestroy) | **Delete** /domains/{domain} | Remove the domain
*DomainApi* | [**DomainsIndex**](DomainApi.md#domainsindex) | **Get** /domains | Get the list of domains
*DomainApi* | [**DomainsNameserversCheck**](DomainApi.md#domainsnameserverscheck) | **Get** /domains/{domain}/ns-keys/check | Check NS to find whether domain is activated
*DomainApi* | [**DomainsNameserversDeprecatedCheck**](DomainApi.md#domainsnameserversdeprecatedcheck) | **Put** /domains/{domain}/dns-service/check-ns | Deprecated in favor /ns-keys/check
*DomainApi* | [**DomainsNameserversOptional**](DomainApi.md#domainsnameserversoptional) | **Post** /domains/{domain}/ns-keys/use-optional-keys | Use optional NS keys
*DomainApi* | [**DomainsNameserversReset**](DomainApi.md#domainsnameserversreset) | **Delete** /domains/{domain}/ns-keys | Reset custom Nameserver keys to the default values for the domain
*DomainApi* | [**DomainsNameserversSet**](DomainApi.md#domainsnameserversset) | **Put** /domains/{domain}/ns-keys | Set custom NS records for the domain
*DomainApi* | [**DomainsRegenerate**](DomainApi.md#domainsregenerate) | **Post** /domains/{domain}/regenerate | Regenerate domain settings for edge servers
*DomainApi* | [**DomainsShow**](DomainApi.md#domainsshow) | **Get** /domains/{domain} | Get information of the domain
*DomainApi* | [**DomainsStore**](DomainApi.md#domainsstore) | **Post** /domains/dns-service | Create new domain
*DomainTransferApi* | [**DomainsTransferIndex**](DomainTransferApi.md#domainstransferindex) | **Get** /domains/transfer | Get the list of pending transfers
*DomainTransferApi* | [**DomainsTransferStore**](DomainTransferApi.md#domainstransferstore) | **Post** /domains/{domain}/transfer | Transfer domain to another account
*DomainTransferApi* | [**DomainsTransferUpdate**](DomainTransferApi.md#domainstransferupdate) | **Post** /domains/transfer/change-status | Accept or cancel transferring a domain
*EmailForwardingApi* | [**EmailForwardingsActivate**](EmailForwardingApi.md#emailforwardingsactivate) | **Post** /domains/{domain}/email-forwardings/activate | Activate email forwarding
*EmailForwardingApi* | [**EmailForwardingsAliasesDestroy**](EmailForwardingApi.md#emailforwardingsaliasesdestroy) | **Delete** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId} | Delete an alias
*EmailForwardingApi* | [**EmailForwardingsAliasesIndex**](EmailForwardingApi.md#emailforwardingsaliasesindex) | **Get** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases | List of email forwarding aliases for given domain
*EmailForwardingApi* | [**EmailForwardingsAliasesStore**](EmailForwardingApi.md#emailforwardingsaliasesstore) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases | Create new alias
*EmailForwardingApi* | [**EmailForwardingsAliasesToggleActivation**](EmailForwardingApi.md#emailforwardingsaliasestoggleactivation) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId}/toggle-activation | Toggle alias activation
*EmailForwardingApi* | [**EmailForwardingsAliasesUpdateRecipients**](EmailForwardingApi.md#emailforwardingsaliasesupdaterecipients) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/aliases/{emailForwardingAliasId}/recipients | Update alias recipients
*EmailForwardingApi* | [**EmailForwardingsCatchAllActivate**](EmailForwardingApi.md#emailforwardingscatchallactivate) | **Post** /domains/{domain}/email-forwardings/catch-all/activate | Activate email forwarding catch all
*EmailForwardingApi* | [**EmailForwardingsCatchAllDeactivate**](EmailForwardingApi.md#emailforwardingscatchalldeactivate) | **Post** /domains/{domain}/email-forwardings/catch-all/deactivate | Deactivate email forwarding catch all
*EmailForwardingApi* | [**EmailForwardingsDeactivate**](EmailForwardingApi.md#emailforwardingsdeactivate) | **Post** /domains/{domain}/email-forwardings/deactivate | Deactivate email forwarding
*EmailForwardingApi* | [**EmailForwardingsRecipientsDestroy**](EmailForwardingApi.md#emailforwardingsrecipientsdestroy) | **Delete** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId} | Delete a recipient
*EmailForwardingApi* | [**EmailForwardingsRecipientsIndex**](EmailForwardingApi.md#emailforwardingsrecipientsindex) | **Get** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients | List recipients of an email forwarding
*EmailForwardingApi* | [**EmailForwardingsRecipientsResendVerification**](EmailForwardingApi.md#emailforwardingsrecipientsresendverification) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/resend-verification | Resend Verification
*EmailForwardingApi* | [**EmailForwardingsRecipientsSetDefault**](EmailForwardingApi.md#emailforwardingsrecipientssetdefault) | **Patch** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/set-default | Set default recipient
*EmailForwardingApi* | [**EmailForwardingsRecipientsStore**](EmailForwardingApi.md#emailforwardingsrecipientsstore) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients | Create new recipient
*EmailForwardingApi* | [**EmailForwardingsRecipientsVerify**](EmailForwardingApi.md#emailforwardingsrecipientsverify) | **Post** /domains/{domain}/email-forwardings/{emailForwardingId}/recipients/{emailForwardingRecipientId}/verify | Verify recipient
*EmailForwardingApi* | [**EmailForwardingsStats**](EmailForwardingApi.md#emailforwardingsstats) | **Get** /domains/{domain}/email-forwardings/stats | Show stats of domain&#39;s email forwarding
*FirewallApi* | [**FirewallReprioritize**](FirewallApi.md#firewallreprioritize) | **Post** /domains/{domain}/firewall/actions/reprioritize | Change priority of firewall rules
*FirewallApi* | [**FirewallRulesDestroy**](FirewallApi.md#firewallrulesdestroy) | **Delete** /domains/{domain}/firewall/rules/{id} | Delete firewall rule
*FirewallApi* | [**FirewallRulesIndex**](FirewallApi.md#firewallrulesindex) | **Get** /domains/{domain}/firewall/rules | Get domain&#39;s firewall rules
*FirewallApi* | [**FirewallRulesShow**](FirewallApi.md#firewallrulesshow) | **Get** /domains/{domain}/firewall/rules/{id} | Get firewall rule information
*FirewallApi* | [**FirewallRulesStore**](FirewallApi.md#firewallrulesstore) | **Post** /domains/{domain}/firewall/rules | Create new firewall rule
*FirewallApi* | [**FirewallRulesUpdate**](FirewallApi.md#firewallrulesupdate) | **Patch** /domains/{domain}/firewall/rules/{id} | Update the firewall rule
*FirewallApi* | [**FirewallSettingsIndex**](FirewallApi.md#firewallsettingsindex) | **Get** /domains/{domain}/firewall/settings | Get domain&#39;s firewall configuration
*FirewallApi* | [**FirewallSettingsUpdate**](FirewallApi.md#firewallsettingsupdate) | **Patch** /domains/{domain}/firewall/settings | Update domain&#39;s firewall configuration
*ListApi* | [**ListsDestroy**](ListApi.md#listsdestroy) | **Delete** /dynamic-fields/{id} | Delete List
*ListApi* | [**ListsIndex**](ListApi.md#listsindex) | **Get** /dynamic-fields | Get the list of Lists
*ListApi* | [**ListsShow**](ListApi.md#listsshow) | **Get** /dynamic-fields/{id} | Get an existing List
*ListApi* | [**ListsStore**](ListApi.md#listsstore) | **Post** /dynamic-fields | Store new List
*ListApi* | [**ListsUpdate**](ListApi.md#listsupdate) | **Patch** /dynamic-fields/{id} | Update an existing List
*LoadBalancingApi* | [**LoadBalancersDestroy**](LoadBalancingApi.md#loadbalancersdestroy) | **Delete** /domains/{domain}/load-balancers/{loadBalancerId} | Remove a load balancer
*LoadBalancingApi* | [**LoadBalancersIndex**](LoadBalancingApi.md#loadbalancersindex) | **Get** /domains/{domain}/load-balancers | Get list of load balancers
*LoadBalancingApi* | [**LoadBalancersPoolsDestroy**](LoadBalancingApi.md#loadbalancerspoolsdestroy) | **Delete** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Remove a load balancer pool
*LoadBalancingApi* | [**LoadBalancersPoolsIndex**](LoadBalancingApi.md#loadbalancerspoolsindex) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools | Get the list of pools of a load balancers
*LoadBalancingApi* | [**LoadBalancersPoolsOriginsDestroy**](LoadBalancingApi.md#loadbalancerspoolsoriginsdestroy) | **Delete** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins/{loadBalancerPoolOriginId} | Remove an origin from the pool of the load balancer
*LoadBalancingApi* | [**LoadBalancersPoolsOriginsIndex**](LoadBalancingApi.md#loadbalancerspoolsoriginsindex) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins | Get list of origins of a pool
*LoadBalancingApi* | [**LoadBalancersPoolsOriginsShow**](LoadBalancingApi.md#loadbalancerspoolsoriginsshow) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins/{loadBalancerPoolOriginId} | Get load balancer origin information
*LoadBalancingApi* | [**LoadBalancersPoolsOriginsStore**](LoadBalancingApi.md#loadbalancerspoolsoriginsstore) | **Post** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins | Create a new origin in the pool of the load balancer
*LoadBalancingApi* | [**LoadBalancersPoolsOriginsUpdate**](LoadBalancingApi.md#loadbalancerspoolsoriginsupdate) | **Patch** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId}/origins/{loadBalancerPoolOriginId} | Update an existing origin of the pool
*LoadBalancingApi* | [**LoadBalancersPoolsShow**](LoadBalancingApi.md#loadbalancerspoolsshow) | **Get** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Get load balancer pool information
*LoadBalancingApi* | [**LoadBalancersPoolsStore**](LoadBalancingApi.md#loadbalancerspoolsstore) | **Post** /domains/{domain}/load-balancers/{loadBalancerId}/pools | Create a new pool for the load balancer
*LoadBalancingApi* | [**LoadBalancersPoolsUpdate**](LoadBalancingApi.md#loadbalancerspoolsupdate) | **Put** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Update an existing load balancer pool with origins
*LoadBalancingApi* | [**LoadBalancersPoolsUpdatePool**](LoadBalancingApi.md#loadbalancerspoolsupdatepool) | **Patch** /domains/{domain}/load-balancers/{loadBalancerId}/pools/{loadBalancerPoolId} | Update an existing load balancer pool
*LoadBalancingApi* | [**LoadBalancersPrioritizePool**](LoadBalancingApi.md#loadbalancersprioritizepool) | **Post** /domains/{domain}/load-balancers/{loadBalancerId}/prioritize | Reorder the priority of load balancer pools
*LoadBalancingApi* | [**LoadBalancersRegionsIndex**](LoadBalancingApi.md#loadbalancersregionsindex) | **Get** /load-balancers/regions | Get list of regions for load balancers
*LoadBalancingApi* | [**LoadBalancersSettingsShow**](LoadBalancingApi.md#loadbalancerssettingsshow) | **Get** /domains/{domain}/load-balancers/settings | Get list of domain load balancer global settings
*LoadBalancingApi* | [**LoadBalancersSettingsUpdate**](LoadBalancingApi.md#loadbalancerssettingsupdate) | **Patch** /domains/{domain}/load-balancers/settings | Update domain&#39;s global load balancer settings
*LoadBalancingApi* | [**LoadBalancersShow**](LoadBalancingApi.md#loadbalancersshow) | **Get** /domains/{domain}/load-balancers/{loadBalancerId} | Get load balancer information
*LoadBalancingApi* | [**LoadBalancersStore**](LoadBalancingApi.md#loadbalancersstore) | **Post** /domains/{domain}/load-balancers | Create a new load balancer
*LoadBalancingApi* | [**LoadBalancersUpdate**](LoadBalancingApi.md#loadbalancersupdate) | **Patch** /domains/{domain}/load-balancers/{loadBalancerId} | Update a load balancer
*LoadBalancingApi* | [**LoadBalancersUpdateWithPools**](LoadBalancingApi.md#loadbalancersupdatewithpools) | **Put** /domains/{domain}/load-balancers/{loadBalancerId} | Update a load balancer
*LogForwardersApi* | [**LogForwardersDestroy**](LogForwardersApi.md#logforwardersdestroy) | **Delete** /domains/{domain}/log-forwarders/{logForwarderId} | Delete a log forwarder
*LogForwardersApi* | [**LogForwardersIndex**](LogForwardersApi.md#logforwardersindex) | **Get** /domains/{domain}/log-forwarders | Show list of log forwarders for given domain
*LogForwardersApi* | [**LogForwardersShow**](LogForwardersApi.md#logforwardersshow) | **Get** /domains/{domain}/log-forwarders/{logForwarderId} | Show a log forwarder&#39;s details based on given id
*LogForwardersApi* | [**LogForwardersStore**](LogForwardersApi.md#logforwardersstore) | **Post** /domains/{domain}/log-forwarders | Create new log forwarder
*LogForwardersApi* | [**LogForwardersUpdate**](LogForwardersApi.md#logforwardersupdate) | **Put** /domains/{domain}/log-forwarders/{logForwarderId} | Update a log forwarder
*LogForwardersApi* | [**LogForwardersUpdateStatus**](LogForwardersApi.md#logforwardersupdatestatus) | **Patch** /domains/{domain}/log-forwarders/{logForwarderId}/status | Update a log forwarder&#39;s status
*PageRuleApi* | [**PageRulesDestroy**](PageRuleApi.md#pagerulesdestroy) | **Delete** /domains/{domain}/page-rules/{id} | Delete the page-rule
*PageRuleApi* | [**PageRulesDiffShow**](PageRuleApi.md#pagerulesdiffshow) | **Get** /domains/{domain}/page-rules/{id}/diff | Get the page-rule&#39;s exceptions
*PageRuleApi* | [**PageRulesDiffUpdate**](PageRuleApi.md#pagerulesdiffupdate) | **Patch** /domains/{domain}/page-rules/{id}/diff | Update the page-rule&#39;s exceptions
*PageRuleApi* | [**PageRulesIndex**](PageRuleApi.md#pagerulesindex) | **Get** /domains/{domain}/page-rules | Get list of page-rules
*PageRuleApi* | [**PageRulesPurge**](PageRuleApi.md#pagerulespurge) | **Delete** /domains/{domain}/page-rules/{id}/purge | Purge the page-rule
*PageRuleApi* | [**PageRulesShow**](PageRuleApi.md#pagerulesshow) | **Get** /domains/{domain}/page-rules/{id} | Get the page-rule&#39;s information
*PageRuleApi* | [**PageRulesStatusUpdate**](PageRuleApi.md#pagerulesstatusupdate) | **Patch** /domains/{domain}/page-rules/{id} | Toggle status of the page-rule
*PageRuleApi* | [**PageRulesStore**](PageRuleApi.md#pagerulesstore) | **Post** /domains/{domain}/page-rules | Create new page-rule
*PageRuleApi* | [**PageRulesUpdate**](PageRuleApi.md#pagerulesupdate) | **Put** /domains/{domain}/page-rules/{id} | Update the page-rule
*PlanApi* | [**DomainsPlans**](PlanApi.md#domainsplans) | **Get** /domains/{domain}/plans | Get the list of feature defintions for plans based on different sets
*PlanApi* | [**DomainsPlansUpdate**](PlanApi.md#domainsplansupdate) | **Put** /domains/{domain}/plan | Update the domain&#39;s plan
*PlanApi* | [**DomainsPlansUsages**](PlanApi.md#domainsplansusages) | **Get** /domains/{domain}/plan/usages | Get usages based on features and an estimated cost
*PlanApi* | [**DomainsPlansViolations**](PlanApi.md#domainsplansviolations) | **Get** /domains/{domain}/plan/violations | Get violations based on plans
*PlanApi* | [**PlansIndex**](PlanApi.md#plansindex) | **Get** /plans | Get the list of feature defintions for plans based on different sets
*RateLimitingApi* | [**RateLimitingActionsReprioritize**](RateLimitingApi.md#ratelimitingactionsreprioritize) | **Post** /domains/{domain}/rate-limit/actions/reprioritize | Change priority of Rate limiting&#39;s rules
*RateLimitingApi* | [**RateLimitingRulesDestroy**](RateLimitingApi.md#ratelimitingrulesdestroy) | **Delete** /domains/{domain}/rate-limit/rules/{id} | Delete Rate limiting&#39;s rule
*RateLimitingApi* | [**RateLimitingRulesIndex**](RateLimitingApi.md#ratelimitingrulesindex) | **Get** /domains/{domain}/rate-limit/rules | Get Rate limiting rules
*RateLimitingApi* | [**RateLimitingRulesShow**](RateLimitingApi.md#ratelimitingrulesshow) | **Get** /domains/{domain}/rate-limit/rules/{id} | Get Rate limiting&#39;s rule information
*RateLimitingApi* | [**RateLimitingRulesStore**](RateLimitingApi.md#ratelimitingrulesstore) | **Post** /domains/{domain}/rate-limit/rules | Store new Rate limiting rule
*RateLimitingApi* | [**RateLimitingRulesUpdate**](RateLimitingApi.md#ratelimitingrulesupdate) | **Patch** /domains/{domain}/rate-limit/rules/{id} | Update the Rate limiting&#39;s rule
*RateLimitingApi* | [**RateLimitingSettingsIndex**](RateLimitingApi.md#ratelimitingsettingsindex) | **Get** /domains/{domain}/rate-limit/settings | Get Rate limiting settings
*RateLimitingApi* | [**RateLimitingSettingsUpdate**](RateLimitingApi.md#ratelimitingsettingsupdate) | **Patch** /domains/{domain}/rate-limit/settings | Update domain&#39;s Rate limiting configuration
*RedirectApi* | [**RedirectShow**](RedirectApi.md#redirectshow) | **Get** /domains/{domain}/settings/www-redirect | Get redirect settings
*RedirectApi* | [**RedirectUpdate**](RedirectApi.md#redirectupdate) | **Put** /domains/{domain}/settings/www-redirect | Update redirect settings
*ReportsApi* | [**BulkReportsTrafficsTotal**](ReportsApi.md#bulkreportstrafficstotal) | **Post** /bulk/reports/traffics | Get traffic report for multiple domains
*ReportsApi* | [**BulkReportsVisitorsTotal**](ReportsApi.md#bulkreportsvisitorstotal) | **Post** /bulk/reports/visitors | Get visitor report for multiple domains
*ReportsApi* | [**ReportsAttacksAttackers**](ReportsApi.md#reportsattacksattackers) | **Get** /domains/{domain}/reports/attacks/attackers | Get list of attackers information
*ReportsApi* | [**ReportsAttacksIndex**](ReportsApi.md#reportsattacksindex) | **Get** /domains/{domain}/reports/attacks/list | Get list of attacks details
*ReportsApi* | [**ReportsAttacksMap**](ReportsApi.md#reportsattacksmap) | **Get** /domains/{domain}/reports/attacks/map | Get geo-map of attacks
*ReportsApi* | [**ReportsAttacksShow**](ReportsApi.md#reportsattacksshow) | **Get** /domains/{domain}/reports/attacks | Get report of attacks
*ReportsApi* | [**ReportsAttacksUri**](ReportsApi.md#reportsattacksuri) | **Get** /domains/{domain}/reports/attacks/uri | Get list of URLs under attack
*ReportsApi* | [**ReportsDnsGeo**](ReportsApi.md#reportsdnsgeo) | **Get** /domains/{domain}/reports/dns-geo | Get DNS requests as geo-map
*ReportsApi* | [**ReportsDnsRequests**](ReportsApi.md#reportsdnsrequests) | **Get** /domains/{domain}/reports/dns-requests | Get response time report
*ReportsApi* | [**ReportsErrorLogDetails**](ReportsApi.md#reportserrorlogdetails) | **Get** /domains/{domain}/reports/error-log-details | Get detail of an error
*ReportsApi* | [**ReportsErrorLogs**](ReportsApi.md#reportserrorlogs) | **Get** /domains/{domain}/reports/error-logs | Get list of errors
*ReportsApi* | [**ReportsErrorLogsChart**](ReportsApi.md#reportserrorlogschart) | **Get** /domains/{domain}/reports/error-logs/chart | Get chart view of errors
*ReportsApi* | [**ReportsResponseTimeIndex**](ReportsApi.md#reportsresponsetimeindex) | **Get** /domains/{domain}/reports/response-time | Get response time report
*ReportsApi* | [**ReportsStatusIndex**](ReportsApi.md#reportsstatusindex) | **Get** /domains/{domain}/reports/status | Get status codes pie chart
*ReportsApi* | [**ReportsStatusSummary**](ReportsApi.md#reportsstatussummary) | **Get** /domains/{domain}/reports/status/summary | Get an overview of status codes pie chart
*ReportsApi* | [**ReportsTrafficsMap**](ReportsApi.md#reportstrafficsmap) | **Get** /domains/{domain}/reports/traffics/map | Get traffic as geo-map
*ReportsApi* | [**ReportsTrafficsSaved**](ReportsApi.md#reportstrafficssaved) | **Get** /domains/{domain}/reports/traffics/saved | Get traffic saved to total pie chart
*ReportsApi* | [**ReportsTrafficsTotal**](ReportsApi.md#reportstrafficstotal) | **Get** /domains/{domain}/reports/traffics | Get traffic report for domain
*ReportsApi* | [**ReportsTransportLayerProxiesTraffics**](ReportsApi.md#reportstransportlayerproxiestraffics) | **Get** /domains/{domain}/reports/transport-layer-proxies/{transportLayerProxyId}/traffics | Get traffic report for transport layer proxy
*ReportsApi* | [**ReportsVisitorsHighRequestIps**](ReportsApi.md#reportsvisitorshighrequestips) | **Get** /domains/{domain}/reports/high-request-ips | Get report of IPs with highest number of requests
*ReportsApi* | [**ReportsVisitorsIndex**](ReportsApi.md#reportsvisitorsindex) | **Get** /domains/{domain}/reports/visitors | Get report of visitors for domain
*SSLTLSApi* | [**SslCertDestroy**](SSLTLSApi.md#sslcertdestroy) | **Delete** /domains/{domain}/ssl/certificates/{id} | Delete an unused customer certificate
*SSLTLSApi* | [**SslCertOrderIndex**](SSLTLSApi.md#sslcertorderindex) | **Get** /domains/{domain}/ssl/orders | Get All Managed certificate orders history
*SSLTLSApi* | [**SslCertOrderRetry**](SSLTLSApi.md#sslcertorderretry) | **Post** /domains/{domain}/ssl/orders/action/retry | Retry a previously &#x60;killed&#x60; order
*SSLTLSApi* | [**SslCertStore**](SSLTLSApi.md#sslcertstore) | **Post** /domains/{domain}/ssl/certificates | Upload Certificate
*SSLTLSApi* | [**SslIndex**](SSLTLSApi.md#sslindex) | **Get** /domains/{domain}/ssl | Get SSL settings
*SSLTLSApi* | [**SslUpdate**](SSLTLSApi.md#sslupdate) | **Patch** /domains/{domain}/ssl | Update domain&#39;s SSL settings
*TransportLayerProxyApi* | [**TransportLayerProxiesDestroy**](TransportLayerProxyApi.md#transportlayerproxiesdestroy) | **Delete** /domains/{domain}/transport-layer-proxies/{transportLayerProxyId} | delete a transport layer proxy
*TransportLayerProxyApi* | [**TransportLayerProxiesIndex**](TransportLayerProxyApi.md#transportlayerproxiesindex) | **Get** /domains/{domain}/transport-layer-proxies | Show list of transport layer proxies for given domain
*TransportLayerProxyApi* | [**TransportLayerProxiesShow**](TransportLayerProxyApi.md#transportlayerproxiesshow) | **Get** /domains/{domain}/transport-layer-proxies/{transportLayerProxyId} | Show a transport layer proxy&#39;s details based on given id
*TransportLayerProxyApi* | [**TransportLayerProxiesStore**](TransportLayerProxyApi.md#transportlayerproxiesstore) | **Post** /domains/{domain}/transport-layer-proxies | Create new transport layer proxy
*TransportLayerProxyApi* | [**TransportLayerProxiesUpdate**](TransportLayerProxyApi.md#transportlayerproxiesupdate) | **Put** /domains/{domain}/transport-layer-proxies/{transportLayerProxyId} | Update a transport layer proxy
*TroubleshootApi* | [**TroubleshootsIndex**](TroubleshootApi.md#troubleshootsindex) | **Get** /domains/{domain}/troubleshoots | Show list of troubleshoots for given domain
*TroubleshootApi* | [**TroubleshootsLatest**](TroubleshootApi.md#troubleshootslatest) | **Get** /domains/{domain}/troubleshoots/latest | Show the latest troubleshoot for given domain
*TroubleshootApi* | [**TroubleshootsStore**](TroubleshootApi.md#troubleshootsstore) | **Post** /domains/{domain}/troubleshoots | Create new troubleshoot
*WAFApi* | [**GlobalWafIndex**](WAFApi.md#globalwafindex) | **Get** /waf | Get WAF presets
*WAFApi* | [**GlobalWafShowPackage**](WAFApi.md#globalwafshowpackage) | **Get** /waf/packages/{packageId} | Get WAF package details
*WAFApi* | [**WafPackageReprioritize**](WAFApi.md#wafpackagereprioritize) | **Post** /domains/{domain}/waf/actions/reprioritize-package | Change priority of WAF packages
*WAFApi* | [**WafPackagesDestroy**](WAFApi.md#wafpackagesdestroy) | **Delete** /domains/{domain}/waf/packages/{id} | Delete WAF package from domain
*WAFApi* | [**WafPackagesIndex**](WAFApi.md#wafpackagesindex) | **Get** /domains/{domain}/waf/packages | Get WAF packages
*WAFApi* | [**WafPackagesShow**](WAFApi.md#wafpackagesshow) | **Get** /domains/{domain}/waf/packages/{id} | Get WAF package information
*WAFApi* | [**WafPackagesStore**](WAFApi.md#wafpackagesstore) | **Post** /domains/{domain}/waf/packages | Add new WAF package to domain
*WAFApi* | [**WafPackagesUpdate**](WAFApi.md#wafpackagesupdate) | **Patch** /domains/{domain}/waf/packages/{id} | Update the WAF package
*WAFApi* | [**WafReconfigure**](WAFApi.md#wafreconfigure) | **Post** /domains/{domain}/waf/actions/reconfigure | Reconfigure WAF module using a preset
*WAFApi* | [**WafReprioritize**](WAFApi.md#wafreprioritize) | **Post** /domains/{domain}/waf/actions/reprioritize | Change priority of WAF rules
*WAFApi* | [**WafRulesDestroy**](WAFApi.md#wafrulesdestroy) | **Delete** /domains/{domain}/waf/rules/{id} | Delete WAF rule
*WAFApi* | [**WafRulesIndex**](WAFApi.md#wafrulesindex) | **Get** /domains/{domain}/waf/rules | Get WAF Rules
*WAFApi* | [**WafRulesShow**](WAFApi.md#wafrulesshow) | **Get** /domains/{domain}/waf/rules/{id} | Get WAF rule information
*WAFApi* | [**WafRulesStore**](WAFApi.md#wafrulesstore) | **Post** /domains/{domain}/waf/rules | Create new WAF rule
*WAFApi* | [**WafRulesUpdate**](WAFApi.md#wafrulesupdate) | **Patch** /domains/{domain}/waf/rules/{id} | Update the WAF rule
*WAFApi* | [**WafSettingsIndex**](WAFApi.md#wafsettingsindex) | **Get** /domains/{domain}/waf/settings | Get WAF configuration
*WAFApi* | [**WafSettingsUpdate**](WAFApi.md#wafsettingsupdate) | **Patch** /domains/{domain}/waf/settings | Configure WAF module of the domain


## Documentation For Models

 - [AAAARecord](AAAARecord.md)
 - [AAAARecordValue](AAAARecordValue.md)
 - [ANAMERecord](ANAMERecord.md)
 - [ANAMERecordValue](ANAMERecordValue.md)
 - [ARecord](ARecord.md)
 - [ARecordValue](ARecordValue.md)
 - [Acceleration](Acceleration.md)
 - [AccelerationResponse](AccelerationResponse.md)
 - [AccelerationUpdate](AccelerationUpdate.md)
 - [ActiveHealthCheckIndex200Response](ActiveHealthCheckIndex200Response.md)
 - [ActiveHealthCheckReportsDetails200Response](ActiveHealthCheckReportsDetails200Response.md)
 - [ActiveHealthCheckReportsSummary200Response](ActiveHealthCheckReportsSummary200Response.md)
 - [ApplicationCategory](ApplicationCategory.md)
 - [ApplicationCategoryApplicationsInner](ApplicationCategoryApplicationsInner.md)
 - [ApplicationCategoryNameTranslation](ApplicationCategoryNameTranslation.md)
 - [ApplicationCategoryNameTranslationEn](ApplicationCategoryNameTranslationEn.md)
 - [AppsCategoryIndex200Response](AppsCategoryIndex200Response.md)
 - [AppsCategoryShow200Response](AppsCategoryShow200Response.md)
 - [AppsIndex200Response](AppsIndex200Response.md)
 - [AttackReport](AttackReport.md)
 - [AttackReportCharts](AttackReportCharts.md)
 - [AttackReportChartsAttacks](AttackReportChartsAttacks.md)
 - [AttackReportChartsAttacksSeriesInner](AttackReportChartsAttacksSeriesInner.md)
 - [AttackReportItem](AttackReportItem.md)
 - [AttackReportMap](AttackReportMap.md)
 - [AttackReportMapCharts](AttackReportMapCharts.md)
 - [AttackReportMapChartsAttacksValue](AttackReportMapChartsAttacksValue.md)
 - [AttackReportMapData](AttackReportMapData.md)
 - [AttackReportMapStatisticsInner](AttackReportMapStatisticsInner.md)
 - [AttackReportStatistics](AttackReportStatistics.md)
 - [AttackReportUri](AttackReportUri.md)
 - [AttackReportUriData](AttackReportUriData.md)
 - [BaseDnsRecord](BaseDnsRecord.md)
 - [BaseFirewallRule](BaseFirewallRule.md)
 - [BaseFirewallSettings](BaseFirewallSettings.md)
 - [BaseHealthCheck](BaseHealthCheck.md)
 - [BaseRateLimitRule](BaseRateLimitRule.md)
 - [BulkReportsTrafficsTotalRequest](BulkReportsTrafficsTotalRequest.md)
 - [BulkReportsVisitorsTotalRequest](BulkReportsVisitorsTotalRequest.md)
 - [BulkTrafficReport](BulkTrafficReport.md)
 - [BulkTrafficReportData](BulkTrafficReportData.md)
 - [BulkTrafficReportDataEgressBytes](BulkTrafficReportDataEgressBytes.md)
 - [BulkVisitorReport](BulkVisitorReport.md)
 - [BulkVisitorReportData](BulkVisitorReportData.md)
 - [BypassAction](BypassAction.md)
 - [CAARecord](CAARecord.md)
 - [CAARecordValue](CAARecordValue.md)
 - [CNAMERecord](CNAMERecord.md)
 - [CNAMERecordValue](CNAMERecordValue.md)
 - [CacheSettings](CacheSettings.md)
 - [CacheSettingsData](CacheSettingsData.md)
 - [CachingPurge](CachingPurge.md)
 - [CdnApp](CdnApp.md)
 - [CdnAppData](CdnAppData.md)
 - [CdnAppInstall](CdnAppInstall.md)
 - [CdnAppLike](CdnAppLike.md)
 - [CdnAppLikeStats](CdnAppLikeStats.md)
 - [CdnAppLikeStatsData](CdnAppLikeStatsData.md)
 - [CdnAppTriggerWebhook](CdnAppTriggerWebhook.md)
 - [Certificate](Certificate.md)
 - [CertificateOrder](CertificateOrder.md)
 - [ChallengeAction](ChallengeAction.md)
 - [CloneDomain](CloneDomain.md)
 - [CountryList](CountryList.md)
 - [CountryRequestChart](CountryRequestChart.md)
 - [CountryStatistics](CountryStatistics.md)
 - [CountryTrafficChart](CountryTrafficChart.md)
 - [Currency](Currency.md)
 - [CustomCname](CustomCname.md)
 - [CustomPage](CustomPage.md)
 - [CustomPages](CustomPages.md)
 - [CustomPagesData](CustomPagesData.md)
 - [DKIMRecord](DKIMRecord.md)
 - [DataResponse](DataResponse.md)
 - [DataWithMessageResponse](DataWithMessageResponse.md)
 - [DdosPreflight](DdosPreflight.md)
 - [DdosRule](DdosRule.md)
 - [DdosRuleData](DdosRuleData.md)
 - [DdosRuleResponse](DdosRuleResponse.md)
 - [DdosRulesIndex200Response](DdosRulesIndex200Response.md)
 - [DdosSettings](DdosSettings.md)
 - [DdosSettingsData](DdosSettingsData.md)
 - [DdosSettingsUpdate200Response](DdosSettingsUpdate200Response.md)
 - [DeprecatedNs](DeprecatedNs.md)
 - [DnsGeoReport](DnsGeoReport.md)
 - [DnsGeoReportCharts](DnsGeoReportCharts.md)
 - [DnsGeoReportChartsRequestsValue](DnsGeoReportChartsRequestsValue.md)
 - [DnsGeoReportData](DnsGeoReportData.md)
 - [DnsGeoReportListsInner](DnsGeoReportListsInner.md)
 - [DnsRecord](DnsRecord.md)
 - [DnsRecordCloud](DnsRecordCloud.md)
 - [DnsRecordData](DnsRecordData.md)
 - [DnsRecordGeneric](DnsRecordGeneric.md)
 - [DnsRecordGenericArrayValue](DnsRecordGenericArrayValue.md)
 - [DnsRecordGenericObjectValue](DnsRecordGenericObjectValue.md)
 - [DnsRecordIpFilterMode](DnsRecordIpFilterMode.md)
 - [DnsRecordResponse](DnsRecordResponse.md)
 - [DnsRecordsIndex200Response](DnsRecordsIndex200Response.md)
 - [DnsRequestReport](DnsRequestReport.md)
 - [DnsRequestReportCharts](DnsRequestReportCharts.md)
 - [DnsRequestReportChartsRequests](DnsRequestReportChartsRequests.md)
 - [DnsRequestReportChartsRequestsSeriesInner](DnsRequestReportChartsRequestsSeriesInner.md)
 - [DnsRequestReportData](DnsRequestReportData.md)
 - [DnsRequestReportStatistics](DnsRequestReportStatistics.md)
 - [DnsSec](DnsSec.md)
 - [DnsSecData](DnsSecData.md)
 - [DnsSecStatus](DnsSecStatus.md)
 - [Domain](Domain.md)
 - [DomainCdnApp](DomainCdnApp.md)
 - [DomainPurgeTags](DomainPurgeTags.md)
 - [DomainResponse](DomainResponse.md)
 - [DomainStore](DomainStore.md)
 - [DomainTransferData](DomainTransferData.md)
 - [DomainWafPackage](DomainWafPackage.md)
 - [DomainWafPackageDetails](DomainWafPackageDetails.md)
 - [DomainWafPackageDetailsData](DomainWafPackageDetailsData.md)
 - [DomainWafPackageStore](DomainWafPackageStore.md)
 - [DomainWafPackagesData](DomainWafPackagesData.md)
 - [DomainsAppsStore200Response](DomainsAppsStore200Response.md)
 - [DomainsIndex200Response](DomainsIndex200Response.md)
 - [DomainsNameserversCheck200Response](DomainsNameserversCheck200Response.md)
 - [DomainsNameserversDeprecatedCheck200Response](DomainsNameserversDeprecatedCheck200Response.md)
 - [DomainsPlansUsages200Response](DomainsPlansUsages200Response.md)
 - [DomainsPlansViolations200Response](DomainsPlansViolations200Response.md)
 - [DomainsShow404Response](DomainsShow404Response.md)
 - [DomainsStore422Response](DomainsStore422Response.md)
 - [DomainsStore422ResponseErrors](DomainsStore422ResponseErrors.md)
 - [DomainsTransferIndex200Response](DomainsTransferIndex200Response.md)
 - [DynamicField](DynamicField.md)
 - [DynamicFieldData](DynamicFieldData.md)
 - [DynamicFieldResponse](DynamicFieldResponse.md)
 - [DynamicFieldType](DynamicFieldType.md)
 - [DynamicFieldValue](DynamicFieldValue.md)
 - [EmailForwardingAlias](EmailForwardingAlias.md)
 - [EmailForwardingAliasesListData](EmailForwardingAliasesListData.md)
 - [EmailForwardingAliasesListInner](EmailForwardingAliasesListInner.md)
 - [EmailForwardingAliasesRecipients](EmailForwardingAliasesRecipients.md)
 - [EmailForwardingAliasesStore](EmailForwardingAliasesStore.md)
 - [EmailForwardingAliasesToggleActivation](EmailForwardingAliasesToggleActivation.md)
 - [EmailForwardingRecipient](EmailForwardingRecipient.md)
 - [EmailForwardingRecipientsListData](EmailForwardingRecipientsListData.md)
 - [EmailForwardingRecipientsListInner](EmailForwardingRecipientsListInner.md)
 - [EmailForwardingRecipientsStore](EmailForwardingRecipientsStore.md)
 - [EmailForwardingRecipientsVerify](EmailForwardingRecipientsVerify.md)
 - [EmailForwardingStats](EmailForwardingStats.md)
 - [EmailForwardingStatsData](EmailForwardingStatsData.md)
 - [EmailForwardingsAliasesStore201Response](EmailForwardingsAliasesStore201Response.md)
 - [EmailForwardingsRecipientsStore201Response](EmailForwardingsRecipientsStore201Response.md)
 - [ErrorLog](ErrorLog.md)
 - [ErrorLogChart](ErrorLogChart.md)
 - [ErrorLogChartCharts](ErrorLogChartCharts.md)
 - [ErrorLogChartChartsStatusCode](ErrorLogChartChartsStatusCode.md)
 - [ErrorLogChartChartsStatusCodeSeriesInner](ErrorLogChartChartsStatusCodeSeriesInner.md)
 - [ErrorLogChartData](ErrorLogChartData.md)
 - [ErrorLogChartStatistics](ErrorLogChartStatistics.md)
 - [ErrorLogUpstreamsInner](ErrorLogUpstreamsInner.md)
 - [ErrorLogsData](ErrorLogsData.md)
 - [EstimatedCost](EstimatedCost.md)
 - [ExpectedResponse](ExpectedResponse.md)
 - [FeatureDefinition](FeatureDefinition.md)
 - [FeatureDefinitionMeta](FeatureDefinitionMeta.md)
 - [FeatureDefinitionPlans](FeatureDefinitionPlans.md)
 - [FeaturePlanDefinition](FeaturePlanDefinition.md)
 - [FeaturePlanDefinitionMeta](FeaturePlanDefinitionMeta.md)
 - [FeaturePlanDefinitionMetaLabelsInner](FeaturePlanDefinitionMetaLabelsInner.md)
 - [FeaturePrice](FeaturePrice.md)
 - [FeaturePricing](FeaturePricing.md)
 - [FeatureSet](FeatureSet.md)
 - [FeatureSets](FeatureSets.md)
 - [FeatureUsage](FeatureUsage.md)
 - [FirewallActionDetails](FirewallActionDetails.md)
 - [FirewallRule](FirewallRule.md)
 - [FirewallRuleResponse](FirewallRuleResponse.md)
 - [FirewallRuleView](FirewallRuleView.md)
 - [FirewallRulesIndex200Response](FirewallRulesIndex200Response.md)
 - [FirewallSettings](FirewallSettings.md)
 - [FirewallSettingsIndex200Response](FirewallSettingsIndex200Response.md)
 - [FirewallSettingsView](FirewallSettingsView.md)
 - [HealthCheck](HealthCheck.md)
 - [HealthCheckReportDetail](HealthCheckReportDetail.md)
 - [HealthCheckReportSummary](HealthCheckReportSummary.md)
 - [HealthCheckReportSummaryDetail](HealthCheckReportSummaryDetail.md)
 - [HealthCheckRequestConfig](HealthCheckRequestConfig.md)
 - [HealthCheckResponse](HealthCheckResponse.md)
 - [HealthCheckView](HealthCheckView.md)
 - [HealthCheckZone](HealthCheckZone.md)
 - [HealthCheckZoneName](HealthCheckZoneName.md)
 - [HealthChecksZonesIndex200Response](HealthChecksZonesIndex200Response.md)
 - [HighRequestedIp](HighRequestedIp.md)
 - [HighRequestedIpIp](HighRequestedIpIp.md)
 - [HttpConfig](HttpConfig.md)
 - [ImageResize](ImageResize.md)
 - [ImageResizeResponse](ImageResizeResponse.md)
 - [ListsIndex200Response](ListsIndex200Response.md)
 - [LoadBalancer](LoadBalancer.md)
 - [LoadBalancerData](LoadBalancerData.md)
 - [LoadBalancerOrigin](LoadBalancerOrigin.md)
 - [LoadBalancerOriginData](LoadBalancerOriginData.md)
 - [LoadBalancerOriginResponse](LoadBalancerOriginResponse.md)
 - [LoadBalancerOriginStore](LoadBalancerOriginStore.md)
 - [LoadBalancerPool](LoadBalancerPool.md)
 - [LoadBalancerPoolData](LoadBalancerPoolData.md)
 - [LoadBalancerPoolResponse](LoadBalancerPoolResponse.md)
 - [LoadBalancerPoolStore](LoadBalancerPoolStore.md)
 - [LoadBalancerRegion](LoadBalancerRegion.md)
 - [LoadBalancerResponse](LoadBalancerResponse.md)
 - [LoadBalancerSetting](LoadBalancerSetting.md)
 - [LoadBalancerSettingsData](LoadBalancerSettingsData.md)
 - [LoadBalancerStore](LoadBalancerStore.md)
 - [LoadBalancerUpdate](LoadBalancerUpdate.md)
 - [LoadBalancerUpdatePoolsInner](LoadBalancerUpdatePoolsInner.md)
 - [LoadBalancersIndex200Response](LoadBalancersIndex200Response.md)
 - [LoadBalancersPoolsIndex200Response](LoadBalancersPoolsIndex200Response.md)
 - [LoadBalancersPoolsOriginsIndex200Response](LoadBalancersPoolsOriginsIndex200Response.md)
 - [LoadBalancersRegionsIndex200Response](LoadBalancersRegionsIndex200Response.md)
 - [LoadBalancersSettingsUpdate200Response](LoadBalancersSettingsUpdate200Response.md)
 - [LogForwarder](LogForwarder.md)
 - [LogForwarderAccessLogType](LogForwarderAccessLogType.md)
 - [LogForwarderDNSType](LogForwarderDNSType.md)
 - [LogForwarderDataFormat](LogForwarderDataFormat.md)
 - [LogForwarderDatadogConnectionType](LogForwarderDatadogConnectionType.md)
 - [LogForwarderErrorType](LogForwarderErrorType.md)
 - [LogForwarderEventType](LogForwarderEventType.md)
 - [LogForwarderGeneric](LogForwarderGeneric.md)
 - [LogForwarderKafkaConnectionType](LogForwarderKafkaConnectionType.md)
 - [LogForwarderLogglyConnectionType](LogForwarderLogglyConnectionType.md)
 - [LogForwarderResponse](LogForwarderResponse.md)
 - [LogForwarderS3ConnectionType](LogForwarderS3ConnectionType.md)
 - [LogForwarderSetting](LogForwarderSetting.md)
 - [LogForwarderSummary](LogForwarderSummary.md)
 - [LogForwarderSyslogConnectionType](LogForwarderSyslogConnectionType.md)
 - [LogForwarderWAFType](LogForwarderWAFType.md)
 - [LogForwardersIndex200Response](LogForwardersIndex200Response.md)
 - [MXRecord](MXRecord.md)
 - [MXRecordValue](MXRecordValue.md)
 - [MapTrafficsData](MapTrafficsData.md)
 - [MessageResponse](MessageResponse.md)
 - [MonitoringStatus](MonitoringStatus.md)
 - [NSRecord](NSRecord.md)
 - [NSRecordValue](NSRecordValue.md)
 - [NextUpstreamTcpCodes](NextUpstreamTcpCodes.md)
 - [NsDomain](NsDomain.md)
 - [NsKeys](NsKeys.md)
 - [NsKeysResponse](NsKeysResponse.md)
 - [PTRRecord](PTRRecord.md)
 - [PTRRecordValue](PTRRecordValue.md)
 - [PageRule](PageRule.md)
 - [PageRuleData](PageRuleData.md)
 - [PageRuleDiff](PageRuleDiff.md)
 - [PageRuleDiffData](PageRuleDiffData.md)
 - [PageRuleDiffRedirect](PageRuleDiffRedirect.md)
 - [PageRuleDiffReqCustomHeadersInner](PageRuleDiffReqCustomHeadersInner.md)
 - [PageRuleImageResize](PageRuleImageResize.md)
 - [PageRuleRedirect](PageRuleRedirect.md)
 - [PageRuleResponse](PageRuleResponse.md)
 - [PageRuleSummary](PageRuleSummary.md)
 - [PageRulesDiffUpdate200Response](PageRulesDiffUpdate200Response.md)
 - [PageRulesIndex200Response](PageRulesIndex200Response.md)
 - [PaginatedResponse](PaginatedResponse.md)
 - [PaginatedResponseLinks](PaginatedResponseLinks.md)
 - [PaginatedResponseMeta](PaginatedResponseMeta.md)
 - [PlanInfo](PlanInfo.md)
 - [PlanResponse](PlanResponse.md)
 - [PlanUpdate](PlanUpdate.md)
 - [PlansIndexDomainParameter](PlansIndexDomainParameter.md)
 - [PrioritizePool](PrioritizePool.md)
 - [PrioritizePoolAfter](PrioritizePoolAfter.md)
 - [PrioritizePoolBefore](PrioritizePoolBefore.md)
 - [PurgeTagsIndex200Response](PurgeTagsIndex200Response.md)
 - [RateLimitRule](RateLimitRule.md)
 - [RateLimitRuleData](RateLimitRuleData.md)
 - [RateLimitRuleView](RateLimitRuleView.md)
 - [RateLimitSettings](RateLimitSettings.md)
 - [RateLimitSettingsData](RateLimitSettingsData.md)
 - [RateLimitingRulesIndex200Response](RateLimitingRulesIndex200Response.md)
 - [RateLimitingRulesUpdate200Response](RateLimitingRulesUpdate200Response.md)
 - [RateLimitingSettingsUpdate200Response](RateLimitingSettingsUpdate200Response.md)
 - [Redirect](Redirect.md)
 - [RedirectData](RedirectData.md)
 - [ReportsAttacksAttackers200Response](ReportsAttacksAttackers200Response.md)
 - [ReportsAttacksAttackers200ResponseDataInner](ReportsAttacksAttackers200ResponseDataInner.md)
 - [ReportsAttacksIndex200Response](ReportsAttacksIndex200Response.md)
 - [ReportsAttacksShow200Response](ReportsAttacksShow200Response.md)
 - [ReportsErrorLogDetails200Response](ReportsErrorLogDetails200Response.md)
 - [ReportsVisitorsHighRequestIps200Response](ReportsVisitorsHighRequestIps200Response.md)
 - [ReprioritizeRuleRequest](ReprioritizeRuleRequest.md)
 - [ResponseTime](ResponseTime.md)
 - [ResponseTimeCharts](ResponseTimeCharts.md)
 - [ResponseTimeChartsIr](ResponseTimeChartsIr.md)
 - [ResponseTimeChartsIrSeriesInner](ResponseTimeChartsIrSeriesInner.md)
 - [ResponseTimeData](ResponseTimeData.md)
 - [ResponseTimeStatistics](ResponseTimeStatistics.md)
 - [SPFRecord](SPFRecord.md)
 - [SRVRecord](SRVRecord.md)
 - [SRVRecordValue](SRVRecordValue.md)
 - [SavedTrafficsCharts](SavedTrafficsCharts.md)
 - [SavedTrafficsChartsRequestInner](SavedTrafficsChartsRequestInner.md)
 - [SavedTrafficsChartsTrafficInner](SavedTrafficsChartsTrafficInner.md)
 - [SavedTrafficsData](SavedTrafficsData.md)
 - [SavedTrafficsStatistics](SavedTrafficsStatistics.md)
 - [SavedTrafficsStatisticsTraffic](SavedTrafficsStatisticsTraffic.md)
 - [Ssl](Ssl.md)
 - [SslCertOrderIndex200Response](SslCertOrderIndex200Response.md)
 - [SslResponse](SslResponse.md)
 - [SslUpdate](SslUpdate.md)
 - [StatusCodeReport](StatusCodeReport.md)
 - [StatusCodeReportCharts](StatusCodeReportCharts.md)
 - [StatusCodeReportChartsStatusCode](StatusCodeReportChartsStatusCode.md)
 - [StatusCodeReportChartsStatusCodeSeriesInner](StatusCodeReportChartsStatusCodeSeriesInner.md)
 - [StatusCodeReportData](StatusCodeReportData.md)
 - [StatusCodeReportStatistics](StatusCodeReportStatistics.md)
 - [StatusCodeReportStatisticsStatusCodes](StatusCodeReportStatisticsStatusCodes.md)
 - [StatusCodeSummary](StatusCodeSummary.md)
 - [StatusCodeSummaryCharts](StatusCodeSummaryCharts.md)
 - [StatusCodeSummaryChartsStatusCodeInner](StatusCodeSummaryChartsStatusCodeInner.md)
 - [StatusCodeSummaryData](StatusCodeSummaryData.md)
 - [TLSARecord](TLSARecord.md)
 - [TLSARecordValue](TLSARecordValue.md)
 - [TXTRecord](TXTRecord.md)
 - [TXTRecordValue](TXTRecordValue.md)
 - [TcpConfig](TcpConfig.md)
 - [TrafficCharts](TrafficCharts.md)
 - [TrafficChartsRequests](TrafficChartsRequests.md)
 - [TrafficChartsRequestsSeriesInner](TrafficChartsRequestsSeriesInner.md)
 - [TrafficChartsTraffics](TrafficChartsTraffics.md)
 - [TrafficChartsTrafficsSeriesInner](TrafficChartsTrafficsSeriesInner.md)
 - [TrafficStatistics](TrafficStatistics.md)
 - [TrafficStatisticsTraffics](TrafficStatisticsTraffics.md)
 - [Traffics](Traffics.md)
 - [TrafficsData](TrafficsData.md)
 - [TrafficsMap](TrafficsMap.md)
 - [TrafficsMapCharts](TrafficsMapCharts.md)
 - [TransferDomain](TransferDomain.md)
 - [TransferDomainChangeStatus](TransferDomainChangeStatus.md)
 - [TransportLayerProxiesIndex200Response](TransportLayerProxiesIndex200Response.md)
 - [TransportLayerProxy](TransportLayerProxy.md)
 - [TransportLayerProxyFirewall](TransportLayerProxyFirewall.md)
 - [TransportLayerProxyFirewallsInner](TransportLayerProxyFirewallsInner.md)
 - [TransportLayerProxyMatch](TransportLayerProxyMatch.md)
 - [TransportLayerProxyResponse](TransportLayerProxyResponse.md)
 - [TransportLayerProxyServer](TransportLayerProxyServer.md)
 - [TransportLayerProxyServersInner](TransportLayerProxyServersInner.md)
 - [TransportLayerProxyStore](TransportLayerProxyStore.md)
 - [TransportLayerProxyTrafficCharts](TransportLayerProxyTrafficCharts.md)
 - [TransportLayerProxyTrafficChartsTraffics](TransportLayerProxyTrafficChartsTraffics.md)
 - [TransportLayerProxyTrafficChartsTrafficsSeriesInner](TransportLayerProxyTrafficChartsTrafficsSeriesInner.md)
 - [TransportLayerProxyTrafficsData](TransportLayerProxyTrafficsData.md)
 - [TransportLayerProxyUpdate](TransportLayerProxyUpdate.md)
 - [Troubleshoot](Troubleshoot.md)
 - [TroubleshootDetailsInner](TroubleshootDetailsInner.md)
 - [TroubleshootsIndex200Response](TroubleshootsIndex200Response.md)
 - [TroubleshootsStore201Response](TroubleshootsStore201Response.md)
 - [UpdateBooleanStatus](UpdateBooleanStatus.md)
 - [UpstreamTimeout](UpstreamTimeout.md)
 - [UsageLimit](UsageLimit.md)
 - [Usages](Usages.md)
 - [Violations](Violations.md)
 - [ViolationsViolations](ViolationsViolations.md)
 - [Visitors](Visitors.md)
 - [VisitorsCharts](VisitorsCharts.md)
 - [VisitorsChartsVisitors](VisitorsChartsVisitors.md)
 - [VisitorsChartsVisitorsSeriesInner](VisitorsChartsVisitorsSeriesInner.md)
 - [VisitorsData](VisitorsData.md)
 - [VisitorsStatistics](VisitorsStatistics.md)
 - [VisitorsStatisticsVisitors](VisitorsStatisticsVisitors.md)
 - [WafPackage](WafPackage.md)
 - [WafPackageDetails](WafPackageDetails.md)
 - [WafPackageDetailsData](WafPackageDetailsData.md)
 - [WafPackageProvider](WafPackageProvider.md)
 - [WafPackagesStore200Response](WafPackagesStore200Response.md)
 - [WafPackagesUpdate200Response](WafPackagesUpdate200Response.md)
 - [WafPreset](WafPreset.md)
 - [WafPresetPackagesInner](WafPresetPackagesInner.md)
 - [WafPresetPackagesInnerProvider](WafPresetPackagesInnerProvider.md)
 - [WafPresets](WafPresets.md)
 - [WafPresetsData](WafPresetsData.md)
 - [WafReconfigure](WafReconfigure.md)
 - [WafReprioritize](WafReprioritize.md)
 - [WafRule](WafRule.md)
 - [WafRuleResponse](WafRuleResponse.md)
 - [WafRulesIndex200Response](WafRulesIndex200Response.md)
 - [WafRuleset](WafRuleset.md)
 - [WafRulesetRulesInner](WafRulesetRulesInner.md)
 - [WafRulesets](WafRulesets.md)
 - [WafSettings](WafSettings.md)
 - [WafSettingsData](WafSettingsData.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### UserToken

- **Type**: HTTP Bearer token authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```

### ApiKey

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



