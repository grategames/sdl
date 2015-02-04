// Copyright 2012 The go-sdl2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package mixer provies access to the SDL2_mixer library.
package mixer

// #cgo pkg-config: sdl2
// #cgo LDFLAGS: -lSDL2_mixer 
// #include "SDL2/SDL_mixer.h"
import "C"

import (
	"grate/backend/sdl2"
	"errors"
	"io"
	"io/ioutil"
	"unsafe"
)

// sdlError creates a new sdl.SDLError.
func sdlError(i int) sdl.SDLError {
	return sdl.SDLError{sdl.GetError(), i}
}

const (
	MAJOR_VERSION = C.SDL_MIXER_MAJOR_VERSION
	MINOR_VERSION = C.SDL_MIXER_MINOR_VERSION
	PATCHLEVEL    = C.SDL_MIXER_PATCHLEVEL
)

// VERSION fills v with the compile-time version of the SDL2_mixer library.
func VERSION(v *sdl.Version) {
	v.Major = MAJOR_VERSION
	v.Minor = MINOR_VERSION
	v.Patch = PATCHLEVEL
}

// LinkedVersion gets the version of the dynamically linked SDL2_mixer library.
// It should not be used to fill a version structure, instead you should use
// VERSION.
func LinkedVersion() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.Mix_Linked_Version()))
}

type InitFlags int32

const (
	INIT_FLAC       InitFlags = C.MIX_INIT_FLAC
	INIT_MOD        InitFlags = C.MIX_INIT_MOD
	INIT_MP3        InitFlags = C.MIX_INIT_MP3
	INIT_OGG        InitFlags = C.MIX_INIT_OGG
	INIT_FLUIDSYNTH InitFlags = C.MIX_INIT_FLUIDSYNTH
)

// Init loads dynamic libraries and prepares them for use.  Flags should
// be one or more flags from InitFlags OR'd together.  It returns an error
// if a library fails to load.
func Init(flags InitFlags) error {
	r := C.Mix_Init(C.int(flags))
	if r != C.int(flags) {
		return sdlError(int(r))
	}
	return nil
}

// Quit unloads libraries loaded with Initialize.
func Quit() {
	C.Mix_Quit()
}

type Fading int

const (
	NO_FADING  Fading = C.MIX_NO_FADING
	FADING_OUT Fading = C.MIX_FADING_OUT
	FADING_IN  Fading = C.MIX_FADING_IN
)

type AudioFormat uint16

const (
	AUDIO_U8     AudioFormat = C.AUDIO_U8     // Unsigned 8-bit samples
	AUDIO_S8     AudioFormat = C.AUDIO_S8     // Signed 8-bit samples
	AUDIO_U16LSB AudioFormat = C.AUDIO_U16LSB // Unsigned 16-bit samples
	AUDIO_S16LSB AudioFormat = C.AUDIO_S16LSB // Signed 16-bit samples
	AUDIO_U16MSB AudioFormat = C.AUDIO_U16MSB // As above, but big-endian byte order
	AUDIO_S16MSB AudioFormat = C.AUDIO_S16MSB // As above, but big-endian byte order
	AUDIO_U16    AudioFormat = C.AUDIO_U16
	AUDIO_S16    AudioFormat = C.AUDIO_S16

	// Native audio byte ordering
	AUDIO_U16SYS AudioFormat = C.AUDIO_U16SYS
	AUDIO_S16SYS AudioFormat = C.AUDIO_S16SYS // Same as MIX_DEFAULT_FORMAT
)

// The internal format for an audio chunk.
type Chunk struct {
	ptr *C.Mix_Chunk
}

type MusicType int32

const (
	MUS_NONE    MusicType = C.MUS_NONE
	MUS_CMD     MusicType = C.MUS_CMD
	MUS_WAV     MusicType = C.MUS_WAV
	MUS_MOD     MusicType = C.MUS_MOD
	MUS_MID     MusicType = C.MUS_MID
	MUS_OGG     MusicType = C.MUS_OGG
	MUS_MP3     MusicType = C.MUS_MP3
	MUS_MP3_MAD MusicType = C.MUS_MP3_MAD
	MUS_FLAC    MusicType = C.MUS_FLAC
	MUS_MODPLUG MusicType = C.MUS_MODPLUG
)

