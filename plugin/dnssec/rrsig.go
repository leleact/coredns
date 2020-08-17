package dnssec

import "github.com/miekg/dns"

// newRRSIG return a new RRSIG, with all fields filled out, except the signed data.
func (k *DNSKEY) newRRSIG(signerName string, ttl, incep, expir uint32) *dns.RRSIG {
	sig := new(dns.RRSIG)

	sig.Hdr.Rrtype = dns.TypeRRSIG
	sig.Algorithm = k.K.Algorithm
	sig.KeyTag = k.tag
	sig.SignerName = signerName
	sig.Hdr.Ttl = ttl
	sig.OrigTtl = origTTL

	sig.Inception = incep
	sig.Expiration = expir

	return sig
}

const origTTL = 3600
