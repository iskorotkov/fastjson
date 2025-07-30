package compliance

import (
	"net"
	"net/netip"
	"testing"
)

func TestMarshal_NetIP(t *testing.T) {
	runMarshalTest(t, "ipv4", net.ParseIP("192.168.1.1"), `"192.168.1.1"`)
	runMarshalTest(t, "ipv6", net.ParseIP("::1"), `"::1"`)
	runMarshalTest(t, "ipv6/full", net.ParseIP("2001:db8::1"), `"2001:db8::1"`)
	runMarshalTest(t, "nil", net.IP(nil), "null")
}

func TestUnmarshal_NetIP(t *testing.T) {
	runUnmarshalTest(t, "ipv4", `"192.168.1.1"`, net.ParseIP("192.168.1.1"))
	runUnmarshalTest(t, "ipv6", `"::1"`, net.ParseIP("::1"))
	runUnmarshalTest(t, "null", "null", net.IP(nil))
	runUnmarshalErrorTest[net.IP](t, "invalid", `"not.an.ip"`)
}

func TestMarshal_NetipAddr(t *testing.T) {
	runMarshalTest(t, "ipv4", netip.MustParseAddr("192.168.1.1"), `"192.168.1.1"`)
	runMarshalTest(t, "ipv6", netip.MustParseAddr("::1"), `"::1"`)
	runMarshalTest(t, "zero", netip.Addr{}, `""`)
}

func TestUnmarshal_NetipAddr(t *testing.T) {
	runUnmarshalTest(t, "ipv4", `"192.168.1.1"`, netip.MustParseAddr("192.168.1.1"))
	runUnmarshalTest(t, "ipv6", `"::1"`, netip.MustParseAddr("::1"))
	runUnmarshalErrorTest[netip.Addr](t, "invalid", `"invalid"`)
}

func TestMarshal_NetipAddrPort(t *testing.T) {
	runMarshalTest(t, "ipv4", netip.MustParseAddrPort("192.168.1.1:8080"), `"192.168.1.1:8080"`)
	runMarshalTest(t, "ipv6", netip.MustParseAddrPort("[::1]:8080"), `"[::1]:8080"`)
}

func TestUnmarshal_NetipAddrPort(t *testing.T) {
	runUnmarshalTest(t, "ipv4", `"192.168.1.1:8080"`, netip.MustParseAddrPort("192.168.1.1:8080"))
	runUnmarshalTest(t, "ipv6", `"[::1]:8080"`, netip.MustParseAddrPort("[::1]:8080"))
}

func TestMarshal_NetipPrefix(t *testing.T) {
	runMarshalTest(t, "ipv4", netip.MustParsePrefix("192.168.1.0/24"), `"192.168.1.0/24"`)
	runMarshalTest(t, "ipv6", netip.MustParsePrefix("2001:db8::/32"), `"2001:db8::/32"`)
}

func TestUnmarshal_NetipPrefix(t *testing.T) {
	runUnmarshalTest(t, "ipv4", `"192.168.1.0/24"`, netip.MustParsePrefix("192.168.1.0/24"))
	runUnmarshalTest(t, "ipv6", `"2001:db8::/32"`, netip.MustParsePrefix("2001:db8::/32"))
}