// The internal format for a music chunk.
type Music struct {
	ptr *C.Mix_Music
}

// OpenAudio opens the mixer with a certain audio format.  frequency is the
// output sampling frequency in samples per second (Hz).  format is the
// output sample format.  channels is the number of sound channels in output,
// it has nothing to do with mixing channels.  chunksize is the size of each
// mixer sample.  If it is too small the sound may skip on slow systems, if it
// is too large the sound effects will lag behind the action.  OpenAudio also
// allocates 8 mixing channels.
//
// SDL must be initialized with INIT_AUDIO before calling OpenAudio.
//
// Common settings are:
//  frequency: 44100 or 22050
//  format: AUDIO_S16SYS
//  channels: 2 (stereo) or 1 (mono)
//  chunksize: 512, 1024, 2048, or 4096
func OpenAudio(frequency int, format AudioFormat, channels, chunksize int) error {
	r := C.Mix_OpenAudio(C.int(frequency), C.Uint16(format),
		C.int(channels), C.int(chunksize))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// AllocateChannels sets the number of channels being mixed.  This can be
// called multiple times, even with sounds playing.  If numchans is less than
// the current number of channels, then the higher channels will be stopped,
// freed, and therefore not mixed any longer.  It returns the number of
// channels allocated.
//
// Note: If numchans is zero, AllocateChannels will free all mixing channels,
// however music will still play.
func AllocateChannels(numchans int) int {
	return int(C.Mix_AllocateChannels(C.int(numchans)))
}

// QuerySpec returns what the actual audio device parameters are.  opened is
// the number of times the device was opened or 0 on error.
func QuerySpec() (format AudioFormat, frequency, channels, opened int) {
	opened = int(C.Mix_QuerySpec((*C.int)(unsafe.Pointer(&frequency)),
		(*C.Uint16)(&format),
		(*C.int)(unsafe.Pointer(&channels))))
	return
}

// LoadWAV loads a file into a Chunk.
func LoadWAV(file string) (Chunk, error) {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))
	mode := C.CString("rb")
	defer C.free(unsafe.Pointer(mode))

	r := C.Mix_LoadWAV_RW(C.SDL_RWFromFile(cstr, mode), 1)
	if r == nil {
		return Chunk{}, sdlError(0)
	}
	return Chunk{r}, nil
}

// LoadWAVFromReader loads an io.Reader into a Chunk.
func LoadWAVFromReader(reader io.Reader) (Chunk, error) {
	buff, err := ioutil.ReadAll(reader)
	if err != nil {
		return Chunk{}, err
	}
	if len(buff) == 0 {
		return Chunk{}, errors.New("io.Reader is empty, no chunk created.")
	}
	r := C.Mix_LoadWAV_RW(C.SDL_RWFromMem(unsafe.Pointer(&buff[0]), C.int(len(buff))), 1)
	if r == nil {
		return Chunk{}, sdlError(0)
	}
	return Chunk{r}, nil
}

// LoadMUS loads a file into a Music.
func LoadMUS(file string) (Music, error) {
	cstr := C.CString(file)
	defer C.free(unsafe.Pointer(cstr))

	r := C.Mix_LoadMUS(cstr)
	if r == nil {
		return Music{}, sdlError(0)
	}
	return Music{r}, nil
}

// LoadMUSFromReader loads an io.Reader into a Music.
func LoadMUSFromReader(reader io.Reader) (Music, error) {
	buff, err := ioutil.ReadAll(reader)
	if err != nil {
		return Music{}, err
	}
	if len(buff) == 0 {
		return Music{}, errors.New("io.Reader is empty, no music created.")
	}
	r := C.Mix_LoadMUS_RW(C.SDL_RWFromMem(unsafe.Pointer(&buff[0]), C.int(len(buff))), C.int(0))
	if r == nil {
		return Music{}, sdlError(0)
	}
	return Music{r}, nil
}

