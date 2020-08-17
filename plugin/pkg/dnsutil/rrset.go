package dnsutil

import "github.com/miekg/dns"

type rrset struct {
	qname string
	qtype uint16
}

// RRSets returns the RRSets from rrs as slice of []dns.RR. OPT records are ignored. The ordering of the returned
// slice is not specified.
func RRSets(rrs []dns.RR) [][]dns.RR {
	// TODO: include other meta RRs in the ignore list?
	m := make(map[rrset][]dns.RR)

	for _, r := range rrs {
		if r.Header().Rrtype == dns.TypeOPT {
			continue
		}

		if s, ok := m[rrset{r.Header().Name, r.Header().Rrtype}]; ok {
			s = append(s, r)
			m[rrset{r.Header().Name, r.Header().Rrtype}] = s
			continue
		}

		s := make([]dns.RR, 1, 2)
		s[0] = r
		m[rrset{r.Header().Name, r.Header().Rrtype}] = s
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
