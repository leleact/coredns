package dnsutil

import (
	"testing"

	"github.com/coredns/coredns/plugin/test"
	"github.com/miekg/dns"
)

func TestRRSets(t *testing.T) {
	// define 3 RRsets
	rrs := []dns.RR{
		test.A("example.org. IN A 127.0.0.1"),
		test.A("example.org. IN A 127.0.0.2"),
		test.AAAA("example.org. IN AAAA ::1"),
	}

	sets := RRSets(rrs)
	for i := range sets {
		if sets[i][0].Header().Rrtype == dns.TypeA {
			if len(sets[i]) != 2 {
				t.Errorf("Expected 2 RRs in the %q RRSet", "A")
			}
		}
		if sets[i][0].Header().Rrtype == dns.TypeAAAA {
			if len(sets[i]) != 1 {
				t.Errorf("Expected 1 RRs in the %q RRSet", "AAAA")
			}
		}

	}
}

func TestRRSetsOwners(t *testing.T) {
	// define 3 RRsets
	rrs := []dns.RR{
		test.A("example.org. IN A 127.0.0.1"),
		test.A("example.com. IN A 127.0.0.2"),
		test.AAAA("example.org. IN AAAA ::1"),
	}

	sets := RRSets(rrs)
	for i := range sets {
		if sets[i][0].Header().Rrtype == dns.TypeA {
			if len(sets[i]) != 1 {
				t.Errorf("Expected 1 RRs in the %q RRSet", "A")
			}
		}
		if sets[i][0].Header().Rrtype == dns.TypeAAAA {
			if len(sets[i]) != 1 {
				t.Errorf("Expected 1 RRs in the %q RRSet", "AAAA")
			}
		}

	}
}
