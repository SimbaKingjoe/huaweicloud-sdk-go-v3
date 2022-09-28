# 0.1.5 2022-09-28

### HuaweiCloud SDK APM

- _Features_
  - Support the interfaces `DeleteEnv`, `DeleteApp`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `sk` to the interface `CreateAkSk`
  - Add the response parameter `sk` to the interface `DeleteAkSk`

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `domain_id` to the interface `ListDomains`
  - Changes of the interface `CreateDomain`:
    - Add the request parameter `domain_id`
    - Add the response parameter `domain_id`
  - Add the response parameter `domain_id` to the interface `ShowDomainDetail`
  - Add the response parameter `domain_id` to the interface `DeleteDomain`
  - Add the response parameter `domain_id` to the interface `EnableDomain`
  - Add the response parameter `domain_id` to the interface `DisableDomain`
  - Changes of the interface `UpdateDomainOrigin`:
    - Add the request parameter `domain_id`
    - Add the response parameter `domain_id`
  - Add the response parameters `origin_range_status`, `user_agent_filter`, `origin_request_url_rewrite`, `error_code_redirect_rules` to the interface `ShowDomainFullConfig`
  - Add the request parameters `origin_range_status`, `user_agent_filter`, `origin_request_url_rewrite`, `error_code_redirect_rules` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CloudRTC

- _Features_
  - Support the interfaces `RemoveUsers`, `RemoveRoom`, `UpdateIndividualStreamJob`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `variable_mode`, `share_mode` to the interface `ListVariables`
  - Changes of the interface `UpdateTask`:
    - Add the request parameters `name`, `conditions`, `is_disabled`
    - Add the response parameter `taskInfo`
    - Remove the request parameter `case_name`
    - Remove the response parameter `taskinfo`
    - Modify the type `string` -> `object` of the request parameter `check_end_length`
    - Modify the type `string` -> `object` of the request parameter `check_end_str`
    - Modify the type `string` -> `object` of the request parameter `check_end_type`
  - Changes of the interface `ShowTask`:
    - Add the response parameter `taskInfo`
    - Remove the response parameter `taskinfo`
  - Add the response parameters `respTimeRange`, `progressState`, `createBy`, `statusValue` to the interface `ShowReport`
  - Add the request parameters `case_id`, `name`, `case_type`, `increase_setting`, `stages`, `status`, `temp_id`, `sort` to the interface `UpdateCase`
  - Changes of the interface `UpdateTemp`:
    - Modify the type `string` -> `object` of the request parameter `check_end_length`
    - Modify the type `string` -> `object` of the request parameter `check_end_str`
    - Modify the type `string` -> `object` of the request parameter `check_end_type`
  - Changes of the interface `UpdateTaskRelatedTestCase`:
    - Add the response parameter `taskInfo`
    - Remove the response parameter `taskinfo`
  - Add the response parameters `end_time`, `parallel` to the interface `ShowHistoryRunInfo`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowInstance`:
    - Modify the type `string` -> `int64` of the response parameter `begin_time`
    - Modify the type `string` -> `int64` of the response parameter `end_time`
    - Modify the type `string` -> `int64` of the response parameter `current_time`
    - Modify the type `string` -> `int64` of the response parameter `next_expand_time`
    - Modify the type `string` -> `int64` of the response parameter `expand_effect_time`
    - Modify the type `string` -> `int64` of the response parameter `expand_interval_time`
  - Modify the type `int32` -> `integer` of the request parameter `new_capacity` of the interface `ResizeInstance`
  - Add the response parameters `target_instance_address`, `migration_method`, `task_name`, `target_instance_id`, `source_instance_name`, `target_instance_name`, `migrate_type`, `created_at`, `source_instance_id`, `task_id`, `data_source`, `status` to the interface `ListMigrationTask`
  - Changes of the interface `ListRedislog`:
    - Add the response parameter `backup_id`
    - Remove the response parameter `backupId`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `region` changed to required of the interface `ListFlavors`

### HuaweiCloud SDK IEF

- _Features_
  - Support the following interfaces：
    - `RestartDeploymentsPod`
    - `ShowEdgeNodeUpgradeDetails`
    - `UpgradeEdgeNode`
    - `ListEncryptdatas`
    - `CreateEncryptdatas`
    - `ShowEncryptdatas`
    - `UpdateEncryptdatas`
    - `DeleteEncryptdatas`
    - `ListNodeEncryptdatas`
    - `CreateNodeEncryptdatas`
    - `DeleteNodeEncryptdatas`
    - `ListEncryptdataNodes`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateEdgeNode`:
    - Add the request parameter `UpdateEdgeNodeRequestBody`
    - Remove the request parameter `UpdateEdgeNodeBody`
  - Changes of the interface `EnableDisableEdgeNodes`:
    - Add the request parameter `EnableDisableEdgeNodesRequestBody`
    - Remove the request parameter `node`
  - Add the response parameter `host_pid` to the interface `ListApps`
  - Add the response parameter `host_pid` to the interface `UpdateApp`
  - Add the response parameter `host_pid` to the interface `ShowAppDetail`
  - Add the request parameter `host_pid` to the interface `CreateAppVersions`
  - Add the response parameter `host_pid` to the interface `ListAppVersions`
  - Changes of the interface `UpdateAppVersion`:
    - Add the request parameter `host_pid`
    - Add the response parameter `host_pid`
  - Add the response parameter `host_pid` to the interface `ShowAppVersionDetail`

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The response parameter `width`, `height` changed to not required of the interface `ShowTranscodingsTemplate`
  - Changes of the interface `UpdateTranscodingsTemplate`:
    - Add the request parameter `trans_type`
    - The request parameter `width`, `height` changed to not required
  - Changes of the interface `CreateTranscodingsTemplate`:
    - Add the request parameter `trans_type`
    - The request parameter `width`, `height` changed to not required

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `type`, `idcard_number`, `phone_number`, `province`, `city`, `vaccination_status`, `pcr_test_result`, `pcr_test_organization`, `pcr_test_time`, `pcr_sampling_time`, `reached_city` to the interface `RecognizeHealthCode`

### HuaweiCloud SDK VPCEP

- _Features_
  - Support the following interfaces：
    - `UpdateEndpointServiceName`
    - `UpdateEndpointConnectionsDesc`
    - `BatchAddEndpointServicePermissions`
    - `BatchRemoveEndpointServicePermissions`
    - `UpdateEndpointServicePermissionDesc`
    - `UpdateEndpointPolicy`
    - `DeleteEndpointPolicy`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateEndpointService`:
    - Add the request parameter `description`
    - Add the response parameter `description`
  - Changes of the interface `ListEndpointService`:
    - Add the request parameter `public_border_group`
    - Add the response parameters `description`, `public_border_group`
    - Modify the type `string` -> `enum` of the response parameter `service_type`
    - Modify the type `enum` -> `string` of the response parameter `server_type`
  - Changes of the interface `UpdateEndpointService`:
    - Add the request parameter `description`
    - Add the response parameter `description`
  - Changes of the interface `ListServiceDetails`:
    - Add the response parameter `description`
    - Modify the type `string` -> `enum` of the response parameter `service_type`
  - Remove the response parameters `id`, `marker_id`, `created_at`, `updated_at`, `domain_id`, `status` from the interface `ListServiceConnections`
  - Add the response parameter `description` to the interface `AcceptOrRejectEndpoint`
  - Remove the response parameters `id`, `permission`, `created_at` from the interface `ListServicePermissionsDetails`
  - Changes of the interface `CreateEndpoint`:
    - Add the request parameter `description`
    - Add the response parameters `specification_name`, `description`, `policy_statement`, `enable_status`
  - Changes of the interface `ListEndpoints`:
    - Add the request parameter `public_border_group`
    - Add the response parameters `description`, `policy_statement`, `endpoint_pool_id`, `public_border_group`
  - Add the response parameters `description`, `policy_statement` to the interface `ListEndpointInfoDetails`
  - Remove the response parameters `status`, `id`, `updated`, `version`, `min_version`, `links` from the interface `ListVersionDetails`
  - Remove the response parameters `status`, `id`, `updated`, `version`, `min_version`, `links` from the interface `ListSpecifiedVersionDetails`
  - Changes of the interface `ListResourceInstances`:
    - Add the request parameters `sys_tags`, `without_any_tag`
    - Remove the request parameters `key`, `value`, `key`, `value`, `key`, `value`, `key`, `value`

# 0.1.4 2022-09-26

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `version`, `last_modified` to the interface `CreateDependency`
  - Add the response parameters `version`, `last_modified` to the interface `ListDependencies`

### HuaweiCloud SDK Kafka

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `disk_encrypted_key`, `public_management_connect_address`, `subnet_cidr`, `subnet_name`, `enable_acl` to the interface `ListInstances`
  - Add the response parameters `disk_encrypted_key`, `public_management_connect_address`, `subnet_cidr`, `subnet_name`, `enable_acl` to the interface `ShowInstance`
  - Add the request parameters `oper_type`, `new_broker_num`, `new_product_id` to the interface `ResizeInstance`

### HuaweiCloud SDK Moderation

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the response parameters' type of the interface `CheckImageModeration` is incorrect
- _Change_
  - None

# 0.1.3 2022-09-22

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `network_type`, `alias_urn`, `network_type`, `alias_urn` to the interface `CreateApiV2`
  - Add the response parameters `network_type`, `alias_urn`, `network_type`, `alias_urn` to the interface `ShowDetailsOfApiV2`
  - Changes of the interface `UpdateApiV2`:
    - Add the request parameters `network_type`, `alias_urn`, `network_type`, `alias_urn`
    - Add the response parameters `network_type`, `alias_urn`, `network_type`, `alias_urn`
  - Add the response parameters `network_type`, `alias_urn`, `network_type`, `alias_urn` to the interface `ListApiVersionDetailV2`
  - Remove the response parameters `url_domain`, `id`, `status`, `min_ssl_version` from the interface `UpdateDomainV2`
  - Add the response parameters `req_uri`, `auth_type` to the interface `ListApisUnbindedToAclPolicyV2`
  - Add the response parameters `authorizer_version`, `authorizer_alias_uri` to the interface `ListCustomAuthorizersV2`
  - Add the request parameters `authorizer_version`, `authorizer_alias_uri` to the interface `CreateCustomAuthorizerV2`
  - Add the response parameters `authorizer_version`, `authorizer_alias_uri` to the interface `ShowDetailsOfCustomAuthorizersV2`
  - Changes of the interface `UpdateCustomAuthorizerV2`:
    - Add the request parameters `authorizer_version`, `authorizer_alias_uri`
    - Add the response parameters `authorizer_version`, `authorizer_alias_uri`
  - The request parameter `env_id` changed to required of the interface `ExportApiDefinitionsV2`
  - Changes of the interface `ListTagsV2`:
    - Add the response parameter `tags`
    - Remove the response parameter `responses`
  - Remove the response parameters `id`, `name`, `enable`, `config`, `instance_id`, `update_time` from the interface `CreateFeatureV2`
  - Add the response parameter `ingress_ip_v6` to the interface `ShowDetailsOfInstanceV2`
  - Add the response parameter `ingress_ip_v6` to the interface `UpdateInstanceV2`

### HuaweiCloud SDK CES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `event_type` to the interface `CreateEvents`
  - Add the response parameter `event_type` to the interface `ListEventDetail`

### HuaweiCloud SDK CloudPipeline

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interfaces `ShowAgentStatus`, `RegisterAgent`

### HuaweiCloud SDK CodeCheck

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `check_end_time` to the interface `CheckRecord`
  - Add the response parameter `events` to the interface `ShowTaskDefects`
  - Remove the response parameters `name`, `cfg_key`, `default_value`, `option_value`, `is_required`, `description`, `type`, `status` from the interface `CheckParameters`
  - Add the response parameter `value` to the interface `CheckRulesetParameters`

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListProjectSets`:
    - Add the response parameter `source`
    - Remove the response parameter `status`
  - Add the response parameter `parallel` to the interface `ShowTaskSet`

### HuaweiCloud SDK CTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `group_id`, `stream_id` to the interface `ListTrackers`

### HuaweiCloud SDK DDM

- _Features_
  - Support the interfaces `ResetAdministrator`, `ValidateWeakPassword`, `ResizeFlavor`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `admin_user_name`, `admin_user_password` to the interface `CreateInstance`
  - Add the response parameter `admin_user_name` to the interface `ShowInstance`
  - Add the response parameter `host` to the interface `ListSlowLog`

### HuaweiCloud SDK DevStar

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `none` to the response parameter `deploy_type` to the interface `ShowApplicationV3`
  - Add the response parameter `subscribe_guide` to the interface `ShowApplicationDependentResources`
  - Add the enum values `none` to the response parameter `deploy_type` to the interface `ListApplicationsV6`
  - Add the response parameter `category_name` to the interface `ShowApplicationReleaseRepositories`
  - Add the response parameter `subscribe_guide` to the interface `ShowTemplateV3`
  - Add the response parameter `subscribe_guide` to the interface `ListTemplatesV2`
  - Add the response parameter `subscribe_guide` to the interface `ListTemplates`

### HuaweiCloud SDK EG

- _Features_
  - Support the following interfaces：
    - `ListApiVersions`
    - `ListEventSchema`
    - `CreateEventSchema`
    - `ShowDetailOfEventSchema`
    - `UpdateEventSchema`
    - `DeleteEventSchema`
    - `ListEventSchemaVersions`
    - `CreateEventSchemaVersion`
    - `ShowDetailOfEventSchemaVersion`
    - `DeleteEventSchemaVersion`
    - `DiscoverEventSchemaFromData`
    - `ListConnections`
    - `CreateConnection`
    - `ShowDetailOfConnection`
    - `UpdateConnection`
    - `DeleteConnection`
    - `ListAgencies`
    - `CreateAgencies`
    - `ListTriggers`
    - `UpdateEndpoint`
    - `DeleteEndpoint`
    - `ListEndpoints`
    - `CreateEndpoint`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `id`, `name`, `label`, `description`, `provider_type`, `event_types`, `created_time`, `updated_time`, `channel_id`, `channel_name` from the interface `ListEventSources`
  - Changes of the interface `CreateEventSource`:
    - Add the request parameters `type`, `detail`
    - Add the response parameters `event_types`, `type`, `detail`, `status`
  - Add the response parameters `event_types`, `type`, `detail`, `status` to the interface `ShowDetailOfEventSource`
  - Changes of the interface `UpdateEventSource`:
    - Add the request parameter `detail`
    - Add the response parameters `event_types`, `type`, `detail`, `status`
  - Remove the response parameters `name`, `label`, `metadata` from the interface `ListEventTarget`
  - Add the response parameter `connection_id` to the interface `ListSubscriptions`
  - Changes of the interface `CreateSubscription`:
    - Add the request parameter `connection_id`
    - Add the response parameter `connection_id`
  - Add the response parameter `connection_id` to the interface `ShowDetailOfSubscription`
  - Changes of the interface `UpdateSubscription`:
    - Add the request parameter `connection_id`
    - Add the response parameter `connection_id`
  - Changes of the interface `CreateSubscriptionTarget`:
    - Add the request parameter `connection_id`
    - Add the response parameter `connection_id`
  - Add the response parameter `connection_id to the interface `ShowDetailOfSubscriptionTarget`
  - Changes of the interface `UpdateSubscriptionTarget`:
    - Add the request parameter `connection_id`
    - Add the response parameter `connection_id`
  - Changes of the interface `ListQuotas`:
    - Add the enum values `CONNECTION`, `PRIVATE_ENDPOINT`, `SOURCE_RABBITMQ` to the request parameter `type`
    - Add the enum values `CONNECTION`, `PRIVATE_ENDPOINT`, `SOURCE_RABBITMQ` to the response parameter `type`
    - Modify the type `string` -> `int32` of the response parameter `max`
    - Modify the type `string` -> `int32` of the response parameter `min`
    - Modify the type `string` -> `int32` of the response parameter `quota`
    - Modify the type `string` -> `int32` of the response parameter `used`

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `port_id` to the interface `CreatePublicip`
  - Add the request parameter `port_id` to the interface `CreatePrePaidPublicip`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interface `ListGaussMySqlInstanceDetailInfo`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IoTDA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `error_info` to the interface `ListDeviceMessages`
  - Add the response parameter `error_info` to the interface `ShowDeviceMessage`

### HuaweiCloud SDK Moderation

- _Features_
  - Support the interface `CheckImageModeration`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `url`, `categories` changed to required of the interface `RunCreateAudioModerationJob`

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RecognizeIdCard`:
    - Add the request parameter `detect_copy`
    - Add the response parameter `detect_copy_result`

### HuaweiCloud SDK OMS

- _Features_
  - Support the following interfaces：
    - `ListTaskGroup`
    - `CreateTaskGroup`
    - `ShowTaskGroup`
    - `DeleteTaskGroup`
    - `StopTaskGroup`
    - `StartTaskGroup`
    - `RetryTaskGroup`
    - `UpdateTaskGroup`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `enable_metadata_migration`, `object_overwrite_mode`, `consistency_check`, `enable_requester_pays` to the interface `ListTasks`
  - Changes of the interface `CreateTask`:
    - Add the request parameters `enable_metadata_migration`, `object_overwrite_mode`, `consistency_check`, `enable_requester_pays`
    - Add the response parameters `id`, `task_name`
  - Changes of the interface `ShowTask`:
    - Add the response parameters `enable_metadata_migration`, `object_overwrite_mode`, `consistency_check`, `enable_requester_pays`
    - Modify the type `int64` -> `string` of the request parameter `task_id`
  - Modify the type `int64` -> `string` of the request parameter `task_id` of the interface `DeleteTask`
  - Modify the type `int64` -> `string` of the request parameter `task_id` of the interface `StopTask`
  - Modify the type `int64` -> `string` of the request parameter `task_id` of the interface `StartTask`
  - Modify the type `int64` -> `string` of the request parameter `task_id` of the interface `UpdateBandwidthPolicy`

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interfaces `DownloadImageFile`, `ListScrumProjectStatuses`, `UploadAttachments`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RabbitMQ

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `rabbitmq`, Remove the enum values `true`, `false` from the request parameter `all_failure` to the interface `BatchRestartOrDeleteInstances`

### HuaweiCloud SDK SMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `disks` from the interface `ListTemplates`
  - Remove the response parameter `disks` from the interface `ShowTemplate`
  - Remove the request parameter `disks` from the interface `UpdateMigproject`
  - Remove the response parameter `disks` from the interface `ShowMigproject`

### HuaweiCloud SDK VOD

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `pack_type`, `pack_type` to the interface `PublishAssets`
  - Add the response parameters `pack_type`, `pack_type` to the interface `UnpublishAssets`
  - Add the response parameters `pack_type`, `pack_type` to the interface `ShowAssetMeta`
  - Add the response parameters `pack_type`, `pack_type` to the interface `ShowAssetDetail`
  - Add the response parameters `pack_type`, `pack_type` to the interface `ShowTakeOverTaskDetails`
  - Add the response parameters `pack_type`, `pack_type` to the interface `ShowTakeOverAssetDetails`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `rule` to the interface `DeleteIgnoreRule`

# 0.1.2 2022-09-15

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `root_resource_id`, `parent_resource_id`, `trade_id`, `product_spec_desc` to the interface `ListCustomerselfResourceRecordDetails`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interfaces `ShowTags`, `CreateTags`, `BatchDeleteTags`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowUrlTaskInfo`:
    - Add the response parameter `result`
    - Remove the response parameter `results`
  - Add the response parameter `error_code_cache` to the interface `ShowDomainFullConfig`
  - Add the request parameter `error_code_cache` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the following interfaces：
    - `ListFilesByQuery`
    - `ListBranchesByRepositoryId`
    - `CreateNewBranch`
    - `AssociateIssues`
    - `ListMergeRequest`
    - `ShowMergeRequest`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSubfiles`:
    - Add the response parameters `error`, `result`, `status`
    - Remove the response parameters `trees`, `total`
  - Changes of the interface `ShowStatisticalData`:
    - Add the response parameters `error`, `result`, `status`
    - Remove the response parameters `repoName`, `commitCount`, `repoSize`, `lastCommitTime`, `codeLines`, `branchCount`, `archiveUrl`
  - Modify the type `string` -> `boolean` of the request parameter `force` of the interface `CreateCommit`
  - Changes of the interface `AddProtectBranchV2`:
    - Modify the type `int32` -> `enum` of the request parameter `push_access_level`
    - Modify the type `int32` -> `enum` of the request parameter `merge_access_level`

### HuaweiCloud SDK CSE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `create_revision`, `update_revision` to the interface `UploadKie`

### HuaweiCloud SDK CSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `CloseKibanaPublicReq` to the interface `UpdateCloseKibana`
  - Add the request parameter `payInfo` to the interface `CreateCluster`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `enterprise_project_name`, `update_at`, `product_type`, `storage_type`, `launched_at`, `cache_mode`, `support_slow_log_flag`, `db_number`, `replica_count`, `sharding_count`, `bandwidth_info` to the interface `ShowInstance`
  - Add the response parameter `backupId` to the interface `ListRedislog`
  - Add the response parameter `instance_id` to the interface `ShowIpWhitelist`
  - Add the request parameter `instance_id` to the interface `UpdateIpWhitelist`
  - Add the response parameters `updated_at`, `created_at`, `status` to the interface `ListBackgroundTask`

### HuaweiCloud SDK EVS

- _Features_
  - Support the following interfaces：
    - `ShowVersion`
    - `ListVersions`
    - `CinderShowVolumeTransfer`
    - `CinderDeleteVolumeTransfer`
    - `CinderListVolumeTransfers`
    - `CinderCreateVolumeTransfer`
    - `CinderAcceptVolumeTransfer`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IEF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ListApps`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `UpdateApp`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ShowAppDetail`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ListAppVersions`
  - Add the request parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `CreateAppVersions`
  - Changes of the interface `UpdateAppVersion`:
    - Add the request parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu`
    - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ShowAppVersionDetail`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `ttl_in_days` to the interface `CreateLogStream`
  - Add the request parameter `whether_to_rows` to the interface `ListStructuredLogsWithTimeRange`
  - The request parameter `isAnalysis` changed to not required of the interface `UpdateStructTemplate`
  - The request parameter `isAnalysis` changed to not required of the interface `CreateStructTemplate`

### HuaweiCloud SDK Moderation

- _Features_
  - Support the interfaces `RunCreateVideoModerationJob`, `RunQueryVideoModerationJob`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `url` changed to not required of the interface `RunCreateAudioModerationJob`

# 0.1.1 2022-09-08

### HuaweiCloud SDK IAM

- _Features_
  - Support the custom credential `IamCredentials`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `root_resource_id`, `parent_resource_id`, `trade_id`, `product_spec_desc` to the interface `ListCustomerselfResourceRecordDetails`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interfaces `ShowTags`, `CreateTags`, `BatchDeleteTags`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowUrlTaskInfo`:
    - Add the response parameter `result`
    - Remove the response parameter `results`
  - Add the response parameter `error_code_cache` to the interface `ShowDomainFullConfig`
  - Add the request parameter `error_code_cache` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the following interfaces：
    - `ListFilesByQuery`
    - `ListBranchesByRepositoryId`
    - `CreateNewBranch`
    - `AssociateIssues`
    - `ListMergeRequest`
    - `ShowMergeRequest`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSubfiles`:
    - Add the response parameters `error`, `result`, `status`
    - Remove the response parameters `trees`, `total`
  - Changes of the interface `ShowStatisticalData`:
    - Add the response parameters `error`, `result`, `status`
    - Remove the response parameters `repoName`, `commitCount`, `repoSize`, `lastCommitTime`, `codeLines`, `branchCount`, `archiveUrl`
  - Modify the type `string` -> `boolean` of the request parameter `force` of the interface `CreateCommit`
  - Changes of the interface `AddProtectBranchV2`:
    - Modify the type `int32` -> `enum` of the request parameter `push_access_level`
    - Modify the type `int32` -> `enum` of the request parameter `merge_access_level`

### HuaweiCloud SDK CSE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `create_revision`, `update_revision` to the interface `UploadKie`

### HuaweiCloud SDK CSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `CloseKibanaPublicReq` to the interface `UpdateCloseKibana`
  - Add the request parameter `payInfo` to the interface `CreateCluster`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `enterprise_project_name`, `update_at`, `product_type`, `storage_type`, `launched_at`, `cache_mode`, `support_slow_log_flag`, `db_number`, `replica_count`, `sharding_count`, `bandwidth_info` to the interface `ShowInstance`
  - Add the response parameter `backupId` to the interface `ListRedislog`
  - Add the response parameter `instance_id` to the interface `ShowIpWhitelist`
  - Add the request parameter `instance_id` to the interface `UpdateIpWhitelist`
  - Add the response parameters `updated_at`, `created_at`, `status` to the interface `ListBackgroundTask`

### HuaweiCloud SDK EVS

- _Features_
  - Support the following interfaces：
    - `ShowVersion`
    - `ListVersions`
    - `CinderShowVolumeTransfer`
    - `CinderDeleteVolumeTransfer`
    - `CinderListVolumeTransfers`
    - `CinderCreateVolumeTransfer`
    - `CinderAcceptVolumeTransfer`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IEF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ListApps`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `UpdateApp`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ShowAppDetail`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ListAppVersions`
  - Add the request parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `CreateAppVersions`
  - Changes of the interface `UpdateAppVersion`:
    - Add the request parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu`
    - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu`
  - Add the response parameters `cpu`, `memory`, `gpu`, `npu`, `cpu`, `memory`, `gpu`, `npu` to the interface `ShowAppVersionDetail`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `ttl_in_days` to the interface `CreateLogStream`
  - Add the request parameter `whether_to_rows` to the interface `ListStructuredLogsWithTimeRange`
  - The request parameter `isAnalysis` changed to not required of the interface `UpdateStructTemplate`
  - The request parameter `isAnalysis` changed to not required of the interface `CreateStructTemplate`

### HuaweiCloud SDK Moderation

- _Features_
  - Support the interfaces `RunCreateVideoModerationJob`, `RunQueryVideoModerationJob`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `url` changed to not required of the interface `RunCreateAudioModerationJob`

# 0.0.108 2022-09-01

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - [Issue 56](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/56) Fix the issue of the RequestHandler
- _Change_
  - None

### HuaweiCloud SDK BSSINTL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `indirect_partner_id` to the interface `ListCustomerOrders`
  - Add the request parameter `indirect_partner_id` to the interface `ShowCustomerOrderDetails`
  - Add the request parameter `indirect_partner_id` to the interface `ListCustomerOnDemandResources`

### HuaweiCloud SDK CC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the enum values `er` from the response parameter `used_scene` from the interface `ListCloudConnections`
  - Remove the request parameter `used_scene` from the interface `CreateCloudConnection`
  - Remove the enum values `er` from the response parameter `used_scene` from the interface `ShowCloudConnection`
  - Remove the enum values `er` from the response parameter `used_scene` from the interface `UpdateCloudConnection`
  - Remove the enum values `er` from the response parameter `type` from the interface `ListNetworkInstances`
  - Remove the enum values `er` from the request parameter `type` from the interface `CreateNetworkInstance`
  - Remove the enum values `er` from the response parameter `type` from the interface `ShowNetworkInstance`
  - Remove the enum values `er` from the response parameter `type` from the interface `UpdateNetworkInstance`

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The response parameter `id`, `type` changed to not required of the interface `ShowJobs`
  - The request parameter `id`, `type` changed to not required of the interface `UpdateJob`
  - The request parameter `id`, `type` changed to not required of the interface `CreateAndStartRandomClusterJob`
  - The request parameter  `id`, `type` changed to not required of the interface `CreateJob`

### HuaweiCloud SDK CSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ordeId` to the interface `CreateCluster`
  - Changes of the interface `ShowClusterDetail`:
    - Add the response parameters `vpcepIp`, `elbWhiteListResp`
    - Remove the response parameters `elbWhiteList`, `action`
  - Remove the request parameter `isAutoPay` from the interface `UpdateUnbindPublic`
  - Changes of the interface `ListYmlsJob`:
    - Add the response parameter `configList`
    - Remove the response parameter `configurations`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `version`, `last_modified` from the interface `CreateDependency`
  - Remove the response parameters `version`, `last_modified` from the interface `ListDependencies`

### HuaweiCloud SDK IAM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateDomainProtectPolicy`:
    - Add the request parameters `allow_user`, `mobile`, `admin_check`, `email`, `scene`
    - Remove the response parameter `operation_protection`
  - Remove the response parameter `operation_protection` from the interface `ShowDomainProtectPolicy`
  - The request parameter `maximum_consecutive_identical_chars`, `minimum_password_age`, `minimum_password_length`, `number_of_recent_passwords_disallowed`, `password_not_username_or_invert`, `password_validity_period`, `password_char_combination` changed to not required of the interface `UpdateDomainPasswordPolicy`
  - The request parameter `account_validity_period`, `custom_info_for_login`, `lockout_duration`, `login_failed_times`, `period_with_login_failures`, `session_timeout`, `show_recent_login_info` changed to not required of the interface `UpdateDomainLoginPolicy`
  - Add the enum values `mapping` to the request parameter `type` to the interface `ShowDomainQuota`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `UpdateDbUserComment`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `comment` to the interface `CreateDbUser`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListStatistics`:
    - Modify the type `array` -> `string` of the request parameter `hosts`
    - Modify the type `array` -> `string` of the request parameter `instances`
  - Changes of the interface `ListQpsTimeline`:
    - Modify the type `array` -> `string` of the request parameter `hosts`
    - Modify the type `array` -> `string` of the request parameter `instances`
  - Modify the type `array` -> `string` of the request parameter `instances` of the interface `ListBandwidthTimeline`
  - Changes of the interface `ListTopAbnormal`:
    - Modify the type `array` -> `string` of the request parameter `hosts`
    - Modify the type `array` -> `string` of the request parameter `instances`
  - Changes of the interface `ListOverviewsClassification`:
    - Modify the type `array` -> `string` of the request parameter `hosts`
    - Modify the type `array` -> `string` of the request parameter `instances`

# 0.0.107 2022-08-29

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `batch_create_in_multi_az` to the interface `CreatePostPaidServers`

### HuaweiCloud SDK IMS

- _Features_
  - Support the interface `ShowJobProgress`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `SetReadOnlySwitch`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `FailoverRequestBody` to the interface `StartFailover`

# 0.0.106 2022-08-25

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateJob`:
    - Add the request parameters `rows_written`, `rows_read`, `files_written`, `extended-configs`, `value`, `extended-configs`, `value`, `extended-configs`, `value`
    - Remove the request parameters `files_writte`, `values`, `values`, `values`
  - Changes of the interface `ShowJobs`:
    - Add the response parameters `rows_written`, `rows_read`, `files_written`, `extended-configs`, `value`, `extended-configs`, `value`, `extended-configs`, `value`
    - Remove the response parameters `files_writte`, `values`, `values`, `values`
  - Changes of the interface `CreateAndStartRandomClusterJob`:
    - Add the request parameters `rows_written`, `rows_read`, `files_written`, `extended-configs`, `value`, `extended-configs`, `value`, `extended-configs`, `value`
    - Remove the request parameters `files_writte`, `values`, `values`, `values`
  - Changes of the interface `CreateJob`:
    - Add the request parameters `rows_written`, `rows_read`, `files_written`, `extended-configs`, `value`, `extended-configs`, `value`, `extended-configs`, `value`
    - Remove the request parameters `files_writte`, `values`, `values`, `values`