// Free frees c.
func (c Chunk) Free() {
	C.Mix_FreeChunk(c.ptr)
	c.ptr = nil
}

// Free frees m.
func (m Music) Free() {
	C.Mix_FreeMusic(m.ptr)
	m.ptr = nil
}

// GetNumChunkDecoders gets the number of chunk decoders mixer provides.  You
// must successfully call OpenAudio before calling the function.
func GetNumChunkDecoders() int {
	return int(C.Mix_GetNumChunkDecoders())
}

// GetChunkDecoder gets the name of the chunk decoder at index.  You must
// successfully call OpenAudio before calling the function.
func GetChunkDecoder(index int) string {
	return C.GoString(C.Mix_GetChunkDecoder(C.int(index)))
}

// GetNumMusicDecoders gets the number of music decoders mixer provides.  You
// must successfully call OpenAudio before calling the function.
func GetNumMusicDecoders() int {
	return int(C.Mix_GetNumMusicDecoders())
}

// GetMusicDecoder gets the name of the music decoder at index.  You must
// successfully call OpenAudio before calling the function.
func GetMusicDecoder(index int) string {
	return C.GoString(C.Mix_GetMusicDecoder(C.int(index)))
}

// GetType gets the music format of m, or the currently playing music, if m is
// a zero value Music.
func (m Music) GetType() MusicType {
	return MusicType(C.Mix_GetMusicType(m.ptr))
}

// ReserveChannels reserves the first channels (0 -> n-1) for the application,
// i.e. don't allocate them dynamically to the next sample if requested with a
// -1 value below.  Returns the number of reserved channels.
func ReserveChannels(num int) int {
	return int(C.Mix_ReserveChannels(C.int(num)))
}

// PlayChannel plays an audio chunk on a specific channel.  If the specified
// channel is -1, play on the first free channel. If loops is greater then
// zero, loop the sound that many times. If loops is -1, loop inifinitely.
// Returns which channel was used to play the sound.
func PlayChannel(channel int, chunk Chunk, loops int) (int, error) {
	r := int(C.Mix_PlayChannelTimed(C.int(channel), chunk.ptr,
		C.int(loops), -1))
	if r == -1 {
		return r, sdlError(r)
	}
	return r, nil
}

// PlayChannelTimed is the same as PlayChannel, but the sound is played at
// most ticks milliseconds.
func PlayChannelTimed(channel int, chunk Chunk, loops, ticks int) (int, error) {
	r := int(C.Mix_PlayChannelTimed(C.int(channel), chunk.ptr,
		C.int(loops), C.int(ticks)))
	if r == -1 {
		return r, sdlError(r)
	}
	return r, nil
}

