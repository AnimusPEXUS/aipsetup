package cachepresets

import "time"

type CachePreset struct {
	// disables the cache. cache files, though, should not be deleted until
	// CacheSingleFileLifetime passes
	PassThrough bool

	DirectoryListingTimeout time.Duration
	CacheSingleFileLifetime time.Duration
}

// TODO: looks like all this cachesettings package is unused
