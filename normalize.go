package hostutils

// Normalize Unpack and pack hosts
func Normalize(hosts []string) (packedHosts []string) {
	return Pack(Unpack(hosts))
}

// NormalizeString Unpack and pack hosts
func NormalizeString(hosts string) (packedHosts []string) {
	return Pack(Unpack([]string{hosts}))
}