// Play plays m loops number of times.  If loops is -1 m will loop forever.
// The previous is music is halted, or if fading out it waits (blocking) for
// that to finish.
func (m Music) Play(loops int) error {
	r := C.Mix_PlayMusic(m.ptr, C.int(loops))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// FadeIn is the same as Play, but m fades in over ms milliseconds.
func (m Music) FadeIn(loops, ms int) error {
	r := C.Mix_FadeInMusic(m.ptr, C.int(loops), C.int(ms))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// FadeInPos is the same as FadeIn, but the music will be started at position.
// position has different meanings for different types of music files, see
// SetMusicPosition for more information.
func (m Music) FadeInPos(loops, ms int, position float64) error {
	r := C.Mix_FadeInMusicPos(m.ptr, C.int(loops), C.int(ms),
		C.double(position))
	if r != 0 {
		return sdlError(int(r))
	}
	return nil
}

// Same as PlayChannel, but the chunk fades in over ms milliseconds.
func FadeInChannel(channel int, chunk Chunk, loops, ms int) (int, error) {
	r := int(C.Mix_FadeInChannelTimed(C.int(channel), chunk.ptr,
		C.int(loops), C.int(ms), -1))
	if r == -1 {
		return r, sdlError(r)
	}
	return r, nil
}

// Same as PlayChannelTimed, but the chunk fades in over ms milliseconds.
func FadeInChannelTimed(channel int, chunk Chunk, loops, ms, ticks int) (int, error) {
	r := int(C.Mix_FadeInChannelTimed(C.int(channel), chunk.ptr,
		C.int(loops), C.int(ms), C.int(ticks)))
	if r == -1 {
		return r, sdlError(r)
	}
	return r, nil
}

// Volume sets the volume for any allocated channel.  If channel is -1 then all
// channels are set at once.  The volume is applied during the final mix,
// along with the sample volume.  So setting this volume to 64 will halve the
// output of all samples played on the specified channel.  All channels
// default to a volume of 128, which is the max.  Newly allocated channels will
// have the max volume set, so setting all channels volumes does not affect
// subsequent channel allocations.
//
// If volume is less then 0 then the volume will not be set.
//
// Volume returns the current volume of the channel.  If channel is -1, the
// average volume is returned.
func Volume(channel, volume int) int {
	return int(C.Mix_Volume(C.int(channel), C.int(volume)))
}

// Volume sets the volume for c.  The volume setting will take effect when
// the chunk is used on a channel, be mixed into the output.
//
// If volume is less then 0 then the volume will not be set.
//
// Volume returns the previous volume of c.
func (c Chunk) Volume(volume int) int {
	return int(C.Mix_VolumeChunk(c.ptr, C.int(volume)))
}

// VolumeMusic sets the volume for music. Setting the volume during a fade will
// not work, the faders use this function to perform their effect!  Setting
// volume while using an external music player set by SetMusicCMD will have
// no effect.
//
// If volume is less then 0 then the volume will not be set.
//
// VolumeMusic returns the previous volume setting
func VolumeMusic(volume int) int {
	return int(C.Mix_VolumeMusic(C.int(volume)))
}

// HaltChannel halts the playback on channel, or all channels if channel is -1.
func HaltChannel(channel int) {
	C.Mix_HaltChannel(C.int(channel))
}

// HaltMusic halts the playback of music.  This interrupts music fader effects.
func HaltMusic() {
	C.Mix_HaltMusic()
}

// ExpireChannel halts playback on channel, or all channels if channel is -1,
// after ticks milliseconds.
//
// ExpireChannel returns the number of channels set to expire.  Whether or not
// they are active.
func ExpireChannel(channel, ticks int) int {
	return int(C.Mix_ExpireChannel(C.int(channel), C.int(ticks)))
}

// FadeOutChannel gradually fades out playback on channel which, or all
// channels if which is -1, over ms milliseconds starting from now.  The
// channel(s) will be halted after the fade out is completed.  Only channels
// that are playing are set to fade out, including paused channels.
//
// FadeOutChannel returns the number of channel set to fade out.
func FadeOutChannel(which, ms int) int {
	return int(C.Mix_FadeOutChannel(C.int(which), C.int(ms)))
}

// FadeOutMusic gradually fades out music over ms milliseconds starting from
// now.  The music will be halted after the fade is completed.
func FadeOutMusic(ms int) error {
	r := int(C.Mix_FadeOutMusic(C.int(ms)))
	if r == 0 {
		return sdlError(r)
	}
	return nil
}

// FadingMusic tells you if the music is fading in, out, or not at all.  It
// does not tell you if the channel is playing anything, or paused.
func FadingMusic() Fading {
	return Fading(C.Mix_FadingMusic())
}

// FadingChannel tells you if which channel is fading in, out, or not at all.
// It does not tell you if the channel is playing anything, or paused.
//
// Note which can not be -1.
func FadingChannel(which int) Fading {
	return Fading(C.Mix_FadingChannel(C.int(which)))
}

// Pause pause playback on channel, or all channels if channel is -1.  You may
// still halt a paused channel.
//
// Note: Only channels which are actively playing will be paused.
func Pause(channel int) {
	C.Mix_Pause(C.int(channel))
}

// Resume unpauses channel, or all playing and paused channels if channel is
// -1.
func Resume(channel int) {
	C.Mix_Resume(C.int(channel))
}

// Paused tells you if channel is paused, or not.  If channel is -1 it returns
// the number of paused channels, otherwise it returns 0 if channel is not
// paused or 1 if it is paused.
func Paused(channel int) int {
	return int(C.Mix_Paused(C.int(channel)))
}

// PauseMusic pauses the music playback.  You may halt paused music.
//
// Note: Music can only be paused if it is actively playing.
func PauseMusic() {
	C.Mix_PauseMusic()
}

// ResumeMusics unpauses the music.  This is safe to use on halted, paused,
// and already playing music.
func ResumeMusic() {
	C.Mix_ResumeMusic()
}

// RewindMusic rewinds the music to the start. This is safe to use on halted,
// paused, and already playing music.  It is not useful to rewind the music
// Immediately after starting playback, because it starts at the beginning by
// default.
//
// RewindwMusic does not work for all music formats.
func RewindMusic() {
	C.Mix_RewindMusic()
}

// SetMusicPosition sets the postion of the currently playing music.  The
// position takes different meaning for different music sources.  It only
// works on the music types listed below.
//
// Position meanings:
//  MOD:
//    position is cast to a uint16 and used for a pattern number in the
//    module.  Passing zero is similar to rewinding the song.
//
//  OGG, FLAC, MP3_MAD, and MODPLUG:
//    Jumps to the position seconds from the beginning of the song.
//
//  MP3:
//    Jumps to the position seconds from the current position in the stream.
//    So you may want to call RewindMusic before this.  Does not go in
//    Reverse.
func SetMusicPosition(position float64) int {
	return int(C.Mix_SetMusicPosition(C.double(position)))
}

// PausedMusic tells you if music is paused, or not.
func PausedMusic() bool {
	r := C.Mix_PausedMusic()
	if r == 1 {
		return true
	}
	return false
}

// Playing tells you if channel is playing, or not.  It does not check if the
// channel has been paused.
//
// Playing returns zero if the channel is not playing, or 1 if it is playing.
// If channel is -1, it returns the number of channels that are playing.
func Playing(channel int) int {
	return int(C.Mix_Playing(C.int(channel)))
}

// PlayingMusic tells you if music is actively playing, or not.  It does not
// check if the channel has been paused.
func PlayingMusic() bool {
	r := C.Mix_PlayingMusic()
	if r == 1 {
		return true
	}
	return false
}

// SetMusicCMD sets the system command that is used to play music.  The command
// should be a complete command, as if typed in to the command line, but it
// should expect the filename to be added as teh last argument.  Set command to
// and empty string to turn off using an external command for music.
//
// When SetMusicCMD is called any music playing is halted.  The music file to
// play is set by calling LoadMUS(filename), and the filename is appended as
// the last argument on the commandline.  The command will be sent signals
// SIGTERM to halt, SIGSTOP to pause, and SIGCONT to resume.  The command
// program should react correctly to those signals for it to function properly
// with SDL2_Mixer.  Volume should be set in the music player's command if the
// music player supports that.
//
// Notes: Playing music through an external command may not work if the sound
// driver does not support multiple openings of the audio device.  Also
// commands are not totally portable, so be careful.
func SetMusicCMD(command string) int {
	cstr := C.CString(command)
	defer C.free(unsafe.Pointer(cstr))

	return int(C.Mix_SetMusicCMD(cstr))
}

func SetSoundFonts(paths string) int {
	cstr := C.CString(paths)
	defer C.free(unsafe.Pointer(cstr))

	return int(C.Mix_SetSoundFonts(cstr))
}

func GetSoundFonts() string {
	return C.GoString(C.Mix_GetSoundFonts())
}

// CloseAudio shuts down and cleans up the mixer API.  After calling this all
// audio is stopped, the device is closed, and the mixer functions should not
// be used.  You may, of course use OpenAudio to start the functionality again.
//
// Note: This function does not do anything until you have called it the same
// number of times that you called OpwnAudio.  You may use QuerySpec to find
// out how many times CloseAudio needs to be called before the device is
// actually closed.
func CloseAudio() {
	C.Mix_CloseAudio()
}