### HuaweiCloud SDK ELB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the following interfaces：
    - `ListMasterSlavePools`
    - `CreateMasterSlavePool`
    - `ShowMasterSlavePool`
    - `DeleteMasterSlavePool`
  - Add the response parameter `waf_failure_action` to the interface `ListLoadBalancers`
  - Add the request parameter `waf_failure_action` to the interface `CreateLoadBalancer`
  - Add the response parameter `waf_failure_action` to the interface `ShowLoadBalancer`
  - Changes of the interface `UpdateLoadBalancer`:
    - Add the request parameter `waf_failure_action`
    - Add the response parameter `waf_failure_action`
    - Remove the request parameter `cloud_service_console_url`
  - Add the response parameters `enc_certificate`, `enc_private_key` to the interface `ListCertificates`
  - Add the request parameters `enc_certificate`, `enc_private_key` to the interface `CreateCertificate`
  - Add the response parameters `enc_certificate`, `enc_private_key` to the interface `ShowCertificate`
  - Changes of the interface `UpdateCertificate`:
    - Add the request parameters `enc_certificate`, `enc_private_key`
    - Add the response parameters `enc_certificate`, `enc_private_key`
  - Add the response parameter `sni_match_algo` to the interface `ListListeners`
  - Add the request parameter `sni_match_algo` to the interface `CreateListener`
  - Add the response parameter `sni_match_algo` to the interface `ShowListener`
  - Changes of the interface `UpdateListener`:
    - Add the request parameter `sni_match_algo`
    - Add the response parameter `sni_match_algo`
  - Add the request parameter `instance_id` to the interface `ListMembers`

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RecognizeWebImage`:
    - Add the request parameter `detect_font`
    - Add the response parameters `font_list`, `font_scores`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `SetDatabaseUserPrivilege`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VOD

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int32` -> `int64` of the request parameter `size` of the interface `CheckMd5Duplication`

### HuaweiCloud SDK WAF

- _Features_
  - Support the interface `ListRequestTimeline`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `description` to the interface `UpdateGeoipRule`

# 0.0.105 2022-08-22

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `uri` -> `string` of the request parameter `object` of the interface `CreateRecordIndex`

# 0.0.104 2022-08-18

### HuaweiCloud SDK BSSINTL

- _Features_
  - Support the interfaces `ListIndirectPartners`, `ListCosts`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `indirect_partner_id` to the interface `ListSubCustomers`
  - Add the request parameter `indirect_partner_id` to the interface `CreateSubCustomer`
  - Add the request parameter `indirect_partner_id` to the interface `ShowSubCustomerBudget`
  - Add the request parameter `indirect_partner_id` to the interface `UpdateSubCustomerBudget`
  - Add the request parameter `indirect_partner_id` to the interface `FreezeSubCustomers`
  - Add the request parameter `indirect_partner_id` to the interface `UnfreezeSubCustomers`

### HuaweiCloud SDK CloudIDE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `arch` to the interface `ListProjectTemplates`

### HuaweiCloud SDK CPH

- _Features_
  - Support the service `Cloud Phone`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ECS

- _Features_
  - Support the interface `ListServersByTag`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `destination_type` changed to required of the interface `NovaCreateServers`

### HuaweiCloud SDK EG

- _Features_
  - Support the service `EventGrid`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateFunction`:
    - Add the response parameters `enable_dynamic_memory`, `is_stateful_function`, `enable_auth_in_header`, `custom_image`
    - The request parameter `file`, `link` changed to required
    - Modify the type `int32` -> `string` of the response parameter `user_id`
    - Modify the type `int32` -> `string` of the response parameter `user_group_id`
    - The response parameter `concurrent_num` changed to required
    - The response parameter `mount_share_path` changed to not required
  - Changes of the interface `ListFunctions`:
    - Add the response parameter `fail_count`
    - Remove the request parameter `X-Auth-Token`
    - The response parameter `concurrent_num` changed to required
  - Changes of the interface `ShowFunctionCode`:
    - Remove the response parameter `id`
    - The response parameter `file`, `link`, `concurrent_num` changed to required
  - Changes of the interface `UpdateFunctionCode`:
    - Remove the response parameter `id`
    - The request parameter `file`, `link` changed to required
    - The response parameter `file`, `link`, `concurrent_num` changed to required
  - Changes of the interface `ShowFunctionConfig`:
    - Add the response parameters `is_stateful_function`, `enable_auth_in_header`, `custom_image`
    - Remove the response parameter `id`
    - Modify the type `int32` -> `string` of the response parameter `user_id`
    - Modify the type `int32` -> `string` of the response parameter `user_group_id`
    - The response parameter `concurrent_num` changed to required
    - The response parameter `mount_share_path` changed to not required
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the response parameters `enable_auth_in_header`, `custom_image`
    - Remove the request parameter `X-Auth-Token`
    - Remove the response parameter `id`
    - Modify the type `int32` -> `string` of the request parameter `user_id`
    - Modify the type `int32` -> `string` of the request parameter `user_group_id`
    - The request parameter `concurrent_num` changed to required
    - The request parameter `mount_share_path` changed to not required
    - Modify the type `int32` -> `string` of the response parameter `user_id`
    - Modify the type `int32` -> `string` of the response parameter `user_group_id`
    - The response parameter `concurrent_num` changed to required
    - The response parameter `mount_share_path` changed to not required
  - Changes of the interface `UpdateFunctionMaxInstanceConfig`:
    - Remove the response parameter `id`
    - Modify the type `int32` -> `string` of the response parameter `user_id`
    - Modify the type `int32` -> `string` of the response parameter `user_group_id`
    - The response parameter `concurrent_num` changed to required
    - The response parameter `mount_share_path` changed to not required
  - Changes of the interface `CreateFunctionVersion`:
    - Remove the response parameter `id`
    - Modify the type `int32` -> `string` of the response parameter `user_id`
    - Modify the type `int32` -> `string` of the response parameter `user_group_id`
    - The response parameter `concurrent_num` changed to required
    - The response parameter `mount_share_path` changed to not required
  - Changes of the interface `ListFunctionVersions`:
    - Add the response parameters `is_stateful_function`, `enable_auth_in_header`, `custom_image`, `reserved_instance_idle_mode`
    - Remove the response parameters `log_group_id`, `log_stream_id`
    - The response parameter `concurrent_num` changed to required
  - Add the enum values `SMN`, `RABBITMQ`, `DEDICATEDGATEWAY`, `OPENSOURCEKAFKA`, `APIC`, `GAUSSMONGO`, `EVENTGRID` to the request parameter `trigger_type_code` to the interface `CreateFunctionTrigger`
  - Changes of the interface `ListFunctionTriggers`:
    - Add the enum values `RABBITMQ`, `DEDICATEDGATEWAY`, `OPENSOURCEKAFKA`, `APIC`, `GAUSSMONGO`, `EVENTGRID` to the response parameter `trigger_type_code`
    - Add the enum values `DISABLE`, Remove the enum values `DISABLED` from the response parameter `trigger_status`
  - Add the enum values `RABBITMQ`, `DEDICATEDGATEWAY`, `OPENSOURCEKAFKA`, `APIC`, `GAUSSMONGO`, `EVENTGRID` to the request parameter `trigger_type_code` to the interface `DeleteFunctionTrigger`
  - Changes of the interface `ShowFunctionTrigger`:
    - Add the enum values `RABBITMQ`, `DEDICATEDGATEWAY`, `OPENSOURCEKAFKA`, `APIC`, `GAUSSMONGO`, `EVENTGRID` to the request parameter `trigger_type_code`
    - Add the enum values `RABBITMQ`, `DEDICATEDGATEWAY`, `OPENSOURCEKAFKA`, `APIC`, `GAUSSMONGO`, `EVENTGRID` to the response parameter `trigger_type_code`
    - Add the enum values `DISABLE`, Remove the enum values `DISABLED` from the response parameter `trigger_status`
  - Changes of the interface `UpdateTrigger`:
    - Add the enum values `DISABLE`, Remove the enum values `DISABLED` from the request parameter `trigger_status`
    - Add the enum values `RABBITMQ`, `DEDICATEDGATEWAY`, `OPENSOURCEKAFKA`, `APIC`, `GAUSSMONGO`, `EVENTGRID` to the request parameter `trigger_type_code`
    - The request parameter `trigger_status` changed to not required
  - Modify the type `float` -> `int32` of the response parameter `value` of the interface `ListStatistics`
  - Changes of the interface `UpdateFunctionReservedInstancesCount`:
    - Add the request parameter `UpdateFunctionReservedInstancesCountRequestBody`
    - Add the response parameters `idle_mode`, `tactics_config`
    - Remove the request parameter `UpdateFunctionReservedInstancesRequestBody`
  - Changes of the interface `CreateDependency`:
    - Add the response parameters `version`, `last_modified`
    - Modify the type `enum` -> `string` of the response parameter `runtime`
  - Changes of the interface `ListDependencies`:
    - Add the request parameters `maxitems`, `ispublic`
    - Add the response parameters `version`, `last_modified`
    - Modify the type `int32` -> `int64` of the response parameter `count`
  - Changes of the interface `ShowDependcy`:
    - Add the response parameters `version`, `last_modified`
    - Modify the type `enum` -> `string` of the response parameter `runtime`
  - Changes of the interface `UpdateDependcy`:
    - Add the request parameter `UpdateDependcyRequestBody`
    - Remove the request parameter `UpdateDependencyRequestBody`
    - Modify the type `enum` -> `string` of the response parameter `runtime`
  - Remove the response parameters `content`, `last_modified` from the interface `CreateEvent`
  - Remove the response parameters `content`, `last_modified` from the interface `UpdateEvent`
  - Changes of the interface `ImportFunction`:
    - Add the request parameter `package`
    - Remove the request parameter `X-Auth-Token`
    - The response parameter `concurrent_num` changed to required
  - Add the request parameter `X-Auth-Token` to the interface `EnableLtsLogs`
  - Add the response parameter `group_name` to the interface `ShowLtsLogDetails`
  - The request parameter `request_id` changed to required of the interface `CancelAsyncInvocation`
  - Changes of the interface `CreateWorkflow`:
    - Add the request parameter `duration`
    - Add the enum values `SMN`, `APIG`, `APIG_DE` to the request parameter `trigger_type`
  - Changes of the interface `ListWorkflow`:
    - Add the request parameters `enterprise_project`, `mode`
    - Remove the response parameters `id`, `workflow_urn`, `name`, `description`, `created_time`, `updated_time`, `created_by`
  - Add the request parameter `X-WorkflowRun-MergeFnParameters` to the interface `StartWorkflowExecution`
  - Remove the response parameters `workflow_id`, `workflow_urn`, `execution_id`, `status`, `begin_time`, `end_time`, `last_update_time`, `created_by` from the interface `ListWorkflowExecutions`
  - Changes of the interface `ShowWorkflowExecution`:
    - Add the request parameter `X-Get-Workflow-Full-History-Data`
    - Modify the type `string` -> `object` of the response parameter `workflow_urn`
  - Changes of the interface `ShowWorkFlow`:
    - Remove the response parameters `name`, `description`, `triggers`, `start`, `functions`, `states`, `constants`, `retries`, `mode`, `express_config`, `enterprise_project_id`
    - Modify the type `string` -> `object` of the response parameter `workflow_urn`
    - The response parameter `id`, `workflow_urn`, `created_time`, `updated_time`, `created_by` changed to required
  - Changes of the interface `UpdateWorkFlow`:
    - Add the request parameter `duration`
    - Add the enum values `SMN`, `APIG`, `APIG_DE` to the request parameter `trigger_type`
    - Modify the type `string` -> `object` of the response parameter `workflow_urn`
    - The response parameter `id`, `workflow_urn`, `name`, `description`, `created_time`, `updated_time`, `created_by` changed to required

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RecognizeIdCard`:
    - Add the request parameter `detect_reproduce`
    - Add the response parameter `detect_reproduce_result`

### HuaweiCloud SDK ProjectMan

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateIssueV4`:
    - Add the request parameter `field_name`
    - Add the response parameter `field_name`
  - Add the response parameter `field_name` to the interface `ListIssuesV4`
  - Changes of the interface `UpdateIssueV4`:
    - Add the request parameter `field_name`
    - Add the response parameter `field_name`
  - Add the response parameter `field_name` to the interface `ListChildIssuesV4`
  - Changes of the interface `CreateSystemIssueV4`:
    - Add the request parameter `field_name`
    - Add the response parameter `field_name`

### HuaweiCloud SDK ROMA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `limit`, `offset` to the interface `ListNotification`

### HuaweiCloud SDK VOD

- _Features_
  - Support the following interfaces：
    - `ListTranscodeTemplate`
    - `UpdateTranscodeTemplate`
    - `CreateTranscodeTemplate`
    - `DeleteTranscodeTemplate`
    - `ListTemplateGroupCollection`
    - `UpdateTemplateGroupCollection`
    - `CreateTemplateGroupCollection`
    - `DeleteTemplateGroupCollection`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListIgnoreRule`:
    - Add the response parameter `domain`
    - Remove the response parameter `domains`
  - Add the response parameter `policyid` to the interface `ListGeoipRule`
  - Add the request parameter `description` to the interface `UpdateGeoipRule`

# 0.0.103 2022-08-11

### HuaweiCloud SDK APM

- _Features_
  - Support the following interfaces：
    - `ListOpenRegion`
    - `ListSupportedRegion`
    - `ShowTopologyTree`
    - `ShowMonitorItemViewConfig`
    - `ListEnvTags`
    - `ShowTopology`
    - `ShowEventDetail`
    - `ShowSpanSearch`
    - `ShowTraceEvents`
    - `ShowTrend`
    - `ShowSumTable`
    - `ShowRawTable`
    - `ShowClobDetail`
    - `ListEnvInstances`
    - `ShowEnvMonitorItems`
    - `ListApps`
    - `ListAppEnvs`
    - `ShowAkSks`
    - `CreateAkSk`
    - `DeleteAkSk`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - Support the interface `ListCosts`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CDN

- _Features_
  - Support the interface `ShowUrlTaskInfo`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ipv6_accelerate` to the interface `ShowDomainFullConfig`
  - Add the request parameter `ipv6_accelerate` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CSMS

- _Features_
  - Support the interfaces `UploadSecretBlob`, `DownloadSecretBlob`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the following interfaces：
    - `ListGaussMySqlDatabaseUser`
    - `CreateGaussMySqlDatabaseUser`
    - `DeleteGaussMySqlDatabaseUser`
    - `ResetGaussMySqlDatabasePassword`
    - `AddDatabasePermission`
    - `DeleteDatabasePermission`
    - `ListGaussMySqlDatabaseCharsets`
    - `ListGaussMySqlDatabase`
    - `CreateGaussMySqlDatabase`
    - `DeleteGaussMySqlDatabase`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `enableHyG`, `trafficIpList`, `cryptAlgorithm`, `enableHttps`, `tags` to the interface `ListGraphs`
  - Add the response parameters `enableHyG`, `trafficIpList`, `cryptAlgorithm`, `enableHttps`, `tags` to the interface `ShowGraph`

### HuaweiCloud SDK Kafka

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreatePostPaidInstance`:
    - Add the request parameter `broker_num`
    - Add the enum values `c6.2u4g.cluster`, `c6.4u8g.cluster`, `c6.8u16g.cluster`, `c6.12u24g.cluster`, `c6.16u32g.cluster` to the request parameter `specification`
    - Add the enum values `250`, `500`, `1000`, `1500`, `2000` to the request parameter `partition_num`
    - Add the enum values `dms.physical.storage.high.v2`, `dms.physical.storage.ultra.v2` to the request parameter `storage_spec_code`
    - The request parameter `specification` changed to not required
  - Add the response parameters `description`, `access_user`, `ssl_two_way_enable`, `cert_replaced`, `public_boundwidth`, `agent_enable`, `public_access_enabled`, `node_num`, `new_spec_billing_enable`, `broker_num` to the interface `ListInstances`
  - Add the response parameters `description`, `access_user`, `ssl_two_way_enable`, `cert_replaced`, `public_boundwidth`, `agent_enable`, `public_access_enabled`, `node_num`, `new_spec_billing_enable`, `broker_num` to the interface `ShowInstance`
  - The request parameter `engine` changed to not required of the interface `ShowInstanceExtendProductInfo`
  - Changes of the interface `ShowPartitionBeginningMessage`:
    - Add the response parameter `offset`
    - Remove the response parameter `message_offset`
  - Changes of the interface `ShowPartitionEndMessage`:
    - Add the response parameter `offset`
    - Remove the response parameter `message_offset`
  - Add the response parameter `product_alias` to the interface `ListEngineProducts`

### HuaweiCloud SDK Moderation

- _Features_
  - Support the interfaces `RunCreateAudioModerationJob`, `RunQueryAudioModerationJob`
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `ad_glossaries` from the interface `RunImageModeration`

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeMacaoIdCard`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `count` from the interface `CreateRestoreInstance`

### HuaweiCloud SDK SWR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `filter` to the interface `ListNamespaces`
  - Add the request parameters `limit`, `offset`, `order_column`, `order_type` to the interface `ListReposDetails`
  - Add the request parameter `filter` to the interface `ListRepositoryTags`
  - Add the request parameters `namespace`, `name`, `center`, `limit`, `offset`, `order_column`, `order_type` to the interface `ListSharedReposDetails`
  - Changes of the interface `ListRetentionHistories`:
    - Add the request parameter `filter`
    - Remove the request parameters `offset`, `limit`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `DeleteIgnoreRule`:
    - Add the response parameter `advanced`
    - Remove the response parameter `rule`
  - Changes of the interface `CreateIgnoreRule`:
    - Add the request parameter `advanced`
    - Add the response parameter `advanced`
  - Add the response parameter `advanced` to the interface `ListIgnoreRule`
  - Changes of the interface `ListStatistics`:
    - Remove the response parameter `host`
    - Modify the type `string` -> `array` of the request parameter `instances`
    - Modify the type `string` -> `array` of the request parameter `hosts`
  - Changes of the interface `ListQpsTimeline`:
    - Modify the type `string` -> `array` of the request parameter `instances`
    - Modify the type `string` -> `array` of the request parameter `hosts`
  - Modify the type `string` -> `array` of the request parameter `instances` of the interface `ListBandwidthTimeline`
  - Changes of the interface `ListTopAbnormal`:
    - Modify the type `string` -> `array` of the request parameter `instances`
    - Modify the type `string` -> `array` of the request parameter `hosts`
  - Changes of the interface `ListOverviewsClassification`:
    - Modify the type `string` -> `array` of the request parameter `instances`
    - Modify the type `string` -> `array` of the request parameter `hosts`
  - Add the response parameters `geoip_enable`, `load_balance_enable`, `ipv6_protection_enable`, `policy_sharing_enable`, `ip_group`, `robot_action_enable`, `http2_enable`, `timeout_config_enable` to the interface `ShowConsoleConfig`
  - Add the response parameter `producer` to the interface `CreateValueList`
  - Modify the type `string` -> `enum` of the response parameter `type` of the interface `ListValueList`
  - Remove the response parameter `timestamp` from the interface `UpdateValueList`
  - Add the response parameter `payload_location` to the interface `ListEvent`
  - Changes of the interface `CreateHost`:
    - Add the request parameters `web_tag`, `exclusive_ip`, `paid_type`, `lb_algorithm`, `weight`
    - Add the response parameters `lb_algorithm`, `web_tag`, `block_page`, `extend`, `weight`, `ipv6`
  - Changes of the interface `ListHost`:
    - Add the response parameters `region`, `web_tag`, `ipv6`
    - Remove the response parameter `timeout_config`
  - Modify the type `string` -> `enum` of the response parameter `back_protocol` of the interface `ListHostRoute`
  - Add the response parameters `web_tag`, `ipv6` to the interface `DeleteHost`
  - Changes of the interface `UpdateHost`:
    - Add the request parameters `web_tag`, `exclusive_ip`, `paid_type`, `circuit_breaker`
    - Add the response parameters `projectid`, `extend`, `traffic_mark`, `circuit_breaker`, `access_progress`, `weight`, `ipv6`
    - Remove the request parameter `lb_algorithm`
    - Remove the response parameter `ipv6_enable`
    - Modify the type `enum` -> `string` of the response parameter `protocol`
    - Modify the type `boolean` -> `string` of the response parameter `web_tag`
    - Modify the type `string` -> `enum` of the response parameter `lb_algorithm`
  - Add the response parameters `domainid`, `projectid`, `enterprise_project_id`, `locked`, `tls`, `cipher`, `block_page`, `extend`, `traffic_mark`, `circuit_breaker`, `lb_algorithm`, `web_tag`, `flag`, `description`, `http2_enable`, `access_progress`, `weight` to the interface `ShowHost`
  - Add the response parameters `robot_action`, `modulex_options` to the interface `CreatePolicy`
  - Add the response parameters `robot_action`, `modulex_options`, `hosts` to the interface `ListPolicy`
  - Add the response parameters `robot_action`, `modulex_options` to the interface `DeletePolicy`
  - Add the response parameters `robot_action`, `modulex_options`, `hosts` to the interface `UpdatePolicyProtectHost`
  - Changes of the interface `UpdatePolicy`:
    - Add the request parameters `level`, `full_detection`, `robot_action`, `modulex_options`, `hosts`, `bind_host`, `extend`
    - Add the response parameters `robot_action`, `modulex_options`
  - Add the response parameters `robot_action`, `modulex_options`, `hosts` to the interface `ShowPolicy`
  - Add the enum values `custom`, `ignore` to the request parameter `ruletype` to the interface `UpdatePolicyRuleStatus`
  - Changes of the interface `CreateWhiteblackipRule`:
    - Add the request parameter `ip_group_id`
    - Add the response parameters `name`, `ip_group`, `status`, `description`
    - The request parameter `addr` changed to not required
  - Add the response parameters `name`, `ip_group` to the interface `ListWhiteblackipRule`
  - Add the response parameter `ip_group` to the interface `DeleteWhiteBlackIpRule`
  - Changes of the interface `UpdateWhiteblackipRule`:
    - Add the request parameter `ip_group_id`
    - Add the response parameters `name`, `ip_group`
    - The request parameter `addr` changed to not required
  - Add the response parameters `timestamp`, `status`, `description` to the interface `CreatePrivacyRule`
  - Add the response parameter `description` to the interface `ListPrivacyRule`
  - Add the response parameters `timestamp`, `status`, `description` to the interface `UpdatePrivacyRule`
  - Changes of the interface `CreateGeoipRule`:
    - Add the request parameters `name`, `status`
    - Add the response parameters `name`, `status`
  - Add the response parameters `name`, `status` to the interface `ListGeoipRule`
  - Add the response parameters `name`, `status` to the interface `DeleteGeoipRule`
  - Changes of the interface `UpdateGeoipRule`:
    - Add the request parameter `name`
    - Add the response parameter `name`
  - Remove the response parameters `content`, `key` from the interface `ListCertificates`
  - Changes of the interface `ListCompositeHosts`:
    - Add the response parameters `hostid`, `web_tag`, `access_progress`, `premium_waf_instances`, `description`, `exclusive_ip`, `region`
    - Remove the response parameters `pci_dss`, `pci_3ds`, `cname`, `is_dual_az`, `ipv6`
  - Changes of the interface `ShowCompositeHost`:
    - Add the response parameters `hostid`, `web_tag`, `access_progress`, `premium_waf_instances`, `description`, `exclusive_ip`, `region`
    - Remove the response parameters `pci_dss`, `pci_3ds`, `cname`, `is_dual_az`, `ipv6`
  - Changes of the interface `CreatePremiumHost`:
    - Add the request parameters `block_page`, `description`, `weight`
    - Add the response parameters `server`, `proxy`, `locked`, `timestamp`, `tls`, `cipher`, `extend`, `flag`, `description`, `enterprise_project_id`, `protect_status`, `access_status`, `block_page`
    - Modify the type `string` -> `enum` of the response parameter `protocol`
  - Changes of the interface `ListPremiumHost`:
    - Add the response parameters `extend`, `region`, `description`, `web_tag`, `hostid`
    - Remove the response parameters `mode`, `pool_ids`
  - Changes of the interface `DeletePremiumHost`:
    - Add the response parameters `extend`, `description`, `web_tag`, `host_id`
    - Remove the response parameters `mode`, `pool_ids`
  - Changes of the interface `UpdatePremiumHost`:
    - Add the response parameters `description`, `projectid`, `enterprise_project_id`, `web_tag`, `lb_algorithm`, `access_progress`, `weight`
    - Remove the request parameters `flag`, `extend`
    - Remove the response parameters `mode`, `pool_ids`, `project_id`, `access_code`
  - Changes of the interface `ShowPremiumHost`:
    - Add the response parameters `description`, `projectid`, `enterprise_project_id`, `web_tag`, `access_progress`, `weight`
    - Remove the response parameters `mode`, `pool_ids`, `project_id`, `access_code`
  - Changes of the interface `UpdateCertificate`:
    - Add the request parameters `content`, `key`
    - The request parameter `name` changed to required

# 0.0.102 2022-08-08

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of incorrect splicing of query parameters of enumeration types
- _Change_
  - None

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interfaces:
    - `UpdateSqlFilterControl`
    - `ShowSqlFilterControl`
    - `SetSqlFilterRule`
    - `ShowSqlFilterRule`
    - `DeleteSqlFilterRule`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.101 2022-08-02

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `flavors` to the interface `ShowGaussMySqlFlavors`
  - Add the response parameter `instance` to the interface `ShowGaussMySqlInstanceInfo`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - Support the interface `UpgradeDbVersion`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `patch_available` to the interface `ListInstances`

### HuaweiCloud SDK Image

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `language` from the interface `RunImageDescription`

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `CreateRecordIndex`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interfaces `CreateProjectDomain`, `ListProjectDomains`, `UpdateProjectDomain`, `CancelProjectDomain`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SIS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `auto` to the request parameter `audio_format` to the interface `RecognizeShortAudio`

# 0.0.100 2022-07-28

### HuaweiCloud SDK CBS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interfaces `CreateTbSession`, `ExecuteTbSession`, `DeleteTbSession`
  - Modify the type `string` -> `int32` of the request parameter `top` of the interface `CollectHotQuestions`

### HuaweiCloud SDK DRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `db_type`, `db_type`, `key`, `value`, `key`, `value` changed to not required of the interface `BatchCreateJobs`
  - The request parameter `id`, `object_type`, `object_name` changed to not required of the interface `BatchSetObjects`
  - The request parameter `name`, `alarm_to_user`, `db_type`, `db_type`, `node_type`, `engine_type`, `net_type`, `store_db_info`, `key`, `value` changed to not required of the interface `BatchUpdateJob`
  - The response parameter `db_type`, `db_type`, `db_type`, `db_type` changed to not required of the interface `BatchListJobDetails`
  - The request parameter `id`, `select` changed to not required of the interface `BatchChangeData`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interfaces `ShowDedicatedResourceInfo`, `SetGaussMySqlProxyWeight`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `proxy`, `master_node`, `readonly_nodes` to the interface `ShowGaussMySqlProxy`
  - Add the response parameter `proxy_list` to the interface `ShowGaussMySqlProxyList`
  - Add the response parameter `proxy_flavor_groups` to the interface `ShowGaussMySqlProxyFlavors`
  - Changes of the interface `ShowGaussMySqlBackupList`:
    - Add the enum values `BUILDING`, `COMPLETED`, `FAILED`, `AVAILABLE` to the response parameter `status`
    - Add the enum values `auto`, `manual` to the response parameter `type`
    - Add the enum values `0`, `1`, `2` to the response parameter `backup_level`

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateInstance`:
    - Add the request parameters `period_type`, `period_num`, `is_auto_renew`, `is_auto_pay`
    - Add the enum values `prePaid` to the request parameter `charge_mode`
  - Add the request parameter `is_auto_pay` to the interface `RunInstanceAction`
  - Changes of the interface `CreateRestoreInstance`:
    - Add the request parameters `period_type`, `period_num`, `is_auto_renew`, `is_auto_pay`
    - Add the enum values `prePaid` to the request parameter `charge_mode`
  - Modify the type `string` -> `boolean` of the request parameter `is_auto_pay` of the interface `ResizeInstanceFlavor`

### HuaweiCloud SDK GSL

- _Features_
  - Support the interface `ShowMonthUsages`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK KMS

- _Features_
  - Support the following interfaces：
    - `ListKeyStores`
    - `CreateKeyStore`
    - `ShowKeyStore`
    - `DeleteKeyStore`
    - `EnableKeyStore`
    - `DisableKeyStore`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `keystore_id` to the interface `CreateKey`
  - Add the response parameters `keystore_id`, `key_label` to the interface `ListKeys`
  - Add the response parameters `keystore_id`, `key_label` to the interface `ListKeyDetail`
  - Add the response parameters `keystore_id`, `key_label` to the interface `ListKmsByTags`

### HuaweiCloud SDK NLP

- _Features_
  - Support the interface `RunConstituencyParser`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.99 2022-07-21

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `string` -> `int32` of the response parameter `status` of the interface `ListLatelyApiStatisticsV2`

### HuaweiCloud SDK CES

