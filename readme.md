# Cybersecurity toolkit for Gophers

## V0

Recon:
- [x] Default port scanner - perform basic port scanning (not stealthy)
- [ ] TCP SYN port scanner - perform stealthy TCP port scanning (https://github.com/google/gopacket/blob/master/examples/synscan/main.go)

Network attacks:
- [x] HTTP flood attack
- [ ] Slow loris attack
- [ ] TCP SYN flood attack
- [ ] UDP flood attack
- [ ] NTP amplification attack
- [ ] DNS cache poisoning

Remote access:
- [x] SSH remote shell execution

Other utilities:
- [x] Sample HTTP user agents

## Ideas for V1
- Steganography (fsutil)
- Botnet framework (netutil)
- File encryption (fsutil)
- Binary injection (fsutil)
- Remote shell execution (netutil)
- Password cracking (dictionary, timing and brute force attacks)
- DNS amplification attack (https://github.com/say4n/dns.amplify)
- Packet capture (https://github.com/google/gopacket/blob/master/examples/pcapdump/main.go)
- Packet replay (https://github.com/google/gopacket/blob/master/examples/pcaplay/main.go)
- Packet crafting (Packet crafting - spoof IP and mac addresses, craft fake packets)
