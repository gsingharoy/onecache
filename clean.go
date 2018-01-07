package onecache

// deletes expired records
func (c *Cache) clean() {
  foundExpired := false
	for k, v := range c.recordMap {
		if v.hasExpired() {
      if !foundExpired{
        // Lock only if an entry for expiration has been found
        foundExpired = true
        cMutex.Lock()
        defer cMutex.Unlock()
      }
			delete(c.recordMap, k)
		}
	}
}
