package tikv

// Config is the configuration of TiKV.
type Config struct {
	Server    ServerConfig    `toml:"server"`
	Storage   StorageConfig   `toml:"storage"`
	Raftstore RaftstoreConfig `toml:"raftstore"`
	Rocksdb   DbConfig        `toml:"rocksdb"`
}

// DbConfig is the rocksdb config.
type DbConfig struct {
	WalRecoveryMode                  int64         `toml:"wal-recovery-mode"`
	WalDir                           string        `toml:"wal-dir"`
	WalTTLSeconds                    int64         `toml:"wal-ttl-seconds"`
	WalSizeLimit                     string        `toml:"wal-size-limit"`
	MaxTotalWalSize                  string        `toml:"max-total-wal-size"`
	MaxBackgroundJobs                int64         `toml:"max-background-jobs"`
	MaxManifestFileSize              string        `toml:"max-manifest-file-size"`
	CreateIfMissing                  bool          `toml:"create-if-missing"`
	MaxOpenFiles                     int64         `toml:"max-open-files"`
	EnableStatistics                 bool          `toml:"enable-statistics"`
	StatsDumpPeriod                  string        `toml:"stats-dump-period"`
	CompactionReadaheadSize          string        `toml:"compaction-readahead-size"`
	InfoLogMaxSize                   string        `toml:"info-log-max-size"`
	InfoLogRollTime                  string        `toml:"info-log-roll-time"`
	InfoLogKeepLogFileNum            int64         `toml:"info-log-keep-log-file-num"`
	InfoLogDir                       string        `toml:"info-log-dir"`
	RateBytesPerSec                  string        `toml:"rate-bytes-per-sec"`
	RateLimiterMode                  int64         `toml:"rate-limiter-mode"`
	AutoTuned                        bool          `toml:"auto-tuned"`
	BytesPerSync                     string        `toml:"bytes-per-sync"`
	WalBytesPerSync                  string        `toml:"wal-bytes-per-sync"`
	MaxSubCompactions                int64         `toml:"max-sub-compactions"`
	WritableFileMaxBufferSize        string        `toml:"writable-file-max-buffer-size"`
	UseDirectIoForFlushAndCompaction bool          `toml:"use-direct-io-for-flush-and-compaction"`
	EnablePipelinedWrite             bool          `toml:"enable-pipelined-write"`
	Defaultcf                        CfConfig      `toml:"defaultcf"`
	Writecf                          CfConfig      `toml:"writecf"`
	Lockcf                           CfConfig      `toml:"lockcf"`
	Raftcf                           CfConfig      `toml:"raftcf"`
	Titan                            TitanDBConfig `toml:"titan"`
}

// CfConfig is the config of a cf
type CfConfig struct {
	BlockSize                       string        `toml:"block-size"`
	BlockCacheSize                  string        `toml:"block-cache-size"`
	DisableBlockCache               bool          `toml:"disable-block-cache"`
	CacheIndexAndFilterBlocks       bool          `toml:"cache-index-and-filter-blocks"`
	PinL0FilterAndIndexBlocks       bool          `toml:"pin-l0-filter-and-index-blocks"`
	UseBloomFilter                  bool          `toml:"use-bloom-filter"`
	OptimizeFiltersForHits          bool          `toml:"optimize-filters-for-hits"`
	WholeKeyFiltering               bool          `toml:"whole-key-filtering"`
	BloomFilterBitsPerKey           int64         `toml:"bloom-filter-bits-per-key"`
	BlockBasedBloomFilter           bool          `toml:"block-based-bloom-filter"`
	ReadAmpBytesPerBit              int64         `toml:"read-amp-bytes-per-bit"`
	CompressionPerLevel             []string      `toml:"compression-per-level"`
	WriteBufferSize                 string        `toml:"write-buffer-size"`
	MaxWriteBufferNumber            int64         `toml:"max-write-buffer-number"`
	MinWriteBufferNumberToMerge     int64         `toml:"min-write-buffer-number-to-merge"`
	MaxBytesForLevelBase            string        `toml:"max-bytes-for-level-base"`
	TargetFileSizeBase              string        `toml:"target-file-size-base"`
	Level0FileNumCompactionTrigger  int64         `toml:"level0-file-num-compaction-trigger"`
	Level0SlowdownWritesTrigger     int64         `toml:"level0-slowdown-writes-trigger"`
	Level0StopWritesTrigger         int64         `toml:"level0-stop-writes-trigger"`
	MaxCompactionBytes              string        `toml:"max-compaction-bytes"`
	CompactionPri                   int64         `toml:"compaction-pri"`
	DynamicLevelBytes               bool          `toml:"dynamic-level-bytes"`
	NumLevels                       int64         `toml:"num-levels"`
	MaxBytesForLevelMultiplier      int64         `toml:"max-bytes-for-level-multiplier"`
	CompactionStyle                 int64         `toml:"compaction-style"`
	DisableAutoCompactions          bool          `toml:"disable-auto-compactions"`
	SoftPendingCompactionBytesLimit string        `toml:"soft-pending-compaction-bytes-limit"`
	HardPendingCompactionBytesLimit string        `toml:"hard-pending-compaction-bytes-limit"`
	ForceConsistencyChecks          bool          `toml:"force-consistency-checks"`
	PropSizeIndexDistance           int64         `toml:"prop-size-index-distance"`
	PropKeysIndexDistance           int64         `toml:"prop-keys-index-distance"`
	EnableDoublySkiplist            bool          `toml:"enable-doubly-skiplist"`
	Titan                           TitanCfConfig `toml:"titan"`
}