- _Features_
  - Support the following interfaces：
    - `ListAlarmRules`
    - `CreateAlarmRules`
    - `BatchDeleteAlarmRules`
    - `BatchEnableAlarmRules`
    - `ListAlarmRuleResources`
    - `DeleteAlarmRuleResources`
    - `AddAlarmRuleResources`
    - `ListAlarmRulePolicies`
    - `UpdateAlarmRulePolicies`
    - `ListAgentDimensionInfo`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListAlarmHistories`:
    - Add the response parameter `datapoints`
    - Remove the response parameters `data_points`, `type`, `notification_list`, `type`, `notification_list`
    - Modify the type `string` -> `enum` of the response parameter `status`
    - Modify the type `int32` -> `enum` of the response parameter `level`
    - Modify the type `string` -> `enum` of the response parameter `type`
    - Modify the type `integer` -> `enum` of the response parameter `period`
    - Modify the type `float` -> `double` of the response parameter `value`
    - Modify the type `integer` -> `enum` of the response parameter `suppress_duration`

### HuaweiCloud SDK CloudIDE

- _Features_
  - Support the interface `ShowInstanceStatusInfo`
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameters `instance_user_domain_name`, `instance_user_name` from the interface `CreateInstance`
  - Remove the request parameters `instance_user_domain_name`, `instance_user_name` from the interface `CreateInstanceBy3rd`

### HuaweiCloud SDK HSS

- _Features_
  - Support the interfaces:
    - `ListHostStatus`
    - `ListPasswordComplexity`
    - `ListRiskConfigCheckRules`
    - `ListRiskConfigHosts`
    - `ListRiskConfigs`
    - `ListSecurityEvents`
    - `ListVulnerabilities`
    - `ListWeakPasswordUsers`
    - `ShowCheckRuleDetail`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Image

- _Features_
  - Support the interface `RunImageDescription`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `security_group_rules_links` to the interface `NeutronListSecurityGroupRules`

# 0.0.98 2022-07-14

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `effective_tag_pairs`, `cost_unit_pairs` to the interface `ListCustomerBillsMonthlyBreakDown`

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `tls_version` to the interface `UpdateDomainFullConfig`
  - Add the response parameter `tls_version` to the interface `ShowDomainFullConfig`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `tags`, `cpu_type` to the interface `ShowInstance`

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `date-time` -> `string` of the response parameter `create_time` of the interface `ListPublicips`
  - Modify the type `date-time` -> `string` of the response parameter `create_time` of the interface `ShowPublicip`

### HuaweiCloud SDK Image

- _Features_
  - Support the interfaces `RunImageMediaTagging`, `RunImageMainObjectDetection`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `port_filter`, `ovs_hybrid_plug` to the interface `ListPorts`
  - Add the response parameters `port_filter`, `ovs_hybrid_plug` to the interface `UpdatePort`
  - Add the response parameters `port_filter`, `ovs_hybrid_plug` to the interface `ShowPort`
  - Add the response parameter `remote_address_group_id` to the interface `CreateSecurityGroup`
  - Add the response parameter `remote_address_group_id` to the interface `ListSecurityGroups`
  - Add the response parameter `remote_address_group_id` to the interface `ShowSecurityGroup`
  - Add the response parameter `remote_address_group_id` to the interface `ListSecurityGroupRules`
  - Add the response parameter `remote_address_group_id` to the interface `ShowSecurityGroupRule`
  - Add the response parameter `remote_address_group_id` to the interface `NeutronListSecurityGroups`
  - Add the response parameter `remote_address_group_id` to the interface `NeutronUpdateSecurityGroup`
  - Add the response parameter `remote_address_group_id` to the interface `NeutronShowSecurityGroup`
  - Add the response parameter `remote_address_group_id` to the interface `NeutronListSecurityGroupRules`
  - Add the response parameter `remote_address_group_id` to the interface `NeutronShowSecurityGroupRule`

# 0.0.97 2022-07-07

### HuaweiCloud SDK APM

- _Features_
  - Support the interfaces `SearchApplication`, `SaveMonitorItemConfig`, `ListEnvMonitorItem`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - Support the interfaces `UpdateClusterEip`, `ShowClusterEndpoints`, `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListNodes`:
    - Add the response parameters `isStatic`, `privateIPv6IP`
    - The response parameter `key`, `effect` changed to required
  - Changes of the interface `CreateNode`:
    - Add the request parameter `isStatic`
    - The request parameter `key`, `effect` changed to required
  - Changes of the interface `ShowNode`:
    - Add the response parameters `isStatic`, `privateIPv6IP`
    - The response parameter `key`, `effect` changed to required
  - Changes of the interface `DeleteNode`:
    - Add the response parameters `isStatic`, `privateIPv6IP`
    - The response parameter `key`, `effect` changed to required
  - Changes of the interface `UpdateNode`:
    - Add the response parameters `isStatic`, `privateIPv6IP`
    - The response parameter `key`, `effect` changed to required
  - The request parameter `key`, `effect` changed to required of the interface `AddNode`
  - The request parameter `key`, `effect` changed to required of the interface `ResetNode`
  - Changes of the interface `ListNodePools`:
    - Add the response parameter `isStatic`
    - The response parameter `key`, `effect` changed to required
  - Changes of the interface `CreateNodePool`:
    - Add the request parameter `isStatic`
    - The request parameter `key`, `effect` changed to required
  - Changes of the interface `ShowNodePool`:
    - Add the response parameter `isStatic`
    - The response parameter `key`, `effect` changed to required
  - Changes of the interface `DeleteNodePool`:
    - Add the response parameter `isStatic`
    - The response parameter `key`, `effect` changed to required
  - Changes of the interface `UpdateNodePool`:
    - Add the response parameter `isStatic`
    - The request parameter `key`, `effect` changed to required
    - The response parameter `key`, `effect` changed to required

### HuaweiCloud SDK CloudRTC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `filling_policy` to the interface `CreateMixJob`
  - Changes of the interface `UpdateMixJob`:
    - Add the request parameter `filling_policy`
    - Add the response parameter `filling_policy`
  - Add the response parameter `filling_policy` to the interface `ShowMixJob`

### HuaweiCloud SDK CSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `CREATING` from the interface `ListClustersDetails`
  - Remove the response parameter `CREATING` from the interface `ShowClusterDetail`
  - Add the request parameter `unBindPublicReq` to the interface `UpdateUnbindPublic`
  - Changes of the interface `ListYmlsJob`:
    - Add the response parameter `configurations`
    - Remove the response parameter `configList`
  - Changes of the interface `ListYmls`:
    - Add the response parameter `configurations`
    - Remove the response parameter `configList`

### HuaweiCloud SDK CTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateNotification`:
    - Add the request parameter `filter`
    - Add the response parameter `filter`
  - Add the request parameter `filter` to the interface `CreateNotification`
  - Add the response parameter `filter` to the interface `ListNotifications`

### HuaweiCloud SDK ELB

- _Features_
  - Support the interfaces `ListMasterSlavePools`, `CreateMasterSlavePool`, `ShowMasterSlavePool`, `DeleteMasterSlavePool`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSystemSecurityPolicies`:
    - Modify the type `array` -> `string` of the response parameter `protocols`
    - Modify the type `array` -> `string` of the response parameter `ciphers`
  - Add the request parameter `X-Auth-Token` to the interface `ListQuotaDetails`
  - Add the request parameter `public_border_group` to the interface `ListAvailabilityZones`
  - Changes of the interface `CreateLoadBalancer`:
    - Add the request parameters `id`, `global_eip_ids`
    - Remove the request parameter `min_l4_flavor_id`
    - The request parameter `X-Auth-Token` changed to required
  - Changes of the interface `ListLoadBalancers`:
    - Add the response parameters `global_eips`, `public_border_group`
    - Remove the response parameter `min_l4_flavor_id`
    - The request parameter `X-Auth-Token` changed to required
  - Changes of the interface `UpdateLoadBalancer`:
    - Add the request parameter `cloud_service_console_url`
    - Add the response parameters `global_eips`, `public_border_group`
    - Remove the request parameter `min_l4_flavor_id`
    - Remove the response parameter `min_l4_flavor_id`
  - Changes of the interface `ShowLoadBalancer`:
    - Add the response parameters `global_eips`, `public_border_group`
    - Remove the response parameter `min_l4_flavor_id`
  - Add the request parameter `X-Auth-Token` to the interface `ChangeLoadbalancerChargeMode`
  - Remove the request parameters `enc_certificate`, `enc_private_key` from the interface `CreateCertificate`
  - Remove the response parameters `enc_certificate`, `enc_private_key` from the interface `ListCertificates`
  - Changes of the interface `UpdateCertificate`:
    - Remove the request parameters `enc_certificate`, `enc_private_key`
    - Remove the response parameters `enc_certificate`, `enc_private_key`
  - Remove the response parameters `enc_certificate`, `enc_private_key` from the interface `ShowCertificate`
  - Add the request parameter `quic_config` to the interface `CreateListener`
  - Add the response parameter `quic_config` to the interface `ListListeners`
  - Changes of the interface `UpdateListener`:
    - Add the request parameter `quic_config`
    - Add the response parameter `quic_config`
  - Add the response parameter `quic_config` to the interface `ShowListener`
  - Add the request parameters `vpc_id`, `type` to the interface `CreatePool`
  - Changes of the interface `ListPools`:
    - Add the request parameters `vpc_id`, `type`
    - Add the response parameters `created_at`, `updated_at`, `vpc_id`, `type`
  - Changes of the interface `UpdatePool`:
    - Add the request parameters `X-Auth-Token`, `vpc_id`, `type`
    - Add the response parameters `created_at`, `updated_at`, `vpc_id`, `type`
  - Add the response parameters `created_at`, `updated_at`, `vpc_id`, `type` to the interface `ShowPool`
  - Modify the type `enum` -> `string` of the request parameter `project_id` of the interface `CreateMember`
  - Changes of the interface `ListMembers`:
    - Add the response parameters `status`, `loadbalancers`, `created_at`, `updated_at`
    - Remove the request parameter `instance_id`
  - Add the response parameters `status`, `loadbalancers`, `created_at`, `updated_at` to the interface `UpdateMember`
  - Add the response parameters `status`, `loadbalancers`, `created_at`, `updated_at` to the interface `ShowMember`
  - Add the response parameters `status`, `loadbalancers`, `created_at`, `updated_at` to the interface `ListAllMembers`
  - Add the response parameters `created_at`, `updated_at` to the interface `ListHealthMonitors`
  - Add the response parameters `created_at`, `updated_at` to the interface `UpdateHealthMonitor`
  - Add the response parameters `created_at`, `updated_at` to the interface `ShowHealthMonitor`
  - Add the request parameter `redirect_pools_config` to the interface `CreateL7Policy`
  - Add the response parameters `redirect_pools_config`, `created_at`, `updated_at` to the interface `ListL7Policies`
  - Changes of the interface `UpdateL7Policy`:
    - Add the request parameter `redirect_pools_config`
    - Add the response parameters `redirect_pools_config`, `created_at`, `updated_at`
  - Add the response parameters `redirect_pools_config`, `created_at`, `updated_at` to the interface `ShowL7Policy`
  - Add the request parameter `X-Auth-Token` to the interface `BatchUpdatePoliciesPriority`
  - Add the response parameters `created_at`, `updated_at` to the interface `ListL7Rules`
  - Add the response parameters `created_at`, `updated_at` to the interface `UpdateL7Rule`
  - Add the response parameters `created_at`, `updated_at` to the interface `ShowL7Rule`
  - Changes of the interface `UpdateIpList`:
    - Remove the request parameters `name`, `ip_list`, `description`
    - The request parameter `X-Auth-Token` changed to required
  - Changes of the interface `BatchDeleteIpList`:
    - Add the request parameter `BatchDeleteIpListRequestBody`
    - Remove the request parameter `BatchDeleteIpGroupIpListRequestBody`
    - The request parameter `X-Auth-Token` changed to required
  - Changes of the interface `BatchCreateMembers`:
    - Add the request parameter `BatchCreateMembersRequestBody`
    - Add the response parameter `status`
    - Remove the request parameter `BatchCreateMemberRequestBody`
  - Changes of the interface `BatchDeleteMembers`:
    - Add the request parameter `BatchDeleteMembersRequestBody`
    - Remove the request parameter `BatchDeleteMemberRequestBody`
  - Changes of the interface `UpdateLogtank`:
    - Add the request parameter `UpdateLogtankRequestBody`
    - Remove the request parameter `logtank`

### HuaweiCloud SDK IoTEdge

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the following interfaces:
    - `CreateAccessCode`
    - `ShowProtocolMappings`
    - `UploadProtocolMappings`
    - `BatchUpdateConfigs`
    - `ShowExternalEntity`
  - Add the enum values `COMPOSITE_APPLICATION`, `DATA_COLLECTION` to the request parameter `function_type` to the interface `CreateEdgeApp`
  - Add the response parameters `sdk_version`, `deploy_multi_instance` to the interface `BatchListEdgeAppVersions`
  - Add the request parameters `sdk_version`, `deploy_multi_instance`, `ext_devices`, `tcp_socket`, `period_seconds`, `failure_threshold`, `tcp_socket`, `period_seconds`, `failure_threshold` to the interface `CreateEdgeApplicationVersion`
  - Changes of the interface `UpdateEdgeApplicationVersion`:
    - Add the request parameters `deploy_multi_instance`, `sdk_version`, `ext_devices`, `tcp_socket`, `period_seconds`, `failure_threshold`, `tcp_socket`, `period_seconds`, `failure_threshold`
    - Add the response parameters `sdk_version`, `deploy_multi_instance`
    - The request parameter `arch` changed to not required
  - Add the response parameters `deploy_multi_instance`, `sdk_version`, `tcp_socket`, `period_seconds`, `failure_threshold`, `tcp_socket`, `period_seconds`, `failure_threshold`, `ext_devices` to the interface `ShowEdgeApplicationVersion`
  - Add the response parameters `sdk_version`, `deploy_multi_instance` to the interface `UpdateEdgeApplicationVersionState`
  - Add the request parameters `edge_node_id`, `verify_code`, `time_out`, `arch`, `base_path`, `hardware_model` to the interface `CreateEdgeNode`
  - Add the response parameters `ha_config`, `hardware_model` to the interface `ShowEdgeNode`
  - Add the request parameter `CreateInstallCmdRequestBody` to the interface `CreateInstallCmd`
  - Changes of the interface `BatchListModules`:
    - Add the response parameters `control_status`, `container_settings`
    - Add the enum values `DELETE_SUCCESS`, `STOPPED` to the response parameter `state`
    - Add the enum values `GATEWAY_MANAGER`, `COMPOSITE_APPLICATION`, `DATA_COLLECTION` to the response parameter `function_type`
  - Add the request parameters `module_name`, `container_settings` to the interface `CreateModule`
  - Changes of the interface `UpdateModule`:
    - Add the request parameters `module_name`, `container_settings`
    - Add the response parameter `control_status`
    - The request parameter `app_version` changed to not required
    - Add the enum values `DELETE_SUCCESS`, `STOPPED` to the response parameter `state`
    - Add the enum values `GATEWAY_MANAGER`, `COMPOSITE_APPLICATION`, `DATA_COLLECTION` to the response parameter `function_type`
  - Changes of the interface `ShowModule`:
    - Add the response parameters `control_status`, `container_settings`
    - Add the enum values `DELETE_SUCCESS`, `STOPPED` to the response parameter `state`
    - Add the enum values `GATEWAY_MANAGER`, `COMPOSITE_APPLICATION`, `DATA_COLLECTION` to the response parameter `function_type`

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `title` to the interface `RecognizeVatInvoice`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `SetSensitiveSlowLog`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RMS

- _Features_
  - Support the interface `ListSchemas`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SIS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `chinese_16k_travel` to the request parameter `property` to the interface `RecognizeShortAudio`
  - Add the enum values `chinese_16k_media` to the request parameter `property` to the interface `PushTranscriberJobs`
  - Add the response parameter `audio_duration` to the interface `CollectTranscriberJob`
  - Add the enum values `chinese_huaxiaomei_common`, `chinese_huaxiaofei_common` to the request parameter `property` to the interface `RunTts`

# 0.0.96 2022-06-30

### HuaweiCloud SDK Core

- _Features_
  - Support federal authentication
  - Support authentication management
- _Bug Fix_
  - [Issue 53](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/53) Fix the issue of panic when the client sends a request
- _Change_
  - None

### HuaweiCloud SDK UGO

- _Features_
  - Support `Database and Application Migration` service.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `mobile_phone` changed to required of the interface `SendVerificationMessageCode`

### HuaweiCloud SDK BSSINTL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `email` changed to required of the interface `SendVerificationMessageCode`

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `cache_url_parameter_filter` to the interface `ShowDomainFullConfig`
  - Add the request parameter `cache_url_parameter_filter` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CloudIDE

- _Features_
  - Support the interface `UploadExtensionFile`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DNS

- _Features_
  - Support the interfaces `CreateRecordSetWithBatchLines`, `BatchUpdateRecordSetWithLine`, `BatchDeleteRecordSetWithLine`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `records` changed to not required of the interface `CreateRecordSetWithLine`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - Support the interface `UpdateFunctionMaxInstanceConfig`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `lb_port` to the interface `ListInstances`

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `total_count` to the interface `ListComponentInfos`

### HuaweiCloud SDK IEF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateTag`:
    - Add the request parameter `CreateTagRequestBody`
    - Remove the request parameter `tag`
  - Add the request parameters `sort`, `state` to the interface `ListEdgeNodes`

### HuaweiCloud SDK OCR

- _Features_
  - Support the interfaces `RecognizeHkIdCard`, `RecognizeCambodianIdCard`, `RecognizeExitEntryPermit`, `RecognizeMainlandTravelPermit`
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int32` -> `float` of the response parameter `direction` of the interface `RecognizeGeneralText`

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interfaces `CreateProjectModule`, `ListProjectModules`, `UpdateProjectModule`, `DeleteProjectModule`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK UGO

- _Features_
  - Support the following interfaces:
    - `ListQuotas`
    - `RunSqlConversion`
    - `ListEvaluationProjects`
    - `CreateEvaluationProject`
    - `ShowEvaluationProjectStatus`
    - `ShowEvaluationProjectDetail`
    - `ConfirmTargetDbType`
    - `DeleteEvaluationProject`
    - `ListMigrationProjects`
    - `CreateMigrationProject`
    - `ShowMigrationProjectStatus`
    - `CheckPermission`
    - `ListPermissionCheckResult`
    - `ShowMigrationProjectDetail`
    - `CommitSyntaxConversion`
    - `ListSyntaxConversionProgress`
    - `CommitVerification`
    - `ListVerificationProgress`
    - `DownloadFailureReport`
    - `DeleteMigrationProject`
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `MigrateSqlStatement`
  - Changes of the interface `ListApiVersions`:
    - Add the request parameter `X-Auth-Token`
    - The response parameter `id`, `links`, `version`, `status`, `updated` changed to required
  - Changes of the interface `ShowApiVersionInfo`:
    - Add the request parameter `X-Auth-Token`
    - The response parameter `id`, `links`, `version`, `status`, `updated` changed to required

### HuaweiCloud SDK WAF

- _Features_
  - Support the interface `ListOverviewsClassification`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `host` to the interface `ListStatistics`
  - Add the response parameter `timeout_config` to the interface `ListHost`
  - Add the response parameter `timeout_config` to the interface `ShowHost`
  - Changes of the interface `UpdateHost`:
    - Add the request parameters `block_page`, `traffic_mark`, `flag`, `extend`, `http2_enable`, `ipv6_enable`, `lb_algorithm`, `timeout_config`
    - Add the response parameters `http2_enable`, `ipv6_enable`, `lb_algorithm`, `timeout_config`
  - Changes of the interface `UpdatePremiumHost`:
    - Add the request parameters `mode`, `locked`, `protect_status`, `access_status`, `timestamp`, `pool_ids`, `block_page`, `traffic_mark`, `flag`, `extend`, `circuit_breaker`, `timeout_config`
    - Add the response parameter `timeout_config`
  - Add the response parameter `timeout_config` to the interface `ShowPremiumHost`

# 0.0.95 2022-06-25

### HuaweiCloud SDK CodeCheck

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `cyclomatic_complexity_per_file`, `file_duplication_ratio` to the interface `ShowTaskDetail`

### HuaweiCloud SDK DDS

- _Features_
  - Support the following interfaces:
    - `ShowEntityConfiguration`
    - `UpdateEntityConfiguration`
    - `ShowConfigurationParameter`
    - `UpdateConfigurationParameter`
    - `DeleteConfiguration`
    - `ListConfigurations`
    - `CreateConfiguration`
    - `SwitchConfiguration`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `configurations`, `charge_info` to the interface `CreateInstance`
  - Changes of the interface `ResizeInstanceVolume`:
    - Add the request parameter `is_auto_pay`
    - Add the response parameter `order_id`
  - Changes of the interface `AddShardingNode`:
    - Add the request parameter `is_auto_pay`
    - Add the response parameter `order_id`
  - Changes of the interface `ResizeInstance`:
    - Add the request parameter `is_auto_pay`
    - Add the response parameter `order_id`
  - Add the request parameters `configurations`, `charge_info` to the interface `RestoreNewInstance`

# 0.0.94 2022-06-23

### HuaweiCloud SDK CSE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UploadKie`:
    - Add the request parameter `upload_file`
    - Remove the request parameter `UploadFile`

### HuaweiCloud SDK CSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateCluster`:
    - Add the request parameters `httpsEnable`, `authorityEnable`, `adminPwd`
    - The request parameter `availability_zone` changed to required
  - Changes of the interface `ListClustersDetails`:
    - Add the response parameters `publicKibanaResp`, `elbWhiteList`, `publicIp`, `bandwidthSize`, `diskEncrypted`, `backupAvailable`, `enterpriseProjectId`, `ip`
    - Remove the response parameter `enterprise_project_id`
  - Remove the request parameter `X-Auth-Token` from the interface `DeleteCluster`
  - Changes of the interface `ShowClusterDetail`:
    - Add the response parameters `publicKibanaResp`, `elbWhiteList`, `publicIp`, `vpcId`, `subnetId`, `securityGroupId`, `bandwidthSize`, `diskEncrypted`, `backupAvailable`, `enterpriseProjectId`, `period`, `ip`
    - Remove the response parameter `enterprise_project_id`
  - Changes of the interface `ListFlavors`:
    - Add the response parameters `type`, `availableAZ`
  - Changes of the interface `StartVpecp`:
    - Modify the type `string` -> `boolean` of the request parameter `endpointWithDnsName`
  - Changes of the interface `CreateCluster`:
    - Add the request parameter `payInfo`
    - Add the response parameter `cluster`
    - Remove the response parameter `schema`
  - Changes of the interface `RestartCluster`:
    - Add the request parameter `RestartClusterRequestBody`
    - Remove the request parameters `X-Auth-Token`, `RollingRestartReq`

### HuaweiCloud SDK DRS

- _Features_
  - Support the interfaces `ListAvailableZone`, `UpdateTuningParams`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `master_az`, `slave_az` to the interface `BatchCreateJobs`
  - Add the request parameter `slot_name` to the interface `BatchSetPolicy`

### HuaweiCloud SDK ELB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListQuotaDetails`:
    - Add the request parameter `quota_key`
    - Remove the request parameter `type`
  - Changes of the interface `ListListeners`:
    - Add the request parameters `loadbalancer_id`, `connection_limit`, `admin_state_up`, `http2_enable`, `enterprise_project_id`
    - Remove the request parameters `member_timeout`, `client_timeout`, `keepalive_timeout`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `subnet_id` to the interface `ListInstances`
  - Add the request parameter `subnet_id` to the interface `ExpandInstanceNode`

### HuaweiCloud SDK VSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `AuthorizeDomains`:
    - Add the response parameter `usage_notice`
    - Add the enum values `free` to the request parameter `auth_mode`

# 0.0.93 2022-06-19

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that Region could not be configured by environment variables
- _Change_
  - None

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the response parameter's type of the interface `ListRangeQueryAomPromGet` is incorrect
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeIdDocument`
- _Bug Fix_
  - None
- _Change_
  - None
  
### HuaweiCloud SDK RMS

- _Features_
  - Support the interface `RunQuery`, `CreateStoredQuery`, `ListStoredQueries`, `ShowStoredQuery`, `UpdateStoredQuery`, `DeleteStoredQuery`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `domain_id`, `project_id`, `region`, `company_name`, `product_name` changed to required of the interface `CheckProductHealthy`
  - Changes of the interface `ImportEvents`:
    - The request parameter `type`, `domain_id`, `business`, `checkitem_id`, `checkpoint_id`, `spec_id`, `patch_id` changed to required
    - The request parameter `type`, `domain_id`, `product_feature`, `arrive_time`, `verification_state` changed to not required

### HuaweiCloud SDK VSS

- _Features_
  - Support the interfaces `ShowReportStatus`, `DownloadTaskReport`, `ExecuteGenerateReport`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.92 2022-06-09

### HuaweiCloud SDK CSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateCluster`:
    - Add the response parameter `schema`
    - Remove the response parameters `id`, `name`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `UpdateFunctionCode`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `ShowFunctionCode`
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the request parameter `domain_names`
    - Add the response parameter `domain_names`
    - Modify the type `string` -> `enum` of the response parameter `runtime`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `ShowFunctionConfig`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `CreateFunctionVersion`
  - Add the request parameter `option` to the interface `ListStatistics`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `CreateDependency`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `ListDependencies`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `UpdateDependency`
  - Modify the type `string` -> `enum` of the response parameter `runtime` of the interface `ShowDependency`
  - The request parameter `content` changed to required of the interface `UpdateEvent`
  - Add the request parameter `marker` to the interface `ListFunctionAsyncInvocations`
  - The request parameter `workflow_urns` changed to required of the interface `BatchDeleteWorkflows`
  - The request parameter `name`, `trigger_name`, `trigger_type`, `bucket`, `events`, `prefix`, `suffix`, `start`, `name`, `operation`, `id`, `name`, `type`, `end`, `transition`, `ref_name`, `arguments`, `constants`, `name` changed to required of the interface `CreateWorkflow`
  - The request parameter `input` changed to required of the interface `StartWorkflowExecution`
  - Add the response parameters `node_name`, `execution_id`, `request_id` to the interface `ShowWorkflowExecution`
  - The request parameter `name`, `trigger_name`, `trigger_type`, `bucket`, `events`, `prefix`, `suffix`, `start`, `name`, `operation`, `id`, `name`, `type`, `end`, `transition`, `ref_name`, `arguments`, `constants`, `name` changed to required of the interface `UpdateWorkFlow`
  - Changes of the interface `ShowWorkFlow`:
    - Add the response parameters `lts_group_id`, `lts_stream_id`
    - The response parameter `name`, `trigger_name`, `trigger_type`, `bucket`, `events`, `prefix`, `suffix`, `start`, `name`, `operation`, `id`, `name`, `type`, `end`, `transition`, `ref_name`, `arguments`, `constants`, `name` changed to required
  - The request parameter `input` changed to required of the interface `StartSyncWorkflowExecution`

### HuaweiCloud SDK GSL

- _Features_
  - Support the following interfaces:
    - `BatchSetTags`
    - `ListTags`
    - `CreateTag`
    - `DeleteTag`
    - `BatchSetAttributes`
    - `ListAttributes`
    - `CreateAttribute`
    - `UpdateAttribute`
    - `EnableAttribute`
    - `DisableAttribute`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `tag_id` to the interface `ListSimCards`

### HuaweiCloud SDK Moderation

- _Features_
  - Support the interface `RunTextModeration`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `availabilityZoneId` to the interface `ListClusters`
  - Add the response parameter `availabilityZoneId` to the interface `ShowClusterDetails`
  - Add the response parameters `availability_zone_id`, `tags` to the interface `ListHosts`

# 0.0.91 2022-06-02

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `instance_id` to the interface `ListFlavors`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interface `ChangeGaussMySqlProxySpecification`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Meeting

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `maxAudienceParties`, `expireDate`, `commercialMaxAudienceParties` to the interface `SearchCorpVmr`

### HuaweiCloud SDK NAT

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListNatGateways`:
    - Modify the type `date-time` -> `string` of the request parameter `created_at`
    - Modify the type `date-time` -> `string` of the response parameter `created_at`
  - Modify the type `date-time` -> `string` of the response parameter `created_at` of the interface `UpdateNatGateway`
  - Modify the type `date-time` -> `string` of the response parameter `created_at` of the interface `ShowNatGateway`
  - Changes of the interface `ListNatGatewayDnatRules`:
    - Modify the type `date-time` -> `string` of the request parameter `created_at`
    - Modify the type `date-time` -> `string` of the response parameter `created_at`
  - Modify the type `date-time` -> `string` of the response parameter `created_at` of the interface `UpdateNatGatewayDnatRule`
  - Modify the type `date-time` -> `string` of the response parameter `created_at` of the interface `ShowNatGatewayDnatRule`
  - Changes of the interface `ListNatGatewaySnatRules`:
    - Modify the type `date-time` -> `string` of the request parameter `created_at`
    - Modify the type `date-time` -> `string` of the response parameter `created_at`
  - Modify the type `date-time` -> `string` of the response parameter `created_at` of the interface `UpdateNatGatewaySnatRule`
  - Modify the type `date-time` -> `string` of the response parameter `created_at` of the interface `ShowNatGatewaySnatRule`

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `words_block_count`, `words_block_list` to the interface `RecognizeHealthCode`
  - Modify the type `float` -> `object` of the response parameter `confidence` of the interface `RecognizePcrTestRecord`

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the following interfaces:
    - `ListIssueCustomFields`
    - `ListIssuesSfV4`
    - `ListProjectIssuesRecordsV4`
    - `ListWorkitemStatusRecordsV4`
    - `ListWorkitems`
    - `ShowIssuesWrokFlowConfig`
    - `ShowWorkItemWrokflowConfig`
    - `ListAssociatedIssues`
    - `ListAssociatedWikis`
    - `ListIssueAssociatedCommits`
    - `ListAssociatedTestCases`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateIssueV4`:
    - Add the request parameter `new_custom_fields`
    - Add the response parameters `new_custom_fields`, `new_name`
  - Changes of the interface `ListIssuesV4`:
    - Add the request parameter `custom_fields`
    - Add the response parameters `new_custom_fields`, `new_name`
  - Add the response parameters `new_custom_fields`, `new_name` to the interface `ShowIssueV4`
  - Changes of the interface `UpdateIssueV4`:
    - Add the request parameter `new_custom_fields`
    - Add the response parameters `new_custom_fields`, `new_name`
  - Add the response parameters `new_custom_fields`, `new_name` to the interface `ListChildIssuesV4`
  - Changes of the interface `CreateSystemIssueV4`:
    - Add the request parameter `new_custom_fields`
    - Add the response parameters `new_custom_fields`, `new_name`

# 0.0.90 2022-05-26

### HuaweiCloud SDK BCS

- _Features_
  - Support the interfaces `HandleUnionMemberQuitList`, `BatchRemovePeersFromChannel`, `DeleteChannel`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `is_delete_ief`, `is_delete_lightpeer`, `ief_nodes_id` to the interface `DeleteBlockchain`
  - Changes of the interface `CreateNewBlockchain`:
    - Add the request parameter `cluster_platform_type`
    - Remove the request parameters `user_name`, `password`
    - Modify the type `int64` -> `string` of the request parameter `node_flavor`
    - Modify the type `int64` -> `string` of the request parameter `cce_flavor`
    - Modify the type `int64` -> `string` of the request parameter `init_node_pwd`
    - Modify the type `int64` -> `string` of the request parameter `az`
    - Modify the type `int64` -> `string` of the request parameter `cluster_platform_type`
  - Add the request parameter `CreateBlockchainCertByUserNameRequestBody` to the interface `CreateBlockchainCertByUserName`
  - Add the request parameter `file` to the interface `UnfreezeCert`
  - Add the request parameter `file` to the interface `FreezeCert`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `features`, `sub_status` to the interface `ListInstances`
  - Add the response parameters `features`, `transparent_client_ip_enable`, `sub_status` to the interface `ShowInstance`
  - Add the request parameter `execute_immediately` to the interface `ResizeInstance`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `volume_type`, `hw:passthrough` to the interface `AttachServerVolume`

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the interfaces `ListComponentInfos`, `SwitchShard`, `ResizeInstanceFlavor`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `rate_limit`, `prefetch_block`, `filesplit_size` from the interface `ShowBackupPolicy`
  - Modify the type `object` -> `string` of the response parameter `memberof` of the interface `ListDbUsers`

