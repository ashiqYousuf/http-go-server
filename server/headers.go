package server

import "strings"

func (server *Server) getHeaderMap(headers []string) map[string]string {
	headersMap := make(map[string]string, len(headers))
	for _, header := range headers {
		headerParts := strings.Split(header, ":")
		headersMap[headerParts[0]] = headerParts[1]
	}

	return headersMap
}
