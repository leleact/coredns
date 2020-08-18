package dnsutil

import (
	"strings"

	"github.com/miekg/dns"
)

type rrset struct {
	qname string
	qtype uint16
}

// RRSets returns the RRSets from rrs as slice of []dns.RR. OPT records are ignored.
// The ordering of the returned slice is not specified.
func RRSets(rrs []dns.RR) [][]dns.RR {
	m := make(map[rrset][]dns.RR)

	for _, r := range rrs {
		// TODO: include other meta RRs in the ignore list?
		if r.Header().Rrtype == dns.TypeOPT {
			continue
		}

		n := strings.ToLower(r.Header().Name)
		t := r.Header().Rrtype

		m[rrset{n, t}] = append(m[rrset{n, t}], r)
	}
	if len(m) == 0 {
		return nil
	}

	sets := make([][]dns.RR, len(m))
	i := 0
	for _, v := range m {
		sets[i] = v
		i++
	}

	return sets
}