### HuaweiCloud SDK KMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ValidateSignature`:
    - Add the response parameter `signature_valid`
    - Remove the response parameter `signature_vaild`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateAomMappingRules`:
    - Add the request parameter `deployments_prefix`
    - Add the response parameter `deployments_prefix`
  - Changes of the interface `UpdateAomMappingRules`:
    - Add the request parameter `deployments_prefix`
    - Add the response parameter `deployments_prefix`
  - Add the response parameter `deployments_prefix` to the interface `ShowAomMappingRules`
  - Add the response parameter `deployments_prefix` to the interface `ShowAomMappingRule`

### HuaweiCloud SDK Moderation

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `ocr_text`, `error_code`, `error_msg` to the interface `RunCheckResult`
  - Changes of the interface `RunImageBatchModeration`:
    - Add the request parameters `moderation_rule`, `ad_categories`, `show_ocr_text`
    - Add the response parameters `ocr_text`, `error_code`, `error_msg`
  - Add the request parameters `moderation_rule`, `ad_categories`, `show_ocr_text` to the interface `RunTaskSumbit`

# 0.0.89 2022-05-19

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `group_id` to the interface `ListEnvironmentVariablesV2`
  - Add the response parameters `update_time`, `create_time`, `id`, `name`, `sign_type`, `sign_key`, `sign_secret`, `sign_algorithm` to the interface `UpdateSignatureKeyV2`
  - Add the request parameter `loadbalancer_provider` to the interface `CreateInstanceV2`
  - Add the response parameter `loadbalancer_provider` to the interface `ListInstancesV2`
  - Add the response parameters `endpoint_service`, `endpoint_services`, `node_ips`, `publicips`, `privateips`, `loadbalancer_provider` to the interface `ShowDetailsOfInstanceV2`
  - Add the response parameters `endpoint_service`, `endpoint_services`, `node_ips`, `publicips`, `privateips`, `loadbalancer_provider` to the interface `UpdateInstanceV2`

### HuaweiCloud SDK CSE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `X-Enterprise-Project-ID` from the interface `ListEngines`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `dry_run` to the interface `ResizeServer`
  - Add the request parameter `dry_run` to the interface `ResizePostPaidServer`
  - Add the request parameter `dry_run` to the interface `AttachServerVolume`

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListClusters`:
    - Add the response parameters `GroupName`, `NodeNum`, `NodeSize`, `NodeSpecId`, `VmProductId`, `VmSpecCode`, `NodeProductId`, `RootVolumeSize`, `RootVolumeProductId`, `RootVolumeType`, `RootVolumeResourceSpecCode`, `RootVolumeResourceType`, `DataVolumeType`, `DataVolumeCount`, `DataVolumeSize`, `DataVolumeProductId`, `DataVolumeResourceSpecCode`, `DataVolumeResourceType`, `GroupName`, `NodeNum`, `NodeSize`, `NodeSpecId`, `VmProductId`, `VmSpecCode`, `NodeProductId`, `RootVolumeSize`, `RootVolumeProductId`, `RootVolumeType`, `RootVolumeResourceSpecCode`, `RootVolumeResourceType`, `DataVolumeType`, `DataVolumeCount`, `DataVolumeSize`, `DataVolumeProductId`, `DataVolumeResourceSpecCode`, `DataVolumeResourceType`
    - Remove the response parameters `groupName`, `nodeNum`, `nodeSize`, `nodeSpecId`, `vmProductId`, `vmSpecCode`, `nodeProductId`, `rootVolumeSize`, `rootVolumeProductId`, `rootVolumeType`, `rootVolumeResourceSpecCode`, `rootVolumeResourceType`, `dataVolumeType`, `dataVolumeCount`, `dataVolumeSize`, `dataVolumeProductId`, `dataVolumeResourceSpecCode`, `dataVolumeResourceType`, `groupName`, `nodeNum`, `nodeSize`, `nodeSpecId`, `vmProductId`, `vmSpecCode`, `nodeProductId`, `rootVolumeSize`, `rootVolumeProductId`, `rootVolumeType`, `rootVolumeResourceSpecCode`, `rootVolumeResourceType`, `dataVolumeType`, `dataVolumeCount`, `dataVolumeSize`, `dataVolumeProductId`, `dataVolumeResourceSpecCode`, `dataVolumeResourceType`
  - Changes of the interface `ShowClusterDetails`:
    - Add the response parameters `GroupName`, `NodeNum`, `NodeSize`, `NodeSpecId`, `VmProductId`, `VmSpecCode`, `NodeProductId`, `RootVolumeSize`, `RootVolumeProductId`, `RootVolumeType`, `RootVolumeResourceSpecCode`, `RootVolumeResourceType`, `DataVolumeType`, `DataVolumeCount`, `DataVolumeSize`, `DataVolumeProductId`, `DataVolumeResourceSpecCode`, `DataVolumeResourceType`, `GroupName`, `NodeNum`, `NodeSize`, `NodeSpecId`, `VmProductId`, `VmSpecCode`, `NodeProductId`, `RootVolumeSize`, `RootVolumeProductId`, `RootVolumeType`, `RootVolumeResourceSpecCode`, `RootVolumeResourceType`, `DataVolumeType`, `DataVolumeCount`, `DataVolumeSize`, `DataVolumeProductId`, `DataVolumeResourceSpecCode`, `DataVolumeResourceType`
    - Remove the response parameters `groupName`, `nodeNum`, `nodeSize`, `nodeSpecId`, `vmProductId`, `vmSpecCode`, `nodeProductId`, `rootVolumeSize`, `rootVolumeProductId`, `rootVolumeType`, `rootVolumeResourceSpecCode`, `rootVolumeResourceType`, `dataVolumeType`, `dataVolumeCount`, `dataVolumeSize`, `dataVolumeProductId`, `dataVolumeResourceSpecCode`, `dataVolumeResourceType`, `groupName`, `nodeNum`, `nodeSize`, `nodeSpecId`, `vmProductId`, `vmSpecCode`, `nodeProductId`, `rootVolumeSize`, `rootVolumeProductId`, `rootVolumeType`, `rootVolumeResourceSpecCode`, `rootVolumeResourceType`, `dataVolumeType`, `dataVolumeCount`, `dataVolumeSize`, `dataVolumeProductId`, `dataVolumeResourceSpecCode`, `dataVolumeResourceType`
  - The request parameter `data_volume_type`, `data_volume_count`, `data_volume_size` changed to required of the interface `UpdateClusterScaling`

### HuaweiCloud SDK VOD

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `M3U8` to the request parameter `video_type` to the interface `UploadMetaDataByUrl`
  - Add the response parameter `sign_url` to the interface `PublishAssets`
  - Add the response parameter `sign_url` to the interface `UnpublishAssets`
  - Add the response parameter `sign_url` to the interface `ShowAssetMeta`
  - Add the response parameter `sign_url` to the interface `ShowAssetDetail`
  - Add the response parameter `sign_url` to the interface `ShowTakeOverTaskDetails`
  - Add the response parameter `sign_url` to the interface `ShowTakeOverAssetDetails`

### HuaweiCloud SDK VPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `neutron:VIP_PORT`, Remove the enum values `network:VIP_PORT` from the response parameter `device_owner` to the interface `ListPorts`
  - Add the enum values `neutron:VIP_PORT`, Remove the enum values `network:VIP_PORT` from the response parameter `device_owner` to the interface `UpdatePort`
  - Add the enum values `neutron:VIP_PORT`, Remove the enum values `network:VIP_PORT` from the response parameter `device_owner` to the interface `ShowPort`

# 0.0.88 2022-05-12

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the interfaces `AddProtectBranchV2`, `AddTagV2`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateTask`:
    - Modify the type `string` -> `array` of the request parameter `bodys`
    - Modify the type `string` -> `object` of the request parameter `data`
    - Modify the type `string` -> `array` of the response parameter `bodys`
    - Modify the type `string` -> `object` of the response parameter `data`
  - Changes of the interface `ShowTask`:
    - Modify the type `string` -> `array` of the response parameter `bodys`
    - Modify the type `string` -> `object` of the response parameter `data`
  - Changes of the interface `UpdateCase`:
    - Modify the type `string` -> `array` of the request parameter `bodys`
    - Modify the type `string` -> `object` of the request parameter `data`
  - Changes of the interface `UpdateTemp`:
    - Modify the type `string` -> `array` of the request parameter `bodys`
    - Modify the type `string` -> `object` of the request parameter `data`

### HuaweiCloud SDK CSS

- _Features_
  - Support the interface `DownloadCert`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DWS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `nodes` to the interface `ListClusterDetails`
  - Add the response parameter `nodes` to the interface `ListClusters`

### HuaweiCloud SDK FRS

- _Features_
  - Support the following interfaces:
    - `DetectLiveByUrlIntl`
    - `DetectLiveByFileIntl`
    - `DetectLiveByBase64Intl`
    - `DetectFaceByFileIntl`
    - `DetectFaceByUrlIntl`
    - `DetectFaceByBase64Intl`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the following interfaces:
    - `ListBackups`
    - `CreateManualBackup`
    - `DeleteManualBackup`
    - `ListRestoreTimes`
    - `CreateRestoreInstance`
    - `CreateDatabase`
    - `CreateDbUser`
    - `CreateDatabaseSchemas`
    - `AllowDbPrivileges`
    - `SetDbUserPwd`
    - `ListDatabases`
    - `ListDbUsers`
    - `ListDatabaseSchemas`
    - `ShowBackupPolicy`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSimCards`:
    - Add the request parameters `min_used_flow`, `max_used_flow`, `min_left_flow`, `max_left_flow`
    - Remove the request parameters `min_flow`, `max_flow`

### HuaweiCloud SDK IAM

- _Features_
  - Support the interface `ShowDomainRoleAssignments`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Moderation

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RunImageModeration`:
    - Add the request parameter `show_ocr_text`
    - Add the response parameter `ocr_text`

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeHealthCode`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RestoreExistInstance`:
    - Add the request parameter `RestoreExistingInstanceRequestBody`
    - Remove the request parameter `RestoreToExistingInstanceRequest`

### HuaweiCloud SDK RocketMQ

- _Features_
  - Support the service `RocketMQ`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.87 2022-05-05

### HuaweiCloud SDK Moderation

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `byte` -> `string` of the request parameter `image` of the interface `RunImageModeration`

### HuaweiCloud SDK RES

- _Features_
  - Support the service `Recommender System`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.86 2022-04-28

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `string` -> `int32` of the request parameter `priority` of the interface `AddOrUpdateServiceDiscoveryRules`
  - Modify the type `string` -> `int32` of the response parameter `priority` of the interface `ListServiceDiscoveryRules`

### HuaweiCloud SDK CSE

- _Features_
  - Support the service `Cloud Service Engine`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DevStar

- _Features_
  - Support the interface `ConfirmDeploymentJob`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `cci` to the interface `CreateDeploymentJobs`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - Support the interfaces `CancelAsyncInvocation`, `StartSyncWorkflowExecution`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListFunctionStatistics`:
    - Modify the type `int32` -> `int64` of the response parameter `timestamp`
    - Modify the type `int32` -> `double` of the response parameter `value`
  - Changes of the interface `ListStatistics`:
    - Modify the type `int32` -> `int64` of the response parameter `timestamp`
    - Modify the type `int32` -> `double` of the response parameter `value`
  - Add the response parameter `enable_async_status_log` to the interface `ListFunctionAsyncInvokeConfig`
  - Add the response parameter `enable_async_status_log` to the interface `ShowFunctionAsyncInvokeConfig`
  - Changes of the interface `UpdateFunctionAsyncInvokeConfig`:
    - Add the request parameter `enable_async_status_log`
    - Add the response parameter `enable_async_status_log`
  - Changes of the interface `CreateWorkflow`:
    - Add the request parameters `mode`, `express_config`
    - Add the enum values `End` to the request parameter `type`
  - Changes of the interface `ShowWorkFlow`:
    - Add the response parameters `mode`, `express_config`
    - Add the enum values `End` to the response parameter `type`
  - Changes of the interface `UpdateWorkFlow`:
    - Add the request parameters `mode`, `express_config`
    - Add the enum values `End` to the request parameter `type`
  - Changes of the interface `ShowTenantMetric`:
    - Add the request parameters `start_time`, `end_time`
    - Modify the type `int32` -> `int64` of the response parameter `timestamp`
    - Modify the type `int32` -> `double` of the response parameter `value`
  - Changes of the interface `ShowWorkFlowMetric`:
    - Add the request parameters `start_time`, `end_time`
    - Modify the type `int32` -> `int64` of the response parameter `timestamp`
    - Modify the type `int32` -> `double` of the response parameter `value`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interface `ShowGaussMySqlProxyList`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateGaussMySqlReadonlyNode`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Remove the response parameters `size`, `order_id` from the interface `ExpandGaussMySqlInstanceVolume`
  - Changes of the interface `ListGaussMySqlDedicatedResources`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Add the request parameter `CloseMysqlProxyRequest` to the interface `DeleteGaussMySqlProxy`
  - Changes of the interface `CreateGaussMySqlProxy`:
    - Add the request parameters `proxy_name`, `proxy_mode`, `nodes_read_weight`
    - The request parameter `flavor_ref`, `node_num` changed to required
  - Add the request parameter `proxy_id` to the interface `ExpandGaussMySqlProxy`
  - Changes of the interface `ListGaussMySqlErrorLog`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
    - The request parameter `node_id` changed to not required
  - Changes of the interface `ListGaussMySqlSlowLog`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Changes of the interface `ListGaussMySqlConfigurations`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Changes of the interface `ShowGaussMySqlJobInfo`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Changes of the interface `ListInstanceTags`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Changes of the interface `ListProjectTags`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`
  - Changes of the interface `BatchTagAction`:
    - Add the request parameter `X-Auth-Token`
    - Remove the request parameter `x-auth-token`

### HuaweiCloud SDK Meeting

- _Features_
  - Support the interface `ShowDepartment`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `deptCodes` to the interface `SearchCorpDir`

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the response body's type of the interface `RecognizeMyanmarDriverLicense` is incorrect.
- _Change_
  - None

### HuaweiCloud SDK WAF

- _Features_
  - Support the interfaces `DeleteIgnoreRule`, `CreateIgnoreRule`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `id`, `policyid`, `timestamp`, `description`, `status`, `url`, `rule`, `domain`, `url_logic`, `advanced` from the interface `ListIgnoreRule`
  - Add the response parameter `producer` to the interface `ListValueList`
  - Add the response parameters `process_time`, `request_body` to the interface `ListEvent`
  - Changes of the interface `ShowEvent`:
    - Add the response parameters `process_time`, `request_body`
    - Modify the type `string` -> `object` of the response parameter `headers`
    - Modify the type `string` -> `int32` of the response parameter `response_size`
    - Modify the type `string` -> `int64` of the response parameter `response_time`
  - Modify the type `string` -> `enum` of the response parameter `paid_type` of the interface `ListHost`
  - Add the response parameters `flag`, `http2_enable` to the interface `CreateHost`
  - Add the response parameter `protect_status` to the interface `UpdateHostProtectStatus`
  - Modify the type `string` -> `enum` of the response parameter `paid_type` of the interface `DeleteHost`
  - Modify the type `string` -> `enum` of the response parameter `protocol` of the interface `UpdateHost`
  - Add the response parameters `id`, `name`, `level`, `action`, `options`, `full_detection`, `hosts`, `bind_host`, `timestamp`, `extend` to the interface `DeletePolicy`
  - Add the response parameters `id`, `policyid`, `timestamp`, `description`, `status` to the interface `UpdatePolicyRuleStatus`
  - Add the response parameters `id`, `policyid`, `timestamp`, `description`, `status`, `url`, `category`, `index` to the interface `DeletePrivacyRule`
  - Add the response parameter `region` to the interface `DeletePremiumHost`

# 0.0.85 2022-04-21

### HuaweiCloud SDK AS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `iam_agency_name` to the interface `CreateScalingGroup`
  - Add the response parameter `iam_agency_name` to the interface `ListScalingGroups`
  - Add the request parameter `iam_agency_name` to the interface `UpdateScalingGroup`
  - Add the response parameter `iam_agency_name` to the interface `ShowScalingGroup`

### HuaweiCloud SDK BSS

- _Features_
  - Support the interface `ListConsumeSubCustomers`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `sub_service_type_code`, `sub_service_type_name`, `sub_resource_type_code`, `sub_resource_type_name`, `sub_resource_id`, `sub_resource_name` to the interface `ListCustomerBillsMonthlyBreakDown`

### HuaweiCloud SDK CloudDeploy

- _Features_
  - Support the interfaces `ListDeployTasks`, `ListDeployTaskHistoryByDate`, `ShowProjectSuccessRate`, `ListTaskSuccessRate`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `domain_name_info` to the interface `ShowInstance`

### HuaweiCloud SDK DWS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `tags` to the interface `ListClusters`
  - Remove the response parameters `cluster_id`, `size`, `name`, `description`, `finished`, `started`, `id`, `type`, `status` from the interface `ListSnapshotDetails`

### HuaweiCloud SDK IES

- _Features_
  - Support the service `Intelligent EdgeSite`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `restore_all_database` from the interface `RestoreToExistingInstance`

# 0.0.84 2022-04-14

### HuaweiCloud SDK Core

- _Features_
  - Support for obtaining temporary authentication information from instance's metadata
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `sources`, `origin_protocol`, `force_redirect`, `compress` to the interface `UpdateDomainFullConfig`
  - Changes of the interface `ShowDomainFullConfig`:
    - Add the response parameters `sources`, `origin_protocol`, `force_redirect`, `compress`
    - Modify the type `string` -> `int32` of the response parameter `certificate_source`

### HuaweiCloud SDK CloudBuild

- _Features_
  - Support the interface `ShowJobSuccessRatio`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `ShowTasksRulesets`, `CheckRulesetParameters`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CSMS

- _Features_
  - Support the `Cloud Secret Management Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `rename_commands` to the interface `UpdateInstance`

### HuaweiCloud SDK DRS

- _Features_
  - Support the interface `BatchSetSmn`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `X-Language` changed to not required of the interface `ShowQuotas`
  - The request parameter `X-Language` changed to not required of the interface `BatchCreateJobs`
  - The request parameter `X-Language` changed to not required of the interface `BatchValidateConnections`
  - The request parameter `X-Language` changed to not required of the interface `BatchValidateClustersConnections`
  - The request parameter `X-Language` changed to not required of the interface `BatchSetObjects`
  - The request parameter `X-Language` changed to not required of the interface `BatchCheckJobs`
  - The request parameter `X-Language` changed to not required of the interface `BatchCheckResults`
  - Changes of the interface `BatchSetSpeed`:
    - Add the request parameter `is_utc`
    - The request parameter `X-Language` changed to not required
  - The request parameter `X-Language` changed to not required of the interface `BatchShowParams`
  - The request parameter `X-Language` changed to not required of the interface `UpdateParams`
  - The request parameter `X-Language` changed to not required of the interface `BatchStartJobs`
  - The request parameter `X-Language` changed to not required of the interface `BatchRestoreTask`
  - The request parameter `X-Language` changed to not required of the interface `BatchStopJobs`
  - The request parameter `X-Language` changed to not required of the interface `BatchDeleteJobs`
  - Changes of the interface `BatchUpdateJob`:
    - The request parameter `X-Language` changed to not required
    - The request parameter `endpoints`, `protocol` changed to not required
  - The request parameter `X-Language` changed to not required of the interface `BatchResetPassword`
  - The request parameter `X-Language` changed to not required of the interface `BatchSetDefiner`
  - The request parameter `X-Language` changed to not required of the interface `CreateCompareTask`
  - The request parameter `X-Language` changed to not required of the interface `ListCompareResult`
  - The request parameter `X-Language` changed to not required of the interface `BatchListProgresses`
  - Changes of the interface `BatchListJobDetails`:
    - Add the response parameter `is_utc`
    - The request parameter `X-Language` changed to not required
    - Modify the type `string` -> `object` of the response parameter `alarm_notify`
  - Changes of the interface `ShowJobList`:
    - The request parameter `X-Language` changed to not required
    - Modify the type `string` -> `boolean` of the response parameter `billing_tag`
    - Modify the type `string` -> `boolean` of the response parameter `node_newFramework`
  - The request parameter `X-Language` changed to not required of the interface `BatchListJobStatus`
  - The request parameter `X-Language` changed to not required of the interface `BatchChangeData`
  - The request parameter `X-Language` changed to not required of the interface `BatchSwitchover`
  - The request parameter `X-Language` changed to not required of the interface `BatchListRposAndRtos`
  - The request parameter `X-Language` changed to not required of the interface `ShowMonitoringData`
  - The request parameter `X-Language` changed to not required of the interface `BatchListStructProcess`
  - The request parameter `X-Language` changed to not required of the interface `BatchListStructDetail`
  - The request parameter `X-Language` changed to not required of the interface `BatchUpdateUser`
  - The request parameter `X-Language` changed to not required of the interface `ListUsers`
  - Changes of the interface `BatchSetPolicy`:
    - Modify the type `string` -> `boolean` of the request parameter `export_snapshot`
    - The request parameter `X-Language` changed to not required
    - The request parameter `conflict_policy`, `filter_ddl_policy` changed to not required

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `ip_eq` to the interface `ListServersDetails`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `encrypted_user_data` to the interface `UpdateFunctionConfig`

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the interfaces `RestartInstance`, `ShowInstanceConfiguration`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateInstance`:
    - Remove the request parameter `solution`
    - The request parameter `sharding_num`, `coordinator_num` changed to not required
  - Add the response parameter `count` to the interface `ListConfigurations`
  - Add the response parameter `total` to the interface `ListFlavors`

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSimPricePlans`:
    - Modify the type `date` -> `date-time` of the response parameter `create_time`
    - Modify the type `date` -> `date-time` of the response parameter `active_time`
    - Modify the type `date` -> `date-time` of the response parameter `stop_time`

### HuaweiCloud SDK IEC

- _Features_
  - Support the following interfaces:
    - `ListRelatedRoutetables`
    - `ListRoutetables`
    - `CreateRoutetable`
    - `ShowRoutetable`
    - `UpdateRoutetable`
    - `DeleteRoutetable`
    - `AssociateSubnet`
    - `DisassociateSubnet`
    - `ListRoutes`
    - `CreateRoutes`
    - `UpdateRoutes`
    - `DeleteRoutes`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `action`, `priority` to the interface `ListSecurityGroups`
  - Add the response parameters `action`, `priority` to the interface `CreateSecurityGroup`
  - Add the response parameters `action`, `priority` to the interface `ShowSecurityGroup`
  - Add the response parameters `action`, `priority` to the interface `ListSecurityGroupRules`
  - Changes of the interface `CreateSecurityGroupRule`:
    - Add the request parameters `action`, `priority`
    - Add the response parameters `action`, `priority`
  - Add the response parameters `action`, `priority` to the interface `ShowSecurityGroupRule`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateAomMappingRules`:
    - The request parameter `target_log_group_id`, `target_log_group_name`, `target_log_stream_id`, `target_log_stream_name` changed to required
    - The request parameter `container_name` changed to not required
    - The response parameter `target_log_group_id`, `target_log_group_name`, `target_log_stream_id`, `target_log_stream_name` changed to required
    - The response parameter `container_name` changed to not required
  - Changes of the interface `ShowAomMappingRules`:
    - The response parameter `target_log_group_id`, `target_log_group_name`, `target_log_stream_id`, `target_log_stream_name` changed to required
    - The response parameter `container_name` changed to not required
  - Changes of the interface `CreateAomMappingRules`:
    - The request parameter `target_log_group_id`, `target_log_group_name`, `target_log_stream_id`, `target_log_stream_name` changed to required
    - The request parameter `container_name` changed to not required
    - The response parameter `target_log_group_id`, `target_log_group_name`, `target_log_stream_id`, `target_log_stream_name` changed to required
    - The response parameter `container_name` changed to not required
  - Changes of the interface `ShowAomMappingRule`:
    - The response parameter `target_log_group_id`, `target_log_group_name`, `target_log_stream_id`, `target_log_stream_name` changed to required
    - The response parameter `container_name` changed to not required
  - The request parameter `host_id_list` changed to required of the interface `ListHost`
  - The request parameter `host_group_id_list` changed to required of the interface `ListHostGroup`
  - The request parameter `access_config_name_list`, `host_group_name_list`, `log_group_name_list`, `log_stream_name_list` changed to required of the interface `ListAccessConfig`

### HuaweiCloud SDK Moderation

- _Features_
  - Support the interface `RunModerationAudio`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SCM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListCertificates`:
    - Add the request parameters `enterprise_project_id`, `deploy_support`
    - Add the response parameter `enterprise_project_id`
    - The response parameter `id`, `name`, `domain`, `sans`, `signature_algorithm`, `deploy_support`, `type`, `brand`, `expire_time`, `domain_type`, `validity_period`, `status`, `domain_count`, `wildcard_count`, `description`, `total_count` changed to required
  - Changes of the interface `ImportCertificate`:
    - Add the request parameter `enterprise_project_id`
    - The response parameter `certificate_id` changed to required
  - Changes of the interface `ShowCertificate`:
    - Add the response parameter `enterprise_project_id`
    - The response parameter `id`, `status`, `order_id`, `name`, `type`, `brand`, `push_support`, `revoke_reason`, `signature_algorithm`, `issue_time`, `not_before`, `not_after`, `validity_period`, `validation_method`, `domain_type`, `domain`, `sans`, `domain_count`, `wildcard_count` changed to required
  - The response parameter `certificate`, `certificate_chain`, `private_key` changed to required of the interface `ExportCertificate`

### HuaweiCloud SDK SMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `process_trace` to the interface `ListTasks`
  - Add the response parameter `process_trace` to the interface `ShowTask`
  - Add the request parameter `process_trace` to the interface `UpdateTask`
  - Changes of the interface `UpdateTaskSpeed`:
    - Add the request parameters `process_trace`, `compress_rate`
    - The request parameter `migrate_speed` changed to not required

# 0.0.83 2022-04-07

### HuaweiCloud SDK BSS

- _Features_
  - Support the interface `ListStoredValueCards`
- _Bug Fix_
  - None
- _Change_
  - Remove the interfaces `ListSubCustomerDiscounts`, `BatchSetSubCustomerDiscount`
  - Add the response parameters `resource_type_name`, `service_type_name` to the interface `ShowRefundOrderDetails`
  - Add the response parameter `service_type_name` to the interface `ListCustomerOrders`
  - Add the response parameters `service_type_name`, `service_type_name` to the interface `ShowCustomerOrderDetails`
  - Add the response parameters `resource_type_name`, `service_type_name` to the interface `ListPayPerUseCustomerResources`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ListCustomerOnDemandResources`
  - Add the response parameters `cloud_service_type_name`, `resource_type_name` to the interface `ListSubcustomerMonthlyBills`
  - Add the response parameters `cloud_service_type_name`, `resource_type_name`, `period_type` to the interface `ListCustomerselfResourceRecordDetails`
  - Add the response parameters `cloud_service_type_name`, `resource_type_name` to the interface `ListCustomerselfResourceRecords`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ShowCustomerMonthlySum`
  - Changes of the interface `ListCustomerBillsFeeRecords`:
    - Add the request parameters `bill_date_begin`, `bill_date_end`
    - Add the response parameters `service_type_name`, `resource_type_name`
  - Add the response parameters `resource_type_name`, `service_type_name` to the interface `ListUsageTypes`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ListSubCustomerBillDetail`
  - Add the response parameters `service_type_name`, `resource_type_name` to the interface `ListCustomerBillsMonthlyBreakDown`
  - Add the response parameter `service_type_name` to the interface `ListFreeResourceInfos`
  - Add the response parameter `service_type_name` to the interface `ListIncentiveDiscountPolicies`

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameters `kind`, `apiVersion`, `status` from the interface `UpdateNodePool`

### HuaweiCloud SDK DSC

- _Features_
  - Support the following interfaces:
    - `CreateDocWatermarkByAddress`
    - `ShowDocWatermarkByAddress`
    - `ShowImageWatermarkWithImage`
    - `CreateImageWatermarkByAddress`
    - `ShowImageWatermarkByAddress`
    - `ShowImageWatermarkWithImageByAddress`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateImageWatermark`:
    - Add the request parameter `image_watermark`
    - The request parameter `blind_watermark` changed to required

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `price_plan_list` from the interface `StopSimCard`
  - Remove the request parameter `price_plan_list` from the interface `ResetSimCard`

### HuaweiCloud SDK Meeting

- _Features_
  - Support the interface `StartMeeting`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeWaybillElectronic`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `print_code` to the interface `RecognizeVatInvoice`
  - Changes of the interface `RecognizeVehicleLicense`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeTaxiInvoice`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Add the response parameters `type`, `accumulated_scores`, `status`, `generation_date`, `current_time` to the interface `RecognizeDriverLicense`
  - Changes of the interface `RecognizeTrainTicket`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeBankcard`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ApplyConfigurationAsync`, `UpdateInstanceConfigurationAsync`, `DeleteSqlserverDatabaseEx`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the interfaces `ListConfigurations`, `ListDatastores`, `ListFlavors`, `ListStorageTypes`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateInstance`:
    - Add the request parameter `solution`
    - Add the enum values `centralization_standard` to the request parameter `mode`
    - Add the enum values `ESSD` to the request parameter `type`

# 0.0.82 2022-03-31

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of concurrent read/write authentication information cache
- _Change_
  - None

### HuaweiCloud SDK CC

- _Features_
  - Support the service `Cloud Connect`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - Support the interface `BatchShowNodesInformation`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.81 2022-03-25

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `responseStatus` to the interface `DeleteserviceDiscoveryRules`
  - Add the response parameter `responseStatus` to the interface `AddOrUpdateServiceDiscoveryRules`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interfaces(v2):
    - `ShowDomainLocationStats`
    - `ShowDomainStats`
    - `ShowTopUrl`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DAS

- _Features_
  - Support the interface `ShowSqlExplain`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `group_name`, `replication_ip` to the interface `ListRedislog`

### HuaweiCloud SDK DNS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `routers` from the interface `ListPublicZones`
  - Add the request parameters `marker`, `limit`, `offset`, `line_id`, `tags`, `status`, `type`, `name`, `id`, `sort_key`, `sort_dir`, `search_mode` to the interface `ShowRecordSetByZone`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - Support the following interfaces:
    - `ListWorkflows`
    - `CreateWorkflow`
    - `BatchDeleteWorkflows`
    - `ListWorkflowExecutions`
    - `StartWorkflowExecution`
    - `ShowWorkflowExecution`
    - `ShowWorkFlow`
    - `UpdateWorkFlow`
    - `ShowTenantMetric`
    - `ShowWorkFlowMetric`
    - `RetryWorkFlow`
    - `StopWorkFlow`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListSimCards`:
    - Add the request parameters `min_flow`, `max_flow`, `order_id`, `filter_downtime_period`
    - Modify the type `date` -> `date-time` of the response parameter `device_status_date`
    - Modify the type `date` -> `date-time` of the response parameter `expire_time`
  - Add the request parameter `price_plan_list` to the interface `StopSimCard`
  - Add the request parameter `price_plan_list` to the interface `ResetSimCard`
  - Changes of the interface `ShowSimCard`:
    - Modify the type `date` -> `date-time` of the response parameter `device_status_date`
    - Modify the type `date` -> `date-time` of the response parameter `expire_time`

### HuaweiCloud SDK IMS

- _Features_
  - Support the interfaces `ListVersions`, `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - The request parameter `os_type` changed to not required of the interface `CreateDataImage`

### HuaweiCloud SDK IoTDA

- _Features_
  - Support the interface `ResetFingerprint`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `RecognizeVatInvoice`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeIdCard`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`
  - Changes of the interface `RecognizeDriverLicense`:
    - Add the request parameter `return_text_location`
    - Add the response parameter `text_location`

### HuaweiCloud SDK VSS

- _Features_
  - Support the following interfaces:
    - `ShowDomainSettings`
    - `UpdateDomainSettings`
    - `ListTaskHistories`
    - `ListPortResults`
    - `ListBusinessRisks`
    - `UpdateFalsePositive`
    - `CancelTasks`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `domain_id` to the interface `ListDomains`
  - Add the response parameter `hit_details` to the interface `ShowResults`

# 0.0.80 2022-03-10

