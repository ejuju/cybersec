# Cybersecurity toolkit for Gophers

## V0

Port scanning:
- [x] Default port scanner - perform basic port scanning (not stealthy)
- [ ] TCP SYN port scanner - perform stealthy TCP port scanning (https://github.com/google/gopacket/blob/v1.1.19/examples/synscan/main.go)

DOS attacks:
- [x] HTTP flood attack
- [ ] DNS amplification attack
- [ ] NTP amplification attack
- [ ] TCP SYN flood attack
- [ ] Slow loris attack

User agents:
- [x] Random user agents - for random HTTP user agents useful for DOS attacks

## Ideas for V1
- File encryption (fsutil)
- Steganography (fsutil)
- Botnet framework (netutil)
- Binary injection (fsutil)
- Remote shell execution (netutil)
- Password cracking (dictionary, timing and brute force attacks)
- Packet capture (https://github.com/google/gopacket/blob/master/examples/pcapdump/main.go)
- Packet replay (https://github.com/google/gopacket/blob/master/examples/pcaplay/main.go)
- Packet crafting (Packet crafting - spoof IP and mac addresses, craft fake packets)
