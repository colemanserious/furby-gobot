package sounds

import (
	"golang.org/x/mobile/asset"
	"golang.org/x/mobile/exp/audio"
	"log"
)

// PlayFile plays a given wav file, fully pathed
func PlayFile(fileName string) (err error) {

	a, err := asset.Open(fileName)
	if err != nil {
		log.Println(err)
		return err
	}
	defer a.Close()

	player, err := audio.NewPlayer(a, 0, 0)
	if err != nil {
		log.Println(err)
		return err
	}
	player.Seek(0)
	player.Play()

	return

}
