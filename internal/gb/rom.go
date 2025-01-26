package gb

import (
	"fmt"
	"os"
)

var CartridgeTypes = map[byte]string{
	0x00: "ROM ONLY",
	0x01: "MBC1",
	0x02: "MBC1+RAM",
	0x03: "MBC1+RAM+BATTERY",
	0x05: "MBC2",
	0x06: "MBC2+BATTERY",
	0x08: "ROM+RAM 9",
	0x09: "ROM+RAM+BATTERY 9",
	0x0B: "MMM01",
	0x0C: "MMM01+RAM",
	0x0D: "MMM01+RAM+BATTERY",
	0x0F: "MBC3+TIMER+BATTERY",
	0x10: "MBC3+TIMER+RAM+BATTERY 10",
	0x11: "MBC3",
	0x12: "MBC3+RAM 10",
	0x13: "MBC3+RAM+BATTERY 10",
	0x19: "MBC5",
	0x1A: "MBC5+RAM",
	0x1B: "MBC5+RAM+BATTERY",
	0x1C: "MBC5+RUMBLE",
	0x1D: "MBC5+RUMBLE+RAM",
	0x1E: "MBC5+RUMBLE+RAM+BATTERY",
	0x20: "MBC6",
	0x22: "MBC7+SENSOR+RUMBLE+RAM+BATTERY",
	0xFC: "POCKET CAMERA",
	0xFD: "BANDAI TAMA5",
	0xFE: "HuC3",
	0xFF: "HuC1+RAM+BATTERY",
}

type ROMSize struct {
	Value            byte
	Size             string
	NumberOfROMBanks string
}

var ROMSizes = map[byte]ROMSize{
	0x00: {Value: 0x00, Size: "32 KiB", NumberOfROMBanks: "2 (no banking)"},
	0x01: {Value: 0x01, Size: "64 KiB", NumberOfROMBanks: "4"},
	0x02: {Value: 0x02, Size: "128 KiB", NumberOfROMBanks: "8"},
	0x03: {Value: 0x03, Size: "256 KiB", NumberOfROMBanks: "16"},
	0x04: {Value: 0x04, Size: "512 KiB", NumberOfROMBanks: "32"},
	0x05: {Value: 0x05, Size: "1 MiB", NumberOfROMBanks: "64"},
	0x06: {Value: 0x06, Size: "2 MiB", NumberOfROMBanks: "128"},
	0x07: {Value: 0x07, Size: "4 MiB", NumberOfROMBanks: "256"},
	0x08: {Value: 0x08, Size: "8 MiB", NumberOfROMBanks: "512"},
	0x52: {Value: 0x52, Size: "1.1 MiB", NumberOfROMBanks: "72 11"},
	0x53: {Value: 0x53, Size: "1.2 MiB", NumberOfROMBanks: "80 11"},
	0x54: {Value: 0x54, Size: "1.5 MiB", NumberOfROMBanks: "96 11"},
}

type RAMSize struct {
	Code     byte
	SRAMSize string
	Comment  string
}

var RAMSizes = map[byte]RAMSize{
	0x00: {Code: 0x00, SRAMSize: "0", Comment: "No RAM"},
	0x01: {Code: 0x01, SRAMSize: "-", Comment: "Unused"},
	0x02: {Code: 0x02, SRAMSize: "8 KiB", Comment: "1 bank"},
	0x03: {Code: 0x03, SRAMSize: "32 KiB", Comment: "4 banks of 8 KiB each"},
	0x04: {Code: 0x04, SRAMSize: "128 KiB", Comment: "16 banks of 8 KiB each"},
	0x05: {Code: 0x05, SRAMSize: "64 KiB", Comment: "8 banks of 8 KiB each"},
}

var DestinationCodes = map[byte]string{
	0x00: "Japan (and possibly overseas)",
	0x01: "Overseas only",
}

var NewLicenseeCodes = map[string]string{
	"00": "None",
	"01": "Nintendo Research & Development 1",
	"08": "Capcom",
	"13": "EA (Electronic Arts)",
	"18": "Hudson Soft",
	"19": "B-AI",
	"20": "KSS",
	"22": "Planning Office WADA",
	"24": "PCM Complete",
	"25": "San-X",
	"28": "Kemco",
	"29": "SETA Corporation",
	"30": "Viacom",
	"31": "Nintendo",
	"32": "Bandai",
	"33": "Ocean Software/Acclaim Entertainment",
	"34": "Konami",
	"35": "HectorSoft",
	"37": "Taito",
	"38": "Hudson Soft",
	"39": "Banpresto",
	"41": "Ubi Soft1",
	"42": "Atlus",
	"44": "Malibu Interactive",
	"46": "Angel",
	"47": "Bullet-Proof Software2",
	"49": "Irem",
	"50": "Absolute",
	"51": "Acclaim Entertainment",
	"52": "Activision",
	"53": "Sammy USA Corporation",
	"54": "Konami",
	"55": "Hi Tech Expressions",
	"56": "LJN",
	"57": "Matchbox",
	"58": "Mattel",
	"59": "Milton Bradley Company",
	"60": "Titus Interactive",
	"61": "Virgin Games Ltd.3",
	"64": "Lucasfilm Games4",
	"67": "Ocean Software",
	"69": "EA (Electronic Arts)",
	"70": "Infogrames5",
	"71": "Interplay Entertainment",
	"72": "Broderbund",
	"73": "Sculptured Software6",
	"75": "The Sales Curve Limited7",
	"78": "THQ",
	"79": "Accolade",
	"80": "Misawa Entertainment",
	"83": "lozc",
	"86": "Tokuma Shoten",
	"87": "Tsukuda Original",
	"91": "Chunsoft Co.8",
	"92": "Video System",
	"93": "Ocean Software/Acclaim Entertainment",
	"95": "Varie",
	"96": "Yonezawa/s'pal",
	"97": "Kaneko",
	"99": "Pack-In-Video",
	"9H": "Bottom Up",
	"A4": "Konami (Yu-Gi-Oh!)",
	"BL": "MTO",
	"DK": "Kodansha",
}

