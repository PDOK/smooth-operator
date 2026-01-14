package uptime_utils

import (
	"crypto/sha1"
	"encoding/hex"
	"strings"
)

func GetUptimeAnnotations(customResourceAnnotations map[string]string, id string, name string, url string, tags []string) map[string]string {

	ingressRouteAnnotations := make(map[string]string)

	ingressRouteAnnotations[UptimePrefix+"id"] = GetUptimeId(id)
	ingressRouteAnnotations[UptimePrefix+"name"] = name
	ingressRouteAnnotations[UptimePrefix+"url"] = url
	ingressRouteAnnotations[UptimePrefix+"tags"] = strings.Join(tags, ",")

	PassUptimeAnnotationsFormCRToIngressRoute(customResourceAnnotations, ingressRouteAnnotations)

	return ingressRouteAnnotations
}

func PassUptimeAnnotationsFormCRToIngressRoute(customResourceAnnotations map[string]string, ingressRouteAnnotations map[string]string) {
	for key, value := range customResourceAnnotations {
		if len(key) >= len(UptimePrefix) && key[:len(UptimePrefix)] == UptimePrefix {
			ingressRouteAnnotations[key] = value
		}
	}
}

func GetUptimeId(seed string) string {
	sum := sha1.Sum([]byte(seed))
	return hex.EncodeToString(sum[:])
}
