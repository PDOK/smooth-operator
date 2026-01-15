package uptimeutils

import (
	"crypto/sha1" //nolint:gosec
	"encoding/hex"
	"sort"
	"strings"
)

func GetUptimeAnnotations(customResourceAnnotations map[string]string, id string, name string, url string, customResourceLabels map[string]string) map[string]string {

	tags := []string{
		"public-stats",
	}
	for _, v := range customResourceLabels {
		tags = append(tags, v)
	}
	sort.Strings(tags)

	ingressRouteAnnotations := make(map[string]string)

	ingressRouteAnnotations[UptimePrefix+"id"] = GetUptimeID(id)
	ingressRouteAnnotations[UptimePrefix+"name"] = name
	ingressRouteAnnotations[UptimePrefix+"url"] = url
	ingressRouteAnnotations[UptimePrefix+"tags"] = strings.Join(tags, ",")

	PassUptimeAnnotationsFromCRToIngressRoute(customResourceAnnotations, ingressRouteAnnotations)

	return ingressRouteAnnotations
}

func PassUptimeAnnotationsFromCRToIngressRoute(customResourceAnnotations map[string]string, ingressRouteAnnotations map[string]string) {
	for key, value := range customResourceAnnotations {
		if strings.HasPrefix(key, UptimePrefix) {
			ingressRouteAnnotations[key] = value
		}
	}
}

func GetUptimeID(seed string) string {
	sum := sha1.Sum([]byte(seed)) //nolint:gosec
	return hex.EncodeToString(sum[:])
}