var OldLicenseeCodes = map[byte]string{
	0x00: "None",
	0x01: "Nintendo",
	0x08: "Capcom",
	0x09: "HOT-B",
	0x0A: "Jaleco",
	0x0B: "Coconuts Japan",
	0x0C: "Elite Systems",
	0x13: "EA (Electronic Arts)",
	0x18: "Hudson Soft",
	0x19: "ITC Entertainment",
	0x1A: "Yanoman",
	0x1D: "Japan Clary",
	0x1F: "Virgin Games Ltd.3",
	0x24: "PCM Complete",
	0x25: "San-X",
	0x28: "Kemco",
	0x29: "SETA Corporation",
	0x30: "Infogrames5",
	0x31: "Nintendo",
	0x32: "Bandai",
	0x33: "New licensee code should be used instead.",
	0x34: "Konami",
	0x35: "HectorSoft",
	0x38: "Capcom",
	0x39: "Banpresto",
	0x3C: "Entertainment Interactive (stub)",
	0x3E: "Gremlin",
	0x41: "Ubi Soft1",
	0x42: "Atlus",
	0x44: "Malibu Interactive",
	0x46: "Angel",
	0x47: "Spectrum HoloByte",
	0x49: "Irem",
	0x4A: "Virgin Games Ltd.3",
	0x4D: "Malibu Interactive",
	0x4F: "U.S. Gold",
	0x50: "Absolute",
	0x51: "Acclaim Entertainment",
	0x52: "Activision",
	0x53: "Sammy USA Corporation",
	0x54: "GameTek",
	0x55: "Park Place13",
	0x56: "LJN",
	0x57: "Matchbox",
	0x59: "Milton Bradley Company",
	0x5A: "Mindscape",
	0x5B: "Romstar",
	0x5C: "Naxat Soft14",
	0x5D: "Tradewest",
	0x60: "Titus Interactive",
	0x61: "Virgin Games Ltd.3",
	0x67: "Ocean Software",
	0x69: "EA (Electronic Arts)",
	0x6E: "Elite Systems",
	0x6F: "Electro Brain",
	0x70: "Infogrames5",
	0x71: "Interplay Entertainment",
	0x72: "Broderbund",
	0x73: "Sculptured Software6",
	0x75: "The Sales Curve Limited7",
	0x78: "THQ",
	0x79: "Accolade15",
	0x7A: "Triffix Entertainment",
	0x7C: "MicroProse",
	0x7F: "Kemco",
	0x80: "Misawa Entertainment",
	0x83: "LOZC G.",
	0x86: "Tokuma Shoten",
	0x8B: "Bullet-Proof Software2",
	0x8C: "Vic Tokai Corp.16",
	0x8E: "Ape Inc.17",
	0x8F: "I’Max18",
	0x91: "Chunsoft Co.8",
	0x92: "Video System",
	0x93: "Tsubaraya Productions",
	0x95: "Varie",
	0x96: "Yonezawa19/S’Pal",
	0x97: "Kemco",
	0x99: "Arc",
	0x9A: "Nihon Bussan",
	0x9B: "Tecmo",
	0x9C: "Imagineer",
	0x9D: "Banpresto",
	0x9F: "Nova",
	0xA1: "Hori Electric",
	0xA2: "Bandai",
	0xA4: "Konami",
	0xA6: "Kawada",
	0xA7: "Takara",
	0xA9: "Technos Japan",
	0xAA: "Broderbund",
	0xAC: "Toei Animation",
	0xAD: "Toho",
	0xAF: "Namco",
	0xB0: "Acclaim Entertainment",
	0xB1: "ASCII Corporation or Nexsoft",
	0xB2: "Bandai",
	0xB4: "Square Enix",
	0xB6: "HAL Laboratory",
	0xB7: "SNK",
	0xB9: "Pony Canyon",
	0xBA: "Culture Brain",
	0xBB: "Sunsoft",
	0xBD: "Sony Imagesoft",
	0xBF: "Sammy Corporation",
	0xC0: "Taito",
	0xC2: "Kemco",
	0xC3: "Square",
	0xC4: "Tokuma Shoten",
	0xC5: "Data East",
	0xC6: "Tonkin House",
	0xC8: "Koei",
	0xC9: "UFL",
	0xCA: "Ultra Games",
	0xCB: "VAP, Inc.",
	0xCC: "Use Corporation",
	0xCD: "Meldac",
	0xCE: "Pony Canyon",
	0xCF: "Angel",
	0xD0: "Taito",
	0xD1: "SOFEL (Software Engineering Lab)",
	0xD2: "Quest",
	0xD3: "Sigma Enterprises",
	0xD4: "ASK Kodansha Co.",
	0xD6: "Naxat Soft14",
	0xD7: "Copya System",
	0xD9: "Banpresto",
	0xDA: "Tomy",
	0xDB: "LJN",
	0xDD: "Nippon Computer Systems",
	0xDE: "Human Ent.",
	0xDF: "Altron",
	0xE0: "Jaleco",
	0xE1: "Towa Chiki",
	0xE2: "Yutaka # Needs more info",
	0xE3: "Varie",
	0xE5: "Epoch",
	0xE7: "Athena",
	0xE8: "Asmik Ace Entertainment",
	0xE9: "Natsume",
	0xEA: "King Records",
	0xEB: "Atlus",
	0xEC: "Epic/Sony Records",
	0xEE: "IGS",
	0xF0: "A Wave",
	0xF3: "Extreme Entertainment",
	0xFF: "LJN",
}