### HuaweiCloud SDK BCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `operation_id` to the interface `BatchCreateChannels`
  - The request parameter `fabric_version` changed to not required of the interface `CreateNewBlockchain`

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `cluster_id` changed to not required of the interface `DeleteAddonInstance`

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `X-Auth-Token` from the interface `ShowTopUrl`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainLocationStats`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainItemDetails`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainStats`
  - Remove the request parameter `X-Auth-Token` from the interface `ShowDomainItemLocationDetails`

### HuaweiCloud SDK DWS

- _Features_
  - Support the following interfaces:
    - `ListClusterDetails`
    - `DeleteCluster`
    - `ResetPassword`
    - `ListSnapshots`
    - `CreateSnapshot`
    - `RestartCluster`
    - `ListClusters`
    - `CreateCluster`
    - `RestoreCluster`
    - `ResizeCluster`
    - `ListNodeTypes`
    - `ListSnapshotDetails`
    - `DeleteSnapshot`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ELB

- _Features_
  - Support the following interfaces:
    - `ListLogtanks`
    - `CreateLogtank`
    - `ShowLogtank`
    - `UpdateLogtank`
    - `DeleteLogtank`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `https_cps` to the interface `ListFlavors`
  - Add the response parameter `https_cps` to the interface `ShowFlavor`
  - The request parameter `X-Auth-Token` changed to not required of the interface `ListLoadBalancers`
  - The request parameter `X-Auth-Token` changed to not required of the interface `CreateLoadBalancer`
  - The request parameter `X-Auth-Token` changed to not required of the interface `UpdateIpList`
  - The request parameter `X-Auth-Token` changed to not required of the interface `BatchDeleteIpList`

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `date` -> `date-time` of the response parameter `act_date` of the interface `ListSimCards`
  - Modify the type `date` -> `date-time` of the response parameter `act_date` of the interface `ShowSimCard`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `count` to the interface `ListLogs`

### HuaweiCloud SDK Meeting

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `sortLevel` to the interface `AddDepartment`
  - Add the request parameter `sortLevel` to the interface `UpdateDepartment`
  - Add the response parameters `sortLevel`, `sortLevel` to the interface `ShowDeptAndChildDept`
  - Add the response parameter `sortLevel` to the interface `SearchDepartmentByName`

# 0.0.79 2022-03-07

### HuaweiCloud SDK Core

- _Features_
  - Support SK derived authentication
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - Support the interfaces `UpdateClusterEip`, `ShowClusterEndpoints`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CES

- _Features_
  - Support the following interfaces (V2):
    - `ListAlarms`
    - `CreateAlarm`
    - `DeleteAlarm`
    - `UpdateAlarmAction`
    - `ListAlarmResources`
    - `DeleteAlarmResources`
    - `AddAlarmResources`
    - `AddResourceGroupsResourcesBatch`
    - `DeleteResourceGroupsResourcesBatch`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `gaussdbv5`, `postgresql` to the request parameter `engine_type` to the interface `BatchCreateJobs`
  - Changes of the interface `BatchValidateConnections`:
    - Add the request parameter `kafka_security_config`
    - Add the enum values `postgresql` to the request parameter `db_type`
  - Changes of the interface `BatchUpdateUser`:
    - Add the request parameter `is_sync_object_privilege`
    - Add the response parameters `no_privileges`, `parent_account`, `no_parent_account`
  - Add the response parameters `no_privileges`, `parent_account`, `no_parent_account` to the interface `ListUsers`
  - Add the request parameters `topic_policy`, `topic`, `partition_policy`, `kafka_data_format`, `topic_name_format`, `partitions_num`, `replication_factor`, `is_fill_materialized_view`, `export_snapshot` to the interface `BatchSetPolicy`

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `ip_version` of the interface `CreatePrePaidPublicip`: `integer` -> `enum`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `order_id` to the interface `ShrinkInstanceNode`

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `error_code` to the interface `ListEditingJob`

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `request_id` from the interface `DownloadSlowlog`

# 0.0.78 2022-02-25

### HuaweiCloud SDK AS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `alarm_id` to the interface `ListAllScalingV2Policies`
  - Add the enum values `GPSSD` to the request parameter `volume_type` to the interface `CreateScalingConfig`
  - Add the response parameter `min` to the interface `ShowResourceQuota`
  - Add the response parameter `min` to the interface `ShowPolicyAndInstanceQuota`

### HuaweiCloud SDK BMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateBaremetalServerMetadata`:
    - Modify the type of the request body `MetaData` -> `UpdateBaremetalServerMetadataReq`
    - Remove the response parameter `key`

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The response parameter `from-connector-name`, `to-link-name` changed to required of the interface `ShowJobs`
  - The request parameter `from-connector-name`, `to-link-name` changed to required of the interface `UpdateJob`
  - Changes of the interface `CreateAndStartRandomClusterJob`:
    - The request parameter `from-connector-name`, `to-link-name` changed to required
    - Modify the type `int32` -> `float` of the response parameter `progress`
    - Modify the type `boolean` -> `string` of the response parameter `isStopingIncrement`
  - Add the response parameter `submissions` to the interface `StopJob`
  - The request parameter `from-connector-name`, `to-link-name` changed to required of the interface `CreateJob`
  - Changes of the interface `StartJob`:
    - Modify the type `int32` -> `float` of the response parameter `progress`
    - Modify the type `boolean` -> `string` of the response parameter `isStopingIncrement`
  - Modify the type `double` -> `float` of the response parameter `progress` of the interface `ShowJobStatus`
  - Modify the type `double` -> `float` of the response parameter `progress` of the interface `ShowSubmissions`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interfaces `ShowDomainLocationStats`, `ShowDomainFullConfig`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowDomainStats`:
    - Add the request parameter `service_area`
    - Remove the request parameters `X-Auth-Token`, `country`, `district`, `isp`
    - Remove the response parameters `start_time`, `end_time`, `interval`, `action`, `stat_type`, `group_by`
  - Add the request parameter `https` to the interface `UpdateDomainFullConfig`

### HuaweiCloud SDK CloudIDE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `instance_domain_id`, `instance_user_id` to the interface `CreateInstance`
  - The request parameter `instance_user_domain_name`, `instance_user_name` changed to not required of the interface `CreateInstanceBy3rd`

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interface `CheckRecord`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudRTC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `array` -> `string` of the request parameter `mid` of the interface `ListRtcClientQosDetails`

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the interface `ShowStatisticCommitV3`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListProjectSets`:
    - Add the response parameters `CreateTime`, `UpdateTime`, `external_params`, `variables_no_file`
    - Remove the response parameters `create_time`, `update_time`, `group`
  - The request parameter `name` changed to required of the interface `UpdateProject`
  - Changes of the interface `ShowTask`:
    - Add the response parameters `parallel`, `contents`, `sort`, `related_temp_running_data`
    - Remove the response parameter `content`
  - Changes of the interface `UpdateTask`:
    - Add the request parameters `contents`, `sort`, `related_temp_running_data`
    - Add the response parameters `parallel`, `contents`, `sort`, `related_temp_running_data`
    - Remove the request parameter `content`
    - Remove the response parameter `content`
    - The request parameter `name` changed to required
  - Changes of the interface `ShowReport`:
    - Add the response parameters `performance`, `minNetworkTraffic`, `avgNetworkTraffic`, `maxNetworkTraffic`, `branchId`, `branchName`, `projectId`, `serviceId`
    - Remove the response parameters `progressState`, `statusValue`
    - Modify the type `float` -> `double` of the response parameter `averageRespTime`
    - Modify the type `float` -> `double` of the response parameter `avgRecBytes`
    - Modify the type `int32` -> `double` of the response parameter `avgSentBytes`
    - Modify the type `string` -> `double` of the response parameter `avgTranRespTime`
    - Modify the type `int32` -> `double` of the response parameter `currentThreadNum`
    - Modify the type `int32` -> `double` of the response parameter `errorCount`
    - Modify the type `int32` -> `double` of the response parameter `errorEventsCount`
    - Modify the type `int32` -> `double` of the response parameter `failedAssert`
    - Modify the type `int32` -> `double` of the response parameter `failedOthers`
    - Modify the type `int32` -> `double` of the response parameter `failedParsed`
    - Modify the type `int32` -> `double` of the response parameter `failedRefused`
    - Modify the type `int32` -> `double` of the response parameter `failedTimeout`
    - Modify the type `int32` -> `double` of the response parameter `max`
    - Modify the type `int32` -> `double` of the response parameter `maxRecBytes`
    - Modify the type `int32` -> `double` of the response parameter `maxRespTime`
    - Modify the type `int32` -> `double` of the response parameter `maxSentBytes`
    - Modify the type `int32` -> `double` of the response parameter `maxTranRespTime`
    - Modify the type `int32` -> `double` of the response parameter `min`
    - Modify the type `int32` -> `double` of the response parameter `requests`
    - Modify the type `int32` -> `double` of the response parameter `result`
    - Modify the type `int32` -> `double` of the response parameter `status`
    - Modify the type `int32` -> `double` of the response parameter `successCount`
    - Modify the type `int32` -> `double` of the response parameter `successRate`
    - Modify the type `int32` -> `double` of the response parameter `sum1xx`
    - Modify the type `int32` -> `double` of the response parameter `sum2xx`
    - Modify the type `int32` -> `double` of the response parameter `sum3xx`
    - Modify the type `int32` -> `double` of the response parameter `sum4xx`
    - Modify the type `int32` -> `double` of the response parameter `sum5xx`
    - Modify the type `int32` -> `double` of the response parameter `taskStatus`
    - Modify the type `int32` -> `double` of the response parameter `tp50`
    - Modify the type `int32` -> `double` of the response parameter `tp75`
    - Modify the type `int32` -> `double` of the response parameter `tp90`
    - Modify the type `int32` -> `double` of the response parameter `tp95`
    - Modify the type `int32` -> `double` of the response parameter `tp99`
    - Modify the type `float` -> `double` of the response parameter `tps`
    - Modify the type `string` -> `double` of the response parameter `tranTPS`
    - Modify the type `string` -> `double` of the response parameter `transactionSuccess`
    - Modify the type `int32` -> `double` of the response parameter `transactionalSuccessRate`
    - Modify the type `int32` -> `double` of the response parameter `transactionalTps`
    - Modify the type `int32` -> `double` of the response parameter `transactionalTpsSuccess`
    - Modify the type `string` -> `double` of the response parameter `transactions`
    - Modify the type `int32` -> `double` of the response parameter `vum`
    - Modify the type `float` -> `double` of the response parameter `avgResponseTime`
    - Modify the type `int32` -> `double` of the response parameter `caseRetry`
    - Modify the type `int32` -> `double` of the response parameter `completeNum`
    - Modify the type `int32` -> `double` of the response parameter `duration`
    - Modify the type `int32` -> `double` of the response parameter `executedNum`
    - Modify the type `int32` -> `double` of the response parameter `kpiCaseCount`
    - Modify the type `int32` -> `double` of the response parameter `kpiCaseExecuteCount`
    - Modify the type `int32` -> `double` of the response parameter `kpiCasePassCount`
    - Modify the type `int32` -> `double` of the response parameter `maxUsers`
    - Modify the type `int32` -> `double` of the response parameter `passNum`
    - Modify the type `int32` -> `double` of the response parameter `stage`
    - Modify the type `int32` -> `double` of the response parameter `totalNum`
  - Changes of the interface `UpdateCase`:
    - Add the request parameters `contents`, `sort`
    - Remove the request parameter `content`
  - Add the request parameter `contents` to the interface `CreateTemp`
  - Changes of the interface `UpdateTemp`:
    - Modify the type `array` -> `string` of the request parameter `bodys`
    - The request parameter `name` changed to required
  - Add the request parameter `is_quoted` to the interface `CreateVariable`
  - Changes of the interface `ShowHistoryRunInfo`:
    - Modify the type `int32` -> `double` of the response parameter `run_id`
    - Modify the type `int32` -> `double` of the response parameter `run_type`
    - Modify the type `int32` -> `double` of the response parameter `continue_time`

### HuaweiCloud SDK CSS

- _Features_
  - Support the following interfaces:
    - `UpdateFlavor`
    - `UpdateFlavorByType`
    - `UpdateShrinkNodes`
    - `UpdateShrinkCluster`
    - `ListLogsJob`
    - `ShowClusterDetail`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `id`, `instances` from the interface `UpdateExtendCluster`
  - Remove the request parameter `status` from the interface `StartConnectivityTest`

### HuaweiCloud SDK DDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `group_id` to the interface `ExpandInstanceNodes`
  - Add the request parameter `group_id` to the interface `ShrinkInstanceNodes`
  - The request parameter `shard_unit` changed to not required of the interface `CreateDatabase`

### HuaweiCloud SDK DGC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListConnections`:
    - Add the response parameters `type`, `description`
    - Remove the response parameter `connectionType`
    - Modify the type `string` -> `int32` of the response parameter `total`
    - The response parameter `name` changed to required
  - Changes of the interface `CreateConnection`:
    - Add the request parameters `type`, `description`
    - Remove the request parameter `connectionType`
    - The request parameter `name` changed to required
  - Changes of the interface `ShowConnection`:
    - Add the response parameters `type`, `description`
    - Remove the response parameter `connectionType`
    - The response parameter `name` changed to required
  - Changes of the interface `UpdateConnection`:
    - Add the request parameters `type`, `description`
    - Remove the request parameter `connectionType`
    - The request parameter `name` changed to required
  - Changes of the interface `ExecuteScript`:
    - Add the response parameter `instanceId`
    - Remove the response parameter `jobId`
    - Modify the type `string` -> `object` of the request parameter `params`

### HuaweiCloud SDK ELB

- _Features_
  - Support the interfaces `BatchCreateMembers`, `BatchDeleteMembers`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ListFunctions`
  - Changes of the interface `CreateFunction`:
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the request parameter `runtime`
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ShowFunctionCode`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `UpdateFunctionCode`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ShowFunctionConfig`
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the request parameter `runtime`
    - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ListFunctionVersions`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `CreateFunctionVersion`
  - Add the enum values `Java11`, `Node.js14.18`, `Python3.9` to the request parameter `runtime` to the interface `CreateDependency`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `PHP 7.3` from the request parameter `runtime` to the interface `UpdateDependency`
  - Add the enum values `Java8`, `Java11`, `Node.js6.10`, `Node.js8.10`, `Node.js10.16`, `Node.js12.13`, `Node.js14.18`, `Python2.7`, `Python3.6`, `Python3.9`, `Go1.8`, `Go1.x`, `PHP7.3`, Remove the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3` from the response parameter `runtime` to the interface `ImportFunction`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the interfaces `UpdateAuditLog`, `ShowAuditLog`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `ListSingleStreamDetail`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK LTS

- _Features_
  - Support the interfaces `UpdateStructConfig`, `CreateStructConfig`, `ListStructTemplate`, `ListBreifStructTemplate`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `tag` to the interface `ListLogGroups`
  - Add the response parameter `tag` to the interface `ListLogStream`

### HuaweiCloud SDK Meeting

- _Features_
  - Support the service `Meeting`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interfaces `BatchUpdateChildNickNames`, `ListIterationHistories`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `updated_time_interval`, `include_deleted` to the interface `ListProjectIterationsV4`
  - Add the request parameters `include_deleted`, `updated_time_interval` to the interface `ListIssuesV4`
  - Add the response parameters `description`, `order`, `accessories` to the interface `ShowIssueV4`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ListSlowLogFile`, `StopInstance`, `StartupInstance`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SCM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `sans`, `signature_algorithm`, `deploy_support` to the interface `ListCertificates`
  - Add the request parameters `enc_certificate`, `enc_private_key` to the interface `ImportCertificate`
  - Changes of the interface `ShowCertificate`:
    - Add the response parameter `signature_algorithm`
    - Remove the response parameter `signature_algrithm`
  - Add the response parameters `enc_certificate`, `enc_private_key` to the interface `ExportCertificate`

### HuaweiCloud SDK VOD

- _Features_
  - Support the interface `ListDomainLogs`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `delete_type` to the interface `DeleteAssets`

### HuaweiCloud SDK VPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `subnetpool_id` to the interface `NeutronListSubnets`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListWhiteblackipRule`:
    - Add the response parameter `addr`
    - Remove the response parameter `ip`

# 3.0.77 2022-02-10

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateAlarmRule`:
    - Modify the type `string` -> `enum` of the request parameter `statistic`
    - The request parameter `alarm_level`, `comparison_operator`, `evaluation_periods`, `metric_name`, `namespace`, `period`, `statistic`, `threshold`, `unit` changed to not required

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListRateOnPeriodDetail`:
    - Add the request parameter `fee_installment_mode`
    - Add the response parameters `installment_official_website_amount`、`installment_period_type`、`installment_official_discount_amount`、`installment_amount`

### HuaweiCloud SDK CBR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListProtectable`:
    - Modify the type `string` -> `boolean` of the response parameter `result`
    - Modify the type `string` -> `int32` of the response parameter `size`
  - Changes of the interface `ShowProtectable`:
    - Modify the type `string` -> `boolean` of the response parameter `result`
    - Modify the type `string` -> `int32` of the response parameter `size`

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `kind`, `apiVersion`, `metadata`, `spec`, `status` from the interface `CreateAddonInstance`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `ListNodes`
  - Add the request parameter `isStatic` to the interface `CreateNode`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `DeleteNode`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `ShowNode`
  - Add the response parameters `isStatic`, `privateIPv6IP` to the interface `UpdateNode`
  - Changes of the interface `RemoveNode`:
    - The request parameter `uid` changed to required
    - The response parameter `uid` changed to required
  - Changes of the interface `MigrateNode`:
    - The request parameter `uid` changed to required
    - The response parameter `uid` changed to required
  - Add the response parameter `isStatic` to the interface `ListNodePools`
  - Add the request parameter `isStatic` to the interface `CreateNodePool`
  - Add the response parameter `isStatic` to the interface `DeleteNodePool`
  - Add the response parameter `isStatic` to the interface `ShowNodePool`
  - Changes of the interface `UpdateNodePool`:
    - Add the request parameter `isStatic`
    - Add the response parameter `isStatic`

### HuaweiCloud SDK CSS

- _Features_
  - Support the following interfaces:
    - `UpdateOndemandClusterToPeriod`
    - `UpdateClusterName`
    - `ResetPassword`
    - `StartKibanaPublic`
    - `UpdateCloseKibana`
    - `UpdateAlterKibana`
    - `UpdatePublicKibanaWhitelist`
    - `StopPublicKibanaWhitelist`
    - `CreateCnf`
    - `UpdateCnf`
    - `StartPipeline`
    - `StopPipeline`
    - `AddFavorite`
    - `StartConnectivityTest`
    - `ListTemplates`
    - `ListConfs`
    - `ListPipelines`
    - `ListActions`
    - `ShowGetConfDetail`
    - `DeleteConf`
    - `DeleteTemplate`
    - `StartLogs`
    - `StopLogs`
    - `ShowGetLogSetting`
    - `UpdateLogSetting`
    - `StartLogAutoBackupPolicy`
    - `StopLogAutoBackupPolicy`
    - `CreateLogBackup`
    - `ShowLogBackup`
    - `CreateBindPublic`
    - `UpdateUnbindPublic`
    - `UpdatePublicBandWidth`
    - `StartPublicWhitelist`
    - `StopPublicWhitelist`
    - `StartVpecp`
    - `StopVpecp`
    - `ShowVpcepConnection`
    - `UpdateVpcepConnection`
    - `UpdateVpcepWhitelist`
    - `UpdateYmls`
    - `ListYmlsJob`
    - `ListYmls`
    - `ListClustersDetails`
    - `CreateCluster`
    - `DeleteCluster`
    - `RestartCluster`
    - `UpdateExtendCluster`
    - `UpdateExtendInstanceStorage`
    - `ListFlavors`
    - `ShowClusterTag`
    - `CreateClustersTags`
    - `ListClustersTags`
    - `UpdateBatchClustersTags`
    - `DeleteClustersTags`
    - `ShowIkThesaurus`
    - `CreateLoadIkThesaurus`
    - `DeleteIkThesaurus`
    - `StartAutoSetting`
    - `UpdateSnapshotSetting`
    - `ShowAutoCreatePolicy`
    - `CreateAutoCreatePolicy`
    - `CreateSnapshot`
    - `ListSnapshots`
    - `StopSnapshot`
    - `RestoreSnapshot`
    - `DeleteSnapshot`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - Support the interfaces `CreateOnlineMigrationTask`, `SetOnlineMigrationTask`, `BatchStopMigrationTasks`, `StopMigrationTaskSync`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DevStar

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowApplicationDependentResources`:
    - Add the request parameters `X-Auth-Token`, `limit`, `offset`
    - Add the response parameter `count`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateFunction`:
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the request parameter `runtime`
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ListFunctions`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `UpdateFunctionCode`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ShowFunctionCode`
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the request parameter `runtime`
    - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ShowFunctionConfig`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `CreateFunctionVersion`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ListFunctionVersions`
  - Add the enum values `Go1.x` to the request parameter `runtime` to the interface `CreateDependency`
  - Add the enum values `Go1.x` to the request parameter `runtime` to the interface `UpdateDependency`
  - Add the enum values `Java 8`, `Node.js 6.10`, `Node.js 8.10`, `Node.js 10.16`, `Node.js 12.13`, `Python 2.7`, `Python 3.6`, `Go 1.8`, `Go 1.x`, `PHP 7.3`, Remove the enum values `Python2.7`, `Python3.6`, `Go1.8`, `Java8`, `Node.js6.10`, `Node.js8.10`, `Custom`, `PHP7.3` from the response parameter `runtime` to the interface `ImportFunction`

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `alias` to the interface `ShowGaussMySqlInstanceInfo`
  - Add the response parameter `job_id` to the interface `CreateGaussMySqlBackup`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int32` -> `string` of the response parameter `port` of the interface `ListInstances`

### HuaweiCloud SDK Live

- _Features_
  - Support the interfaces `ListTranscodeTaskCount`, `ListAreaDetail`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `publish_domain` to the interface `ListRecordData`
  - Remove the request parameter `plan_record_time` from the interface `CreateRecordRule`
  - Remove the response parameter `plan_record_time` from the interface `ListRecordRules`
  - Remove the response parameter `plan_record_time` from the interface `ShowRecordRule`
  - Changes of the interface `UpdateRecordRule`:
    - Remove the request parameter `plan_record_time`
    - Remove the response parameter `plan_record_time`
  - Remove the request parameter `on_demand_callback_url` from the interface `CreateRecordCallbackConfig`
  - Remove the response parameter `on_demand_callback_url` from the interface `ListRecordCallbackConfigs`
  - Remove the response parameter `on_demand_callback_url` from the interface `ShowRecordCallbackConfig`
  - Remove the request parameter `on_demand_callback_url` from the interface `UpdateRecordCallbackConfig`
  
### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `restore_all_database` to the interface `RestoreToExistingInstance`
  - Remove the request parameter `is_open_recycle_policy` from the interface `StartRecyclePolicy`

### HuaweiCloud SDK SIS

- _Features_
  - Support the `Voice Interaction Service`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.76 2022-01-25

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `content_type` to the interface `CreateApiV2`
  - Add the response parameters `publish_time`, `roma_app_name`, `ld_api_id`, `api_group_info`, `content_type` to the interface `ShowDetailsOfApiV2`
  - Changes of the interface `UpdateApiV2`:
    - Add the request parameter `content_type`
    - Add the response parameters `publish_time`, `roma_app_name`, `ld_api_id`, `api_group_info`, `content_type`
  - Add the response parameter `content_type` to the interface `ListApiRuntimeDefinitionV2`
  - Changes of the interface `ListApiVersionDetailV2`:
    - Add the response parameters `roma_app_name`, `ld_api_id`, `api_group_info`, `content_type`
    - Modify the type `date-time` -> `string` of the response parameter `publish_time`

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowJobs`:
    - Remove the response parameter `simple`
    - The response parameter `name`, `values` changed to required
  - The request parameter `name`, `values` changed to required of the interface `UpdateJob`
  - The request parameter `name`, `values` changed to required of the interface `CreateAndStartRandomClusterJob`
  - The request parameter `name`, `values` changed to required of the interface `CreateJob`
  - Remove the response parameter `config_status` from the interface `ListClusters`

### HuaweiCloud SDK CES

- _Features_
  - Support the interface `ListAlarmHistories`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `DeleteRuleset`, `SetDefaulTemplate`, `ShowTasklog`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `endpoint_id` to the interface `CreateTask`
  - Add the request parameter `custom_attributes` to the interface `CreateRuleset`
  - Changes of the interface `ListTemplateRules`:
    - Add the request parameter `tags`
    - Add the response parameter `rule_config_list`

### HuaweiCloud SDK DevStar

- _Features_
  - Support the interfaces `ShowRepositoryByCloudIde`, `ListTemplates`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IAM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `session_user_id` to the interface `CreateLoginToken`

### HuaweiCloud SDK Kafka

- _Features_
  - Support the interface `ListEngineProducts`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `dr_enable` to the interface `ListInstances`
  - Add the response parameter `dr_enable` to the interface `ShowInstance`
  - Changes of the interface `ListProducts`:
    - Add the response parameters `Hourly`, `Monthly`
    - Remove the response parameters `hourly`, `honthly`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `language` from the interface `CreateSqlAlarmRule`
  - Remove the request parameter `language` from the interface `UpdateSqlAlarmRule`
  - Changes of the interface `ListSqlAlarmRules`:
    - Add the response parameters `template_name`, `status`
    - Remove the response parameter `language`
  - Remove the request parameters `language`, `eps_id`, `eps_name`, `is_time_range_relative` from the interface `CreateKeywordsAlarmRule`
  - Changes of the interface `UpdateKeywordsAlarmRule`:
    - Remove the request parameters `language`, `eps_id`, `eps_name`, `is_time_range_relative`
    - Remove the response parameters `eps_id`, `eps_name`, `is_time_range_relative`
  - Changes of the interface `ListKeywordsAlarmRules`:
    - Add the response parameters `template_name`, `status`
    - Remove the response parameters `language`, `eps_id`, `eps_name`, `is_time_range_relative`
  - Changes of the interface `ListAccessConfig`:
    - Add the request parameter `access_config_tag_list`
    - Add the response parameter `access_config_tag`
  - Changes of the interface `CreateAccessConfig`:
    - Add the request parameter `access_config_tag`
    - Add the response parameter `access_config_tag`
  - Changes of the interface `UpdateAccessConfig`:
    - Add the request parameter `access_config_tag`
    - Add the response parameter `access_config_tag`
  - Add the response parameter `access_config_tag` to the interface `DeleteAccessConfig`

### HuaweiCloud SDK RabbitMQ

- _Features_
  - Support the interface `ListEngineProducts`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.75 2022-01-17

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the request body structure of the interface `CreateNodePool` is incorrect
- _Change_
  - None

### HuaweiCloud SDK FunctionGraph

- _Features_
  - Support the interface `ListFunctionAsyncInvocations`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `used` to the interface `ShowGaussMySqlInstanceInfo`
  - Add the request parameter `monitor_switch` to the interface `UpdateInstanceMonitor`
  - Modify the type of the request parameter `period` of the interface `UpdateInstanceMonitor`: `string` -> `int32`
  - Remove the response parameter `mode` from the interface `ShowGaussMySqlProjectQuotas`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `ListApiVersionNew`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `az_desc` to the interface `ListFlavors`

### HuaweiCloud SDK ROMA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `id`, `roles` to the interface `AddUserToApp`
  - Remove the response parameters `apps`, `tasks` from the interface `DownloadAssetArchive`
  - Remove the request parameters `apps`, `tasks` from the interface `ImportAsset`
  - The request parameter `tasks` changed to required of the interface `DeleteAsset`
  - Add the response parameter `access_user` to the interface `ShowMqsInstance`

# 0.0.74 2022-01-10

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `coupon_max_use_quantity`, `coupon_group` to the interface `ListOrderCouponsByOrderId`

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowQuotas`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `CreateCluster`
  - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `ListClusters`
  - Changes of the interface `UpdateCluster`:
    - Add the request parameters `customSan`, `containerNetwork`
    - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain`
  - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `ShowCluster`
  - Add the response parameters `customSan`, `offloadCluster`, `cidrs`, `flavor`, `faultDomain` to the interface `DeleteCluster`
  - Add the request parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `CreateNode`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `ListNodes`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `UpdateNode`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `ShowNode`
  - Add the response parameters `faultDomain`, `offloadNode`, `offloadNode` to the interface `DeleteNode`
  - Add the request parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `CreateNodePool`
  - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `ListNodePools`
  - Changes of the interface `UpdateNodePool`:
    - Add the request parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode`
    - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode`
  - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `ShowNodePool`
  - Add the response parameters `podSecurityGroups`, `faultDomain`, `offloadNode`, `offloadNode` to the interface `DeleteNodePool`

### HuaweiCloud SDK CDN

- _Features_
  - Support the interface `UpdateDomainFullConfig`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudIDE

- _Features_
  - Support the interface `ListStacks`
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListStacksByTag`

### HuaweiCloud SDK Cloudtest

- _Features_
  - Support the interface `ShowPlanList`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowPlans`:
    - Modify the type `int64` -> `int32` of the request parameter `limit`
    - Modify the type `int64` -> `int32` of the request parameter `offset`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ecs_tenant_private_ip` to the interface `ListMigrationTask`
  - Add the response parameter `ecs_tenant_private_ip` to the interface `ShowMigrationTask`
  - Add the response parameter `ecs_tenant_private_ip` to the interface `StopMigrationTask`

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `lb_ip_address` from the interface `ListInstances`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `delete_on_termination` to the interface `CreateServers`

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `UpdateFunctionConfig`:
    - Add the request parameter `is_stateful_function`
    - Add the response parameter `is_stateful_function`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `lb_ip_address` to the interface `ListInstances`

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `GaussDB(for openGauss)`, Remove the enum values `GaussDB(openGauss)` from the request parameter `type` to the interface `CreateInstance`
  - Remove the response parameter `related_instance` from the interface `ListInstances`

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The response parameter `active_at` changed to not required of the interface `GlanceListImages`
  - The response parameter `active_at` changed to not required of the interface `GlanceShowImage`
  - The response parameter `active_at` changed to not required of the interface `GlanceUpdateImage`

### HuaweiCloud SDK IoTAnalytics

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `site_id` changed to required of the interface `CreateDatasource`
  - The request parameter `site_id` changed to required of the interface `UpdateDataSource`

### HuaweiCloud SDK KPS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListKeypairs`:
    - Add the request parameters `limit`, `marker`
    - Add the response parameter `page_info`
  - Add the request parameters `limit`, `offset` to the interface `ListFailedTask`
  - Add the request parameters `limit`, `offset` to the interface `ListRunningTask`

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `StartInstanceSingleToHaAction`:
    - Add the request parameter `ad_domain_info`
    - Remove the request parameter `password`

### HuaweiCloud SDK ROMA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `connect_dn` from the interface `ShowMqsInstance`
  - Remove the response parameter `policies` from the interface `ListMqsInstanceTopics`
  - Add the response parameter `roma_app_type` to the interface `ShowDetailsOfAppV2`

### HuaweiCloud SDK VPCEP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `pool_id` to the interface `CreateEndpointService`
  - Add the response parameter `domain_id` to the interface `ListEndpointService`
  - Add the response parameter `pool_id` to the interface `UpdateEndpointService`
  - Add the response parameters `enable_status`, `specification_name` to the interface `ListEndpointInfoDetails`

### HuaweiCloud SDK VSS

- _Features_
  - Support the interfaces `AuthorizeDomains`, `ShowTasks`, `CreateTasks`, `ShowResults`
- _Bug Fix_
  - None
- _Change_
  - Add the enum values `skip` to the response parameter `auth_status` to the interface `ListDomains`

# 0.0.73 2021-12-25

