package ffmpeg

type Preset string

const (
	PresetUltraFast Preset = "ultrafast"
	PresetSuperFast Preset = "superfast"
	PresetVeryFast  Preset = "veryfast"
	PresetFaster    Preset = "faster"
	PresetFast      Preset = "fast"
	PresetMedium    Preset = "medium"
	PresetSlow      Preset = "slow"
	PresetSlower    Preset = "slower"
	PresetVerySlow  Preset = "veryslow"
	PresetPlacebo   Preset = "placebo"
)

var presets = []Preset{
	PresetUltraFast, PresetSuperFast, PresetVeryFast,
	PresetFaster, PresetFast, PresetMedium, PresetSlow,
	PresetSlower, PresetVerySlow, PresetPlacebo,
}

func FormatPreset(p Preset) string {
	for _, preset := range presets {
		if preset == p {
			return string(p)
		}
	}
	return string(PresetMedium)
}
