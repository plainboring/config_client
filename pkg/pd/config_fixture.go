package pd

// ConfigFixture is a pd configuration fixture
var ConfigFixture = `
Version = false
ConfigCheck = false
client-urls = ""
peer-urls = ""
advertise-client-urls = ""
advertise-peer-urls = ""
name = ""
data-dir = ""
ForceNewCluster = false
EnableGRPCGateway = false
initial-cluster = ""
initial-cluster-state = ""
join = ""
lease = 0
log-file = ""
log-level = ""
quota-backend-bytes = 0
auto-compaction-mode = ""
auto-compaction-retention = ""
enable-prevote = false
namespace-classifier = ""
DisableStrictReconfigCheck = false

[log]
  level = ""
  format = ""
  disable-timestamp = false
  development = false
  disable-caller = false
  disable-stacktrace = false
  disable-error-verbose = false
  [log.file]
    filename = ""
    log-rotate = false
    max-size = 0
    max-days = 0
    max-backups = 0

[tso-save-interval]
  Duration = 0

[metric]
  job = ""
  address = ""
  [metric.interval]
    Duration = 0

[schedule]
  max-snapshot-count = 3
  max-pending-peer-count = 16
  max-merge-region-size = 20
  max-merge-region-keys = 200000
  leader-schedule-limit = 4
  leader-schedule-strategy = "count"
  region-schedule-limit = 2048
  replica-schedule-limit = 64
  merge-schedule-limit = 8
  hot-region-schedule-limit = 4
  hot-region-cache-hits-threshold = 3
  store-balance-rate = 15.0
  tolerant-size-ratio = 0.0
  low-space-ratio = 0.8
  high-space-ratio = 0.6
  scheduler-max-waiting-operator = 3
  disable-raft-learner = false
  disable-remove-down-replica = false
  disable-replace-offline-replica = false
  disable-make-up-replica = false
  disable-remove-extra-replica = false
  disable-location-replacement = false
  disable-namespace-relocation = false
  [schedule.split-merge-interval]
    Duration = 3600000000000
  [schedule.patrol-region-interval]
    Duration = 100000000
  [schedule.max-store-down-time]
    Duration = 1800000000000

  [[schedule.schedulers]]
    type = "balance-region"
    disable = false

  [[schedule.schedulers]]
    type = "balance-leader"
    disable = false

  [[schedule.schedulers]]
    type = "hot-region"
    disable = false

  [[schedule.schedulers]]
    type = "label"
    disable = false
  [schedule.SchedulersPayload]
    balance-hot-region-scheduler = "null"
    balance-leader-scheduler = "null"
    balance-region-scheduler = "null"
    label-scheduler = "null"

[replication]
  max-replicas = 3

[pd-server]
  use-region-storage = true

[ClusterVersion]
  Major = 0
  Minor = 0
  Patch = 0
  PreRelease = ""
  Metadata = ""

[tick-interval]
  Duration = 0

[election-interval]
  Duration = 0

[security]
  cacert-path = ""
  cert-path = ""
  key-path = ""

[HeartbeatStreamBindInterval]
  Duration = 0

[LeaderPriorityCheckInterval]
  Duration = 0`