### HuaweiCloud SDK CDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `id` to the interface `CreateLink`
  - Add the response parameters `endpointDomainName`, `isScheduleBootOff`, `failedReasons`, `components`, `createFrom`, `resourceId`, `flavorType`, `workSpaceId`, `trial` to the interface `ShowClusterDetail`
  - Add the request parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type` to the interface `UpdateJob`
  - Add the response parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type` to the interface `ShowJobs`
  - Changes of the interface `CreateAndStartRandomClusterJob`:
    - Add the request parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type`
    - Add the response parameter `submissions`
    - Remove the response parameters `name`, `validation-result`
  - Add the request parameters `is_incre_job`, `flag`, `files_read`, `external_id`, `type`, `execute_start_date`, `delete_rows`, `enabled`, `bytes_written`, `id`, `is_use_sql`, `update_rows`, `group_name`, `bytes_read`, `execute_update_date`, `write_rows`, `files_writte`, `is_incrementing`, `execute_create_date`, `id`, `type`, `id`, `type`, `id`, `type` to the interface `CreateJob`
  - Add the response parameter `execute-date` to the interface `StartJob`
  - Add the request parameter `id` to the interface `UpdateLink`
  - Add the response parameter `id` to the interface `ShowLink`
  - Changes of the interface `ListClusters`:
    - Add the response parameters `bakExpectedStartTime`, `bakKeepDay`, `createFrom`, `resourceId`, `flavorType`, `workSpaceId`, `trial`, `components`
    - Remove the response parameter `version`

### HuaweiCloud SDK CloudBuild

- _Features_
  - Support the interface `ShowHistoryDetails`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudTable

- _Features_
  - Support the service `CloudTable`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Cloudtest

- _Features_
  - Support the interfaces `ShowPlanJournals`, `ShowIssuesByPlanId`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `CheckParameters`, `ListTaskParameter`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `rule_desc` to the interface `ListRules`
  - Add the response parameters `is_devcloud_project_default`, `is_default_template` to the interface `ListRulesets`

### HuaweiCloud SDK CodeHub

- _Features_
  - Support the service `CodeHub`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `backup_format` to the interface `CopyInstance`

### HuaweiCloud SDK DevStar

- _Features_
  - Support the following interfaces:
    - `ShowApplicationV3`
    - `UpdateApplication`
    - `ShowApplicationDependentResources`
    - `DeleteApplicationV4`
    - `ShowApplicationResDeleteStatus`
    - `ListApplicationsV6`
    - `ShowDeploymentJobs`
    - `CreateDeploymentJobs`
    - `ShowPipelineLastStatusV2`
    - `ListPipelineTemplates`
    - `StartPipeline`
    - `ListProjectsV4`
    - `ShowRepositoryStatisticalDataV2`
    - `CheckRepositoryDuplicateName`
    - `ShowApplicationReleaseRepositories`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `dependents` to the interface `ShowTemplateV3`
  - Add the response parameters `dependents`, `dependent_services` to the interface `ListTemplatesV2`
  - Add the response parameter `show_type` to the interface `ShowJobDetail`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `updated_at` from the interface `NovaListServerActions`

### HuaweiCloud SDK IEF

- _Features_
  - Support the service `Intelligent EdgeFabric`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IoTAnalytics

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateDatasource`:
    - Add the request parameter `site_id`
    - Remove the request parameter `instance_id`
  - Changes of the interface `ShowAllDataSource`:
    - Add the response parameter `site_id`
    - Remove the response parameter `instance_id`
  - Changes of the interface `UpdateDataSource`:
    - Add the request parameter `site_id`
    - Add the response parameter `site_id`
    - Remove the request parameter `instance_id`
    - Remove the response parameter `instance_id`
  - Changes of the interface `ShowDataSource`:
    - Add the response parameter `site_id`
    - Remove the response parameter `instance_id`

### HuaweiCloud SDK Kafka

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `partition_num` changed to not required of the interface `CreatePostPaidInstance`
  - Add the response parameters `result`, `instance_id` to the interface `RestartManager`
  - Changes of the interface `ListProducts`:
    - Add the response parameters `hourly`, `honthly`
    - Remove the response parameters `Hourly`, `Monthly`

### HuaweiCloud SDK KPS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int64` -> `string` of the response parameter `task_time` of the interface `ListFailedTask`
  - Modify the type `int64` -> `string` of the response parameter `task_time` of the interface `ListRunningTask`

### HuaweiCloud SDK LTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `offset`, `limit` to the interface `ListCharts`
  - Add the request parameters `offset`, `limit` to the interface `ListNotificationTemplates`

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `CreateTranscodingTask`:
    - Add the request parameter `auto_volume_setting`
    - Add the enum values `original` to the request parameter `volume`
  - Add the response parameter `av_parameters` to the interface `ListTranscodingTask`
  - Add the request parameter `template_id` to the interface `CreateWatermarkTemplate`
  - Add the response parameter `template_id` to the interface `ListWatermarkTemplate`
  - Add the request parameter `template_id` to the interface `UpdateWatermarkTemplate`

### HuaweiCloud SDK MRS

- _Features_
  - Support the following interfaces:
    - `CreateScalingPolicy`
    - `ShowClusterDetails`
    - `UpdateClusterScaling`
    - `ListHosts`
    - `CreateAndExecuteJob`
    - `ListExecuteJob`
    - `ShowJobExes`
    - `DeleteJobExecution`
    - `CreateCluster`
    - `ShowAgencyMapping`
    - `UpdateAgencyMapping`
    - `ShowJobExeListNew`
    - `CreateExecuteJob`
    - `ShowSingleJobExe`
    - `StopJob`
    - `ShowSqlResultWithJob`
    - `BatchDeleteJobs`
    - `ExecuteSql`
    - `ShowSqlResult`
    - `CancelSql`
    - `ShowHdfsFileList`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `group_type`, `success_record_error_reason`, `skip_record_error_reason`, `save_prefix` to the interface `ListTasks`
  - Add the response parameters `group_type`, `success_record_error_reason`, `skip_record_error_reason`, `save_prefix` to the interface `ShowTask`

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ListApiVersion`, `ShowApiVersion`, `SearchQueryScaleComputeFlavors`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ROMA

- _Features_
  - Support the service `ROMA Connect`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SA

- _Features_
  - Support the service `Situation Awareness`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.72 2021-12-17

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowVersion`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudIDE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ShowInstance`:
    - Add the response parameters `bundle_url`, `visitor_id`, `visitor_name`, `visitor_domain_name`
    - Remove the response parameters `action_list`, `role`, `role_id`, `sub_org`
  - Changes of the interface `ListOrgInstances`:
    - Add the response parameters `visitor_id`, `visitor_name`, `visitor_domain_name`
    - Remove the response parameters `action_list`, `role`, `role_id`, `sub_org`
  - Changes of the interface `ListInstances`:
    - Add the response parameters `visitor_id`, `visitor_name`, `visitor_domain_name`
    - Remove the response parameters `action_list`, `role`, `role_id`, `sub_org`

### HuaweiCloud SDK CloudRTC

- _Features_
  - Support the interfaces `ListRtcAbnormalEvents`, `ListRtcAbnormalEventDimension`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListEvents`:
    - Modify the type `string` -> `int32` of the response parameter `event_count`
    - Modify the type `string` -> `int64` of the response parameter `latest_occur_time`
  - Modify the type `string` -> `double` of the response parameter `variance` of the interface `BatchListMetricData`
  - Modify the type `string` -> `int32` of the response parameter `type_statistics` of the interface `ListResourceGroup`
  - Changes of the interface `ListEventDetail`:
    - Modify the type `string` -> `array` of the response parameter `event_users`
    - Modify the type `string` -> `array` of the response parameter `event_sources`

### HuaweiCloud SDK IoTAnalytics

- _Features_
  - Support the service `IoTAnalytics`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `input`, `edit_settings` to the interface `CreateEditingJob`
  - Add the response parameters `input`, `edit_settings` to the interface `ListEditingJob`

### HuaweiCloud SDK OCR

- _Features_
  - Support the following interfaces:
    - `RecognizeThailandIdcard`
    - `RecognizeMyanmarIdcard`
    - `RecognizeMyanmarDriverLicense`
    - `RecognizeChileIdCard`
    - `RecognizeThailandLicensePlate`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.71 2021-12-10

### HuaweiCloud SDK AOM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Update the enum values to `[AND, OR, NOT]` of the request parameter `relation` of the interface `ListEvents`

### HuaweiCloud SDK APM

- _Features_
  - Support the interfaces `ShowMasterAddress`, `ListAkSk`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK AS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `allowed_address_pairs` to the interface `ListScalingGroups`
  - Add the request parameter `allowed_address_pairs` to the interface `CreateScalingGroup`
  - Add the request parameter `allowed_address_pairs` to the interface `UpdateScalingGroup`
  - Add the response parameter `allowed_address_pairs` to the interface `ShowScalingGroup`

### HuaweiCloud SDK BCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type `int32` -> `int64` of the response parameter `node_cnt` of the interface `ShowBlockchainDetail`
  - Add the request parameters `limit`, `offset` to the interface `ShowBlockchainFlavors`

### HuaweiCloud SDK CampusGo

- _Features_
  - Support the service `CampusGo`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CGS

- _Features_
  - Support the `Container Guard Service`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudIDE

- _Features_
  - Support the interfaces `ShowExtensionAuthorization`, `CreateExtensionAuthorization`, `CheckInstanceAccess`, `UpdateInstanceActivity`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CloudRTC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `stream_id` to the interface `ListRtcClientQosDetails`

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support the interfaces `ShowTaskCmetrics`, `ListTemplateRules`, `ListTaskRuleset`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `is_access`, `trigger_type` to the interface `ShowTaskDetail`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `readonly_domain_name` to the interface `ListInstances`
  - Add the response parameter `readonly_domain_name` to the interface `ShowInstance`

### HuaweiCloud SDK DDM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `time_zone` to the interface `CreateInstance`
  - Add the response parameters `instanceId`, `instanceName`, `jobId` to the interface `RestartInstance`

### HuaweiCloud SDK DSC

- _Features_
  - Support the interface `ShowOpenApiCalledRecords`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `readonly_password` to the interface `CreateDocWatermark`
  - Changes of the interface `ShowScanJobs`:
    - Add the request parameter `offset`
    - Remove the request parameter `page`
  - Changes of the interface `ShowScanJobResults`:
    - Add the request parameter `offset`
    - Remove the request parameter `page`

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameters `landmark`, `gender`, `yaw_angle`, `roll_angle`, `pitch_angle`, `headpose`, `smile`, `skin`, `ethnic` from the interface `DetectFaceByFile`
  - Remove the response parameters `landmark`, `gender`, `yaw_angle`, `roll_angle`, `pitch_angle`, `headpose`, `smile`, `skin`, `ethnic` from the interface `DetectFaceByUrl`
  - Remove the response parameters `landmark`, `gender`, `yaw_angle`, `roll_angle`, `pitch_angle`, `headpose`, `smile`, `skin`, `ethnic` from the interface `DetectFaceByBase64`

### HuaweiCloud SDK GaussDB

- _Features_
  - Support the following interfaces:
    - `ListInstanceTags`
    - `ListProjectTags`
    - `BatchTagAction`
    - `ShowInstanceMonitorExtend`
    - `UpdateInstanceMonitor`
- _Bug Fix_
  - None
- _Change_
  - Changes of the interface `ListGaussMySqlInstances`:
    - Add the request parameters `private_ip`, `tags`
    - Add the response parameter `tags`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `region` changed to required of the interface `ListFlavors`

### HuaweiCloud SDK GES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `CheckMetadata`

### HuaweiCloud SDK HiLens

- _Features_
  - Support the service `Huawei HiLens`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK KMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `encrypted_privatekey` to the interface `ImportKeyMaterial`

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `default_record_config` changed to required of the interface `CreateRecordRule`

### HuaweiCloud SDK LTS

- _Features_
  - Support the following interfaces:
    - `ListLogHistogram`
    - `ListHost`
    - `ListHostGroup`
    - `UpdateHostGroup`
    - `CreateHostGroup`
    - `DeleteHostGroup`
    - `ListAccessConfig`
    - `UpdateAccessConfig`
    - `CreateAccessConfig`
    - `DeleteAccessConfig`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `RestoreExistInstance`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SWR

- _Features_
  - Support the interface `ListQuotas`
- _Bug Fix_
  - None
- _Change_
  - Modify the name `UpdateNamespaceAuthReq` -> `UpdateNamespaceAuthRequestBody` of the request body of the interface `UpdateNamespaceAuth`

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - The request parameter `to`, `from` changed to required of the interface `ListStatistics`
  - The request parameter `from`, `to` changed to required of the interface `ListQpsTimeline`
  - The request parameter `from`, `to` changed to required of the interface `ListBandwidthTimeline`
  - The request parameter `from`, `to` changed to required of the interface `ListTopAbnormal`
  - Add the response parameter `cookie` to the interface `ShowEvent`
  - Changes of the interface `CreatePremiumHost`:
    - Add the request parameter `Content-Type`
    - Modify the type `string` -> `enum` of the request parameter `type`
  - Add the request parameter `Content-Type` to the interface `ListPremiumHost`
  - Changes of the interface `UpdatePremiumHost`:
    - Add the request parameter `Content-Type`
    - Modify the type `string` -> `enum` of the response parameter `type`
  - Add the request parameter `Content-Type` to the interface `DeletePremiumHost`
  - Changes of the interface `ShowPremiumHost`:
    - Add the request parameter `Content-Type`
    - Modify the type `string` -> `enum` of the response parameter `type`
  - Add the request parameter `Content-Type` to the interface `UpdatePremiumHostProtectStatus`

# 0.0.70 2021-11-29

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `roma_app_type` to the interface `ResettingAppSecretV2`.

### HuaweiCloud SDK BCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameter `values` of the interface `ListEntityMetric`: `object` -> `array`

### HuaweiCloud SDK CBR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `provider_id` to the interface `ListBackups`.

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `cidrs` to the interface `ShowCluster`.

### HuaweiCloud SDK DBSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the response parameter of the interfaces `SwitchAgent` and `SwitchRiskRule`: `status` -> `result`

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `port` to the interface `UpdateInstance`.

### HuaweiCloud SDK DSC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `start_time` to the interface `ShowScanJobs`.

### HuaweiCloud SDK EVS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `CinderExportToImage`.

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `os:scheduler_hints` to the interfaces `NovaShowServer` and `NovaListServersDetails`.

### HuaweiCloud SDK ELB

- _Features_
  - Support the following interfaces:
    - `ListApiVersions`
    - `ListSecurityPolicies`
    - `CreateSecurityPolicy`
    - `ShowSecurityPolicy`
    - `UpdateSecurityPolicy`
    - `DeleteSecurityPolicy`
    - `ListSystemSecurityPolicies`
    - `ListQuotaDetails`
    - `ChangeLoadbalancerChargeMode`
    - `BatchUpdatePoliciesPriority`
    - `UpdateIpList`
    - `BatchDeleteIpList`
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ShowQuotaDefaults`.
  - Add the response parameters `flavor_sold_out` and `lcu` to the interfaces `ListFlavors` and `ShowFlavor`, and remove the response parameter `availability_zone_ids`.
  - Add the response parameter `members_per_pool` to the interface `ShowQuota`.
  - Add the request parameters `enc_certificate` and `enc_private_key` to the interfaces `CreateCertificate` and `UpdateCertificate`.
  - Add the response parameters `enc_certificate` and `enc_private_key` to the interfaces `ListCertificates` and `ShowCertificate`.
  - Add the request parameters `prepaid_options`, `autoscaling` and `id` to the interface `CreateLoadBalancer`.
  - Add the request parameters `elb_virsubnet_type` and `autoscaling` and the response parameters `autoscaling` and `ip_version` to the interface `ListLoadBalancers`.
  - Add the request parameters `elb_virsubnet_type` and `autoscaling` and the response parameters `loadbalancer_id`, `order_id`, `autoscaling` and `ip_version` to the interface `UpdateLoadBalancer`.
  - Add the response parameters `autoscaling` and `ip_version` to the interface `ShowLoadBalancer`.
  - Add the response parameters `id`, `type` and `provisioning_status` to the interface `ShowLoadBalancerStatus`.
  - Add the request parameters `security_policy_id` and `enhance_l7policy_enable` to the interface `CreateListener`.
  - Add the request parameters `enhance_l7policy_enable` and `member_instance_id` and the response parameters `security_policy_id`, `transparent_client_ip_enable` and `enhance_l7policy_enable` to the interface `ListListeners`.
  - Add the request parameters `enhance_l7policy_enable` and `member_instance_id` and the response parameters `security_policy_id`, `transparent_client_ip_enable` and `enhance_l7policy_enable` to the interface `UpdateListener`.
  - Add the response parameters `security_policy_id`, `transparent_client_ip_enable` and `enhance_l7policy_enable` to the interface `ShowListener`.
  - Add the request parameters `listener_id` and `member_instance_id` to the interface `ListPools`.
  - Add the request parameters `ip_version` and `member_type` and the response parameters `member_type` and `instance_id` to the interface `ListMembers`.
  - Add the response parameters `member_type` and `instance_id` to the interfaces `UpdateMember`, `ShowMember` and `ListAllMembers`.
  - Add the request parameters `priority`, `redirect_url_config`, `fixed_response_config` and `conditions` to the interface `CreateL7Policy`.
  - Add the request parameter `priority` and the response parameters `redirect_url_config` and `fixed_response_config` to the interface `ListL7Policies`.
  - Add the request parameters `priority`, `redirect_url_config`, `fixed_response_config` and `rules` and the response parameters `member_type` and `instance_id` to the interface `UpdateL7Policy`.
  - Add the response parameters `redirect_url_config` and `fixed_response_config` to the interface `ShowL7Policy`.
  - Add the request parameter `conditions` to the interface `CreateL7Rule`.
  - Add the response parameter `conditions` to the interfaces `ListL7Rules` and `ShowL7Rule`.
  - Add the request and response parameter `conditions` to the interface `UpdateL7Rule`.

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ListVersionAliases` to the interface `ListVersionAliases`.
  - Update the request parameter `name` to required of the interfaces  `CreateDependency` and `UpdateDependency`.
  - Update the request parameters `name` and `content` to required of the interface `CreateEvent`.

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `enable_force_switch` to the interface `CreateInstance`.

### HuaweiCloud SDK GES

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `graphSizeTypeIndex` of the interface `CreateGraph`: `integer` -> `string`

### HuaweiCloud SDK KMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Update the request parameter `key_alias` to required of the interface `CreateKey`.
  - Update the request parameter `key_id` to required of the interfaces `EnableKey`, `CancelKeyDeletion`, `ListKeys`, `ListKeyDetail`, `ShowPublicKey`, `ListGrants`, `DeleteImportedKeyMaterial`, `EnableKeyRotation`, `DisableKeyRotation` and `ShowKeyRotationStatus`.
  - Update the request parameters `key_id` and `pending_days` to required of the interface `DeleteKey`.
  - Add the request parameter `enterprise_project_id` to the interface `ListKeys`.
  - Update the request parameter `random_data_length` to required of the interface `CreateRandom`.
  - Update the request parameters `key_id` and `datakey_length` to required of the interfaces `CreateDatakey` and `CreateDatakeyWithoutPlaintext`.
  - Update the request parameters `key_id`, `plain_text` and `datakey_plain_length` to required of the interface `EncryptDatakey`.
  - Update the request parameters `key_id`, `cipher_text` and `datakey_cipher_length` to required of the interface `DecryptDatakey`.
  - Update the request parameters `key_id` and `key_alias` to required of the interface `UpdateKeyAlias`.
  - Update the request parameters `key_id` and `key_description` to required of the interface `UpdateKeyDescription`.
  - Update the request parameters `key_id`, `grantee_principal` and `operations` to required of the interface `CreateGrant`.
  - Update the request parameters `key_id` and `grant_id` to required of the interfaces `CancelGrant` and `CancelSelfGrant`.
  - Update the request parameters `key_id` and `plain_text` to required of the interface `EncryptData`.
  - Update the request parameter `cipher_text` to required of the interface `DecryptData`.
  - Update the request parameters `key_id` and `wrapping_algorithm` to required of the interface `CreateParametersForImport`.
  - Update the request parameters `key_id`, `import_token` and `encrypted_key_material` to required of the interface `ImportKeyMaterial`.
  - Update the request parameters `key_id` and `rotation_interval` to required of the interface `UpdateKeyRotationInterval`.
  - Add the request parameter `sequence` to the interface `CreateKmsTag`.

### HuaweiCloud SDK RDS

- _Features_
  - Support the following interfaces:
    - `ChangeProxyScale`
    - `SearchQueryScaleFlavors`
    - `ShowInformationAboutDatabaseProxy`
    - `StartDatabaseProxy`
    - `StopDatabaseProxy`
    - `UpdateReadWeight`
    - `ChangeTheDelayThreshold`
    - `ShowDrReplicaStatus`
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameter `size` of the interface `ListPostgresqlDatabases`: `int32` -> `int64`

### HuaweiCloud SDK SMS

- _Features_
  - None
- _Bug Fix_
  - Fix the problem that the enumeration value contains the Chinese description and causes the parameter error.
- _Change_
  - None

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `items` to the interface `ListIgnoreRule`.
  - Add the request parameter `attacks` to the interface `ListEvent`.
  - Add the response parameter `host_id` to the interface `ShowEvent`.
  - Add the response parameter `certificatename` to the interface `UpdateHost`.

# 0.0.69 2021-11-25

### HuaweiCloud SDK AOM

- _Features_
  - Support the service `Application Operations Management`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK APIG

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `CheckBackendConnectivity`.

### HuaweiCloud SDK APM

- _Features_
  - Support the service `Application Performance Management`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BCS

- _Features_
  - Support the service `Application Performance Management`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `statistic_type` and response parameter `bill_date` to the interface `ListCustomerselfResourceRecordDetails`.

### HuaweiCloud SDK CBH

- _Features_
  - Support the service `Cloud Bastion Host`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `platformVersion` to the interface `ShowCluster`.
  - Add the enumeration values `RollingBack` and `RollbackFailed` to the request parameter `status` of the interface `ListClusters`.

### HuaweiCloud SDK CloudRTC

- _Features_
  - Support the following interfaces:
    - `ListRtcRealtimeScaleDimension`
    - `ListRtcRealtimeQuality`
    - `ListRtcRealtimeNetwork`
    - `ListRtcHistoryUsage`
    - `ListRtcHistoryScale`
    - `ListRtcHistoryQuality`
    - `ListRtcClientQosDetails`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCheck

- _Features_
  - Support CodeCheck service.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DAS

- _Features_
  - None
- _Bug Fix_
  - Correct the enumeration value of the request parameter `X-Language` of some interfaces.
- _Change_
  - None

### HuaweiCloud SDK DBSS

- _Features_
  - Support the `Database Security Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `tags` to the interfaces `CreateInstance` and `ListInstances`.

### HuaweiCloud SDK DeH

- _Features_
  - Support the service `Dedicated Host`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request and response parameter `alias` to the interfaces `CreatePrePaidPublicip` and `CreatePublicip`.
  - Add the response parameter `alias` to the interfaces `ShowPublicip` and `UpdatePublicip`.

### HuaweiCloud SDK GES

- _Features_
  - Support the interfaces `ResizeGraph`, `ExpandGraph` and `UploadFromObs`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK LTS

- _Features_
  - Support the following interfaces:
    - `CreateTransfer`
    - `DeleteTransfer`
    - `UpdateTransfer`
    - `ListTransfers`
    - `ListLogStreams`
    - `RegisterDmsKafkaInstance`
    - `CreateNotificationTemplate`
    - `UpdateNotificationTemplate`
    - `ListNotificationTemplates`
    - `DeleteNotificationTemplate`
    - `ShowNotificationTemplate`
    - `ListNotificationTemplate`
    - `UpdateAlarmRuleStatus`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK KMS

- _Features_
  - Support the interfaces of version `V2`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK NLP

- _Features_
  - Support the service `Natural Language Processing`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the response parameter `extracted_data` of the interface `RecognizeHandwriting`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `sort` to the interface `ListSlowlogStatistics`.

### HuaweiCloud SDK SFSTurbo

- _Features_
  - Support the service `Scalable File Service - Turbo`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK TMS

- _Features_
  - Support the interface `ShowTagQuota`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VPCEP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameters `created_at` and `updated_at` of the interfaces `ListServicePublicDetails`, `ListServiceDetails`, `ListServiceConnections`, `AcceptOrRejectEndpoint`, `ListEndpoints`, `UpdateEndpointWhite` and `ListEndpointInfoDetails`.
  - Update the request parameters `vpc_id` and `port_id` to required of the interface `CreateEndpointService`.
  - Remove the response parameter `error` of the interface `AcceptOrRejectEndpoint`.
  - Modify the type of the response parameter `updated` of the interfaces `ListVersionDetails` and `ListSpecifiedVersionDetails`: `datetime` -> `string`
  - Update the request parameter `action` to required of the interfaces `ListResourceInstances` and `BatchAddOrRemoveResourceInstance`.

### HuaweiCloud SDK WAF

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `from` and `to` to the interface `ListEvent`.
  - Add the request parameter `name` to the interface `ListWhiteblackipRule`.

# 0.0.68 2021-11-12

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of parsing errors when the response body type is `array` or `map`.
  - Fix the problem of automatic escaping of request body.
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `customSan` to the interface `UpdateNode`.

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `force_redirect_https` to the interface `UpdateHttpsInfo`.

### HuaweiCloud SDK CloudDeploy

- _Features_
  - Support the service `CloudDeploy`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the response parameter of the interface `ListAuditLogs`: `total_count` -> `total_record`.

### HuaweiCloud SDK DSC

- _Features_
  - Support the interfaces `ShowScanJobs` and `ShowScanJobResults`.
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `marked_file_password` to the interface `CreateDocWatermark`.

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `offset` and `limit`.
  - Add the response parameters `id` and `spec_code`.

### HuaweiCloud SDK GES

- _Features_
  - Support `Graph Engine Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameters `sim_card_id`, `partner`, `package_type` and `sim_type` from the interface `ListProPricePlans`.
  - Remove the request parameter `tag_id` from the interface `ListSimCards`.
  - Remove the request parameter `sim_price_plan_id` from the interface `ListSimPricePlans`.
  - Remove the request parameter `price_plan_list` from the interfaces `StopSimCard` and `ResetSimCard`.
  - Remove the response parameters `sn`, `supply_code`, `bundle_id` and `test_type` from the interfaces `ShowSimCard` and `ListSimCards`.
  - Remove the response parameter `sim_price_plan_list` from the interface `StopSimCard`.
  - Remove the response parameter `order_id` from the interface `ListSimPools`.
  - Remove the response parameters `partner` and `partner_pid` from the interface `ListSimPricePlans`.

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `results` to the interface `ShowJob`.

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `country` and `protocol` to the interface `ListBandwidthDetail` and `ListUsersOfStream`.
  - Add the request parameter `protocol` to the interface `ListDomainTrafficDetail`, `ListDomainBandwidthPeak` and `ListDomainTrafficSummary`.
  - Add the request parameter `stream` to the interface `ListTranscodeData`.
  - Add the request parameters `stream`, `start_time` and `end_time` to the interface `ListHistoryStreams`.

### HuaweiCloud SDK LTS

- _Features_
  - Support the following interfaces:
    - `ShowStructTemplate`
    - `CreateStructTemplate`
    - `UpdateStructTemplate`
    - `DeleteStructTemplate`
    - `ShowAomMappingRules`
    - `CreateAomMappingRules`
    - `UpdateAomMappingRules`
    - `DeleteAomMappingRules`
    - `ShowAomMappingRule`
    - `ListNotificationTopics`
    - `CreateSqlAlarmRule`
    - `UpdateSqlAlarmRule`
    - `ListSqlAlarmRules`
    - `DeleteSqlAlarmRule`
    - `CreateKeywordsAlarmRule`
    - `UpdateKeywordsAlarmRule`
    - `ListKeywordsAlarmRules`
    - `DeleteKeywordsAlarmRule`
    - `ListActiveOrHistoryAlarms`
    - `DeleteActiveAlarms`
    - `ListCharts`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the request parameter of the interface `CreateEditingJob`: `const` -> `consts`
  - Remove the request parameter `index`, and add the request parameters `overlay_input`, `const` and `mix` to the interface `CreateEditingJob`.
  - Add the response parameter `output` to the interface `ListEditingJob`.
  - Add the response parameters `hls_init_count` and `hls_init_interval`, add the optional value `EFFICIENCY` to the request parameter `video_enhance` of the interface `CreateTranscodingTask`.

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `job_id`, `user` and `queue` to the interface(V2) `GetJobExeListNew`.

### HuaweiCloud SDK OCR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `confidence` to the interface `RecognizeGeneralTable`.

### HuaweiCloud SDK ProjectMan

- _Features_
  - Support the interface `CreateSystemIssueV4`.
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `sequence` to the interface `ListIrs`.
  - Add the request parameter `status_id` to the interface `BatchUpdateIrs`.

### HuaweiCloud SDK RDS

- _Features_
  - Support the interfaces `ListSlowLogsNew`, `ListErrorLogsNew` and `UpgradeDbVersion`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VAS