// TitanCfConfig is the titian config.
type TitanCfConfig struct {
	MinBlobSize             string  `toml:"min-blob-size"`
	BlobFileCompression     string  `toml:"blob-file-compression"`
	BlobCacheSize           string  `toml:"blob-cache-size"`
	MinGcBatchSize          string  `toml:"min-gc-batch-size"`
	MaxGcBatchSize          string  `toml:"max-gc-batch-size"`
	DiscardableRatio        float64 `toml:"discardable-ratio"`
	SampleRatio             float64 `toml:"sample-ratio"`
	MergeSmallFileThreshold string  `toml:"merge-small-file-threshold"`
	BlobRunMode             string  `toml:"blob-run-mode"`
}

// TitanDBConfig is the config a titian db.
type TitanDBConfig struct {
	Enabled         bool   `toml:"enabled"`
	Dirname         string `toml:"dirname"`
	DisableGc       bool   `toml:"disable-gc"`
	MaxBackgroundGc int64  `toml:"max-background-gc"`
	// The value of this field will be truncated to seconds.
	PurgeObsoleteFilesPeriod string `toml:"purge-obsolete-files-period"`
}

// StorageConfig is the config of storage
type StorageConfig struct {
	DataDir                        string           `toml:"data-dir"`
	MaxKeySize                     int64            `toml:"max-key-size"`
	SchedulerNotifyCapacity        int64            `toml:"scheduler-notify-capacity"`
	SchedulerConcurrency           int64            `toml:"scheduler-concurrency"`
	SchedulerWorkerPoolSize        int64            `toml:"scheduler-worker-pool-size"`
	SchedulerPendingWriteThreshold string           `toml:"scheduler-pending-write-threshold"`
	BlockCache                     BlockCacheConfig `toml:"block-cache"`
}

// BlockCacheConfig is the config of a block cache
type BlockCacheConfig struct {
	Shared              bool    `toml:"shared"`
	Capacity            string  `toml:"capacity"`
	NumShardBits        int64   `toml:"num-shard-bits"`
	StrictCapacityLimit bool    `toml:"strict-capacity-limit"`
	HighPriPoolRatio    float64 `toml:"high-pri-pool-ratio"`
	MemoryAllocator     string  `toml:"memory-allocator"`
}

// ServerConfig is the configuration of TiKV server.
type ServerConfig struct {
	GrpcCompressionType              string            `toml:"grpc-compression-type"`
	GrpcConcurrency                  int64             `toml:"grpc-concurrency"`
	GrpcConcurrentStream             int64             `toml:"grpc-concurrent-stream"`
	GrpcRaftConnNum                  int64             `toml:"grpc-raft-conn-num"`
	GrpcStreamInitialWindowSize      string            `toml:"grpc-stream-initial-window-size"`
	GrpcKeepaliveTime                string            `toml:"grpc-keepalive-time"`
	GrpcKeepaliveTimeout             string            `toml:"grpc-keepalive-timeout"`
	ConcurrentSendSnapLimit          int64             `toml:"concurrent-send-snap-limit"`
	ConcurrentRecvSnapLimit          int64             `toml:"concurrent-recv-snap-limit"`
	EndPointRecursionLimit           int64             `toml:"end-point-recursion-limit"`
	EndPointStreamChannelSize        int64             `toml:"end-point-stream-channel-size"`
	EndPointBatchRowLimit            int64             `toml:"end-point-batch-row-limit"`
	EndPointStreamBatchRowLimit      int64             `toml:"end-point-stream-batch-row-limit"`
	EndPointEnableBatchIfPossible    bool              `toml:"end-point-enable-batch-if-possible"`
	EndPointRequestMaxHandleDuration string            `toml:"end-point-request-max-handle-duration"`
	SnapMaxWriteBytesPerSec          string            `toml:"snap-max-write-bytes-per-sec"`
	SnapMaxTotalSize                 string            `toml:"snap-max-total-size"`
	StatsConcurrency                 int64             `toml:"stats-concurrency"`
	HeavyLoadThreshold               int64             `toml:"heavy-load-threshold"`
	HeavyLoadWaitDuration            string            `toml:"heavy-load-wait-duration"`
	Labels                           map[string]string `toml:"labels"`
}

