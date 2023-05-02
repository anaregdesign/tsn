package main

import (
	"context"
	"encoding/json"
	"github.com/anaregdesign/tsn"
	"strings"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	c := tsn.NewTokenCache(ctx, 10)

	paragraph := `
And so even though we face the difficulties of today and tomorrow, I still have a dream. It is a dream deeply rooted in the American dream.
I have a dream that one day this nation will rise up and live out the true meaning of its creed：“We hold these truths to be self-evident, that all men are created equal.”
I have a dream that one day on the red hills of Georgia, the sons of former slaves and the sons of former slave owners will be able to sit down together at the table of brotherhood.
I have a dream that one day even the state of Mississippi, a state sweltering with the heat of injustice, sweltering with the heat of oppression, will be transformed into an oasis of freedom and justice.
I have a dream that my four little children will one day live in a nation where they will not be judged by the color of their skin but by the content of their character.
I have a dream today!
I have a dream that one day, down in Alabama, with its vicious racists, with its governor having his lips dripping with the words of “interposition” and “nullification” -- one day right there in Alabama little black boys and black girls will be able to join hands with little white boys and white girls as sisters and brothers.
I have a dream today!
I have a dream that one day every valley shall be exalted, and every hill and mountain shall be made low, the rough places will be made plain, and the crooked places will be made straight; “and the glory of the Lord shall be revealed and all flesh shall see it together.”
`
	paragraph = strings.Replace(paragraph, "\n", " ", -1)
	paragraph = strings.Replace(paragraph, ".", " ", -1)
	paragraph = strings.Replace(paragraph, ",", " ", -1)
	c.Add(strings.Split(paragraph, " "))

	g := c.Neighbor("slaves", 1, 5)
	if jsonString, err := json.MarshalIndent(g, "", "\t"); err == nil {
		println(string(jsonString))
	}

	cancel()
}
