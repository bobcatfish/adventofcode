# http://adventofcode.com/2016/day/7
import collections
import re


IP = collections.namedtuple("IP",[ "others", "hypernets"])


def get_ips(filename):
    ips = []
    with open(filename) as f:
        for l in f.readlines():
            l = l.strip()
            others, hypernets = [], []
            for result in re.findall('([^\[\]]*)(\[[^\]]*\])?', l):
                if result[0]:
                    others.append(result[0])
                if result[1]:
                    hypernets.append(re.findall("\[(.*)\]", result[1])[0])
            ips.append(IP(others, hypernets))
    return ips


def is_abba(s):
    return s[0] == s[3] and s[0] != s[1] and s[1] == s[2]


def is_aba(s):
    return s[0] == s[2] and s[0] != s[1]


def contains_abba(s):
    for i in range(len(s) - 4 + 1):
        if is_abba(s[i:i+4]):
            return True
    return False


def is_tls_enabled(ip):
    if any(contains_abba(i) for i in ip.hypernets):
        return False
    return any(contains_abba(i) for i in ip.others)


def is_sls_enabled(ip):
    for o in ip.others:
        for i in range(len(o) - 3 + 1):
            s = o[i:i+3]
            if is_aba(s) and any("".join([s[1], s[0], s[1]]) in h for h in ip.hypernets):
                return True
    return False
    

if __name__ == "__main__":
    ips = get_ips("input")

    enabled = []
    for i, ip in enumerate(ips):
        #if is_tls_enabled(ip):
        if is_sls_enabled(ip):
            enabled.append(ip)

    print len(enabled)