// RaftstoreConfig is the configuration of TiKV raftstore component.
type RaftstoreConfig struct {
	// true for high reliability, prevent data loss when power failure.
	SyncLog bool `toml:"sync-log"`

	// raft-base-tick-interval is a base tick interval (ms).
	RaftBaseTickInterval     string `toml:"raft-base-tick-interval"`
	RaftHeartbeatTicks       int64  `toml:"raft-heartbeat-ticks"`
	RaftElectionTimeoutTicks int64  `toml:"raft-election-timeout-ticks"`
	// When the entry exceed the max size, reject to propose it.
	RaftEntryMaxSize string `toml:"raft-entry-max-size"`

	// Interval to gc unnecessary raft log (ms).
	RaftLogGCTickInterval string `toml:"raft-log-gc-tick-interval"`
	// A threshold to gc stale raft log, must >= 1.
	RaftLogGCThreshold int64 `toml:"raft-log-gc-threshold"`
	// When entry count exceed this value, gc will be forced trigger.
	RaftLogGCCountLimit int64 `toml:"raft-log-gc-count-limit"`
	// When the approximate size of raft log entries exceed this value
	// gc will be forced trigger.
	RaftLogGCSizeLimit string `toml:"raft-log-gc-size-limit"`
	// When a peer is not responding for this time, leader will not keep entry cache for it.
	RaftEntryCacheLifeTime string `toml:"raft-entry-cache-life-time"`
	// When a peer is newly added, reject transferring leader to the peer for a while.
	RaftRejectTransferLeaderDuration string `toml:"raft-reject-transfer-leader-duration"`

	// Interval (ms) to check region whether need to be split or not.
	SplitRegionCheckTickInterval string `toml:"split-region-check-tick-interval"`
	/// When size change of region exceed the diff since last check, it
	/// will be checked again whether it should be split.
	RegionSplitCheckDiff string `toml:"region-split-check-diff"`
	/// Interval (ms) to check whether start compaction for a region.
	RegionCompactCheckInterval string `toml:"region-compact-check-interval"`
	// delay time before deleting a stale peer
	CleanStalePeerDelay string `toml:"clean-stale-peer-delay"`
	/// Number of regions for each time checking.
	RegionCompactCheckStep int64 `toml:"region-compact-check-step"`
	/// Minimum number of tombstones to trigger manual compaction.
	RegionCompactMinTombstones int64 `toml:"region-compact-min-tombstones"`
	/// Minimum percentage of tombstones to trigger manual compaction.
	/// Should between 1 and 100.
	RegionCompactTombstonesPercent int64  `toml:"region-compact-tombstones-percent"`
	PdHeartbeatTickInterval        string `toml:"pd-heartbeat-tick-interval"`
	PdStoreHeartbeatTickInterval   string `toml:"pd-store-heartbeat-tick-interval"`
	SnapMgrGCTickInterval          string `toml:"snap-mgr-gc-tick-interval"`
	SnapGCTimeout                  string `toml:"snap-gc-timeout"`
	LockCfCompactInterval          string `toml:"lock-cf-compact-interval"`
	LockCfCompactBytesThreshold    string `toml:"lock-cf-compact-bytes-threshold"`

	NotifyCapacity  int64 `toml:"notify-capacity"`
	MessagesPerTick int64 `toml:"messages-per-tick"`

	/// When a peer is not active for max-peer-down-duration
	/// the peer is considered to be down and is reported to PD.
	MaxPeerDownDuration string `toml:"max-peer-down-duration"`

	/// If the leader of a peer is missing for longer than max-leader-missing-duration
	/// the peer would ask pd to confirm whether it is valid in any region.
	/// If the peer is stale and is not valid in any region, it will destroy itself.
	MaxLeaderMissingDuration string `toml:"max-leader-missing-duration"`
	/// Similar to the max-leader-missing-duration, instead it will log warnings and
	/// try to alert monitoring systems, if there is any.
	AbnormalLeaderMissingDuration string `toml:"abnormal-leader-missing-duration"`
	PeerStaleStateCheckInterval   string `toml:"peer-stale-state-check-interval"`

	LeaderTransferMaxLogLag int64 `toml:"leader-transfer-max-log-lag"`

	SnapApplyBatchSize string `toml:"snap-apply-batch-size"`

	// Interval (ms) to check region whether the data is consistent.
	ConsistencyCheckInterval string `toml:"consistency-check-interval"`

	ReportRegionFlowInterval string `toml:"report-region-flow-interval"`

	// The lease provided by a successfully proposed and applied entry.
	RaftStoreMaxLeaderLease string `toml:"raft-store-max-leader-lease"`

	// Right region derive origin region id when split.
	RightDeriveWhenSplit bool `toml:"right-derive-when-split"`

	AllowRemoveLeader bool `toml:"allow-remove-leader"`

	/// Max log gap allowed to propose merge.
	MergeMaxLogGap int64 `toml:"merge-max-log-gap"`
	/// Interval to re-propose merge.
	MergeCheckTickInterval string `toml:"merge-check-tick-interval"`

	UseDeleteRange bool `toml:"use-delete-range"`

	CleanupImportSstInterval string `toml:"cleanup-import-sst-interval"`

	ApplyMaxBatchSize int64 `toml:"apply-max-batch-size"`
	ApplyPoolSize     int64 `toml:"apply-pool-size"`

	StoreMaxBatchSize int64 `toml:"store-max-batch-size"`
	StorePoolSize     int64 `toml:"store-pool-size"`
	HibernateRegions  bool  `toml:"hibernate-regions"`
}
