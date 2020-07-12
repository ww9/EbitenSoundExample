package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/audio"
	"github.com/hajimehoshi/ebiten/audio/wav"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Create the audio context that is used to play all sounds
var AudioContext = func() *audio.Context {
	c, err := audio.NewContext(48000)
	if err != nil {
		panic(err)
	}
	return c
}()

func main() {
	ebiten.SetWindowSize(400, 400)
	ebiten.SetWindowTitle("Ebiten Sound Player Example")
	game := &Game{}
	game.Sound1 = NewWavPlayer(Bsound_wav)
	game.Sound2 = NewWavPlayer(Bsound2_wav) // <--- errors with "bits per sample must be 8 or 16 but was 24". Change Bsound2_wav to Bsound_wav to run the program
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func NewWavPlayer(wavBytes []byte) *audio.Player {
	stream, err := wav.Decode(AudioContext, audio.BytesReadSeekCloser(wavBytes))
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(stream)
	if err != nil {
		log.Fatal(err)
	}
	player, _ := audio.NewPlayerFromBytes(AudioContext, b)
	return player
}

type Game struct {
	Sound1 *audio.Player
	Sound2 *audio.Player
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 400
}

func (g *Game) Update(*ebiten.Image) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, _ := ebiten.CursorPosition()
		if x <= 200 && !g.Sound1.IsPlaying() {
			g.Sound1.Rewind()
			g.Sound1.Play()
		} else if !g.Sound2.IsPlaying() {
			g.Sound2.Rewind()
			g.Sound2.Play()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawRect(screen, 0, 0, 200, 400, color.RGBA{255, 0, 0, 255})
	ebitenutil.DrawRect(screen, 200, 0, 400, 400, color.RGBA{0, 255, 0, 255})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f (Click to play sound)", ebiten.CurrentTPS()))
}