- _Features_
  - Support `Video Analysis Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK VPC

- _Features_
  - Support interfaces(V3): `AddVpcExtendCidr`、`RemoveVpcExtendCidr`、`ListVpcs`、`ShowVpc`
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.67 2021-10-25

### HuaweiCloud SDK BCS

- _Features_
  - Support the interface `ShowBlockchainFlavors`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `account_manager_id` and `account_manager_name` to the interface `ListIndirectPartners`.

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `create_time` and response parameter `task_type` to the interface `ShowHistoryTaskDetails`.
  - Remove the request parameter `create_time` from the interface `ShowHistoryTasks`.

### HuaweiCloud SDK DNS

- _Features_
  - Support the interface `ShowDomainQuota`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `bandwidth_type` to the interface `CreateSharedBandwidth`.

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `single` to the interfaces `AddFacesByFile`, `AddFacesByBase64` and `AddFacesByUrl`.

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of request and response parameters `num` and `size` of the interface `CreateInstance`: `integer` -> `string`.

### HuaweiCloud SDK GSL

- _Features_
  - Support the service `Global SIM Link`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ImageSearch

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the request parameter of the interface `RunSearchPicture`: `isCrop` -> `is_crop`
  - Add the request parameter `box` to the interface `RunSearchPicture`.

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `current_task`, `image_name` and `process_percent` to the interface `ShowJob`.

### HuaweiCloud SDK IoTDA

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `status` to the interface `ListDevices`.
  - Add the request parameter `file_path` to the interface `CreateRuleAction`.

### HuaweiCloud SDK OCR

- _Features_
  - Support the interfaces `RecognizeInsurancePolicy`, `RecognizeFinancialStatement` and `RecognizeQualificationCertificate`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `database_name` to the interface `ShowBackupDownloadLink`.
  - Add the response parameter `max_iops` and `expiration_time` to the interface `ListInstances`.

### HuaweiCloud SDK SDRS

- _Features_
  - Support the `Storage Disaster Recovery Service`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.66 2021-10-19

### HuaweiCloud SDK DCS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `backup_id` to the interface `ShowMigrationTask`.
  - Add the following response parameters to the interface `ListMigrationTask`:
    - `source_instance_name`
    - `source_instance_id`
    - `target_instance_addrs`
    - `target_instance_id`
  - Modify the type of the response parameter `total_usec_sum` of the interface `ListDiagnosisTasks`: `integer` -> `double`

### HuaweiCloud SDK EIP

- _Features_
  - Support the following interfaces:
    - `ListCommonPools`
    - `ListPublicBorderGroups`
    - `ListPublicipPool`
    - `ShowPublicipPool`
    - `ListShareBandwidthTypes`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `allow_share_bandwidth_type_any` to the interface `ListPublicips`.
  - Modify the type of the request parameter `limit` of the interface `NeutronListFloatingIps`.
  - Modify the name of the request body of the interface `NeutronUpdateFloatingIp`: `floatingip` -> `NeutronUpdateFloatingIpRequestBody`
  - Add the response parameter `allow_share_bandwidth_types` to the interface `ShowPublicip`.

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `hls_init_count` and `hls_init_interval` to the interface `CreateTranscodingTask`.

### HuaweiCloud SDK VPCEP

- _Features_
  - Support the service `VPC Endpoint`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.65 2021-10-11

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the issue that there is no status code when the response body is a stream.
- _Change_
  - None

### HuaweiCloud SDK APIG

- _Features_
  - Support the interfaces:
    - `CreateAppCodeV2`
    - `CreateAppCodeAutoV2`
    - `ListAppCodesV2`
    - `ShowDetailsOfAppCodeV2`
    - `DeleteAppCodeV2`
    - `DebugApiV2`
    - `BatchPublishOrOfflineApiV2`
    - `ListApiVersionsV2`
    - `ChangeApiVersionV2`
    - `ListApiRuntimeDefinitionV2`
    - `ListApiVersionDetailV2`
    - `DeleteApiByVersionIdV2`
    - `ListAclStrategiesV2`
    - `BatchDeleteAclV2`
    - `CreateAclStrategyV2`
    - `ShowDetailsOfAclPolicyV2`
    - `UpdateAclStrategyV2`
    - `DeleteAclV2`
    - `BatchDeleteApiAclBindingV2`
    - `CreateApiAclBindingV2`
    - `DeleteApiAclBindingV2`
    - `ListAclPolicyBindedToApiV2`
    - `ListApisBindedToAclPolicyV2`
    - `ListApisUnbindedToAclPolicyV2`
    - `ListCustomAuthorizersV2`
    - `CreateCustomAuthorizerV2`
    - `ShowDetailsOfCustomAuthorizersV2`
    - `UpdateCustomAuthorizerV2`
    - `DeleteCustomAuthorizerV2`
    - `ExportApiDefinitionsV2`
    - `ImportApiDefinitionsV2`
    - `CreateVpcChannelV2`
    - `ListVpcChannelsV2`
    - `ShowDetailsOfVpcChannelV2`
    - `DeleteVpcChannelV2`
    - `UpdateVpcChannelV2`
    - `AddingBackendInstancesV2`
    - `ListBackendInstancesV2`
    - `DeleteBackendInstanceV2`
    - `ListLatelyApiStatisticsV2`
    - `ListLatelyGroupStatisticsV2`
    - `CreateGatewayResponseV2`
    - `ListGatewayResponsesV2`
    - `ShowDetailsOfGatewayResponseV2`
    - `UpdateGatewayResponseV2`
    - `DeleteGatewayResponseV2`
    - `ShowDetailsOfGatewayResponseTypeV2`
    - `UpdateGatewayResponseTypeV2`
    - `DeleteGatewayResponseTypeV2`
    - `ListTagsV2`
    - `CreateFeatureV2`
    - `ListFeaturesV2`
    - `CreateInstanceV2`
    - `ListInstancesV2`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListSkuInventories`.

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `ShowTask`:
    - `create_time`
    - `description`
    - `operate_mode`
    - `related_temp_running_data`
    - `run_status`
    - `update_time`

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameter `reverse_binding` to `optional` of the interface `DisassociateServerVirtualIp`.

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the optional value to `2,4,6,7,8,11,12,13,21` of the request parameter `attributes` of the interfaces `DetectFaceByFile`,`DetectFaceByBase64` and `DetectFaceByUrl`.

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `instance_mode` to the interface `ListInstances`.

### HuaweiCloud SDK IoTEdge

- _Features_
  - Support the service `IoTEdge`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `vpcId` to the interface `ListClusters`, and modify the type of the response parameter `start_time` of this interface: `string` -> `integer`

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `unchangeable_param`,`dry_run` and `count` to the interface `CreateInstance`.
  - Add the response parameter `enable_ssl` to the interface `CreateRestoreInstance`.

# 0.0.64 2021-09-29

### HuaweiCloud SDK Core

- _Features_
  - None
- _Bug Fix_
  - Fix the issue that RequestHandler is added in the wrong order.
- _Change_
  - None

### HuaweiCloud SDK BSS

- _Features_
  - Support the interface `ListSubCustomerBillDetail`.
- _Bug Fix_
  - None
- _Change_
  - Remove the interfaces `ListSubCustomerResFeeRecords` and `ListFreeResources`.
  - Modify the name of the interface: `ListParnterAdjustRecords` -> `ListPartnerAdjustRecords`

### HuaweiCloud SDK CPTS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add request body to the interface `CreateTemp`.

### HuaweiCloud SDK DNS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the response parameter `resource_detail` of the interface `ListTag`: `string` -> `object`
  - Modify the type of the response parameter `masters` of the interfaces `CreatePrivateZone`,`UpdatePublicZone` and `DeletePublicZone`: `string` -> `array`
  - Modify the type of the request parameter `ttl` of the interfaces `CreatePrivateZone` and `UpdatePublicZone`: `string` -> `integer`
  - Modify the type of the request parameters `limit` and `offset` of the interfaces `ListRecordSets`,`ListRecordSetsWithLine` and `ListRecordSetsByZone`: `string` -> `integer`
  - Add the response parameter `tags` to the interfaces `CreatePrivateZone`,`ListRecordSetsByZone`,`ShowRecordSetWithLine`,`ListPtrRecords` and `ListPublicZones`.

### HuaweiCloud SDK ECS

- _Features_
  - Support the following interfaces:
    - `ListServerTags`
    - `BatchAttachSharableVolumes`
    - `ShowServerAutoRecovery`
    - `BatchResetServersPassword`
    - `ReinstallServerWithoutCloudInit`
    - `ChangeServerOsWithoutCloudInit`
    - `BatchUpdateServersName`
    - `ShowServerPassword`
    - `AssociateServerVirtualIp`
    - `MigrateServer`
    - `ShowServerBlockDevice`
    - `DisassociateServerVirtualIp`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GaussDB

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the following interfaces:
    - `ShowMysqlEngineVersion` -> `ShowGaussMySqlEngineVersion`
    - `ShowMysqlFlavors` -> `ShowGaussMySqlFlavors`
    - `CreateMysqlInstance` -> `CreateGaussMySqlInstance`
    - `ShowMysqlInstanceList` -> `ListGaussMySqlInstances`
    - `DeleteMysqlInstance` -> `DeleteGaussMySqlInstance`
    - `ShowMysqlInstanceInfo` -> `ShowGaussMySqlInstanceInfo`
    - `CreateMysqlReadonlyNode` -> `CreateGaussMySqlReadonlyNode`
    - `DeleteMysqlReadonlyNode` -> `DeleteGaussMySqlReadonlyNode`
    - `ExpandMysqlInstanceVolume` -> `ExpandGaussMySqlInstanceVolume`
    - `UpdateMysqlBackupPolicy` -> `UpdateGaussMySqlBackupPolicy`
    - `UpdateMysqlInstanceName` -> `UpdateGaussMySqlInstanceName`
    - `ResetMysqlPassword` -> `ResetGaussMySqlPassword`
    - `ChangeMysqlInstanceSpecification` -> `ChangeGaussMySqlInstanceSpecification`
    - `ListDedicatedResources`  -> `ListGaussMySqlDedicatedResources`
    - `CreateMysqlProxy` -> `CreateGaussMySqlProxy`
    - `DeleteMysqlProxy` -> `DeleteGaussMySqlProxy`
    - `ShowMysqlProxy` -> `ShowGaussMySqlProxy`
    - `ShowMysqlProxyFlavors` -> `ShowGaussMySqlProxyFlavors`
    - `ExpandMysqlProxy` -> `ExpandGaussMySqlProxy`
    - `ListMysqlErrorLog` -> `ListGaussMySqlErrorLog`
    - `ListMysqlSlowLog` -> `ListGaussMySqlSlowLog`
    - `ShowMysqlProjectQuotas` -> `ShowGaussMySqlProjectQuotas`
    - `ShowMysqlQuotas` -> `ShowGaussMySqlQuotas`
    - `SetMysqlQuotas` -> `SetGaussMySqlQuotas`
    - `UpdateMysqlQuotas` -> `UpdateGaussMySqlQuotas`
    - `CreateMysqlBackup` -> `CreateGaussMySqlBackup`
    - `ShowMysqlBackupList` -> `ShowGaussMySqlBackupList`
    - `ShowMysqlBackupPolicy` -> `ShowGaussMySqlBackupPolicy`
    - `ListMysqlConfigurations` -> `ListGaussMySqlConfigurations`
    - `ShowMysqlJobInfo` -> `ShowGaussMySqlJobInfo`

### HuaweiCloud SDK ProjectMan

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `id` and `name` to the interface `ListIssueRecordsV4`.
  - Add the response parameter `status` to the interface `ListProjectIterationsV4`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the optional value `bigmen` to the response parameter `group_type` of the interface `ListFlavors`.

# 0.0.63 2021-09-26

### HuaweiCloud SDK DRS

- _Features_
  - Support the interface `BatchSetPolicy`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - Support the interfaces `CreateEditingJob`,`ListEditingJob` and `DeleteEditingJob`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK WAF

- _Features_
  - Support the following interfaces:
    - `ListIgnoreRule`
    - `ListStatistics`
    - `ListQpsTimeline`
    - `ListBandwidthTimeline`
    - `ListResponseCodeTimeline`
    - `ListTopAbnormal`
    - `ShowConsoleConfig`
- _Bug Fix_
  - None
- _Change_
  - None

# 3.0.62 2021-09-24

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - Fix the issue that the interface `ListRecordContents` cannot be found.
- _Change_
  - None

# 0.0.61 2021-09-24

### HuaweiCloud SDK BSS

- _Features_
  - Support the interfaces `ListParnterAdjustRecords` and `ListFreeResourceInfos`.
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListSubCustomerBillDetail`.

### HuaweiCloud SDK CCE

- _Features_
  - Support the interface `ShowQuotas`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Classroom

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the optional request parameter `testcases` to the interface `ApplyJudgement`.

### HuaweiCloud SDK Cloudtest

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameter `testcase_number` to `optional`, and remove the request parameter of the interface `ShowTestCaseDetailV2`.

### HuaweiCloud SDK GaussDBforopenGauss

- _Features_
  - Support the service `GaussDB(for openGauss)`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `ListRecordContents`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SWR

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `domain_id`, `scanned` and `tag_type` to the interface `ListRepositoryTags`.

# 0.0.60 2021-09-16

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `platformVersion` to the interface `ShowCluster`.

### HuaweiCloud SDK CDN

- _Features_
  - Support the interface `ShowDomainStats`.
- _Bug Fix_
  - Fix the issue of no response data when calling the interface `ShowDomainItemLocationDetails`.
- _Change_
  - None

### HuaweiCloud SDK DDM

- _Features_
  - Support the interfaces `ListSlowLog` and `ListReadWriteRatio`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DSC

- _Features_
  - Support the service `Data Security Center`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK GSL

- _Features_
  - Support the service `Global SIM Link`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IEC

- _Features_
  - Support the service `Intelligent EdgeCloud`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK IMS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the optional request parameter `__support_amd` to the interface `CreateDataImage`.
  - Add the response parameter `__support_amd` to the interface `ListImages`.

### HuaweiCloud SDK KMS

- _Features_
  - Support the interfaces `ShowPublicKey` and `Sign`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeInvoiceVerification`.
- _Bug Fix_
  - None
- _Change_
  - None

# 3.0.59 2021-09-10

### HuaweiCloud SDK BSS

- _Features_
  - Support the interfaces `ListSubCustomerBillDetail`, `ListResourceUsageSummary` and `ListResourceUsage`.
- _Bug Fix_
  - None
- _Change_
  - Remove the interface `ListResourceUsages`.

### HuaweiCloud SDK CBS

- _Features_
  - Support the interfaces `CreateTbSession`, `ExecuteTbSession` and `DeleteTbSession`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - Support the interfaces `AddNode` and `ResetNode`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CDN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `CreateDomain`:
    - `range_status`
    - `follow_status`
    - `origin_status`
    - `auto_refresh_preheat`
  - Add the required request parameter `switch` and optional request parameter `redirect_type` to the interface `UpdateDomainMultiCertificates`.
  - Add the required request parameter `switch` and optional request parameter `redirect_type` to the interface `UpdateHttpsInfo`.
  - Add the optional request parameter `create_time` to the interface `ShowHistoryTasks`.

### HuaweiCloud SDK DAS

- _Features_
  - Support the `Data Admin Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK DDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `status` and `fail_reason` to the interface `ShowJobDetail`.

### HuaweiCloud SDK EVS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameter `size` of the interface `CreateVolume` to `required`.

### HuaweiCloud SDK IVS

- _Features_
  - Support the service `Identity Verification Solution`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK OCR

- _Features_
  - Support the interface `RecognizeInvoiceVerification`.
- _Bug Fix_
  - None
- _Change_
  - Add the optional request parameter `return_verification` to the interface `RecognizeIdCard`.

### HuaweiCloud SDK RDS

- _Features_
  - Support the interface `UpdateDatabase`.
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `alias` to the interface `ListInstances`.
  - Add the optional request parameter `comment` to the interface `CreateDatabase`.

# 3.0.58 2021-08-31

### HuaweiCloud SDK CodeCraft

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `score` of the interfaces `CreateCompetitionScore` and `UpdateCompetitionScore`: `string` -> `double`.

### HuaweiCloud SDK CPTS

- _Features_
  - Support the service `CPTS`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK FRS

- _Features_
  - Support the following interfaces:
    - `DetectLiveByUrl`
    - `DetectLiveFaceByUrl`
    - `DetectLiveByFile`
    - `DetectLiveFaceByFile`
    - `DetectLiveByBase64`
    - `DetectLiveFaceByBase64`
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameters `video_frame_rate`,`audio_frame_rate`,`audio_bitrate` and `resolution`.

### HuaweiCloud SDK MRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameters `data_volume_count`, `data_volume_size` and `data_volume_type` to `optional` of the interface `ClusterScaling`.
  - Add the request parameter `auto_create_default_security_group` to the interface `CreateCluster`.

### HuaweiCloud SDK OCR

- _Features_
  - Support the service `Optical Character Recognition`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK SCM

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Set the request parameters `name`, `certificate`, `certificate_chain` and `private_key` to `required` of the interface `ImportCertificate`.
  - Set the request parameters `target_project` and `target_service` to `required` of the interface `PushCertificate`.

### HuaweiCloud SDK SMN

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `enterprise_project_id`, `name` and `fuzzy_name` to the interface `ListTopics`.
  - Add the request parameters `protocol`, `status` and `endpoint` to the interface `ListSubscriptions`.

# 0.0.57 2021-08-25

### HuaweiCloud SDK CBS

- _Features_
  - Support the `Conversational Bot Service`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CodeCraft

- _Features_
  - Support the `Conversational Bot Service`.
- _Bug Fix_
  - None
- _Change_
  - Modify the type of the request parameter `score` of the interfaces `CreateCompetitionScore` and `UpdateCompetitionScore`: `float` -> `string`

### HuaweiCloud SDK DDS

- _Features_
  - Support the following interfaces:
    - `ShowJobDetail`
    - `SwitchSlowlogDesensitization`
    - `ListFlavorInfos`
    - `UpdateInstanceRemark`
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `job_id` to the interfaces `RestoreInstance`, `CreateManualBackup` and `RestoreInstanceFromCollection`.
  - Add the response parameter `remark` to the interface `ListInstances`.

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `ListServerInterfaces`:
    - `delete_on_termination`
    - `driver_mode`
    - `min_rate`
    - `multiqueue_num`
    - `pci_address`
  - Add the response parameters `cpu_options` and `hypervisor` to the interface `ListServersDetails`.

### HuaweiCloud SDK EIP

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request and response parameter `public_border_group` to the interface `BatchCreateSharedBandwidths`.
  - Add the response parameter `public_border_group` to the interface `AddPublicipsIntoSharedBandwidth`.

### HuaweiCloud SDK FRS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the interface: `AuthorizeFaceRecognitionService` -> `ShowSubscribes`.

### HuaweiCloud SDK FunctionGraph

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameters `function_urn` and `type` to the interface `ExportFunction`.
  - -The optional value of the request parameter `filter` of the interface `ListStatistics` is modified to [`monitor_data`, `monthly_report`]

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
  - Support the following interfaces:
    - `ListDedicatedResources`
    - `ListFlavorInfos`
    - `ListConfigurationTemplates`
    - `ListInstancesByResourceTags`
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `dedicated_resource_id` to the interface `CreateInstance`.
  - Add the response parameter `dedicated_resource_id` to the interface `ListInstances`.

### HuaweiCloud SDK ImageSearch

- _Features_
  - Support the service `ImageSearch`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - Support the interface `RunRecord`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK MPC

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `aspect_ratio` of the interface `CreateTransTemplate`.
  - Remove the request parameter `GOP_structure` and `sr_factor` of the interface `CreateTranscodingTask`.
  - Remove the request parameter `GOP_structure` and `sr_factor` of the interface `CreateTemplateGroup`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Modify the name of the response parameter of the interface `ListJobInfoDetail`: `taskDetail` -> `task_detail`
  - Add the response parameter `count` to the interface `ListJobInfoDetail`.

# 0.0.56 2021-08-17

### HuaweiCloud SDK AntiDDoS

- _Features_
  - Support the service `Anti-DDoS`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK CCE

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the following response parameters to the interface `ListNodePools`:
    - `annotations`
    - `updateTimestamp`
    - `creationTimestamp`
    - `creatingNode`
    - `deletingNode`
    - `conditions`
    - `enterprise_project_id`
  - Add the response parameters `orderID` and `upgradefrom` to the interface `ListClusters`.

### HuaweiCloud SDK CloudTest

- _Features_
  - Support the service `CloudTest`.
- _Bug Fix_
  - None
- _Change_
  - None

### HuaweiCloud SDK ECS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the response parameter `ecs:instance_architecture` to the interface `ListServerResizeFlavors`.
  - Add the request parameters `tenancy` and `dedicated_host_id` to the interface `NovaCreateServers`.

### HuaweiCloud SDK ELB

- _Features_
  - None
- _Bug Fix_
  - Fix the problem of abnormal creation of LB.
- _Change_
  - None

### HuaweiCloud SDK Live

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Remove the request parameter `key` from the interface `CreateRecordCallbackConfig`.
  - Add the response parameter `sign_type` to the interface `ListRecordCallbackConfigs`.
  - Add the response parameters `status_describe` and `service_area` to the interface `ShowDomain`.

### HuaweiCloud SDK RDS

- _Features_
  - None
- _Bug Fix_
  - None
- _Change_
  - Add the request parameter `readonly` to the interfaces `AllowSqlserverDbUserPrivilege` and `RevokeSqlserverDbUserPrivilege`.

### HuaweiCloud SDK SMS

- _Features_
  - Support `Server Migration Service`.
- _Bug Fix_
  - None
- _Change_
  - None

# 0.0.55 2021-08-10

### HuaweiCloud SDK AS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the request parameters `key` and `value` of the interface `ListResourceInstances` as required parameters.

### HuaweiCloud SDK CloudBuild

- _Features_
    - Support the service `CloudBuild`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of the response parameter of the interfaces `ListBandwidths` and `ShowPublicip`: `publicip_border_group` -> `public_border_group`

### HuaweiCloud SDK EVS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `server_id` to the interface `ListVolumes`.

### HuaweiCloud SDK FRS

- _Features_
    - Support the `Face Recognition Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `order_id` from the interface `CreateDeployment`.

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the request parameter `value` of the interface `UpdateImage` as a required parameter.

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `new_partition_numbers` to the interface `UpdateInstanceTopic`.

### HuaweiCloud SDK LTS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `highlight` to the interface `ListLogs`.
    - Add the optional value `backwards` to the request parameter `search_type` of the interface `ListLogs`.

### HuaweiCloud SDK MPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `dynamic_range` from the interface `ListTranscodingTask`.
    - Remove the request parameter `tenant_id` from the interface `CreateTransTemplate`.
    - Remove the request parameters `percent` and `frame_type` from the interface `CreateThumbnailsTask`.
    - Remove the request parameter `black_enhance` from the interface `CreateTranscodingTask`.

### HuaweiCloud SDK MRS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `start_time` and `state` to the interface `DescribeCluster`.
    
### HuaweiCloud SDK SWR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the following response parameters to the interface `ShowAccessDomain`:
        - `namespace`
        - `repository`
        - `access_domain`
        - `permit`
        - `deadline`
        - `description`
        - `creator_id`
        - `creator_name`
        - `created`
        - `updated`
        - `status`

### HuaweiCloud SDK VPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `enable_dhcp` to the interface `NeutronListSubnets`.
    - Add the response parameter `security_groups_links` to the interface `NeutronListSecurityGroups`.

# 0.0.54 2021-07-27

### HuaweiCloud SDK Classroom

- _Features_
    - Support the interfaces `ApplyJudgement`,`ShowJudgementDetail`,`ShowJudgementFile`.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.53 2021-07-26

### HuaweiCloud SDK CDN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameters `urls`, `task_id` of the interface `ShowHistoryTasks`.
    - Remove the response parameters `task_id`, `process_reason`, modify the type of the request parameter `process_reason`:`integer`->`string`
    - Remove the request parameters `user_domain_id`, `task_id` of the interface `ShowTopUrl`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - Support the interface `ShowPlans`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `dcs_cluster_proxy2_node` to the interface `UpdateConfigurations`.

### HuaweiCloud SDK DDM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the request parameter `extend_authority` of the interface `UpdateUser`.

### HuaweiCloud SDK DDS

- _Features_
    - Support the interface `UpdateClientNetwork`.
- _Bug Fix_
    - None
- _Change_
    - Change the request parameters `start_time`,`stop_time` from `optional` to `required` of the interface `SetBalancerWindow`.
    - Add the request parameter `port` and response parameter `port` to the interface `CreateInstance`.

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support the interface `EnableLtsLogs`.
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `concurrent_num`,`id`,`encrypted_user_data`.
    - Add the response parameters `func_vpc_id`,`encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` to the interface `ListFunctions`, and remove the response parameters `version_description`,`last_modified_utc`,`dependencies` of this interface.
    - Remove the request parameter `name`,`last_modified`,`alias_urn` of the interface `UpdateVersionAlias`.
    - Add the response parameters `encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` of the interface `ShowFunctionConfig`, and remove the response parameters `version_description`,`concurrency` of this interface.
    - Add the response parameters `encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` to the interface `ListFunctionVersions`, and remove the response parameters `version_description`,`concurrency`,`depend_list`.
    - Add the response parameters `encrypted_user_data`,`long_time`,`log_group_id`,`log_stream_id`,`type` to the interface `ListFunctionVersions`, remove the response parameters `last_modified_utc`,`concurrency`.
    - Modify the type of the request parameter `size` of the interface `UpdateTrigger`: `string`->`integer`
    - Modify the type of the response parameter `size` of the interface `ShowDependency`: `string`->`integer`
    - Modify the type of the response parameter `size` of the interface `UpdateDependency`: `string`->`integer`

### HuaweiCloud SDK HSS

- _Features_
    - Support the interface `ListEvents`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `domain_source` of the interface `ShowDomain`.

### HuaweiCloud SDK MPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the response parameter `language` of the interface `ListTranscodeDetail`.
    - Remove the request parameter `project_id`,`tenant_project_id`,`domain_name`,`canonical_grant_id` of the interface `CreateThumbnailsTask`.
    - Add the response parameter `audit_report` to the interface `ListTranscodeDetail`.
    - Remove the response parameter `output_url` of the interface `QueryTranscodingsTask`.
    - Add the request parameter `audit` to the interface, and remove the request parameter `special_effect`,`quality_enhance`,`template_extend` of this interface.
    - Remove the response parameter `template_id`,`error` of the interface `ListWatermarkTemplate`.
    - Remove the request parameter `multidrm`,`preview_duration` of the interface `CreateVodTranscodingTask`.

### HuaweiCloud SDK VOD

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the request parameter `auto_publish` of the interface `CreateAssetByFileUpload`, and configure the optional values `0`,`1`.

### HuaweiCloud SDK WAF

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameters `response_time`,`response_size` of the interface `ListEvent`: `string`->`integer`.

# 0.0.52 2021-07-16

### HuaweiCloud SDK AS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `description` to the interface `CreateScalingV2Policy`.
    - Add the response parameter `description` to the interfaces `ShowScalingV2Policy`, `ShowScalingGroup`.

### HuaweiCloud SDK DCS

- _Features_
    - Support more interfaces:
        - `CreateDiagnosisTask`
        - `CreateRedislog`
        - `CreateRedislogDownloadLink`
        - `ListDiagnosisTasks`
        - `ListRedislog`
        - `ListSlowlog`
        - `ShowDiagnosisTaskDetails`
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `include_delete` to the interface `ListInstances`.

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - [Issue 40](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/40): Fix the issue that the type of the response parameter `__lazyloading` is incorrectly defined.
- _Change_
    - None

# 0.0.51 2021-07-09

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the issue of authentication failure when the query parameter name contains uppercase letters.
- _Change_
    - None

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - Fix the issue that the data structure of the response parameter `addresses` of the interface `ListBareMetalServers` is incorrectly defined.
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `smn_notify`,`threshold` to the interface `ListProtectable`.
    - Add the request parameter `add_policy_ids` and the response parameters `without_any_tag`,`smn_notify`,`threshold` to the interface `AssociateVaultPolicy`.

### HuaweiCloud SDK CCE

- _Features_
    - Support the interfaces `RemoveNode`,`MigrateNode`.
- _Bug Fix_
    - None
- _Change_
    - Add the request parameter `tobedeleted` to the interface `DeleteCluster`.

### HuaweiCloud SDK CDN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change the request parameters `start_time`,`end_time` from optional to required of the interface `ShowTopUrl`, and add the optional value `outside_mainland_china` to the request parameter `domain_name`.
    - Add the request parameter `service_area` to the interface `ShowDomainItemDetails`.

### HuaweiCloud SDK DDM

- _Features_
    - Support the `Distributed Database Middleware` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DNS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameters `masters`, `zones` of the interface `CreatePublicZone`: `string`->`array`

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameter `publicip_border_group` to the interfaces `CreateSharedBandwidth`,`ListBandwidths`.

### HuaweiCloud SDK GaussDB

- _Features_
    - Support `GaussDB` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `__root_origin`,`checksum`,`size` to the interfaces `GlanceCreateImageMetadata`.
    - Remove the request parameters `deleted`, `deleted_at` of the interface `GlanceAddImageMember`, and add the following request parameters:
        - `__lazyloading`
        - `__os_feature_list`
        - `__root_origin`
        - `__sequence_num`
        - `__support_agent_list`
        - `__system__cmkid`
        - `active_at`
        - `hw_vif_multiqueue_enabled`
        - `max_ram`
        - `__image_location`
        - `__is_config_init`
        - `__account_code`

### HuaweiCloud SDK IoTDA

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `edge_node_ids`, `last_update_time` to the interface `ListRules`.

### HuaweiCloud SDK LTS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameter `context` of the interface `ListStructuredLogsWithTimeRange`: `string`->`array`
    - Modify the name of the interfaces::
        - `UpdateLogContents`->`ListLogs`
        - `UpdateLogContents2`->`ListQueryStructuredLogs`
        - `UpdateLogContents3`->`ListStructuredLogsWithTimeRange`

### HuaweiCloud SDK MRS

- _Features_
    - Support the interface `ListClusters`.
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameter `clusterType` of the interface `DescribeCluster`: `string`->`integer`

### HuaweiCloud SDK SWR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the response parameters `domain_id`,`priority` to the interface `ShowRepository`.
    - Add the response parameter `template` to the interface `CreateRetention`.

# 0.0.50 2021-06-29

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add a request parameter `storage` to interfaces `CreateNodePool`,`ShowNodePool`,`UpdateNodePool`,`DeleteNodePool`.

### HuaweiCloud SDK DRS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the parameter `selected` of the interface `BatchUpdateUser`: `string`->`boolean`

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of the response parameter `port` of the interface `ListInstances`: `string`->`integer`.
    - Modify the name of response parameter of the interface `ListInstances`: `storage_engine`->`mode`
    - Remove a response parameter `node_name` and add a response parameter `time` to the interface `ListSlowLogs`.

### HuaweiCloud SDK MRS

- _Features_
    - None
- _Bug Fix_
    - Fix the issue of incorrect definition of some parameters.
- _Change_
    - Remove the parameters `start_time`,`state` of the interface `CreateCluster`.

### HuaweiCloud SDK NAT

- _Features_
    - None
- _Bug Fix_
    - Fix the issue that the request parameter `project_id` of the interface `ListNatGateways` is duplicated.
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of response parameters `port`,`node_num` of the interface `ShowInformationAboutDatabaseProxy`: `string`->`integer`

# 0.0.49 2021-06-25

### HuaweiCloud SDK APIG

- _Features_
    - Support more ineterfaces:
        - `ListGatewayResponsesV2`
        - `UpdateGatewayResponseV2`
        - `DeleteGatewayResponseV2`
        - `UpdateGatewayResponseTypeV2`
        - `DeleteGatewayResponseTypeV2`
        - `DeleteInstancesV2`
        - `UpdateInstanceV2`
        - `ListInstancesV2`
        - `RemoveEipV2`
        - `UpdateEngressEipV2`
        - `RemoveEngressEipV2`
        - `ListFeaturesV2`
        - `UpdateDomainV2`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BMS

- _Features_
    - Support interface `ChangeBaremetalServerOs`.
- _Bug Fix_
    - None
- _Change_
    - Modify the name of reponse parameter of interface `ChangeBaremetalServerName`: `server_tags`->`sys_tags`.

### HuaweiCloud SDK CDN

- _Features_
    - Support interface `ShowQuota`.
- _Bug Fix_
    - None
- _Change_
    - Modify the type of request parameter `url` of interface `ShowHistoryTaskDetails`: `integer`->`string`.

### HuaweiCloud SDK CloudRTC

- _Features_
    - Support `CloudRTC` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DRS

- _Features_
    - Support interface `ShowQuotas`.
- _Bug Fix_
    - None
- _Change_
    - Modify the type of request parameters `is_transfer`,`selected` of interface `BatchUpdateUser`: `string`->`boolean`.

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request parameters`permission_type`,`display_name`,`catalog`,`type` of interface `KeystoneListPermissions`.

### HuaweiCloud SDK LTS