type Header struct {
	Checksum         byte
	Title            []byte
	Logo             []byte
	ManufacturerCode []byte
	NewLicenseeCode  string
	OldLicenseeCode  byte
	CGBFlag          byte
	SGBFlag          byte
	CartridgeType    byte
	ROMSize          byte
	RAMSize          byte
	DestinationCode  byte
	// GlobalChecksum       []byte
	// MaskROMVersionNumber byte
}

type ROM struct {
	Header Header
	data   []byte
}

func (r *ROM) Load(rom_data []byte) {
	checksum := rom_data[0x014D]
	checksum_acc := 0
	for address := 0x0134; address <= 0x014C; address++ {
		checksum_acc = checksum_acc - int(rom_data[address]) - 1
	}

	if int(checksum) != checksum_acc&0xFF {
		fmt.Println("(ROM) error: checksum missmatch")
		os.Exit(1)
	}

	r.Header.Checksum = checksum
	r.Header.Title = rom_data[0x0134:0x0143]
	r.Header.Logo = rom_data[0x0104:0x0133]
	r.Header.ManufacturerCode = rom_data[0x013F:0x0142]
	r.Header.NewLicenseeCode = string(rom_data[0x0144]) + string(rom_data[0x0145])
	r.Header.OldLicenseeCode = rom_data[0x014B]
	r.Header.CGBFlag = rom_data[0x0143]
	r.Header.SGBFlag = rom_data[0x0146]
	r.Header.CartridgeType = rom_data[0x0147]
	r.Header.ROMSize = rom_data[0x0148]
	r.Header.RAMSize = rom_data[0x0149]
	r.Header.DestinationCode = rom_data[0x014A]
	// r.Header.MaskROMVersionNumber = rom_data[0x014C]
	// r.Header.GlobalChecksum = rom_data[0x014E:0x014F]

	r.data = rom_data
}

func (r *ROM) Read(address uint16) byte {
	return r.data[address]
}

func (r *ROM) Write(address uint16, value byte) {
}

func (r *ROM) String() string {
	CartridgeType := CartridgeTypes[r.Header.CartridgeType]
	ROMSize := ROMSizes[r.Header.ROMSize]
	RAMSize := RAMSizes[r.Header.RAMSize]
	DestinationCode := DestinationCodes[r.Header.DestinationCode]
	NewLicenseeCode := NewLicenseeCodes[r.Header.NewLicenseeCode]
	OldLicenseeCode := OldLicenseeCodes[r.Header.OldLicenseeCode]

	return fmt.Sprintf(
		`
-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

Checksum             : %2.2X
Title                : %s
ManufacturerCode     : %2.2X
NewLicenseeCode      : %s
OldLicenseeCode      : %s
CGBFlag              : %2.2X
SGBFlag              : %2.2X
CartridgeType        : %s
ROMSize              : %s
RAMSize              : %s # %s
DestinationCode      : %s

-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-_-

`,
		r.Header.Checksum,
		string(r.Header.Title),
		r.Header.ManufacturerCode,
		NewLicenseeCode,
		OldLicenseeCode,
		r.Header.CGBFlag,
		r.Header.SGBFlag,
		CartridgeType,
		ROMSize.Size,
		RAMSize.SRAMSize,
		RAMSize.Comment,
		DestinationCode,
	)
}
