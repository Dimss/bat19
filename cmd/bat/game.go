package main

import (
	"bytes"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/atomicgo/cursor"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"
)

var (
	sceneTemplate               = "assets/scene.tpl"
	batSpaces                   = 3
	batHeight                   = 7
	refreshRate   time.Duration = 700
	batsInScene                 = 4
)

type Bomb struct {
	Sprite    string
	LocationX int
	LocationY int
	Height    int
}

type Bat struct {
	Sprite   string
	Location int
	Height   int
}

type Scene struct {
	BatsInScene int
	RefreshRate time.Duration
}

func NewScene() *Scene {
	return &Scene{
		BatsInScene: batsInScene,
		RefreshRate: refreshRate,
	}
}

func NewBat(spriteLocation string) *Bat {
	b, err := os.ReadFile(spriteLocation)
	if err != nil {
		log.Fatal(err)
	}

	return &Bat{
		Sprite:   string(b),
		Location: rand.Intn(50) + 0,
		Height:   batHeight,
	}
}

func NewBomb() *Bomb {
	b, err := os.ReadFile("assets/bomb")
	if err != nil {
		log.Fatal(err)
	}

	return &Bomb{
		Sprite:    string(b),
		LocationX: rand.Intn(50) + 0,
		LocationY: rand.Intn(getSceneHeight()-6) + 0,
		Height:    5,
	}
}

func (b *Bomb) printBomb() {
	cursor.Up(b.LocationY)
	s := "{{ indent .LocationX .Sprite }}"
	t, buf := new(template.Template).Funcs(sprig.TxtFuncMap()).Funcs(customTemplateFuncs()), new(strings.Builder)
	template.Must(t.Parse(s)).Execute(buf, b)
	fmt.Printf("\033[31m%s\033[0m", buf.String())
}

func (s *Scene) getSceneHeight() int {
	return getSceneHeight()
}

func (s *Scene) generateBats(maxBats int) (bats []*Bat) {
	for i := 0; i < maxBats; i++ {
		if i%2 == 0 {
			bats = append(bats, NewBat("assets/bat1"))
		} else {
			bats = append(bats, NewBat("assets/bat2"))
		}
	}
	return
}

func (s *Scene) play() {
	fmt.Println(renderScene(sceneTemplate, s.generateBats(s.BatsInScene)))
	for {
		b := NewBomb()
		b.printBomb()
		fmt.Println()
		time.Sleep(1000 * time.Millisecond)
		cursor.Down(b.LocationY - 5)
		cursor.ClearLinesUp(s.getSceneHeight())
		cursor.ClearLine()
		scene := renderScene(sceneTemplate, s.generateBats(s.BatsInScene))
		fmt.Println(scene)
		time.Sleep(s.RefreshRate * time.Millisecond)
	}
}

func customTemplateFuncs() map[string]interface{} {
	return map[string]interface{}{
		"batSpace": func() string {
			var s []string
			for i := 0; i < batSpaces; i++ {
				s = append(s, "\n")
			}
			return strings.Join(s, "")
		},
		"color": func() string {
			rand.Seed(time.Now().UnixNano())
			colors := []string{"\u001B[32m", "\u001B[35m", "\u001B[33m", "\u001B[36m", "\033[97m", "\033[34m"}
			return colors[rand.Intn(len(colors)-0)+0]
		},
		"colorReset": func() string {
			return "\033[0m"
		},
	}
}

func renderScene(templateLocation string, bats []*Bat) string {
	t, err := os.ReadFile(templateLocation)
	if err != nil {
		log.Fatal(err)
	}
	var tpl bytes.Buffer
	tmpl, _ := template.New("t").
		Funcs(sprig.TxtFuncMap()).
		Funcs(customTemplateFuncs()).
		Parse(string(t))
	if err := tmpl.Execute(&tpl, bats); err != nil {
		log.Fatal(err)
	}
	return tpl.String()
}

func getSceneHeight() int {
	return batsInScene*batHeight + (batSpaces * (batsInScene - 1))
}