- _Features_
    - Support `Log Tank Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Meeting

- _Features_
    - Support interface `InviteShare`.
- _Bug Fix_
    - None
- _Change_
    - Add request parameter `multiPicSaveOnly` to interface `SetMultiPicture`.
    - Add reponse parameter `leftReason` to interface `SearchHisMeetings`.

### HuaweiCloud SDK VOD

- _Features_
    - Support `Video on Demand` service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.48 2021-06-21

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add response parameters `server_tags`,`enterprise_project_id`,`group` to interface `ChangeBaremetalServerName`.

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - [Issue 22](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/22): Modify the optional value of response parameter `status` of interface `ListAddonInstances`.
- _Change_
    - None

### HuaweiCloud SDK CDN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove the request parameter `user_domain_id` of interface `ListDomains`.
    - Modify the name of interface: `ShowRefer` -> `ShowReferer`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request parameters to interface `ShowTemplateDetail`:
        - `template_url`
        - `create_time`
        - `last_modify_time`
        - `can_update`
        - `can_delete`
        - `need_hub`

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces:
        - `CreateRecordCallbackConfig`
        - `ShowRecordCallbackConfig`
        - `UpdateRecordCallbackConfig`
        - `DeleteRecordCallbackConfig`
        - `ListRecordCallbackConfigs`
        - `UpdateRecordRule`
        - `ShowRecordRule`
- _Bug Fix_
    - None
- _Change_
    - Modify the name of some interfaces:
        - `CreateRecordConfig` -> `CreateRecordRule`
        - `DeleteRecordConfig` -> `DeleteRecordRule`
        - `ListRecordConfigs` -> `ListRecordRules`
    - Remove some interfaces:
        - `ShowTraffic`
        - `ShowBandwidth`
        - `ShowOnlineUsers`

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the type of response parameter `partitions` of interface `ShowGroups`: `array[string]` -> `array[integer]`

### HuaweiCloud SDK RabbitMQ

- _Features_
    - None
- _Bug Fix_
    - Fix the issue of compilation failure.
- _Change_
    - None

# 0.0.47 2021-06-10

### HuaweiCloud SDK BSS

- _Features_
    - Support interfaces `ListFreeResources`,`ListFreeResourceUsages`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CDN

- _Features_
    - Support `Content Delivery Network` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DRS

- _Features_
    - Support `Data Replication Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support interfaces
        - `ImportFunction`
        - `ExportFunction`
        - `AsyncInvokeReservedFunction`
        - `DeleteReservedInstanceById`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK OSM

- _Features_
    - Support `Online Service Management`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support interfaces `SetBinlogClearPolicy`,`ShowBinlogClearPolicy`.
- _Bug Fix_
    - None
- _Change_
    - Add request parameters `offset`,`limit` to interface `ListOffSiteInstances`.

# 0.0.46 2021-06-04

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - [Issue 20](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/20): Fix the issue that the type of `extendParam`
      is defined incorrectly.
- _Change_
    - Add the request parameter `tobedeleted` to the interface `DeleteCluster`.

### HuaweiCloud SDK DDS

- _Features_
    - Support the interface `ShowQuotas`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IoTDA

- _Features_
    - Support interface `ListComplexQueryDevice`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK GaussDBforNoSQL

- _Features_
    - Support `GaussDBforNoSQL` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support the interface `ShowQuotas`
- _Bug Fix_
    - None
- _Change_
    - Modify the type of request parameter `restart` of the interface `StartInstanceRestartAction`: string -> object

# 0.0.45 2021-05-25

### HuaweiCloud SDK AS

- _Features_
    - Support more interfaces:
        - `ListApiVersions`
        - `ShowApiVersion`
        - `BatchProtectScalingInstances`
        - `BatchRemoveScalingInstances`
        - `CreateScalingTagInfo`
        - `BatchResumeScalingPolicies`
        - `BatchPauseScalingPolicies`
        - `PauseScalingGroup`
        - `BatchSetScalingInstancesStandby`
        - `BatchUnsetScalingInstancesStandby`
        - `ResumeScalingPolicy`
        - `PauseScalingPolicy`
- _Bug Fix_
    - None
- _Change_
    - Modify the operation name of the following interfaces:
        - from `ExecuteScalingPolicies` to `BatchDeleteScalingPolicies`
        - from `EnableOrDisableScalingGroup` to `ResumeScalingGroup`
        - from `UpdateScalingGroupInstance` to `BatchAddScalingInstances`
        - from `CompleteLifecycleAction` to `AttachCallbackInstanceLifeCycleHook`
    - Remove the interface: `DeleteScalingTags`
    - Add the parameter `enterprise_project_id` to the interface `ListScalingGroups`.
    - Add the parameter `log_id` to the interface `ListScalingActivityV2Logs`.

### HuaweiCloud SDK BSS

- _Features_
    - Support interface `ListCustomerBillsMonthlyBreakDown` and `ListOrderDiscounts`.
- _Bug Fix_
    - None
- _Change_
    - Add query parameter _bill_date_begin_ and _bill_date_end_ to interface `ListSubCustomerResFeeRecords`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - Support interface: `StopPipelineNew`.
- _Bug Fix_
    - None
- _Change_
    - Remove interfaces `StartPipeline`, `StopPipeline`.

### HuaweiCloud SDK CloudPipeline

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove interfaces `StartPipeline`,`StopPipeline`.

### HuaweiCloud SDK DMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of interface from `ShowProjectTags` to `ShowQueueProjectTags`.

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change request parameter `offset` of interface `ListEnterpriseProject` from required to optional.

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support more interfaces:
        - `ListFunctionAsyncInvokeConfig`
        - `ShowFunctionAsyncInvokeConfig`
        - `DeleteFunctionAsyncInvokeConfig`
        - `UpdateFunctionAsyncInvokeConfig`
- _Bug Fix_
    - None
- _Change_
    - Modify the name of request parameter of interfaces `DeleteVersionAlias`,`UpdateVersionAlias`
      ,`ShowVersionAlias`: `name` -> `alias_name`
    - Modify the name of request parameter of interfaces `DeleteFunctionTrigger`,`UpdateTrigger`
      ,`ShowFunctionTrigger`: `triggerId` -> `trigger_id`

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the parameter `access_mode` to interface `CreateUsers`.
    - Change the parameter `authentication_code` of interface `DeleteBindingDevice` from required to optional.

### HuaweiCloud SDK Kafka

- _Features_
    - Support more interfaces:
        - `CreateInstanceUser`
        - `BatchDeleteInstanceUsers`
        - `ShowInstanceUsers`
        - `ShowTopicAccessPolicy`
        - `UpdateTopicAccessPolicy`
        - `ShowKafkaTopicPartitionDiskusage`
        - `ShowInstanceMessages`
        - `ResetUserPassword`
- _Bug Fix_
    - None
- _Change_
    - Modify the name of the interface:
        - from `ShowInstanceTags` to `ShowKafkaTags`
        - from `ShowProjectTags` to `ShowKafkaProjectTags`
        - from `BatchCreateOrDeleteInstanceTag` to `BatchCreateOrDeleteKafkaTag`
    - Modify the request body name of the interface:
        - from `BatchCreateOrDeleteInstanceTagRequestBody` to `BatchCreateOrDeleteKafkaTagRequestBody`
    - Modify the data type of parameter `sink_max_tasks` in the request body of interface `UpdateSinkTaskQuota` from
      String to Integer.

### HuaweiCloud SDK OMS

- _Features_
    - Support `Object Storage Migration Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RabbitMQ

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of the following interfaces:
        - from `BatchCreateOrDeleteInstanceTag` to `BatchCreateOrDeleteRabbitMqTag`;
        - from `ShowProjectTags` to `ShowRabbitMqProjectTags`;
        - from `ShowInstanceTags` to `ShowRabbitMqTags`.
    - Modify the request body name of interface `BatchCreateOrDeleteInstanceTag`
      from `BatchCreateOrDeleteInstanceTagRequestBody` to
      `BatchCreateOrDeleteRabbitMqTagRequestBody`.

# 0.0.43-rc 2021-05-14

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Solve the issue of abnormal parsing result when using interface `NovaShowKeypair` to obtain the secret key.
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add new result values `CLOUDSSD` and `LOCALSSD` to response field `type` of interface `ListInstances`.
    - Add an optional request parameter `complete_version` to interface `ListBackups`.
    - Change request parameter `type` of interface `ListSlowlogStatistics` from optional to required.

# 0.0.42-rc 2021-05-10

### HuaweiCloud SDK BMS

- _Features_
    - Support interface `BatchCreateBaremetalServerTags`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support interfaces `MigrateAz`, `ListAz2Migrate`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/17): Fix the issue that `EpDetailType` enum
      is defined incorrectly.
- _Change_
    - None

### HuaweiCloud SDK KPS

- _Features_
    - None
- _Bug Fix_
    - [Issue 19](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/19): Fix the issue of type of response body
      in `ListKeypairs`.
- _Change_
    - None

### HuaweiCloud SDK MRS

- _Features_
    - Support more interfaces:
        - `BatchDeleteClusterTags`
        - `CreateClusterTag`
        - `DeleteClusterTag`
        - `ListClusterTags`
        - `ListAllTags`
        - `BatchCreateClusterTags`
        - `ListClustersByTags`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of response body of interface `ListOffSiteInstances`: `OffSiteInstanceListResponse`
      -> `OffSiteInstanceListResponseBody`
    - Modify the name of response field of interface `ListOffSiteInstances`: `offsite_backup_instances`
      -> `offsite_backup_instance`

# 0.0.41-rc 2021-04-30

### HuaweiCloud SDK BCS

- _Features_
    - Support interface `ListOpRecord`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support interfaces:
        - `ShowShardingBalancer`
        - `SetBalancerSwitch`
        - `SetBalancerWindow`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK HSS

- _Features_
    - Support interface `ListHosts`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add optional values to request parameter `type` of the interface `ShowDomainQuota`:
        - `assigment_group_mp`
        - `assigment_agency_mp`
        - `assigment_group_ep`
        - `assigment_user_ep`

### HuaweiCloud SDK IoTDA

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Remove interfaces:
        - `ListSubscriptions`
        - `CreateSubscription`
        - `UpdateSubscription`
        - `ShowSubscription`
        - `DeleteSubscription`

### HuaweiCloud SDK MPC

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request parameters `language`、`sky_switch` to the interface `CreateMpeCallBack`.
    - Update optional values of request parameter `subtitle_type` of interface `CreateTranscodingTask`.

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add a field `project_code` to response body of the interface `ShowProjectInfoV4`.

# 0.0.40-rc 2021-04-15

### HuaweiCloud SDK Moderation

- _Features_
    - Support `Moderation` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Image

- _Features_
    - Support `Image` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support more interfaces about database management operations.
        - `CreateSqlserverDatabase`
        - `DeleteSqlserverDatabase`
        - `ListSqlserverDatabases`
    - Support more interfaces about user management operations.
        - `CreateSqlserverDbUser`
        - `ListSqlserverDbUsers`
        - `ListAuthorizedSqlserverDbUsers`
        - `DeleteSqlserverDbUser`
        - `AllowSqlserverDbUserPrivilege`
        - `RevokeSqlserverDbUserPrivilege`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces `DeleteDatabaseUser`,`DeleteDatabaseRole`,`ShowConnectionStatistics`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add fields `closed_time` ,`id` ,`created_time` to reponse body of interfaces `ListIssuesV4`, `ListChildIssuesV4`.

### HuaweiCloud SDK VPC

- _Features_
    - None
- _Bug Fix_
    - Fix the bug, open the tags of the VPC and subnet.
- _Change_
    - None

# 0.0.39-rc 2021-03-30

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - Fix the issue that the interface for querying messages does not contain the timestamp field.
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add the attribute `name` to the response body `IssueResponseV4` of the interfaces `ShowIssueV4`
      and `UpdateIssueV4`.
    - Change the attribute `work_time` to `work_date` in `ShowProjectWorkHoursResponseBody` in the response body of the
      interfaces `ShowProjectWorkHours` and `ListProjectWorkHours`.

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change the request parameter `protocol`  of the interface `PublishMessage`  from mandatory to optional.
    - Change the attribute `subject`  of the class `PublishMessageRequestBody` in the request body of the
      interface `PublishMessage`  from mandatory to optional.

# 0.0.38-rc 2021-03-26

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that authorization failed in auto acquisition of domain id.
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of deserialization error of the response of interface `ListLiveStreamsOnline`.
- _Change_
    - None

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change the property `protocol` in `ListMessageTemplates` from required to optional.

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that some fields in the response body of interface `ListSlowlogStatistics` are empty.
- _Change_
    - None

# 0.0.37-rc 2021-03-19

### HuaweiCloud SDK Core

- _Features_
    - Support the file upload feature.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of deserialization failure of response body of interface `ListFlavors`.
- _Change_
    - None

# 0.0.36-rc 2021-03-16

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Add request field `enterprise_project_id` in interface `CreatePrePaidPublicip`.

### HuaweiCloud SDK ProjectMan

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that the SDK can not be used.
- _Change_
    - None

# 0.0.35-rc 2021-03-15

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of deserialization failure when the response body of the API contains special characters such
      as `\n`.
- _Change_
    - If the `endpoint` input by the user does not contain a protocol prefix, the `https` prefix will be automatically
      added.

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Class adjustment in interface `CreateAlarmRequestBody`: change class definition of property `metric`
      from `MetricInfoForAlarm` to `MetricForAlarm`.
    - Make the property `name` in class `MetricsDimension` required, which affects interfaces of `BatchListMetricData`
      , `CreateMetricData`, `CreateResourceGroup` , and `UpdateResourceGroup`.

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces:
        - `RestoreNewInstance`
        - `ListSessions`
        - `DeleteSession`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - Support more interface: `ShowServerGroup`.
- _Bug Fix_
    - None
- _Change_
    - Change the interface name from `ShowWindowsServerPassword` to `ShowServerPassword`.
    - Change the interface name from `DeleteWindowsServerPassword` to `DeleteServerPassword`.

### HuaweiCloud SDK ELB

- _Features_
    - Support more interface: `ListAllMembers`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK FunctionGraph

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `ListDependencies` adjustment: change the data type of property `size` of the response definition from
      string to int64.

### HuaweiCloud SDK IAM

- _Features_
    - Support more interfaces:
        - `KeystoneShowIdentityProvider`
        - `KeystoneCreateIdentityProvider`
        - `KeystoneUpdateIdentityProvider`
        - `KeystoneDeleteIdentityProvider`
        - `CreateTokenWithIdToken`
- _Bug Fix_
    - None
- _Change_
    - Do not support interface `CreateUnscopeTokenByIdpInitiated` anymore.

### HuaweiCloud SDK IMS

- _Features_
    - Support more interfaces:
        - `ListImageByTags` which mead list images queried by tags.
        - `ListImagesTags` which means list all tags of all images in current account.
        - `ListImageTags` which means list all tags of specified image.
        - `AddImageTag`
        - `DeleteImageTag`
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IoTDA

- _Features_
    - Support `IoT Device Access` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - Support more interfaces:
        - `CreateCustomfields`
        - `ShowBugsPerDeveloper`
        - `ShowCompletionRate`
        - `ShowBugDensityV2`
        - `ShowProjectInfoV4`
- _Bug Fix_
    - Change the incorrect name of interface from `ShowtIssueCompletionRate` to `ShowIssueCompletionRate`.
- _Change_
    - Change the data type of property `created_time` and `updated_time` in class `ListProjectV4ResponseBody` from
      string to int64.

### HuaweiCloud SDK RDS

- _Features_
    - Support `Postgresql` related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK SCM

- _Features_
    - Support `SSL Certificate Management` service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.34-rc 2021-02-27

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Optimize the description of `README` and the format of `CHANGELOG`.

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `WindowsBaremetalServerCleanPwd` to `DeleteWindowsBareMetalServerPassword`.
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - All properties with type `float32` have been changed to `float64`.

### HuaweiCloud SDK CCE

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of property `Content-Type` is required when sending requests and returns
      error `Unsupported Content Type`.
- _Change_
    - None

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of circular dependency in the `CreateAlarmResponse` class.
- _Change_
    - None

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces: `DownloadSlowlog` and `DownloadErrorlog`.
- _Bug Fix_
    - Correct operation name from `ModifyConfigurationParameter` to `UpdateConfigurationParameter`, change the class
      name of this operation from `ModifyConfigurationParameterRequestBody` to `UpdateConfigurationParameterRequestBody`
      .
- _Change_
    - None

### HuaweiCloud SDK DGC

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `ModifyJob` to `UpdateJob`.
    - Correct operation name from `ModifyScript` to `UpdateScript`.
    - Correct operation name from `ModifyResource` to `UpdateResource`.
- _Change_
    - None

### HuaweiCloud SDK DIS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of circular dependency in the classes of `CommitCheckpointRequest`, `PutRecordsRequest`
      , `CreateAppRequest`, `UpdatePartitionCountRequest`.
- _Change_
    - None

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `ListMenbers` to `ListMembers`.
- _Change_
    - None

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `ModifyEnterpriseProject` to `UpdateEnterpriseProject`.
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - Correct property from `pwd_stength` to `pwd_strength` in class `KeystoneUserResult`.
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - Correct operation name from `DoManualBackup` to `CreateManualBackup`.
    - Correct operation name from `ModifyConfiguration` to `UpdateConfiguration`.
    - Correct operation name from `ModifyInstanceConfiguration` to `UpdateInstanceConfiguration`.
    - Fix the problem of circular dependency in the classes of `CreateInstanceResponse`
      and `CreateConfigurationResponse`.
    - Fix the unavailable problem of operation `CreateInstance`.
- _Change_
    - Add property `is_auto_pay` to the operation `StartInstanceAction` in the scenario of changing a single-node system
      to a primary/standby mode.

# 0.0.33-rc 2021-02-07

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that request sends fail when the data type of request body is `string`.
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `ListOsVersions` adjustment: change the data type of `os_bit` which is the property of response of the
      interface from string to integer.

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces: ListLiveSampleLogs, CreateDomain, DeleteDomain, UpdateDomain, ShowDomain,
      CreateDomainMapping, DeleteDomainMapping
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK MRS

- _Features_
    - Support `MapReduce Service`.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.32-rc 2021-01-30

### HuaweiCloud SDK DDS

- _Features_
    - Support more interfaces: `ListApiVersion`, `ShowApiVersion`,`SetAuditlogPolicy`, `ShowAuditlogPolicy`
      , `ListAuditlogs`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DNS

- _Features_
    - Support `Domain Name Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change interface name from `UpdateAutoTerminateTimeServer` to `UpdateServerAutoTerminateTime`.

### HuaweiCloud SDK EVS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `CinderCreateVolume` is supported to specify the id of dedicated storage pool using
      property `OS-SCH-HNT:scheduler_hints`.
    - Modify property type of `allocated` of class `QuotaDetails` from `String` to `Integer`.

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Increases the property `access_mode` of response class of interface `ShowUser`.

# 0.0.31-rc 2021-01-25

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - The default value of `DefaultTimeout` is set to 120 seconds.

### HuaweiCloud SDK BSS

- _Features_
    - Support more interface: ListOrderDiscounts.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DNS

- _Features_
    - Support `Domain Name Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - Support more interface: `UpdateAutoTerminateTimeServer`.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.30-rc 2021-01-15

### HuaweiCloud SDK Core

- _Features_
    - Support function `ValueOf` to get region information.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CloudBuild

- _Features_
    - Support more interface: `ShowListHistory`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DGC

- _Features_
    - Support more interfaces: `Job` related interfaces, `Script` related interfaces, `Resource` related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the data type of response field `is_domain_owner` from string to boolean of interface `ShowUser`
      and `CreateUser`.

### HuaweiCloud SDK Live

- _Features_
    - Support more interface: `ListLiveStreamsOnline`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support more interfaces: ShowOffSiteBackupPolicy, SetOffSiteBackupPolicy, ListOffSiteBackups,
      ListOffSiteRestoreTimes, ListOffSiteRestoreTimes
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK SWR

- _Features_
    - Support `Software Repository for Container` service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.29-beta 2020-12-31

### HuaweiCloud SDK BMS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of interface: ListBareMetalServers.
    - Fix the problem of interface: ListBareMetalServerDetails.
    - Modify interface fields: ShowBaremetalServerInterfaceAttachments.
- _Change_
    - None

### HuaweiCloud SDK CloudIDE

- _Features_
    - Support more interface: ShowAccountStatus.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - Modify the interface return data type to prevent deserialization failure:
        - ListSlowlog: change data type of `Tags` from Tag to ResourceTag.
        - ListInstances: change data type of `duration` from int32 to string.
        - ShowBigkeyScanTaskDetails: change data type of `db` from int32 to string.
        - ShowHotkeyTaskDetails: change data type of `db` from int32 to string.
- _Change_
    - None

### HuaweiCloud SDK DGC

- _Features_
    - Support `Data Lake Governance Center` service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DIS

- _Features_
    - Support `Data Ingestion Service`.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface modification: response type of interface `CreateInstance` adjustment.

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - Modify the request parameters of interface `PublishMessage` from Object to Map<String, String>
- _Change_
    - None

# 0.0.28-beta 2020-12-28

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix response exception when using temporary AK/SK.
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - Change property type of `port` from string to integer.
- _Change_
    - None

# 0.0.27-beta 2020-12-25

### HuaweiCloud SDK DCS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of compilation failure in Mac OS.
- _Change_
    - Query parameter in interface `ListInstances` modification: id → instance_id.

### HuaweiCloud SDK DDS

- _Features_
    - Support Document Database Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Kafka

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of compilation failure in Mac OS.
- _Change_
    - None

### HuaweiCloud SDK RabbitMQ

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of compilation failure in Mac OS.
- _Change_
    - None

### HuaweiCloud SDK RMS

- _Features_
    - Support more interfaces: `Resources` related interfaces and `Region` related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.26-beta 2020-12-23

### HuaweiCloud SDK Core

- _Features_
    - Support Endpoint Resolver: it's supported to use {Service}Region when initializing {ServiceClient} which can
      automatically backfill endpoint. After choosing a region, the projectId/domainId will be backfilled automatically.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - Support more interfaces: ListMeasureUnits.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Update interface: ShowMetricData

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces: ShowStreamPortrait.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK MPC

- _Features_
    - Support more interfaces: QualityEnhanceTemplate related interfaces and MergeChannelsTask related interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK RDS

- _Features_
    - Support Relational Database Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK SMN

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Update field type in message_template_name.

# 0.0.25-beta 2020-12-15

### HuaweiCloud SDK CCE

- _Features_
    - Support Cloud Container Engine service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that sending request to interface `CreateListener` returns empty response.
    - Fix the problem that sending request to interface `CreateListener` returns response with wrong type.
- _Change_
    - None

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support more interfaces: UpdateFunctionReservedInstances.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK NAT

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that using interface `BatchCreateNatGatewayDnatRules` failed.
- _Change_
    - None

# 0.0.24-beta 2020-12-04

### HuaweiCloud SDK SMN

- _Features_
    - Support Simple Message Notification service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.23-beta 2020-11-30

### HuaweiCloud SDK BCS

- _Features_
    - Support BlockChain Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BMS

- _Features_
    - Support Bare Metal Server service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - Support more interfaces: ListUsageTypes, ModPeriodToOnDemand.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - Support more interfaces: MigrateVaultResource
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CES

- _Features_
    - Support more interfaces:
        - ListEvents
        - ListEventDetail
        - CreateResourceGroup
        - UpdateResourceGroup
        - DeleteResourceGroup
        - ListResourceGroup
        - UpdateAlarm
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK DCS

- _Features_
    - Support Distributed Cache Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - [Issue 21](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/21) Open related interface.

### HuaweiCloud SDK FunctionGraph

- _Features_
    - Support FunctionGraph service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - Support more interfaces.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Name of service client modification: LiveAPIClient → LiveClient.

# 0.0.22-beta 2020-11-17

### HuaweiCloud SDK AS

- _Features_
    - None
- _Bug Fix_
    - [Issue 8](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/8) Fix the problem that creating scaling
      policy failed.
- _Change_
    - None

### HuaweiCloud SDK DMS

- _Features_
    - Support Distributed Message Services, provide Kafka premium instances and RabbitMQ premium instances with
      dedicated resources.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Property adjustment:  increase property `dry_run` in interfaces `CreatePostPaidServers` and `CreateServers` which
      means whether parameters will be checked before sending real requests.

### HuaweiCloud SDK NAT

- _Features_
    - Support NAT Gateway service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK VPC

- _Features_
    - Support more interfaces: interfaces related to Network ACLs.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.21-beta 2020-11-11

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Update core code from [Pull requests #11](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/pull/11).
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface adjustment: property `object_type` in interface `CreateVault` support key `turbo`.
    - Interface adjustment: property `description` in interface `UpdatePolicy` is removed.

### HuaweiCloud SDK CES

- _Features_
    - Add examples of interface response and adjust some filed description.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CloudPipeline

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify the name of generated Client class: devcloudpipeline_client → cloudPipeline_client
    - Modify the name of generated Meta class: devcloudpipeline_meta → cloudPipeline_meta

### HuaweiCloud SDK DevStar

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Modify interface parameters and adjust sample code.

# 0.0.20-beta 2020-11-02

### HuaweiCloud SDK CES

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface adjustment: property `alarm_type` in class `CreateAlarmRequestBody` support key `RESOURCE_GROUP`.
    - Interface adjustment: property `meta_data` in class `ListAlarmHistoriesResponse` only returns total number of
      alarm histories.

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Modify wrong parameter in class `CreateL7ruleRequestBody`.
- _Change_
    - None

# 0.0.19-beta 2020-10-31

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix: fix the problem that when query parameter contains enumerated variables the request will fail.
    - [Issue 7](https://github.com/huaweicloud/huaweicloud-sdk-go-v3/issues/7) resolve the problem of using
      json.Marshal()
      returns object{}.
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - Support more interfaces: interfaces related to `TAG`.
- _Bug Fix_
    - [Issue 17](https://github.com/huaweicloud/huaweicloud-sdk-java-v3/issues/17) fix the problem of interface:
      ListBackups.
- _Change_
    - None

### HuaweiCloud SDK CTS

- _Features_
    - Support more interface: ListQuotas
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EPS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust interfaces' names, replace abbreviations of `EP` with `EnterpriseProject`, the involved interfaces are:
        1. ListEP → ListEnterpriseProject
        2. CreateEP → CreateEnterpriseProject
        3. ShowEP → ShowEnterpriseProject
        4. ModifyEP → ModifyEnterpriseProject
        5. EnableEP → EnableEnterpriseProject
        6. ShowEPQuota → ShowEnterpriseProjectQuota
        7. ShowResourceBindEP → ShowResourceBindEnterpriseProject
        8. DisableEP → DisableEnterpriseProject

### HuaweiCloud SDK Iam

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust interfaces' names, the involved interfaces are:
        1. KeystoneCreateUserTokenByPasswordAndMFA → KeystoneCreateUserTokenByPasswordAndMfa
        2. CreateUnscopeTokenByIDPInitiated → CreateUnscopeTokenByIdpInitiated

### HuaweiCloud SDK ProjectMan

- _Features_
    - Support more interfaces: iteration information, user information, project members, project information, project
      indicators, project statistics, etc.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.18-beta 2020-10-20

### HuaweiCloud SDK ELB

- _Features_
    - Support more interfaces of version v2.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.17-beta 2020-10-14

### HuaweiCloud SDK BSS

- _Features_
    - Partner center supports exporting product catalog prices.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - Support more interfaces of version v2 of Live.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.16-beta 2020-10-12

### HuaweiCloud SDK CTS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Delete deprecated interfaces of version v1.

### HuaweiCloud SDK DevStar

- _Features_
    - None
- _Bug Fix_
    - Modify the credential type of DevStarClient: the correct credential type is GlobalCredentials.
- _Change_
    - None

# 0.0.15-beta 2020-09-30

### HuaweiCloud SDK AS

- _Features_
    - Support Auto Scaling service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.14-beta 2020-09-24

### HuaweiCloud SDK BSS

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that the class `BssClient` cannot be loaded.
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Interface `ListPublicips` adjustment: enterprise_project_id does not support multi-value query.

# 0.0.13-beta 2020-09-16

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Fix parameter type of interfaces.
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Update interfaces.

### HuaweiCloud SDK EIP

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that not support transfer multiple values.
- _Change_
    - None

### HuaweiCloud SDK ELB

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that some parameters are wrong.
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust descriptions of interfaces.

### HuaweiCloud SDK Live

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Adjust descriptions of banned interface.

# 0.0.12.1-beta 2020-09-16

### HuaweiCloud SDK ECS

- _Features_
    - None
- _Bug Fix_
    - Fix parameter type of interfaces.
- _Change_
    - None

### HuaweiCloud SDK All

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - When the optional parameter type is list, the parameter type will be changed to a pointer, for example, []string
      to *[]string

# 0.0.12-beta 2020-09-11

### HuaweiCloud SDK KPS

- _Features_
    - Support KPS service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EVS

- _Features_
    - Support EVS service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - Fix response value type definition errors.
- _Change_
    - None

# 0.0.11-beta 2020-09-09

### HuaweiCloud SDK CBR

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that resources related interfaces have wrong data type.
- _Change_
    - None

### HuaweiCloud SDK VPC

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that security group related interfaces have wrong data type.
- _Change_
    - None

# 0.0.10-beta 2020-09-04

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that the enumeration code cannot be generated for integer enumeration parameters without format
      defined in yaml.
- _Change_
    - Modify User Agent of Http Request header.

# 0.0.9-beta 2020-08-28

### HuaweiCloud SDK CloudPipeline

- _Features_
    - Support CloudPipeline service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - Support more APIs: tags related APIs and shared bandwidth related APIs.
- _Bug Fix_
    - Interface BatchCreateBandwidth: modify field billing_info.
- _Change_
    - None

### HuaweiCloud SDK IAM

- _Features_
    - Support more APIs: consistency of console related APIs.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ProjectMan

- _Features_
    - Support Project Management service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK VPC

- _Features_
    - Support more APIs: security group, security group rules, available IP count related APIs.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.8-beta 2020-08-25

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - None
- _Change_
    - Change optional enum variable type to pointer.

### HuaweiCloud SDK ELB

- _Features_
    - Support Elastic Load Balance service.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.7-beta 2020-08-20

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem that some enum variables are unreadable.
- _Change_
    - Add 'E' as prefix to enum Variables which start with '_'.

# 0.0.6-beta 2020-08-20

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of missing the imports when the struct contains enum variables.
- _Change_
    - None

# 0.0.5-beta 2020-08-19

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the deserialization failure of enum variables.
    - Fix the deserialization failure of error response in some scenarios.
- _Change_
    - None

# 0.0.4-beta 2020-08-18

### HuaweiCloud SDK Core

- _Features_
    - None
- _Bug Fix_
    - Fix the problem of missing default values of Go basic type when serializing.
- _Change_
    - None

# 0.0.3-beta 2020-08-14

### HuaweiCloud SDK APIG

- _Features_
    - Support API Gateway service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK BSS

- _Features_
    - Support Business Support System.
- _Bug Fix_
    - None
- _Change_
    - None

# 0.0.2-beta 2020-8-11

### HuaweiCloud SDK Core

- _Features_
    - Support temporary AK/SK authentication mode.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CBR

- _Features_
    - Support Cloud Backup and Recovery service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK CloudIDE

- _Features_
    - Support Cloud IDE service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK ECS

- _Features_
    - Support Elastic Cloud Server service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EIP

- _Features_
    - Support Elastic IP service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK EVS

- _Features_
    - Support Elastic Volume Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK IMS

- _Features_
    - Support Image Management Service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK Live

- _Features_
    - Support Live service.
- _Bug Fix_
    - None
- _Change_
    - None

### HuaweiCloud SDK MPC

- _Features_
    - Support Media Processing Center.
- _Bug Fix_
    - None
- _Change_
    - None

## 3.0.1-beta 2020-07-30

### First Release

- _Supported Services_
    - Classroom
    - Cloud Trace Service(CTS)
    - DevStar
    - Enterprise Project Management Service(EPS)
    - Identity and Access Management(IAM)
    - Tag Management Service(TMS)
    - Virtual Private Cloud(VPC)
